package uni5pay

import "net/http"

// The config represent the top-level credentials required to call the API.
//
// These values are provided by the bank once you've setup your account.
type Config struct {
	// Your Uni5Pay+ Merchant ID
	MerchantID string `min:"32" type:"string" required:"true"`

	// Your Uni5Pay+ API key.
	MerchantKey string `min:"32" type:"string" required:"true"`
}

// The request paramaters for generating a new QR code.
type CodeInput struct {
	// Pointer to the credentials struct.
	//
	// Config is a required field.
	Config *Config `type:"structure" required:"true"`

	// Total amount to charge the customer.
	//
	// Amount is a required field.
	Amount float64 `min:"1" type:"number" required:"true"`

	// Currency code in ISO-4217 format.
	//
	// Currency is a required field.
	Currency string `lenght:"3" type:"string" required:"true"`

	// The URL where the instant payment notification is sent to.
	//
	// Use the VerifyCallback API to verify callbacks securely.
	CallbackURL string `type:"string" required:"false"`

	// The URL where the customer is redirected to if the payment fails.
	RedirectFailedURL string `type:"string" required:"false"`

	// The URL where the customer is redirected to once the payment succeeds.
	RedirectSuccessURL string `type:"string" required:"false"`
}

// The response parameters.
//
// Returned fields:
//
//   - ExtOrderNo
//     The external order number.
//
//   - OrderNo
//     The order number of the transaction.
//
//   - QrCode
//     The QR code secret used to generate the image.
//
//   - DeepLink
//     The URL for in-app linking. This opens the Uni5Pay+ app.
type CodeOutput struct {
	ExtOrderNo string `json:"ext_order_no"`
	OrderNo    string `json:"order_no"`
	QrCode     string `json:"qr_code"`
	Deeplink   string `json:"deep_link"`
}

// The request parameters for refunding a transaction.
type RefundInput struct {
	// Pointer to the credentials struct.
	//
	// Config is a required field.
	Config *Config `type:"structure" required:"true"`

	// The external order.
	//
	// ExtOrderNo is a required field.
	ExtOrderNo string `min:"8" type:"string" required:"true"`

	// Total amount to refund the customer.
	//
	// Amount is a required field.
	Amount float64 `min:"1" type:"number" required:"true"`

	// Currency code in ISO-4217 format.
	//
	// Currency is a required field.
	Currency string `lenght:"3" type:"string" required:"true"`
}

// The response parameters.
type RefundOutput struct{}

// The request parameters for verifying a transaction.
type VerifyInput struct {
	// Pointer to the credentials struct.
	//
	// Config is a required field.
	Config *Config `type:"structure" required:"true"`

	// The external order.
	//
	// ExtOrderNo is a required field.
	ExtOrderNo string `min:"8" type:"string" required:"true"`
}

// The response parameters.
//
// Returned fields:
//
//   - Status
//     The current status of the transaction.
type VerifyOutput struct {
	Status string `json:"status"`
}

// The request parameters for verifying a transaction.
type CallbackInput struct {
	// Pointer to the credentials struct.
	//
	// Config is a required field.
	Config *Config `type:"structure" required:"true"`

	// Pointer to the full HTTP request.
	//
	// Request is a required field.
	Request *http.Request
}

type reqRes struct {
	RspCode    string `json:"rspCode"`
	RspMessage string `json:"rspMsg"`
}

type codeReq struct {
	MchOrderNo string `json:"mchtOrderNo"`
	TerminalID string `json:"terminalId"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`

	UrlSuccess string `json:"url_success,omitempty"`
	UrlFailed  string `json:"url_failure,omitempty"`
	UrlNotify  string `json:"url_notify,omitempty"`
}

type codeRes struct {
	QrCode      string `json:"qrCode"`
	DeepLink    string `json:"deepLink"`
	ExtOrderNo  string `json:"extOrderNo"`
	OrderNo     string `json:"orderNo"`
	MchtOrderNo string `json:"mchtOrderNo"`
}

type refundReq struct {
	ExtOrderNo string `json:"OrigextOrderNo"`
	TerminalID string `json:"terminalId"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
}

type refundRes struct {
	ExtOrderNo string `json:"extOrderNo"`
	OrderNo    string `json:"orderNo"`
}

type verifyReq struct {
	ExtOrderNo string `json:"extOrderNo"`
}

type verifyRes struct {
	Status string `json:"status"`
}
