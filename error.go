package uni5pay

import "errors"

var (
	errSystem        = errors.New("internal server error")
	errAuth          = errors.New("invalid or unauthorized merchant key")
	errClient        = errors.New("missing one or more required parameters")
	errInvalidURL    = errors.New("invalid URL format")
	errCallbackParam = errors.New("callback URL cannot contain reserved parameters: signature or timestamp")
	errConflict      = errors.New("unable to process request at this time")
	errCurrency      = errors.New("unsupported currency code")
	errSignature     = errors.New("invalid callback signature")
)
