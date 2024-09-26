package v1_test

import (
	"errors"
	"github.com/labstack/echo/v4"
	v1 "github.com/nable-fusion/fusion-api-template/api/v1"
	mocks2 "github.com/nable-fusion/fusion-api-template/mocks/internal_/service"
	"github.com/nable-fusion/fusion-cloud-common/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type ApiTestSuite struct {
	suite.Suite
}

func (suite *ApiTestSuite) SetupTest() {
	_ = log.NewLogger(log.NewConfig())
}

// TestGetWelcomePageOnSuccessReturns200AndMessage verifies that a successful GetWelcomePage
// returns a 200 status and an expected message.
// This requires a GET with a URL=/, but no headers are particularly needed since the
// request has no body to transmit.
// No failure case for this is tested because it really has no reason to generate one;
// the health check endpoint covers that eventuality.
func (suite *ApiTestSuite) TestGetWelcomePageOnSuccessReturns200AndMessage() {
	// assemble request context
	e := echo.New()
	subjectUrl := "/"
	req := httptest.NewRequest(http.MethodGet, subjectUrl, strings.NewReader(""))
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	// mock the service
	mockService := mocks2.ServiceInterface{}
	// assemble test artifact and exercise the target API method
	testSubject := v1.NewApi(&mockService)
	err := testSubject.GetWelcomePage(context)
	assert.NoError(suite.T(), err)
	// verify http response
	assert.Equal(suite.T(), http.StatusOK, rec.Code)
	assert.Equal(suite.T(), echo.MIMETextPlainCharsetUTF8, rec.Header().Get(echo.HeaderContentType))
	expectedMessage := "Welcome to the Fusion Api Template!"
	expectedMessageLength := len(expectedMessage)
	actualMessage := rec.Body.String()
	actualMessageLengthViaExamination := len(actualMessage)
	actualMessageLengthViaHeader := rec.Header().Get(echo.HeaderContentLength)
	actualMessageLengthViaHeaderParsed, parseErr := strconv.Atoi(actualMessageLengthViaHeader)
	assert.NoError(suite.T(), parseErr)
	assert.Less(suite.T(), 0, actualMessageLengthViaHeaderParsed)
	assert.Equal(suite.T(), expectedMessageLength, actualMessageLengthViaHeaderParsed)
	assert.Equal(suite.T(), actualMessageLengthViaExamination, actualMessageLengthViaHeaderParsed)
	assert.Equal(suite.T(), expectedMessage, actualMessage)
}

// TestGetHealthCheckOnSuccessReturns204 verifies that a successful GetHealthCheck
// returns a 204 but no message.
// This requires a GET with a URL=/, but no headers are particularly needed since the
// request has no body to transmit.
func (suite *ApiTestSuite) TestGetHealthCheckOnSuccessReturns204() {
	// assemble request context
	e := echo.New()
	subjectUrl := "/"
	req := httptest.NewRequest(http.MethodGet, subjectUrl, strings.NewReader(""))
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	// mock the service
	mockService := mocks2.ServiceInterface{}
	mockService.On("IntegrationPointsHealthCheck").
		Once().
		Return(nil)
	// assemble test artifact and exercise the target API method
	testSubject := v1.NewApi(&mockService)
	err := testSubject.GetHealthCheck(context)
	// verify the result is not an error
	if err != nil {
		assert.Fail(suite.T(), err.Error())
	}
	// verify http response
	assert.Equal(suite.T(), http.StatusNoContent, rec.Code)
	expectedMessage := ""
	actualMessage := rec.Body.String()
	assert.Equal(suite.T(), expectedMessage, actualMessage)
	// we don't verify a zero content-length because it is neither required nor expected no body
}

// TestGetHealthCheckOnFailureReturns503 verifies that an unsuccessful GetHealthCheck
// returns a 503 but no message.
// This requires a GET with a URL=/, but no headers are particularly needed since the
// request has no body to transmit.
func (suite *ApiTestSuite) TestGetHealthCheckOnFailureReturns503() {
	// assemble request context
	e := echo.New()
	subjectUrl := "/"
	req := httptest.NewRequest(http.MethodGet, subjectUrl, strings.NewReader(""))
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	// mock the service
	mockService := mocks2.ServiceInterface{}
	mockService.On("IntegrationPointsHealthCheck").
		Once().
		Return(errors.New("integration point unavailable"))
	// assemble test artifact and exercise the target API method
	testSubject := v1.NewApi(&mockService)
	err := testSubject.GetHealthCheck(context)
	// verify the result is not an error (because an HTTP response will announce the error)
	if err != nil {
		assert.Fail(suite.T(), err.Error())
	}
	// verify http response
	assert.Equal(suite.T(), http.StatusServiceUnavailable, rec.Code)
	expectedMessage := ""
	actualMessage := rec.Body.String()
	assert.Equal(suite.T(), expectedMessage, actualMessage)
	// we don't verify a zero content-length because it is neither required nor expected no body
}

func TestApiTestSuite(t *testing.T) {
	suite.Run(t, new(ApiTestSuite))
}
