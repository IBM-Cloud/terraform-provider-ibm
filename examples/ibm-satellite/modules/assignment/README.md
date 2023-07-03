# This Module is used to create satellite storage assignment

This module creates a `satellite storage assignment` based on a storage template of your choice. For more information on storage templates and their parameters refer -> https://cloud.ibm.com/docs/satellite?topic=satellite-storage-template-ov&interface=ui 

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
module "satellite-storage-assignment" {
  assignment_name = var.assignment_name
  cluster = var.cluster
  config = var.config
  controller = var.controller
} 
```

### Assigning a Configuration to a cluster
```hcl
resource "ibm_satellite_storage_assignment" "odf_assignment" {
  assignment_name = var.assignment_name
  config = var.config
  cluster = var.cluster
  controller = var.controller
}
```

### Assigning a Configuration to Cluster Groups
```hcl
resource "ibm_satellite_storage_assignment" "odf_assignment" {
  assignment_name = var.assignment_name
  config = var.config
  groups = var.groups
}
```

### Updating the Configuration Revision to a cluster
```hcl
resource "ibm_satellite_storage_assignment" "odf_assignment" {
  assignment_name = var.assignment_name
  config = var.config
  cluster = var.cluster
  controller = var.controller
  update_config_revision = true
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| assignment_name | Name of the Assignment. | `string` | true |
| groups | One or more cluster groups on which you want to apply the configuration. Note that at least one cluster group is required. | `list[string]` | true |
| cluster | ID of the Satellite cluster or Service Cluster that you want to apply the configuration to. | `string` | true |
| config | Storage Configuration Name or ID. | `string` | true |
| controller | The Name or ID of the Satellite Location. | `string` | true |
| update_config_revision | Update an assignment to the latest available storage configuration version. | `bool` | false |

## Note
  * You cannot use the `groups` argument with `cluster` & `controller`, this is applicable when creating assignments to cluster groups.
  * Similarly `cluster` & `controller` are to be used together and cannot be used with `groups`, this is applicable when creating assignments to clusters.


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->