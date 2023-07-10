# IBM Cloud Kubernetes Service Ingress Instance and Secrets Example

This example shows how to register an existing IBM Cloud Secrets Manager Instance with an IBM Cloud Kubernetes Service or Redhat Openshift cluster.
Kubernetes secrets of type TLS or opaque can than be created in the corresponding cluster using a Secrets Manager CRN.  

## Usage

To run this example you need to execute:

```sh
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you want to deregister the instance.

## Example usage

Register a Secrets Manager Instance with your cluster:

```terraform
resource "ibm_container_ingress_instance" "cluster_instance" {
  instance_crn = var.sm_instance_crn
  is_default = var.default_instance
  cluster  = var.cluster_name_or_id
  secret_group_id = var.sm_secret_group_id
}
```

```terraform
data "ibm_container_ingress_instance" "ingress_instance" {
    instance_name = ibm_container_ingress_instance.cluster_instance.instance_name
    cluster = var.cluster_name_or_id
}
```

```terraform
resource "ibm_container_ingress_secret_tls" "container_ingress_secret_tls" {
    cluster  = var.cluster_name_or_id
    secret_name = var.tls_secret_name
    secret_namespace = var.tls_secret_namespace
    cert_crn = var.secret_cert_crn
    persistence = true
}
```

```terraform
// Create a ibm_container_ingress_secret_tls data source
data "ibm_container_ingress_secret_tls" "ingress_secret_tls" {
    secret_name= ibm_container_ingress_secret_tls.container_ingress_secret_tls.secret_name
    secret_namespace= ibm_container_ingress_secret_tls.container_ingress_secret_tls.secret_namespace
    cluster = var.cluster_name_or_id
}
```

```terraform
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
```

```terraform
// Create a ibm_container_ingress_secret_opaque data source
data "ibm_container_ingress_secret_opaque" "ingress_secret_opaque" {
    secret_name= ibm_container_ingress_secret_opaque.container_ingress_secret_opaque.secret_name
    secret_namespace= ibm_container_ingress_secret_opaque.container_ingress_secret_opaque.secret_namespace
    cluster = var.cluster_name_or_id
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
  required_version = ">= 0.12"

## Providers

| Name | Version |
|------|---------|
| ibm  | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_crn | CRN for the Secrets Manager instance | `string` | yes |
| cluster | The cluster id or name to register the instance to or secrets in | `string` | yes |
| secret_group_id | The Secrets Manager secret group that default TLS certs will be uploaded to | `string` | no |
| is_default | Whether the registered instance will be used for storing the default TLS certificates for the cluster | `bool` | no |
| secret_name | The name of the Kubernetes secret that will be created | `string` | yes |
| secret_namespace | The namespace of the Kubernetes secret that will be created | `string` | yes |
| cert_crn | CRN of a Secrets Manager secret of type certificate | `string` | yes |
| field_name | The name of the Kubernetes secret field that will be created | `string` | yes |
| crn | CRN of a Secrets Manager secret for an opaque secret field | `string` | yes |
{: caption="inputs"}