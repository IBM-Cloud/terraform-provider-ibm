resource "null_resource" "install_plugin" {
    provisioner "local-exec" {
        command = <<EOT
helm ibmc install ibm-charts/ibm-object-storage-plugin --name ibm-object-storage-plugin --kubeconfig "${data.ibm_container_cluster_config.cluster_config.config_file_path}"
sleep 60
        EOT
    }
    depends_on = [null_resource.helm_obj_storage_plugin]
}
