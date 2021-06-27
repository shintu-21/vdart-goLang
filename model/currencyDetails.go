package model

type CurrencyDetails struct {
	CurrencyDetails []CurrencyDetail `json:"currency_details"`
}

type CurrencyDetail struct {
	ID                 string `json:"id,omitempty"`
	FullName           string `json:"fullName,omitempty"`
	Crypto             bool   `json:"crypto,omitempty"`
	PayinEnabled       bool   `json:"payinEnabled,omitempty"`
	PayinPaymentID     bool   `json:"payinPaymentId,omitempty"`
	PayinConfirmations int    `json:"payinConfirmations,omitempty"`
	PayoutEnabled      bool   `json:"payoutEnabled,omitempty"`
	PayoutIsPaymentID  bool   `json:"payoutIsPaymentId,omitempty"`
	TransferEnabled    bool   `json:"transferEnabled,omitempty"`
	Delisted           bool   `json:"delisted,omitempty"`
	PayoutFee          string `json:"payoutFee,omitempty"`
	PrecisionPayout    int    `json:"precisionPayout,omitempty"`
	PrecisionTransfer  int    `json:"precisionTransfer,omitempty"`
	LowProcessingTime  int    `json:"lowProcessingTime,omitempty"`
	HighProcessingTime int    `json:"highProcessingTime,omitempty"`
	AvgProcessingTime  int    `json:"avgProcessingTime,omitempty"`
}
