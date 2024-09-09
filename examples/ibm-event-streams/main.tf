
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
     throughput            = "300"
     storage_size          = "4096"
     service-endpoints     = "private"
     private_ip_allowlist  = "[10.0.0.0/32,10.0.0.1/32]"
     metrics               = "[topic,consumers]"
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
  schema_id = "tf_schema"
  schema = <<SCHEMA
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
