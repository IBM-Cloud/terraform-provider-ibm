# Examples for Partner Center Sell

These examples illustrate how to use the resources and data sources associated with Partner Center Sell.

The following resources are supported:
* ibm_onboarding_resource_broker
* ibm_onboarding_catalog_deployment
* ibm_onboarding_catalog_plan
* ibm_onboarding_catalog_product
* ibm_onboarding_iam_registration
* ibm_onboarding_product
* ibm_onboarding_registration

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Partner Center Sell resources

### Resource: ibm_onboarding_resource_broker

```hcl
resource "ibm_onboarding_resource_broker" "onboarding_resource_broker_instance" {
  env = var.onboarding_resource_broker_env
  auth_username = var.onboarding_resource_broker_auth_username
  auth_password = var.onboarding_resource_broker_auth_password
  auth_scheme = var.onboarding_resource_broker_auth_scheme
  resource_group_crn = var.onboarding_resource_broker_resource_group_crn
  state = var.onboarding_resource_broker_state
  broker_url = var.onboarding_resource_broker_broker_url
  allow_context_updates = var.onboarding_resource_broker_allow_context_updates
  catalog_type = var.onboarding_resource_broker_catalog_type
  type = var.onboarding_resource_broker_type
  name = var.onboarding_resource_broker_name
  region = var.onboarding_resource_broker_region
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| env | The environment to fetch this object from. | `string` | false |
| auth_username | The authentication username to reach the broker. | `string` | false |
| auth_password | The authentication password to reach the broker. | `string` | false |
| auth_scheme | The supported authentication scheme for the broker. | `string` | true |
| resource_group_crn | The cloud resource name of the resource group. | `string` | false |
| state | The state of the broker. | `string` | false |
| broker_url | The URL associated with the broker application. | `string` | true |
| allow_context_updates | Whether the resource controller will call the broker for any context changes to the instance. Currently, the only context related change is an instance name update. | `bool` | false |
| catalog_type | To enable the provisioning of your broker, set this parameter value to `service`. | `string` | false |
| type | The type of the provisioning model. | `string` | true |
| name | The name of the broker. | `string` | true |
| region | The region where the pricing plan is available. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| account_id | The ID of the account in which you manage the broker. |
| crn | The cloud resource name (CRN) of the broker. |
| created_at | The time when the service broker was created. |
| updated_at | The time when the service broker was updated. |
| deleted_at | The time when the service broker was deleted. |
| created_by | The details of the user who created this broker. |
| updated_by | The details of the user who updated this broker. |
| deleted_by | The details of the user who deleted this broker. |
| guid | The globally unique identifier of the broker. |
| url | The URL associated with the broker. |

### Resource: ibm_onboarding_catalog_deployment

```hcl
resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
  product_id = var.onboarding_catalog_deployment_product_id
  catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
  catalog_plan_id = ibm_onboarding_catalog_plan.onboarding_catalog_plan_instance.onboarding_catalog_plan_id
  env = var.onboarding_catalog_deployment_env
  object_id = var.onboarding_catalog_deployment_object_id
  name = var.onboarding_catalog_deployment_name
  active = var.onboarding_catalog_deployment_active
  disabled = var.onboarding_catalog_deployment_disabled
  kind = var.onboarding_catalog_deployment_kind
  overview_ui = var.onboarding_catalog_deployment_overview_ui
  tags = var.onboarding_catalog_deployment_tags
  object_provider = var.onboarding_catalog_deployment_object_provider
  metadata = var.onboarding_catalog_deployment_metadata
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| product_id | The unique ID of the resource. | `string` | true |
| catalog_product_id | The unique ID of this global catalog product. | `string` | true |
| catalog_plan_id | The unique ID of this global catalog plan. | `string` | true |
| env | The environment to fetch this object from. | `string` | false |
| object_id | The desired ID of the global catalog object. | `string` | false |
| name | The programmatic name of this deployment. | `string` | true |
| active | Whether the service is active. | `bool` | true |
| disabled | Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled. | `bool` | true |
| kind | The kind of the global catalog object. | `string` | true |
| overview_ui | The object that contains the service details from the Overview page in global catalog. | `` | false |
| tags | A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog. | `list(string)` | false |
| object_provider | The provider or owner of the product. | `` | true |
| metadata | Global catalog deployment metadata. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| geo_tags |  |
| url | The global catalog URL of your product. |
| catalog_deployment_id | The ID of a global catalog object. |

