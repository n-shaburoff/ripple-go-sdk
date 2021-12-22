package ripple_go_sdk

import (
	"bytes"
	"encoding/json"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
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
}

type servicer struct {
	http        *http.Client
	url         *url.URL
	accessToken string
}

func (c *servicer) Authorize(data resources.Authorization) error {
	addr, err := url.Parse(AuthUrl)
	if err != nil {
		return errors.Wrap(err, "failed to parse url")
	}

	c.url = addr

	response, err := c.Post(data, authorizationPath)
	if err != nil {
		return errors.Wrap(err, "error sending authorization request")
	}

	var body resources.AuthorizationResponse

	err = json.Unmarshal(response, &body)
	if err != nil {
		return errors.Wrap(err, "error unmarshalling response")
	}

	// setting JWT
	c.accessToken = body.AccessToken

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
