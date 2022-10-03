---
layout: "ibm"
page_title: "IBM : ibm_cm_catalog"
description: |-
  Manages cm_catalog.
subcategory: "Catalog Management API"
---

# ibm_cm_catalog

Provides a resource for cm_catalog. This allows cm_catalog to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cm_catalog" "cm_catalog" {
  catalog_filters {
		include_all = true
		category_filters = { "key": { example: "object" } }
		id_filters {
			include {
				filter_terms = [ "filter_terms" ]
			}
			exclude {
				filter_terms = [ "filter_terms" ]
			}
		}
  }
  features {
		title = "title"
		title_i18n = { "key": "inner" }
		description = "description"
		description_i18n = { "key": "inner" }
  }
  syndication_settings {
		remove_related_components = true
		clusters {
			region = "region"
			id = "id"
			name = "name"
			resource_group_name = "resource_group_name"
			type = "type"
			namespaces = [ "namespaces" ]
			all_namespaces = true
		}
		history {
			namespaces = [ "namespaces" ]
			clusters {
				region = "region"
				id = "id"
				name = "name"
				resource_group_name = "resource_group_name"
				type = "type"
				namespaces = [ "namespaces" ]
				all_namespaces = true
			}
			last_run = "2021-01-31T09:44:12Z"
		}
		authorization {
			token = "token"
			last_run = "2021-01-31T09:44:12Z"
		}
  }
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `catalog_filters` - (Optional, List) Filters for account and catalog filters.
Nested scheme for **catalog_filters**:
	* `category_filters` - (Optional, Map) Filter against offering properties.
	* `id_filters` - (Optional, List) Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
	Nested scheme for **id_filters**:
		* `exclude` - (Optional, List) Offering filter terms.
		Nested scheme for **exclude**:
			* `filter_terms` - (Optional, List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
		* `include` - (Optional, List) Offering filter terms.
		Nested scheme for **include**:
			* `filter_terms` - (Optional, List) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `include_all` - (Optional, Boolean) -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.
* `catalog_icon_url` - (Optional, String) URL for an icon associated with this catalog.
* `disabled` - (Optional, Boolean) Denotes whether a catalog is disabled.
* `features` - (Optional, List) List of features associated with this catalog.
Nested scheme for **features**:
	* `description` - (Optional, String) Feature description.
	* `description_i18n` - (Optional, Map) A map of translated strings, by language code.
	* `title` - (Optional, String) Heading.
	* `title_i18n` - (Optional, Map) A map of translated strings, by language code.
* `kind` - (Optional, String) Kind of catalog. Supported kinds are offering and vpe.
* `label` - (Optional, String) Display Name in the requested language.
* `label_i18n` - (Optional, Map) A map of translated strings, by language code.
* `metadata` - (Optional, Map) Catalog specific metadata.
* `owning_account` - (Optional, String) Account that owns catalog.
* `resource_group_id` - (Optional, String) Resource group id the catalog is owned by.
* `short_description` - (Optional, String) Description in the requested language.
* `short_description_i18n` - (Optional, Map) A map of translated strings, by language code.
* `syndication_settings` - (Optional, List) Feature information.
Nested scheme for **syndication_settings**:
	* `authorization` - (Optional, List) Feature information.
	Nested scheme for **authorization**:
		* `last_run` - (Optional, String) Date and time last updated.
		* `token` - (Optional, String) Array of syndicated namespaces.
	* `clusters` - (Optional, List) Syndication clusters.
	Nested scheme for **clusters**:
		* `all_namespaces` - (Optional, Boolean) Syndicated to all namespaces on cluster.
		* `id` - (Optional, String) Cluster ID.
		* `name` - (Optional, String) Cluster name.
		* `namespaces` - (Optional, List) Syndicated namespaces.
		* `region` - (Optional, String) Cluster region.
		* `resource_group_name` - (Optional, String) Resource group ID.
		* `type` - (Optional, String) Syndication type.
	* `history` - (Optional, List) Feature information.
	Nested scheme for **history**:
		* `clusters` - (Optional, List) Array of syndicated namespaces.
		Nested scheme for **clusters**:
			* `all_namespaces` - (Optional, Boolean) Syndicated to all namespaces on cluster.
			* `id` - (Optional, String) Cluster ID.
			* `name` - (Optional, String) Cluster name.
			* `namespaces` - (Optional, List) Syndicated namespaces.
			* `region` - (Optional, String) Cluster region.
			* `resource_group_name` - (Optional, String) Resource group ID.
			* `type` - (Optional, String) Syndication type.
		* `last_run` - (Optional, String) Date and time last syndicated.
		* `namespaces` - (Optional, List) Array of syndicated namespaces.
	* `remove_related_components` - (Optional, Boolean) Remove related components.
* `tags` - (Optional, List) List of tags associated with this catalog.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cm_catalog.
* `created` - (String) The date-time this catalog was created.
* `crn` - (String) CRN associated with the catalog.
* `offerings_url` - (String) URL path to offerings.
* `rev` - (String) Cloudant revision.
* `updated` - (String) The date-time this catalog was last updated.
* `url` - (String) The url for this specific catalog.

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

You can import the `ibm_cm_catalog` resource by using `id`. Unique ID.

# Syntax
```
$ terraform import ibm_cm_catalog.cm_catalog <id>
```
