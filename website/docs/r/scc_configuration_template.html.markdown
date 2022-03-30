---
layout: "ibm"
page_title: "IBM : ibm_scc_template"
description: |-
  Manages scc_template.
subcategory: "Security and Compliance Center"
---

# ibm_scc_template

Provides a resource for scc_template. This allows scc_template to be created, updated and deleted. For more information about Security and Compliance Center templates, see [Defining Templates](https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-templates-define&interface=ui).

## Example Usage

```hcl
resource "ibm_scc_template" "scc_template_tf_example" {
  account_id  = var.account_id
  name        = "Terraform template"
  description = "COS buckets in us-south should send write data events to Activity Tracker by default"
  target {
    service_name  = "cloud-object-storage"
    resource_kind = "bucket"
    additional_target_attributes {
      name  = "location"
      value = "us-south"
    }
  }
  customized_defaults {
    property = "activity_tracking.write_data_events"
    value    = "true"
  }
}
```

In the above example, new COS buckets in `us-south` will send write data events to Activity Tracker by default.

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Required, String) Your IBM Cloud account ID. 
* `name` - (Required, String) A human-readable alias to assign to your template.
    * Constraints: The maximum length is `32` characters. The minimum length is `1` character.
* `description` - (Required, String) An extended description of your template.
    * Constraints: The maximum length is `256` characters. The minimum length is `1` character.
* `target` - (Required, List) The properties that describe the resource that you want to target with the rule or template. 
    Nested scheme for ** target**:
  * `service_name` - (Required, String) The programmatic name of the IBM Cloud service that you want to target with the
    rule or template. 
    * Constraints: The value must match regular expression `/^[a-z-]*$/`.
  * `resource_kind` - (Required, String) The type of resource that you want to target.
* `additional_target_attributes` - (Optional, List) An extra qualifier for the resource kind. When you include additional attributes, only the resources that match the definition are included in the rule or template. 
    * Nested scheme for **additional_target_attributes**:
      * `name` - (Required, String) The name of the additional attribute that you want to use to further qualify the target. Options differ depending on the service or resource that you are targeting with a rule or template. For more information, refer to the service documentation. 
      * `value` - (Required, String) The value that you want to apply to `name` field. Options differ depending on the rule or template that you configure. For more information, refer to the service documentation. 
* `customized_defaults` - (Required, List) A list of default property values to apply to your template. 
    Nested scheme for **customized_defaults**:
  * `property` - (Required, String) The name of the resource property that you want to configure. Property options differ depending on the service or resource that you are targeting with a template. To view a list of properties that are compatible with templates, refer to the service documentation. 
  * `value` - (Required, String) The custom value that you want to apply as the default for the resource property in the `name` field. This value is used to override the default value that is provided by IBM when a resource is created. Value options differ depending on the resource that you are configuring. To learn more about your options, refer to the service documentation.


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the ibm_scc_template.
* `created_by` - (Optional, String) The unique identifier for the user or application that created the resource.
* `creation_date` - (Optional, String) The date the resource was created.
* `modification_date` - (Optional, String) The date the resource was last modified.
* `modified_by` - (Optional, String) The unique identifier for the user or application that last modified the resource.
* `version` - Version of the ibm_scc_template.

## Import

You can import the `ibm_scc_template` resource by using `template_id`. The UUID that uniquely identifies the template.

# Syntax
```
$ terraform import ibm_scc_template.scc_template <template_id>
```

# Example
```
$ terraform import ibm_scc_template.scc_template template-702d1db7-ca4a-414b-8464-2b517a065c14
```
