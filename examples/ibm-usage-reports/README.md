# Example for UsageReportsV4

This example illustrates how to use the UsageReportsV4

The following types of resources are supported:

* billing_report_snapshot

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## UsageReportsV4 resources

billing_report_snapshot resource:

```hcl
resource "billing_report_snapshot" "billing_report_snapshot_instance" {
  account_id = var.billing_report_snapshot_account_id
  interval = var.billing_report_snapshot_interval
  versioning = var.billing_report_snapshot_versioning
  report_types = var.billing_report_snapshot_report_types
  cos_reports_folder = var.billing_report_snapshot_cos_reports_folder
  cos_bucket = var.billing_report_snapshot_cos_bucket
  cos_location = var.billing_report_snapshot_cos_location
}
```

## UsageReportsV4 data sources

billing_snapshot_list data source:

```hcl
data "billing_snapshot_list" "billing_snapshot_list_instance" {
  account_id = var.billing_snapshot_list_account_id
  month = var.billing_snapshot_list_month
  date_from = var.billing_snapshot_list_date_from
  date_to = var.billing_snapshot_list_date_to
}
```

## Assumptions

1. TODO

## Notes

1. TODO

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
| account_id | Account ID for which billing report snapshot is configured. | `string` | true |
| interval | Frequency of taking the snapshot of the billing reports. | `string` | true |
| versioning | A new version of report is created or the existing report version is overwritten with every update. | `string` | false |
| report_types | The type of billing reports to take snapshot of. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage]. | `list(string)` | false |
| cos_reports_folder | The billing reports root folder to store the billing reports snapshots. Defaults to "IBMCloud-Billing-Reports". | `string` | false |
| cos_bucket | The name of the COS bucket to store the snapshot of the billing reports. | `string` | true |
| cos_location | Region of the COS instance. | `string` | true |
| account_id | Account ID for which the billing report snapshot is requested. | `string` | true |
| month | The month for which billing report snapshot is requested.  Format is yyyy-mm. | `string` | true |
| date_from | Timestamp in milliseconds for which billing report snapshot is requested. | `number` | false |
| date_to | Timestamp in milliseconds for which billing report snapshot is requested. | `number` | false |

## Outputs

| Name | Description |
|------|-------------|
| billing_report_snapshot | billing_report_snapshot object |
| billing_snapshot_list | billing_snapshot_list object |
