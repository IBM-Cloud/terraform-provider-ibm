resource "null_resource" "run_ssh" {
      depends_on = [var.module_depends_on]
  triggers = {
    value = "${length(var.module_depends_on)}"
  }
  count = var.ipcount
  connection {
    type     = "ssh"
    user     = "root"
    # host = var.hostip[4]
    host = "${element(var.hostip, count.index)}"
    private_key = "${file(var.private_ssh_key)}"
    # password = "mfaFTp7f"
    # password = "${element(var.passwords, count.index)}"

  }

  provisioner "file" {
    source      = var.path
    destination = "/tmp/attach.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "subscription-manager refresh",
      "subscription-manager repos --enable=*",
      "nohup sudo bash /tmp/attach.sh &",
      "sleep 120",
    ]
  }

}

resource "null_resource" "module_depends_on" {
    // depends_on = [null_resource.run_ssh]
  triggers = {
    value = "${length(var.module_depends_on)}"
  }
}

output "trigger" {
  value = null_resource.run_ssh.0.triggers.value
}

// output "depends_on_var" {
//   value = var.ipcount
// }
