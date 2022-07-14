data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_cd_toolchain" "toolchain_instance" {
  name        = var.toolchain_name
  description = var.toolchain_description
  resource_group_id = data.ibm_resource_group.resource_group.id
}

output "toolchain_id" {
  value = ibm_cd_toolchain.toolchain_instance.id
}