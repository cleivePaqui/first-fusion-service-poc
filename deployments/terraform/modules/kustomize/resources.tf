terraform {
  required_providers {
    kustomization = {
      source  = "kbst/kustomization"
      version = "0.8.0"
    }
  }
}

provider "kustomization" {
  kubeconfig_path = "~/.kube/config"
}

data "kustomization" "fusion" {
  path = var.kustomization_path
}

resource "kustomization_resource" "fusion" {
  for_each = data.kustomization.fusion.ids

  manifest = data.kustomization.fusion.manifests[each.value]
}