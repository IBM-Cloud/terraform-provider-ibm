# This example shows how to create and IBM Cloud Internet Services instance, pools and global load balancer

This sample configuration will configure CIS instance, a health check monitor, origin pool and global load balancer. Also see the example  `ibm-website-multi-region` for an example of using CIS with a working website deployed across multiple regions. 

## Costs

This sample uses chargable services and **will** incur costs. Billing for the CIS service instance is pro-rata'd for the remaining duration of the month it is deployed in. Execution of `terraform destroy` will result in deletion of all resources including the CIS service instance. Billing for VSIs and Cloud Load Balancing will terminate on the hour. The billing for the CIS service instance will be pro-rata'd to the end of the month. For each delete and recreate of the environment a new CIS service instance will be created and result in an additional billing instance pro-rata'd for the month. 

To avoid additional CIS service instance costs if the sample confifuration is executed additional times, after creation the `ibm_cis` resource should be removed from the configuration and replaced with a `ibm_cis` data source. All dependent CIS Terraform resource definitions must also be updated to reference the `data source`. A replacement data-source sample configuration is included in the `dns.tf` file. 

The following steps can be used to replace the `ibm_cis` resource with a data source to preserve the CIS service instance.

  1. Create the initial environment
    `terraform apply`

  2. Delete the ibm_cis resource from the Terraform state.
    `terraform state rm ibm_cis.web_domain`
 

  3. Update `dns.tf` to replace the `ibm_cis` resource with a data source. 
    - Comment out lines 17 - 89
    - Uncomment lines 93 - 163
    - Save file

  4. Refresh the Terraform state file with the new data source.
    `terraform plan`

  5. Delete the existing web site deployment
    `terraform destroy`


## Dependencies

- User has IAM security rights to create and configure an Internet Services instance
- DNS Domain registration

## Configuration 

The following variables need to be set in the `terraform.tf` file before use

* `softlayer_username` is an Infrastructure user name. Go to https://control.bluemix.net/account/user/profile, scroll down, and check API Username.
* `softlayer_api_key` is an Infrastructure API Key. Go to https://control.bluemix.net/account/user/profile, scroll down, and check Authentication Key.
* `bluemix_api_key` - An API key for IBM Cloud services. If you don't have one already, go to https://console.bluemix.net/iam/#/apikeys and create a new key.


Customise the variables in `variables.tf` to your local environment and chosen DNS domain name. 

* `domain` in the DNS Domain for web server registed with the DNS registrar. The DNS domain must be pre-registered with the IBM Cloud [Domain Registration Service](https://cloud.ibm.com/classic/services/domains). 
* `dns_name`. DNS name (prefix) for website, including '.',e.g. 'www.' Can be "" for website to be at root of domain. 
* `datacenter1`. Name of origin pool in region 1. 
* `datacenter2`. Name of origin pool in region 2. 
* `resource_group`. Name of the Resource Group configured resources will be allocated to. Default is "Default". 

  

## Running the configuration 
```shell
terraform init
terraform plan
```

For apply phase

```shell
terraform apply
```

For destroy see notes under **Costs** for how to preserve the CIS service instance to avoid additional billing costs for further instances. Otherwise destroy all resources. 

```shell
terraform destroy
```  
