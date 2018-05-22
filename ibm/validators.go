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

func validateAuthProtocol(v interface{}, k string) (ws []string, errors []error) {
	authProtocol := v.(string)
	if authProtocol != "MD5" && authProtocol != "SHA1" && authProtocol != "SHA256" {
		errors = append(errors, fmt.Errorf(
			"%q auth protocol can be MD5 or SHA1 or SHA256", k))
	}
	return
}

func validateList(v interface{}, k string) (ws []string, errors []error) {
	authProtocol := v.([]interface{})
	if len(authProtocol) > 1 {
		errors = append(errors, fmt.Errorf(
			"%q Members of list can be only one", k))
	}
	return
}

func validateEncyptionProtocol(v interface{}, k string) (ws []string, errors []error) {
	encyptionProtocol := v.(string)
	if encyptionProtocol != "DES" && encyptionProtocol != "3DES" && encyptionProtocol != "AES128" && encyptionProtocol != "AES192" && encyptionProtocol != "AES256" {
		errors = append(errors, fmt.Errorf(
			"%q encryption protocol can be DES or 3DES or AES128 or AES192 or AES256", k))
	}
	return
}

func validateDiffieHellmanGroup(v interface{}, k string) (ws []string, errors []error) {
	diffieHellmanGroup := v.(int)
	if diffieHellmanGroup != 0 && diffieHellmanGroup != 1 && diffieHellmanGroup != 2 && diffieHellmanGroup != 5 {
		errors = append(errors, fmt.Errorf(
			"%q Diffie Hellman Group can be 0 or 1 or 2 or 5", k))
	}
	return
}

func validatekeylife(v interface{}, k string) (ws []string, errors []error) {
	keylife := v.(int)
	if keylife < 120 || keylife > 172800 {
		errors = append(errors, fmt.Errorf(
			"%q keylife value can be between 120 and 172800", k))
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
		"NAS/FTP":     true,
		"Portable":    true,
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
