package uni5pay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func equals(v, s string) bool {
	return v == s
}

func empty(v string) bool {
	return equals(v, "")
}

func endpoint(path string) string {
	return fmt.Sprintf("%s://%s/v%d/%s", protocol, host, version, path)
}

func request(end, key string, body []byte) (*[]byte, error) {
	req, err := http.NewRequest(http.MethodPost, end, bytes.NewBuffer(body))
	if err != nil {
		return nil, errSystem
	}

	req.Header.Set("apiKey", key)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Uni5Pay+ Go Client by Seco")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, errSystem
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusBadRequest {
			return nil, errClient
		}

		if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
			return nil, errAuth
		}

		if res.StatusCode > http.StatusBadRequest && res.StatusCode < http.StatusInternalServerError {
			return nil, errConflict
		}

		return nil, errSystem
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errSystem
	}

	var code reqRes
	err = json.Unmarshal(bytes, &code)
	if err != nil {
		return nil, errSystem
	}

	if !equals(code.RspCode, success) {
		msg := strings.ToLower(code.RspMessage)
		return nil, errors.New(msg)
	}

	return &bytes, nil
}

func sign(pld, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(pld))
	return hex.EncodeToString(h.Sum(nil))
}

func compare(pld, ecp string) bool {
	return hmac.Equal([]byte(pld), []byte(ecp))
}

func convertCurr(c string) *string {
	var SRD string = "968"
	var USD string = "840"
	var EUR string = "978"

	switch c {
	case "SRD":
		return &SRD
	case "USD":
		return &USD
	case "EUR":
		return &EUR
	default:
		return nil
	}
}

func validateURL(raw string) error {
	if empty(strings.TrimSpace(raw)) {
		return errInvalidURL
	}

	u, err := url.ParseRequestURI(raw)
	if err != nil {
		return errInvalidURL
	}

	if !equals(u.Scheme, "https") {
		return errInvalidURL
	}

	if empty(u.Host) {
		return errInvalidURL
	}

	return nil
}

func signURL(callbackURL, signature string, timestamp int64) (*string, error) {
	err := validateURL(callbackURL)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(callbackURL)
	if err != nil {
		return nil, errCallbackParam
	}
	q := u.Query()

	if _, exists := q["signature"]; exists {
		return nil, errCallbackParam
	}
	if _, exists := q["timestamp"]; exists {
		return nil, errCallbackParam
	}

	q.Set("signature", signature)
	q.Set("timestamp", fmt.Sprintf("%d", timestamp))

	u.RawQuery = q.Encode()
	r := u.String()

	return &r, nil
}
