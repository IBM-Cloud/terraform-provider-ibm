# IBM Watson Query examples

This example shows 1 usage scenario.

#### Scenario 1: Create a Watson Query service instance.

```terraform
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

```

## Dependencies

- The owner of the `ibmcloud_api_key` has permission to create Watson Query instance under specified resource group.

## Configuration

- `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.

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
