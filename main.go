package ripple_go_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"os"
	"time"
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

func (c *client) CreateQuoteCollection(data resources.CreateQuoteCollection) (*resources.CreateQuoteCollectionResponse, error) {
	err := c.checkAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	response, err := c.Do.Post(data, createQuoteCollectionPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending create quote collection request")
	}

	var body resources.CreateQuoteCollectionResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) AcceptQuote(data resources.AcceptQuote) (*resources.Payment, error) {
	err := c.checkAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	response, err := c.Do.Post(data, acceptQuotePath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending accept quote request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) SettlePayment(paymentID string) (*resources.Payment, error) {
	err := c.checkAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqBody := resources.SettlePayment{}
	reqPath := fmt.Sprintf("%s%s/settle", settlePaymentPath, paymentID)
	response, err := c.Do.Post(reqBody, reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending settle payment request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) GetPaymentByID(paymentID string) (*resources.Payment, error) {
	err := c.checkAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqPath := fmt.Sprintf("%s%s", getPaymentByIDPath, paymentID)
	response, err := c.Do.Get(reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending get payment by id request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) checkAccessToken() error {
	nowTime := time.Now()
	difference := nowTime.Sub(c.Do.GetTokenTime()).Seconds()

	if difference > 3600 {
		err := c.Do.Authorize(authReqBody())
		if err != nil {
			return errors.Wrap(err, "failed to refresh access token")
		}
	}
	return nil
}

func authReqBody() resources.Authorization {
	return resources.Authorization{
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Audience:     audience,
	}
}
