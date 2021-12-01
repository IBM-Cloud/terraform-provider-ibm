# This Module is used to register public IP address to satellite location & cluster dns records.

This module register public IP address to control plane & openshift cluster subdomain DNS records.
 
## Usage

```
terraform init
```
```
terraform plan
```
```
terraform apply
```
```
terraform destroy
```
## Example Usage

``` hcl
module "satellite-dns" {
  source = "./modules/dns"

  location          = var.location
  cluster           = var.cluster
  control_plane_ips = var.control_plane_ips
  cluster_ips       = var.cluster_ips
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Inputs

| Name              | Description                                                       | Type     | Default | Required |
|-------------------|-------------------------------------------------------------------|----------|---------|----------|
| location          | Name of the Location                                              | string   | n/a     | yes      |
| cluster           | Name of the cluster                                               | string   | n/a     | yes      |
| control_plane_ips | Public IP addresses                                               | list     | n/a     | yes      |
| cluster_ips       | Public IP addresses                                               | list     | n/a     | yes      |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Note

All optional fields are given value `null` in varaible.tf file. User can configure the same by overwriting with appropriate values.

