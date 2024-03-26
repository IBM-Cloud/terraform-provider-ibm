# Example for CatalogManagementV1

This example illustrates how to use the CatalogManagementV1

These types of resources are supported:

* ibm_cm_catalog
* ibm_cm_offering
* ibm_cm_version
* cm_offering_instance

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.


## CatalogManagementV1 resources

ibm_cm_catalog resource:

```hcl
resource "ibm_cm_catalog" "cm_catalog_instance" {
  label = var.cm_catalog_label
  label_i18n = var.cm_catalog_label_i18n
  short_description = var.cm_catalog_short_description
  short_description_i18n = var.cm_catalog_short_description_i18n
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
  features = var.cm_catalog_features
  disabled = var.cm_catalog_disabled
  resource_group_id = var.cm_catalog_resource_group_id
  owning_account = var.cm_catalog_owning_account
  catalog_filters = var.cm_catalog_catalog_filters
  kind = var.cm_catalog_kind
  metadata = var.cm_catalog_metadata
}
```
ibm_cm_offering resource:

```hcl
resource "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  url = var.cm_offering_url
  crn = var.cm_offering_crn
  label = var.cm_offering_label
  label_i18n = var.cm_offering_label_i18n
  name = var.cm_offering_name
  offering_icon_url = var.cm_offering_offering_icon_url
  offering_docs_url = var.cm_offering_offering_docs_url
  offering_support_url = var.cm_offering_offering_support_url
  tags = var.cm_offering_tags
  keywords = var.cm_offering_keywords
  rating = var.cm_offering_rating
  created = var.cm_offering_created
  updated = var.cm_offering_updated
  short_description = var.cm_offering_short_description
  short_description_i18n = var.cm_offering_short_description_i18n
  long_description = var.cm_offering_long_description
  long_description_i18n = var.cm_offering_long_description_i18n
  features = var.cm_offering_features
  kinds = var.cm_offering_kinds
  pc_managed = var.cm_offering_pc_managed
  publish_approved = var.cm_offering_publish_approved
  share_with_all = var.cm_offering_share_with_all
  share_with_ibm = var.cm_offering_share_with_ibm
  share_enabled = var.cm_offering_share_enabled
  permit_request_ibm_public_publish = var.cm_offering_permit_request_ibm_public_publish
  ibm_publish_approved = var.cm_offering_ibm_publish_approved
  public_publish_approved = var.cm_offering_public_publish_approved
  public_original_crn = var.cm_offering_public_original_crn
  publish_public_crn = var.cm_offering_publish_public_crn
  portal_approval_record = var.cm_offering_portal_approval_record
  portal_ui_url = var.cm_offering_portal_ui_url
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  catalog_name = var.cm_offering_catalog_name
  metadata = var.cm_offering_metadata
  disclaimer = var.cm_offering_disclaimer
  hidden = var.cm_offering_hidden
  provider = var.cm_offering_provider
  provider_info = var.cm_offering_provider_info
  repo_info = var.cm_offering_repo_info
  image_pull_keys = var.cm_offering_image_pull_keys
  support = var.cm_offering_support
  media = var.cm_offering_media
  deprecate_pending = var.cm_offering_deprecate_pending
  product_kind = var.cm_offering_product_kind
  badges = var.cm_offering_badges
}
```
ibm_cm_version resource:

```hcl
resource "ibm_cm_version" "cm_version_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.offering_id
  tags = var.cm_version_tags
  content = var.cm_version_content
  name = var.cm_version_name
  label = var.cm_version_label
  install_kind = var.cm_version_install_kind
  target_kinds = var.cm_version_target_kinds
  format_kind = var.cm_version_format_kind
  product_kind = var.cm_version_product_kind
  import_sha = var.cm_version_sha
  version = var.cm_version_version
  flavor = var.cm_version_flavor
  import_metadata = var.cm_version_metadata
  working_directory = var.cm_version_working_directory
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
  include_config = var.cm_version_include_config
  is_vsi = var.cm_version_is_vsi
  repotype = var.cm_version_repotype
  x_auth_token = var.cm_version_x_auth_token
}
```
cm_offering_instance resource:

```hcl
resource "cm_offering_instance" "cm_offering_instance_instance" {
  x_auth_refresh_token = var.cm_offering_instance_x_auth_refresh_token
  rev = var.cm_offering_instance_rev
  url = var.cm_offering_instance_url
  crn = var.cm_offering_instance_crn
  label = var.cm_offering_instance_label
  catalog_id = var.cm_offering_instance_catalog_id
  offering_id = var.cm_offering_instance_offering_id
  kind_format = var.cm_offering_instance_kind_format
  version = var.cm_offering_instance_version
  version_id = var.cm_offering_instance_version_id
  cluster_id = var.cm_offering_instance_cluster_id
  cluster_region = var.cm_offering_instance_cluster_region
  cluster_namespaces = var.cm_offering_instance_cluster_namespaces
  cluster_all_namespaces = var.cm_offering_instance_cluster_all_namespaces
  schematics_workspace_id = var.cm_offering_instance_schematics_workspace_id
  install_plan = var.cm_offering_instance_install_plan
  channel = var.cm_offering_instance_channel
  created = var.cm_offering_instance_created
  updated = var.cm_offering_instance_updated
  metadata = var.cm_offering_instance_metadata
  resource_group_id = var.cm_offering_instance_resource_group_id
  location = var.cm_offering_instance_location
  disabled = var.cm_offering_instance_disabled
  account = var.cm_offering_instance_account
  last_operation = var.cm_offering_instance_last_operation
  kind_target = var.cm_offering_instance_kind_target
  sha = var.cm_offering_instance_sha
}
```

