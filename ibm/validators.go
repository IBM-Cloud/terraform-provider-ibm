package ibm

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/hashicorp/terraform/helper/schema"
	homedir "github.com/mitchellh/go-homedir"
)

func validateSecondaryIPCount(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value != 4 && value != 8 {
		errors = append(errors, fmt.Errorf(
			"%q must be either 4 or 8", k))
	}
	return
}

func validateServiceTags(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 2048 {
		errors = append(errors, fmt.Errorf(
			"%q must contain tags whose maximum length is 2048 characters", k))
	}
	return
}

func validateAllowedStringValue(validValues []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		input := v.(string)
		existed := false
		for _, s := range validValues {
			if s == input {
				existed = true
				break
			}
		}
		if !existed {
			errors = append(errors, fmt.Errorf(
				"%q must contain a value from %#v, got %q",
				k, validValues, input))
		}
		return

	}
}

func validateRoutePath(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	//Somehow API allows this
	if value == "" {
		return
	}

	if (len(value) < 2) || (len(value) > 128) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must contain from 2 to 128 characters ", k, value))
	}
	if !(strings.HasPrefix(value, "/")) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must start with a forward slash '/'", k, value))

	}
	if strings.Contains(value, "?") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must not contain a '?'", k, value))
	}

	return
}

func validateRoutePort(v interface{}, k string) (ws []string, errors []error) {
	return validatePortRange(1024, 65535)(v, k)
}

func validateAppPort(v interface{}, k string) (ws []string, errors []error) {
	return validatePortRange(1024, 65535)(v, k)
}

func validatePortRange(start, end int) func(v interface{}, k string) (ws []string, errors []error) {
	f := func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if (value < start) || (value > end) {
			errors = append(errors, fmt.Errorf(
				"%q (%d) must be in the range of %d to %d", k, value, start, end))
		}
		return
	}
	return f
}

func validateDomainName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if !(strings.Contains(value, ".")) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must contain a '.',example.com,foo.example.com", k, value))
	}

	return
}

func validateAppInstance(v interface{}, k string) (ws []string, errors []error) {
	instances := v.(int)
	if instances < 0 {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be greater than 0", k, instances))
	}
	return

}

func validateWorkerNum(v interface{}, k string) (ws []string, errors []error) {
	workerNum := v.(int)
	if workerNum <= 0 {
		errors = append(errors, fmt.Errorf(
			"%q  must be greater than 0", k))
	}
	return

}

func validateAppZipPath(v interface{}, k string) (ws []string, errors []error) {
	path := v.(string)
	applicationZip, err := homedir.Expand(path)
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"%q (%q) home directory in the given path couldn't be expanded", k, path))
	}
	if !helpers.FileExists(applicationZip) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) doesn't exist", k, path))
	}

	return

}

func validateNotes(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 1000 {
		errors = append(errors, fmt.Errorf(
			"%q should not exceed 1000 characters", k))
	}
	return
}

func validatePublicBandwidth(v interface{}, k string) (ws []string, errors []error) {
	bandwidth := v.(int)
	if bandwidth < 0 {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must be greater than 0", k, bandwidth))
		return
	}
	validBandwidths := []int{250, 1000, 5000, 10000, 20000}
	for _, b := range validBandwidths {
		if b == bandwidth {
			return
		}
	}
	errors = append(errors, fmt.Errorf(
		"%q (%d) must be one of the value from %d", k, bandwidth, validBandwidths))
	return

}

func validateMaxConn(v interface{}, k string) (ws []string, errors []error) {
	maxConn := v.(int)
	if maxConn < 1 || maxConn > 64000 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 64000",
			k))
		return
	}
	return
}

func validateWeight(v interface{}, k string) (ws []string, errors []error) {
	weight := v.(int)
	if weight < 0 || weight > 100 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 100",
			k))
	}
	return
}

func validateSizePerZone(v interface{}, k string) (ws []string, errors []error) {
	sizePerZone := v.(int)
	if sizePerZone <= 0 {
		errors = append(errors, fmt.Errorf(
			"%q must be greater than 0",
			k))
	}
	return
}

func validateInterval(v interface{}, k string) (ws []string, errors []error) {
	interval := v.(int)
	if interval < 2 || interval > 60 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 2 and 60",
			k))
	}
	return
}

func validateMaxRetries(v interface{}, k string) (ws []string, errors []error) {
	maxRetries := v.(int)
	if maxRetries < 1 || maxRetries > 10 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 10",
			k))
	}
	return
}

func validateTimeout(v interface{}, k string) (ws []string, errors []error) {
	timeout := v.(int)
	if timeout < 1 || timeout > 59 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 59",
			k))
	}
	return
}

func validateURLPath(v interface{}, k string) (ws []string, errors []error) {
	urlPath := v.(string)
	if len(urlPath) > 250 || !strings.HasPrefix(urlPath, "/") {
		errors = append(errors, fmt.Errorf(
			"%q should start with ‘/‘ and has a max length of 250 characters.",
			k))
	}
	return
}

