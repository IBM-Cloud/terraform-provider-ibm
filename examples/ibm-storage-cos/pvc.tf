resource "ibm_resource_key" "resourceKey" {
  name = var.pvc_config.secret_name
  role = "Writer"
  resource_instance_id = ibm_resource_instance.cos_instance.id
  parameters = {"HMAC": true}
  depends_on = [null_resource.install_plugin]
}

resource "kubernetes_secret" "pvc_secret" {
  count = length(var.configure_namespace)
  metadata {
    name = var.pvc_config.secret_name
    namespace = element(var.configure_namespace, count.index)
     }
  data = {
    "access-key" = lookup(ibm_resource_key.resourceKey.credentials, "cos_hmac_keys.access_key_id", "")
    "secret-key" = lookup(ibm_resource_key.resourceKey.credentials, "cos_hmac_keys.secret_access_key", "")
  }

  type = "ibm/ibmc-s3fs"
  depends_on = [ibm_resource_key.resourceKey]

}

resource "kubernetes_persistent_volume_claim" "example" {
  metadata {
    name = var.pvc_config.name
    namespace = var.pvc_config.namespace
    annotations = {
        "ibm.io/auto-create-bucket" = var.pvc_config.auto_create
        "ibm.io/auto-delete-bucket" = var.pvc_config.auto_delete
        "ibm.io/bucket" = var.pvc_config.bucket_name 
        "ibm.io/secret-name" = var.pvc_config.secret_name
        "ibm.io/endpoint" = "https://s3.${var.cos_endpoint}.${var.region}.cloud-object-storage.appdomain.cloud"
    }
  }
  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "8Gi"
      }
    }
    storage_class_name = var.pvc_config.storage_class_name
  }
  depends_on = [kubernetes_secret.pvc_secret]
}

