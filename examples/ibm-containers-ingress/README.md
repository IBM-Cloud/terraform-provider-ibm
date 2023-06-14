# IBM Cloud Kubernetes Service Ingress Instance Example

This example shows how to register an existing IBM Cloud Secrets Manager Instance with an IBM Cloud Kubernetes Service or Redhat Openshift cluster.

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
  is_default = true
  cluster  = var.cluster_name_or_id
}
```

```terraform
data "ibm_container_ingress_instance" "ingress_instance" {
    instance_name = ibm_container_ingress_instance.cluster_instance.instance_name
    cluster = var.cluster_name_or_id
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm  | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| instance_crn | CRN for the Secrets Manager instance | `string` | yes |
| cluster | The clusterID or name to register the instance to | `string` | yes |
| secret_group_id | The Secrets Manager secret group that default TLS certs will be uploaded to | `string` | no |
| is_default | Whether the registered instance will be used for storing the default TLS certificates for the cluster | `string` | no |
{: caption="inputs"}

