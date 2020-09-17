resource "null_resource" "dns" {
    depends_on = [null_resource.module_depends_on]

 provisioner "local-exec" {
    command = ". ${path.cwd}/scripts/dnsregister.sh"
    environment = {
    ip0 = var.host_ip.0
    ip1 = var.host_ip.1
    ip2 = var.host_ip.2
    location = var.location
    }
  }
}

resource "null_resource" "module_depends_on" {
    // depends_on = [null_resource.dns]

  triggers = {
    value = "${length(var.module_depends_on)}"
  }
}

output "trigger" {
  value = null_resource.module_depends_on.triggers.value
}