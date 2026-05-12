---
layout: "ibm"
page_title: "IBM : ibm_code_engine_domain_mapping"
description: |-
  Get information about code_engine_domain_mapping
subcategory: "Code Engine"
---

# ibm_code_engine_domain_mapping

Provides a read-only data source to retrieve information about a code_engine_domain_mapping. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_domain_mapping" "code_engine_domain_mapping" {
  project_id = data.ibm_code_engine_project.code_engine_project.project_id
  name       = "my-domain-mapping"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your domain mapping.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)+$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_domain_mapping.

* `domain_mapping_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `cname_target` - (String) The value of the CNAME record that must be configured in the DNS settings of the domain, to route traffic properly to the target Code Engine region.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `component` - (List) A reference to another component.
Nested schema for **component**:
	* `name` - (String) The name of the referenced component.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
	* `resource_type` - (String) The type of the referenced resource.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.

* `created_at` - (String) The timestamp when the resource was created.

* `entity_tag` - (String) The version of the domain mapping instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `href` - (String) When you provision a new domain mapping, a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.

* `resource_type` - (String) The type of the Code Engine resource.
  * Constraints: Allowable values are: `domain_mapping_v2`.

* `status` - (String) The current status of the domain mapping.
  * Constraints: Allowable values are: `ready`, `failed`, `deploying`.

* `status_details` - (List) The detailed status of the domain mapping.
Nested schema for **status_details**:
	* `reason` - (String) Optional information to provide more context in case of a 'failed' or 'warning' status.
	  * Constraints: Allowable values are: `ready`, `domain_already_claimed`, `app_ref_failed`, `failed_reconcile_ingress`, `deploying`, `failed`.

* `tls_secret` - (String) The name of the TLS secret that includes the certificate and private key of this domain mapping.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

* `user_managed` - (Boolean) Specifies whether the domain mapping is managed by the user or by Code Engine.

* `visibility` - (String) Specifies whether the domain mapping is reachable through the public internet, or private IBM network, or only through other components within the same Code Engine project.
  * Constraints: Allowable values are: `custom`, `private`, `project`, `public`.

