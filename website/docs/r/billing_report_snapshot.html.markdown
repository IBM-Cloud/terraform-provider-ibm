---
layout: "ibm"
page_title: "IBM : ibm_billing_report_snapshot"
description: |-
  Manages billing_report_snapshot.
subcategory: "Usage Reports"
---

# ibm_billing_report_snapshot

Create, update, and delete billing_report_snapshots with this resource.

## Example Usage

```hcl
resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
  cos_bucket = "bucket_name"
  cos_location = "us-south"
  cos_reports_folder = "IBMCloud-Billing-Reports"
  interval = "daily"
  versioning = "new"
}
```

## Example usage with service-to-service authorization

```hcl
resource "ibm_iam_authorization_policy" "policy" {
  source_service_name         = "billing"
  target_service_name         = "cloud-object-storage"
  target_resource_instance_id = "cos_instance_id"
  roles                       = ["Object Writer", "Content Reader"]
}

resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
  cos_bucket = "bucket_name"
  cos_location = "us-south"
  cos_reports_folder = "IBMCloud-Billing-Reports"
  interval = "daily"
  versioning = "new"
  report_types = ["account_summary", "account_resource_instance_usage"]
  depends_on = [ ibm_iam_authorization_policy.policy ]
}
```
If service-to-service authorization already exists in the specific COS bucket, the resource `policy` can be omitted and the `depends_on` flag can also be removed from resource `billing_report_snapshot_instance`. For more information, about IAM service authorizations, see [using authorizations to grant access between services](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/iam_authorization_policy).

## Argument Reference

You can specify the following arguments for this resource.

* `cos_bucket` - (Required, Forces new resource, String) The name of the COS bucket to store the snapshot of the billing reports.
* `cos_location` - (Required, Forces new resource, String) Region of the COS instance.
* `cos_reports_folder` - (Optional, Forces new resource, String) The billing reports root folder to store the billing reports snapshots. Defaults to "IBMCloud-Billing-Reports".
  * Constraints: The default value is `IBMCloud-Billing-Reports`.
* `interval` - (Required, Forces new resource, String) Frequency of taking the snapshot of the billing reports.
  * Constraints: Allowable values are: `daily`.
* `report_types` - (Optional, Forces new resource, List) The type of billing reports to take snapshot of. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage].
  * Constraints: Allowable list items are: `account_summary`, `enterprise_summary`, `account_resource_instance_usage`.
* `versioning` - (Optional, Forces new resource, String) A new version of report is created or the existing report version is overwritten with every update.
  * Constraints: The default value is `new`. Allowable values are: `new`, `overwrite`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the billing_report_snapshot.
* `account_type` - (String) Type of account. Possible values are [enterprise, account].
  * Constraints: Allowable values are: `account`, `enterprise`.
* `compression` - (String) Compression format of the snapshot report.
* `content_type` - (String) Type of content stored in snapshot report.
* `cos_endpoint` - (String) The endpoint of the COS instance.
* `created_at` - (Integer) Timestamp in milliseconds when the snapshot configuration was created.
* `history` - (List) List of previous versions of the snapshot configurations.
Nested schema for **history**:
	* `account_id` - (String) Account ID for which billing report snapshot is configured.
	* `account_type` - (String) Type of account. Possible values [enterprise, account].
	  * Constraints: Allowable values are: `account`, `enterprise`.
	* `compression` - (String) Compression format of the snapshot report.
	* `content_type` - (String) Type of content stored in snapshot report.
	* `cos_bucket` - (String) The name of the COS bucket to store the snapshot of the billing reports.
	* `cos_endpoint` - (String) The endpoint of the COS instance.
	* `cos_location` - (String) Region of the COS instance.
	* `cos_reports_folder` - (String) The billing reports root folder to store the billing reports snapshots. Defaults to "IBMCloud-Billing-Reports".
	  * Constraints: The default value is `IBMCloud-Billing-Reports`.
	* `end_time` - (Integer) Timestamp in milliseconds when the snapshot configuration ends.
	* `interval` - (String) Frequency of taking the snapshot of the billing reports.
	  * Constraints: Allowable values are: `daily`.
	* `report_types` - (List) The type of billing reports to take snapshot of. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage].
	  * Constraints: Allowable list items are: `account_summary`, `enterprise_summary`, `account_resource_instance_usage`.
	* `start_time` - (Integer) Timestamp in milliseconds when the snapshot configuration was created.
	* `state` - (String) Status of the billing snapshot configuration. Possible values are [enabled, disabled].
	  * Constraints: Allowable values are: `enabled`, `disabled`.
	* `updated_by` - (String) Account that updated the billing snapshot configuration.
	* `versioning` - (String) A new version of report is created or the existing report version is overwritten with every update.
	  * Constraints: The default value is `new`. Allowable values are: `new`, `overwrite`.
* `last_updated_at` - (Integer) Timestamp in milliseconds when the snapshot configuration was last updated.
* `state` - (String) Status of the billing snapshot configuration. Possible values are [enabled, disabled].
  * Constraints: Allowable values are: `enabled`, `disabled`.


## Import

You can import the `ibm_billing_report_snapshot` resource by using `account_id`. Account ID for which billing report snapshot is configured.

# Syntax
<pre>
$ terraform import ibm_billing_report_snapshot.billing_report_snapshot &lt;account_id&gt;
</pre>

# Example
```
$ terraform import ibm_billing_report_snapshot.billing_report_snapshot abc
```
