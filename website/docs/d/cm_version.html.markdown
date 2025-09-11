---
layout: "ibm"
page_title: "IBM : ibm_cm_version"
description: |-
  Get information about ibm_cm_version
subcategory: "Catalog Management"
---

# ibm_cm_version

Provides a read-only data source for ibm_cm_version. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cm_version" "cm_version" {
	version_loc_id = "version_loc_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `version_loc_id` - (Required, String) The version locator. A dotted value of `catalogID`.`versionID`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the ibm_cm_version.
* `catalog_id` - (String) Catalog ID.

* `configuration` - (List) List of user solicited overrides.
Nested scheme for **configuration**:
	* `custom_config` - (List) Render type.
	Nested scheme for **custom_config**:
		* `associations` - (List) List of parameters that are associated with this configuration.
		Nested scheme for **associations**:
			* `parameters` - (List) Parameters for this association.
			Nested scheme for **parameters**:
				* `name` - (String) Name of this parameter.
				* `options_refresh` - (Boolean) Refresh options.
		* `config_constraints` - (Map) Map of constraint parameters that will be passed to the custom widget.
		* `grouping` - (String) Determines where this configuration type is rendered (3 sections today - Target, Resource, and Deployment).
		* `grouping_index` - (Integer) Determines the order that this configuration item shows in that particular grouping.
		* `original_grouping` - (String) Original grouping type for this configuration (3 types - Target, Resource, and Deployment).
		* `type` - (String) ID of the widget type.
	* `default_value` - (Map) The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.
	* `description` - (String) Key description.
	* `display_name` - (String) Display name for configuration type.
	* `hidden` - (Boolean) Hide values.
	* `key` - (String) Configuration key.
	* `options` - (List) List of options of type.
	* `required` - (Boolean) Is key required to install.
	* `type` - (String) Value type (string, boolean, int).
	* `type_metadata` - (String) The original type, as found in the source being onboarded.
	* `value_constraint` - (String) Constraint associated with value, e.g., for string type - regx:[a-z].
	* `value_constraints` - (List) Validation rules for this input value.
	Nested scheme for **value_constraints**:
		* `description` - (String) The value to display if the inptu value does not match the specified constraint.
		* `type` - (String) Type of constraint.
		* `value` - (String) Contstraint value. For type regex, this is a regular expression in Javascript notation.

* `created` - (String) The date and time this version was created.

* `crn` - (String) Version's CRN.

* `deprecate_pending` - (List) Deprecation information for an Offering.
Nested scheme for **deprecate_pending**:
	* `deprecate_date` - (String) Date of deprecation.
	* `deprecate_state` - (String) Deprecation state.
	* `description` - (String)

* `deprecated` - (Boolean) read only field, indicating if this version is deprecated.

* `entitlement` - (List) Entitlement license info.
Nested scheme for **entitlement**:
	* `image_repo_name` - (String) Image repository name.
	* `part_numbers` - (List) list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.
	* `product_id` - (String) Product ID.
	* `provider_id` - (String) Provider ID.
	* `provider_name` - (String) Provider name.

* `flavor` - (List) Version Flavor Information.  Only supported for Product kind Solution.
Nested scheme for **flavor**:
	* `index` - (Integer) Order that this flavor should appear when listed for a single version.
	* `label` - (String) Label for this flavor.
	* `label_i18n` - (Map) A map of translated strings, by language code.
	* `name` - (String) Programmatic name for this flavor.

* `iam_permissions` - (List) List of IAM permissions that are required to consume this version.
Nested scheme for **iam_permissions**:
	* `resources` - (List) Resources for this permission.
	Nested scheme for **resources**:
		* `description` - (String) Resource description.
		* `name` - (String) Resource name.
		* `role_crns` - (List) Role CRNs for this permission.
	* `role_crns` - (List) Role CRNs for this permission.
	* `service_name` - (String) Service name.

* `image_manifest_url` - (String) If set, denotes a url to a YAML file with list of container images used by this version.

* `image_pull_key_name` - (String) ID of the image pull key to use from Offering.ImagePullKeys.

* `install` - (List) Script information.
Nested scheme for **install**:
	* `delete_script` - (String) Optional script that if run will remove the installed version.
	* `instructions` - (String) Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
	* `instructions_i18n` - (Map) A map of translated strings, by language code.
	* `scope` - (String) Optional value indicating if this script is scoped to a namespace or the entire cluster.
	* `script` - (String) Optional script that needs to be run post any pre-condition script.
	* `script_permission` - (String) Optional iam permissions that are required on the target cluster to run this script.

