provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision onboarding_resource_broker resource instance
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

// Provision onboarding_catalog_deployment resource instance
resource "ibm_onboarding_catalog_deployment" "onboarding_catalog_deployment_instance" {
  product_id = var.onboarding_catalog_deployment_product_id
  catalog_product_id = var.onboarding_catalog_deployment_catalog_product_id
  catalog_plan_id = var.onboarding_catalog_deployment_catalog_plan_id
  env = var.onboarding_catalog_deployment_env
  name = var.onboarding_catalog_deployment_name
  active = var.onboarding_catalog_deployment_active
  disabled = var.onboarding_catalog_deployment_disabled
  kind = var.onboarding_catalog_deployment_kind
  overview_ui {
    en {
      display_name = "display_name"
      description = "description"
      long_description = "long_description"
    }
  }
  tags = var.onboarding_catalog_deployment_tags
  object_provider {
    name = "name"
    email = "email"
  }
  metadata {
    rc_compatible = true
    ui {
      strings {
        en {
          bullets {
            description = "description"
            description_i18n = { "key" = "inner" }
            title = "title"
            title_i18n = { "key" = "inner" }
          }
          media {
            caption = "caption"
            caption_i18n = { "key" = "inner" }
            thumbnail = "thumbnail"
            type = "image"
            url = "url"
          }
        }
      }
      urls {
        doc_url = "doc_url"
        terms_url = "terms_url"
      }
      hidden = true
      side_by_side_index = 1.0
    }
    service {
      rc_provisionable = true
      iam_compatible = true
    }
  }
}

// Provision onboarding_catalog_plan resource instance
resource "ibm_onboarding_catalog_plan" "onboarding_catalog_plan_instance" {
  product_id = var.onboarding_catalog_plan_product_id
  catalog_product_id = var.onboarding_catalog_plan_catalog_product_id
  env = var.onboarding_catalog_plan_env
  name = var.onboarding_catalog_plan_name
  active = var.onboarding_catalog_plan_active
  disabled = var.onboarding_catalog_plan_disabled
  kind = var.onboarding_catalog_plan_kind
  overview_ui {
    en {
      display_name = "display_name"
      description = "description"
      long_description = "long_description"
    }
  }
  tags = var.onboarding_catalog_plan_tags
  object_provider {
    name = "name"
    email = "email"
  }
  metadata {
    rc_compatible = true
    ui {
      strings {
        en {
          bullets {
            description = "description"
            description_i18n = { "key" = "inner" }
            title = "title"
            title_i18n = { "key" = "inner" }
          }
          media {
            caption = "caption"
            caption_i18n = { "key" = "inner" }
            thumbnail = "thumbnail"
            type = "image"
            url = "url"
          }
        }
      }
      urls {
        doc_url = "doc_url"
        terms_url = "terms_url"
      }
      hidden = true
      side_by_side_index = 1.0
    }
    pricing {
      type = "free"
      origin = "global_catalog"
    }
  }
}

// Provision onboarding_catalog_product resource instance
resource "ibm_onboarding_catalog_product" "onboarding_catalog_product_instance" {
  product_id = var.onboarding_catalog_product_product_id
  env = var.onboarding_catalog_product_env
  name = var.onboarding_catalog_product_name
  active = var.onboarding_catalog_product_active
  disabled = var.onboarding_catalog_product_disabled
  kind = var.onboarding_catalog_product_kind
  overview_ui {
    en {
      display_name = "display_name"
      description = "description"
      long_description = "long_description"
    }
  }
  tags = var.onboarding_catalog_product_tags
  images {
    image = "image"
  }
  object_provider {
    name = "name"
    email = "email"
  }
  metadata {
    rc_compatible = true
    ui {
      strings {
        en {
          bullets {
            description = "description"
            description_i18n = { "key" = "inner" }
            title = "title"
            title_i18n = { "key" = "inner" }
          }
          media {
            caption = "caption"
            caption_i18n = { "key" = "inner" }
            thumbnail = "thumbnail"
            type = "image"
            url = "url"
          }
        }
      }
      urls {
        doc_url = "doc_url"
        terms_url = "terms_url"
      }
      hidden = true
      side_by_side_index = 1.0
    }
    service {
      rc_provisionable = true
      iam_compatible = true
    }
    other {
      pc {
        support {
          url = "url"
          status_url = "status_url"
          locations = [ "locations" ]
          languages = [ "languages" ]
          process = "process"
          process_i18n = { "key" = "inner" }
          support_type = "community"
          support_escalation {
            contact = "contact"
            escalation_wait_time {
              value = 1.0
              type = "type"
            }
            response_wait_time {
              value = 1.0
              type = "type"
            }
          }
          support_details {
            type = "support_site"
            contact = "contact"
            response_wait_time {
              value = 1.0
              type = "type"
            }
            availability {
              times {
                day = 1.0
                start_time = "start_time"
                end_time = "end_time"
              }
              timezone = "timezone"
              always_available = true
            }
          }
        }
      }
    }
  }
}