func validateSecurityRuleDirection(v interface{}, k string) (ws []string, errors []error) {
	validDirections := map[string]bool{
		"ingress": true,
		"egress":  true,
	}

	value := v.(string)
	_, found := validDirections[value]
	if !found {
		strarray := make([]string, 0, len(validDirections))
		for key := range validDirections {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid security group rule direction %q. Valid types are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateSecurityRuleEtherType(v interface{}, k string) (ws []string, errors []error) {
	validEtherTypes := map[string]bool{
		"IPv4": true,
		"IPv6": true,
	}

	value := v.(string)
	_, found := validEtherTypes[value]
	if !found {
		strarray := make([]string, 0, len(validEtherTypes))
		for key := range validEtherTypes {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid security group rule ethernet type %q. Valid types are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

//validateIP...
func validateIP(v interface{}, k string) (ws []string, errors []error) {
	address := v.(string)
	if net.ParseIP(address) == nil {
		errors = append(errors, fmt.Errorf(
			"%q must be a valid ip address",
			k))
	}
	return
}

//validateCIDR...
func validateCIDR(v interface{}, k string) (ws []string, errors []error) {
	address := v.(string)
	_, _, err := net.ParseCIDR(address)
	if err != nil {
		errors = append(errors, fmt.Errorf(
			"%q must be a valid cidr address",
			k))
	}
	return
}

//validateRemoteIP...
func validateRemoteIP(v interface{}, k string) (ws []string, errors []error) {
	_, err1 := validateCIDR(v, k)
	_, err2 := validateIP(v, k)

	if len(err1) != 0 && len(err2) != 0 {
		errors = append(errors, fmt.Errorf(
			"%q must be a valid remote ip address (cidr or ip)",
			k))
	}
	return
}

func validateSecurityRuleProtocol(v interface{}, k string) (ws []string, errors []error) {
	validProtocols := map[string]bool{
		"icmp": true,
		"tcp":  true,
		"udp":  true,
	}

	value := v.(string)
	_, found := validProtocols[value]
	if !found {
		strarray := make([]string, 0, len(validProtocols))
		for key := range validProtocols {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid security group rule ethernet type %q. Valid types are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateNamespace(ns string) error {
	os := strings.Split(ns, "_")
	if len(os) < 2 || (len(os) == 2 && (len(os[0]) == 0 || len(os[1]) == 0)) {
		return fmt.Errorf(
			"Namespace is (%s), it must be of the form <org>_<space>, provider can't find the auth key if you use _ as well", ns)
	}
	return nil
}

func validateJSONString(v interface{}, k string) (ws []string, errors []error) {
	if _, err := normalizeJSONString(v); err != nil {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON: %s", k, err))
	}
	if err := validateKeyValue(v); err != nil {
		errors = append(errors, fmt.Errorf("%q contains an invalid JSON: %s", k, err))
	}
	return
}

func validateKeyValue(jsonString interface{}) error {
	var j [](map[string]interface{})
	if jsonString == nil || jsonString.(string) == "" {
		return nil
	}
	s := jsonString.(string)
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return err
	}
	for _, v := range j {
		_, exists := v["key"]
		if !exists {
			return errors.New("'key' is missing from json")
		}
		_, exists = v["value"]
		if !exists {
			return errors.New("'value' is missing from json")
		}
	}
	return nil
}

func validateActionName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if strings.HasPrefix(value, "/") {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must not start with a forward slash '/'.The action name should be like 'myaction' or utils/cloudant'", k, value))

	}

	const alphaNumeric = "abcdefghijklmnopqrstuvwxyz0123456789/_@.-"

	for _, char := range value {
		if !strings.Contains(alphaNumeric, strings.ToLower(string(char))) {
			errors = append(errors, fmt.Errorf(
				"%q (%q) The name of the package contains illegal characters", k, value))
		}
	}

	return
}

func validateActionKind(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	kindList := []string{"php:7.1", "nodejs:8", "swift:3", "nodejs", "blackbox", "java", "sequence", "nodejs:6", "python:3", "python", "python:2", "swift", "swift:3.1.1"}
	if !stringInSlice(value, kindList) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) Invalid kind is provided.Supported list of kinds of actions are (%q)", k, value, kindList))
	}
	return
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func validateFunctionName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	var validName = regexp.MustCompile(`\A([\w]|[\w][\w@ .-]*[\w@.-]+)\z`)
	if !validName.MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) The name contains illegal characters", k, value))

	}
	return
}

func validateBindedPackageName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	if !(strings.HasPrefix(value, "/")) {
		errors = append(errors, fmt.Errorf(
			"%q (%q) must start with a forward slash '/'.The package name should be '/whisk.system/cloudant', '/test@in.ibm.com_new/utils' or '/_/utils'", k, value))

	}

	index := strings.LastIndex(value, "/")

	if index < 2 || index == len(value)-1 {
		errors = append(errors, fmt.Errorf(
			"%q (%q) is not a valid bind package name.The package name should be '/whisk.system/cloudant','/test@in.ibm.com_new/utils' or '/_/utils'", k, value))

	}

	return
}

func validateStorageType(v interface{}, k string) (ws []string, errors []error) {
	validEtherTypes := map[string]bool{
		"Endurance":   true,
		"Performance": true,
	}

	value := v.(string)
	_, found := validEtherTypes[value]
	if !found {
		strarray := make([]string, 0, len(validEtherTypes))
		for key := range validEtherTypes {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid storage type %q. Valid types are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateRole(v interface{}, k string) (ws []string, errors []error) {
	validRolesTypes := map[string]bool{
		"Writer":        true,
		"Reader":        true,
		"Manager":       true,
		"Administrator": true,
		"Operator":      true,
		"Viewer":        true,
		"Editor":        true,
	}

	value := v.(string)
	_, found := validRolesTypes[value]
	if !found {
		strarray := make([]string, 0, len(validRolesTypes))
		for key := range validRolesTypes {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid role %q. Valid roles are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateDayOfWeek(v interface{}, k string) (ws []string, errors []error) {
	validDayTypes := map[string]bool{
		"SUNDAY":    true,
		"MONDAY":    true,
		"TUESDAY":   true,
		"WEDNESDAY": true,
		"THURSDAY":  true,
		"FRIDAY":    true,
		"SATURDAY":  true,
	}

	value := v.(string)
	_, found := validDayTypes[value]
	if !found {
		strarray := make([]string, 0, len(validDayTypes))
		for key := range validDayTypes {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid day %q. Valid days are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateScheduleType(v interface{}, k string) (ws []string, errors []error) {
	validSchdTypes := map[string]bool{
		"HOURLY": true,
		"DAILY":  true,
		"WEEKLY": true,
	}

	value := v.(string)
	_, found := validSchdTypes[value]
	if !found {
		strarray := make([]string, 0, len(validSchdTypes))
		for key := range validSchdTypes {
			strarray = append(strarray, key)
		}
		errors = append(errors, fmt.Errorf(
			"%q contains an invalid schedule type %q. Valid schedules are %q.",
			k, value, strings.Join(strarray, ",")))
	}
	return
}

func validateHour(start, end int) func(v interface{}, k string) (ws []string, errors []error) {
	f := func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if (value < start) || (value > end) {
			errors = append(errors, fmt.Errorf(
				"%q (%d) must be in the range of %d to %d", k, value, start, end))
		}
		return
	}
	return f
}

func validateMinute(start, end int) func(v interface{}, k string) (ws []string, errors []error) {
	f := func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if (value < start) || (value > end) {
			errors = append(errors, fmt.Errorf(
				"%q (%d) must be in the range of %d to %d", k, value, start, end))
		}
		return
	}
	return f
}

func validateDatacenterOption(v []interface{}, allowedValues []string) error {
	for _, option := range v {
		if option == nil {
			return fmt.Errorf("Provide a valid `datacenter_choice`")
		}
		values := option.(map[string]interface{})
		for k := range values {
			if !stringInSlice(k, allowedValues) {
				return fmt.Errorf(
					"%q Invalid values are provided in `datacenter_choice`. Supported list of keys are (%q)", k, allowedValues)
			}

		}
	}
	return nil
}

func validateLBTimeout(v interface{}, k string) (ws []string, errors []error) {
	timeout := v.(int)
	if timeout <= 0 || timeout > 3600 {
		errors = append(errors, fmt.Errorf(
			"%q must be between 1 and 3600",
			k))
	}
	return
}

// validateRecordType ensures that the dns record type is valid
func validateRecordType(t string, proxied bool) error {
	switch t {
	case "A", "AAAA", "CNAME":
		return nil
	case "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI":
		if !proxied {
			return nil
		}
	default:
		return fmt.Errorf(
			`Invalid type %q. Valid types are "A", "AAAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CAA", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA" or "URI".`, t)
	}

	return fmt.Errorf("Type %q cannot be proxied", t)
}

// validateRecordName ensures that based on supplied record type, the name content matches
// Currently only validates A and AAAA types
func validateRecordName(t string, value string) error {
	switch t {
	case "A":
		// Must be ipv4 addr
		addr := net.ParseIP(value)
		if addr == nil || !strings.Contains(value, ".") {
			return fmt.Errorf("A record must be a valid IPv4 address, got: %q", value)
		}
	case "AAAA":
		// Must be ipv6 addr
		addr := net.ParseIP(value)
		if addr == nil || !strings.Contains(value, ":") {
			return fmt.Errorf("AAAA record must be a valid IPv6 address, got: %q", value)
		}
	case "TXT":
		// Must be printable ASCII
		for i := 0; i < len(value); i++ {
			char := value[i]
			if (char < 0x20) || (0x7F < char) {
				return fmt.Errorf("TXT record must contain printable ASCII, found: %q", char)
			}
		}
	}

	return nil
}
