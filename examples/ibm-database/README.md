# This example shows how to create an IBM Cloud Database and configure connectivity from a VSI

This sample configuration will configure an ICD instance and a VSI. Security Groups are configured to allow the VSI to access the ICD instance on the dynamically defined ICD connection port. Whitelisting is configured on the ICD instance to only allow access from the VSI IP address. 

## Costs

This sample uses chargable services and **will** incur costs for the time the services are deployed. Execution of `terraform destroy` will result in deletion of all resources including the ICD service instance. Billing for VSIs and ICD will terminate on the hour. 


## Dependencies

- User has IAM security rights to create and configure an IBM Cloud Database instance and VSIs

## Configuration 

The following variables need to be set in the `terraform.tfvars` file before use

* `softlayer_username` is an Infrastructure user name. Go to https://control.bluemix.net/account/user/profile, scroll down, and check API Username.
* `softlayer_api_key` is an Infrastructure API Key. Go to https://control.bluemix.net/account/user/profile, scroll down, and check Authentication Key.
* `bluemix_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://console.bluemix.net/iam/#/apikeys and create a new key.


The example is deployed in the eu-gb region. The `region` parameter in provider.tf must be set to the same region as the ICD instance will be deployed in as defined by the `location` parameter on the ibm_database resource. 

## Outputs 

The composed connection string for the default admin ID for the deployed ICD PostgreSQL database is displayed. Connection string usage is dependent on the type of ICD service deployed and there may be multiple strings for different hosts in the DB cluster. Consult the ICD documentation about the use of connection strings, e.g. https://cloud.ibm.com/docs/services/databases-for-etcd?topic=databases-for-etcd-connection-strings#connection-strings 


## Running the configuration 
```shell
terraform init
terraform plan
```

For apply phase

```shell
terraform apply
```

```shell
terraform destroy
```  
