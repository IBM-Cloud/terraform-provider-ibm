# Example for ALL VPC Classic resources

This example illustrates how to use the vpc classic resources to create Classic Infrastructure on IBM Cloud.

These types of resources are supported:

## VPC Classic Resources

* [ibm_is_vpc](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-vps)
* [ibm_is_vpc_route](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-route)
* [ibm_is_vpc_address_prefix](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#address-prefix)
* [ibm_is_subnet](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#subnet)
* [ibm_is_lb](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb)
* [ibm_is_lb_pool](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb-pool)
* [ibm_is_lb_pool_member](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb-pool-member)
* [ibm_is_lb_listener](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb-listener)
* [ibm_is_lb_listener_policy](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb-listener-policy)
* [ibm_is_lb_listener_policy_rule](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#lb-listener-policy-rule)
* [ibm_is_vpn_gateway](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#vpn-gateway)
* [ibm_is_vpn_gateway_connection](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#vpn-gateway-connection)
* [ibm_is_ssh_key](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#ssh-key)
* [ibm_is_instance](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-instance)
* [ibm_is_floating_ip](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-floating-ip)
* [ibm_is_security_group](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#sec-group)
* [ibm_is_security_group_network_interface_attachment](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#sec-group-netint)
* [ibm_is_security_group_rule](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#sec-group-rule)
* [ibm_is_ipsec_policy](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-ipsec)
* [ibm_is_ike_policy](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-ike-policy)
* [ibm_is_volume](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#volume)
* [ibm_is_network_acl](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#network-acl)
* [ibm_is_public_gateway](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources#provider-public-gateway)

##  VPC Classic Data Sources
* [ibm_is_image](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#image)
* [ibm_is_images](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#images)
* [ibm_is_instance_profile](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#instance-profile)
* [ibm_is_instance_profiles](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#instance-profiles)
* [ibm_is_region](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#region)
* [ibm_is_ssh_key](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#ssh-key)
* [ibm_is_subnet](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#subnet)
* [ibm_is_subnets](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#vpc-subnets)
* [ibm_is_vpc](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#vpc)
* [ibm_is_zone](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#zone)
* [ibm_is_zones](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources#zones)

## Terraform versions

Terraform 0.12. Pin module version to `~> v1.5.2`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.25.0`. Branch - `terraform_v0.11.x`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Notes

1. Provider information (Eg: generation, api key, region) is declared in provider.tf
2. The Variables used in main.tf are declared in variables.tf

## Examples

* [VPC Classic Infrastructure Resources](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-resources)
* [VPC Classic Infrastructure Data Sources](https://cloud.ibm.com/docs/terraform?topic=terraform-vpc-gen1-data-sources)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

Single OpenAPI document or directory of documents.

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud_api_key | The IBM Cloud API Key of an account in which the vpc resources are to be provisioned. It can be exported as an environment variable if not declared in the provider. | `string` | no |
| generation | The generation of IBM Cloud VPC infrastructure. Value: `1`| `string` | yes |
| region | The region where the resource has to be provisioned. Default: `us-south`| `string` | yes |
| zone1 | The zone in which route1, addprefix1, subnet1, instance1, vol1, vol2, publicgateway1 are created . Default: `us-south-1`| `string` | yes |
| zone2 | The zone in which subnet2, instance2 are created . Default: `us-south-2`| `string` | yes |
| ssh\_public\_key | The name of the API Gateway Endpoint resource. Default: `~/.ssh/id_rsa.pub`| `string` | yes |
| image | ID of the virtual server Image used in instance1, instance2. Default: `fc538f61-7dd6-4408-978c-c6b85b69fe76` | `string` | yes |
| profile | Name of the profile that is to be associated with instances Default: `b-2x8`| `string` | yes |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
