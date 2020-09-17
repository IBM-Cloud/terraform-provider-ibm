

resource "null_resource" "create_cluster" {
  depends_on = [var.module_depends_on]
  provisioner "local-exec" {
    command = ". ${path.cwd}/scripts/createcluster.sh"
    environment = {
    cluster_name = var.cluster_name
    location = var.location
    }
  }

}

resource "null_resource" "module_depends_on" {

  triggers = {
    value = "${length(var.module_depends_on)}"
  }
}

output "trigger" {
  value = null_resource.module_depends_on.triggers.value
}