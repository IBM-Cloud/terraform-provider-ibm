variable "ibmcloud_api_key" {
    type = string
    description = "IBM Cloud API Key"
}

variable "iamAPIKey" {
    type = string
    description = "Your IBM Cloud API Key"
}

variable "location" {
    type = string
    description = "The satellite location where you want to create your configuration"
}

variable "configName" {
    type = string
    description = "The name of your storage configuration"
}

variable "storageTemplateName" {
    type = string
    description = "The storage template for your configuration."
}

variable "storageTemplateVersion" {
    type = string
    description = "The version of the storage template."
}

variable "region" {
    type = string
    description = "Enter Satellite Location Region"
}

variable "odfVersion" {
    type = string
    default = "4.15.0"
    description = "Provide the ODF Version you wish to install on your cluster"
}

variable "numOfOsd" {
type = string
default = "1"
description = "Number of Osd"
}

variable "osdDevicePaths" {
type = string
description = "IDs of the disks to be used for OSD pods if using local disks or standard classic cluster"
default = null
}

variable "ocsUpgrade" {
    type = string
    default = "false"
    description = "Set to true to upgrade Ocscluster"

}

variable "clusterEncryption" {
    type = string
    default = "false"
    description = "Enable at-rest encryption of all disks in the storage cluster."
}


variable "billingType" {
    type = string
    default = "advanced"
    description = "Choose between advanced and essentials"
}

variable "ignoreNoobaa" {
    type = bool
    default = false
    description = "Set to true if you do not want MultiCloudGateway"
}

variable "performCleanup" {
    type = bool
    default = false
    description = "Set to true if you want to perform cleanup during assignment deletion"
}

variable "ibmCosEndpoint" {
    type = string
    default = null
    description = "The IBM COS regional public endpoint"
}

variable "ibmCosLocation" {
    type = string
    default = null
    description = "The location constraint that you want to use when creating your bucket. For example us-east-standard."
}

variable "ibmCosSecretKey" {
    type = string
    default = null
    description = "Your IBM COS HMAC secret access key."
}

variable "ibmCosAccessKey" {
    type = string
    default = null
    description = "Your IBM COS HMAC access key ID."
}

variable "kmsApiKey" {
    type = string
    default = null
    description = "IAM API key to access the KMS instance. The API key that you provide must have at least Viewer access to the KMS instance."
}

variable "kmsRootKey" {
    type = string
    default = null
    description = "KMS root key of your instance."
}

variable "osdSize" {
    type = string
    default = "250Gi"
    description = "Enter the size for the storage devices that you want to provision for the Object Storage Daemon (OSD) pods."
}

variable "osdStorageClassName" {
    type = string
    default = "ibmc-vpc-block-metro-10iops-tier"
    description = "Enter the storage class to be used to provision block volumes for Object Storage Daemon (OSD) pods."

}

variable "autoDiscoverDevices" {
    type = string
    default = "false"
    description = "Set to true if automatically discovering local disks"
}

variable "kmsEncryption" {
    type = string
    default = "false"
    description = "Set to true to enable HPCS Encryption"
}

variable "kmsInstanceName" {
    type = string
    default = null
    description = "Please provide HPCS service name"
}

variable "kmsSecretName" {
    type = string
    default = null
    description = "Please provide the HPCS secret name"
}

variable "workerPools" {
    type = string
    default =  null
    description = "Provide the names/ID of the workerpool on which to install ODF. Specify either workerpool or worker nodes to select storage nodes. If none of them specified, ODF will install on all workers."
}

variable "workerNodes" {
    type = string
    default =  null
    description = "Provide the names of the worker nodes on which to install ODF. Leave blank to install ODF on all worker nodes."
}

variable "kmsInstanceId" {
    type = string
    default = null
    description = "Please provide HPCS Service ID"
}

variable "kmsBaseUrl" {
    type = string
    default = null
    description = "Please provide HPCS Base URL"
}

variable "kmsTokenUrl" {
    type = string
    default = null
    description = "Please provide HPCS token URL"
}

variable "encryptionInTransit" {
    type = bool
    default = false
    description = "Enter true to enable in-transit encryption. Enabling in-transit encryption does not affect the existing mapped or mounted volumes. After a volume is mapped/mounted, it retains the encryption settings that were used when it was initially mounted. To change the encryption settings for existing volumes, they must be remounted again one-by-one."
}

variable "disableNoobaaLB" {
    type = bool
    default = false
    description = "Specify true to disable to NooBaa public load balancer."
}

variable "cluster" {
    type = string
    description = "Cluster ID or Name you wish to assign your configuration to."
}

variable "assignmentName" {
    type = string
    description = "Name of your storage assignment to a cluster"
}

variable "updateConfigRevision" {
    type = bool
    default = false
    description = "Set to true if you want to update the assignment with the latest configuration revision"
}

variable "deleteAssignments" {
    type = bool
    default = false
    description = "Set to true if you want to delete all the assignments of the configuration, during storage configuration destroy"
}

variable "updateAssignments" {
    type = bool
    default = false
    description = "Set to true if you want to update all the configuration's assignments with the latest revision"
}

variable "taintNodes" {
    type = bool
    default = false
    description = "Specify true to taint the selected worker nodes so that only OpenShift Data Foundation pods can run on those nodes. Use this option only if you limit ODF to a subset of nodes in your cluster."
}

variable "addSingleReplicaPool" {
    type = bool
    default = false
    description = "Specify true to create a single replica pool without data replication, increasing the risk of data loss, data corruption, and potential system instability."
}

variable "prepareForDisasterRecovery" {
    type = bool
    default = false
    description = "Specify true to set up the storage system for disaster recovery service with the essential configurations in place. This allows seamless implementation of disaster recovery strategies for your workloads."
}

variable "enableNFS" {

    type = bool
    default = false
    description = "Enabling this allows you to create exports using Network File System (NFS) that can then be accessed internally or externally from the OpenShift cluster."

}

variable "useCephRBDAsDefaultStorageClass" {

    type = bool
    default = false
    description = "Enable to set the Ceph RADOS block device (RBD) storage class as the default storage class during the deployment of OpenShift Data Foundation"

}

variable "resourceProfile" {

    type = string
    default = "balanced"
    description = "Provides an option to choose a resource profile based on the availability of resources during deployment. Choose between lean, balanced and performance."

}
