package integration_test

import (
	"github.com/nable-fusion/fusion-api-template/test/flows/responses"

	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Health >", func() {
	It("Successfully returns healthy", Label("SMOKE"), func() {
		response := apiClient.SendHealthRequest()
		responses.ConfirmHealthSuccessResponse(response)
	})
})