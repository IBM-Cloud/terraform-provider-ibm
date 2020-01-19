resource "null_resource" "helm_obj_storage_plugin" {
    provisioner "local-exec" {
        command = <<EOT
helm repo add ibm-charts https://icr.io/helm/ibm-charts   
helm repo update                                          
helm fetch --untar ibm-charts/ibm-object-storage-plugin   
helm plugin install ./ibm-object-storage-plugin/helm-ibmc 
ibmcloud ks cluster pull-secret apply --cluster "${var.cluster_id}" 
        EOT
    }
    depends_on = [null_resource.helm_init_local]
}
