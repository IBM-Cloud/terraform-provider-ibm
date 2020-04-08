variable "cluster_id" {
  default = "bp6grr6d08osr67qso50"
}
variable "role_binding" {
  default = "rolebinding"
}
variable "chart_name" {
  default = "logDNA_chart"
}
variable "logDNA_tags" {
  default = ""
}
//ibm_resource_key datasource to get ingestion key...
data "ibm_resource_key" "resourceKey" {
  name = "myobjectkey"
}
resource "random_id" "name" {
  byte_length = 4
}
data "ibm_container_cluster_config" "clusterConfig" {
  cluster_name_id = var.cluster_id
  config_dir      = "/tmp"
}
provider "helm" {
  service_account = "default"
  namespace = "kube-system"
  kubernetes {
    config_path = data.ibm_container_cluster_config.clusterConfig.config_file_path
  }
}
provider "kubernetes" {
  config_path = data.ibm_container_cluster_config.clusterConfig.config_file_path
}
resource "kubernetes_role_binding" "example" {
  depends_on = [data.ibm_container_cluster_config.clusterConfig]
  metadata {
    name = "${var.role_binding}${random_id.name.hex}"
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subject {
    kind      = "User"
    name      = "admin"
    api_group = "rbac.authorization.k8s.io"
  }
  subject {
    kind      = "ServiceAccount"
    name      = "default"
    namespace = "kube-system"
  }
  subject {
    kind      = "Group"
    name      = "system:masters"
    api_group = "rbac.authorization.k8s.io"
  }
}

resource "helm_release" "test" {
  name       = "${var.chart_name}${random_id.name.hex}"
  namespace  = "default"
  chart      = "stable/logdna-agent"
  wait       = true
  depends_on = [kubernetes_cluster_role_binding.tiller]
  set {
    name  = "logdna.key"
    value = data.ibm_resource_key.resourceKey.credentials.ingestion_key
  }
  set {
    name  = "logdna.tags"
    value = var.logDNA_tags
  }
}
