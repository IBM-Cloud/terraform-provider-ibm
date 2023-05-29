# Example for CatalogManagementV1

This example illustrates how to use the CatalogManagementV1

These types of resources are supported:

* ibm_cm_catalog
* ibm_cm_offering
* ibm_cm_version
* cm_offering_instance

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CatalogManagementV1 resources

ibm_cm_catalog resource:

```hcl
resource "ibm_cm_catalog" "cm_catalog_instance" {
  label = var.cm_catalog_label
  short_description = var.cm_catalog_short_description
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
}
```
ibm_cm_offering resource:

```hcl
resource "ibm_cm_offering" "cm_offering_instance" {
  catalog_id= var.cm_offering_catalog_id
  label = var.cm_offering_label
  offering_icon_url = var.cm_offering_offering_icon_url
  offering_docs_url = var.cm_offering_offering_docs_url
  offering_support_url = var.cm_offering_offering_support_url
  tags = var.cm_offering_tags
}
```
ibm_cm_version resource:

```hcl
resource "ibm_cm_version" "cm_version_instance" {
  catalog_id = var.cm_version_catalog_id
  offering_id = var.cm_version_offering_id
  tags = var.cm_version_tags
  target_kinds = var.cm_version_target_kinds
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
}
```
cm_offering_instance resource:

```hcl
resource "cm_offering_instance" "cm_offering_instance_instance" {
  label = var.cm_offering_instance_label
  catalog_id = var.cm_offering_instance_catalog_id
  offering_id = var.cm_offering_instance_offering_id
  kind_format = var.cm_offering_instance_kind_format
  version = var.cm_offering_instance_version
  cluster_id = var.cm_offering_instance_cluster_id
  cluster_region = var.cm_offering_instance_cluster_region
  cluster_namespaces = var.cm_offering_instance_cluster_namespaces
  cluster_all_namespaces = var.cm_offering_instance_cluster_all_namespaces
}
```

## CatalogManagementV1 Data sources

ibm_cm_catalog data source:

```hcl
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = var.cm_catalog_catalog_id
}
```
ibm_cm_offering data source:

```hcl
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = var.cm_offering_catalog_id
  offering_id = var.cm_offering_offering_id
}
```
ibm_cm_version data source:

```hcl
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}
```
cm_offering_instance data source:

```hcl
data "cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
```