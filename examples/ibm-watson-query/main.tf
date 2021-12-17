
data "ibm_resource_group" "group" {
  name = "Default"
}

#### Scenario 1: Create Watson Query service instance
resource "ibm_resource_instance" "wq_instance_1" {
  name              = "terraform-integration-1"
  service           = "data-virtualization"
  plan              = "data-virtualization-enterprise" # "data-virtualization-enterprise-dev","data-virtualization-enterprise-preprod","data-virtualization-enterprise-dev-stable"
  location          = "us-south" # "eu-gb", "eu-de", "jp-tok"
  resource_group_id = data.ibm_resource_group.group.id

  # timeouts {
  #   create = "15m" # use 3h when creating enterprise instance, add additional 1h for each level of non-default throughput, add additional 30m for each level of non-default storage_size
  #   update = "15m" # use 1h when updating enterprise instance, add additional 1h for each level of non-default throughput, add additional 30m for each level of non-default storage_size
  #   delete = "15m"
  # }
}

