package kms_test

import (
	"fmt"
)

const (
	ibmKMSResourceNamePrefix = "test-acc-ibm-kms"
)

// TODO: adding sweepers to clean up after tests would be nice.
// Reference: https://developer.hashicorp.com/terraform/plugin/sdkv2/testing/acceptance-tests/sweepers

// func init() {
// 	resource.AddTestSweepers("IBM_KMS", &resource.Sweeper{
// 		Name: "IBM_KMS",
// 		F: func(region string) error {
// 			rsConClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
// 			if err != nil {
// 				return flex.FmtErrorf("Error getting client: %s", err)
// 			}
// 			conn := client.(*ExampleClient)

// 			instances, err := conn.DescribeComputeInstances()
// 			if err != nil {
// 				return flex.FmtErrorf("Error getting instances: %s", err)
// 			}
// 			for _, instance := range instances {
// 				if strings.HasPrefix(instance.Name, "test-acc") {
// 					err := conn.DestroyInstance(instance.ID)

// 					if err != nil {
// 						log.Printf("Error destroying %s during sweep: %s", instance.Name, err)
// 					}
// 				}
// 			}
// 			return nil
// 		},
// 	})
// }

func addPrefixToResourceName(resourceName string) string {
	return fmt.Sprintf("%s-%s", ibmKMSResourceNamePrefix, resourceName)
}

func convertMapToTerraformConfigString(mapToConv map[string]string) string {
	// For maximum flexibility, this will not escape or parse anything
	output := "{"
	for k, v := range mapToConv {
		output += fmt.Sprintf("%s = %s \n", k, v)
	}
	output += "}"
	return output
}

func wrapQuotes(input string) string {
	return fmt.Sprintf("\"%s\"", input)
}
