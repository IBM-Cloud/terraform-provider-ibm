
resource "kubernetes_namespace" "add_name_spaces" {
  count = length(var.configure_namespace)
  metadata {
    annotations = {
      name = element(var.configure_namespace, count.index)
    }

    labels = {
      name  = element(var.configure_namespace, count.index)
    }

    name = element(var.configure_namespace, count.index)
  }
}
