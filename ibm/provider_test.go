package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var cfOrganization string
var cfSpace string
var ibmid1 string
var ibmid2 string

func init() {
	cfOrganization = os.Getenv("IBM_ORG")
	if cfOrganization == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ORG for testing ibm_org  resource Some tests for that resource will fail if this is not set correctly")
	}
	cfSpace = os.Getenv("IBM_SPACE")
	if cfSpace == "" {
		fmt.Println("[WARN] Set the environment variable IBM_SPACE for testing ibm_space  resource Some tests for that resource will fail if this is not set correctly")
	}
	ibmid1 = os.Getenv("IBM_ID1")
	if ibmid1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID1 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	ibmid2 = os.Getenv("IBM_ID2")
	if ibmid2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID2 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}
}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ibm": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("BM_API_KEY"); v == "" {
		t.Fatal("BM_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("SL_API_KEY"); v == "" {
		t.Fatal("SL_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("SL_USERNAME"); v == "" {
		t.Fatal("SL_USERNAME must be set for acceptance tests")
	}
}
