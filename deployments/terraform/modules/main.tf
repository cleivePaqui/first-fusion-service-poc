provider "aws" {
  region = var.aws-region
  allowed_account_ids = [
    var.aws-account-id
  ]
}

terraform {
  required_version = ">= 0.12"

  backend "s3" {
    encrypt = "true"
  }
}

locals {
  tags = {
    owner = "Fusion"
    organisation = "N-able"
    app = var.fusion-app-name
  }
  region = var.aws-region
}

module "kustomization" {
  source = "../modules/kustomize"
  kustomization_path = var.kustomization_path
}
