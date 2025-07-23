# Deploying and Managing Openshift Data Foundation

This example shows how to deploy and manage the Openshift Data Foundation (ODF) on IBM Cloud VPC based RedHat Openshift cluster. Note this template is still in development, so please be advised before using in production.

This sample configuration will deploy the ODF, scale and upgrade it using the "ibm_container_addons" and "kubernetes_manifest" from the ibm terraform provider and kubernetes provider respectively.

For more information, about

* ODF Deployment, see [Deploying OpenShift Data Foundation on VPC clusters](https://cloud.ibm.com/docs/openshift?topic=openshift-deploy-odf-vpc&interface=ui)
* ODF Management, see [Managing your OpenShift Data Foundation deployment](https://cloud.ibm.com/docs/openshift?topic=openshift-ocs-manage-deployment&interface=ui)

#### Folder Structure

```ini
├── openshift-data-foundation
│   ├── addon
│   │   ├── ibm-odf-addon
│   │   │   ├── main.tf
│   │   ├── ocscluster
│   │   │   ├── main.tf
│   │   ├── createaddon.sh
│   │   ├── createcrd.sh
│   │   ├── updatecrd.sh
│   │   ├── updateodf.sh
│   │   ├── deleteaddon.sh
│   │   ├── deletecrd.sh
│   │   ├── main.tf
│   │   ├── variables.tf
│   │   ├── schematics.tfvars
```

* `ibm-odf-addon` - This folder is used to deploy a specific Version of Openshift-Data-Foundation with the `odfDeploy` parameter set to false i.e the add-on is installed without the ocscluster using the IBM-Cloud Terraform Provider.
* `ocscluster` - This folder is used to deploy the `OcsCluster` CRD with the given parameters set in the `schematics.tfvars` file.
* `addon` - This folder contains scripts to create the CRD and deploy the ODF add-on on your cluster. `The main.tf` file contains the `null_resource` to internally call the above two folders, and perform the required actions.

#### Note

You do not have to change anything in the `ibm-odf-addon` and `ocscluster` folders. You just have to input the required parameters in the `schematics.tfvars` file under the `addon` folder, and run terraform.

## Usage

### Option 1 - Command Line Interface

To run this example on your Terminal, first download this directory i.e `examples/openshift-data-foundation/`

```bash
$ cd addon/4.18.0
```

```bash
$ terraform init
$ terraform plan --var-file schematics.tfvars
$ terraform apply --var-file schematics.tfvars
```

Run `terraform destroy --var-file schematics.tfvars` when you don't need these resources.

### Option 2 - IBM Cloud Schematics

To Deploy & Manage the Openshift-Data-Foundation add-on using `IBM Cloud Schematics` please follow the below documentation

https://cloud.ibm.com/docs/schematics?topic=schematics-get-started-terraform

Please note you have to change the `terraform` keyword in the scripts to `terraform1.x` where `x` is the version of terraform you use in IBM Schematics, for example if you're using terraform version 1.3 in schematics make sure to change `terraform` -> `terraform1.3` in the .sh files.

## Example usage

### Deployment of ODF

The default schematics.tfvars is given below, the user should just change the value of the parameters in accorandance to their requirment.

```hcl
ibmcloud_api_key = "" # Enter your API Key
cluster = "" # Enter the Cluster ID
region = "us-south" # Enter the region

# For add-on deployment
odfVersion = "4.18.0"

# For CRD Creation and Management
autoDiscoverDevices = "false"
billingType = "advanced"
clusterEncryption = "false"
hpcsBaseUrl = null
hpcsEncryption = "false"
hpcsInstanceId = null
hpcsSecretName = null
hpcsServiceName = null
hpcsTokenUrl = null
ignoreNoobaa = "false"
numOfOsd = "1"
ocsUpgrade = "false"
osdDevicePaths = null
osdSize = "512Gi"
osdStorageClassName = "ibmc-vpc-block-metro-10iops-tier"
workerPools = null
workerNodes = null
encryptionInTransit = false
taintNodes = false
addSingleReplicaPool = false
ignoreNoobaa = false
disableNoobaaLB = false
enableNFS = false
useCephRBDAsDefaultStorageClass = false
resourceProfile = "balanced"
```

### Scale-Up of ODF

The following variables in the `schematics.tfvars` file can be edited

* numOfOsd - To scale your storage
* workerNodes - To increase the number of Worker Nodes with ODF
* workerPools -  To increase the number of Storage Nodes by adding more nodes using workerpool

```hcl
# For CRD Management
numOfOsd = "1" -> "2"
workerNodes = null -> "worker_1_ID,worker_2_ID"
```

### Upgrade of ODF

The following variables in the `schematics.tfvars` file should be changed in order to upgrade the ODF add-on and the Ocscluster CRD.

* odfVersion - Specify the version you wish to upgrade to
* ocsUpgrade - Must be set to `true` to upgrade the CRD

```hcl
# For ODF add-on upgrade
odfVersion = "4.17.0" -> "4.18.0"

# For Ocscluster upgrade
ocsUpgrade = "false" -> "true"
```

## Examples

* [ ODF Deployment & Management ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/openshift-data-foundation/deployment)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.14.8 |

## Providers

| Name | Version |
|------|---------|
| ibm | latest |
| kubernetes | latest |

## Inputs

| Name | Description | Type | Required | Default
|------|-------------|------|----------|--------|
| ibmcloud_api_key | IBM Cloud API Key | `string` | yes | -
| cluster | Name of the cluster. | `string` | yes | -
| region | Region of the cluster | `string` | yes | -
| odfVersion | Version of the ODF add-on | `string` | yes | 4.16.0
| osdSize | Enter the size for the storage devices that you want to provision for the Object Storage Daemon (OSD) pods | `string` | yes | 512Gi
| numOfOsd | The Number of OSD | `string` | yes | 1
| osdStorageClassName | Enter the storage class to be used to provision block volumes for Object Storage Daemon (OSD) pods | `string` | yes | ibmc-vpc-block-metro-10iops-tier
| autoDiscoverDevices | Set to true if automatically discovering local disks | `string` | no | true
| billingType | Set to true if automatically discovering local disks | `string` | no | advanced
| clusterEncryption | To enable at-rest encryption of all disks in the storage cluster | `string` | no | false
| hpcsEncryption | Set to true to enable HPCS Encryption | `string` | no | false
| hpcsBaseUrl | The HPCS Base URL | `string` | no | null
| hpcsInstanceId | The HPCS Service ID | `string` | no | null
| hpcsSecretName |  The HPCS secret name | `string` | no | null
| hpcsServiceName | The HPCS service name | `string` | no | null
| hpcsTokenUrl | The HPCS Token URL | `string` | no | null
| ignoreNoobaa | Set to true if you do not want MultiCloudGateway | `bool` | no | false
| ocsUpgrade | Set to true to upgrade Ocscluster | `string` | no | false
| osdDevicePaths | IDs of the disks to be used for OSD pods if using local disks or standard classic cluster | `string` | no | null
| workerPools | A list of the worker pools names where you want to deploy ODF. Either specify workerpool or workernodes to deploy ODF, if not specified ODF will deploy on all nodes | `string` | no | null
| workerNodes | Provide the names of the worker nodes on which to install ODF. Leave blank to install ODF on all worker nodes | `string` | no | null
| encryptionInTransit |To enable in-transit encryption. Enabling in-transit encryption does not affect the existing mapped or mounted volumes. After a volume is mapped/mounted, it retains the encryption settings that were used when it was initially mounted. To change the encryption settings for existing volumes, they must be remounted again one-by-one. | `bool` | no | false
| taintNodes | Specify true to taint the selected worker nodes so that only OpenShift Data Foundation pods can run on those nodes. Use this option only if you limit ODF to a subset of nodes in your cluster. | `bool` | no | false
| addSingleReplicaPool | Specify true to create a single replica pool without data replication, increasing the risk of data loss, data corruption, and potential system instability. | `bool` | no | false
| disableNoobaaLB | Specify true to disable to NooBaa public load balancer. | `bool` | no | false
| enableNFS | Enabling this allows you to create exports using Network File System (NFS) that can then be accessed internally or externally from the OpenShift cluster. | `bool` | no | false
| useCephRBDAsDefaultStorageClass | Enable to set the Ceph RADOS block device (RBD) storage class as the default storage class during the deployment of OpenShift Data Foundation | `bool` | no | false
| resourceProfile | Provides an option to choose a resource profile based on the availability of resources during deployment. Choose between lean, balanced and performance. | `string` | yes | balanced


Refer - https://cloud.ibm.com/docs/openshift?topic=openshift-deploy-odf-vpc&interface=ui#odf-vpc-param-ref

## Note

* Users should only change the values of the variables within quotes, variables should be left untouched with the default values if they are not set.
* `workerNodes` takes a string containing comma separated values of the names of the worker nodes you wish to enable ODF on.
* On `terraform apply --var-file=schematics.tfvars`, the add-on is enabled and the custom resource is created.
* During ODF update, please do not tamper with the `ocsUpgrade` variable, just change the value to true within quotation, without changing the format of the variable.
* During the `Upgrade of Odf` scenario on IBM Schematics, please make sure to change the value of `ocsUpgrade` to `false` after. Locally this is automatically handled using `sed`.
