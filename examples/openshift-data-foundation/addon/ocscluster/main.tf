terraform {
  required_providers {
    kubernetes = {
      version = "2.18.1"
    }
  }
}

provider "kubernetes" {
config_path = var.kube_config_path
}


resource "kubernetes_manifest" "ocscluster_ocscluster_auto" {
  manifest = {
    "apiVersion" = "ocs.ibm.io/v1"
    "kind" = "OcsCluster"
    "metadata" = {
      "name" = "ocscluster-auto"
    }
    "spec" = {
      "autoDiscoverDevices" = var.autoDiscoverDevices,
      "billingType" = var.billingType,
      "clusterEncryption" = var.clusterEncryption,
      "hpcsBaseUrl" = var.hpcsBaseUrl,
      "hpcsEncryption" = var.hpcsEncryption,
      "hpcsInstanceId" = var.hpcsInstanceId,
      "hpcsSecretName" = var.hpcsSecretName,
      "hpcsServiceName" = var.hpcsServiceName,
      "hpcsTokenUrl" = var.hpcsTokenUrl,
      "ignoreNoobaa" = var.ignoreNoobaa,
      "numOfOsd" = var.numOfOsd,
      "ocsUpgrade" = var.ocsUpgrade,
      "osdDevicePaths" = var.osdDevicePaths,
      "osdSize" = var.osdSize,
      "osdStorageClassName" = var.osdStorageClassName,
      "workerNodes" = var.workerNodes==null ? null : split(",", var.workerNodes)
    }
  }

  field_manager {
    force_conflicts = true
  }
}