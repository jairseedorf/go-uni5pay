package uni5pay

import (
	"encoding/json"
	"fmt"
	"strconv"
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
		payload := fmt.Sprintf("%s:%s:%d",
			terminal,
			input.Config.MerchantID,
			timestamp,
		)

		signature := sign(payload, input.Config.MerchantKey)
		callback, err := signURL(input.CallbackURL, signature, timestamp)
		if err != nil {
			return nil, err
		}

		body.UrlNotify = *callback
	}

	if !empty(input.RedirectSuccessURL) {
		err := validateURL(input.RedirectSuccessURL)
		if err != nil {
			return nil, err
		}

		body.UrlSuccess = input.RedirectSuccessURL
	}

	if !empty(input.RedirectFailedURL) {
		err := validateURL(input.RedirectFailedURL)
		if err != nil {
			return nil, err
		}

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

	payload := fmt.Sprintf("%s:%s:%s",
		terminal,
		input.Config.MerchantID,
		timestamp,
	)
	ecpSig := sign(payload, input.Config.MerchantKey)

	if !compare(recSig, ecpSig) {
		return errSignature
	}

	return nil
}
