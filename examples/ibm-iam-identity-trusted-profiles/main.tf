
// Provision iam_trusted_profiles resource instance
resource "ibm_iam_trusted_profiles" "iam_trusted_profiles_instance" {
  name = "profile789"
  description = "my nice profile desc"
  account_id = "29671bdd39ef4a1e9c07c16b13466908"
}
