



resource "ibm_cis_ruleset_rule" "config1" {
    cis_id    = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:1a9174b6-0106-417a-844b-c8eb43a72f63::"
    domain_id = "601b728b86e630c744c81740f72570c3"
    ruleset_id = "eb5efc50d5ec49d8b0b0f44b357b8d7b"
      rule {
        action =  "log"

        description = "Anagha test rule 1"
        enabled = false
        expression = "true"
        position {
          index = 1
          
        }
      }
}

resource "ibm_cis_ruleset_rule" "config2" {
    cis_id    = "crn:v1:staging:public:internet-svcs-ci:global:a/01652b251c3ae2787110a995d8db0135:1a9174b6-0106-417a-844b-c8eb43a72f63::"
    domain_id = "601b728b86e630c744c81740f72570c3"
    ruleset_id = "eb5efc50d5ec49d8b0b0f44b357b8d7b"
      rule {
        action =  "log"

        description = "Anagha test rule 2"
        enabled = true
        expression = "true"
        position {
          index = 2
          after = "3"
          
        }
      }
}

