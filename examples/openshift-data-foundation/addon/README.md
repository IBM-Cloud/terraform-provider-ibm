# [Tech Preview] Deploying and Managing Openshift Data Foundation

This example shows how to deploy and manage the Openshift Data Foundation (ODF) on IBM Cloud VPC based RedHat Openshift cluster. Note this template is still in development, so please be advised before using in production.

This sample configuration will deploy the ODF, scale and upgrade it using the "ibm_container_addons" and "kubernetes_manifest" from the ibm terraform provider and kubernetes provider respectively.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan --var-file input.tfvars
$ terraform apply --var-file input.tfvars
```

Run `terraform destroy` when you don't need these resources.

## Example usage


### Create & Deployment:

The default input.tfvars is given below, the user should just change the value of the parameters in accorandance to their requirment. 

```hcl
# For add-on deployment
odfVersion = "4.12.0"

# For CRD Creation and Management
kube_config_path = "~/.kube/config"
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
osdSize = "250Gi"
osdStorageClassName = "ibmc-vpc-block-metro-10iops-tier"
workerNodes = null
```
### Editing the Ocscluster custom resource

The following variables in the `input.tfvars` file can be edited

* numOfOsd - To scale your storage
* workerNodes - To increase the number of Worker Nodes with ODF

```hcl
# For CRD Management
numOfOsd = "1" -> "2"
workerNodes = null -> "worker_1_ID,worker_2_ID"
```

### Update

The following variables in the `input.tfvars` file should be changed in order to upgrade the ODF add-on and the Ocscluster CRD.

* odfVersion - Specify the version you wish to upgrade to
* ocsUpgrade - Must be set to `true` to upgrade the CRD 

```hcl
# For ODF add-on upgrade
odfVersion = "4.12.0" -> "4.13.0"

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

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud_api_key | IBM Cloud API Key | `string` | yes |
| cluster | Name of the cluster. | `string` | yes |
| region | Region of the cluster | `string` | yes |
| odfVersion | Version of the ODF add-on | `string` | yes |
| kube_config_path |The Cluster config with absolute path | `string` | yes |
| osdSize | Enter the size for the storage devices that you want to provision for the Object Storage Daemon (OSD) pods | `string` | yes
| numOfOsd | The Number of OSD | `string` | yes
| osdStorageClassName | Enter the storage class to be used to provision block volumes for Object Storage Daemon (OSD) pods | `string` | yes
| autoDiscoverDevices | Set to true if automatically discovering local disks | `string` | no |
| clusterEncryption | To enable at-rest encryption of all disks in the storage cluster | `string` | no |
| hpcsBaseUrl | The HPCS Base URL | `string` | no |
| hpcsInstanceId | The HPCS Service ID | `string` | no |
| hpcsSecretName |  The HPCS secret name | `string` | no |
| hpcsServiceName | The HPCS service name | `string` | no |
| hpcsTokenUrl | The HPCS Token URL | `string` | no
| ignoreNoobaa | Set to true if you do not want MultiCloudGateway | `string` | no
| hpcsTokenUrl | The HPCS Token URL | `string` | no
| ocsUpgrade | Set to true to upgrade Ocscluster | `string` | no
| osdDevicePaths | IDs of the disks to be used for OSD pods if using local disks or standard classic cluster | `string` | no
| workerNodes | Provide the names of the worker nodes on which to install ODF. Leave blank to install ODF on all worker nodes | `string` | no



## Note

* Users should only change the values of the variables within quotes, variables should be left untouched with the default values if they are not set.
* `workerNodes` takes a string containing comma separated values of the names of the worker nodes you wish to enable ODF on.
* On `terraform apply --var-file=input.tfvars`, the add-on is enabled and the custom resource is created.
* During ODF update, please do not tamper with the `ocsUpgrade` variable, just change the value to true within quotation, without changing the format of the variable. 