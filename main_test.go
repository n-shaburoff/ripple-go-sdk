package ripple_go_sdk

import (
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var grantType = os.Getenv("GRANT_TYPE")
var clientID = os.Getenv("CLIENT_ID")
var clientSecret = os.Getenv("CLIENT_SECRET")
var audience = os.Getenv("AUDIENCE")
var authUrl = os.Getenv("AUTH_URL")
var baseURL = os.Getenv("BASE_URL")

func TestClient_GetQuoteCollection(t *testing.T) {
	servicer := NewServicer(authUrl, baseURL, grantType, clientID, clientSecret, audience)
	client, err := NewClient(servicer)
	assert.NoError(t, err)

	quoteCollection, err := client.CreateQuoteCollection(resources.CreateQuoteCollection{
		SendingAddress:   "trans1_usd_everest@test.mlt.everest",
		ReceivingAddress: "trans_php_everesttestpeer@test.cloud.everesttestpeer",
		Amount:           1,
		Currency:         "USD",
		QuoteType:        "SENDER_AMOUNT",
		PaymentMethod:    nil,
	})

	fmt.Println(quoteCollection)
	assert.NoError(t, err)
}
