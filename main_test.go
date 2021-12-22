package ripple_go_sdk

import (
	"fmt"
	"github.com/n-shaburoff/ripple-go-sdk/resources"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestServicer_Authorize(t *testing.T) {
	svc := servicer{
		http: http.DefaultClient,
	}

	err := svc.Authorize(resources.Authorization{
		GrantType:    "client_credentials",
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Audience:     audience,
	})

	fmt.Println(svc.accessToken)
	assert.NoError(t, err)
}
