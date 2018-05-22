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
						"ibm_ssl_certificate.myssllll", "orderApproverEmailAddress", "admin@pune.in"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "serverType", "apache2"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "serverCount", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "validityMonths", "24"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "sslType", "SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"),
				),
			},
		},
	})
}

var testAccIBMSSLCertificateConfig_basic = `
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
  technicalContactSameAsOrgAddressFlag = "false"
  technicalContact = {
	  tech_address = {
		  tech_addressLine1= "Baner"
		  tech_addressLine2= "Mahalunge"
		  tech_city="Pune"
		  tech_countryCode= "AF"
		  tech_state="OT"
		  tech_postalCode= "411007"
	  }
	  tech_organizationName= "IBMPune"
	  tech_phoneNumber= "952741928"
	  tech_firstName = "poy"
	  tech_lastName = "joy"
	  tech_emailAddress = "email.prefix@yahoo.com"
	  tech_title= "something"
  }
  billingContact = {
	  billing_address = {
		  billing_addressLine1= "Baner"
		  billing_addressLine2= "Mahalunge"
		  billing_city="Pune"
		  billing_countryCode= "AF"
		  billing_state="OT"
		  billing_postalCode= "411007"
	  }
	  billing_organizationName= "IBMPune"
	  billing_phoneNumber= "9876542102"
	  billing_firstName = "pqr"
	  billing_lastName = "qwe"
	  billing_emailAddress = "email@domain.com"
	  billing_title= "plzeru"
  }
  administrativeContact = {
	  admin_address = {
		  admin_addressLine1= "Baner"
		  admin_addressLine2= "Mahalunge"
		  admin_city="Pune"
		  admin_countryCode= "AF"
		  admin_state="OT"
		  admin_postalCode= "411007"
	  }
	  admin_organizationName= "IBMPune"
	  admin_phoneNumber= "952741928"
	  admin_firstName = "lkj"
	  admin_lastName = "ply"
	  admin_emailAddress = "email.prefix@domain.com"
	  admin_title= "see"
  }
  administrativeContactSameAsTechnicalFlag = "false"
  billingContactSameAsTechnicalFlag = "false"	
  sslType="SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
  renewalFlag= true
  serverCount= 1
  serverType= "apache2"
  validityMonths= 24
  orderApproverEmailAddress= "admin@pune.in"	
}
`
