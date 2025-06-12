terraform {
  required_providers {
    kubernetes = {
      version = ">= 2.18.1"
    }
    ibm = {
      source = "IBM-Cloud/ibm"
      version = ">= 1.56.0"
    }
  }
}

provider "ibm" {
  region             = var.region
  ibmcloud_api_key = var.ibmcloud_api_key
}


data "ibm_container_cluster_config" "cluster_vpc" {
  cluster_name_id = var.cluster
}

provider "kubernetes" {
  host                   = data.ibm_container_cluster_config.cluster_vpc.host
  token                  = data.ibm_container_cluster_config.cluster_vpc.token
  cluster_ca_certificate = data.ibm_container_cluster_config.cluster_vpc.ca_certificate
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
      "osdDevicePaths" = var.osdDevicePaths==null ? null : split(",", var.osdDevicePaths),
      "osdSize" = var.osdSize,
      "osdStorageClassName" = var.osdStorageClassName,
      "workerPools" = var.workerPools==null ? null : split(",", var.workerPools),
      "workerNodes" = var.workerNodes==null ? null : split(",", var.workerNodes),
      "encryptionInTransit" = var.encryptionInTransit,
      "taintNodes" = var.taintNodes,
      "addSingleReplicaPool" = var.addSingleReplicaPool,
      "prepareForDisasterRecovery" = var.prepareForDisasterRecovery,
      "disableNoobaaLB" = var.disableNoobaaLB
    }
  }

  field_manager {
    force_conflicts = true
  }
}