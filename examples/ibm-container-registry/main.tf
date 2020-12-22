provider "ibm" {
    region = "eu-gb"
}
data "ibm_resource_group" "rg" {
  name = var.resource_group_name
}
resource "ibm_cr_namespace" "namespace" {
  name              = var.name
  resource_group_id = data.ibm_resource_group.rg.id
}
