# How to Use This Repository

### Introduction
This repository can be used to quickly stand up a simple API Server.

The steps needed to achieve this are:
1. Apply template data to your repository
2. Create Docker Image repositories
3. Create Seed Job

### Apply template data to your repository

This project has been generated from a template in the `fusion-service-generator` repository
using some default `api-template` configuration values.

You can provide your own configuration values and regenerate the template.

Update the file `api-template.json` to reflect your own repository's functionality. The existing
values should provide an idea of what is expected.

Run the make target:
```
make update-from-template
```

This will reapply the template with your new values.

Note: This process can be reapplied as many times as you like. So if you do not
have some data available you can update the config file later
and re-run the make target to update the repo, without editing the files directly.

### Create Docker Image Repositories

Checkout the [fusion-cloud-foundation](https://github.com/nable-fusion/fusion-cloud-foundation) repository

Update the file `deployments/foundation/main.tf`, adding a new entry like below:

```
# Api Template Repository
module "fusion_my_api_server_repositories" {
  source = "../terraform/modules/ecr"
  app-name = "MyApiServer"
  release-repo-name = "fusion/my-api-server"
  snapshot-repo-name = "fusion/my-api-server-snapshot"
}
```

Create a PR for this change.

Once merged, this will allow the "Build and Push..." stages in the Jenkinsfile to push versions
of your application to the ECR (Elastic Container Repository)

### Create seed job in bld-common

Checkout the [bld-common](https://github.com/logicnow/bld-common) project.

Update the file `seed/jobs/Fusion/seed.yaml`, adding a new entry like below:

```
  - type: "multibranchPipelineJob"
    pipelineName: "Fusion/fusion-my-api-server"
    repoOwner: "nable-fusion"
    repository: "fusion-my-api-server"
    scriptPath: "Jenkinsfile"
    githubCredentials: "github-pat"
```

Create a PR for this change.

Once merged, this will allow you to see your applications builds in [NorthStar Jenkins](https://common.build.n-able.dev/)
