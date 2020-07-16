# Example for Certificate Manager resources

This example illustrates how to use the Certificate Manager resources to import a certifictae and to order a certificate into the Certificate Manager service instance.

These types of resources are supported:

* [ Import Certificates ](https://cloud.ibm.com/docs/terraform?topic=terraform-cert-manager-resources#cert-manager)
* [ Order Certificates ](https://cloud.ibm.com/docs/terraform?topic=terraform-cert-manager-resources#certmanager-order)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.4.0`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## Certificate Manager Resources

`Import existing Certificates`:

```hcl
resource "ibm_certificate_manager_import" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = var.import_name
  data = {
    content = file(var.cert_file_path)
  }
}
```

`Create ssl certificates using null resource and Import Certificates`:
```hcl

resource "ibm_certificate_manager_import" "cert" {

  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = var.import_name
  data = {
    content  = data.local_file.cert.content
    priv_key = data.local_file.key.content
  }
}

```
`Order Certificates`:
```hcl

resource "ibm_certificate_manager_order" "cert" {
  certificate_manager_instance_id = ibm_resource_instance.cm.id
  name                            = var.order_name
  description                     = var.order_description
  domains                         = [ibm_cis_domain.example.domain]
  rotate_keys                     = var.rotate_key
  domain_validation_method        = var.dvm
  dns_provider_instance_crn       = ibm_cis.instance.id
}

```
##  Certificate Manager Data Source
`List all certificates:`

```hcl

data "ibm_certificate_manager_certificates" "certs"{
    certificate_manager_instance_id=data.ibm_resource_instance.cm.id
}

```

## Assumptions

1. It's assumed that user has valid domain ownership while ordering certificates using IBM CIS.
2. [ Certificate Ordering Limitations ](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-ordering-certificates#certificate-ordering-limitations)
3. Before ordering certificates using IBM CIS [ Set up ordering certificates using CIS ](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-ordering-certificates#cis)
4. [ Certificate Ordering Limits ](https://cloud.ibm.com/docs/certificate-manager?topic=certificate-manager-limits#api-limits)
5. [ API Documentation for CMS ](https://cloud.ibm.com/apidocs/certificate-manager)

## Notes

1. Terraform IBM provider v1.4.0 (via Terraform 0.12) doesn't supports ordering certificates using `Other DNS provider`.
2. With `auto_renew_enabled`, certificates are automatically renewed 31 days before they expire. If your certificate expires in less than 31 days, you must renew it by updating `rotate_keys`. After you do so, your future certificates are renewed automatically.

## Examples

* [ Certificate Manager Import Certificates ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-certificate-manager/ibm-certificate-manager-import)
* [ Certificate Manager Order Certificates ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-certificate-manager/ibm-certificate-manager-order)


<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| region | THe region where the resource has to be provisioned. Default: `us-south`| `string` | yes |
| cms\_name | The name of the Certificate Manager Service Instance. | `string` | yes |
| cis\_name | The name of the CIS Instance resource. | `string` | yes |
| cis\_plan | The Plan of CIS Instance resource. Default: `standard` | `string` | yes |
| domain | Valid CIS domain | `string` | yes |
| order\_name | Name of certificate that has to be orderd.| `string` | yes |
| order\_description | Description of certificate that has to be orderd| `string` | no |
| rotate\_key | Rotate Keys. Default: `false` | `bool` | Required while Renewing certificate. |
| dvm | Domain Validation Method of the CIS Domain. Default: `dns-01` | `string` | yes |
| import\_name | Name of certificate that has to be imported. | `string` | yes |
| cert\_file\_path | Path of the certificate file that has to be imported. | `string` | yes |
| ssl\_region | Region of SSL certificate that is been generated. | `string` | Required while generating a certificate using null resource. |
| host | Host of SSL certificate that is been generated. | `string` | Required while generating a certificate using null resource. |
| ssl\_key | Private Key file name of SSL certificate. Default: `private_key.key` | `string` | Required while generating a certificate using null resource. |
| ssl\_cert | SSL Certificate file name. Default: `certificate.pem` | `string` | Required while generating a certificate using null resource. |

## Outputs

| Name | Description |
|------|-------------|
| cert_order_id | ID of the ordered Certificate |
| expires_on | Indicates when the ordered certificate expires. |
| cert_import_id | ID of the Imported Certificate |
| cert_import_content | Content of Imported Certificate. |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