// Provision onboarding_iam_registration resource instance
resource "ibm_onboarding_iam_registration" "onboarding_iam_registration_instance" {
  product_id = var.onboarding_iam_registration_product_id
  env = var.onboarding_iam_registration_env
  name = var.onboarding_iam_registration_name
  enabled = var.onboarding_iam_registration_enabled
  service_type = var.onboarding_iam_registration_service_type
  actions {
    id = "id"
    roles = [ "roles" ]
    description {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    display_name {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    options {
      hidden = true
    }
  }
  additional_policy_scopes = var.onboarding_iam_registration_additional_policy_scopes
  display_name {
    default = "default"
    en = "en"
    de = "de"
    es = "es"
    fr = "fr"
    it = "it"
    ja = "ja"
    ko = "ko"
    pt_br = "pt_br"
    zh_tw = "zh_tw"
    zh_cn = "zh_cn"
  }
  parent_ids = var.onboarding_iam_registration_parent_ids
  resource_hierarchy_attribute {
    key = "key"
    value = "value"
  }
  supported_anonymous_accesses {
    attributes {
      account_id = "account_id"
      service_name = "service_name"
    }
    roles = [ "roles" ]
  }
  supported_attributes {
    key = "key"
    options {
      operators = [ "stringEquals" ]
      hidden = true
      supported_attributes = [ "supported_attributes" ]
      policy_types = [ "access" ]
      is_empty_value_supported = true
      is_string_exists_false_value_supported = true
      key = "key"
      resource_hierarchy {
        key {
          key = "key"
          value = "value"
        }
        value {
          key = "key"
        }
      }
    }
    display_name {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    description {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    ui {
      input_type = "input_type"
      input_details {
        type = "type"
        values {
          value = "value"
          display_name {
            default = "default"
            en = "en"
            de = "de"
            es = "es"
            fr = "fr"
            it = "it"
            ja = "ja"
            ko = "ko"
            pt_br = "pt_br"
            zh_tw = "zh_tw"
            zh_cn = "zh_cn"
          }
        }
        gst {
          query = "query"
          value_property_name = "value_property_name"
          label_property_name = "label_property_name"
          input_option_label = "input_option_label"
        }
        url {
          url_endpoint = "url_endpoint"
          input_option_label = "input_option_label"
        }
      }
    }
  }
  supported_authorization_subjects {
    attributes {
      service_name = "service_name"
      resource_type = "resource_type"
    }
    roles = [ "roles" ]
  }
  supported_roles {
    id = "id"
    description {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    display_name {
      default = "default"
      en = "en"
      de = "de"
      es = "es"
      fr = "fr"
      it = "it"
      ja = "ja"
      ko = "ko"
      pt_br = "pt_br"
      zh_tw = "zh_tw"
      zh_cn = "zh_cn"
    }
    options {
      access_policy = { "key" = "inner" }
      policy_type = [ "access" ]
      account_type = "enterprise"
    }
  }
  supported_network {
    environment_attributes {
      key = "key"
      values = [ "values" ]
      options {
        hidden = true
      }
    }
  }
}

// Provision onboarding_product resource instance
resource "ibm_onboarding_product" "onboarding_product_instance" {
  type = var.onboarding_product_type
  primary_contact {
    name = "name"
    email = "email"
  }
  eccn_number = var.onboarding_product_eccn_number
  ero_class = var.onboarding_product_ero_class
  unspsc = var.onboarding_product_unspsc
  tax_assessment = var.onboarding_product_tax_assessment
  support {
    escalation_contacts {
      name = "name"
      email = "email"
      role = "role"
    }
  }
}

// Provision onboarding_registration resource instance
resource "ibm_onboarding_registration" "onboarding_registration_instance" {
  account_id = var.onboarding_registration_account_id
  company_name = var.onboarding_registration_company_name
  primary_contact {
    name = "name"
    email = "email"
  }
  default_private_catalog_id = var.onboarding_registration_default_private_catalog_id
  provider_access_group = var.onboarding_registration_provider_access_group
}
