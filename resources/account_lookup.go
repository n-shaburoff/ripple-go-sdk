package resources

type AccountLookUp struct {
	DestinationAddress  string `json:"destination_address"`
	AccountLookupID     string `json:"account_lookup_id"`
	AccountLookupStatus string `json:"account_lookup_status"`
	FirstName           string `json:"first_name"`
	MiddleName          string `json:"middle_name"`
	LastName            string `json:"last_name"`
	OrgName             string `json:"org_name"`
	CountryCode         string `json:"country_code"`
	AccountID           string `json:"account_id"`
	AccountIDType       string `json:"account_id_type"`
	ResultStatus        string `json:"result_status"`
}
