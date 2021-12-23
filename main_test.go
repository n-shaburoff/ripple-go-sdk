package ripple_go_sdk

import (
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServicer_Authorize(t *testing.T) {
	svc := NewServicer()

	err := svc.Authorize(authReqBody())

	fmt.Println(svc.accessToken)
	assert.NoError(t, err)
}

func TestClient_CreateQuoteCollection(t *testing.T) {
	ripple, err := NewClient()
	assert.NoError(t, err)

	_, err = ripple.CreateQuoteCollection(resources.CreateQuoteCollection{
		SendingAddress:             "test.irl.rfc",
		ReceivingAddress:           "test.xrapid.rfc",
		Amount:                     25,
		QuoteType:                  "SENDER_AMOUNT",
		Currency:                   "USD",
		PaymentMethod:              "spei",
		DigitalAssetOrigination:    false,
		EnableQuotePerPayoutMethod: false,
	})
	assert.NoError(t, err)
}
