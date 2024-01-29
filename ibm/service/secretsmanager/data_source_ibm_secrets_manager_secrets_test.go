// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

//func TestAccIBMSecretsManagerSecretsDataSourceBasic(t *testing.T) {
//	resource.Test(t, resource.TestCase{
//		PreCheck:  func() { acc.TestAccPreCheck(t) },
//		Providers: acc.TestAccProviders,
//		Steps: []resource.TestStep{
//			{
//				Config: testAccCheckIBMSecretsManagerSecretsDataSourceConfigBasic(),
//				Check: resource.ComposeTestCheckFunc(
//					resource.TestCheckResourceAttr("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secret_type", acc.SecretsManagerSecretType),
//					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "id"),
//					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secret_type"),
//					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "metadata.#"),
//					resource.TestCheckResourceAttrSet("data.ibm_secrets_manager_secrets.secrets_manager_secrets", "secrets.#"),
//				),
//			},
//		},
//	})
//}
//
//func testAccCheckIBMSecretsManagerSecretsDataSourceConfigBasic() string {
//	return fmt.Sprintf(`
//		data "ibm_secrets_manager_secrets" "secrets_manager_secrets" {
//			instance_id = "%s"
//			secret_type = "%s"
//		}
//
//		output "WorkSpaceValues" {
//			value = data.ibm_secrets_manager_secrets.secrets_manager_secrets.secret_type
//		}
//	`, acc.SecretsManagerInstanceID, acc.SecretsManagerSecretType)
//}
