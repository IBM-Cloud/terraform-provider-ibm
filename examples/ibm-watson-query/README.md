# IBM Watson Query examples

This example shows 1 usage scenario.

#### Scenario 1: Create a Watson Query service instance.

```terraform
resource "ibm_resource_instance" "es_instance_1" {
  name              = "terraform-integration-1"
  service           = "messagehub"
  plan              = "standard" # "lite", "enterprise-3nodes-2tb"
  location          = "us-south" # "us-east", "eu-gb", "eu-de", "jp-tok", "au-syd"
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
