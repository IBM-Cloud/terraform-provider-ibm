package ibm

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
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
						"ibm_ssl_certificate.myssllll", "email", "admin@pune.in"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "server_type", "apache2"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "server_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "validity_month", "24"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "title", "something"),
					resource.TestCheckResourceAttr(
						"ibm_ssl_certificate.myssllll", "ssl_type", "SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"),
				),
			},
		},
	})
}

var testAccIBMSSLCertificateConfig_basic = `
resource "ibm_ssl_certificate" "myssllll"{
	csr= "-----BEGIN CERTIFICATE REQUEST-----\nMIIC2jCCAcICAQAwgYAxCzAJBgNVBAYTAklOMRMwEQYDVQQIDApNYWhhcmFzaHRh\nMQ0wCwYDVQQHDARQdW5lMRAwDgYDVQQKDAdJQk1QdW5lMQwwCgYDVQQLDANJQk0x\nFDASBgNVBAMMC2libS5wdW5lLmluMRcwFQYJKoZIhvcNAQkBFghyYkBncy5pbjCC\nASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALZTfLDkU/q8ory+a+aBIAB7\nezyMbF4iqEpO3mVmlIfENBRq/MomhTiYID2jVWhPS4KytXa9bANgAG+NN5zMxqgH\nwUVa4fO2Th//o9RdRPZQcC3IGR+b3tpGSV+KxjsW9XV82lQp6l8HxfSBo12PZcwU\nlu53BXMiVKU6ZHSjz6eDF4MGTuPM1QWs/oexb9y6Yqqj2IicrkviJMuKhh0FN6H7\nXT2MHIF8OHT1xhs8w7U1SdUwhBa5dHvj7vqi33DU0p3s8JdXlzE9PUXKuAzXav1D\nngwE1iT5KFHkEFUr7plH5wGsy/y4tnu2xVzfGMyECS9EYmUtOA5M5RpqF5DdXGEC\nAwEAAaAUMBIGCSqGSIb3DQEJAjEFDANnc2wwDQYJKoZIhvcNAQELBQADggEBAGB7\n3lv/6fSn9rgTiHszLL9pOU9ytjOVrhNjFjDzQL73VQ0+Isb7aPHnWrLz4kT9m/60\nmgy/dHsOIF8KRP1LpOs5BYwlstD3Ss57XR8GatnrLMN4lZCjacL6A8RPhwr3x29W\nMyFntvu2caAL4ZpZpWMKHtoemXijCiFXa9Z4pZFBk4V7k0/DIEXEeyYazsSaXeTw\nXr4IFPmk7VS/NkLAht2hRhllN5NHGf/gzTsmgrgKclXtf1Z7EotnDTTIt0dFVtk1\nVX2Z7kvx9/QWbDVhPEz2uOrJnCoAm+0OpQfFc4THcP0uv0Y49B3WUG89mAjlWQKa\nU7hhc8gZ77+eaBQKD6k=\n-----END CERTIFICATE REQUEST-----"
	address1= "Baner"
	address2= "Mahalunge"
	city="Pune"
	country_code= "AF"
	state="OT"
	postal_code= "411007"
	org_name= "IBMPune"
	phone_no= "952741928"
	first_name= "ravi"
	ssl_type="SSL_CERTIFICATE_QUICKSSL_PREMIUM_2_YEAR"
	last_name= "bele"
	renewal= true
	server_count= 1
	server_type= "apache2"
	validity_month= 24
	email= "admin@pune.in"
	title= "something"
}
`
