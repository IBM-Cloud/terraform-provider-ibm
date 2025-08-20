// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func init() {
	testlogger := os.Getenv("TF_LOG")
	if testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}
}

// ProtoV6ProviderFactories is used by acceptance tests.
var ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"ibm": providerserver.NewProtocol6WithError(provider.NewFrameworkProvider(version.Version)()),
}

// TestAccCheckResourceAttrNotEmpty will verify that the resource named "resourceName"
// contains an attribute named "attr" that contains a non-empty value.
func TestAccCheckResourceAttrNotEmpty(resourceName, attr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource '%s' not found", resourceName)
		}

		value := rs.Primary.Attributes[attr]
		if value == "" {
			return fmt.Errorf("resource attribute '%s.%s' should not be empty", resourceName, attr)
		}
		return nil
	}
}

// TestAccExtractImportId is used by testcases to compute the id to be used for import testing.
// The id-related attributes specified in "idAttrs" will be retrieved from the named resource and their
// values are then combined to form the import id.  In addition, the id-related attribute values are stored
// in a map so that we can later verify that the resource was in fact imported.
func TestAccExtractImportId(resourceName string, idCache map[string]string, idAttrs []string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("resource '%s' not found", resourceName)
		}
		idParts := []string{}
		for _, idAttr := range idAttrs {
			v, ok := rs.Primary.Attributes[idAttr]
			if !ok {
				return "", fmt.Errorf("attribute '%s' not found in resource '%s'", idAttr, resourceName)
			}
			idCache[idAttr] = v
			idParts = append(idParts, v)
		}
		return strings.Join(idParts, "/"), nil
	}
}

// TestAccVerifyImportState is used to verify the state after an import test has run.
// Specifically, this function will examine the terraform state looking for a resource that contains
// the attributes specified in "idAttrs" and verify that the values match those contained in "idCache".
func TestAccVerifyImportState(resourceName string, idCache map[string]string, idAttrs []string) resource.ImportStateCheckFunc {
	return func(is []*terraform.InstanceState) error {
		if len(idCache) != len(idAttrs) {
			return fmt.Errorf("idCache has not been initialized")
		}
		for _, rs := range is {
			// Build a map containing the id-related attributes from "rs".
			rsMap := make(map[string]string)
			foundAllAttrs := true
			for _, idAttr := range idAttrs {
				v, ok := rs.Attributes[idAttr]
				if !ok {
					foundAllAttrs = false
					break
				}
				rsMap[idAttr] = v
			}
			if !foundAllAttrs {
				continue
			}
			if reflect.DeepEqual(idCache, rsMap) {
				return nil
			}
		}
		return fmt.Errorf("resource '%s' not found after import", resourceName)
	}
}
