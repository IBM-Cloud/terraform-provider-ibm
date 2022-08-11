---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_alb_cert"
description: |-
  Manages IBM container Application Load Balancer certificate.
---

# ibm_container_alb_cert
Create, update, or delete an SSL certificate that you store in IBM Cloud Certificate Manager for an Ingress Application Load Balancer (ALB). For more information, about container ALB certificate, see [setting up Kubernetes Ingress](https://cloud.ibm.com/docs/containers?topic=containers-ingress-types).

## Example usage
The following example adds an SSL certificate that is stored in IBM Cloud Certificate Manager to an Ingress ALB that is set up in a cluster that is named `myCluster`. 

```terraform
resource "ibm_container_alb_cert" "cert" {
  cert_crn    = "crn:v1:bluemix:public:cloudcerts:us-south:a/e9021a4dc47e3d:faadea8e-a7f4-408f-8b39-2175ed17ae62:certificate:3f2ab474fbbf9564582"
  secret_name = "test-sec"
  cluster_id  = "myCluster"
}

```

## Timeouts
The `ibm_container_alb_cert` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create**: The creation of the SSL certificate is considered `failed` if no response is received for 10 minutes.
- **Delete**: The deletion of the SSL certificate is considered `failed` if no response is received for 10 minutes.
- **Update**: The update of the SSL certificate is considered `failed` if no response is received for 10 minutes.


## Argument reference
Review the argument references that you can specify for your resource. 

- `cert_crn` - (Required, String) The CRN of the certificate that you uploaded to IBM Cloud Certificate Manager.
- `cluster_id` - (Required, Forces new resource, String) The ID of the cluster that hosts the Ingress ALB that you want to configure for SSL traffic.
- `secret_name` - (Required, Forces new resource, String) The name of the ALB certificate secret.
- `namespace` - (String)  Optional- The namespace in which the secret is created. Default value is `ibm-cert-store`.
- `persistence`-(Optional, Bool) Persist the secret data in your cluster. If the secret is later deleted from the command line or OpenShift web console, the secret is automatically re-created in your cluster.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `cluster_crn` - (String) The CRN of the cluster that hosts the Ingress ALB.
- `cloud_cert_instance_id` - (String) The IBM Cloud Certificate Manager instance ID from which the certificate was downloaded.
- `domain_name` - (String) The domain name of the certificate.
- `expires_on` - Date - The date the certificate expires.
- `id` - (String) The unique identifier of the certificate in the format `<cluster_name_id>/<secret_name>`.
- `issuer_name` - (String) The name of the issuer of the certificate. 
- `status` - (String) The Status of the secret.

## Import
The `ibm_container_alb_cert` can be imported by using cluster_id, and secret_name.

**Example**

```
$ terraform import ibm_container_alb_cert.example 166179849c9a469581f28939874d0c82/mysecret
```
