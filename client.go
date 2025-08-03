package uni5pay

// Initialize Uni5Pay+ credentials
func New(input Config) *Config {
	if empty(input.MerchantID) || empty(input.MerchantKey) {
		return nil
	}

	return &Config{
		MerchantID:  input.MerchantID,
		MerchantKey: input.MerchantKey,
	}
}

// GenerateCode API operation for Uni5Pay+
//
// This operation requests a new QR code from the Uni5Pay+ service.
//
// See the Uni5Pay+ API manual for usage information.
func GenerateCode(input CodeInput) (*CodeOutput, error) {
	config := input.Config
	if config == nil {
		return nil, errClient
	}

	return requestQrCode(CodeInput{
		Config:      config,
		Amount:      input.Amount,
		Currency:    input.Currency,
		CallbackURL: input.CallbackURL,
	})
}

// VerifyTransaction API operation for Uni5Pay+
//
// This operation requests the current status for a transaction from the Uni5Pay+ service.
//
// See the Uni5Pay+ API manual for usage information.
func VerifyTransaction(input VerifyInput) (*VerifyOutput, error) {
	config := input.Config
	if config == nil {
		return nil, errClient
	}

	return verifyTransaction(VerifyInput{
		Config:     config,
		ExtOrderNo: input.ExtOrderNo,
	})
}

// RefundTransaction API operation for Uni5Pay+
//
// This operation requests a refund for a transaction from the Uni5Pay+ service.
//
// See the Uni5Pay+ API manual for usage information.
func RefundTransaction(input RefundInput) (*RefundOutput, error) {
	config := input.Config
	if config == nil {
		return nil, errClient
	}

	return requestRefund(RefundInput{
		Config:     config,
		ExtOrderNo: input.ExtOrderNo,
		Amount:     input.Amount,
		Currency:   input.Currency,
	})
}

// VerifyCallback API operation for Uni5Pay+
//
// This operation validates the instant payment notification sent from the Uni5Pay+ service once a payment is successful.
//
// Once you generate a new QR code and provide a callback URL using the SDK, it creates an HMAC
// signature, signs it using your API key, and appends it to the URL. Once the payment succeeds, the IPN is sent
// to the callback URL and the signature can be used to validate the request on your end.
func VerifyCallback(input CallbackInput) error {
	config := input.Config
	if config == nil {
		return errClient
	}

	return verifyCallback(input)
}
