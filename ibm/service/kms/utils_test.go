package kms_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

const (
	ibmKMSResourceNamePrefix = "test-acc-ibm-kms"
)

func init() {
	resource.AddTestSweepers("IBM_KMS", &resource.Sweeper{
		Name: "IBM_KMS",
		F: func(region string) error {
			client, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			conn := client.(*ExampleClient)

			instances, err := conn.DescribeComputeInstances()
			if err != nil {
				return fmt.Errorf("Error getting instances: %s", err)
			}
			for _, instance := range instances {
				if strings.HasPrefix(instance.Name, "test-acc") {
					err := conn.DestroyInstance(instance.ID)

					if err != nil {
						log.Printf("Error destroying %s during sweep: %s", instance.Name, err)
					}
				}
			}
			return nil
		},
	})
}

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