* `is_consumable` - (Boolean) Is the version able to be shared.

* `kind_id` - (String) Kind ID.

* `licenses` - (List) List of licenses the product was built with.
Nested scheme for **licenses**:
	* `description` - (String) License description.
	* `id` - (String) License ID.
	* `name` - (String) license name.
	* `type` - (String) type of license e.g., Apache xxx.
	* `url` - (String) URL for the license text.

* `long_description` - (String) Long description for version.

* `long_description_i18n` - (Map) A map of translated strings, by language code.

* `metadata` - (List) Generic data to be included with content being onboarded. Required for virtual server image for VPC.
Nested scheme for **metadata**:
	* `source_url` - (String) Source URL for the version.
    * `working_directory` - (String) Working directory of source files.
	* `example_name` - (String) Name of example directory that contains source files in existing examples directory.
	* `start_deploy_time` - (String) The time validation starts.
	* `end_deploy_time` - (String) The time validation ends.
	* `est_deploy_time` - (String) The estimated time validation takes.
	* `usage` - (String) Usage text for the version.
	* `usage_template` - (String) Usage text for the version.
	Nested scheme for **modules**:
		* `name` - (String) Name of the module.
		* `source` - (String) Source of the module.
		Nested scheme for **offering_reference**:
			* `name` (String) Name of the offering.
			* `id` (String) ID of the offering.
			* `kind` (String) Kind of the offering.
			* `version` (String) Version of the offering.
			* `flavor` (String) Flavor of the offering.
			* `flavors` (List) Flavors of the offering.
			* `catalog_id` (String) Catalog ID of the offering.
			* `metadata` (String) Metadata of the offering.
	* `terraform_version` - (String) Version's terraform version.
	* `validated_terraform_version` - (String) Version's validated terraform version.
	* `version_name` - (String) Name of the version.
	* `vsi_vpc` - (List) VSI information for this version.
  	Nested scheme for **vsi_vpc**:
    	* `file` - (List) Details for the stored image file. Required for virtual server image for VPC.
    	Nested scheme for **file**:
    		* `size` - (Integer) Size of the stored image file rounded up to the next gigabyte. Required for virtual server image for VPC.
    	* `images` - (List) Image operating system. Required for virtual server image for VPC.
    	Nested scheme for **images**:
    		* `id` - (String) Programmatic ID of virtual server image. Required for virtual server image for VPC.
    		* `name` - (String) Programmatic name of virtual server image. Required for virtual server image for VPC.
    		* `region` - (String) Region the virtual server image is available in. Required for virtual server image for VPC.
    	* `minimum_provisioned_size` - (Integer) Minimum size (in gigabytes) of a volume onto which this image may be provisioned. Required for virtual server image for VPC.
    	* `operating_system` - (List) Operating system included in this image. Required for virtual server image for VPC.
    	Nested scheme for **operating_system**:
    		* `architecture` - (String) Operating system architecture. Required for virtual server image for VPC.
    		* `dedicated_host_only` - (Boolean) Images with this operating system can only be used on dedicated hosts or dedicated host groups. Required for virtual server image for VPC.
    		* `display_name` - (String) Unique, display-friendly name for the operating system. Required for virtual server image for VPC.
    		* `family` - (String) Software family for this operating system. Required for virtual server image for VPC.
    		* `href` - (String) URL for this operating system. Required for virtual server image for VPC.
    		* `name` - (String) Globally unique name for this operating system Required for virtual server image for VPC.
    		* `vendor` - (String) Vendor of the operating system. Required for virtual server image for VPC.
    		* `version` - (String) Major release version of this operating system. Required for virtual server image for VPC.

* `offering_id` - (String) Offering ID.

* `outputs` - (List) List of output values for this version.
Nested scheme for **outputs**:
	* `description` - (String) Output description.
	* `key` - (String) Output key.

* `package_version` - (String) Version of the package used to create this version.

* `pre_install` - (List) Optional pre-install instructions.
Nested scheme for **pre_install**:
	* `delete_script` - (String) Optional script that if run will remove the installed version.
	* `instructions` - (String) Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
	* `instructions_i18n` - (Map) A map of translated strings, by language code.
	* `scope` - (String) Optional value indicating if this script is scoped to a namespace or the entire cluster.
	* `script` - (String) Optional script that needs to be run post any pre-condition script.
	* `script_permission` - (String) Optional iam permissions that are required on the target cluster to run this script.

* `repo_url` - (String) Content's repo URL.

* `required_resources` - (List) Resource requirments for installation.
Nested scheme for **required_resources**:
	* `type` - (String) Type of requirement.
	  * Constraints: Allowable values are: `mem`, `disk`, `cores`, `targetVersion`, `nodes`.
	* `value` - (Map) mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.

* `rev` - (String) Cloudant revision.

* `sha` - (String) hash of the content.

* `single_instance` - (Boolean) Denotes if single instance can be deployed to a given cluster.

* `solution_info` - (List) Version Solution Information.  Only supported for Product kind Solution.
Nested scheme for **solution_info**:
	* `architecture_diagrams` - (List) Architecture diagrams for this solution.
	Nested scheme for **architecture_diagrams**:
		* `description` - (String) Description of this diagram.
		* `description_i18n` - (Map) A map of translated strings, by language code.
		* `diagram` - (List) Offering Media information.
		Nested scheme for **diagram**:
			* `api_url` - (String) CM API specific URL of the specified media item.
			* `caption` - (String) Caption for this media item.
			* `caption_i18n` - (Map) A map of translated strings, by language code.
			* `thumbnail_url` - (String) Thumbnail URL for this media item.
			* `type` - (String) Type of this media item.
			* `url` - (String) URL of the specified media item.
			* `url_proxy` - (List) Offering URL proxy information.
			Nested scheme for **url_proxy**:
				* `sha` - (String) SHA256 fingerprint of image.
				* `url` - (String) URL of the specified media item being proxied.
	* `cost_estimate` - (List) Cost estimate definition.
	Nested scheme for **cost_estimate**:
		* `currency` - (String) Cost estimate currency.
		* `diff_total_hourly_cost` - (String) Difference in total hourly cost.
		* `diff_total_monthly_cost` - (String) Difference in total monthly cost.
		* `past_total_hourly_cost` - (String) Past total hourly cost.
		* `past_total_monthly_cost` - (String) Past total monthly cost.
		* `projects` - (List) Cost estimate projects.
		Nested scheme for **projects**:
			* `breakdown` - (List) Cost breakdown definition.
			Nested scheme for **breakdown**:
				* `resources` - (List) Resources.
				Nested scheme for **resources**:
					* `cost_components` - (List) Cost components.
					Nested scheme for **cost_components**:
						* `hourly_cost` - (String) Cost component hourly cost.
						* `hourly_quantity` - (String) Cost component hourly quantity.
						* `monthly_cost` - (String) Cost component monthly cist.
						* `monthly_quantity` - (String) Cost component monthly quantity.
						* `name` - (String) Cost component name.
						* `price` - (String) Cost component price.
						* `unit` - (String) Cost component unit.
					* `hourly_cost` - (String) Hourly cost.
					* `metadata` - (Map) Resource metadata.
					* `monthly_cost` - (String) Monthly cost.
					* `name` - (String) Resource name.
				* `total_hourly_cost` - (String) Total hourly cost.
				* `total_monthly_c_ost` - (String) Total monthly cost.
			* `diff` - (List) Cost breakdown definition.
			Nested scheme for **diff**:
				* `resources` - (List) Resources.
				Nested scheme for **resources**:
					* `cost_components` - (List) Cost components.
					Nested scheme for **cost_components**:
						* `hourly_cost` - (String) Cost component hourly cost.
						* `hourly_quantity` - (String) Cost component hourly quantity.
						* `monthly_cost` - (String) Cost component monthly cist.
						* `monthly_quantity` - (String) Cost component monthly quantity.
						* `name` - (String) Cost component name.
						* `price` - (String) Cost component price.
						* `unit` - (String) Cost component unit.
					* `hourly_cost` - (String) Hourly cost.
					* `metadata` - (Map) Resource metadata.
					* `monthly_cost` - (String) Monthly cost.
					* `name` - (String) Resource name.
				* `total_hourly_cost` - (String) Total hourly cost.
				* `total_monthly_c_ost` - (String) Total monthly cost.
			* `metadata` - (Map) Project metadata.
			* `name` - (String) Project name.
			* `past_breakdown` - (List) Cost breakdown definition.
			Nested scheme for **past_breakdown**:
				* `resources` - (List) Resources.
				Nested scheme for **resources**:
					* `cost_components` - (List) Cost components.
					Nested scheme for **cost_components**:
						* `hourly_cost` - (String) Cost component hourly cost.
						* `hourly_quantity` - (String) Cost component hourly quantity.
						* `monthly_cost` - (String) Cost component monthly cist.
						* `monthly_quantity` - (String) Cost component monthly quantity.
						* `name` - (String) Cost component name.
						* `price` - (String) Cost component price.
						* `unit` - (String) Cost component unit.
					* `hourly_cost` - (String) Hourly cost.
					* `metadata` - (Map) Resource metadata.
					* `monthly_cost` - (String) Monthly cost.
					* `name` - (String) Resource name.
				* `total_hourly_cost` - (String) Total hourly cost.
				* `total_monthly_c_ost` - (String) Total monthly cost.
			* `summary` - (List) Cost summary definition.
			Nested scheme for **summary**:
				* `no_price_resource_counts` - (Map) No price resource counts.
				* `total_detected_resources` - (Integer) Total detected resources.
				* `total_no_price_resources` - (Integer) Total no price resources.
				* `total_supported_resources` - (Integer) Total supported resources.
				* `total_unsupported_resources` - (Integer) Total unsupported resources.
				* `total_usage_based_resources` - (Integer) Total usage based resources.
				* `unsupported_resource_counts` - (Map) Unsupported resource counts.
		* `summary` - (List) Cost summary definition.
		Nested scheme for **summary**:
			* `no_price_resource_counts` - (Map) No price resource counts.
			* `total_detected_resources` - (Integer) Total detected resources.
			* `total_no_price_resources` - (Integer) Total no price resources.
			* `total_supported_resources` - (Integer) Total supported resources.
			* `total_unsupported_resources` - (Integer) Total unsupported resources.
			* `total_usage_based_resources` - (Integer) Total usage based resources.
			* `unsupported_resource_counts` - (Map) Unsupported resource counts.
		* `time_generated` - (String) When this estimate was generated.
		* `total_hourly_cost` - (String) Total hourly cost.
		* `total_monthly_cost` - (String) Total monthly cost.
		* `version` - (String) Cost estimate version.
	* `dependencies` - (List) Dependencies for this solution.
	Nested scheme for **dependencies**:
		* `catalog_id` - (String) Optional - If not specified, assumes the Public Catalog.
		* `flavors` - (List) Optional - List of dependent flavors in the specified range.
		* `id` - (String) Optional - Offering ID - not required if name is set.
		* `name` - (String) Optional - Programmatic Offering name.
		* `version` - (String) Required - Semver value or range.
	* `features` - (List) Features - titles only.
	Nested scheme for **features**:
		* `description` - (String) Feature description.
		* `description_i18n` - (Map) A map of translated strings, by language code.
		* `title` - (String) Heading.
		* `title_i18n` - (Map) A map of translated strings, by language code.

