---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_profile_import"
description: |-
  To Import profile.
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_profile_import

Provides a resource for profiles. This allows profiles to be import a profile that you formatted locally.

## Example Usage

```hcl
resource "ibm_scc_posture_profile_import" "profiles" {
  file      = "/local_path_to_import_profile.csv"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `file` - (Required, Forces new resource, String) The import data file that you want to use to import a profile.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the profiles.

## Import

You can import the `ibm_scc_posture_profile_import` resource by using `id`. An identifier of the profile.

# Syntax
```
$ terraform import ibm_scc_posture_profile_import.profiles <id>
```

# Example
```
$ terraform import ibm_scc_posture_profile_import.profiles 1
```
!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_profile_import is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
