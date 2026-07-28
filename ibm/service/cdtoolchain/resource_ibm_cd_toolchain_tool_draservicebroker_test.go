package cdtoolchain_test

import (
        "fmt"
        "testing"

        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
        "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
        "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

        acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
        "github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
        "github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
        "github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
)

func TestAccIBMCdToolchainToolDRAServiceBroker_Basic(t *testing.T) {
	var conf cdtoolchainv2.GetToolByIDResponse
	toolchainID := fmt.Sprintf("tf_toolchain_%d", acctest.RandIntRange(10, 100))
	toolName := fmt.Sprintf("tf_tool_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCdToolchainToolDRAServiceBrokerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainToolDRAServiceBrokerBasic(toolchainID, toolName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCdToolchainToolDRAServiceBrokerExists("ibm_cd_toolchain_tool_draservicebroker.tool_draservicebroker", conf),
					resource.TestCheckResourceAttr("ibm_cd_toolchain_tool_draservicebroker.tool_draservicebroker", "name", toolName),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainToolDRAServiceBrokerExists(n string, obj cdtoolchainv2.GetToolByIDResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		cdToolchainClient, err := testAccProvider.Meta().(ClientSession).CdToolchainV2()
		if err != nil {
			return err
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		getToolByIDResponse, _, err := cdToolchainClient.GetToolByIDWithContext(context.TODO(), getToolByIDOptions)
		if err != nil {
			return err
		}

		obj = *getToolByIDResponse
		return nil
	}
}

func testAccCheckIBMCdToolchainToolDRAServiceBrokerDestroy(s *terraform.State) error {
	cdToolchainClient, err := testAccProvider.Meta().(ClientSession).CdToolchainV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_cd_toolchain_tool_draservicebroker" {
			continue
		}

		getToolByIDOptions := &cdtoolchainv2.GetToolByIDOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getToolByIDOptions.SetToolchainID(parts[0])
		getToolByIDOptions.SetToolID(parts[1])

		_, _, err = cdToolchainClient.GetToolByIDWithContext(context.TODO(), getToolByIDOptions)
		if err == nil {
			return fmt.Errorf("DRA service broker tool still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMCdToolchainToolDRAServiceBrokerBasic(toolchainID string, toolName string) string {
	return fmt.Sprintf(`
resource "ibm_cd_toolchain_tool_draservicebroker" "tool_draservicebroker" {
  toolchain_id = "%s"
  name = "%s"
  parameters {
    dashboard_url = "https://draservicebroker.example.com/dashboard"
    broker_id = "broker-id"
    api_key_id = "api-key-id"
    api_key_secret = "api-key-secret"
  }
}
`, toolchainID, toolName)
}
