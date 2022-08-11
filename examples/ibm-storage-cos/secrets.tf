data "kubernetes_secret" "default-icr-io" {
  metadata {
    name = "default-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-icr-io" {

  metadata {
    name = "kube-system-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [null_resource.helm_obj_storage_plugin]

}

data "kubernetes_secret" "default-us-icr-io" {
  metadata {
    name = "default-us-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-default-us-icr-io" {

  metadata {
    name = "kube-system-default-us-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-us-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [kubernetes_secret.secret-kube-system-icr-io]

}

data "kubernetes_secret" "default-uk-icr-io" {
  metadata {
    name = "default-uk-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-default-uk-icr-io" {

  metadata {
    name = "kube-system-default-uk-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-uk-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [kubernetes_secret.secret-kube-system-default-us-icr-io]

}

data "kubernetes_secret" "default-de-icr-io" {
  metadata {
    name = "default-de-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-default-de-icr-io" {

  metadata {
    name = "kube-system-default-de-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-de-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [kubernetes_secret.secret-kube-system-default-uk-icr-io]

}

data "kubernetes_secret" "default-au-icr-io" {
  metadata {
    name = "default-au-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-default-au-icr-io" {

  metadata {
    name = "kube-system-default-au-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-au-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [kubernetes_secret.secret-kube-system-default-de-icr-io]

}

data "kubernetes_secret" "default-jp-icr-io" {
  metadata {
    name = "default-jp-icr-io"
    namespace = "default"
  }
}

resource "kubernetes_secret" "secret-kube-system-default-jp-icr-io" {

  metadata {
    name = "kube-system-default-jp-icr-io"
    namespace = "kube-system"
  }

  data = {
    ".dockerconfigjson" = lookup(data.kubernetes_secret.default-jp-icr-io.data, ".dockerconfigjson", "")
  }

  type = "kubernetes.io/dockerconfigjson"
  depends_on = [kubernetes_secret.secret-kube-system-default-au-icr-io]

}

