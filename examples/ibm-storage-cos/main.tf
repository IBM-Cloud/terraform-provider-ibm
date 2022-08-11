resource "ibm_resource_instance" "cos_instance" {
  name     = var.service_instance_name
  service  = "cloud-object-storage"
  plan     = "standard"
  location = "global"
}

resource "ibm_container_bind_service" "bind_service" {
  cluster_name_id     = var.cluster_id
  service_instance_id = element(split(":", ibm_resource_instance.cos_instance.id), 7)
  namespace_id        = "default"
  role                = "Writer"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id     = var.cluster_id
}

