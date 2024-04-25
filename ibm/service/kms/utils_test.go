package kms_test

import "fmt"

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
