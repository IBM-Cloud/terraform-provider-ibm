
provider "helm" {
    service_account = "default"
    namespace = "kube-system"
    kubernetes {
      config_path = data.ibm_container_cluster_config.cluster_config.config_file_path
      load_config_file = true
    }
}

provider "kubernetes" {
    config_path = data.ibm_container_cluster_config.cluster_config.config_file_path
    load_config_file = true
}

resource "kubernetes_cluster_role_binding" "tiller" {
  metadata {
    name = "tiller-instanc2"
  }
  role_ref {
      api_group = "rbac.authorization.k8s.io"
      kind = "ClusterRole"
      name = "cluster-admin"
  }
  subject {
      kind = "User"
      name = "cluster-admin"
      api_group = "rbac.authorization.k8s.io"
  }
  subject {
      kind = "ServiceAccount"
      name = "default"
      namespace = "kube-system"
  }
  subject {
      kind = "Group"
      name = "system:masters"
      api_group = "rbac.authorization.k8s.io"
  }
}

resource "null_resource" "helm_init_local" {
    provisioner "local-exec" {
        command = <<EOT

sleep 10
helm init --kubeconfig "${data.ibm_container_cluster_config.cluster_config.config_file_path}"
sleep 10
        EOT
    }
    depends_on = [kubernetes_cluster_role_binding.tiller, data.ibm_container_cluster_config.cluster_config]
}
