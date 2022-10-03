---
layout: "ibm"
page_title: "IBM : ibm_cm_offering"
description: |-
  Manages cm_offering.
subcategory: "Catalog Management API"
---

# ibm_cm_offering

Provides a resource for cm_offering. This allows cm_offering to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_offering" "cm_offering" {
  badges {
		id = "id"
		label = "label"
		label_i18n = { "key": "inner" }
		description = "description"
		description_i18n = { "key": "inner" }
		icon = "icon"
		authority = "authority"
		tag = "tag"
		learn_more_links {
			first_party = "first_party"
			third_party = "third_party"
		}
		constraints {
			type = "type"
		}
  }
  catalog_identifier = ibm_cm_catalog.cm_catalog.id
  deprecate_pending {
		deprecate_date = "2021-01-31T09:44:12Z"
		deprecate_state = "deprecate_state"
		description = "description"
  }
  features {
		title = "title"
		title_i18n = { "key": "inner" }
		description = "description"
		description_i18n = { "key": "inner" }
  }
  image_pull_keys {
		name = "name"
		value = "value"
		description = "description"
  }
  kinds {
		id = "id"
		format_kind = "format_kind"
		install_kind = "install_kind"
		target_kind = "target_kind"
		metadata = { "key": null }
		tags = [ "tags" ]
		additional_features {
			title = "title"
			title_i18n = { "key": "inner" }
			description = "description"
			description_i18n = { "key": "inner" }
		}
		created = "2021-01-31T09:44:12Z"
		updated = "2021-01-31T09:44:12Z"
		versions {
			id = "id"
			rev = "rev"
			crn = "crn"
			version = "version"
			flavor {
				name = "name"
				label = "label"
				label_i18n = { "key": "inner" }
				index = 1
			}
			sha = "sha"
			created = "2021-01-31T09:44:12Z"
			updated = "2021-01-31T09:44:12Z"
			offering_id = "offering_id"
			catalog_id = "catalog_id"
			kind_id = "kind_id"
			tags = [ "tags" ]
			repo_url = "repo_url"
			source_url = "source_url"
			tgz_url = "tgz_url"
			configuration {
				key = "key"
				type = "type"
				display_name = "display_name"
				value_constraint = "value_constraint"
				description = "description"
				required = true
				options = [ null ]
				hidden = true
				custom_config {
					type = "type"
					grouping = "grouping"
					original_grouping = "original_grouping"
					grouping_index = 1
					config_constraints = { "key": null }
					associations {
						parameters {
							name = "name"
							options_refresh = true
						}
					}
				}
				type_metadata = "type_metadata"
			}
			outputs {
				key = "key"
				description = "description"
			}
			iam_permissions {
				service_name = "service_name"
				role_crns = [ "role_crns" ]
				resources {
					name = "name"
					description = "description"
					role_crns = [ "role_crns" ]
				}
			}
			metadata = { "key": null }
			validation {
				validated = "2021-01-31T09:44:12Z"
				requested = "2021-01-31T09:44:12Z"
				state = "state"
				last_operation = "last_operation"
				target = { "key": null }
				message = "message"
			}
			required_resources {
				type = "mem"
			}
			single_instance = true
			install {
				instructions = "instructions"
				instructions_i18n = { "key": "inner" }
				script = "script"
				script_permission = "script_permission"
				delete_script = "delete_script"
				scope = "scope"
			}
			pre_install {
				instructions = "instructions"
				instructions_i18n = { "key": "inner" }
				script = "script"
				script_permission = "script_permission"
				delete_script = "delete_script"
				scope = "scope"
			}
			entitlement {
				provider_name = "provider_name"
				provider_id = "provider_id"
				product_id = "product_id"
				part_numbers = [ "part_numbers" ]
				image_repo_name = "image_repo_name"
			}
			licenses {
				id = "id"
				name = "name"
				type = "type"
				url = "url"
				description = "description"
			}
			image_manifest_url = "image_manifest_url"
			deprecated = true
			package_version = "package_version"
			state {
				current = "current"
				current_entered = "2021-01-31T09:44:12Z"
				pending = "pending"
				pending_requested = "2021-01-31T09:44:12Z"
				previous = "previous"
			}
			version_locator = "version_locator"
			long_description = "long_description"
			long_description_i18n = { "key": "inner" }
			whitelisted_accounts = [ "whitelisted_accounts" ]
			image_pull_key_name = "image_pull_key_name"
			deprecate_pending {
				deprecate_date = "2021-01-31T09:44:12Z"
				deprecate_state = "deprecate_state"
				description = "description"
			}
			solution_info {
				architecture_diagrams {
					diagram {
						url = "url"
						api_url = "api_url"
						url_proxy {
							url = "url"
							sha = "sha"
						}
						caption = "caption"
						caption_i18n = { "key": "inner" }
						type = "type"
						thumbnail_url = "thumbnail_url"
					}
					description = "description"
					description_i18n = { "key": "inner" }
				}
				features {
					title = "title"
					title_i18n = { "key": "inner" }
					description = "description"
					description_i18n = { "key": "inner" }
				}
				cost_estimate {
					version = "version"
					currency = "currency"
					projects {
						name = "name"
						metadata = { "key": null }
						past_breakdown {
							total_hourly_cost = "total_hourly_cost"
							total_monthly_c_ost = "total_monthly_c_ost"
							resources {
								name = "name"
								metadata = { "key": null }
								hourly_cost = "hourly_cost"
								monthly_cost = "monthly_cost"
								cost_components {
									name = "name"
									unit = "unit"
									hourly_quantity = "hourly_quantity"
									monthly_quantity = "monthly_quantity"
									price = "price"
									hourly_cost = "hourly_cost"
									monthly_cost = "monthly_cost"
								}
							}
						}
						breakdown {
							total_hourly_cost = "total_hourly_cost"
							total_monthly_c_ost = "total_monthly_c_ost"
							resources {
								name = "name"
								metadata = { "key": null }
								hourly_cost = "hourly_cost"
								monthly_cost = "monthly_cost"
								cost_components {
									name = "name"
									unit = "unit"
									hourly_quantity = "hourly_quantity"
									monthly_quantity = "monthly_quantity"
									price = "price"
									hourly_cost = "hourly_cost"
									monthly_cost = "monthly_cost"
								}
							}
						}
						diff {
							total_hourly_cost = "total_hourly_cost"
							total_monthly_c_ost = "total_monthly_c_ost"
							resources {
								name = "name"
								metadata = { "key": null }
								hourly_cost = "hourly_cost"
								monthly_cost = "monthly_cost"
								cost_components {
									name = "name"
									unit = "unit"
									hourly_quantity = "hourly_quantity"
									monthly_quantity = "monthly_quantity"
									price = "price"
									hourly_cost = "hourly_cost"
									monthly_cost = "monthly_cost"
								}
							}
						}
						summary {
							total_detected_resources = 1
							total_supported_resources = 1
							total_unsupported_resources = 1
							total_usage_based_resources = 1
							total_no_price_resources = 1
							unsupported_resource_counts = { "key": 1 }
							no_price_resource_counts = { "key": 1 }
						}
					}
					summary {
						total_detected_resources = 1
						total_supported_resources = 1
						total_unsupported_resources = 1
						total_usage_based_resources = 1
						total_no_price_resources = 1
						unsupported_resource_counts = { "key": 1 }
						no_price_resource_counts = { "key": 1 }
					}
					total_hourly_cost = "total_hourly_cost"
					total_monthly_cost = "total_monthly_cost"
					past_total_hourly_cost = "past_total_hourly_cost"
					past_total_monthly_cost = "past_total_monthly_cost"
					diff_total_hourly_cost = "diff_total_hourly_cost"
					diff_total_monthly_cost = "diff_total_monthly_cost"
					time_generated = "2021-01-31T09:44:12Z"
				}
				dependencies {
					catalog_id = "catalog_id"
					id = "id"
					name = "name"
					version = "version"
					flavors = [ "flavors" ]
				}
			}
			is_consumable = true
		}
		plans {
			id = "id"
			label = "label"
			name = "name"
			short_description = "short_description"
			long_description = "long_description"
			metadata = { "key": null }
			tags = [ "tags" ]
			additional_features {
				title = "title"
				title_i18n = { "key": "inner" }
				description = "description"
				description_i18n = { "key": "inner" }
			}
			created = "2021-01-31T09:44:12Z"
			updated = "2021-01-31T09:44:12Z"
			deployments {
				id = "id"
				label = "label"
				name = "name"
				short_description = "short_description"
				long_description = "long_description"
				metadata = { "key": null }
				tags = [ "tags" ]
				created = "2021-01-31T09:44:12Z"
				updated = "2021-01-31T09:44:12Z"
			}
		}
  }
  media {
		url = "url"
		api_url = "api_url"
		url_proxy {
			url = "url"
			sha = "sha"
		}
		caption = "caption"
		caption_i18n = { "key": "inner" }
		type = "type"
		thumbnail_url = "thumbnail_url"
  }
  provider_info {
		id = "id"
		name = "name"
  }
  rating {
		one_star_count = 1
		two_star_count = 1
		three_star_count = 1
		four_star_count = 1
  }
  repo_info {
		token = "token"
		type = "type"
  }
  support {
		url = "url"
		process = "process"
		process_i18n = { "key": "inner" }
		locations = [ "locations" ]
		support_details {
			type = "type"
			contact = "contact"
			response_wait_time {
				value = 1
				type = "type"
			}
			availability {
				times {
					day = 1
					start_time = "start_time"
					end_time = "end_time"
				}
				timezone = "timezone"
				always_available = true
			}
		}
		support_escalation {
			escalation_wait_time {
				value = 1
				type = "type"
			}
			response_wait_time {
				value = 1
				type = "type"
			}
			contact = "contact"
		}
		support_type = "support_type"
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `badges` - (Optional, List) A list of badges for this offering.
Nested scheme for **badges**:
	* `authority` - (Optional, String) Authority for the current badge.
	* `constraints` - (Optional, List) An optional set of constraints indicating which versions in an Offering have this particular badge.
	Nested scheme for **constraints**:
		* `rule` - (Optional, Map) Rule for the current constraint.
		* `type` - (Optional, String) Type of the current constraint.
	* `description` - (Optional, String) Description of the current badge.
	* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `icon` - (Optional, String) Icon for the current badge.
	* `id` - (Optional, String) ID of the current badge.
	* `label` - (Optional, String) Display name for the current badge.
	* `label_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `learn_more_links` - (Optional, List) Learn more links for a badge.
	Nested scheme for **learn_more_links**:
		* `first_party` - (Optional, String) First party link.
		* `third_party` - (Optional, String) Third party link.
	* `tag` - (Optional, String) Tag for the current badge.
* `catalog_id` - (Optional, String) The id of the catalog containing this offering.
* `catalog_identifier` - (Required, Forces new resource, String) Catalog identifier.
* `catalog_name` - (Optional, String) The name of the catalog.
* `created` - (Optional, String) The date and time this catalog was created.
* `crn` - (Optional, String) The crn for this specific offering.
* `deprecate_pending` - (Optional, List) Deprecation information for an Offering.
Nested scheme for **deprecate_pending**:
	* `deprecate_date` - (Optional, String) Date of deprecation.
	* `deprecate_state` - (Optional, String) Deprecation state.
	* `description` - (Optional, String)
* `disclaimer` - (Optional, String) A disclaimer for this offering.
* `features` - (Optional, List) list of features associated with this offering.
Nested scheme for **features**:
	* `description` - (Optional, String) Feature description.
	* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `title` - (Optional, String) Heading.
	* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
* `hidden` - (Optional, Boolean) Determine if this offering should be displayed in the Consumption UI.
* `ibm_publish_approved` - (Deprecated, Optional, Boolean) Indicates if this offering has been approved for use by all IBMers.
* `image_pull_keys` - (Optional, List) Image pull keys for this offering.
Nested scheme for **image_pull_keys**:
	* `description` - (Optional, String) Key description.
	* `name` - (Optional, String) Key name.
	* `value` - (Optional, String) Key value.
* `keywords` - (Optional, List) List of keywords associated with offering, typically used to search for it.
* `kinds` - (Optional, List) Array of kind.
Nested scheme for **kinds**:
	* `additional_features` - (Optional, List) List of features associated with this offering.
	Nested scheme for **additional_features**:
		* `description` - (Optional, String) Feature description.
		* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
		* `title` - (Optional, String) Heading.
		* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `created` - (Optional, String) The date and time this catalog was created.
	* `format_kind` - (Optional, String) content kind, e.g., helm, vm image.
	* `id` - (Optional, String) Unique ID.
	* `install_kind` - (Optional, String) install kind, e.g., helm, operator, terraform.
	* `metadata` - (Optional, Map) Open ended metadata information.
	* `plans` - (Optional, List) list of plans.
	Nested scheme for **plans**:
		* `additional_features` - (Optional, List) list of features associated with this offering.
		Nested scheme for **additional_features**:
			* `description` - (Optional, String) Feature description.
			* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
			* `title` - (Optional, String) Heading.
			* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
		* `created` - (Optional, String) the date'time this catalog was created.
		* `deployments` - (Optional, List) list of deployments.
		Nested scheme for **deployments**:
			* `created` - (Optional, String) the date'time this catalog was created.
			* `id` - (Optional, String) unique id.
			* `label` - (Optional, String) Display Name in the requested language.
			* `long_description` - (Optional, String) Long description in the requested language.
			* `metadata` - (Optional, Map) open ended metadata information.
			* `name` - (Optional, String) The programmatic name of this offering.
			* `short_description` - (Optional, String) Short description in the requested language.
			* `tags` - (Optional, List) list of tags associated with this catalog.
			* `updated` - (Optional, String) the date'time this catalog was last updated.
		* `id` - (Optional, String) unique id.
		* `label` - (Optional, String) Display Name in the requested language.
		* `long_description` - (Optional, String) Long description in the requested language.
		* `metadata` - (Optional, Map) open ended metadata information.
		* `name` - (Optional, String) The programmatic name of this offering.
		* `short_description` - (Optional, String) Short description in the requested language.
		* `tags` - (Optional, List) list of tags associated with this catalog.
		* `updated` - (Optional, String) the date'time this catalog was last updated.
	* `tags` - (Optional, List) List of tags associated with this catalog.
	* `target_kind` - (Optional, String) target cloud to install, e.g., iks, open_shift_iks.
	* `updated` - (Optional, String) The date and time this catalog was last updated.
	* `versions` - (Optional, List) list of versions.
	Nested scheme for **versions**:
		* `catalog_id` - (Optional, String) Catalog ID.
		* `configuration` - (Optional, List) List of user solicited overrides.
		Nested scheme for **configuration**:
			* `custom_config` - (Optional, List) Render type.
			Nested scheme for **custom_config**:
				* `associations` - (Optional, List) List of parameters that are associated with this configuration.
				Nested scheme for **associations**:
					* `parameters` - (Optional, List) Parameters for this association.
					Nested scheme for **parameters**:
						* `name` - (Optional, String) Name of this parameter.
						* `options_refresh` - (Optional, Boolean) Refresh options.
				* `config_constraints` - (Optional, Map) Map of constraint parameters that will be passed to the custom widget.
				* `grouping` - (Optional, String) Determines where this configuration type is rendered (3 sections today - Target, Resource, and Deployment).
				* `grouping_index` - (Optional, Integer) Determines the order that this configuration item shows in that particular grouping.
				* `original_grouping` - (Optional, String) Original grouping type for this configuration (3 types - Target, Resource, and Deployment).
				* `type` - (Optional, String) ID of the widget type.
			* `default_value` - (Optional, Map) The default value.  To use a secret when the type is password, specify a JSON encoded value of $ref:#/components/schemas/SecretInstance, prefixed with `cmsm_v1:`.
			* `description` - (Optional, String) Key description.
			* `display_name` - (Optional, String) Display name for configuration type.
			* `hidden` - (Optional, Boolean) Hide values.
			* `key` - (Optional, String) Configuration key.
			* `options` - (Optional, List) List of options of type.
			* `required` - (Optional, Boolean) Is key required to install.
			* `type` - (Optional, String) Value type (string, boolean, int).
			* `type_metadata` - (Optional, String) The original type, as found in the source being onboarded.
			* `value_constraint` - (Optional, String) Constraint associated with value, e.g., for string type - regx:[a-z].
		* `created` - (Optional, String) The date and time this version was created.
		* `crn` - (Optional, String) Version's CRN.
		* `deprecate_pending` - (Optional, List) Deprecation information for an Offering.
		Nested scheme for **deprecate_pending**:
			* `deprecate_date` - (Optional, String) Date of deprecation.
			* `deprecate_state` - (Optional, String) Deprecation state.
			* `description` - (Optional, String)
		* `deprecated` - (Optional, Boolean) read only field, indicating if this version is deprecated.
		* `entitlement` - (Optional, List) Entitlement license info.
		Nested scheme for **entitlement**:
			* `image_repo_name` - (Optional, String) Image repository name.
			* `part_numbers` - (Optional, List) list of license entitlement part numbers, eg. D1YGZLL,D1ZXILL.
			* `product_id` - (Optional, String) Product ID.
			* `provider_id` - (Optional, String) Provider ID.
			* `provider_name` - (Optional, String) Provider name.
		* `flavor` - (Optional, List) Version Flavor Information.  Only supported for Product kind Solution.
		Nested scheme for **flavor**:
			* `index` - (Optional, Integer) Order that this flavor should appear when listed for a single version.
			* `label` - (Optional, String) Label for this flavor.
			* `label_i18n` - (Optional, Map) A map of translated strings, by language code.
			* `name` - (Optional, String) Programmatic name for this flavor.
		* `iam_permissions` - (Optional, List) List of IAM permissions that are required to consume this version.
		Nested scheme for **iam_permissions**:
			* `resources` - (Optional, List) Resources for this permission.
			Nested scheme for **resources**:
				* `description` - (Optional, String) Resource description.
				* `name` - (Optional, String) Resource name.
				* `role_crns` - (Optional, List) Role CRNs for this permission.
			* `role_crns` - (Optional, List) Role CRNs for this permission.
			* `service_name` - (Optional, String) Service name.
		* `id` - (Optional, String) Unique ID.
		* `image_manifest_url` - (Optional, String) If set, denotes a url to a YAML file with list of container images used by this version.
		* `image_pull_key_name` - (Optional, String) ID of the image pull key to use from Offering.ImagePullKeys.
		* `install` - (Optional, List) Script information.
		Nested scheme for **install**:
			* `delete_script` - (Optional, String) Optional script that if run will remove the installed version.
			* `instructions` - (Optional, String) Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
			* `instructions_i18n` - (Optional, Map) A map of translated strings, by language code.
			* `scope` - (Optional, String) Optional value indicating if this script is scoped to a namespace or the entire cluster.
			* `script` - (Optional, String) Optional script that needs to be run post any pre-condition script.
			* `script_permission` - (Optional, String) Optional iam permissions that are required on the target cluster to run this script.
		* `is_consumable` - (Optional, Boolean) Is the version able to be shared.
		* `kind_id` - (Optional, String) Kind ID.
		* `licenses` - (Optional, List) List of licenses the product was built with.
		Nested scheme for **licenses**:
			* `description` - (Optional, String) License description.
			* `id` - (Optional, String) License ID.
			* `name` - (Optional, String) license name.
			* `type` - (Optional, String) type of license e.g., Apache xxx.
			* `url` - (Optional, String) URL for the license text.
		* `long_description` - (Optional, String) Long description for version.
		* `long_description_i18n` - (Optional, Map) A map of translated strings, by language code.
		* `metadata` - (Optional, Map) Open ended metadata information.
		* `offering_id` - (Optional, String) Offering ID.
		* `outputs` - (Optional, List) List of output values for this version.
		Nested scheme for **outputs**:
			* `description` - (Optional, String) Output description.
			* `key` - (Optional, String) Output key.
		* `package_version` - (Optional, String) Version of the package used to create this version.
		* `pre_install` - (Optional, List) Optional pre-install instructions.
		Nested scheme for **pre_install**:
			* `delete_script` - (Optional, String) Optional script that if run will remove the installed version.
			* `instructions` - (Optional, String) Instruction on step and by whom (role) that are needed to take place to prepare the target for installing this version.
			* `instructions_i18n` - (Optional, Map) A map of translated strings, by language code.
			* `scope` - (Optional, String) Optional value indicating if this script is scoped to a namespace or the entire cluster.
			* `script` - (Optional, String) Optional script that needs to be run post any pre-condition script.
			* `script_permission` - (Optional, String) Optional iam permissions that are required on the target cluster to run this script.
		* `repo_url` - (Optional, String) Content's repo URL.
		* `required_resources` - (Optional, List) Resource requirments for installation.
		Nested scheme for **required_resources**:
			* `type` - (Optional, String) Type of requirement.
			  * Constraints: Allowable values are: `mem`, `disk`, `cores`, `targetVersion`, `nodes`.
			* `value` - (Optional, Map) mem, disk, cores, and nodes can be parsed as an int.  targetVersion will be a semver range value.
		* `rev` - (Optional, String) Cloudant revision.
		* `sha` - (Optional, String) hash of the content.
		* `single_instance` - (Optional, Boolean) Denotes if single instance can be deployed to a given cluster.
		* `solution_info` - (Optional, List) Version Solution Information.  Only supported for Product kind Solution.
		Nested scheme for **solution_info**:
			* `architecture_diagrams` - (Optional, List) Architecture diagrams for this solution.
			Nested scheme for **architecture_diagrams**:
				* `description` - (Optional, String) Description of this diagram.
				* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
				* `diagram` - (Optional, List) Offering Media information.
				Nested scheme for **diagram**:
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
			* `cost_estimate` - (Optional, List) Cost estimate definition.
			Nested scheme for **cost_estimate**:
				* `currency` - (Optional, String) Cost estimate currency.
				* `diff_total_hourly_cost` - (Optional, String) Difference in total hourly cost.
				* `diff_total_monthly_cost` - (Optional, String) Difference in total monthly cost.
				* `past_total_hourly_cost` - (Optional, String) Past total hourly cost.
				* `past_total_monthly_cost` - (Optional, String) Past total monthly cost.
				* `projects` - (Optional, List) Cost estimate projects.
				Nested scheme for **projects**:
					* `breakdown` - (Optional, List) Cost breakdown definition.
					Nested scheme for **breakdown**:
						* `resources` - (Optional, List) Resources.
						Nested scheme for **resources**:
							* `cost_components` - (Optional, List) Cost components.
							Nested scheme for **cost_components**:
								* `hourly_cost` - (Optional, String) Cost component hourly cost.
								* `hourly_quantity` - (Optional, String) Cost component hourly quantity.
								* `monthly_cost` - (Optional, String) Cost component monthly cist.
								* `monthly_quantity` - (Optional, String) Cost component monthly quantity.
								* `name` - (Optional, String) Cost component name.
								* `price` - (Optional, String) Cost component price.
								* `unit` - (Optional, String) Cost component unit.
							* `hourly_cost` - (Optional, String) Hourly cost.
							* `metadata` - (Optional, Map) Resource metadata.
							* `monthly_cost` - (Optional, String) Monthly cost.
							* `name` - (Optional, String) Resource name.
						* `total_hourly_cost` - (Optional, String) Total hourly cost.
						* `total_monthly_c_ost` - (Optional, String) Total monthly cost.
					* `diff` - (Optional, List) Cost breakdown definition.
					Nested scheme for **diff**:
						* `resources` - (Optional, List) Resources.
						Nested scheme for **resources**:
							* `cost_components` - (Optional, List) Cost components.
							Nested scheme for **cost_components**:
								* `hourly_cost` - (Optional, String) Cost component hourly cost.
								* `hourly_quantity` - (Optional, String) Cost component hourly quantity.
								* `monthly_cost` - (Optional, String) Cost component monthly cist.
								* `monthly_quantity` - (Optional, String) Cost component monthly quantity.
								* `name` - (Optional, String) Cost component name.
								* `price` - (Optional, String) Cost component price.
								* `unit` - (Optional, String) Cost component unit.
							* `hourly_cost` - (Optional, String) Hourly cost.
							* `metadata` - (Optional, Map) Resource metadata.
							* `monthly_cost` - (Optional, String) Monthly cost.
							* `name` - (Optional, String) Resource name.
						* `total_hourly_cost` - (Optional, String) Total hourly cost.
						* `total_monthly_c_ost` - (Optional, String) Total monthly cost.
					* `metadata` - (Optional, Map) Project metadata.
					* `name` - (Optional, String) Project name.
					* `past_breakdown` - (Optional, List) Cost breakdown definition.
					Nested scheme for **past_breakdown**:
						* `resources` - (Optional, List) Resources.
						Nested scheme for **resources**:
							* `cost_components` - (Optional, List) Cost components.
							Nested scheme for **cost_components**:
								* `hourly_cost` - (Optional, String) Cost component hourly cost.
								* `hourly_quantity` - (Optional, String) Cost component hourly quantity.
								* `monthly_cost` - (Optional, String) Cost component monthly cist.
								* `monthly_quantity` - (Optional, String) Cost component monthly quantity.
								* `name` - (Optional, String) Cost component name.
								* `price` - (Optional, String) Cost component price.
								* `unit` - (Optional, String) Cost component unit.
							* `hourly_cost` - (Optional, String) Hourly cost.
							* `metadata` - (Optional, Map) Resource metadata.
							* `monthly_cost` - (Optional, String) Monthly cost.
							* `name` - (Optional, String) Resource name.
						* `total_hourly_cost` - (Optional, String) Total hourly cost.
						* `total_monthly_c_ost` - (Optional, String) Total monthly cost.
					* `summary` - (Optional, List) Cost summary definition.
					Nested scheme for **summary**:
						* `no_price_resource_counts` - (Optional, Map) No price resource counts.
						* `total_detected_resources` - (Optional, Integer) Total detected resources.
						* `total_no_price_resources` - (Optional, Integer) Total no price resources.
						* `total_supported_resources` - (Optional, Integer) Total supported resources.
						* `total_unsupported_resources` - (Optional, Integer) Total unsupported resources.
						* `total_usage_based_resources` - (Optional, Integer) Total usage based resources.
						* `unsupported_resource_counts` - (Optional, Map) Unsupported resource counts.
				* `summary` - (Optional, List) Cost summary definition.
				Nested scheme for **summary**:
					* `no_price_resource_counts` - (Optional, Map) No price resource counts.
					* `total_detected_resources` - (Optional, Integer) Total detected resources.
					* `total_no_price_resources` - (Optional, Integer) Total no price resources.
					* `total_supported_resources` - (Optional, Integer) Total supported resources.
					* `total_unsupported_resources` - (Optional, Integer) Total unsupported resources.
					* `total_usage_based_resources` - (Optional, Integer) Total usage based resources.
					* `unsupported_resource_counts` - (Optional, Map) Unsupported resource counts.
				* `time_generated` - (Optional, String) When this estimate was generated.
				* `total_hourly_cost` - (Optional, String) Total hourly cost.
				* `total_monthly_cost` - (Optional, String) Total monthly cost.
				* `version` - (Optional, String) Cost estimate version.
			* `dependencies` - (Optional, List) Dependencies for this solution.
			Nested scheme for **dependencies**:
				* `catalog_id` - (Optional, String) Optional - If not specified, assumes the Public Catalog.
				* `flavors` - (Optional, List) Optional - List of dependent flavors in the specified range.
				* `id` - (Optional, String) Optional - Offering ID - not required if name is set.
				* `name` - (Optional, String) Optional - Programmatic Offering name.
				* `version` - (Optional, String) Required - Semver value or range.
			* `features` - (Optional, List) Features - titles only.
			Nested scheme for **features**:
				* `description` - (Optional, String) Feature description.
				* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
				* `title` - (Optional, String) Heading.
				* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
		* `source_url` - (Optional, String) Content's source URL (e.g git repo).
		* `state` - (Optional, List) Offering state.
		Nested scheme for **state**:
			* `current` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
			* `current_entered` - (Optional, String) Date and time of current request.
			* `pending` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
			* `pending_requested` - (Optional, String) Date and time of pending request.
			* `previous` - (Optional, String) one of: new, validated, account-published, ibm-published, public-published.
		* `tags` - (Optional, List) List of tags associated with this catalog.
		* `tgz_url` - (Optional, String) File used to on-board this version.
		* `updated` - (Optional, String) The date and time this version was last updated.
		* `validation` - (Optional, List) Validation response.
		Nested scheme for **validation**:
			* `last_operation` - (Optional, String) Last operation (e.g. submit_deployment, generate_installer, install_offering.
			* `message` - (Optional, String) Any message needing to be conveyed as part of the validation job.
			* `requested` - (Optional, String) Date and time of last validation was requested.
			* `state` - (Optional, String) Current validation state - <empty>, in_progress, valid, invalid, expired.
			* `target` - (Optional, Map) Validation target information (e.g. cluster_id, region, namespace, etc).  Values will vary by Content type.
			* `validated` - (Optional, String) Date and time of last successful validation.
		* `version` - (Optional, String) Version of content type.
		* `version_locator` - (Optional, String) A dotted value of `catalogID`.`versionID`.
		* `whitelisted_accounts` - (Optional, List) Whitelisted accounts for version.
* `label` - (Optional, String) Display Name in the requested language.
* `label_i18n` - (Optional, Map) A map of translated strings, by language code.
* `long_description` - (Optional, String) Long description in the requested language.
* `long_description_i18n` - (Optional, Map) A map of translated strings, by language code.
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
* `metadata` - (Optional, Map) Map of metadata values for this offering.
* `name` - (Optional, String) The programmatic name of this offering.
* `offering_docs_url` - (Optional, String) URL for an additional docs with this offering.
* `offering_icon_url` - (Optional, String) URL for an icon associated with this offering.
* `offering_support_url` - (Optional, String) [deprecated] - Use offering.support instead.  URL to be displayed in the Consumption UI for getting support on this offering.
* `pc_managed` - (Optional, Boolean) Offering is managed by Partner Center.
* `permit_request_ibm_public_publish` - (Deprecated, Optional, Boolean) Is it permitted to request publishing to IBM or Public.
* `portal_approval_record` - (Optional, String) The portal's approval record ID.
* `portal_ui_url` - (Optional, String) The portal UI URL.
* `product_kind` - (Optional, String) The product kind.  Valid values are module, solution, or empty string.
* `provider` - (Deprecated, Optional, String) Deprecated - Provider of this offering.
* `provider_info` - (Optional, List) Information on the provider for this offering, or omitted if no provider information is given.
Nested scheme for **provider_info**:
	* `id` - (Optional, String) The id of this provider.
	* `name` - (Optional, String) The name of this provider.
* `public_original_crn` - (Optional, String) The original offering CRN that this publish entry came from.
* `public_publish_approved` - (Deprecated, Optional, Boolean) Indicates if this offering has been approved for use by all IBM Cloud users.
* `publish_approved` - (Optional, Boolean) Offering has been approved to publish to permitted to IBM or Public Catalog.
* `publish_public_crn` - (Optional, String) The crn of the public catalog entry of this offering.
* `rating` - (Optional, List) Repository info for offerings.
Nested scheme for **rating**:
	* `four_star_count` - (Optional, Integer) Four start rating.
	* `one_star_count` - (Optional, Integer) One start rating.
	* `three_star_count` - (Optional, Integer) Three start rating.
	* `two_star_count` - (Optional, Integer) Two start rating.
* `repo_info` - (Optional, List) Repository info for offerings.
Nested scheme for **repo_info**:
	* `token` - (Optional, String) Token for private repos.
	* `type` - (Optional, String) Public or enterprise GitHub.
* `share_enabled` - (Optional, Boolean) Denotes sharing including access list availability of an Offering is enabled.
* `share_with_all` - (Optional, Boolean) Denotes public availability of an Offering - if share_enabled is true.
* `share_with_ibm` - (Optional, Boolean) Denotes IBM employee availability of an Offering - if share_enabled is true.
* `short_description` - (Optional, String) Short description in the requested language.
* `short_description_i18n` - (Optional, Map) A map of translated strings, by language code.
* `support` - (Optional, List) Offering Support information.
Nested scheme for **support**:
	* `locations` - (Optional, List) A list of country codes indicating where support is provided.
	* `process` - (Optional, String) Support process as provided by an ISV.
	* `process_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `support_details` - (Optional, List) A list of support options (e.g. email, phone, slack, other).
	Nested scheme for **support_details**:
		* `availability` - (Optional, List) Times when support is available.
		Nested scheme for **availability**:
			* `always_available` - (Optional, Boolean) Is this support always available.
			* `times` - (Optional, List) A list of support times.
			Nested scheme for **times**:
				* `day` - (Optional, Integer) The day of the week, represented as an integer.
				* `end_time` - (Optional, String) HOURS:MINUTES:SECONDS using 24 hour time (e.g. 8:15:00).
				* `start_time` - (Optional, String) HOURS:MINUTES:SECONDS using 24 hour time (e.g. 8:15:00).
			* `timezone` - (Optional, String) Timezone (e.g. America/New_York).
		* `contact` - (Optional, String) Contact for the current support detail.
		* `response_wait_time` - (Optional, List) Time descriptor.
		Nested scheme for **response_wait_time**:
			* `type` - (Optional, String) Valid values are hour or day.
			* `value` - (Optional, Integer) Amount of time to wait in unit 'type'.
		* `type` - (Optional, String) Type of the current support detail.
	* `support_escalation` - (Optional, List) Support escalation policy.
	Nested scheme for **support_escalation**:
		* `contact` - (Optional, String) Escalation contact.
		* `escalation_wait_time` - (Optional, List) Time descriptor.
		Nested scheme for **escalation_wait_time**:
			* `type` - (Optional, String) Valid values are hour or day.
			* `value` - (Optional, Integer) Amount of time to wait in unit 'type'.
		* `response_wait_time` - (Optional, List) Time descriptor.
		Nested scheme for **response_wait_time**:
			* `type` - (Optional, String) Valid values are hour or day.
			* `value` - (Optional, Integer) Amount of time to wait in unit 'type'.
	* `support_type` - (Optional, String) Support type for this product.
	* `url` - (Optional, String) URL to be displayed in the Consumption UI for getting support on this offering.
* `tags` - (Optional, List) List of tags associated with this catalog.
* `updated` - (Optional, String) The date and time this catalog was last updated.
* `url` - (Optional, String) The url for this specific offering.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cm_offering.
* `offering_id` - (String) unique id.
* `rev` - (String) Cloudant revision.

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
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
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
The `id` property can be formed from `catalog_identifier`, and `offering_id` in the following format:

```
<catalog_identifier>/<offering_id>
```
* `catalog_identifier`: A string. Catalog identifier.
* `offering_id`: A string. Offering identification.

# Syntax
```
$ terraform import ibm_cm_offering.cm_offering <catalog_identifier>/<offering_id>
```
