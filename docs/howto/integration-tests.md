# Integration Tests

## Dependencies

- You have completed your Golang setup as per the [Fusion Developet Getting Started](https://portal.n-able.dev/docs/rm/system/fusion-common-developer-docs/howto/getting-started/) documentation.

- To verify if `Ginkgo` is already installed:

  ```shell
  ginkgo version
  ```

- To install `Ginkgo` (if not installed):

  ```shell
  make integration-tests-install
  ```

## Secrets

To run the tests you may require secret values that must be set as environment variables.

All secrets and instructions on how to retrieve them:

| Secret                           | How to Obtain                                                                                                           | Northstar Secret ID                  |
| -------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------ |
| Placeholder Cerberus Credentials | [MSP/Fusion/Placeholder Cerberus Credentials](https://secretserver.n-able.com/secretserver/app/#/secrets/14185/general) | fsn-placeholder-cerberus-credentials |

- **Add the Cerberus Credentials to your .zprofile or system environment variables as:**

  - `CERBERUS_FSN_PLACEHOLDER_CLIENT_ID`
  - `CERBERUS_FSN_PLACEHOLDER_CLIENT_SECRET`

- To set your test run environment variables:

  ```shell
  export CERBERUS_CLIENT_ID=$CERBERUS_FSN_PLACEHOLDER_CLIENT_ID
  export CERBERUS_CLIENT_SECRET=$CERBERUS_FSN_PLACEHOLDER_CLIENT_SECRET
  ```

  **This must be repeated on every new shell instance**

  **Using this approach prevents the secrets being logged to your shell history.** It also makes it easier to switch between different repositories without having to change your local environment.

## Run Tests

Setup your local secret environment variables as per the [Secrets](#secrets) section.

### Targeting DEV Deployment

- Run:

  ```shell
  make integration-tests-dev
  ```

### Targeting STG Deployment

- Run:

  ```shell
  make integration-tests-stg
  ```

### Targeting Local Deployment

- Run:

  ```shell
  make integration-tests-local
  ```

## Upgrading to Latest Test Automation Lib Version

- Run:

  ```shell
  go get github.com/nable-fusion/test-automation-lib-go/v2@latest
  ```

  ```shell
  go mod tidy
  ```

## References

- [API Testing Good Practices](https://portal.n-able.dev/docs/default/component/qa-standards/Convergence/Topics/test-automation-api-good-practices/)
- [General Automated Testing Good Practices](https://portal.n-able.dev/docs/default/component/qa-standards/Convergence/Topics/test-automation-good-practices/)
- [Ginkgo Setup Guide](https://portal.n-able.dev/docs/default/component/qa-standards/Convergence/How-To-Guides/how-to-setup-ginkgo/)
- [Ginkgo Tool Evaluation](https://portal.n-able.dev/docs/default/component/qa-standards/Fusion/Reference/test-automation-tool-evaluation/)
