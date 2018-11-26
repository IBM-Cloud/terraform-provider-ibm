---
layout: "ibm"
page_title: "IBM : service_key"
sidebar_current: "docs-ibm-resource-service-key"
description: |-
  Manages IBM Service Key.
---

# ibm\_service_key

Provides a service key resource. This allows service keys to be created, updated, and deleted.

## Example Usage

```hcl
data "ibm_service_instance" "service_instance" {
  name = "mycloudant"
}

resource "ibm_service_key" "serviceKey" {
  name                  = "mycloudantkey"
  service_instance_guid = "${data.ibm_service_instance.service_instance.id}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) A descriptive name used to identify a service key.
* `parameters` - (Optional, map) Arbitrary parameters to pass along to the service broker. Must be a JSON object.
* `service_instance_guid` - (Required, string) The GUID of the service instance associated with the service key.
* `tags` - (Optional, array of strings) Tags associated with the service key instance.  
  **NOTE**: `Tags` are managed locally and not stored on the IBM Cloud service endpoint at this moment.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the new service key.
* `credentials` - The credentials associated with the key.
