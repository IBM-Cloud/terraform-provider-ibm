resource "null_resource" "assign_host_to_cluster" {
    depends_on = [null_resource.module_depends_on]
  triggers = {
    value = "${length(var.module_depends_on)}"
  }
  provisioner "local-exec" {
    command = ". ${path.cwd}/scripts/assigncluster.sh"
    environment = {
    hostname = var.host_vm
    location = var.location
    cluster_name=var.cluster_name
    zone = var.zone
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
  value = null_resource.assign_host_to_cluster.triggers.value
}