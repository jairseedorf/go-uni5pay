package uni5pay

// Initialize Uni5Pay+ credentials
func (input Config) New() *Config {
	if empty(input.MerchantID) || empty(input.MerchantKey) {
		return nil
	}

	return &Config{
		MerchantID:  input.MerchantID,
		MerchantKey: input.MerchantKey,
	}
}

// Generate a new QR payment code.
func (input CodeInput) GenerateCode() (*CodeOutput, error) {
	config := input.Config.New()
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

// Get the current status of a payment.
func (input VerifyInput) VerifyTransaction() (*VerifyOutput, error) {
	config := input.Config.New()
	if config == nil {
		return nil, errClient
	}

	return verifyTransaction(VerifyInput{
		Config:     config,
		ExtOrderNo: input.ExtOrderNo,
		Amount:     input.Amount,
		Currency:   input.Currency,
	})
}

// Refund a payment
func (input RefundInput) RefundTransaction() (*RefundOutput, error) {
	config := input.Config.New()
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

// Verify a callback signature
func (input CallbackInput) VerifyCallback() error {
	config := input.Config.New()
	if config == nil {
		return errClient
	}

	return verifyCallback(input)
}
