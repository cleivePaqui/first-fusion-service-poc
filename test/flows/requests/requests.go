package requests

import (
	"context"

	apiV1 "github.com/nable-fusion/fusion-api-template/api/v1"

	"github.com/nable-fusion/test-automation-lib-go/v2/pkg/errorutilities"
)

/*
Adding wrapper functions for all requests here provides multiple benefits:
- The tests do not have to perform the error check on sending a request.
- There is no confusion as to which auto-gened request function to use (there are multiple per endpoint).
- The tests do not need to keep setting the context.
- The tests are protected from the API operation being renamed.
- If we ever replace our usage of `github.com/deepmap/oapi-codegen` there would be less impact on the test files.

The intention is that multiple API clients may be instantiated throughout tests when we want clients with valid auth, expired auth, invalid auth, no auth etc.
*/

type APIClient struct {
	httpClient *apiV1.ClientWithResponses
}

func NewAPIClient(httpClient *apiV1.ClientWithResponses) *APIClient {
	return &APIClient{
		httpClient: httpClient,
	}
}

func (c *APIClient) SendWelcomeRequest() *apiV1.GetWelcomePageResponse {
	response, err := c.httpClient.GetWelcomePageWithResponse(context.Background())
	errorutilities.ConfirmErrorHasNotOccurredWithMessage(err, "Failed to Send Welcome Request")

	return response
}

func (c *APIClient) SendHealthRequest() *apiV1.GetHealthCheckResponse {
	response, err := c.httpClient.GetHealthCheckWithResponse(context.Background())
	errorutilities.ConfirmErrorHasNotOccurredWithMessage(err, "Failed to Send Health Request")

	return response
}
