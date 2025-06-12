---
layout: "ibm"
page_title: "IBM : ibm_cm_account"
description: |-
  Manages cm_account.
subcategory: "Catalog Management"
---

# ibm_cm_account

Create, update, and delete cm_accounts with this resource.

## Example Usage

```hcl
resource "ibm_cm_account" "cm_account_instance" {
  region_filter = "geo:eu"

  account_filters {
    category_filters {
      category_name = "provider"
      filter {
        filter_terms = [
          "ibm_third_party",
        ]
      }
      include = false
    }
    category_filters {
      category_name = "category"
      filter {
        filter_terms = [
          "watson",
          "ai",
          "blockchain",
        ]
      }
      include = false
    }
    id_filters {
      exclude {}
      include {
        filter_terms = [
          "9dcb8ea2-30b4-4adf-8821-0d35f0a9d74f-global",
        ]
      }
    }
    include_all = true
  }
}
```


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cm_account.
* `account_filters` - (List, Optional) Filters for account and catalog filters.
Nested schema for **account_filters**:
	* `category_filters` - (List, Optional) Filter against offering properties.
	Nested schema for **category_filters**:
    	* `category_name` - (String, Required) Name of the category.
    	* `include` -  (Boolean, Optional) Whether to include the category in the catalog filter.
    	* `filter` - (List, Optional) Filter terms related to the category.
		Nested schema for **filter**:
			* `filter_terms` - (List, Optional) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `id_filters` - (List, Optional) Filter on offering ID's. There is an include filter and an exclule filter. Both can be set.
	Nested schema for **id_filters**:
		* `exclude` - (List, Optional) Offering filter terms.
		Nested schema for **exclude**:
			* `filter_terms` - (List, Optional) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
		* `include` - (List, Optional) Offering filter terms.
		Nested schema for **include**:
			* `filter_terms` - (List, Optional) List of values to match against. If include is true, then if the offering has one of the values then the offering is included. If include is false, then if the offering has one of the values then the offering is excluded.
	* `include_all` - (Boolean, Optional) -> true - Include all of the public catalog when filtering. Further settings will specifically exclude some offerings. false - Exclude all of the public catalog when filtering. Further settings will specifically include some offerings.
* `hide_ibm_cloud_catalog` - (Boolean, Optional) Hide the public catalog in this account.
* `region_filter` - (String, Optional) Region filter string.
* `rev` - (String) Cloudant revision.


## Import

You can import the `ibm_cm_account` resource by using `id`. Account identification.

# Syntax
<pre>
$ terraform import ibm_cm_account.cm_account &lt;id&gt;
</pre>
