---
layout: "ibm"
page_title: "IBM : ibm_code_engine_domain_mapping"
description: |-
  Manages code_engine_domain_mapping.
subcategory: "Code Engine"
---

# ibm_code_engine_domain_mapping

Create, update, and delete code_engine_domain_mappings with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_domain_mapping" "code_engine_domain_mapping_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name      = "www.example.com"
  tls_secret = "my-tls-secret"
  component {
		name          = "my-app-1"
		resource_type = "app_v2"
  }
}
```

## Timeouts

code_engine_domain_mapping provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating a code_engine_domain_mapping.
* `update` - (Default 10 minutes) Used for updating a code_engine_domain_mapping.

## Argument Reference

You can specify the following arguments for this resource.

* `component` - (Required, List) A reference to another component.
Nested schema for **component**:
	* `name` - (Required, String) The name of the referenced component.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?$/`.
	* `resource_type` - (Required, String) The type of the referenced resource.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/.+/`.
* `name` - (Required, Forces new resource, String) The name of the domain mapping.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)+$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `tls_secret` - (Required, String) The name of the TLS secret that includes the certificate and private key of this domain mapping.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_domain_mapping.
* `domain_mapping_id` - (String) The identifier of the resource.
    * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `cname_target` - (String) The value of the CNAME record that must be configured in the DNS settings of the domain, to route traffic properly to the target Code Engine region.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
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
* `user_managed` - (Boolean) Specifies whether the domain mapping is managed by the user or by Code Engine.
* `visibility` - (String) Specifies whether the domain mapping is reachable through the public internet, or private IBM network, or only through other components within the same Code Engine project.
  * Constraints: Allowable values are: `custom`, `private`, `project`, `public`.

* `etag` - ETag identifier for code_engine_domain_mapping.

## Import

You can import the `ibm_code_engine_domain_mapping` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `www.example.com`. The name of the domain mapping.

# Syntax
<pre>
$ terraform import ibm_code_engine_domain_mapping.code_engine_domain_mapping &lt;project_id&gt;/&lt;name&gt;
</pre>
