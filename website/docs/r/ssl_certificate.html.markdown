---
layout: "ibm"
page_title: "IBM : ssl_certificate"
sidebar_current: "docs-ibm-resource-ssl-certificate"
description: |-
  Order New IBM SSL Certificate.
---

# ibm\_ssl_certificate

Provides an SSL certificate resource. This allows SSL certificates to be requested, and delete request for ssl certificates.

For additional details, see the [IBM Cloud (SoftLayer) security certificates Request docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate/Request).

## Example Usage

In the following example, you can use a certificate on file:

```hcl
resource "ibm_ssl_certificate" "my_ssllllll" {
  	certificate_signing_request= "-----BEGIN CERTIFICATE REQUEST-----\nMIIC2jCCAcICAQAwgYAxCzAJBgNVBAYTAklOMRMwEQYDVQQIDApNYWhhcmFzaHRh\nMQ0wCwYDVQQHDARQdW5lMRAwDgYDVQQKDAdJQk1QdW5lMQwwCgYDVQQLDANJQk0x\nFDASBgNVBAMMC2libS5wdW5lLmluMRcwFQYJKoZIhvcNAQkBFghyYkBncy5pbjCC\nASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALZTfLDkU/q8ory+a+aBIAB7\nezyMbF4iqEpO3mVmlIfENBRq/MomhTiYID2jVWhPS4KytXa9bANgAG+NN5zMxqgH\nwUVa4fO2Th//o9RdRPZQcC3IGR+b3tpGSV+KxjsW9XV82lQp6l8HxfSBo12PZcwU\nlu53BXMiVKU6ZHSjz6eDF4MGTuPM1QWs/oexb9y6Yqqj2IicrkviJMuKhh0FN6H7\nXT2MHIF8OHT1xhs8w7U1SdUwhBa5dHvj7vqi33DU0p3s8JdXlzE9PUXKuAzXav1D\nngwE1iT5KFHkEFUr7plH5wGsy/y4tnu2xVzfGMyECS9EYmUtOA5M5RpqF5DdXGEC\nAwEAAaAUMBIGCSqGSIb3DQEJAjEFDANnc2wwDQYJKoZIhvcNAQELBQADggEBAGB7\n3lv/6fSn9rgTiHszLL9pOU9ytjOVrhNjFjDzQL73VQ0+Isb7aPHnWrLz4kT9m/60\nmgy/dHsOIF8KRP1LpOs5BYwlstD3Ss57XR8GatnrLMN4lZCjacL6A8RPhwr3x29W\nMyFntvu2caAL4ZpZpWMKHtoemXijCiFXa9Z4pZFBk4V7k0/DIEXEeyYazsSaXeTw\nXr4IFPmk7VS/NkLAht2hRhllN5NHGf/gzTsmgrgKclXtf1Z7EotnDTTIt0dFVtk1\nVX2Z7kvx9/QWbDVhPEz2uOrJnCoAm+0OpQfFc4THcP0uv0Y49B3WUG89mAjlWQKa\nU7hhc8gZ77+eaBQKD6k=\n-----END CERTIFICATE REQUEST-----"
	organization_information = {
		org_address = {
			org_address_line1= "abc"
			org_address_line2= "xyz"
			org_city="pune"
			org_country_code= "IN"
			org_state="MH"
			org_postal_code= "411045"
		}
		org_organization_name= "GSLAB"
		org_phone_number= "8657072955"
		org_fax_number = ""
	}	
	technical_contact_same_as_org_address_flag = "false"
	technical_contact = {
		tech_address = {
			tech_address_line1= "fcb"
			tech_address_line2= "pqr"
			tech_city="pune"
			tech_country_code= "IN"
			tech_state="MH"
			tech_postal_code= "411045"
		}
		tech_organization_name= "IBM"
		tech_phone_number= "8657072955"
		tech_fax_number = ""
		tech_first_name = "qwerty"
		tech_last_name = "ytrewq"
		tech_email_address = "abc@gmail.com"
		tech_title= "SSL CERT"
	}
	billing_contact = {
		billing_address = {
			billing_address_line1= "plk"
			billing_address_line2= "PLO"
			billing_city="PUNE"
			billing_country_code= "IN"
			billing_state="MH"
			billing_postal_code= "411045"
		}
		billing_organization_name= "IBM"
		billing_phone_number= "8657072955"
		billing_fax_number = ""
		billing_first_name = "ERTYU"
		billing_last_name = "SDFGHJK"
		billing_email_address = "kjjj@gsd.com"
		billing_title= "PFGHJK"
	}
	administrative_contact = {
		admin_address = {
			admin_address_line1= "fghds"
			admin_address_line2= "twyu"
			admin_city="pune"
			admin_country_code= "IN"
			admin_state="MH"
			admin_postal_code= "411045"
		}
		admin_organization_name = "GSLAB"
		admin_phone_number = "8657072955"
		admin_fax_number = ""
		admin_first_name = "DFGHJ"
		admin_last_name = "dfghjkl"
		admin_email_address = "fghjk@gshhds.com"
		admin_title= "POIUYGHJK"
	}	
	administrative_contact_same_as_technical_flag = "false"
	billing_contact_same_as_technical_flag = "false"	
	billing_address_same_as_organization_flag = "false"
	administrative_address_same_as_organization_flag = "false"
	ssl_type="SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
	renewal_flag= true
	server_count= 1
	server_type= "apache2"
	validity_months= 24
	order_approver_email_address= "admin@pune.in"	
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
* `organization_information`- (Required, set) Organization information from issuer belongs to.
	* `org_address`- (Required, string) Organization address of the issuer.
		* `org_addressLine1` - (Required, string) The address of organization who is requesting for ssl certificate.
		* `org_addressLine2` - (optional, string) The address of organization who is requesting for ssl certificate.
		* `org_city` - (Required, string) The city of organization which is requesting for ssl certificate .
		* `org_postalCode` - (Required, integer) The postal code for the city of organization.
		* `org_state` - (Required, string) The two letter state code of organization who is requesting for ssl certificate. Allowed value for country which doesn't have states is `OT`.
		* `org_countryCode` - (Required, string) The two letter country code of organization.
	* `org_organizationName` - (Required, string) Name of organization.
	* `org_phone_number` - (Required, string) Phone number of organization
	* `org_fax_number` - (Optional, string) Fax number for organization
* `technical_contact` - (Required, set) Technical contact details of issuer.
	* `tech_address` - (Optional, set) Technical address details
		* `tech_addressLine1` - (Required, string) The address for technical contact.
		* `tech_addressLine2` - (Optional, string) The address for technical contact.
		* `tech_city` - (Required, string) The city for technical contact.
		* `tech_postalCode` - (Required, integer) The postal code for technical contact.
		* `tech_state` - (Required, string) The two letter state code of technical contact. Allowed value for country which doesn't have states is `OT`.
		* `tech_countryCode` - (Required, string) The two letter country code for technical contact.
	* `tech_organizationName` - (Required, string) Name of organization for technical contact.
	* `tech_firstName` - (Required, string) The first name for technical contact.
	* `tech_lastName` - (Required, string) The last name for technical contact.
	* `tech_title` - (Required, string) The title for for technical contact.
	* `tech_emailAddress` -(Required, string) email address for technical contact.
	* `tech_phone_number`- (Required, string) phone number for technical detail.
	* `tech_fax_number` - (Optional, string) Fax number for technical detail.
* `administrative_contact` - (Optional, set) Administrator contact details.
	* `admin_address` - (Optional, set) Administrator address details.
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
	* `admin_phone_number` - (Optional, string) Phone number of administrator.
	* `admin_fax_number` - (Optional, string) Fax number for administrator.

* `billing_contact` - (Optional, set) Billing Contact details.
	* `billing_address` - (Optional, set) Billing address details.
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
	* `billing_emailAddress` - (Optional, string) email address for billing contact.
	* `billing_phone_number` - (Optional, string) Phone number for billing contact.
	* `billing_fax_number` - (Optional, string) Fax number for billing contact.
* `technicalContactSameAsOrgAddressFlag` -(Optional, bool) If your organization address and technical contact address is the same make this flag as true and skip technical contact address details.
* `administrativeContactSameAsTechnicalFlag` -(Required, bool)- If your technical contact details and administrative contact details is the same then make this as true and skip details of administrative contact.
* `billingContactSameAsTechnicalFlag` -(Required, bool)- If your technical contact details and billing contact details is the same then make this as true and skip details of billing contact. 
* `administrativeAddressSameAsOrganizationFlag` -(Required,bool)If administrative address is same as organization address then make this flag as true and skip address details.
* `billingAddressSameAsOrganizationFlag` -(Required,bool)If billing address is same as organization address then make this flag as true and skip address details. 