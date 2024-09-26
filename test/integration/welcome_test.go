package integration_test

import (
	"github.com/nable-fusion/fusion-api-template/test/flows/responses"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Welcome >", func() {
	It("Successfully returns welcome message", func() {
		response := apiClient.SendWelcomeRequest()
		responses.ConfirmWelcomeSuccessResponse(response)
	})
})
