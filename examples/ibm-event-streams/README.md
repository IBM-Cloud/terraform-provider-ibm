# IBM Event Streams examples

This example shows several Event Streams usage scenarios.

## Creating Event Streams instances

Event Streams service instances are created with the `"ibm_resource_instance"` resource type.

The following `"ibm_resource_instance"` arguments are required:

- `name`: The service instance name, as it will appear in the Event Streams UI and CLI.

- `service`: Use `"messagehub"` for an Event Streams instance.

- `plan`: One of `"lite"`, `"standard"`, or `"enterprise-3nodes-2tb"`. For more information about the plans, see [Choosing your plan](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-plan_choose). Note: `"enterprise-3nodes-2tb"` selects the Enterprise plan.

- `location`: The region where the service instance will be provisioned. For a list of regions, see [Region and data center locations](https://cloud.ibm.com/docs/overview?topic=overview-locations).

- `resource_group_id`: The ID of the resource group in which the instance will be provisioned. For more information about resource groups, see [Managing resource groups](https://cloud.ibm.com/docs/account?topic=account-rgs).

The `parameters/parameters_json` argument is optional and provides additional provision or update options. Supported parameters are:

- `throughput`: One of `"150"` (the default), `"300"`, `"450"`. The maximum capacity in MB/s for producing or consuming messages. For more information see [Scaling Enterprise plan capacity](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_scaling_capacity). *Note:* See [Scaling combinations](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_scaling_capacity#ES_scaling_combinations) for allowed combinations of `throughput` and `storage_size`.
    - Example:  `throughput  =  "300"`

- `storage_size`: One of `"2048"` (the default), `"4096"`, `"6144"`, `"8192"`, `"10240"`, or `"12288"`.  The amount of storage capacity in GB. For more information see [Scaling Enterprise plan capacity](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_scaling_capacity). *Note:* See [Scaling combinations](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_scaling_capacity#ES_scaling_combinations) for allowed combinations of `throughput` and `storage_size`.
    - Example:  `storage_size  =  "4096"`

- `service-endpoints`: One of `"public"` (the default), `"private"`, or `"public-and-private"`. For enterprise instance only. For more information see [Restricting network access](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-restrict_access).
    - Example:  `service-endpoints  =  "private"`

- `private_ip_allowlist`: **Deprecated** An array of CIDRs specifying a private IP allowlist. For enterprise instance only. For more information see [Specifying an IP allowlist](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-restrict_access#specify_allowlist). This feature has been deprecated in favor of context-based restrictions.
    - Example:  `private_ip_allowlist  =  "[10.0.0.0/32,10.0.0.1/32]"`

- `metrics`: An array of strings, allowed values are `"topic"`, `"partition"`, and `"consumers"`. Enables additional enhanced metrics for the instance. For enterprise instance only. For more information on enhanced metrics, see [Enabling enhanced Event Streams metrics](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-metrics#opt_in_enhanced_metrics).
    - Example:  `metrics  =  "[topic,partition]"`

- `kms_key_crn`: The CRN (as a string) of a customer-managed root key provisioned with an IBM Cloud Key Protect or Hyper Protect Crypto Service. If provided, this key is used to encrypt all data at rest. For more information on customer-managed encryption, see [Managing encryption in Event Streams](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-managing_encryption).
    - Example:  `kms_key_crn  =  "crn:v1:prod:public:kms:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:20adf7eb-e095-4dec-08cf-0b7d81e32db6:key:3fa9d921-d3b6-3516-a1ec-d54e27e7638b"`

- `mirroring`: To enable mirroring in the cluster using `parameters_json`. For enterprise instance only. If defined `source_crn` (source cluster CRN as a string), `source_alias` (alias for source cluster as a string), `target_alias` (alias for target cluster as a string) are required. `options` are optional. 
    - Example: 
  
```terraform
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
```

The `timeouts` argument is used to specify how long the IBM Cloud terraform provider will wait for the provision, update, or deprovision of the service instance. Values of 15 minutes are sufficient for standard and lite plans. For enterprise plans:
- Use "3h" for create. Add an additional 1 hour for each level of non-default throughput, and an additional 30 minutes for each level of non-default storage size. For example with `throughput = "300"` (one level over default) and `storage_size = "8192"` (three levels over default), use 3 hours + 1 * 1 hour + 3 * 30 minutes = 5.5 hours.
- Use "1h" for update. If increasing the throughput or storage size, add an additional 1 hour for each level of non-default throughput, and an additional 30 minutes for each level of non-default storage size.
- Use "1h" for delete.

## Scenarios

#### Scenario 1: Create an Event Streams standard-plan service instance.

This creates a standard plan instance in us-south.

```terraform
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
```

#### Scenario 2: Create an Event Streams enterprise service instance with non-default attributes

This creates an enterprise plan instance in us-east with 300 MB/s throughput, 4 TB storage, private endpoints with an allowlist, and enhanced metrics for topics and consumer groups. The timeouts are calculated as described above.

```terraform
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
```

#### Scenario 3: Create a topic on an existing Event Streams instance.

For more information on topics and topic parameters, see [Topics and partitions](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-apache_kafka&interface=ui#kafka_topics_partitions) and [Using the administration Kafka Java client API](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-kafka_java_api).

```terraform
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
```

#### Scenario 4: Create a schema on an existing Event Streams Enterprise instance

For more information on the Event Streams schema registry, see [Using Event Streams Schema Registry](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-ES_schema_registry).

```terraform
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
```

#### Scenario 5: Apply access tags to an Event Streams service instance

Tags are applied using the `"ibm_resource_tag"` terraform resource.
For more information about tagging, see the documentation for the `"ibm_resource_tag"` resource and [Tagging](https://cloud.ibm.com/apidocs/tagging).

```terraform
data "ibm_resource_instance" "es_instance_5" {
  name              = "terraform-integration-5"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_resource_tag" "tag_example_on_es" {
  tags        = ["example:tag"]
  tag_type    = "access"
  resource_id = data.ibm_resource_instance.es_instance_5.id
}
```

#### Scenario 6: Set default and user quotas on an existing Event Streams instance.

This code sets the default quota to 32768 bytes/second for producers and 16384 bytes/second for consumers.
It sets a quota for user `iam-ServiceId-00001111-2222-3333-4444-555566667777` to 65536 bytes/second for producers and no limit (-1) for consumers.
For more information on quotas, see [Setting Kafka quotas](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-enabling_kafka_quotas).

```terraform
data "ibm_resource_instance" "es_instance_6" {
  name              = "terraform-integration-6"
  resource_group_id = data.ibm_resource_group.group.id
}

resource "ibm_event_streams_quota" "default_quota" {
  resource_instance_id = data.ibm_resource_instance.es_instance_6.id
  entity               = "default"
  producer_byte_rate   = 32768
  consumer_byte_rate   = 16384
}

resource "ibm_event_streams_quota" "user00001111_quota" {
  resource_instance_id = data.ibm_resource_instance.es_instance_6.id
  entity               = "iam-ServiceId-00001111-2222-3333-4444-555566667777"
  producer_byte_rate   = 65536
  consumer_byte_rate   = -1
}
```

#### Scenario 7: Create a target Event Streams service instance with mirroring enabled and its mirroring config
data "ibm_resource_instance" "es_instance_source" {
  name              = "terraform-integration-source"
  resource_group_id = data.ibm_resource_group.group.id
}
# setup s2s at service level for mirroring to work
resource "ibm_iam_authorization_policy" "service-policy" {
  source_service_name         = "messagehub"
  target_service_name         = "messagehub"
  roles                       = ["Reader"]
  description                 = "test mirroring setup via terraform"
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

#### Scenario 8: Connect to an existing Event Streams instance and its topics.

This scenario uses a fictitious `"kafka_consumer_app"` resource to demonstrate how a consumer application could be configured.
The resource uses three configuration properties:

1. The Kafka broker hostnames used to connect to the service instance.
2. An API key for reading from the topics.
3. The names of the topic(s) which the consumer should read.

The broker hostnames would be required by any consumer or producer application. After the Event Streams service instance has been created, they are available in the `extensions` attribute of the service instance, as an array named `"kafka_brokers_sasl"`. This is shown in the example.

An API key would also be required by any application. This key would typically be created with reduced permissions to restrict the operations it can perform, for example only allowing it to read from certain topics. See [Managing authentication to your Event Streams instance](https://cloud.ibm.com/docs/EventStreams?topic=EventStreams-security) for more information on creating keys. The example assumes the key is provided as a terraform variable.

The topic names can be provided as strings, or can be taken from topic data sources as shown in the example.

```terraform
# Use an existing instance
data "ibm_resource_instance" "es_instance_7" {
  name              = "terraform-integration-7"
  resource_group_id = data.ibm_resource_group.group.id
}

# Use an existing topic on that instance
data "ibm_event_streams_topic" "es_topic_7" {
  resource_instance_id = data.ibm_resource_instance.es_instance_7.id
  name                 = "my-es-topic"
}

# The FICTITIOUS consumer application, configured with brokers, API key, and topics
resource "kafka_consumer_app" "es_kafka_app" {
  bootstrap_server = lookup(data.ibm_resource_instance.es_instance_7.extensions, "kafka_brokers_sasl", [])
  apikey           = var.es_reader_api_key
  topics           = [data.ibm_event_streams_topic.es_topic_7.name]
}
```
#### Scenario 8: Create a target Event Streams service instance with mirroring enabled and its mirroring config
```terraform
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
```

## Dependencies

- The owner of the `ibmcloud_api_key` has permission to create Event Streams instance under specified resource group and has Manager role to the created instance in order to create topic.

- The schema registry is available only for Event Streams Enterprise plan service instances.

## Configuration

- `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/apikeys and create a new key.

## Running the configuration

For planning phase

```bash
terraform init
terraform plan
```

For apply phase

```bash
terraform apply
```

For destroy

```bash
terraform destroy
```
