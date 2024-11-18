# This example shows how to create an IBM Cloud Db2 SaaS and configure connectivity from a VSI

This sample configuration to provision an IBM Db2 SaaS instance. 

## Costs

This sample uses chargable services and **will** incur costs for the time the services are deployed. Execution of `terraform destroy` will result in deletion of all resources including the Db2 SaaS service instance. Billing for Db2 SaaS will terminate on the hour. 


## Dependencies

- User has IAM security rights to create and configure an IBM Cloud Db2 SaaS Instance.

## Configuration 

The terraform template requires you to provide values for the terraform variables. 
Copy the file `variables.tfvars.example` as `variables.tfvars`. Provide appropriate values to the variables within the file. 

The following variables need to be set in the `terraform.tfvars` file before use:

* `ibmcloud_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://cloud.ibm.com/iam/#/apikeys and create a new key.
* `region` - IBM Cloud region where your Db2 SaaS will be created.
* `resource_group` - Resource group within which Db2 SaaS will be created.


The example is deployed in the us-south region. The `region` parameter in main.tf must be set to the same region as the Db2 SaaS instance will be deployed in as defined by the `location` parameter on the ibm_db2 resource. 

## Outputs 

The composed connection string of Db2 SaaS Instance CRN.


## Running the configuration 
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
