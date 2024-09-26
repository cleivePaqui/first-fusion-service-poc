# Service Description

## Summary
This is a template repository which has the basic setup needed to create
and deploy a basic application service.

<img src="reference/img/app_overview.drawio.png" alt="Diagram" />

Primarily, the application ...

### Project Structure
The project can be divided into 3 broad categories:

#### Application Code
|                |                                                                                                                                                         |
|----------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
| /cmd           | Contains the main.go which launches the application and the Dockerfile used to create its container                                                     |
| /internal      | Contains a Service implementation - this is the area that someone using the template will want to change to reflect how they are providing API services |
| /test          | Contains code for automation tests                                                                                                                      |
| go.mod, go.sum | Top level files needed for dependency management in Go applications                                                                                     | 

#### Infrastructure
|               |                                                                                         |
|---------------|-----------------------------------------------------------------------------------------|
| /build        | Contains config files for configuration and additional Makefile helpers                 |
| /deployments  | Contains kubernetes and terraform files - configuration for deployment                  |
| Jenkinsfile   | Contains pipelines used by Jenkins to deploy the application                            |
| Makefile      | Contains tools and helpers, some of which are used in Jenkinsfile as part of deployment |
| skaffold.yaml | Used to run the application in a local kubernetes cluster                               | 

#### Everything else
|                   |                                                                |
|-------------------|----------------------------------------------------------------|
| /docs             | Contains all the repository documentation                      |
| .dockerignore     | Needed for running tests locally                               |
| .gitignore        | Needed for general repository hygiene                          |
| catalog-info.yaml | Needed for repository documentation to be visible in Backstage |
| README.md         | Needed for general repository hygiene                          | 
