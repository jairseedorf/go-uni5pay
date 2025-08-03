package uni5pay

import "errors"

var errSystem = errors.New("internal server error")
var errAuth = errors.New("invalid or unauthorized merchant key")
var errClient = errors.New("missing one or more required parameters")
var errConflict = errors.New("unable to process request at this time")
var errCurrency = errors.New("unsupported currency code")
var errSignature = errors.New("invalid callback signature")
