package ripple_go_sdk

import (
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
)

var clientID = os.Getenv("CLIENT_ID")
var clientSecret = os.Getenv("CLIENT_SECRET")
var audience = os.Getenv("AUDIENCE")

const (
	authorizationPath         = "/oauth/token"
	createQuoteCollectionPath = "/v4/quote_collections"
	acceptQuotePath           = "/v4/payments/accept"
	settlePaymentPath         = "/v4/payments/"
	getPaymentByIDPath        = "/v4/payments/"
)

type Client interface {
	CreateQuoteCollection(data resources.CreateQuoteCollection) (*resources.CreateQuoteCollectionResponse, error)
	AcceptQuote(data resources.AcceptQuote) (*resources.Payment, error)
	SettlePayment(paymentID string) (*resources.Payment, error)
	GetPaymentByID(paymentID string) (*resources.Payment, error)
}

type client struct {
	Do Service
}

func NewClient(url *url.URL, cli *http.Client) (Client, error) {
	svc := servicer{
		http: cli,
	}

	err := svc.Authorize(resources.Authorization{
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Audience:     audience,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize")
	}

	svc.url = url

	return &client{
		&svc,
	}, nil
}

func (c client) CreateQuoteCollection(data resources.CreateQuoteCollection) (*resources.CreateQuoteCollectionResponse, error) {
	panic("implement me")
}

func (c client) AcceptQuote(data resources.AcceptQuote) (*resources.Payment, error) {
	panic("implement me")
}

func (c client) SettlePayment(paymentID string) (*resources.Payment, error) {
	panic("implement me")
}

func (c client) GetPaymentByID(paymentID string) (*resources.Payment, error) {
	panic("implement me")
}