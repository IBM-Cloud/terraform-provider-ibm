#####################################################
# vpc worker replace/update
# Copyright 2023 IBM
#####################################################

#####################################################
# Read each worker information attached to cluster
#####################################################
data ibm_resource_group group {
  name = var.resource_group
}

data ibm_container_cluster_config cluster_config {
    cluster_name_id   = var.cluster_name
    resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_container_vpc_worker" "worker" {
  count                         = length(var.worker_list)
  cluster_name                  = var.cluster_name
  replace_worker                = element(var.worker_list, count.index)
  resource_group_id             = data.ibm_resource_group.group.id
  kube_config_path              = data.ibm_container_cluster_config.cluster_config.config_file_path
  sds                           = "ODF"
  sds_timeout                   = "30m"

  timeouts {
    create = (var.create_timeout != null ? var.create_timeout : null)
    delete = (var.delete_timeout != null ? var.delete_timeout : null)
  }
}