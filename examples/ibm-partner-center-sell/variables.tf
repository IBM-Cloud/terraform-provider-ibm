variable "ibmcloud_api_key" {
  description = "IBM Cloud API key"
  type        = string
}

// Resource arguments for onboarding_resource_broker
variable "onboarding_resource_broker_env" {
  description = "The environment to fetch this object from."
  type        = string
  default     = "env"
}
variable "onboarding_resource_broker_auth_username" {
  description = "The authentication username to reach the broker."
  type        = string
  default     = "apikey"
}
variable "onboarding_resource_broker_auth_password" {
  description = "The authentication password to reach the broker."
  type        = string
  default     = "auth_password"
}
variable "onboarding_resource_broker_auth_scheme" {
  description = "The supported authentication scheme for the broker."
  type        = string
  default     = "bearer"
}
variable "onboarding_resource_broker_resource_group_crn" {
  description = "The cloud resource name of the resource group."
  type        = string
  default     = "crn:v1:bluemix:public:resource-controller::a/4a5c3c51b97a446fbb1d0e1ef089823b::resource-group:4fae20bd538a4a738475350dfdc1596f"
}
variable "onboarding_resource_broker_state" {
  description = "The state of the broker."
  type        = string
  default     = "active"
}
variable "onboarding_resource_broker_broker_url" {
  description = "The URL associated with the broker application."
  type        = string
  default     = "https://broker-url-for-my-service.com"
}
variable "onboarding_resource_broker_allow_context_updates" {
  description = "Whether the resource controller will call the broker for any context changes to the instance. Currently, the only context related change is an instance name update."
  type        = bool
  default     = false
}
variable "onboarding_resource_broker_catalog_type" {
  description = "To enable the provisioning of your broker, set this parameter value to `service`."
  type        = string
  default     = "service"
}
variable "onboarding_resource_broker_type" {
  description = "The type of the provisioning model."
  type        = string
  default     = "provision_through"
}
variable "onboarding_resource_broker_name" {
  description = "The name of the broker."
  type        = string
  default     = "brokername"
}
variable "onboarding_resource_broker_region" {
  description = "The region where the pricing plan is available."
  type        = string
  default     = "global"
}

// Resource arguments for onboarding_catalog_deployment
variable "onboarding_catalog_deployment_product_id" {
  description = "The unique ID of the resource."
  type        = string
  default     = "product_id"
}
variable "onboarding_catalog_deployment_catalog_product_id" {
  description = "The unique ID of this global catalog product."
  type        = string
  default     = "catalog_product_id"
}
variable "onboarding_catalog_deployment_catalog_plan_id" {
  description = "The unique ID of this global catalog plan."
  type        = string
  default     = "catalog_plan_id"
}
variable "onboarding_catalog_deployment_env" {
  description = "The environment to fetch this object from."
  type        = string
  default     = "env"
}
variable "onboarding_catalog_deployment_object_id" {
  description = "The desired ID of the global catalog object."
  type        = string
  default     = "object_id"
}
variable "onboarding_catalog_deployment_name" {
  description = "The programmatic name of this deployment."
  type        = string
  default     = "deployment-eu-de"
}
variable "onboarding_catalog_deployment_active" {
  description = "Whether the service is active."
  type        = bool
  default     = true
}
variable "onboarding_catalog_deployment_disabled" {
  description = "Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled."
  type        = bool
  default     = false
}
variable "onboarding_catalog_deployment_kind" {
  description = "The kind of the global catalog object."
  type        = string
  default     = "deployment"
}
variable "onboarding_catalog_deployment_tags" {
  description = "A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog."
  type        = list(string)
  default     = ["eu-gb"]
}

// Resource arguments for onboarding_catalog_plan
variable "onboarding_catalog_plan_product_id" {
  description = "The unique ID of the resource."
  type        = string
  default     = "product_id"
}
variable "onboarding_catalog_plan_catalog_product_id" {
  description = "The unique ID of this global catalog product."
  type        = string
  default     = "catalog_product_id"
}
variable "onboarding_catalog_plan_env" {
  description = "The environment to fetch this object from."
  type        = string
  default     = "env"
}
variable "onboarding_catalog_plan_object_id" {
  description = "The desired ID of the global catalog object."
  type        = string
  default     = "object_id"
}
variable "onboarding_catalog_plan_name" {
  description = "The programmatic name of this plan."
  type        = string
  default     = "free-plan2"
}
variable "onboarding_catalog_plan_active" {
  description = "Whether the service is active."
  type        = bool
  default     = true
}
variable "onboarding_catalog_plan_disabled" {
  description = "Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled."
  type        = bool
  default     = false
}
variable "onboarding_catalog_plan_kind" {
  description = "The kind of the global catalog object."
  type        = string
  default     = "plan"
}
variable "onboarding_catalog_plan_tags" {
  description = "A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog."
  type        = list(string)
  default     = ["ibm_created"]
}
variable "onboarding_catalog_plan_pricing_tags" {
  description = "A list of tags that carry information about the pricing information of your product."
  type        = list(string)
  default     = [ "pricing_tags" ]
}

