# This Module is used to create satellite link

This module creates `satellite link` for the specified location.

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
module "satellite-link" {
  source = "./modules/link"

  location       = module.satellite-location.location_id
  crn            = module.satellite-location.location_crn
}
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name                          | Description                                                       | Type     | Default | Required |
|-------------------------------|-------------------------------------------------------------------|----------|---------|----------|
| ibmcloud_api_key              | IBM Cloud API Key.                                                | string   | n/a     | yes      |
| location                      | satellite location ID.                                            | string   | n/a     | yes      |
| crn                           | satellite location CRN.                                           | string   | n/a     | yes      |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->