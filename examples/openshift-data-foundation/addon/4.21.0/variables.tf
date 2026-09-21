variable "ibmcloud_api_key" {

    type = string
    description = "IBM Cloud API Key"

}

variable "cluster" {

    type = string
    description = "Cluster ID"
}

variable "region" {

    type = string
    description = "Enter Cluster Region"

}

variable "odfVersion" {

    type = string
    default = "4.21.0"
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

variable "hpcsEncryption" {

    type = string
    default = "false"
    description = "Set to true to enable HPCS Encryption"

}

variable "hpcsServiceName" {

    type = string
    default = null
    description = "Please provide HPCS service name"
}

variable "hpcsSecretName" {

    type = string
    default = null
    description = "Please provide the HPCS secret name"
}

variable "workerPools" {

    type = string
    default =  null
    description = "A list of the worker pool names where you want to deploy ODF. Either specify workerpool or workernodes to deploy ODF, if not specified ODF will deploy on all nodes"
}

variable "workerNodes" {

    type = string
    default =  null
    description = "Provide the names of the worker nodes on which to install ODF. Leave blank to install ODF on all worker nodes."
}

variable "hpcsInstanceId" {

    type = string
    default = null
    description = "Please provide HPCS Service ID"
}

variable "hpcsBaseUrl" {

    type = string
    default = null
    description = "Please provide HPCS Base URL"
}

variable "hpcsTokenUrl" {

    type = string
    default = null
    description = "Please provide HPCS token URL"
}

variable "encryptionInTransit" {

    type = bool
    default = false
    description = "Enter true to enable in-transit encryption. Enabling in-transit encryption does not affect the existing mapped or mounted volumes. After a volume is mapped/mounted, it retains the encryption settings that were used when it was initially mounted. To change the encryption settings for existing volumes, they must be remounted again one-by-one."

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

variable "disableNoobaaLB" {

    type = bool
    default = false
    description = "Specify true to disable to NooBaa public load balancer."

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

variable "setDefaultStorageClassForVirtualization" {

    type = bool
    default = false
    description = "If enabled, RBD virtualization StorageClass will be marked as the default for KubeVirt VM disks (persistent volumes) upon installation."

}

variable "enableAutomaticCapacityScaling" {

    type = bool
    default = false
    description = "To automatically add additional raw capacity equivalent to the configured deployment size whenever used capacity reaches 70%."

}

variable "clusterExpansionLimit" {

    type = string
    default = "12Ti"
    description = "The maximum limit to which the cluster can expand in the cloud. Automatic capacity scaling is suspended if exceeded."

}

variable "enableAutomaticBackup" {

    type = bool
    default = false
    description = "Enable automatic scheduled backups for the MultiCloud Object Gateway metadata database. It can be enabled only when `ignoreNoobaa` is set to false."

}

variable "backupFrequency" {

    type = string
    default = "@daily"
    description = "Defines how often automatic backups of the MultiCloud Object Gateway metadata database are created. Supported values are @daily, @weekly, and @monthly."

}

variable "numOfBackupCopies" {

    type = string
    default = "5"
    description = "Specify the number of automatic backup copies to retain for the MultiCloud Object Gateway metadata database. Supported values are between 1 and 12."

}