// Resource arguments for onboarding_catalog_product
variable "onboarding_catalog_product_product_id" {
  description = "The unique ID of the resource."
  type        = string
  default     = "product_id"
}
variable "onboarding_catalog_product_env" {
  description = "The environment to fetch this object from."
  type        = string
  default     = "env"
}
variable "onboarding_catalog_product_object_id" {
  description = "The desired ID of the global catalog object."
  type        = string
  default     = "object_id"
}
variable "onboarding_catalog_product_name" {
  description = "The programmatic name of this product."
  type        = string
  default     = "1p-service-08-06"
}
variable "onboarding_catalog_product_active" {
  description = "Whether the service is active."
  type        = bool
  default     = true
}
variable "onboarding_catalog_product_disabled" {
  description = "Determines the global visibility for the catalog entry, and its children. If it is not enabled, all plans are disabled."
  type        = bool
  default     = false
}
variable "onboarding_catalog_product_kind" {
  description = "The kind of the global catalog object."
  type        = string
  default     = "service"
}
variable "onboarding_catalog_product_tags" {
  description = "A list of tags that carry information about your product. These tags can be used to find your product in the IBM Cloud catalog."
  type        = list(string)
  default     = ["keyword","support_ibm"]
}

// Resource arguments for onboarding_iam_registration
variable "onboarding_iam_registration_product_id" {
  description = "The unique ID of the resource."
  type        = string
  default     = "product_id"
}
variable "onboarding_iam_registration_env" {
  description = "The environment to fetch this object from."
  type        = string
  default     = "env"
}
variable "onboarding_iam_registration_name" {
  description = "The IAM registration name, which must be the programmatic name of the product."
  type        = string
  default     = "pet-store"
}
variable "onboarding_iam_registration_enabled" {
  description = "Whether the service is enabled or disabled for IAM."
  type        = bool
  default     = true
}
variable "onboarding_iam_registration_service_type" {
  description = "The type of the service."
  type        = string
  default     = "service"
}
variable "onboarding_iam_registration_additional_policy_scopes" {
  description = "List of additional policy scopes."
  type        = list(string)
  default     = ["pet-store"]
}
variable "onboarding_iam_registration_parent_ids" {
  description = "The list of parent IDs for product access management."
  type        = list(string)
  default     = []
}
variable "onboarding_iam_registration_supported_action_control" {
  description = "The list that indicates which actions are part of the service restrictions."
  type        = list(string)
  default     = [ "supported_action_control" ]
}

// Resource arguments for onboarding_product
variable "onboarding_product_type" {
  description = "The type of the product."
  type        = string
  default     = "service"
}
variable "onboarding_product_eccn_number" {
  description = "The Export Control Classification Number of your product."
  type        = string
  default     = "eccn_number"
}
variable "onboarding_product_ero_class" {
  description = "The ERO class of your product."
  type        = string
  default     = "ero_class"
}
variable "onboarding_product_unspsc" {
  description = "The United Nations Standard Products and Services Code of your product."
  type        = number
  default     = 1.0
}
variable "onboarding_product_tax_assessment" {
  description = "The tax assessment type of your product."
  type        = string
  default     = "tax_assessment"
}

// Resource arguments for onboarding_registration
variable "onboarding_registration_account_id" {
  description = "The ID of your account."
  type        = string
  default     = "4a5c3c51b97a446fbb1d0e1ef089823b"
}
variable "onboarding_registration_company_name" {
  description = "The name of your company that is displayed in the IBM Cloud catalog."
  type        = string
  default     = "Beautiful Company"
}
variable "onboarding_registration_default_private_catalog_id" {
  description = "The default private catalog in which products are created."
  type        = string
  default     = "default_private_catalog_id"
}
variable "onboarding_registration_provider_access_group" {
  description = "The onboarding access group for your team."
  type        = string
  default     = "provider_access_group"
}
