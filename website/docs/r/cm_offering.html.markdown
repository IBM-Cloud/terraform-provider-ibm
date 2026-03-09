---
layout: "ibm"
page_title: "IBM : ibm_cm_offering"
description: |-
  Manages ibm_cm_offering.
subcategory: "Catalog Management"
---

# ibm_cm_offering

Provides a resource for ibm_cm_offering. This allows ibm_cm_offering to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_offering" "cm_offering" {
  catalog_id = ibm_cm_catalog.cm_catalog.id
  label = "offering label"
  name = "offering name"
  offering_icon_url = "icon_url"
  tags = [ "tag1", "tag2" ]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `catalog_id` - (Required, Forces new resource, String) Catalog identifier.
* `offering_id` - (Optional, Forces new resource, String) Offering identifier, provide to import an existing offering.
* `hidden` - (Optional, Boolean) Determine if this offering should be displayed in the Consumption UI.
* `label` - (Optional, String) Display Name in the requested language.
* `long_description` - (Optional, String) Long description in the requested language.
* `name` - (Optional, String) The programmatic name of this offering.
* `offering_icon_url` - (Optional, String) URL for an icon associated with this offering.
* `offering_docs_url` - (Optional, String) URL for additional docs of this offering.
* `product_kind` - (Optional, String) The kind of the product.  Valid values are "solution" and "software"
* `provider_info` - (Optional, List) Information on the provider of this offering.
Nested scheme for **provider_info**:
	* `id` - (Optional, String) The provider ID.
	* `name` - (Optional, String) The provider name.
* `features` - (Optional, List) List of features associated with this offering.
Nested scheme for **features**:
	* `description` - (Optional, String) Feature description.
	* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `title` - (Optional, String) Heading.
	* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
* `media` - (Optional, List) A list of media items related to this offering.
Nested scheme for **media**:
	* `api_url` - (Optional, String) CM API specific URL of the specified media item.
	* `caption` - (Optional, String) Caption for this media item.
	* `caption_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `thumbnail_url` - (Optional, String) Thumbnail URL for this media item.
	* `type` - (Optional, String) Type of this media item.
	* `url` - (Optional, String) URL of the specified media item.
	* `url_proxy` - (Optional, List) Offering URL proxy information.
	Nested scheme for **url_proxy**:
		* `sha` - (Optional, String) SHA256 fingerprint of image.
		* `url` - (Optional, String) URL of the specified media item being proxied.
* `short_description` - (Optional, String) Short description in the requested language.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `keywords` - (Optional, List) List of keywords associated with an offering, typically used to search for it.
* `deprecate` - (Optional, Boolean) Specify if this offering should be deprecated.
* `share_with_access_list` - (Optional, List) List of account, enterprise, or enterprise group IDs.  Enterprise IDs should be prefixed with `-ent-` and enterpries group IDs should be prefixed with `-entgrp-`.
* `share_with_all` - (Optional, Boolean) Denotes public availability of an Offering - if share_enabled is true.
* `share_with_ibm` - (Optional, Boolean) Denotes IBM employee availability of an Offering - if share_enabled is true.
* `share_enabled` - (Optional, Boolean) Denotes sharing including access list availability of an Offering is enabled.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the ibm_cm_offering.
* `badges` - (List) A list of badges for this offering.
Nested scheme for **badges**:
	* `authority` - (String) Authority for the current badge.
	* `constraints` - (List) An optional set of constraints indicating which versions in an Offering have this particular badge.
	Nested scheme for **constraints**:
		* `rule` - (Map) Rule for the current constraint.
		* `type` - (String) Type of the current constraint.
	* `description` - (String) Description of the current badge.
	* `icon` - (String) Icon for the current badge.
	* `id` - (String) ID of the current badge.
	* `label` - (String) Display name for the current badge.
	* `learn_more_links` - (List) Learn more links for a badge.
	Nested scheme for **learn_more_links**:
		* `first_party` - (String) First party link.
		* `third_party` - (String) Third party link.
	* `tag` - (String) Tag for the current badge.

* `catalog_id` - (String) The id of the catalog containing this offering.

* `catalog_name` - (String) The name of the catalog.

* `created` - (String) The date and time this catalog was created.

* `crn` - (String) The crn for this specific offering.

* `deprecate_pending` - (List) Deprecation information for an Offering.
Nested scheme for **deprecate_pending**:
	* `deprecate_date` - (String) Date of deprecation.
	* `deprecate_state` - (String) Deprecation state.
	* `description` - (String)

* `disclaimer` - (String) A disclaimer for this offering.

* `features` - (List) list of features associated with this offering.
Nested scheme for **features**:
	* `description` - (String) Feature description.
	* `description_i18n` - (Map) A map of translated strings, by language code.
	* `title` - (String) Heading.
	* `title_i18n` - (Map) A map of translated strings, by language code.

* `hidden` - (Boolean) Determine if this offering should be displayed in the Consumption UI.

* `ibm_publish_approved` - (Deprecated, Boolean) Indicates if this offering has been approved for use by all IBMers.

* `image_pull_keys` - (List) Image pull keys for this offering.
Nested scheme for **image_pull_keys**:
	* `description` - (String) Key description.
	* `name` - (String) Key name.
	* `value` - (String) Key value.

* `keywords` - (List) List of keywords associated with offering, typically used to search for it.

* `kinds` - (List) Array of kind.
Nested scheme for **kinds**:
	* `additional_features` - (List) List of features associated with this offering.
	Nested scheme for **additional_features**:
		* `description` - (String) Feature description.
		* `description_i18n` - (Map) A map of translated strings, by language code.
		* `title` - (String) Heading.
		* `title_i18n` - (Map) A map of translated strings, by language code.
	* `created` - (String) The date and time this catalog was created.
	* `format_kind` - (String) content kind, e.g., helm, vm image.
	* `id` - (String) Unique ID.
	* `install_kind` - (String) install kind, e.g., helm, operator, terraform.
	* `metadata` - (Map) Open ended metadata information.
	* `plans` - (List) list of plans.
	Nested scheme for **plans**:
		* `additional_features` - (List) list of features associated with this offering.
		Nested scheme for **additional_features**:
			* `description` - (String) Feature description.
			* `description_i18n` - (Map) A map of translated strings, by language code.
			* `title` - (String) Heading.
			* `title_i18n` - (Map) A map of translated strings, by language code.
		* `created` - (String) the date'time this catalog was created.
		* `deployments` - (List) list of deployments.
		Nested scheme for **deployments**:
			* `created` - (String) the date'time this catalog was created.
			* `id` - (String) unique id.
			* `label` - (String) Display Name in the requested language.
			* `long_description` - (String) Long description in the requested language.
			* `metadata` - (Map) open ended metadata information.
			* `name` - (String) The programmatic name of this offering.
			* `short_description` - (String) Short description in the requested language.
			* `tags` - (List) list of tags associated with this catalog.
			* `updated` - (String) the date'time this catalog was last updated.
		* `id` - (String) unique id.
		* `label` - (String) Display Name in the requested language.
		* `long_description` - (String) Long description in the requested language.
		* `metadata` - (Map) open ended metadata information.
		* `name` - (String) The programmatic name of this offering.
		* `short_description` - (String) Short description in the requested language.
		* `tags` - (List) list of tags associated with this catalog.
		* `updated` - (String) the date'time this catalog was last updated.
	* `tags` - (List) List of tags associated with this catalog.
	* `target_kind` - (String) target cloud to install, e.g., iks, open_shift_iks.
	* `updated` - (String) The date and time this catalog was last updated.
	* `versions` - (List) list of versions.
	Nested scheme for **versions**:
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
		* `id` - (String) Unique ID.
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
		* `metadata` - (Map) Open ended metadata information.
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
		* `version_locator` - (String) A dotted value of `catalogID`.`versionID`.
		* `whitelisted_accounts` - (List) Whitelisted accounts for version.

* `label` - (String) Display Name in the requested language.

* `label_i18n` - (Map) A map of translated strings, by language code.

* `long_description` - (String) Long description in the requested language.

* `long_description_i18n` - (Map) A map of translated strings, by language code.

* `media` - (List) A list of media items related to this offering.
Nested scheme for **media**:
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

* `metadata` - (Map) Map of metadata values for this offering.

* `name` - (String) The programmatic name of this offering.

* `offering_docs_url` - (String) URL for an additional docs with this offering.

* `offering_icon_url` - (String) URL for an icon associated with this offering.

* `offering_support_url` - (String) [deprecated] - Use offering.support instead.  URL to be displayed in the Consumption UI for getting support on this offering.

* `pc_managed` - (Boolean) Offering is managed by Partner Center.

* `permit_request_ibm_public_publish` - (Deprecated, Boolean) Is it permitted to request publishing to IBM or Public.

* `portal_approval_record` - (String) The portal's approval record ID.

* `portal_ui_url` - (String) The portal UI URL.

* `product_kind` - (String) The product kind.  Valid values are module, solution, or empty string.

* `provider` - (Deprecated, String) Deprecated - Provider of this offering.

* `provider_info` - (List) Information on the provider for this offering, or omitted if no provider information is given.
Nested scheme for **provider_info**:
	* `id` - (String) The id of this provider.
	* `name` - (String) The name of this provider.

* `public_original_crn` - (String) The original offering CRN that this publish entry came from.

* `public_publish_approved` - (Deprecated, Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.

* `publish_approved` - (Boolean) Offering has been approved to publish to permitted to IBM or Public Catalog.

* `publish_public_crn` - (String) The crn of the public catalog entry of this offering.

* `rating` - (List) Repository info for offerings.
Nested scheme for **rating**:
	* `four_star_count` - (Integer) Four start rating.
	* `one_star_count` - (Integer) One start rating.
	* `three_star_count` - (Integer) Three start rating.
	* `two_star_count` - (Integer) Two start rating.

* `repo_info` - (List) Repository info for offerings.
Nested scheme for **repo_info**:
	* `token` - (String) Token for private repos.
	* `type` - (String) Public or enterprise GitHub.

* `rev` - (String) Cloudant revision.

* `share_enabled` - (Boolean) Denotes sharing including access list availability of an Offering is enabled.

* `share_with_all` - (Boolean) Denotes public availability of an Offering - if share_enabled is true.

* `share_with_ibm` - (Boolean) Denotes IBM employee availability of an Offering - if share_enabled is true.

* `short_description` - (String) Short description in the requested language.

* `short_description_i18n` - (Map) A map of translated strings, by language code.

* `support` - (List) Offering Support information.
Nested scheme for **support**:
	* `locations` - (List) A list of country codes indicating where support is provided.
	* `process` - (String) Support process as provided by an ISV.
	* `process_i18n` - (Map) A map of translated strings, by language code.
	* `support_details` - (List) A list of support options (e.g. email, phone, slack, other).
	Nested scheme for **support_details**:
		* `availability` - (List) Times when support is available.
		Nested scheme for **availability**:
			* `always_available` - (Boolean) Is this support always available.
			* `times` - (List) A list of support times.
			Nested scheme for **times**:
				* `day` - (Integer) The day of the week, represented as an integer.
				* `end_time` - (String) HOURS:MINUTES:SECONDS using 24 hour time (e.g. 8:15:00).
				* `start_time` - (String) HOURS:MINUTES:SECONDS using 24 hour time (e.g. 8:15:00).
			* `timezone` - (String) Timezone (e.g. America/New_York).
		* `contact` - (String) Contact for the current support detail.
		* `response_wait_time` - (List) Time descriptor.
		Nested scheme for **response_wait_time**:
			* `type` - (String) Valid values are hour or day.
			* `value` - (Integer) Amount of time to wait in unit 'type'.
		* `type` - (String) Type of the current support detail.
	* `support_escalation` - (List) Support escalation policy.
	Nested scheme for **support_escalation**:
		* `contact` - (String) Escalation contact.
		* `escalation_wait_time` - (List) Time descriptor.
		Nested scheme for **escalation_wait_time**:
			* `type` - (String) Valid values are hour or day.
			* `value` - (Integer) Amount of time to wait in unit 'type'.
		* `response_wait_time` - (List) Time descriptor.
		Nested scheme for **response_wait_time**:
			* `type` - (String) Valid values are hour or day.
			* `value` - (Integer) Amount of time to wait in unit 'type'.
	* `support_type` - (String) Support type for this product.
	* `url` - (String) URL to be displayed in the Consumption UI for getting support on this offering.

* `tags` - (List) List of tags associated with this catalog.

* `updated` - (String) The date and time this catalog was last updated.

* `url` - (String) The url for this specific offering.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cm_offering` resource by using `id`.
The `id` value specified should be a combination of the `catalog_id` and the `offering_id` separated by a `:`.  For example 
`00000000-0000-0000-0000-000000000000:11111111-1111-1111-1111-111111111111` .

* `offering_id`: A string. Offering identification.

# Syntax
```
$ terraform import ibm_cm_offering.cm_offering <offering_id>
```
