
# This is not functional terraform code. It is intended as a template for users to remove
# unneeded scenarios and edit the other sections.

# Replace the resource group name with the one in which your resources should be created
data "ibm_resource_group" "group" {
  name = "Default"
}

#### Scenario 1: Create an Event Streams standard-plan service instance.
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "standard"
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id

  timeouts {
    create = "15m"
    update = "15m"
    delete = "15m"
  }
}

#### Scenario 2: Create an Event Streams enterprise service instance with non-default attributes
resource "ibm_resource_instance" "es_instance_2" {
  name              = "terraform-integration-2"
  service           = "messagehub"
  plan              = "enterprise-3nodes-2tb"
  location          = "us-east"
  resource_group_id = data.ibm_resource_group.group.id

  parameters = {
    throughput           = "300"
    storage_size         = "4096"
    service-endpoints    = "private"
    private_ip_allowlist = "[10.0.0.0/32,10.0.0.1/32]"
    metrics              = "[topic,consumers]"
  }

  timeouts {
    create = "330m" # 5.5h
    update = "210m" # 3.5h
    delete = "1h"
  }
}

#### Scenario 3: Create a topic on an existing Event Streams instance.

# the existing instance
data "ibm_resource_instance" "es_instance_3" {
  name              = "terraform-integration-3"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_topic" "es_topic_3" {
  resource_instance_id = data.ibm_resource_instance.es_instance_3.id
  name                 = "my-es-topic"
  partitions           = 1
  config = {
    "cleanup.policy"  = "compact,delete"
    "retention.ms"    = "86400000"
    "retention.bytes" = "1073741824"
    "segment.bytes"   = "536870912"
  }
}

#### Scenario 4: Create a schema on an existing Event Streams Enterprise instance

data "ibm_resource_instance" "es_instance_4" {
  name              = "terraform-integration-4"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_schema" "es_schema" {
  resource_instance_id = data.ibm_resource_instance.es_instance_4.id
  schema_id            = "tf_schema"
  schema               = <<SCHEMA
   {
           "type": "record",
           "name": "record_name",
           "fields" : [
             {"name": "value_1", "type": "long"},
             {"name": "value_2", "type": "string"}
           ]
         }
  SCHEMA
}


#### Scenario 5: Apply access tags to an Event Streams service instance
data "ibm_resource_instance" "es_instance_5" {
  name              = "terraform-integration-5"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_resource_tag" "tag_example_on_es" {
  tags        = ["example:tag"]
  tag_type    = "access"
  resource_id = data.ibm_resource_instance.es_instance_5.id
}

#### Scenario 6: Create a target Event Streams service instance with mirroring enabled and its mirroring config
data "ibm_resource_instance" "es_instance_source" {
  name              = "terraform-integration-source"
  resource_group_id = data.ibm_resource_group.group.id
}
# setup s2s at service level for mirroring to work
resource "ibm_iam_authorization_policy" "service-policy" {
  source_service_name = "messagehub"
  target_service_name = "messagehub"
  roles               = ["Reader"]
  description         = "test mirroring setup via terraform"
}

resource "ibm_resource_instance" "es_instance_target" {
  name              = "terraform-integration-target"
  service           = "messagehub"
  plan              = "enterprise-3nodes-2tb"
  location          = "us-south"
  resource_group_id = data.ibm_resource_group.group.id
  parameters_json = jsonencode(
    {
      mirroring = {
        source_crn   = data.ibm_resource_instance.es_instance_source.id
        source_alias = "source-alias"
        target_alias = "target-alias"
        options = {
          topic_name_transform = {
            type = "rename"
            rename = {
              add_prefix    = "add_prefix"
              add_suffix    = "add_suffix"
              remove_prefix = "remove_prefix"
              remove_suffix = "remove_suffix"
            }
          }
          group_id_transform = {
            type = "rename"
            rename = {
              add_prefix    = "add_prefix"
              add_suffix    = "add_suffix"
              remove_prefix = "remove_prefix"
              remove_suffix = "remove_suffix"
            }
          }
        }
      }
    }
  )
  timeouts {
    create = "3h"
    update = "1h"
    delete = "15m"
  }
}
# Configure a service-to-service binding between both instances to allow both instances to communicate.
resource "ibm_iam_authorization_policy" "instance_policy" {
  source_service_name         = "messagehub"
  source_resource_instance_id = ibm_resource_instance.es_instance_target.guid
  target_service_name         = "messagehub"
  target_resource_instance_id = data.ibm_resource_instance.es_instance_source.guid
  roles                       = ["Reader"]
  description                 = "test mirroring setup via terraform"
}

# Select some topics from the source cluster to mirror.
resource "ibm_event_streams_mirroring_config" "es_mirroring_config" {
  resource_instance_id     = ibm_resource_instance.es_instance_target.id
  mirroring_topic_patterns = ["topicA", "topicB"]
}
