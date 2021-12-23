package ripple_go_sdk

import (
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/config"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/kit/kv"
	"testing"
)

func TestServicer_Authorize(t *testing.T) {
	cfg := config.NewUrler(kv.MustFromEnv())
	svc := NewServicer(cfg)

	err := svc.Authorize(authReqBody())

	fmt.Println(svc.accessToken)
	assert.NoError(t, err)
}

func TestClient_CreateQuoteCollection(t *testing.T) {
	cfg := config.NewUrler(kv.MustFromEnv())
	svc := NewServicer(cfg)
	ripple, err := NewClient(svc)
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
