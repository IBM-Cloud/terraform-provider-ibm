
data "ibm_container_cluster_config" "clusterConfig" {
  cluster_name_id = var.cluster_name
}


provider "helm" {
  service_account = "default"
  namespace       = "kube-system"
  kubernetes {
    config_path = data.ibm_container_cluster_config.clusterConfig.config_file_path
  }
}

resource "random_id" "name" {
  byte_length = 4
}

resource "kubernetes_secret" "example" {
  metadata {
    name      = "px-etcd-certs"
    namespace = "kube-system"
  }

  data = {
    "ca.pem" = var.secret_cert
    username = var.secret_username
    password = var.secret_password
  }

  type = "Opaque"
}

resource "helm_release" "test" {
  name        = "terahelm${random_id.name.hex}"
  namespace   = "kube-system"
  crepository = "https://raw.githubusercontent.com/IBM/charts/master/repo/community/"
  chart       = "portworx"
  wait        = true
  depends_on  = ["kubernetes_secret.example"]
  set {
    name  = "etcd.secret"
    value = "px-etcd-certs"
  }
  set {
    name  = "kvdb"
    value = var.etcd_endpoint
  }
  set {
    name  = "clusterName"
    value = var.cluster_name
  }

}

data "helm_release_status" "test" {
  name     = helm_release.test.metadata.0.name
  revision = 1
}