package postgres

import (
	"database/sql"
	"encoding/json"

	log "github.com/Hirogava/ServiceBuyer/internal/config/logger"
	errors "github.com/Hirogava/ServiceBuyer/internal/errors/db"
	dbModel "github.com/Hirogava/ServiceBuyer/internal/model/db"
	model "github.com/Hirogava/ServiceBuyer/internal/model/request"
	dbService "github.com/Hirogava/ServiceBuyer/internal/service/db"
	"github.com/lib/pq"
)

func (manager *Manager) CreateServiceRequest(request *model.ServiceRequest) error {
	err := dbService.ParseRequest(request)
	if err != nil {
		log.Logger.Error("Error parse request: ", err)
		return err
	}

	tx, err := manager.Conn.Begin()
	if err != nil {
		log.Logger.Error("Error begin transaction: ", err)
		return errors.ErrTxNotStarted
	}
	defer func() error {
		err := tx.Rollback()
		if err != nil {
			log.Logger.Error("Error rollback transaction: ", err)
			return errors.ErrTxNotRolledBack
		}

		return nil
	}()

	var userExists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM \"user\" WHERE id = $1)", request.UserID).Scan(&userExists)
	if err != nil {
		log.Logger.Error("Failed to check user existence", "user_id", request.UserID, "error", err)
		return errors.ErrUserNotFound
	}
	if !userExists {
		log.Logger.Error("User not found", "user_id", request.UserID)
		return errors.ErrUserNotFound
	}

	var serviceID int
	err = tx.QueryRow("SELECT id FROM service WHERE name = $1", request.Name).Scan(&serviceID)
	if err == sql.ErrNoRows {
		log.Logger.Info("Creating new service", "name", request.Name)
		err = tx.QueryRow(
			"INSERT INTO service (name, amount) VALUES ($1, $2) RETURNING id",
			request.Name, request.Cost,
		).Scan(&serviceID)
		if err != nil {
			log.Logger.Error("Failed to create service", "name", request.Name, "error", err)
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				return errors.ErrServiceAlreadyExists
			}
			return err
		}
	} else if err != nil {
		log.Logger.Error("Failed to check service existence", "name", request.Name, "error", err)
		return err
	}

	_, err = manager.Conn.Exec(`
		INSERT INTO user_purchase (user_id, service_id, end_date) 
		VALUES ($1, $2, $3)
		`, request.UserID, serviceID, request.EndDate)
	if err != nil {
		log.Logger.Error("Error insert user purchase: ", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		log.Logger.Error("Error commit transaction: ", err)
		return errors.ErrTxNotCommitted
	}
	
	return nil
}

func (manager *Manager) CountingServiceRequest(req *model.CountingRequest) (*dbModel.CountingResponse, error) {
	err := dbService.ParseCountingRequest(req)
	if err != nil {
		log.Logger.Error("Error parse request: ", err)
		return nil, err
	}

	query := `
		SELECT 
			'' AS user_id,
			COALESCE(SUM(s.amount), 0) AS total_amount,
			ARRAY_AGG(
				DISTINCT JSON_BUILD_OBJECT(
					'id', s.id,
					'name', s.name,
					'amount', s.amount
				)
			) AS services,
			$1::date AS start_date,
			$2::date AS end_date
		FROM user_purchase up
		JOIN service s ON up.service_id = s.id
		WHERE 
			up.created_at >= $1
			AND (up.end_date IS NULL OR up.end_date <= $2)
			AND ($3::uuid IS NULL OR up.user_id = $3::uuid)
			AND ($4::text IS NULL OR s.name = $4);
	`

	var totalAmount float64
	var servicesJSON []byte
	var userIDResult string
	var startDateResult, endDateResult string

	err = manager.Conn.QueryRow(query, req.StartDate, req.EndDate, req.UserID, req.ServiceName).
		Scan(&userIDResult, &totalAmount, &servicesJSON, &startDateResult, &endDateResult)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Logger.Info("No subscriptions found", "user_id", req.UserID, "service_name", req.ServiceName)
			return &dbModel.CountingResponse{
				UserID:    userIDResult,
				Amount:    0,
				Services:  []dbModel.Service{},
				StartDate: req.StartDate.Format("2006-01-02"),
				EndDate:   req.EndDate.Format("2006-01-02"),
			}, nil
		}
		log.Logger.Error("Failed to query subscriptions", "error", err)
		return nil, err
	}

	var services []dbModel.Service
	if err := json.Unmarshal(servicesJSON, &services); err != nil {
		log.Logger.Error("Failed to unmarshal services JSON", "error", err)
		return nil, err
	}

	response := &dbModel.CountingResponse{
		UserID:    userIDResult,
		Amount:    totalAmount,
		Services:  services,
		StartDate: startDateResult,
		EndDate:   endDateResult,
	}

	log.Logger.Info("Subscriptions counted successfully",
		"user_id", response.UserID,
		"total_amount", response.Amount,
		"services_count", len(response.Services),
		"start_date", response.StartDate,
		"end_date", response.EndDate,
	)
	return response, nil
}
