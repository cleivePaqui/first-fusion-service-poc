package integration_test

import (
	"testing"

	"github.com/nable-fusion/fusion-api-template/test/flows/bootstrap"
	"github.com/nable-fusion/fusion-api-template/test/flows/requests"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega/format"
)

func TestSuite(t *testing.T) {
	bootstrap.RegisterSuite(t)
}

var (
	apiClient *requests.APIClient
)

var _ = SynchronizedBeforeSuite(
	func() []byte {
		return bootstrap.SuiteSetup()
	},
	func(envBytes []byte) {
		env := bootstrap.ParseEnvironment(envBytes)

		apiClient = bootstrap.SetupAPIClient(env)

		format.TruncatedDiff = false
	},
)
