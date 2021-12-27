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

func TestClient_CreateQuoteCollection(t *testing.T) {
	servicer := NewServicer(authUrl, baseURL, grantType, clientID, clientSecret, audience)
	client, err := NewClient(servicer)
	assert.NoError(t, err)

	resp, err := client.CreateQuoteCollection(resources.CreateQuoteCollection{
		SendingAddress:             "sf@rn.us.ca.san_francisco",
		ReceivingAddress:           "sf_gbp@rn.us.ca.san_francisco",
		Amount:                     1,
		Currency:                   "USD",
		QuoteType:                  "SENDER_AMOUNT",
		EnableQuotePerPayoutMethod: true,
		DigitalAssetOrigination:    false,
	})

	fmt.Println(resp)
	assert.NoError(t, err)
}
