# Examples for Platform Notifications

These examples illustrate how to use the resources and data sources associated with Platform Notifications.

The following resources are supported:
* ibm_notification_distribution_list_destination

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Platform Notifications resources

### Resource: ibm_notification_distribution_list_destination

```hcl
resource "ibm_notification_distribution_list_destination" "notification_distribution_list_destination_instance" {
  account_id = var.notification_distribution_list_destination_account_id
  destination_id = var.notification_distribution_list_destination_destination_id
  destination_type = var.notification_distribution_list_destination_destination_type
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| account_id | The IBM Cloud account ID. | `string` | true |
| destination_id | The GUID of the Event Notifications instance. | `` | false |
| destination_type | The type of the destination. | `string` | true |

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |
