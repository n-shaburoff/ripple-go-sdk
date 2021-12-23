package ripple_go_sdk

import (
	"bytes"
	"encoding/json"
	"github.com/n-shaburoff/ripple-go-sdk/config"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/kv"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	AuthUrl = "https://auth.eu.prod.ripplenet.net"
	BaseUrl = "https://rfc.test.ripplexcurrent.com"
)

type Service interface {
	Authorize(data resources.Authorization) error
	Resolve(path string) string
	Get(path string) ([]byte, error)
	Post(data interface{}, path string) ([]byte, error)

	CheckAccessToken() error
}

type servicer struct {
	http              *http.Client
	url               *url.URL
	accessToken       string
	tokenExpires      time.Time
	tokenTimeDuration float64
	authUrl           string
	baseUrl           string
}

func NewServicer() *servicer {
	cfg := config.NewUrler(kv.MustFromEnv())

	return &servicer{
		authUrl: cfg.URL().AuthUrl,
		baseUrl: cfg.URL().BaseUrl,
		http:    http.DefaultClient,
	}
}

func (c *servicer) Authorize(data resources.Authorization) error {
	auth, err := url.Parse(c.authUrl)
	if err != nil {
		return errors.Wrap(err,"failed to parse auth url")
	}

	c.url = auth

	response, err := c.Post(data, authorizationPath)
	if err != nil {
		return errors.Wrap(err, "error sending authorization request")
	}

	var body resources.AuthorizationResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling response")
	}

	base, err := url.Parse(c.baseUrl)
	if err != nil {
		return errors.Wrap(err, "failed to parse base url")
	}

	// setting JWT
	c.accessToken = body.AccessToken

	// setting time if getting token
	c.tokenExpires = time.Now()

	// setting token time duration
	c.tokenTimeDuration = float64(body.ExpiresIn)

	// setting base url
	c.url = base

	return nil
}

func (c *servicer) Resolve(path string) string {
	endpoint, err := url.Parse(path)
	if err != nil {
		panic(errors.New("error parsing path"))
	}

	return c.url.ResolveReference(endpoint).String()
}

func (c *servicer) Post(data interface{}, path string) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, errors.Wrap(err, "error marshaling data")
	}

	request, err := http.NewRequest("POST", c.Resolve(path), bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create post request")
	}

	request.Header.Set("Content-Type", "application/json")
	if path != authorizationPath {
		bearer := "Bearer " + c.accessToken
		request.Header.Add("Authorization", bearer)
	}

	r, err := c.http.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error sending request")
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		errorResponse, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Errorf("error: got status code %d, response %s", r.StatusCode, string(errorResponse))
	}

	return ioutil.ReadAll(r.Body)
}

func (c *servicer) Get(path string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.Resolve(path), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create get request")
	}

	request.Header.Set("Content-Type", "application/json")
	if path != authorizationPath {
		bearer := "Bearer " + c.accessToken
		request.Header.Add("Authorization", bearer)
	}

	r, err := c.http.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "error sending request")
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		errorResponse, _ := ioutil.ReadAll(r.Body)
		return nil, errors.Errorf("error: got status code %d, response %s", r.StatusCode, string(errorResponse))
	}

	return ioutil.ReadAll(r.Body)
}

func (c *servicer) CheckAccessToken() error {
	nowTime := time.Now()
	difference := nowTime.Sub(c.tokenExpires).Seconds()

	if difference > c.tokenTimeDuration {
		err := c.Authorize(authReqBody())
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
