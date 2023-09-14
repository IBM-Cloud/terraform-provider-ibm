# Register a secrets manager instance
resource "ibm_container_ingress_instance" "cluster_instance" {
  instance_crn = var.sm_instance_crn
  is_default = true
  cluster  = var.cluster_name_or_id
  secret_group_id = var.sm_secret_group_id
}


// Create a ibm_container_ingress_instance data source
data "ibm_container_ingress_instance" "ingress_instance" {
    instance_name = ibm_container_ingress_instance.cluster_instance.instance_name
    cluster = var.cluster_name_or_id
}


# Create an ingress tls secret
resource "ibm_container_ingress_secret_tls" "container_ingress_secret_tls" {
    cluster  = var.cluster_name_or_id
    secret_name = var.tls_secret_name
    secret_namespace = var.tls_secret_namespace
    cert_crn = var.secret_cert_crn
    persistence = true
}

// Create an ibm_container_ingress_secret_tls data source
data "ibm_container_ingress_secret_tls" "ingress_secret_tls" {
    secret_name= ibm_container_ingress_secret_tls.container_ingress_secret_tls.secret_name
    secret_namespace= ibm_container_ingress_secret_tls.container_ingress_secret_tls.secret_namespace
    cluster = var.cluster_name_or_id
}

# Create an ingress opaque secret
resource "ibm_container_ingress_secret_opaque" "container_ingress_secret_opaque" {
    cluster  = var.cluster_name_or_id
    secret_name = var.opaque_secret_name
    secret_namespace = var.opaque_secret_namespace
    persistence = true
    fields {
        crn = var.field_secret_crn
    }
    fields {
        field_name = var.field_secret_name
        crn = var.field_secret_crn2
    }
}

// Create a ibm_container_ingress_secret_opaque data source
data "ibm_container_ingress_secret_opaque" "ingress_secret_opaque" {
    secret_name= ibm_container_ingress_secret_opaque.container_ingress_secret_opaque.secret_name
    secret_namespace= ibm_container_ingress_secret_opaque.container_ingress_secret_opaque.secret_namespace
    cluster = var.cluster_name_or_id
}