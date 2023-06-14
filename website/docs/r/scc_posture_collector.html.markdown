---
layout: "ibm"
page_title: "IBM : ibm_scc_posture_collector"
description: |-
  Manages collectors.
subcategory: "Security and Compliance Center"
---

# ibm_scc_posture_collector

Provides a resource for collectors. This allows collectors to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_scc_posture_collector" "collectors" {
  description = "sample collector"
  is_public = true
  managed_by = "ibm"
  name = "IBM-collector-sample"
  passphrase = "secret"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `description` - (Optional, String) A detailed description of the collector.
  * Constraints: The default value is ``. The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\._,\\s]*$/`.
  !> **Removal Notification** Resource Removal: Resource ibm_scc_posture_collector is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
* `is_public` - (Required, Boolean) Determines whether the collector endpoint is accessible on a public network. If set to `true`, the collector connects to resources in your account over a public network. If set to `false`, the collector connects to resources by using a private IP that is accessible only through the IBM Cloud private network.
* `is_ubi_image` - (Optional, Boolean) Determines whether the collector has a Ubi image.
* `managed_by` - (Required, String) Determines whether the collector is an IBM or customer-managed virtual machine. Use `ibm` to allow Security and Compliance Center to create, install, and manage the collector on your behalf. The collector is installed in an OpenShift cluster and approved automatically for use. Use `customer` if you would like to install the collector by using your own virtual machine. For more information, check out the [docs](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-collector).
  * Constraints: Allowable values are: `ibm`, `customer`.
* `name` - (Required, String) A unique name for your collector.
  * Constraints: The maximum length is `32` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.,_\\s]*$/`.
* `passphrase` - (Optional, String) To protect the credentials that you add to the service, a passphrase is used to generate a data encryption key. The key is used to securely store your credentials and prevent anyone from accessing them.
  * Constraints: The maximum length is `200` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\._,\\s]*$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the collectors.

## Import

You can import the `ibm_scc_posture_collector` resource by using `id`. An identifier of the collector.

# Syntax
```
$ terraform import ibm_scc_posture_collector.collectors <id>
```

# Example
```
$ terraform import ibm_scc_posture_collector.collectors 1
```

!> **Removal Notification** Resource Removal: Resource ibm_scc_posture_collector is deprecated and being removed.\n This resource will not be available from future release (v1.54.0).
