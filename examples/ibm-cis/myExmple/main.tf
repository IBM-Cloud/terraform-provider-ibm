terraform {
   required_providers {
      ibm = {
         source = "IBM-Cloud/ibm"
         version = "1.43.0"
      }
    }
  }

resource "ibm_cis_domain" "web_domain" {
        cis_id = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:1a9174b6-0106-417a-844b-c8eb43a72f63::"
        domain = "t2.test.com"
  type = "partial"
}
