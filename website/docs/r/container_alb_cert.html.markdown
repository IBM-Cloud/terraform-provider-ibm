---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb_cert"
description: |-
  Manages IBM container alb cert.
---

# ibm\_container_alb_cert

Create, update or delete a Application load balancer certificate. 

## Example Usage

In the following example, you can configure a alb:

```hcl
resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "crn:v1:bluemix:public:cloudcerts:us-south:a/e9021a4dc47e3d:faadea8e-a7f4-408f-8b39-2175ed17ae62:certificate:3f2ab474fbbf9564582"
  secret_name = "test-sec"
  cluster_id  = "myCluster"
}

```

## Timeouts

ibm_container_alb_cert provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `cert_crn` - (Required, string) The certificate CRN.
* `cluster_id` - (Required, Forces new resource, string)  The cluster ID.
* `secret_name` - (Required, Forces new resource, string) The name of the ALB certificate secret. 
* `namespace` - (Optional,Forces new Resource, string) The namespace in which the secret has to be created. Default: `ibm-cert-store`
* `persistence`  - (Optional, bool) Persist the secret data in your cluster. If the secret is later deleted from the CLI or OpenShift web console, the secret is automatically re-created in your cluster.


## Attribute Reference

The following attributes are exported:

* `id` - The ALB cert ID. The id is composed of \<cluster_name_id\>/\<secret_name\>.<br/>
* `domain_name` - The Domain name of the certificate.
* `expires_on` - The Expiry date of the certificate.
* `issuer_name` - The Issuer name of the certificate.
* `cluster_crn` - The cluster crn.
* `status` - The Status of the secret.
* `cloud_cert_instance_id` - Cloud Certificate instance ID from which certificate is downloaded.

## Import

ibm_container_alb_cert can be imported using cluster_id, secret_name eg

```
$ terraform import ibm_container_alb_cert.example 166179849c9a469581f28939874d0c82/mysecret