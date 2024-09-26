package bootstrap

import (
	"testing"

	apiV1 "github.com/nable-fusion/fusion-api-template/api/v1"

	"github.com/nable-fusion/fusion-api-template/test/flows/requests"

	"github.com/nable-fusion/test-automation-lib-go/v2/pkg/environmentutilities"
	"github.com/nable-fusion/test-automation-lib-go/v2/pkg/errorutilities"
	"github.com/nable-fusion/test-automation-lib-go/v2/pkg/jsonutilities"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Environment struct {
	APIServerHost string
}

func RegisterSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Test Suite")
}

func SuiteSetup() []byte {
	environment := Environment{
		APIServerHost: environmentutilities.ConfirmEnvVariableIsSet("API_SERVER_HOST"),
	}

	return jsonutilities.ParseModelAsJson(environment)
}

func ParseEnvironment(envBytes []byte) *Environment {
	env := &Environment{}
	jsonutilities.ParseJsonAsModel(envBytes, &env)

	return env
}

func SetupAPIClient(env *Environment) *requests.APIClient {
	clientWithResponses, err := apiV1.NewClientWithResponses(env.APIServerHost)
	errorutilities.ConfirmErrorHasNotOccurredWithMessage(err, "Failed to Create New Autogened HTTP Client!")

	return requests.NewAPIClient(clientWithResponses)
}

/*
	Not required for template API but added as an example for future reference
*/
/*
func SetupAPIClientWithMTLS(env *Environment) *apiV1.ClientWithResponses {
	httpClient := httputilities.SetupHTTPClientWithMTLS(env.MTLSCertPath, env.MTLSKeyPath)

	apiClientWithResponses, err := apiV1.NewClientWithResponses(
		env.APIServerHost,
		apiV1.WithHTTPClient(httpClient),
	)

	errorutilities.ConfirmErrorHasNotOccurredWithMessage(err, "Failed to Create New HTTP Client!")

	return apiClientWithResponses
}
*/
