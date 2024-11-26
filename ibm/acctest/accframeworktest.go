package acctest

import (
	"fmt"
	"os"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var (
	// TestAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can reattach.
	TestAccProtoV6ProviderFactories map[string]func() (tfprotov6.ProviderServer, error)

	// testAccProviderConfigure ensures Provider is only configured once
)

func init() {
	// Initialize logging
	if testlogger := os.Getenv("TF_LOG"); testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}

	// Initialize provider factories
	frameworkProvider := provider.NewFrameworkProvider(version.Version)
	TestAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		ProviderName: providerserver.NewProtocol6WithError(frameworkProvider()),
	}
}

// TestAccPreCheck verifies required provider attributes are set
func TestAccFrameworkPreCheck(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("IAAS_CLASSIC_API_KEY"); v == "" {
		t.Fatal("IAAS_CLASSIC_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("IAAS_CLASSIC_USERNAME"); v == "" {
		t.Fatal("IAAS_CLASSIC_USERNAME must be set for acceptance tests")
	}
}

// TestAccPreCheckEnterprise verifies enterprise-specific requirements
func TestAccFrameworkPreCheckEnterprise(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
}

// ConfigureProvider returns basic provider configuration
func ConfigureProvider(region string) string {
	return fmt.Sprintf(`
provider "ibm" {
  region = %q
}
`, region)
}

// ConfigureAlternateProvider returns configuration for alternate provider
func ConfigureAlternateProvider() string {
	return fmt.Sprintf(`
provider %q {
  region = %q
}
`, ProviderNameAlternate, RegionAlternate())
}

// ConfigureProviderWithEnterprise returns provider configuration with enterprise settings
func ConfigureProviderWithEnterprise(region string) string {
	return fmt.Sprintf(`
provider "ibm" {
  region = %q
  enterprise_account_id = "%s"
}
`, region, os.Getenv("ENTERPRISE_ACCOUNT_ID"))
}

// Function to add alternate provider if needed
func AddAlternateProvider() {
	if TestAccProtoV6ProviderFactories == nil {
		TestAccProtoV6ProviderFactories = make(map[string]func() (tfprotov6.ProviderServer, error))
	}

	frameworkProvider := provider.NewFrameworkProvider(version.Version)
	TestAccProtoV6ProviderFactories[ProviderNameAlternate] = providerserver.NewProtocol6WithError(frameworkProvider())
}
