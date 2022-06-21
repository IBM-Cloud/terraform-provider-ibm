# Terraform Toolchain Template for Build your own toolchain

This toolchain templates creates an empty toolchain in IBM Cloud region and resource group of your selection. This toolchain has no preconfigured tools. Use this toolchain template as a starting toolchain and keep adding toolchain integrations as per your requirement. 

## Usage

The terraform template requires you to provide values for the terraform variables. 
Copy the file `variables.tfvars.example` as `variables.tfvars`. Provide appropriate values to the variables within the file. 


1. Initialize the terraform project to download the terraform providers and modules
```bash
$ terraform init
```
2. Perform terraform plan with the variables. Run `terraform plan` to see the changes that will be applied to your account after you make any change to the terraform code. 
```bash
$ terraform plan -var-file=./variables.tfvars
```

3. Perform terraform apply with the variables. Run `terraform apply` to apply the changes to the IBM Cloud after that will be applied to your account after you make any change to the terraform code. 

```bash
$ terraform apply -var-file=./variables.tfvars
```

Run `terraform destroy` to clean up and destroy all the resources created for the toolchain.

## Assumptions



## Notes



## Requirements



## Providers



## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key | `string` | true |
| resource_group | Name of the resource group that the toolchain will belong to. | `string` | true |
| region | IBM Cloud Region in which toolchain will be created. Defaults to `us-south` | `string` | false |
| name | Name of tool. | `string` | false |
| toolchain_name | Name of the toolchain | `string` | false |
| toolchain_description | Description of the toolchain | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_id | Toolchain ID of the newly created toolchain resource |