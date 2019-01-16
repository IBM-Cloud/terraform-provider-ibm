---
layout: "ibm"
page_title: "IBM : ssl_certificate"
sidebar_current: "docs-ibm-resource-ssl-certificate"
description: |-
  Manages IBM Compute SSL Certificate.
---

# ibm\_ssl_certificate

Provides an SSL certificate resource. This allows SSL certificates to be requested, and delete request for ssl certificates.

For additional details, see the [IBM Cloud (SoftLayer) security certificates Request docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate/Request).

## Example Usage

In the following example, you can use a certificate on file:

```hcl
resource "ibm_ssl_certificate" "my_ssllllll" {
  	certificateSigningRequest= "-----BEGIN CERTIFICATE REQUEST-----\nMIIC2jCCAcICAQAwgYAxCzAJBgNVBAYTAklOMRMwEQYDVQQIDApNYWhhcmFzaHRh\nMQ0wCwYDVQQHDARQdW5lMRAwDgYDVQQKDAdJQk1QdW5lMQwwCgYDVQQLDANJQk0x\nFDASBgNVBAMMC2libS5wdW5lLmluMRcwFQYJKoZIhvcNAQkBFghyYkBncy5pbjCC\nASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALZTfLDkU/q8ory+a+aBIAB7\nezyMbF4iqEpO3mVmlIfENBRq/MomhTiYID2jVWhPS4KytXa9bANgAG+NN5zMxqgH\nwUVa4fO2Th//o9RdRPZQcC3IGR+b3tpGSV+KxjsW9XV82lQp6l8HxfSBo12PZcwU\nlu53BXMiVKU6ZHSjz6eDF4MGTuPM1QWs/oexb9y6Yqqj2IicrkviJMuKhh0FN6H7\nXT2MHIF8OHT1xhs8w7U1SdUwhBa5dHvj7vqi33DU0p3s8JdXlzE9PUXKuAzXav1D\nngwE1iT5KFHkEFUr7plH5wGsy/y4tnu2xVzfGMyECS9EYmUtOA5M5RpqF5DdXGEC\nAwEAAaAUMBIGCSqGSIb3DQEJAjEFDANnc2wwDQYJKoZIhvcNAQELBQADggEBAGB7\n3lv/6fSn9rgTiHszLL9pOU9ytjOVrhNjFjDzQL73VQ0+Isb7aPHnWrLz4kT9m/60\nmgy/dHsOIF8KRP1LpOs5BYwlstD3Ss57XR8GatnrLMN4lZCjacL6A8RPhwr3x29W\nMyFntvu2caAL4ZpZpWMKHtoemXijCiFXa9Z4pZFBk4V7k0/DIEXEeyYazsSaXeTw\nXr4IFPmk7VS/NkLAht2hRhllN5NHGf/gzTsmgrgKclXtf1Z7EotnDTTIt0dFVtk1\nVX2Z7kvx9/QWbDVhPEz2uOrJnCoAm+0OpQfFc4THcP0uv0Y49B3WUG89mAjlWQKa\nU7hhc8gZ77+eaBQKD6k=\n-----END CERTIFICATE REQUEST-----"
	organizationInformation = {
		org_address = {
			org_addressLine1= "abc"
			org_addressLine2= "cms"
			org_city="xyz"
			org_countryCode= "AF"
			org_state="OT"
			org_postalCode= "411119"
		}
		org_organizationName= "IBMPune"
		org_phoneNumber= "9876543210"
	}	
	technicalContactSameAsOrgAddressFlag = "true"
	technicalContact = {
		tech_organizationName= "IBMPune"
		tech_phoneNumber= "952741928"
		tech_firstName = "poy"
		tech_lastName = "joy"
		tech_emailAddress = "email.prefix@yahoo.com"
		tech_title= "something"
	}
	administrativeContact = {
		admin_organizationName= "IBMPune"
		admin_phoneNumber= "952741928"
		admin_firstName = "poy"
		admin_lastName = "joy"
		admin_emailAddress = "email.prefix@yahoo.com"
		admin_title= "something"
	}
	billingContact = {
		billing_organizationName= "IBMPune"
		billing_phoneNumber= "952741928"
		billing_firstName = "poy"
		billing_lastName = "joy"
		billing_emailAddress = "email.prefix@yahoo.com"
		billing_title= "something"
	}
	administrativeContactSameAsTechnicalFlag = "false"
	billingContactSameAsTechnicalFlag = "false"
	administrativeAddressSameAsOrganizationFlag ="true"
	billingAddressSameAsOrganizationFlag = "true"	
	sslType="SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
	renewalFlag= true
	serverCount= 1
	serverType= "apache2"
	validityMonths= 24
	orderApproverEmailAddress= "admin@pune.in"	
}
```


