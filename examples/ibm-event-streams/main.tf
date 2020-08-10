
data "ibm_resource_group" "group" {
  name = "Default"
}

#### Scenario 1: Create Event Streams service instance and topic
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "standard" # "lite", "enterprise-3nodes-2tb"
  location          = "us-south" # "us-east", "eu-gb", "eu-de", "jp-tok", "au-syd"
  resource_group_id = data.ibm_resource_group.group.id

  # parameters = {
  #   service-endpoint     = "private"                    # for enterprise instance
  #   private_ip_allowlist = ["1.0.0.0/32", "1.0.0.1/32"] # for enterprise instance
  # }

  # timeouts {
  #   create = "15m" # use 3h when creating enterprise instance
  #   update = "15m" # use 1h when updating enterprise instance
  #   delete = "15m"
  # }
}

resource "ibm_event_streams_topic" "es_topic_1" {
  resource_instance_id = ibm_resource_instance.es_instance_1.id
  name                 = "my-es-topic"
  partitions           = 1
  config = {
    "cleanup.policy"  = "compact,delete"
    "retention.ms"    = "86400000"
    "retention.bytes" = "1073741824"
    "segment.bytes"   = "536870912"
  }
}

#### Scenario 2: Create topic on an existing Event Streams instance
data "ibm_resource_instance" "es_instance_2" {
  name              = "terraform-integration-2"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_topic" "es_topic_2" {
  resource_instance_id = data.ibm_resource_instance.es_instance_2.id
  name                 = "my-es-topic"
  partitions           = 1
  config = {
    "cleanup.policy"  = "compact,delete"
    "retention.ms"    = "86400000"
    "retention.bytes" = "1073741824"
    "segment.bytes"   = "536870912"
  }
}

#### Scenario 3: Create a kafka consumer application connecting to an existing Event Streams instance and its topics
data "ibm_resource_instance" "es_instance_3" {
  name              = "terraform-integration-3"
  resource_group_id = data.ibm_resource_group.group.id
}

data "ibm_event_streams_topic" "es_topic_3" {
  resource_instance_id = data.ibm_resource_instance.es_instance_3.id
  name                 = "my-es-topic"
}

resource "kafka_consumer_app" "es_kafka_app" {
  bootstrap_server = lookup(data.ibm_resource_instance.es_instance_3.extensions, "kafka_brokers_sasl", [])
  topics           = [data.ibm_event_streams_topic.es_topic_3.name]
  apikey           = var.es_reader_api_key
}