* `source_url` - (String) Content's source URL (e.g git repo).

* `state` - (List) Offering state.
Nested scheme for **state**:
	* `current` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `current_entered` - (String) Date and time of current request.
	* `pending` - (String) one of: new, validated, account-published, ibm-published, public-published.
	* `pending_requested` - (String) Date and time of pending request.
	* `previous` - (String) one of: new, validated, account-published, ibm-published, public-published.

* `tags` - (List) List of tags associated with this catalog.

* `tgz_url` - (String) File used to on-board this version.

* `updated` - (String) The date and time this version was last updated.

* `validation` - (List) Validation response.
Nested scheme for **validation**:
	* `last_operation` - (String) Last operation (e.g. submit_deployment, generate_installer, install_offering.
	* `message` - (String) Any message needing to be conveyed as part of the validation job.
	* `requested` - (String) Date and time of last validation was requested.
	* `state` - (String) Current validation state - <empty>, in_progress, valid, invalid, expired.
	* `target` - (Map) Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.
	* `validated` - (String) Date and time of last successful validation.

* `version` - (String) Version of content type.

* `version_id` - (String) Unique ID.

* `version_locator` - (String) A dotted value of `catalogID`.`versionID`.

* `whitelisted_accounts` - (List) Whitelisted accounts for version.

