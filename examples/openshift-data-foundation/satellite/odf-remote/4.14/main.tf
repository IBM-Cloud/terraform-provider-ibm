terraform {
  required_providers {
    ibm = {
      source = "IBM-Cloud/ibm"
      version = ">= 1.56.0"
    }
  }
}

provider "ibm" {
    ibmcloud_api_key = var.ibmcloud_api_key
    region = var.region
}

resource "ibm_satellite_storage_configuration" "storage_configuration" {
    location = var.location
    config_name = var.configName
    storage_template_name = var.storageTemplateName
    storage_template_version = var.storageTemplateVersion
    user_config_parameters = {
        "osd-size" = var.osdSize,
        "num-of-osd" = var.numOfOsd,
        "osd-storage-class" = var.osdStorageClassName,
        "billing-type" = var.billingType,
        "cluster-encryption" = var.clusterEncryption,
        "ibm-cos-endpoint"= var.ibmCosEndpoint,
        "ibm-cos-location"= var.ibmCosLocation,
        "ignore-noobaa"= var.ignoreNoobaa,
        "kms-base-url"= var.kmsBaseUrl,
        "kms-encryption"= var.kmsEncryption,
        "kms-instance-id"= var.kmsInstanceId,
        "kms-instance-name"= var.kmsInstanceName,
        "kms-token-url"= var.kmsTokenUrl,
        "odf-upgrade"= var.ocsUpgrade,
        "perform-cleanup"= var.performCleanup,
        "disable-noobaa-LB"= var.disableNoobaaLB,
        "encryption-intransit"= var.encryptionInTransit,
        "worker-pools"=var.workerPools,
        "worker-nodes"= var.workerNodes,
        "add-single-replica-pool" = var.addSingleReplicaPool,
        "taint-nodes" = var.taintNodes,
        "prepare-for-disaster-recovery" = var.prepareForDisasterRecovery
    }
    user_secret_parameters = {
        "iam-api-key"= var.iamAPIKey,
        "ibm-cos-access-key" = var.ibmCosAccessKey,
        "kms-root-key" = var.kmsRootKey,
        "kms-api-key" = var.kmsApiKey
    }
    delete_assignments = var.deleteAssignments
    update_assignments = var.updateAssignments
}

resource "ibm_satellite_storage_assignment" "storage_assignment" {
  assignment_name = var.assignmentName
  cluster = var.cluster
  controller = var.location
  config = var.configName
  depends_on = [ibm_satellite_storage_configuration.storage_configuration]
  update_config_revision = var.updateConfigRevision
}
