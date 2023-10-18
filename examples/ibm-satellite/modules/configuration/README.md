# This Module is used to create satellite storage configuration

This module creates a `satellite storage configuration` based on a storage template of your choice. For more information on storage templates and their parameters refer -> https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui 

## Prerequisite

* Set up the IBM Cloud command line interface (CLI), the Satellite plug-in, and other related CLIs.
* Install cli and plugin package
```console
    ibmcloud plugin install container-service
```
## Usage

```
terraform init
```
```
terraform plan
```
```
terraform apply
```
```
terraform destroy
```
## Example Usage

``` hcl
module "satellite-storage-configuration" {
  source             = "./modules/configuration"
  location = var.location
  config_name = var.config_name
  storage_template_name = var.storage_template_name
  storage_template_version = var.storage_template_version
  user_config_parameters = var.user_config_parameters
  user_secret_parameters = var.user_secret_parameters
  storage_class_parameters = var.storage_class_parameters
} 
```

### Example using the `odf-remote` storage template
``` hcl
resource "ibm_satellite_storage_configuration" "odf_storage_configuration" {
    location = var.location
    config_name = var.config_name
    storage_template_name = "odf-remote"
    storage_template_version = "4.12"
    user_config_parameters = {
        osd-size = "100Gi"
        osd-storage-class = "ibmc-vpc-block-metro-5iops-tier"
        billing-type = "advanced"
        cluster-encryption = "false"
        ibm-cos-endpoint = ""
        ibm-cos-location = ""
        ignore-noobaa = "false"
        kms-base-url = ""
        kms-encryption = "false"
        kms-instance-id = ""
        kms-instance-name = ""
        kms-token-url = ""
        num-of-osd = "1"
        odf-upgrade = "false"
        perform-cleanup = "false"
        worker-nodes = ""	
    }
    user_secret_parameters = {
        iam-api-key = "api-key-value"
    }
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| config_name | The Storage Configuration Name. | `string` | true |
| storage_template_name | The Name of the Storage Template to create the configuration. | `string` | true |
| storage_template_version | The Version of the Storage Template. | `string` | true |
| user_config_parameters | The different configuration parameters available based on the selected storage template | `map` | true |
| user_secret_parameters | The different secrets required based on the selected storage template | `map` | true |
| storage_class_parameters | Define your own storage classes if supported by the storage template | `list[map]` | true |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->