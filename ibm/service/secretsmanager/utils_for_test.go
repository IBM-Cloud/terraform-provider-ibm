// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"encoding/json"
	"fmt"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

var description = "Some description for this secret."
var modifiedDescription = "Modified description for this secret."
var label = "my-label-1"
var modifiedLabel = "modified-label-1"
var expirationDate = "2033-05-30T21:00:00Z"
var modifiedExpirationDate = "2034-06-12T17:34:56Z"
var customMetadata = `{"key1":"value1"}`
var modifiedCustomMetadata = `{"key2":"value2"}`
var rotationPolicy = `{
				auto_rotate = true
				interval = 1
				unit = "day"
			}`
var modifiedRotationPolicy = `{
				auto_rotate = true
				interval = 2
				unit = "month"
			}`

func getSecret(s *terraform.State, resourceName string) (secretsmanagerv2.SecretIntf, error) {
	rs, ok := s.RootModule().Resources[resourceName]
	if !ok {
		return nil, fmt.Errorf("Not found: %s", resourceName)
	}

	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return nil, err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

	id := strings.Split(rs.Primary.ID, "/")
	secretId := id[2]
	getSecretOptions.SetID(secretId)

	secret, _, err := secretsManagerClient.GetSecret(getSecretOptions)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func getClientWithInstanceEndpointTest(originalClient *secretsmanagerv2.SecretsManagerV2) *secretsmanagerv2.SecretsManagerV2 {
	// build the api endpoint
	domain := "appdomain.cloud"
	if strings.Contains(os.Getenv("IBMCLOUD_IAM_API_ENDPOINT"), "test") {
		domain = "test.appdomain.cloud"
	}
	endpoint := fmt.Sprintf("https://%s.%s.secrets-manager.%s", acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, domain)
	newClient := &secretsmanagerv2.SecretsManagerV2{
		Service: originalClient.Service.Clone(),
	}
	newClient.Service.SetServiceURL(endpoint)

	return newClient
}

func verifyAttr(actual, expected, attrName string) error {
	if actual != expected {
		return fmt.Errorf("Wrong %s. Actual: %s. Expected: %s", attrName, actual, expected)
	} else {
		return nil
	}
}

func verifyIntAttr(actual, expected int, attrName string) error {
	if actual != expected {
		return fmt.Errorf("Wrong %s. Actual: %d. Expected: %d", attrName, actual, expected)
	} else {
		return nil
	}
}

func verifyBoolAttr(actual, expected bool, attrName string) error {
	if actual != expected {
		return fmt.Errorf("Wrong %s. Actual: %t. Expected: %t", attrName, actual, expected)
	} else {
		return nil
	}
}

func verifyDateAttr(actualDate *strfmt.DateTime, expected string, attrName string) error {
	var strDate = time.Time(*actualDate).Format(time.RFC3339)
	if strDate != expected {
		return fmt.Errorf("Wrong %s. Actual: %s. Expected: %s", attrName, strDate, expected)
	} else {
		return nil
	}
}

func verifyJsonAttr(actual map[string]interface{}, expected string, attrName string) error {
	actualStr, _ := json.Marshal(actual)
	if string(actualStr) != expected {
		return fmt.Errorf("Wrong %s. Actual: %s. Expected: %s", attrName, string(actualStr), expected)
	} else {
		return nil
	}
}

// extract the value of the auto_rotate attribute from given rotation policy and convert to string
func getAutoRotate(policyIntf secretsmanagerv2.RotationPolicyIntf) string {
	autoRotate := ""
	if policy, ok := policyIntf.(*secretsmanagerv2.RotationPolicy); ok {
		if policy.AutoRotate != nil {
			autoRotate = strconv.FormatBool(*policy.AutoRotate)
		}
	}
	return autoRotate
}

// extract the value of the unit attribute from given rotation policy
func getRotationUnit(policyIntf secretsmanagerv2.RotationPolicyIntf) string {
	unit := ""
	if policy, ok := policyIntf.(*secretsmanagerv2.RotationPolicy); ok {
		unit = *policy.Unit
	}
	return unit
}

// extract the value of the interval attribute from given rotation policy and convert to string
func getRotationInterval(policyIntf secretsmanagerv2.RotationPolicyIntf) string {
	interval := ""
	if policy, ok := policyIntf.(*secretsmanagerv2.RotationPolicy); ok {
		if policy.Interval != nil {
			interval = strconv.FormatInt(*policy.Interval, 10)
		}
	}
	return interval
}

// extract the value of the rotate_keys attribute from given rotation policy and convert to string
func getRotateKeys(policyIntf secretsmanagerv2.RotationPolicyIntf) string {
	rotateKeys := ""
	if policy, ok := policyIntf.(*secretsmanagerv2.RotationPolicy); ok {
		if policy.RotateKeys != nil {
			rotateKeys = strconv.FormatBool(*policy.RotateKeys)
		}
	}
	return rotateKeys
}

func generatePublicCertCommonName() string {
	return acctest.RandStringFromCharSet(4, "123456789") + "." + acc.SecretsManagerPublicCertificateCommonName
}
