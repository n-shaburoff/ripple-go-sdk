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

	resp, err := ripple.CreateQuoteCollection(resources.CreateQuoteCollection{
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
