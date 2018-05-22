---
layout: "ibm"
page_title: "IBM : service_id"
sidebar_current: "docs-ibm-datasource-service-id"
description: |-
  Manages IBM IAM Service ID.
---

# ibm\_service_id

Import the details of an IAM (Identity and Access Management) servicID  on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_iam_service_id" "ds_serviceID" {
  name = "sample"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) Name of the serviceID.

## Attribute Reference

The following attributes are exported:

* `service_ids` - A nested block list of IAM ServiceIDs. Nested `service_ids` blocks have the following structure:
  * `id` - The unique identifier of the serviceID.
  * `bound_to` -  bound to of the serviceID.
  * `crn` -  crn of the serviceID.
  * `description` -  description of the serviceID.
  * `version` -  version of the serviceID.
  * `locked` -  lock state of the serviceID.

  
