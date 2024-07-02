# Openshift Data Foundation - Local Deployment

This example shows how to deploy and manage the Openshift Data Foundation (ODF) on IBM Cloud Satellite based RedHat Openshift cluster.

This sample configuration will deploy the ODF, scale and upgrade it using the "ibm_satellite_storage_configuration" and "ibm_satellite_storage_assignment" resources from the ibm terraform provider.

For more information, about

* ODF Deployment & Management on Satellite, see [OpenShift Data Foundation for local devices](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-odf-local&interface=ui)

## Usage

### Option 1 - Command Line Interface

To run this example on your Terminal, first download this directory i.e `examples/openshift-data-foundation/`

```bash
$ cd satellite
```

```bash
$ terraform init
$ terraform plan --var-file input.tfvars
$ terraform apply --var-file input.tfvars
```

Run `terraform destroy --var-file input.tfvars` when you don't need these resources.

### Option 2 - IBM Cloud Schematics

To Deploy & Manage the Openshift-Data-Foundation add-on using `IBM Cloud Schematics` please follow the below documentation

https://cloud.ibm.com/docs/schematics?topic=schematics-get-started-terraform


## Example usage

### Deployment of ODF Storage Configuration and Assignment

The default input.tfvars is given below, the user should just change the value of the parameters in accorandance to their requirment.

```hcl
# Common for both storage configuration and assignment 
ibmcloud_api_key = ""
location = "" #Location of your storage configuration and assignment
configName = "" #Name of your storage configuration
region = ""


#ODF Storage Configuration
storageTemplateName = "odf-local"
storageTemplateVersion = "4.13"

## User Parameters
autoDiscoverDevices = "true"
osdDevicePaths = ""
billingType = "advanced"
clusterEncryption = "false"
kmsBaseUrl = null
kmsEncryption = "false"
kmsInstanceId = null
kmsInstanceName = null
kmsTokenUrl = null
ibmCosEndpoint = null
ibmCosLocation = null
ignoreNoobaa = false
numOfOsd = "1"
ocsUpgrade = "false"
workerNodes = null
encryptionInTransit = false
disableNoobaaLB = false
performCleanup = false

## Secret Parameters
ibmCosAccessKey = null
ibmCosSecretKey = null
iamAPIKey = "" #Required
kmsApiKey = null
kmsRootKey = null

#ODF Storage Assignment
assignmentName = ""
cluster = ""
updateConfigRevision = false

## NOTE ##
# The following variables will cause issues to your storage assignment lifecycle, so please use only with a storage configuration resource.
deleteAssignments = false
updateAssignments = false
```

Please note with this deployment the storage configuration and it's respective storage assignment is created to your specific satellite cluster in this example, if you'd like more control over the resources you can split it up into different files.

### Scale-Up of ODF

The following variables in the `input.tfvars` file can be edited

* numOfOsd - To scale your storage
* workerNodes - To increase the number of Worker Nodes with ODF

```hcl
numOfOsd = "1" -> "2"
workerNodes = null -> "worker_1_ID,worker_2_ID"
updateConfigRevision = true
```
In this example we set the `updateConfigRevision` parameter to true in order to update our storage assignment with the latest configuration revision i.e the OcsCluster CRD is updated with the latest changes.

You could also use `updateAssignments` to directly update the storage configuration's assignments, but if you have a dependent `storage_assignment` resource, it's lifecycle will be affected. It it recommended to use this parameter when you've only defined the `storage_configuration` resource.

### Upgrade of ODF
**Step 1:**
Follow the [Satellite worker upgrade documentation](https://cloud.ibm.com/docs/satellite?topic=satellite-sat-storage-odf-update&interface=ui) step 1 to step 7 to perform worker upgrade.

**Step 2:**
Follow the below steps to upgrade ODF to next version.
The following variables in the `input.tfvars` file should be changed in order to upgrade the ODF add-on and the Ocscluster CRD.

* storageTemplateVersion - Specify the version you wish to upgrade to
* ocsUpgrade - Must be set to `true` to upgrade the CRD

```hcl
# For ODF add-on upgrade
storageTemplateVersion = "4.13" -> "4.14"
ocsUpgrade = "false" -> "true"
```

Note this operation deletes the existing configuration and it's respective assignments, updates it to the next version and reassigns back to the previous clusters/groups. If used with a dependent assignment resource, it's lifecycle will be affected. It is recommended to perform this scenario when you've only defined the `storage_configuration` resource.

## Examples

* [ ODF Deployment & Management ](https://cloud.ibm.com/docs/satellite?topic=satellite-storage-odf-local&interface=ui)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.14.8 |

## Providers

| Name | Version |
|------|---------|
| ibm | latest |

## Inputs

| Name | Description | Type | Required | Default
|------|-------------|------|----------|--------|
| ibmcloud_api_key | IBM Cloud API Key | `string` | yes | -
| cluster | Name of the cluster. | `string` | yes | -
| region | Region of the cluster | `string` | yes | -
| storageTemplateVersion | Version of the Storage Template (odf-local) | `string` | yes | -
| storageTemplateName | Name of the Storage Template (odf-local)| `string` | yes | -
| numOfOsd | The Number of OSD | `string` | yes | 1
| autoDiscoverDevices | Set to true if automatically discovering local disks | `string` | no | true
| billingType | Set to true if automatically discovering local disks | `string` | no | advanced
| performCleanup |Set to true if you want to perform complete cleanup of ODF on assignment deletion. | `bool` | yes | false
| clusterEncryption | To enable at-rest encryption of all disks in the storage cluster | `string` | no | false
| iamApiKey | Your IAM API key. | `string` | true | -
| kmsEncryption | Set to true to enable HPCS Encryption | `string` | yes | false
| kmsBaseUrl | The HPCS Base URL | `string` | no | null
| kmsInstanceId | The HPCS Service ID | `string` | no | null
| kmsSecretName |  The HPCS secret name | `string` | no | null
| kmsInstanceName | The HPCS service name | `string` | no | null
| kmsTokenUrl | The HPCS Token URL | `string` | no | null
| ignoreNoobaa | Set to true if you do not want MultiCloudGateway | `bool` | no | false
| ocsUpgrade | Set to true to upgrade Ocscluster | `string` | no | false
| osdDevicePaths | IDs of the disks to be used for OSD pods if using local disks or standard classic cluster | `string` | no | null
| workerNodes | Provide the names of the worker nodes on which to install ODF. Leave blank to install ODF on all worker nodes | `string` | no | null
| encryptionInTransit |To enable in-transit encryption. Enabling in-transit encryption does not affect the existing mapped or mounted volumes. After a volume is mapped/mounted, it retains the encryption settings that were used when it was initially mounted. To change the encryption settings for existing volumes, they must be remounted again one-by-one. | `bool` | no | false
| disableNoobaaLB | Specify true to disable to NooBaa public load balancer. | `bool` | no | false

Refer - https://cloud.ibm.com/docs/satellite?topic=satellite-storage-odf-local&interface=ui#odf-local-4.13-parameters

## Note

* Users should only change the values of the variables within quotes, variables should be left untouched with the default values if they are not set.
* `workerNodes` takes a string containing comma separated values of the names of the worker nodes you wish to enable ODF on.
* During ODF Storage Template Update, it is recommended to delete all terraform related assignments before handed, as their lifecycle will be affected, during update new storage assignments are made back internally with new UUIDs.