### Resource: ibm_onboarding_catalog_plan

```hcl
resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
  product_id = var.onboarding_catalog_plan_product_id
  catalog_product_id = ibm_onboarding_catalog_product.onboarding_catalog_product_instance.onboarding_catalog_product_id
  env = var.onboarding_catalog_plan_env
  object_id = var.onboarding_catalog_plan_object_id
  name = var.onboarding_catalog_plan_name
  active = var.onboarding_catalog_plan_active
  disabled = var.onboarding_catalog_plan_disabled
  kind = var.onboarding_catalog_plan_kind
  overview_ui = var.onboarding_catalog_plan_overview_ui
  tags = var.onboarding_catalog_plan_tags
  pricing_tags = var.onboarding_catalog_plan_pricing_tags
  object_provider = var.onboarding_catalog_plan_object_provider
  metadata = var.onboarding_catalog_plan_metadata
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| product_id | The unique ID of the resource. | `string` | true |
| catalog_product_id | The unique ID of this global catalog product. | `string` | true |
| env | The environment to fetch this object from. | `string` | false |
| object_id | The desired ID of the global catalog object. | `string` | false |
| name | The programmatic name of this plan. | `string` | true |
| active | Whether the service is active. | `bool` | true |
| disabled | Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled. | `bool` | true |
| kind | The kind of the global catalog object. | `string` | true |
| overview_ui | The object that contains the service details from the Overview page in global catalog. | `` | false |
| tags | A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog. | `list(string)` | false |
| pricing_tags | A list of tags that carry information about the pricing information of your product. | `list(string)` | false |
| object_provider | The provider or owner of the product. | `` | true |
| metadata | Global catalog plan metadata. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| geo_tags |  |
| url | The global catalog URL of your product. |
| catalog_plan_id | The ID of a global catalog object. |

### Resource: ibm_onboarding_catalog_product

```hcl
resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
  product_id = var.onboarding_catalog_product_product_id
  env = var.onboarding_catalog_product_env
  object_id = var.onboarding_catalog_product_object_id
  name = var.onboarding_catalog_product_name
  active = var.onboarding_catalog_product_active
  disabled = var.onboarding_catalog_product_disabled
  kind = var.onboarding_catalog_product_kind
  overview_ui = var.onboarding_catalog_product_overview_ui
  tags = var.onboarding_catalog_product_tags
  images = var.onboarding_catalog_product_images
  object_provider = var.onboarding_catalog_product_object_provider
  metadata = var.onboarding_catalog_product_metadata
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| product_id | The unique ID of the resource. | `string` | true |
| env | The environment to fetch this object from. | `string` | false |
| object_id | The desired ID of the global catalog object. | `string` | false |
| name | The programmatic name of this product. | `string` | true |
| active | Whether the service is active. | `bool` | true |
| disabled | Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled. | `bool` | true |
| kind | The kind of the global catalog object. | `string` | true |
| overview_ui | The object that contains the service details from the Overview page in global catalog. | `` | false |
| tags | A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog. | `list(string)` | true |
| images | Images from the global catalog entry that help illustrate the service. | `` | false |
| object_provider | The provider or owner of the product. | `` | true |
| metadata | The global catalog service metadata object. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| geo_tags |  |
| pricing_tags | A list of tags that carry information about the pricing information of your product. |
| url | The global catalog URL of your product. |
| group | Flag for group tile legacy service. |
| catalog_product_id | The ID of a global catalog object. |

### Resource: ibm_onboarding_iam_registration

