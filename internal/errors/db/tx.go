package db

import (
	"errors"
)

var (
	ErrTxNotStarted = errors.New("transaction not started")
	ErrTxAlreadyStarted = errors.New("transaction already started")
	ErrTxNotCommitted = errors.New("transaction not committed")
	ErrTxNotRolledBack = errors.New("transaction not rolled back")
)