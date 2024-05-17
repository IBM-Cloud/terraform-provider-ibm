---
layout: "ibm"
page_title: "IBM : ibm_billing_snapshot_list"
description: |-
  Get information about billing_snapshot_list
subcategory: "Usage Reports"
---

# ibm_billing_snapshot_list

Provides a read-only data source to retrieve information about a billing_snapshot_list. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_billing_snapshot_list" "billing_snapshot_list" {
	date_from = 1675209600000
	date_to = 1675987200000
	month = "2023-02"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `date_from` - (Optional, Integer) Timestamp in milliseconds for which billing report snapshot is requested.
* `date_to` - (Optional, Integer) Timestamp in milliseconds for which billing report snapshot is requested.
* `limit` - (Optional, Integer) Number of usage records returned. The default value is 30. Maximum value is 200.
  * Constraints: The default value is `30`. The maximum value is `200`. The minimum value is `1`.
* `month` - (Required, String) The month for which billing report snapshot is requested.  Format is yyyy-mm.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the billing_snapshot_list.
* `count` - (Integer) Number of total snapshots.

* `snapshots` - (List) 
Nested schema for **snapshots**:
	* `account_id` - (String) Account ID for which billing report snapshot is configured.
	* `account_type` - (String) Type of account. Possible values are [enterprise, account].
	  * Constraints: Allowable values are: `account`, `enterprise`.
	* `billing_period` - (List) Period of billing in snapshot.
	Nested schema for **billing_period**:
		* `end` - (String) Date and time of end of billing in the respective snapshot.
		* `start` - (String) Date and time of start of billing in the respective snapshot.
	* `bucket` - (String) The name of the COS bucket to store the snapshot of the billing reports.
	* `charset` - (String) Character encoding used.
	* `compression` - (String) Compression format of the snapshot report.
	* `content_type` - (String) Type of content stored in snapshot report.
	* `created_on` - (String) Date and time of creation of snapshot.
	* `expected_processed_at` - (Integer) Timestamp of snapshot processed.
	* `files` - (List) List of location of reports.
	Nested schema for **files**:
		* `account_id` - (String) Account ID for which billing report is captured.
		* `location` - (String) Absolute path of the billing report in the COS instance.
		* `report_types` - (String) The type of billing report stored. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage].
		  * Constraints: Allowable values are: `account_summary`, `enterprise_summary`, `account_resource_instance_usage`.
	* `month` - (String) Month of captured snapshot.
	  * Constraints: The value must match regular expression `/^\\d{4}\\-(0?[1-9]|1[012])$/`.
	* `processed_at` - (Integer) Timestamp at which snapshot is captured.
	* `report_types` - (List) List of report types configured for the snapshot.
	Nested schema for **report_types**:
		* `type` - (String) The type of billing report of the snapshot. Possible values are [account_summary, enterprise_summary, account_resource_instance_usage].
		  * Constraints: Allowable values are: `account_summary`, `enterprise_summary`, `account_resource_instance_usage`.
		* `version` - (String) Version of the snapshot.
	* `snapshot_id` - (String) Id of the snapshot captured.
	* `state` - (String) Status of the billing snapshot configuration. Possible values are [enabled, disabled].
	  * Constraints: Allowable values are: `enabled`, `disabled`.
	* `version` - (String) Version of the snapshot.