```hcl
resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
  product_id = var.onboarding_iam_registration_product_id
  env = var.onboarding_iam_registration_env
  name = var.onboarding_iam_registration_name
  enabled = var.onboarding_iam_registration_enabled
  service_type = var.onboarding_iam_registration_service_type
  actions = var.onboarding_iam_registration_actions
  additional_policy_scopes = var.onboarding_iam_registration_additional_policy_scopes
  display_name = var.onboarding_iam_registration_display_name
  parent_ids = var.onboarding_iam_registration_parent_ids
  resource_hierarchy_attribute = var.onboarding_iam_registration_resource_hierarchy_attribute
  supported_anonymous_accesses = var.onboarding_iam_registration_supported_anonymous_accesses
  supported_attributes = var.onboarding_iam_registration_supported_attributes
  supported_authorization_subjects = var.onboarding_iam_registration_supported_authorization_subjects
  supported_roles = var.onboarding_iam_registration_supported_roles
  supported_network = var.onboarding_iam_registration_supported_network
  supported_action_control = var.onboarding_iam_registration_supported_action_control
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| product_id | The unique ID of the resource. | `string` | true |
| env | The environment to fetch this object from. | `string` | false |
| name | The IAM registration name, which must be the programmatic name of the product. | `string` | true |
| enabled | Whether the service is enabled or disabled for IAM. | `bool` | false |
| service_type | The type of the service. | `string` | false |
| actions | The product access management action. | `list()` | false |
| additional_policy_scopes | List of additional policy scopes. | `list(string)` | false |
| display_name | The display name of the object. | `` | false |
| parent_ids | The list of parent IDs for product access management. | `list(string)` | false |
| resource_hierarchy_attribute | The resource hierarchy key-value pair for composite services. | `` | false |
| supported_anonymous_accesses | The list of supported anonymous accesses. | `list()` | false |
| supported_attributes | The list of supported attributes. | `list()` | false |
| supported_authorization_subjects | The list of supported authorization subjects. | `list()` | false |
| supported_roles | The list of roles that you can use to assign access. | `list()` | false |
| supported_network | The registration of set of endpoint types that are supported by your service in the `networkType` environment attribute. This constrains the context-based restriction rules specific to the service such that they describe access restrictions on only this set of endpoints. | `` | false |
| supported_action_control | The list that indicates which actions are part of the service restrictions. | `list(string)` | false |

### Resource: ibm_onboarding_product

```hcl
resource "ibm_onboarding_product" "onboarding_product_instance" {
  type = var.onboarding_product_type
  primary_contact = var.onboarding_product_primary_contact
  eccn_number = var.onboarding_product_eccn_number
  ero_class = var.onboarding_product_ero_class
  unspsc = var.onboarding_product_unspsc
  tax_assessment = var.onboarding_product_tax_assessment
  support = var.onboarding_product_support
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| type | The type of the product. | `string` | true |
| primary_contact | The primary contact for your product. | `` | true |
| eccn_number | The Export Control Classification Number of your product. | `string` | false |
| ero_class | The ERO class of your product. | `string` | false |
| unspsc | The United Nations Standard Products and Services Code of your product. | `number` | false |
| tax_assessment | The tax assessment type of your product. | `string` | false |
| support | The support information that is not displayed in the catalog, but available in ServiceNow. | `` | false |

#### Outputs

| Name | Description |
|------|-------------|
| account_id | The IBM Cloud account ID of the provider. |
| private_catalog_id | The ID of the private catalog that contains the product. Only applicable for software type products. |
| private_catalog_offering_id | The ID of the linked private catalog product. Only applicable for software type products. |
| global_catalog_offering_id | The ID of a global catalog object. |
| staging_global_catalog_offering_id | The ID of a global catalog object. |
| approver_resource_id | The ID of the approval workflow of your product. |
| iam_registration_id | IAM registration identifier. |

### Resource: ibm_onboarding_registration

```hcl
resource "ibm_onboarding_registration" "onboarding_registration_instance" {
  account_id = var.onboarding_registration_account_id
  company_name = var.onboarding_registration_company_name
  primary_contact = var.onboarding_registration_primary_contact
  default_private_catalog_id = var.onboarding_registration_default_private_catalog_id
  provider_access_group = var.onboarding_registration_provider_access_group
}
```

#### Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| account_id | The ID of your account. | `string` | true |
| company_name | The name of your company that is displayed in the IBM Cloud catalog. | `string` | true |
| primary_contact | The primary contact for your product. | `` | true |
| default_private_catalog_id | The default private catalog in which products are created. | `string` | false |
| provider_access_group | The onboarding access group for your team. | `string` | false |

#### Outputs

| Name | Description |
|------|-------------|
| created_at | The time when the registration was created. |
| updated_at | The time when the registration was updated. |


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
