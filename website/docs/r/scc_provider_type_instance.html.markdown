---
layout: "ibm"
page_title: "IBM : ibm_scc_provider_type_instance"
description: |-
  Manages scc_provider_type_instance.
subcategory: "Security and Compliance Center"
---

# ibm_scc_provider_type_instance

Create, update, and delete provider type instances with this resource.

## Example Usage

```hcl
resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
  attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
  name = "workload-protection-instance-1"
  provider_type_id = "provider_type_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `attributes` - (Required, Map) The attributes for connecting to the provider type instance.
* `name` - (Required, String) The name for the provider_type instance
* `provider_type_id` - (Required, String) The unique identifier of the provider type instance.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_provider_type_instance.
* `created_at` - (String) The time when resource was created.
* `type` - (String) The type of the provider type.
* `updated_at` - (String) The time when resource was updated.


## Import

You can import the `ibm_scc_provider_type_instance` resource by using `id`.
The `id` property can be formed from `provider_type_id`, and `provider_type_instance_id` in the following format:

```
<provider_type_id>/<provider_type_instance_id>
```
* `provider_type_id`: A string. The provider type ID.
* `provider_type_instance_id`: A string. The provider type instance ID.

# Syntax
```
$ terraform import ibm_scc_provider_type_instance.scc_provider_type_instance <provider_type_id>/<provider_type_instance_id>
```
