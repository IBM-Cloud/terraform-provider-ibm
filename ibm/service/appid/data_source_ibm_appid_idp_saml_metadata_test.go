package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDIDPSAMLMetadataDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDIDPSAMLMetadataDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_saml_metadata.meta", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttrSet("data.ibm_appid_idp_saml_metadata.meta", "metadata"),
				),
			},
		},
	})
}

func setupAppIDIDPSAMLMetadataDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`	
		data "ibm_appid_idp_saml_metadata" "meta" {
			tenant_id = "%s"
		}
	`, tenantID)
}