## Argument Reference

The following arguments are supported:

* `certificateSigningRequest` - (Required, string) The Certificate Signing Request which is specially formatted encrypted message sent from a Secure Sockets Layer (SSL) digital certificate applicant to a certificate authority.
* `sslType` - (Required, string) The ssl certificate type.
* `serverType` - (Required, string) The server type for which we are requesting ssl certficate.
* `serverCount` - (Required, string) The number of servers with provided server tye .
* `validityMonths` - (Required, integer) The validity of ssl certificate in months it should be multiple of 12.
* `orderApproverEmailAddress` - (Required, string) The email of approver to approve ssl certificate request.
* `org_addressLine1` - (Required, string) The address of organization who is requesting for ssl certificate.
* `org_addressLine2` - (optional, string) The address of organization who is requesting for ssl certificate.
* `org_city` - (Required, string) The city of organization which is requesting for ssl certificate .
* `org_postalCode` - (Required, integer) The postal code for the city of organization.
* `org_state` - (Required, string) The two letter state code of organization who is requesting for ssl certificate. Allowed value for country which doesn't have states is `OT`.
* `org_countryCode` - (Required, string) The two letter country code of organization.
* `org_organizationName` - (Required, string) Name of organization.
* `tech_addressLine1` - (Required, string) The address for technical contact.
* `tech_addressLine2` - (optional, string) The address for technical contact.
* `tech_city` - (Required, string) The city for technical contact.
* `tech_postalCode` - (Required, integer) The postal code for technical contact.
* `tech_state` - (Required, string) The two letter state code of technical contact. Allowed value for country which doesn't have states is `OT`.
* `tech_countryCode` - (Required, string) The two letter country code for technical contact.
* `tech_organizationName` - (Required, string) Name of organization for technical contact.
* `tech_firstName` - (Required, string) The first name for technical contact.
* `tech_lastName` - (Required, string) The last name for technical contact.
* `tech_title` - (Required, string) The title for for technical contact.
* `tech_emailAddress` -(Required, string) email address for technical contact.
* `admin_addressLine1` - (Optional, string) The address for administrative contact.
* `admin_addressLine2` - (optional, string) The address for administrative contact.
* `admin_city` - (Optional, string) The city for administrative contact.
* `admin_postalCode` - (Optional, integer) The postal code for administrative contact.
* `admin_state` - (Optional, string) The two letter state code of administrative contact. Allowed value for country which doesn't have states is `OT`.
* `admin_countryCode` - (Optional, string) The two letter country code for administrative contact.
* `admin_organizationName` - (Optional, string) Name of organization for administrative contact.
* `admin_firstName` - (Optional, string) The first name for administrative contact.
* `admin_lastName` - (Optional, string) The last name for administrative contact.
* `admin_title` - (Optional, string) The title for for administrative contact.
* `admin_emailAddress` -(Optional, string) email address for administrative contact.
* `billing_addressLine1` - (Optional, string) The address for billing contact.
* `billing_addressLine2` - (optional, string) The address for billing contact.
* `billing_city` - (Optional, string) The city for billing contact.
* `billing_postalCode` - (Optional, integer) The postal code for billing contact.
* `billing_state` - (Optional, string) The two letter state code of billing contact. Allowed value for country which doesn't have states is `OT`.
* `billing_countryCode` - (Optional, string) The two letter country code for billing contact.
* `billing_organizationName` - (Optional, string) Name of organization for billing contact.
* `billing_firstName` - (Optional, string) The first name for billing contact.
* `billing_lastName` - (Optional, string) The last name for billing contact.
* `billing_title` - (Optional, string) The title for for billing contact.
* `billing_emailAddress` -(Optional, string) email address for billing contact.
* `technicalContactSameAsOrgAddressFlag` -(Optional, bool) If your organization address and technical contact address is the same make this flag as true and skip technical contact address details.
* `administrativeContactSameAsTechnicalFlag` -(Required, bool)- If your technical contact details and administrative contact details is the same then make this as true and skip details of administrative contact.
* `billingContactSameAsTechnicalFlag` -(Required, bool)- If your technical contact details and billing contact details is the same then make this as true and skip details of billing contact. 
* `administrativeAddressSameAsOrganizationFlag` -(Required,bool)If administrative address is same as organization address then make this flag as true and skip address details.
* `billingAddressSameAsOrganizationFlag` -(Required,bool)If billing address is same as organization address then make this flag as true and skip address details. 