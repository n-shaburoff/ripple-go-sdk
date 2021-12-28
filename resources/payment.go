package resources

import "time"

type Payment struct {
	PaymentID    string                   `json:"payment_id"`
	ContractHash string                   `json:"contract_hash"`
	PaymentState string                   `json:"payment_state"`
	ModifiedAt   time.Time                `json:"modified_at"`
	Contract     PaymentContract          `json:"contract"`
	UserInfo     []map[string]interface{} `json:"user_info"`
}

type PaymentContract struct {
	Quote Quote `json:"quote"`
}

type SettlePayment struct{}

type LockPayment struct{}

type CompletePayment struct{}
