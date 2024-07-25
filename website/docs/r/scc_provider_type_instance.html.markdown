---
layout: "ibm"
page_title: "IBM : ibm_scc_provider_type_instance"
description: |-
  Manages scc_provider_type_instance.
subcategory: "Security and Compliance Center"
---

# ibm_scc_provider_type_instance

Create, update, and delete provider type instances with this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
resource "ibm_scc_provider_type_instance" "scc_provider_type_instance_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  attributes = {"wp_crn":"crn:v1:staging:public:sysdig-secure:eu-gb:a/14q5SEnVIbwxzvP4AWPCjr2dJg5BAvPb:d1461d1ae-df1eee12fa81812e0-12-aa259::"}
  name = "workload-protection-instance-1"
  provider_type_id = "provider_type_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `attributes` - (Required, Map) The attributes for connecting to the provider type instance.
* `name` - (Required, String) The name for the provider_type instance
* `provider_type_id` - (Required, String) The unique identifier of the provider type instance.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_provider_type_instance.
* `provider_type_instance_id` - (String) The ID that is associated with the created `provider_type_instance`
* `created_at` - (String) The time when resource was created.
* `type` - (String) The type of the provider type.
* `updated_at` - (String) The time when resource was updated.


## Import

You can import the `ibm_scc_provider_type_instance` resource by using `id`.
The `id` property can be formed from `instance_id`, `provider_type_id`, and `provider_type_instance_id` in the following format:

```bash
<instance_id>/<provider_type_id>/<provider_type_instance_id>
```
* `instance_id`: A string. The instance ID.
* `provider_type_id`: A string. The provider type ID.
* `provider_type_instance_id`: A string. The provider type instance ID.

# Syntax

```bash
$ terraform import ibm_scc_provider_type_instance.scc_provider_type_instance <instance_id>/<provider_type_id>/<provider_type_instance_id>
```

# Example
```bash
$ terraform import ibm_scc_provider_type_instance.scc_provider_type_instance 00000000-1111-2222-3333-444444444444/00000000-1111-2222-3333-444444444444/f3517159-889e-4781-819a-89d89b747c85
```