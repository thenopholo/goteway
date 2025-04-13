package domain

import "errors"

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrDuplicateAPIKey   = errors.New("api key already exist")
	ErrInvoiceNotFound   = errors.New("invoice not found")
	ErrUnathorizedAccess = errors.New("unathorized access")
)
