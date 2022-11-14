# Example for CdTektonPipelineV2

This example illustrates how to use the CdTektonPipelineV2

These types of resources are supported:

* cd_tekton_pipeline
* cd_tekton_pipeline_definition
* cd_tekton_pipeline_trigger_property
* cd_tekton_pipeline_property
* cd_tekton_pipeline_trigger

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## CdTektonPipelineV2 resources

cd_tekton_pipeline resource:

```terraform
resource "cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  enable_notifications = false
  enable_partial_cloning = false
  worker = var.cd_tekton_pipeline_worker
}
```

cd_tekton_pipeline_definition resource:

```terraform
resource "cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  source {
    type = "git"
    properties {
      url = ibm_cd_toolchain_tool_hostedgit.tekton_repo.repo_url
      branch = "master"
      path = ".tekton"
    }
  }
}
```

cd_tekton_pipeline_property resource:

```terraform
resource "cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  name = "env-prop-1"
  value = "Environment text property 1"
  type = "text"
}
```

cd_tekton_pipeline_trigger resource:

```terraform
resource "cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  type = "manual"
  name = "trigger1"
  event_listener = "listener"
  tags = [ "tag1", "tag2" ]
  worker {
    id = "public"
  }
  max_concurrent_runs = 1
}
```

cd_tekton_pipeline_trigger_property resource:

```terraform
resource "cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = ibm_cd_tekton_pipeline.cd_pipeline_instance.pipeline_id
  trigger_id = ibm_cd_tekton_pipeline_trigger.cd_tekton_pipeline_trigger_instance.trigger_id
  name = "trig-prop-1"
  value = "trigger 1 text property"
  type = "text"
}
```

## CdTektonPipelineV2 Data sources

cd_tekton_pipeline data source:

```terraform
data "cd_tekton_pipeline" "cd_tekton_pipeline_instance" {
  id = var.cd_tekton_pipeline_id
}
```

cd_tekton_pipeline_definition data source:

```terraform
data "cd_tekton_pipeline_definition" "cd_tekton_pipeline_definition_instance" {
  pipeline_id = var.cd_tekton_pipeline_definition_pipeline_id
  definition_id = var.cd_tekton_pipeline_definition_definition_id
}
```

cd_tekton_pipeline_property data source:

```terraform
data "cd_tekton_pipeline_property" "cd_tekton_pipeline_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_property_pipeline_id
  property_name = var.cd_tekton_pipeline_property_property_name
}
```

cd_tekton_pipeline_trigger data source:

```terraform
data "cd_tekton_pipeline_trigger" "cd_tekton_pipeline_trigger_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_trigger_id
}
```

cd_tekton_pipeline_trigger_property data source:

```terraform
data "cd_tekton_pipeline_trigger_property" "cd_tekton_pipeline_trigger_property_instance" {
  pipeline_id = var.cd_tekton_pipeline_trigger_property_pipeline_id
  trigger_id = var.cd_tekton_pipeline_trigger_property_trigger_id
  property_name = var.cd_tekton_pipeline_trigger_property_property_name
}
```

## Assumptions

1. Creating a Tekton Pipeline requires creating an IBM CD Toolchain (ibm_cd_toolchain) resource and a Tekton Pipeline tool resource instance (ibm_cd_toolchain_tool_pipeline) in that toolchain. These are included in the example in `main.tf`.
2. The Tekton Pipeline resource also requires a repository to be included in the toolchain, which contains the Tekton Pipeline declaration. An example is included in the `main.tf`.

## Notes

1. Copy the `variables.tfvars.example` file to your own `terraform.tfvars` file and add values for the variables within.
2. Use `terraform apply -var-file="terraform.tfvars"` to generate the toolchain and pipeline using your provided variables.

## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm | >=1.48.0 |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| resource\_group | Resource group within which toolchain will be created | `string` | true |
| region | IBM Cloud region where your toolchain will be created | `string` | true |
| toolchain\_name | Name of the Toolchain | `string` | true |
| toolchain\_description | Description for the Toolchain | `string` | true |
| clone\_repo | URL of the tekton repo to clone, e.g. `https://github.com/open-toolchain/hello-tekton` | `string` | true |
| repo\_name | Name of the new repo that will be created in the toolchain | `string` | true |
| cluster | Name of the cluster where the app will be deployed | `string` | true |
| cluster_namespace | Namespace in the cluster where the app will be deployed. Default = `prod` | `string` | false |
| registry_namespace | IBM Cloud Container Registry namespace where the app image will be built and stored | `string` | true |

## Outputs

| Name | Description |
|------|-------------|
| cd_tekton_pipeline | cd_tekton_pipeline object |
| cd_tekton_pipeline_definition | cd_tekton_pipeline_definition object |
| cd_tekton_pipeline_property | cd_tekton_pipeline_property object |
| cd_tekton_pipeline_trigger | cd_tekton_pipeline_trigger object |
| cd_tekton_pipeline_trigger_property | cd_tekton_pipeline_trigger_property object |
