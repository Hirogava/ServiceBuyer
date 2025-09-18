package postgres

import (
	"database/sql"

	log "github.com/Hirogava/ServiceBuyer/internal/config/logger"
	errors "github.com/Hirogava/ServiceBuyer/internal/errors/db"
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
