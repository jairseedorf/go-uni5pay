# Uni5Pay+ Go SDK

A simple SDK to integrate with the Uni5Pay+ Payment Gatway in Go. This package is not affiliated with the Southern Commercial Bank.

## Prerequesites

In order to use this SDK, you'll need to setup a Uni5Pay+ Merchant Account and obtain a `Merchant ID` and `API Key` from the Southern Commercial Bank.

Visit the [Uni5Pay+](https://uni5pay.sr) website for more details.

## Installation

```bash
go get github.com/jairseedorf/go-uni5pay
```

## Features

- Generate QR codes
- Process refunds
- Verify transactions
- Verify callbacks

## Example Usage

```go
import "github.com/jairseedorf/go-uni5pay"

config := uni5pay.New(Config{
    MerchantID:  os.Getenv("MERCHANT_ID"),
    MerchantKey: os.Getenv("API_KEY")
})

output, err := uni5pay.GenerateCode(uni5pay.CodeInput{
    Config:   config,
    Currency: "SRD",
    Amount:   100,
})
if err != nil {
    log.Fatal(err)
}

fmt.Println("QR Code:", output.QrCode)
```

## Report an Issue

If you have run into a bug or want to discuss a new feature, please [file an issue](https://jairseedorf/go-uni5pay/issues).

## Licensee

Licensed under [MIT License](./LICENSE)
