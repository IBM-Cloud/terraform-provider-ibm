resource "null_resource" "assign_host" {
    depends_on = [null_resource.module_depends_on]
  triggers = {
    value = "${length(var.module_depends_on)}"
  }
    count = var.ip_count

  provisioner "local-exec" {
    command = ". ${path.cwd}/scripts/assign.sh"
    environment = {
    hostname = "${element(var.host_vm, count.index)}"
    index = count.index
    LOCATION = var.location
    }
  }

}

resource "null_resource" "module_depends_on" {
    // depends_on = [null_resource.assign_host]

  triggers = {
    value = "${length(var.module_depends_on)}"
  }
}

output "trigger" {
  value = null_resource.assign_host.0.triggers.value
}