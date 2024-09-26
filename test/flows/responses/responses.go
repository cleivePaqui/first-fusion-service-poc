package responses

import (
	"net/http"

	apiV1 "github.com/nable-fusion/fusion-api-template/api/v1"

	"github.com/nable-fusion/test-automation-lib-go/v2/pkg/httputilities"

	. "github.com/onsi/gomega"
)

func ConfirmWelcomeSuccessResponse(response *apiV1.GetWelcomePageResponse) {
	httputilities.ConfirmStatusCode(response, http.StatusOK)
	Expect(string(response.Body)).To(Equal("Welcome to the Fusion Api Template!"))
	httputilities.ConfirmResponseContentType(response.HTTPResponse, "text/plain; charset=UTF-8")
}

func ConfirmHealthSuccessResponse(response *apiV1.GetHealthCheckResponse) {
	httputilities.ConfirmStatusCode(response, http.StatusNoContent)
	httputilities.ConfirmResponseBodyIsEmpty(response)
	httputilities.ConfirmResponseHeaderDoesNotExist(response.HTTPResponse, "Content-Type")
}
