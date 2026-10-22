// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package watsonx_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var regexpProjectLocation = regexp.MustCompile("^/v2/projects/[a-f0-9-]+$")

// watsonxEnv holds the inputs the watsonx acceptance tests need. They point at
// existing COS, KMS and watsonx.ai Runtime instances, because creating those is
// slow and costs money.
type watsonxEnv struct {
	Region      string
	COSGUID     string
	COSCRN      string
	KMSKeyCRN   string
	RuntimeName string
	RuntimeGUID string
	RuntimeCRN  string
}

func requireWatsonxEnv(t *testing.T, needKMS, needRuntime bool) watsonxEnv {
	e := watsonxEnv{
		Region:      os.Getenv("IBM_WATSONX_REGION"),
		COSGUID:     os.Getenv("IBM_WATSONX_COS_GUID"),
		COSCRN:      os.Getenv("IBM_WATSONX_COS_CRN"),
		KMSKeyCRN:   os.Getenv("IBM_WATSONX_KMS_KEY_CRN"),
		RuntimeName: os.Getenv("IBM_WATSONX_RUNTIME_NAME"),
		RuntimeGUID: os.Getenv("IBM_WATSONX_RUNTIME_GUID"),
		RuntimeCRN:  os.Getenv("IBM_WATSONX_RUNTIME_CRN"),
	}
	if e.Region == "" || e.COSGUID == "" {
		t.Skip("IBM_WATSONX_REGION and IBM_WATSONX_COS_GUID must be set for watsonx acceptance tests")
	}
	if needKMS && e.KMSKeyCRN == "" {
		t.Skip("IBM_WATSONX_KMS_KEY_CRN must be set")
	}
	if needRuntime && (e.COSCRN == "" || e.RuntimeName == "" || e.RuntimeGUID == "" || e.RuntimeCRN == "") {
		t.Skip("IBM_WATSONX_COS_CRN and IBM_WATSONX_RUNTIME_NAME/GUID/CRN must be set")
	}
	return e
}

func testAccWatsonxImportID(region, resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("%s not found in state", resourceName)
		}
		return region + "/" + rs.Primary.ID, nil
	}
}
