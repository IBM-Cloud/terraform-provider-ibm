resource "ibm_iam_authorization_policy" "s2sAuth1" {
  source_service_name         = "toolchain"
  source_resource_instance_id = var.toolchain_id
  target_service_name         = "kms"
  target_resource_instance_id = var.key_protect_instance_guid
  roles                       = ["Viewer", "ReaderPlus"]
}

resource "ibm_cd_toolchain_tool_keyprotect" "keyprotect" {
  toolchain_id = var.toolchain_id
  parameters {
    name           = var.key_protect_integration_name
    location         = var.key_protect_instance_region
    resource_group_name = var.resource_group
    instance_name  = var.key_protect_instance_name
  }
}

output "keyprotect_integration_name" {
  value = var.key_protect_integration_name
  # Before returning this tool integration name
  # used to construct {vault:: secret references,
  # the authorization_policy must have been successfully created,
  # and the tool integration must have been created, 
  # otherwise the secret references would not resolve and 
  # other tools using secret references could give errors during tool integration creation
  depends_on = [
    ibm_iam_authorization_policy.s2sAuth1,
    ibm_cd_toolchain_tool_keyprotect.keyprotect
  ]
}