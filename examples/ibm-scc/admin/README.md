# Example for AdminServiceApiV1

This example illustrates how to use the AdminServiceApiV1

These types of resources are supported:


## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## AdminServiceApiV1 resources
```hcl
resource "ibm_scc_account_settings" "ibm_scc_account_settings_instance" {
  location_id = var.ibm_scc_account_settings_location_id
}
```

## AdminServiceApiV1 Data sources

scc_account_location_settings data source:

```hcl
data "scc_account_location_settings" "scc_account_location_settings_instance" {
}
```
scc_account_location data source:

```hcl
data "scc_account_location" "scc_account_location_instance" {
  location_id = var.scc_account_location_location_id
}
```
scc_account_locations data source:

```hcl
data "scc_account_locations" "scc_account_locations_instance" {
}
```

## Assumptions

- The default location has already been set for you

## Notes

- Running `terraform apply` will output the location your account is operating in as well as the available locations for your Security and Compliance center

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| location_id | The programatic ID of the location that you want to work in. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| available_locations | The available Security and Compliance Center locations |
| location_details | The details of a given location |
| current_location_settings_details | The details of the current account settings |
