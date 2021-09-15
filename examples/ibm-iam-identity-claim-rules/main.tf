
// Provision iam_trusted_profiles_claim_rule resource instance
resource "ibm_iam_trusted_profiles_claim_rule" "iam_trusted_profiles_claim_rule_instance" {
  profile_id = "Profile-15ecfc76-2cba-4ab9-94aa-fe048547f8ff"
  type = "Profile-SAML"
  name = "nicerule"
  realm_name = "https://w3id.sso.ibm.com/auth/sps/samlidp2/saml20"
  expiration = 43200
  conditions {
				claim = "blueGroups"
				operator = "CONTAINS"
				value = "\"cloud-docs-dev\""
			}
}