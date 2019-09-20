
data "ibm_container_cluster_config" "test_config" {
  cluster_name_id = "${var.cluster_name_id}"

}
data "external" "iks_token" {
  program = ["sh", "${path.module}/token.sh"]

  query = {
    server_url = "https://c100-e.us-east.containers.cloud.ibm.com:30129"
  }
}


provider "kubernetes" {
  host      = "https://c100-e.us-east.containers.cloud.ibm.com:30129"
  token    = "${data.external.iks_token.result.token}"
  config_path = "${data.ibm_container_cluster_config.test_config.config_file_path}"
}

resource "kubernetes_secret" "example" {
  metadata {
    name = "${var.secret_name}"
    namespace = "${var.namespace}"

  }

  data = {
    username = "admin"
    password = "Password"
  }

  type = "kubernetes.io/basic-auth"
}

resource "kubernetes_config_map" "example" {
  metadata {
    name = "newconfig3"
    namespace = "${var.namespace}"
  }


  data = {
    "newconfig3.sh" = "${file("${path.module}/wrapper.sh")}"
  }

}
resource "kubernetes_job" "demo412" {

  depends_on = ["kubernetes_secret.example","kubernetes_config_map.example"]
  # depends_on = ["kubernetes_config_map.example"]

  metadata {
    name = "demo412"
    #if not using namespace then it will go to default namespace
    namespace = "${var.namespace}"
  }
  spec {

    template {
      metadata {}
      spec {
        container {
          
          name    = "demo412"
          image   = "us.icr.io/smjtnamespace/ubuntu5:1.0"
            
          env {
          name  = "AA"
          value = "Injected AA"
          }

          env {
          name  = "BB"
          value = "Injected BB"
          }
          command = ["/scripts/newconfig3.sh"]
          
          port= [
            {
            name = "demo412"
            container_port = 2368
            protocol = "TCP"
          }
          ]
          volume_mount=[

          {
            name = "newconfig3"
            mount_path = "/scripts"
          }
          ]
        }
        volume = [
          {
          name = "newconfig3"
          config_map = [

          
           {
            name = "newconfig3"
            default_mode = "0744"
          }
          ]
        }
        ]
      }
    }
  }
}
