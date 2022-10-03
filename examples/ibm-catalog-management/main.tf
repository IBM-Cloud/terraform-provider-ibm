provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
}

// Provision cm_catalog resource instance
resource "ibm_cm_catalog" "cm_catalog_instance" {
  label = var.cm_catalog_label
  label_i18n = var.cm_catalog_label_i18n
  short_description = var.cm_catalog_short_description
  short_description_i18n = var.cm_catalog_short_description_i18n
  catalog_icon_url = var.cm_catalog_catalog_icon_url
  tags = var.cm_catalog_tags
  features {
    title = "title"
    title_i18n = { "key": "inner" }
    description = "description"
    description_i18n = { "key": "inner" }
  }
  disabled = var.cm_catalog_disabled
  resource_group_id = var.cm_catalog_resource_group_id
  owning_account = var.cm_catalog_owning_account
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
  kind = var.cm_catalog_kind
  metadata = var.cm_catalog_metadata
}

// Provision cm_offering resource instance
resource "ibm_cm_offering" "cm_offering_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
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
  rating {
    one_star_count = 1
    two_star_count = 1
    three_star_count = 1
    four_star_count = 1
  }
  created = var.cm_offering_created
  updated = var.cm_offering_updated
  short_description = var.cm_offering_short_description
  short_description_i18n = var.cm_offering_short_description_i18n
  long_description = var.cm_offering_long_description
  long_description_i18n = var.cm_offering_long_description_i18n
  features {
    title = "title"
    title_i18n = { "key": "inner" }
    description = "description"
    description_i18n = { "key": "inner" }
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
      offering_id = ibm_cm_offering.cm_offering.offering_id
      catalog_id = ibm_cm_catalog.cm_catalog.id
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
      version_locator = ibm_cm_offering.cm_offering.offering_id
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
  provider_info {
    id = "id"
    name = "name"
  }
  repo_info {
    token = "token"
    type = "type"
  }
  image_pull_keys {
    name = "name"
    value = "value"
    description = "description"
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
  deprecate_pending {
    deprecate_date = "2021-01-31T09:44:12Z"
    deprecate_state = "deprecate_state"
    description = "description"
  }
  product_kind = var.cm_offering_product_kind
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
}

// Provision cm_version resource instance
resource "ibm_cm_version" "cm_version_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.offering_id
  tags = var.cm_version_tags
  content = var.cm_version_content
  name = var.cm_version_name
  label = var.cm_version_label
  install_kind = var.cm_version_install_kind
  target_kinds = var.cm_version_target_kinds
  format_kind = var.cm_version_format_kind
  product_kind = var.cm_version_product_kind
  sha = var.cm_version_sha
  version = var.cm_version_version
  flavor {
    name = "name"
    label = "label"
    label_i18n = { "key": "inner" }
    index = 1
  }
  metadata {
    operating_system {
      dedicated_host_only = true
      vendor = "vendor"
      name = "name"
      href = "href"
      display_name = "display_name"
      family = "family"
      version = "version"
      architecture = "architecture"
    }
    file {
      size = 1
    }
    minimum_provisioned_size = 1
    images {
      id = "id"
      name = "name"
      region = "region"
    }
  }
  working_directory = var.cm_version_working_directory
  zipurl = var.cm_version_zipurl
  target_version = var.cm_version_target_version
  include_config = var.cm_version_include_config
  is_vsi = var.cm_version_is_vsi
  repotype = var.cm_version_repotype
  x_auth_token = var.cm_version_x_auth_token
}

// Provision cm_offering_instance resource instance
resource "ibm_cm_offering_instance" "cm_offering_instance_instance" {
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
  last_operation {
    operation = "operation"
    state = "state"
    message = "message"
    transaction_id = "transaction_id"
    updated = "2021-01-31T09:44:12Z"
    code = "code"
  }
  kind_target = var.cm_offering_instance_kind_target
  sha = var.cm_offering_instance_sha
}

// Create cm_catalog data source
data "ibm_cm_catalog" "cm_catalog_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
}

// Create cm_offering data source
data "ibm_cm_offering" "cm_offering_instance" {
  catalog_identifier = ibm_cm_catalog.cm_catalog_instance.id
  offering_id = ibm_cm_offering.cm_offering_instance.offering_id
}

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cm_version data source
data "ibm_cm_version" "cm_version_instance" {
  version_loc_id = var.cm_version_version_loc_id
}
*/

// Data source is not linked to a resource instance
// Uncomment if an existing data source instance exists
/*
// Create cm_offering_instance data source
data "ibm_cm_offering_instance" "cm_offering_instance_instance" {
  instance_identifier = var.cm_offering_instance_instance_identifier
}
*/
