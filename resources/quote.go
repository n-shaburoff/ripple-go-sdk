package resources

type CreateQuoteCollection struct {
	SendingAddress             string `json:"sending_address"`
	ReceivingAddress           string `json:"receiving_address"`
	Amount                     int    `json:"amount"`
	QuoteType                  string `json:"quote_type"`
	Currency                   string `json:"currency"`
	PaymentMethod              string `json:"payment_method"`
	DigitalAssetOrigination    bool   `json:"digital_asset_origination"`
	EnableQuotePerPayoutMethod bool   `json:"enable_quote_per_payout_method"`
}

type CreateQuoteCollectionResponse struct {
	QuoteCollectionID string `json:"quote_collection_id"`
	Quotes            Quote  `json:"quotes"`
}

type Quote struct {
	QuoteID            string       `json:"quote_id"`
	CreatedAt          string       `json:"created_at"`
	ExpiresAt          string       `json:"expires_at"`
	Type               string       `json:"type"`
	PriceGuarantee     string       `json:"price_guarantee"`
	SenderAddress      string       `json:"sender_address"`
	ReceiverAddress    string       `json:"receiver_address"`
	Amount             string       `json:"amount"`
	CurrencyCode       string       `json:"currency_code"`
	CurrencyCodeFilter string       `json:"currency_code_filter"`
	ServiceType        string       `json:"service_type"`
	QuoteElements      QuoteElement `json:"quote_elements"`
	LiquidityWarning   string       `json:"liquidity_warning"`
	PaymentMethod      string       `json:"payment_method"`
}

type QuoteElement struct {
	QuoteElementID        string `json:"quote_element_id"`
	QuoteElementType      string `json:"quote_element_type"`
	QuoteElementOrder     string `json:"quote_element_order"`
	SenderAddress         string `json:"sender_address"`
	ReceiverAddress       string `json:"receiver_address"`
	SendingAmount         string `json:"sending_amount"`
	ReceivingAmount       string `json:"receiving_amount"`
	SendingFee            string `json:"sending_fee"`
	ReceivingFee          string `json:"receiving_fee"`
	SendingCurrencyCode   string `json:"sending_currency_code"`
	ReceivingCurrencyCode string `json:"receiving_currency_code"`
	FxRate                string `json:"fx_rate"`
	TransferCurrencyCode  string `json:"transfer_currency_code"`
}

type AcceptQuote struct {
	QuoteID          string                 `json:"quote_id"`
	SenderEndToEndID string                 `json:"sender_end_to_end_id"`
	InternalID       string                 `json:"internal_id"`
	UserInfo         map[string]interface{} `json:"user_info"`
}

