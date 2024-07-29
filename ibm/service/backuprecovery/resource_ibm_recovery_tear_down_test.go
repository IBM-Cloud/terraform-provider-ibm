// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"
)

func TestAccIbmRecoveryTearDownBasic(t *testing.T) {
	// var conf backuprecoveryv1.CancelRecovery

	// resource.Test(t, resource.TestCase{
	// 	PreCheck:  func() { acc.TestAccPreCheck(t) },
	// 	Providers: acc.TestAccProviders,
	// 	Steps: []resource.TestStep{
	// 		resource.TestStep{
	// 			Config: testAccCheckIbmRecoveryTearDownConfigBasic(),
	// 			Check: resource.ComposeAggregateTestCheckFunc(
	// 				testAccCheckIbmRecoveryTearDownDoesNotExist("ibm_recovery_cancel.recovery_cancel_instance", conf),
	// 			),
	// 		},
	// 		resource.TestStep{
	// 			ResourceName:      "ibm_recovery_cancel.recovery_cancel",
	// 			ImportState:       true,
	// 			ImportStateVerify: true,
	// 		},
	// 	},
	// })
}

func testAccCheckIbmRecoveryTearDownConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_recovery_cancel" "recovery_cancel_instance" {
		}
	`)
}

// func testAccCheckIbmRecoveryTearDownDoesNotExist(n string, obj backuprecoveryv1.CancelRecovery) resource.TestCheckFunc {

// return func(s *terraform.State) error {
// 	rs, ok := s.RootModule().Resources[n]
// 	if !ok {
// 		return fmt.Errorf("Not found: %s", n)
// 	}

// 	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
// 	if err != nil {
// 		return err
// 	}

// 	getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

// 	parts, err := flex.SepIdParts(rs.Primary.ID, "/")
// 	if err != nil {
// 		return err
// 	}

// 	getRecoveryByIdOptions.SetID(parts[0])
// 	getRecoveryByIdOptions.SetID(parts[1])

// 	_, _, err = backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
// 	if err != nil {
// 		return nil
// 	}

// }
// }
