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
    default = "4.10.0"
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