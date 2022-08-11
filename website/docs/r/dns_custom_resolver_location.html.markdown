---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_custom_resolver_location"
description: |-
  Manages IBM Private DNS custom resolver locations.
---

# ibm_dns_custom_resolver_location

~> **Deprecated:**

Beginning with version 1.42.0, using the `ibm_dns_custom_resolver_location` resource to **create** or **update** Custom Resolver Location in Terraform is deprecated. Use the composite **Custom Resolver** resource, which can handle locations, instead. 
It is recommended that you do not use `ibm_dns_custom_resolver_location`. Using the deprecated resource can cause an outage. If you have used the `ibm_dns_custom_resolver_location` resource, change it to the composite Custom Resolver resource before running terraform apply. 
For more information, see the [example usage](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs/resources/dns_custom_resolver#example-usage).

If resources were created using Custom Resolver locations resource block, then follow this migration process to avoid unexpected behaviors in your infrastructure. 

1. Take a backup of your terraform state file. 
2. Modify the **terraform.tfstate** file and remove the json blocks of type: "ibm_dns_custom_resolver_location" and "ibm_dns_custom_resolver", because you are going to perform a terraform import of Custom resolver and its locations from Cloud.   
3. Also, modify your .tf file. Remove custom resolver location resources and add locations to the composite custom resolver block as shown in this example:

```
resource "ibm_dns_custom_resolver" "test" {
        name        = "test-customresolver"
        instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
        description = "new test CR"
        enabled     = true
        locations {
            subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet1.crn
            enabled     = true
        }
        locations {
            subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet2.crn
            enabled     = true
        }
    }
```

4. Run the following terraform import command to import ibm_dns_custom_resolver into your state file:

```
terraform import ibm_dns_custom_resolver.test <custom resolver id>:<dns instance id>

Example:   
terraform import ibm_dns_custom_resolver.test 9930d75f-a501-4305-80fe-ead2e38021da:e180df3b-1de8-442d-8014-881822fd27c5
```
Ensure that the resource custom resolver is imported successfully.    


## ibm_dns_custom_resolver_location

Provides a private DNS custom resolver locations resource. This allows DNS custom resolver location to create. For more information, about custom resolver locations, see [add-custom-resolver-location](https://cloud.ibm.com/apidocs/dns-svcs#add-custom-resolver-location).


## Example usage

```terraform

  	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "test-pdns-custom-resolver-vpc"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "test-pdns-cr-subnet1"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.0.0/24"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
		name			= "test-pdns-cr-subnet2"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "us-south-1"
		ipv4_cidr_block	= "10.240.64.0/24"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "test-customresolver"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "new test CR - TF"
		high_availability = false
		enabled 	= true
	}
	resource "ibm_dns_custom_resolver_location" "test1" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet1.crn
		enabled     = true
		cr_enabled	= true
	}
	resource "ibm_dns_custom_resolver_location" "test2" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		subnet_crn  = ibm_is_subnet.test-pdns-cr-subnet2.crn
		enabled     = true
		cr_enabled	= true
	}
```

## Argument reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, String) The GUID of the private DNS service instance.
* `resolver_id` - (Required, String) The unique identifier of a custom resolver.
* `subnet_crn` - (Required, String) The subnet CRN of the VPC.
* `enabled` - (Optional, Bool) The custom resolver location will enabled or disable. Default is 'false'
* `cr_enabled` - (Optional, Bool) Indicates whether to enable or disable the customer resolver. Default is 'true'



## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your resource is created.

* `dns_server_ip` - (Computed, String) Custom resolver location server ip.
* `healthy` - (Computed, Bool) The Custom resolver location will enable.
* `id` - (String) The unique identifier of the IBM DNS custom resolver location.
* `location_id` - (Computed, String) Type of the custom resolver loaction ID.

