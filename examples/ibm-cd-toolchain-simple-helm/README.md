# Terraform Toolchain Template for Develop a Kubernetes app with Helm Toolchain Template

This toolchain templates creates a toolchain with which you can develop a Docker application and its Helm chart together in source control and have it built and deployed automatically to a Kubernetes cluster. The toolchain performs sanity checks prior to building or deploying and ensures privacy by using a private container registry and namespaces for the container registry and the Kubernetes cluster. This toolchain is also leveraging Code Risk Analyzer and Vulnerability Advisor, to ensure only secure images get deployed.

By default, the toolchain uses a sample Node.js "Hello World" app, but you can link to your own Git repository instead as long as it has a Dockerfile and a Helm chart.

You can manage your IBM Cloud Container clusters in the console.

This toolchain uses tools that are part of the Continuous Delivery service. For more information and terms, see [here](https://www.ibm.com/cloud/architecture/tutorials/use-develop-kubernetes-app-helm-toolchain-with-tekton-pipelines)  

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

1. You have an existing Key Protect Instance created in your account. If not follow the process [here](https://cloud.ibm.com/catalog/services/key-protect) to create one. 
2. You have your IBM Cloud API Key stored in the Key Protect Instance as `Import your own key` option with the Key Name as `ibmcloud-api-key`. This is the same key name which the terraform template requires to pass to the CI and PR Pipelines as environment properties.
3. You have an existing Kubernetes Cluster created using IBM Cloud Kubernetes Service. If not follow the process [here](https://cloud.ibm.com/kubernetes/catalog/create). Once created you will need the name of the Kubernetes Cluster to be passed on to the terraform template.
4. You have an existing namespace within the IBM Cloud Container Registry. If not follow the process [here](https://cloud.ibm.com/registry/namespaces). Once created you will need the name of the Container Registry namespace to be passed on to the terraform template.

## Notes



## Requirements



## Providers



## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| ibmcloud\_api\_key | IBM Cloud API key required by the terraform provider to interact with IBM Cloud | `string` | true |
| resource_group | Name of the resource group that the toolchain will belong to. | `string` | true |
| region | IBM Cloud Region in which toolchain will be created. Defaults to `us-south` | `string` | true |
| toolchain_name | Name of the toolchain | `string` | false |
| toolchain_description | Description of the toolchain | `string` | false |
| ibmcloud\_api | IBM Cloud API Endpoint that will be used by the CI/PR Pipelies to interact with IBM Cloud Services | `string` | false |
| app_name | Name of the application that will be deployed by the CI Pipeline | `string` | false |
| app_image_name | Name of the application container image that will be create by the CI Pipeline | `string` | false |
| cluster_name | Name of the IKS Cluster where the application container image will be deployed | `string` | true |
| cluster_namespace | Name of the IKS Cluster Namespace where the application container image will be deployed | `string` | true |
| cluster_region | Region of the IKS Cluster  | `string` | true |
| registry_namespace | Name of the Container Registry Namespace where the application container image will be stored | `string` | true |
| registry_region | Region of the Container Registry where the application container image will be stored | `string` | true |
| kp_name | Name of the Key Protect Instance where the IBM Cloud API Key to be used within the CI/PR Pipeline is stored | `string` | true |
| app_repo | IBM Cloud Hosted GIT Repository where the application source code resides. Default behavior of the template is to clone this repository from this [source application repository](https://us-south.git.cloud.ibm.com/open-toolchain/hello-helm.git) | `string` | false |
| pipeline_repo | IBM Cloud Hosted GIT Repository where the tekton pipeline  code resides. Default behavior of the template is to clone this repository from this [source pipeline repository](https://us-south.git.cloud.ibm.com/open-toolchain/simple-helm-toolchain.git) | `string` | false |
| tekton_tasks_catalog_repo | IBM Cloud Hosted GIT Repository where the common tekton tasks resides. Default behavior of the template is to clone this repository from this [common tekton tasks repository](https://us-south.git.cloud.ibm.com/open-toolchain/tekton-catalog.git) | `string` | false |

## Outputs

| Name | Description |
|------|-------------|
| toolchain_id | Toolchain ID of the newly created toolchain resource |