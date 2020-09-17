
// data "external" "create_location" {
//   // program = ["bash", "${path.root}/setlocation.sh"]
//     program = ["bash", "${path.cwd}/scripts/setlocation.sh"]

//   query = {
//     zone = var.zone
//     location = var.location
//     cos_key = var.cos_key
//     cos_key_id = var.cos_key_id
//     label = var.label
//   }
// }

// output "location_id" {
//   value = "${data.external.create_location.result.location_id}"
// }
// output "path" {
//   value = "${data.external.create_location.result.path}"
// }
resource "null_resource" "create_location" {
    triggers = {
    value = "${length(var.module_depends_on)}"
  }
  provisioner "local-exec" {
    command = ". ${path.cwd}/scripts/setlocation.sh"
    environment = {
    ZONE = var.zone
    LOCATION = var.location
    # COS_KEY = var.cos_key
    # COS_KEY_ID = var.cos_key_id
    LABEL = var.label
    }
  }

}

// resource "null_resource" "module_depends_on" {
//     depends_on = [null_resource.create_location]

//   triggers = {
//     value = "${length(var.module_depends_on)}"
//   }
// }

output "trigger" {
  value = null_resource.create_location.triggers.value
}