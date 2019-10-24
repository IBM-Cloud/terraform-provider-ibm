package ibm

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMSSLCertificate_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccIBMSSLCertificateConfig_basic,

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "order_approver_email_address", "admin@pune.in"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "server_type", "apache2"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "server_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "validity_months", "24"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "ssl_type", "SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"),
				),
			},
		},
	})
}

var testAccIBMSSLCertificateConfig_basic = `
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
`
