package ibm

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"strings"
	"testing"
)

func TestAccIBMAppIDCloudDirectoryTemplate_basic(t *testing.T) {
	htmlBody := "<HTML><HEAD><TITLE>Test title</TITLE></HEAD><BODY>test</BODY></HTML>"
	b64HTML := b64.StdEncoding.EncodeToString([]byte(htmlBody))
	textBody := "This is the test"
	subject := "Please Verify Your Email Address %%{user.displayName} TEST"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDCloudDirectoryTemplateConfig(appIDTenantID, subject, htmlBody, textBody),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "template_name", "USER_VERIFICATION"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "subject", strings.Replace(subject, "%%", "%", 1)),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "html_body", htmlBody),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "base64_encoded_html_body", b64HTML),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_template.test_tpl", "plain_text_body", textBody),
				),
			},
		},
	})
}

func setupAppIDCloudDirectoryTemplateConfig(tenantID string, subject string, htmlBody string, textBody string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_cloud_directory_template" "test_tpl" {
			tenant_id = "%s"
			template_name = "USER_VERIFICATION"
			subject = "%s"
			html_body = "%s"
			plain_text_body = "%s"
		}	
	`, tenantID, subject, htmlBody, textBody)
}
