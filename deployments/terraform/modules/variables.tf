variable "fusion-app-name" {
  description = "Name of the app being deployed"
  default = "Fusion Api Template"
}

variable "aws-region" {
  description = "AWS region in which to apply"
}

variable "aws-account-id" {
  description = "AWS account in which to apply"
}

variable "kustomization_path" {
  description = "path to the kustomization folder"
}

variable "irsa_oidc_provider" {
  description = ""
}
