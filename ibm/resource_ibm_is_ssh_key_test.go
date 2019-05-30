package ibm

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/compute"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISSSHKey_basic(t *testing.T) {
	var key *models.Key
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	name := fmt.Sprintf("terraformsecurityuat_create_step_name_%d", acctest.RandInt())
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: checkKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISKeyConfig(publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISKeyExists("ibm_is_ssh_key.isExampleKey", &key),
					resource.TestCheckResourceAttr(
						"ibm_is_ssh_key.isExampleKey", "name", name),
				),
			},
		},
	})
}

func checkKeyDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	keyC := compute.NewKeyClient(sess)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_network_acl" {
			continue
		}

		_, err := keyC.Get(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("key still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISKeyExists(n string, key **models.Key) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		keyC := compute.NewKeyClient(sess)
		foundKey, err := keyC.Get(rs.Primary.ID)

		if err != nil {
			return err
		}

		*key = foundKey
		return nil
	}
}

func testAccCheckIBMISKeyConfig(publicKey, name string) string {
	return fmt.Sprintf(`
		resource "ibm_is_ssh_key" "isExampleKey" {
			name = "%s"
			public_key = "%s"
		}
	`, name, publicKey)
}
