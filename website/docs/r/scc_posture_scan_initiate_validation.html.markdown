---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_scan_initiate_validation"
description: |-
  To initiate a validation scan.
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_scan_initiate_validation

Provides a resource for initiate a validation scan. This allows to initiate validation scans determine a specified scope's adherence to regulatory controls by validating the configuration of the resources in your scope to the attached profile.

## Example Usage

```hcl
resource "ibm_scc_posture_scan_initiate_validation" "scanInitiateValidation" {
  scope_id = "scope_id"
  profile_id = "profile_id"
  group_profile_id = "group_profile_id"
  name = "name"
  description = "description"
  frequency = 1
  no_of_occurrences = 1
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `scope_id` - (Required, Forces new resource, String) - The unique ID of the scope.
* `profile_id` - (Required, Forces new resource, String) - The unique ID of the profile.
* `group_profile_id` - (Optional, String) - The ID of the profile group.
* `name` - (Optional, String) - The name of a scheduled scan.
* `description` - (Optional, String) - The description of a scheduled scan.
* `frequency` - (Optional, int) - The frequency at which a scan is run specified in milliseconds.
* `no_of_occurrences` - (Optional, int) - The number of times that a scan should be run.
* `end_time` - (Optional, String) - The date on which a scan should stop running specified in UTC.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The correlation identifier of the initiated validation scan.

## Import

You can import the `ibm_scc_posture_scan_initiate_validation` resource by using `id`. An identifier of the initiated validation scan.

# Syntax
```
$ terraform import ibm_scc_posture_scan_initiate_validation.scans <id>
```

# Example
```
$ terraform import ibm_scc_posture_scan_initiate_validation.scans 1
```
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_scan_initiate_validation is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
