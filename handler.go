package uni5pay

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func requestQrCode(input CodeInput) (*CodeOutput, error) {
	curr := convertCurr(input.Currency)
	if curr == nil {
		return nil, errCurrency
	}

	body := codeReq{
		MchOrderNo: input.Config.MerchantID,
		TerminalID: terminal,
		Currency:   *curr,
		Amount:     fmt.Sprintf("%.2f", input.Amount),
	}

	if !empty(input.CallbackURL) {
		timestamp := time.Now().Unix()
		payload := fmt.Sprintf("%s:%s:%s:%d",
			terminal,
			input.Config.MerchantID,
			input.CallbackURL,
			timestamp,
		)
		signature := sign(payload, input.Config.MerchantKey)
		body.UrlNotify = fmt.Sprintf("%s?signature=%s&timestamp=%d",
			input.CallbackURL,
			signature,
			timestamp,
		)
	}
	if !empty(input.RedirectSuccessURL) {
		body.UrlSuccess = input.RedirectSuccessURL
	}
	if !empty(input.RedirectFailedURL) {
		body.UrlFailed = input.RedirectFailedURL
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, errSystem
	}

	req, err := request(endpoint("qrcode_get"), input.Config.MerchantKey, bytes)
	if err != nil {
		return nil, err
	}

	var res codeRes
	err = json.Unmarshal(*req, &res)
	if err != nil {
		return nil, errSystem
	}

	return &CodeOutput{
		ExtOrderNo: res.ExtOrderNo,
		OrderNo:    res.OrderNo,
		QrCode:     res.QrCode,
		Deeplink:   res.DeepLink,
	}, nil
}

func requestRefund(input RefundInput) (*RefundOutput, error) {
	curr := convertCurr(input.Currency)
	if curr == nil {
		return nil, errCurrency
	}

	body := refundReq{
		ExtOrderNo: input.ExtOrderNo,
		TerminalID: terminal,
		Currency:   *curr,
		Amount:     fmt.Sprintf("%.2f", input.Amount),
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, errSystem
	}

	req, err := request(endpoint("order_cancel"), input.Config.MerchantKey, bytes)
	if err != nil {
		return nil, err
	}

	var res refundRes
	err = json.Unmarshal(*req, &res)
	if err != nil {
		return nil, errSystem
	}

	return &RefundOutput{}, nil
}

func verifyTransaction(input VerifyInput) (*VerifyOutput, error) {
	body := verifyReq{
		ExtOrderNo: input.ExtOrderNo,
	}

	bytes, err := json.Marshal(body)
	if err != nil {
		return nil, errSystem
	}

	req, err := request(endpoint("transaction_verify"), input.Config.MerchantKey, bytes)
	if err != nil {
		return nil, err
	}

	var res verifyRes
	err = json.Unmarshal(*req, &res)
	if err != nil {
		return nil, errSystem
	}

	return &VerifyOutput{
		Status: res.Status,
	}, nil
}

func verifyCallback(input CallbackInput) error {
	url := input.Request.URL
	query := url.Query()

	recSig := query.Get("signature")
	timestamp := query.Get("timestamp")

	if empty(recSig) || empty(timestamp) {
		return errSignature
	}

	tsInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil || time.Since(time.Unix(tsInt, 0)) > 10*time.Minute {
		return errSignature
	}

	originalURL := strings.SplitN(url.String(), "?", 2)[0]
	payload := fmt.Sprintf("%s:%s:%s:%s",
		terminal,
		input.Config.MerchantID,
		originalURL,
		timestamp,
	)
	ecpSig := sign(payload, input.Config.MerchantKey)

	if !compare(recSig, ecpSig) {
		return errSignature
	}

	return nil
}
