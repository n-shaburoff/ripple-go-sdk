package ripple_go_sdk

import (
	"encoding/json"
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/pkg/errors"
)

const (
	authorizationPath         = "/oauth/token"
	createQuoteCollectionPath = "/v4/quote_collections"
	acceptQuotePath           = "/v4/payments/accept"
	settlePaymentPath         = "/v4/payments/%s/settle"
	getPaymentByIDPath        = "/v4/payments/%s"
	getAcceptedQuotesPath     = "/v4/payments/?state=ACCEPTED"
	lockPaymentPath           = "/v4/payments/%s/lock"
	completePaymentPath       = "/v4/payments/%s/complete"
	initiateAccountLookupPath = "/v4/account_lookups/request"
)

type Client interface {
	CreateQuoteCollection(data resources.CreateQuoteCollection) (*resources.CreateQuoteCollectionResponse, error)
	AcceptQuote(data resources.AcceptQuote) (*resources.Payment, error)
	SettlePayment(paymentID string) (*resources.Payment, error)
	GetPaymentByID(paymentID string) (*resources.Payment, error)
	GetAcceptedQuotes() (*resources.Payment, error)
	LockPayment(quoteID string, data resources.LockPayment) (*resources.Payment, error)
	CompletePayment(paymentID string, data resources.CompletePayment) (*resources.Payment, error)
	InitiateAccountLookUp(data resources.AccountLookUp) (*resources.AccountLookUp, error)
}

type client struct {
	Do Service
}

func NewClient(svc *servicer) (Client, error) {
	err := svc.Authorize(svc.creds)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authorize")
	}

	return &client{
		svc,
	}, nil
}

func (c *client) CreateQuoteCollection(data resources.CreateQuoteCollection) (*resources.CreateQuoteCollectionResponse, error) {
	err := c.Do.CheckAccessToken()
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
	err := c.Do.CheckAccessToken()
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
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqBody := resources.SettlePayment{}
	reqPath := fmt.Sprintf(settlePaymentPath, paymentID)
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
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqPath := fmt.Sprintf(getPaymentByIDPath, paymentID)
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

func (c *client) GetAcceptedQuotes() (*resources.Payment, error) {
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	response, err := c.Do.Get(getAcceptedQuotesPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending get accepted quotes request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) LockPayment(quoteID string, data resources.LockPayment) (*resources.Payment, error) {
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqPath := fmt.Sprintf(lockPaymentPath, quoteID)
	response, err := c.Do.Post(data, reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending get accepted quotes request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) CompletePayment(paymentID string, data resources.CompletePayment) (*resources.Payment, error) {
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	reqPath := fmt.Sprintf(completePaymentPath, paymentID)
	response, err := c.Do.Post(data, reqPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending complete payment request")
	}

	var body resources.Payment
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}

func (c *client) InitiateAccountLookUp(data resources.AccountLookUp) (*resources.AccountLookUp, error) {
	err := c.Do.CheckAccessToken()
	if err != nil {
		return nil, errors.Wrap(err, "old access token")
	}

	response, err := c.Do.Post(data, initiateAccountLookupPath)
	if err != nil {
		return nil, errors.Wrap(err, "error sending initiate look up account request")
	}

	var body resources.AccountLookUp
	err = json.Unmarshal(response, &body)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling response")
	}
	return &body, nil
}
