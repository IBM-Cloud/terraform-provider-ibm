# This Module is used to create satellite ROKS Cluster

This module creates `openshift route`.

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
module "satellite-route" {
  source = "./modules/route"

  ibmcloud_api_key   = var.ibmcloud_api_key
  cluster_master_url = var.cluster_master_url
  route_name         = var.route_name
}
```
<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name                          | Description                                                       | Type     | Default | Required |
|-------------------------------|-------------------------------------------------------------------|----------|---------|----------|
| ibmcloud_api_key              | IBM Cloud API Key.                                                | string   | n/a     | yes      |
| cluster_master_url            | Satellite Cluster master URL                                      | string   | n/a     | yes      |
| route_name                    | Openshft route name                                               | string   | n/a     | yes      |


<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->