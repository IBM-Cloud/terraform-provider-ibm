# IBM Cloud Hyper Protect Crypto Keys

# Managing HPCS Service Instances using Terraform Resources

This is a collection of resources that make it easier to provision and manage HPCS Instance IBM Cloud Platform:

* Provisioning HPCS Instances - `ibm_resource_instance`
* Initialising HPCS Instance - Initialises and Configured zeroised crypto units `hpcs_init` (null resource)
* Managing Keys on HPCS Instance - [ Key Management Service Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-kp-resources#kms-key)

## HPCS Initialisation Architecture

![HPCS Architecture](references/diagrams/architechture.png?raw=true)
The figure above depicts the basic architecture of the IBM Cloud HPCS Init Terraform Automation.
The main components are..

- **COS Bucket**: HPCS Crypto unit credentials that stored in a Bucket as a json file will be taken as an input by `hpcs-init` terraform module and the secret tke-files that are obtained after execution of template will be stored back as zip file in cos bucket.
- **Terraform**: Reads the terraform configuration files and templates, execute the plan, and communicate with the plugins, manages the resource state and .tfstate file after apply.
- **IBM Cloud TKE Plugin**: The Python script that automates the initialisation process uses IBM CLOUD TKE Plugin

## Terraform versions

Terraform 0.12.

## Usage

Full example is in [main.tf](main.tf)

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

### Provision HPCS Instance

Note: `provision_instance` will determine if the instance has to be provisioned or not. If `provision_instance` is true, count will be 1 and the instance will be provisioned..
```hcl
resource "ibm_resource_instance" "hpcs_instance" {
  count    = (var.provision_instance == true ? 1 : 0)
  name     = var.hpcs_instance_name
  service  = "hs-crypto"
  plan     = var.plan
  location = var.location
  parameters = {
    units = var.units
  }
}
```

### Initialize HPCS Instance

```hcl
resource "null_resource" "hpcs_init" {
  provisioner "local-exec" {
    command = <<EOT
    python ./scripts/init.py
        EOT
    environment = {
      CLOUDTKEFILES = var.tke_files_path
      INPUT_FILE    = file(var.input_file_name)
      HPCS_GUID     = data.ibm_resource_instance.hpcs_instance.guid
    }
  }
}
```

### Manage HPCS Keys
`Note:` To Manage Keys, Instance should be Initialized..

```hcl
resource "ibm_kms_key" "key" {
  instance_id  = data.ibm_resource_instance.hpcs_instance.guid
  key_name     = var.key_name
  standard_key = false
  force_delete = true
}
```

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| provision_instance | Determines if the instance has to be created or not. | `bool` | Yes |
| hpcs_instance_name | Name of HPCS Instance. | `string` | Yes |
| location | Location of HPCS Instance. | `string` | Yes |
| plan | Plan of HPCS Instance.Default: `standard` | `string` | No |
| units | No of crypto units that has to be attached to the instance. | `number` | Yes |
| api_key | Api key of the COS bucket. | `string` | No |
| cos_crn | COS instance CRN. | `string` | No |
| endpoint | COS endpoint. | `string` | No |
| bucket_name | COS bucket name. | `string` | No |
| input_file_name | Input json file name that is present in the cos-bucket or in the local. | `string` | Yes |
| tke_files_path | Path to which tke files has to be exported. | `string` | Yes |
| key\_name | Name of the key. | `string` | Yes |

Note: COS Credententials are required when `download_from_cos` and `upload_to_cos` null resources are used

 Name | Description |
|------|-------------|
| keyID | The ID of the key.|
| InstanceGUID | The GUID of the HPCS Instance.|

## Pre-Requisites for Initialisation:
* Login to IBM Cloud Account using cli `ibmcloud login --apikey= <Your IC Api Key> -a cloud.ibm.com`
* Target Resource group and region `ibmcloud target -g <resource group name>` `ibmcloud target -r <region>`
* Generate oauth-tokens `ibmcloud iam oauth-tokens`. This step should be done as and when token expires. 

## Notes On Initialization:
* The current script adds only one signature key admin.
* The signature key associated with the Admin name given in the json file will be selected as the current signature key.
* If number of master keys added is more than three, Master key registry will be `loaded`, `commited` and `setimmidiate` with last three added master keys.
* Please find the example json [here](references/input.json).
* Input can be fed in two ways either through local or through IBM Cloud Object Storage
* The input file is download from the cos bucket using `download_from_cos` null resource
* Secret TKE Files that are obtained after initialisation can be stored back in the COS Bucket as a Zip File using `upload_to_cos`null resource
* After uploading zip file to COS Bucket all the secret files and input file can be deleted from the local machine using `remove_tke_files` null resource.

## Future Enhancements:
* Automation of Pre-Requisites.
* Capability to add and select one or more admin.
* Integration with Hashicorp vault.