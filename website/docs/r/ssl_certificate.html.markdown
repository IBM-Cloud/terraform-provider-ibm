---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ssl_certificate"
description: |-
  Fetches the IBM Cloud SSL certificate.
---

# ibm_ssl_certificate
Create, update, and delete an SSL certificate resource. This allows SSL certificates to be requested, and delete request for SSL certificates. For more information, about SSL certificates, see [accessing SSL certificates](https://cloud.ibm.com/docs/ssl-certificates?topic=ssl-certificates-accessing-ssl-certificates).

**Note**

For more information, see [IBM Cloud Classic Infrastructure(SoftLayer) security certificates request documentation](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Security_Certificate/).

## Example usage
In the following example, you can use a certificate on file:

```terraform
resource "ibm_ssl_certificate" "my_ssllllll" {
  certificate_signing_request = "-----BEGIN CERTIFICATE REQUEST-----\nCERTIFICATE CONTENT\n-----END CERTIFICATE REQUEST-----"
  organization_information {
    org_address {
      org_address_line1 = "abc"
      org_address_line2 = "xyz"
      org_city          = "pune"
      org_country_code  = "IN"
      org_state         = "MH"
      org_postal_code   = "411045"
    }
    org_organization_name = "GSLAB"
    org_phone_number      = "8657072955"
    org_fax_number        = ""
  }
  technical_contact_same_as_org_address_flag = "false"
  technical_contact {
    tech_address {
      tech_address_line1 = "fcb"
      tech_address_line2 = "pqr"
      tech_city          = "pune"
      tech_country_code  = "IN"
      tech_state         = "MH"
      tech_postal_code   = "411045"
    }
    tech_organization_name = "IBM"
    tech_phone_number      = "8657072955"
    tech_fax_number        = ""
    tech_first_name        = "qwerty"
    tech_last_name         = "ytrewq"
    tech_email_address     = "abc@gmail.com"
    tech_title             = "SSL CERT"
  }
  billing_contact {
    billing_address {
      billing_address_line1 = "plk"
      billing_address_line2 = "PLO"
      billing_city          = "PUNE"
      billing_country_code  = "IN"
      billing_state         = "MH"
      billing_postal_code   = "411045"
    }
    billing_organization_name = "IBM"
    billing_phone_number      = "8657072955"
    billing_fax_number        = ""
    billing_first_name        = "ERTYU"
    billing_last_name         = "SDFGHJK"
    billing_email_address     = "kjjj@gsd.com"
    billing_title             = "PFGHJK"
  }
  administrative_contact {
    admin_address {
      admin_address_line1 = "fghds"
      admin_address_line2 = "twyu"
      admin_city          = "pune"
      admin_country_code  = "IN"
      admin_state         = "MH"
      admin_postal_code   = "411045"
    }
    admin_organization_name = "GSLAB"
    admin_phone_number      = "8657072955"
    admin_fax_number        = ""
    admin_first_name        = "DFGHJ"
    admin_last_name         = "dfghjkl"
    admin_email_address     = "fghjk@gshhds.com"
    admin_title             = "POIUYGHJK"
  }
  administrative_contact_same_as_technical_flag    = "false"
  billing_contact_same_as_technical_flag           = "false"
  billing_address_same_as_organization_flag        = "false"
  administrative_address_same_as_organization_flag = "false"
  ssl_type                                         = "SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
  renewal_flag                                     = true
  server_count                                     = 1
  server_type                                      = "apache2"
  validity_months                                  = 24
  order_approver_email_address                     = "admin@pune.in"
}

```


## Argument reference 
Review the argument references that you can specify for your resource.

- `administrative_contact` - (Optional, Set) Administrator contact details.

  Nested scheme for `administrative_contact`:
  - `admin_address` - (Optional, Set) Administrator address details.

    Nested scheme for `admin_address`:
    - `admin_addressLine1` - (Optional, String) The address for administrative contact.
    - `admin_addressLine2` - (Optional, String) The address for administrative contact.	
    - `admin_city` - (Optional, String) The city for administrative contact.	
    - `admin_countryCode` - (Optional, String) The two letter country code for administrative contact.
    - `admin_postalCode` - (Optional, Integer) The postal code for administrative contact.	
    - `admin_state` - (Optional, String) The two letter state code of administrative contact. Allowed value for country which doesn't have states is `OT`.	
  - `admin_emailAddress` -(Optional, String) email address for administrative contact.
  - `admin_fax_number` - (Optional, String) Fax number for administrator.
  - `admin_firstName` - (Optional, String) The first name for administrative contact.
  - `admin_lastName` - (Optional, String) The last name for administrative contact.
  - `admin_organizationName` - (Optional, String) Name of organization for administrative contact.
   - `admin_phone_number` - (Optional, String) Phone number of administrator.
   - `admin_title` - (Optional, String) The title for administrative contact.
- `administrativeContactSameAsTechnicalFlag` -(Required, Bool) If your technical contact details and administrative contact details are the same then make this as true and skip details of administrative contact.
- `administrativeAddressSameAsOrganizationFlag` -(Required, Bool) If administrative address is same as organization address then make this flag as true and skip address details.
- `billingAddressSameAsOrganizationFlag` -(Required, Bool) If billing address is same as organization address then make this flag as true and skip address details.
- `billingContactSameAsTechnicalFlag` -(Required, Bool) If your technical contact details and billing contact details are the same then make this as true and skip details of billing contact.
- `billing_contact` - (Optional, Set) Billing Contact details.

  Nested scheme for `billing_contact`:
  - `billing_address` - (Optional, Set) Billing address details.

    Nested scheme for `billing_address`:
    - `billing_addressLine1` - (Optional, String) The address for billing contact.	
    - `billing_addressLine2` - (Optional, String) The address for billing contact.	
     - `billing_countryCode` - (Optional, String) The two letter country code for billing contact.
    - `billing_city` - (Optional, String) The city for billing contact.	
    - `billing_postalCode` - (Optional, Integer) The postal code for billing contact.	
    - `billing_state` - (Optional, String) The two letter state code of billing contact. Allowed value for country which doesn't have states is `OT`.	
  - `billing_emailAddress` - (Optional, String) email address for billing contact.
  - `billing_fax_number` - (Optional, String) Fax number for billing contact.
  - `billing_firstName` - (Optional, String) The first name for billing contact.
  - `billing_lastName` - (Optional, String) The last name for billing contact.
  - `billing_organizationName` - (Optional, String) Name of organization for billing contact.
  - `billing_phone_number` - (Optional, String) Phone number for billing contact.
  - `billing_title` - (Optional, String) The title for billing contact.
- `certificateSigningRequest` - (Required, String) The Certificate Signing Request which is specially formatted encrypted message sent from a Secure Sockets Layer (SSL) digital certificate applicant to a certificate authority.
- `orderApproverEmailAddress` - (Required, String) The email of approver to approve SSL certificate request.
- `organization_information` - (Required, Set) Organization information from issuer belongs to.

  Nested scheme for `organization_information`:
  - `org_address` - (Required, String) Organization address of the issuer.	

    Nested scheme for `org_address`:
    - `org_addressLine1` - (Required, String) The address of organization who is requesting for SSL certificate.	
    - `org_addressLine2` - (Optional, String) The address of organization who is requesting for SSL certificate.	
    - `org_city` - (Required, String) The city of organization which is requesting for SSL certificate.	
    - `org_countryCode` - (Required, String) The two letter country code of organization.
    - `org_postalCode` - (Required, Integer) The postal code for the city of organization.	
    - `org_state` - (Required, String) The two letter state code of organization who is requesting for SSL certificate. Allowed value for country which doesn't have states is `OT`.	
  - `org_fax_number` - (Optional, String) Fax number for organization.
  - `org_organizationName` - (Required, String) Name of organization.
  - `org_phone_number` - (Required, String) Phone number of organization
- `sslType` - (Required, String) The SSL certificate type.
- `serverType` - (Required, String) The server type for which we are requesting SSL certificate.
- `serverCount` - (Required, String) The number of servers with provided server type.
- `technical_contact` - (Required, Set) Technical contact details of issuer.

  Nested scheme for `technical_contact`:
  - `tech_address` - (Optional, Set) Technical address details.

    Nested scheme for `tech_address`:	
    - `tech_addressLine1` - (Required, String) The address for technical contact.	
    - `tech_addressLine2` - (Optional, String) The address for technical contact.	
    - `tech_city` - (Required, String) The city for technical contact.
    - `tech_countryCode` - (Required, String) The two letter country code for technical contact.
    - `tech_postalCode` - (Required, Integer) The postal code for technical contact.	
    - `tech_state` - (Required, String) The two letter state code of technical contact. Allowed value for country which doesn't have states is `OT`.	
  - `tech_emailAddress` -(Required, String) email address for technical contact.
  - `tech_fax_number` - (Optional, String) Fax number for technical detail.
  - `tech_firstName` - (Required, String) The first name for technical contact.
  - `tech_lastName` - (Required, String) The last name for technical contact.
  - `tech_organizationName` - (Required, String) Name of organization for technical contact.
  - `tech_phone_number`- (Required, String) phone number for technical detail.
  - `tech_title` - (Required, String) The title for technical contact.
- `technicalContactSameAsOrgAddressFlag` -(Optional, Bool) If your organization address and technical contact address are the same make this flag as true and skip technical contact address details.
- `validityMonths` - (Required, Integer) The validity of SSL certificate in months it should be multiple of 12.
