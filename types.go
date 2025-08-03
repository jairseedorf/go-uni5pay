package uni5pay

import "net/http"

type Config struct {
	MerchantID  string
	MerchantKey string
}

type CodeInput struct {
	Config      *Config
	Amount      float64
	Currency    string
	CallbackURL string

	RedirectFailedURL  string
	RedirectSuccessURL string
}

type CodeOutput struct {
	ExtOrderNo string `json:"ext_order_no"`
	OrderNo    string `json:"order_no"`
	QrCode     string `json:"qr_code"`
	Deeplink   string `json:"deep_link"`
}

type RefundInput struct {
	Config     *Config
	ExtOrderNo string
	Amount     float64
	Currency   string
}

type RefundOutput struct{}

type VerifyInput struct {
	Config     *Config
	ExtOrderNo string
	Amount     float64
	Currency   string
}

type VerifyOutput struct {
	Status string `json:"status"`
}

type CallbackInput struct {
	Config  *Config
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
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
}

type verifyRes struct {
	ExtOrderNo string  `json:"extOrderNo"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Status     string  `json:"status"`
}