## CatalogManagementV1 Data sources

ibm_cm_catalog data source:

```hcl
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
}
```
ibm_cm_offering data source:

```hcl
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_id = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.offering_id
}
```
ibm_cm_version data source:

```hcl
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}
```
cm_offering_instance data source:

```hcl
data "cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
```

## Assumptions

1. TODO

## Notes

1. TODO

## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | 1.13.1 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| label | Display Name in the requested language. | `string` | false |
| label_i18n | A map of translated strings, by language code. | `map(string)` | false |
| short_description | Description in the requested language. | `string` | false |
| short_description_i18n | A map of translated strings, by language code. | `map(string)` | false |
| catalog_icon_url | URL for an icon associated with this catalog. | `string` | false |
| tags | List of tags associated with this catalog. | `list(string)` | false |
| features | List of features associated with this catalog. | `list()` | false |
| disabled | Denotes whether a catalog is disabled. | `bool` | false |
| resource_group_id | Resource group id the catalog is owned by. | `string` | false |
| owning_account | Account that owns catalog. | `string` | false |
| catalog_filters | Filters for account and catalog filters. | `` | false |
| kind | Kind of catalog. Supported kinds are offering and vpe. | `string` | false |
| metadata | Catalog specific metadata. | `map()` | false |
| catalog_id | Catalog identifier. | `string` | true |
| url | The url for this specific offering. | `string` | false |
| crn | The crn for this specific offering. | `string` | false |
| label | Display Name in the requested language. | `string` | false |
| label_i18n | A map of translated strings, by language code. | `map(string)` | false |
| name | The programmatic name of this offering. | `string` | false |
| offering_icon_url | URL for an icon associated with this offering. | `string` | false |
| offering_docs_url | URL for an additional docs with this offering. | `string` | false |
| offering_support_url | [deprecated] - Use offering.support instead.  URL to be displayed in the Consumption UI for getting support on this offering. | `string` | false |
| tags | List of tags associated with this catalog. | `list(string)` | false |
| keywords | List of keywords associated with offering, typically used to search for it. | `list(string)` | false |
| rating | Repository info for offerings. | `` | false |
| created | The date and time this catalog was created. | `` | false |
| updated | The date and time this catalog was last updated. | `` | false |
| short_description | Short description in the requested language. | `string` | false |
| short_description_i18n | A map of translated strings, by language code. | `map(string)` | false |
| long_description | Long description in the requested language. | `string` | false |
| long_description_i18n | A map of translated strings, by language code. | `map(string)` | false |
| features | list of features associated with this offering. | `list()` | false |
| kinds | Array of kind. | `list()` | false |
| pc_managed | Offering is managed by Partner Center. | `bool` | false |
| publish_approved | Offering has been approved to publish to permitted to IBM or Public Catalog. | `bool` | false |
| share_with_all | Denotes public availability of an Offering - if share_enabled is true. | `bool` | false |
| share_with_ibm | Denotes IBM employee availability of an Offering - if share_enabled is true. | `bool` | false |
| share_enabled | Denotes sharing including access list availability of an Offering is enabled. | `bool` | false |
| permit_request_ibm_public_publish | Is it permitted to request publishing to IBM or Public. | `bool` | false |
| ibm_publish_approved | Indicates if this offering has been approved for use by all IBMers. | `bool` | false |
| public_publish_approved | Indicates if this offering has been approved for use by all IBM Cloud users. | `bool` | false |
| public_original_crn | The original offering CRN that this publish entry came from. | `string` | false |
| publish_public_crn | The crn of the public catalog entry of this offering. | `string` | false |
| portal_approval_record | The portal's approval record ID. | `string` | false |
| portal_ui_url | The portal UI URL. | `string` | false |
| catalog_id | The id of the catalog containing this offering. | `string` | false |
| catalog_name | The name of the catalog. | `string` | false |
| metadata | Map of metadata values for this offering. | `map()` | false |
| disclaimer | A disclaimer for this offering. | `string` | false |
| hidden | Determine if this offering should be displayed in the Consumption UI. | `bool` | false |
| provider | Deprecated - Provider of this offering. | `string` | false |
| provider_info | Information on the provider for this offering, or omitted if no provider information is given. | `` | false |
| repo_info | Repository info for offerings. | `` | false |
| image_pull_keys | Image pull keys for this offering. | `list()` | false |
| support | Offering Support information. | `` | false |
| media | A list of media items related to this offering. | `list()` | false |
| deprecate_pending | Deprecation information for an Offering. | `` | false |
| product_kind | The product kind.  Valid values are module, solution, or empty string. | `string` | false |
| badges | A list of badges for this offering. | `list()` | false |
| catalog_id | Catalog identifier. | `string` | true |
| offering_id | Offering identification. | `string` | true |
| tags | Tags array. | `list(string)` | false |
| content | Byte array representing the content to be imported. Only supported for OVA images at this time. | `` | false |
| name | Name of version. Required for virtual server image for VPC. | `string` | false |
| label | Display name of version. Required for virtual server image for VPC. | `string` | false |
| install_kind | Install type. Example: instance. Required for virtual server image for VPC. | `string` | false |
| target_kinds | Deployment target of the content being onboarded. Current valid values are iks, roks, vcenter, power-iaas, terraform, and vpc-x86. Required for virtual server image for VPC. | `list(string)` | false |
| format_kind | Format of content being onboarded. Example: vsi-image. Required for virtual server image for VPC. | `string` | false |
| product_kind | Optional product kind for the software being onboarded.  Valid values are software, module, or solution.  Default value is software. | `string` | false |
| sha | SHA256 fingerprint of the image file. Required for virtual server image for VPC. | `string` | false |
| version | Semantic version of the software being onboarded. Required for virtual server image for VPC. | `string` | false |
| flavor | Version Flavor Information.  Only supported for Product kind Solution. | `` | false |
| metadata | Generic data to be included with content being onboarded. Required for virtual server image for VPC. | `` | false |
| working_directory | Optional - The sub-folder within the specified tgz file that contains the software being onboarded. | `string` | false |
| zipurl | URL path to zip location.  If not specified, must provide content in the body of this call. | `string` | false |
| target_version | The semver value for this new version, if not found in the zip url package content. | `string` | false |
| include_config | Add all possible configuration values to this version when importing. | `bool` | false |
| is_vsi | Indicates that the current terraform template is used to install a virtual server image. | `bool` | false |
| repotype | The type of repository containing this version.  Valid values are 'public_git' or 'enterprise_git'. | `string` | false |
| x_auth_token | Authentication token used to access the specified zip file. | `string` | false |
| x_auth_refresh_token | IAM Refresh token. | `string` | true |
| rev | Cloudant revision. | `string` | false |
| url | url reference to this object. | `string` | false |
| crn | platform CRN for this instance. | `string` | false |
| label | the label for this instance. | `string` | false |
| catalog_id | Catalog ID this instance was created from. | `string` | false |
| offering_id | Offering ID this instance was created from. | `string` | false |
| kind_format | the format this instance has (helm, operator, ova...). | `string` | false |
| version | The version this instance was installed from (semver - not version id). | `string` | false |
| version_id | The version id this instance was installed from (version id - not semver). | `string` | false |
| cluster_id | Cluster ID. | `string` | false |
| cluster_region | Cluster region (e.g., us-south). | `string` | false |
| cluster_namespaces | List of target namespaces to install into. | `list(string)` | false |
| cluster_all_namespaces | designate to install into all namespaces. | `bool` | false |
| schematics_workspace_id | Id of the schematics workspace, for offering instances provisioned through schematics. | `string` | false |
| install_plan | Type of install plan (also known as approval strategy) for operator subscriptions. Can be either automatic, which automatically upgrades operators to the latest in a channel, or manual, which requires approval on the cluster. | `string` | false |
| channel | Channel to pin the operator subscription to. | `string` | false |
| created | date and time create. | `` | false |
| updated | date and time updated. | `` | false |
| metadata | Map of metadata values for this offering instance. | `map()` | false |
| resource_group_id | Id of the resource group to provision the offering instance into. | `string` | false |
| location | String location of OfferingInstance deployment. | `string` | false |
| disabled | Indicates if Resource Controller has disabled this instance. | `bool` | false |
| account | The account this instance is owned by. | `string` | false |
| last_operation | the last operation performed and status. | `` | false |
| kind_target | The target kind for the installed software version. | `string` | false |
| sha | The digest value of the installed software version. | `string` | false |
| catalog_id | Catalog identifier. | `string` | true |
| catalog_id | Catalog identifier. | `string` | true |
| offering_id | Offering identification. | `string` | true |
| version_loc_id | A dotted value of `catalogID`.`versionID`. | `string` | true |
| instance_identifier | Version Instance identifier. | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| ibm_cm_catalog | ibm_cm_catalog object |
| ibm_cm_offering | ibm_cm_offering object |
| ibm_cm_version | ibm_cm_version object |
| cm_offering_instance | cm_offering_instance object |
| ibm_cm_catalog | ibm_cm_catalog object |
| ibm_cm_offering | ibm_cm_offering object |
| ibm_cm_version | ibm_cm_version object |
| cm_offering_instance | cm_offering_instance object |
