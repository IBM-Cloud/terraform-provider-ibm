# IBM Power Examples
This directory contains sample Terraform code to create a Power Systems Virtual Server.
​
## Prerequisites
- An [IBM Cloud Account](https://cloud.ibm.com/registration)
- An IBM Cloud [IAM API key](https://cloud.ibm.com/docs/account?topic=account-userapikey)
- [Terraform](https://www.terraform.io/downloads)
​
## Setup
The following steps are for setting up the example terraform configuration locally:
 - Make a local copy of the files in this directory.
 - Modify the variables in `variables.tf`.
 - To use using a stock image, specified by `image_name` in `variables.tf`, you must first copy the stock image to Cloud Instance. You can copy an image using the [IBM Cloud CLI](https://cloud.ibm.com/docs/cli): `ibmcloud pi image-create <stock-image-id>`.
 - If using a custom or imported image remove `pi_storage_type` from the `ibm_pi_instance` resource since the storage type will default to the custom image's storage type.
  - For developers using a staging environment, export the following environment variables.
```bash
export IBMCLOUD_IAM_API_ENDPOINT="https://iam.test.cloud.ibm.com"
export IBMCLOUD_PI_API_ENDPOINT="<region>.power-iaas.test.cloud.ibm.com"
```
​
## Export vs Define
You can define variables for the provider in `provider.tf` or you can export env variables. More information can be found about exporting env vairables [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#argument-reference).

## Running the Configuration
```bash
# Initalize terraform directory and validate the configuration
terraform init
terraform fmt
terraform validate
​
# Show changes required by the current configuration
terraform plan
​
# Create or update infrastructure
terraform apply
​
# Destroy previously-created infrastructure
terraform destroy
```

## Documentation
 - [IBM Power Systems Virtual Server Docs](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started)
 - [Intro to Terraform](https://www.terraform.io/intro) | [Terraform Overview](https://www.terraform.io/language)
 - [Terraform SDK Docs](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk) (For devs contributing to repo)
​
## IBM Power Systems Terraform Docs
> [Link](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_capture) to IBM Terrafom Docs


| Resource Type | Resource Docs | Data Source Docs |
| :------------ |:------------- | :--------------- |
| Capture | [ibm_pi_capture](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_capture) |  |
| Cloud Connection | [ibm_pi_cloud_connection](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_cloud_connection)<br>[ibm_pi_cloud_connection_network_attach](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_cloud_connection_network_attach) | [ibm_pi_cloud_connection](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_cloud_connection)<br>[ibm_pi_cloud_connections](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_cloud_connections) |
| Cloud Instance |  | [ibm_pi_cloud_instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_cloud_instance) |
| DHCP | [ibm_pi_dhcp](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_dhcp) | [ibm_pi_dhcp](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_dhcp)<br>[ibm_pi_dhcps](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_dhcps) |
| Image | [ibm_pi_image](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_image)<br>[image_export](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_image_export) | [ibm_pi_catalog_images](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_catalog_images)<br>[ibm_pi_image](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_image)<br>[ibm_pi_images](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_images) |
| Instance | [ibm_pi_instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_instance)<br>[ibm_pi_instance_console_language](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_console_language)<br>[ibm_pi_snapshot](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_snapshot) | [ibm_pi_instance](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_instance)<br>[ibm_pi_instances](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_instances)<br>[ibm_pi_instance_console_languages](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_console_languages)<br>[ibm_pi_instance_ip](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_instance_i)<br>[ibm_pi_snapshot](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_pvm_snapshots)<br>[ibm_pi_snapshots](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_pvm_snapshots) |
| Key | [ibm_pi_key](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_key) | [ibm_pi_key](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_key)<br>[ibm_pi_keys](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_keys) |
| Network | [ibm_pi_network](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_network)<br>[ibm_pi_network_port](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_network_port)<br>[ibm_pi_network_port_attach](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_network) | [ibm_pi_network](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_network)<br>[ibm_pi_network_port](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_network)<br>[ibm_pi_public_network](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_public_network) |
| Placement Group | [ibm_pi_placement_group](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_placement_group) | [ibm_pi_placement group](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_placement_group)<br>[ibm_pi_placement_groups](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_placement_groups) |
| SAP | | [ibm_pi_sap_profile](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_sap_profile)<br>[ibm_pi_sap_profiles](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_sap_profiles) |
| Storage Capacity | | [ibm_pi_storage_pool_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_pool_capacity)<br>[ibm_pi_storage_pools_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_pools_capacity)<br>[ibm_pi_system_pools]() |
| Storage Type | | [ibm_pi_storage_type_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_type_capacity)<br>[ibm_pi_storage_types_capacity](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_storage_types_capacity) |
| Tenant | | [ibm_pi_tenant](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_tenant) |
| Volume | [ibm_pi_volume](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_volume)<br>[ibm_pi_volume_attach](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_volume_attach) | [ibm_pi_instance_volumes](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_instance_volumes)<br>[ibm_pi_volume](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/data-sources/pi_volume) |
| VPN | [ibm_pi_ike_policy](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_vpn_ike_policy)<br>[ibm_pi_ipsec_policy](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_vpn_ipsec_policy)<br>[ibm_pi_vpn_connection](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/pi_vpn_connection) |  |
