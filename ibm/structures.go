package ibm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform/flatmap"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/apache/incubator-openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/sl"
)

//HashInt ...
func HashInt(v interface{}) int { return v.(int) }

func expandStringList(input []interface{}) []string {
	vs := make([]string, len(input))
	for i, v := range input {
		vs[i] = v.(string)
	}
	return vs
}

func flattenStringList(list []string) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func expandIntList(input []interface{}) []int {
	vs := make([]int, len(input))
	for i, v := range input {
		vs[i] = v.(int)
	}
	return vs
}

func flattenIntList(list []int) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func newStringSet(f schema.SchemaSetFunc, in []string) *schema.Set {
	var out = make([]interface{}, len(in), len(in))
	for i, v := range in {
		out[i] = v
	}
	return schema.NewSet(f, out)
}

func flattenRoute(in []mccpv2.Route) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v.GUID
	}
	return newStringSet(schema.HashString, vs)
}

func stringSliceToSet(in []string) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v
	}
	return newStringSet(schema.HashString, vs)
}

func flattenServiceBindings(in []mccpv2.ServiceBinding) *schema.Set {
	vs := make([]string, len(in))
	for i, v := range in {
		vs[i] = v.ServiceInstanceGUID
	}
	return newStringSet(schema.HashString, vs)
}

func flattenPort(in []int) *schema.Set {
	var out = make([]interface{}, len(in))
	for i, v := range in {
		out[i] = v
	}
	return schema.NewSet(HashInt, out)
}

func flattenFileStorageID(in []datatypes.Network_Storage) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		if *v.NasType == "NAS" {
			out = append(out, *v.Id)
		}
	}
	return schema.NewSet(HashInt, out)
}

func flattenBlockStorageID(in []datatypes.Network_Storage) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		if *v.NasType == "ISCSI" {
			out = append(out, *v.Id)
		}
	}
	return schema.NewSet(HashInt, out)
}

func flattenSSHKeyIDs(in []datatypes.Security_Ssh_Key) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		out = append(out, *v.Id)
	}
	return schema.NewSet(HashInt, out)
}

func flattenSpaceRoleUsers(in []mccpv2.SpaceRole) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		out = append(out, v.UserName)
	}
	return schema.NewSet(schema.HashString, out)
}

func flattenOrgRole(in []mccpv2.OrgRole, excludeUsername string) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		if excludeUsername == "" {
			out = append(out, v.UserName)
		} else {
			if v.UserName != excludeUsername {
				out = append(out, v.UserName)
			}
		}
	}
	return schema.NewSet(schema.HashString, out)
}

func flattenMapInterfaceVal(m map[string]interface{}) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		out[k] = fmt.Sprintf("%v", v)
	}
	return out
}

func flattenCredentials(creds map[string]interface{}) map[string]string {
	return flattenMapInterfaceVal(creds)
}

func flattenServiceKeyCredentials(creds map[string]interface{}) map[string]string {
	return flattenCredentials(creds)
}

func flattenServiceInstanceCredentials(keys []mccpv2.ServiceKeyFields) []interface{} {
	var out = make([]interface{}, len(keys), len(keys))
	for i, k := range keys {
		m := make(map[string]interface{})
		m["name"] = k.Entity.Name
		m["credentials"] = flatmap.Flatten(k.Entity.Credentials)
		out[i] = m
	}
	return out
}

func flattenIAMPolicyResource(list []iampapv1.Resources, iamClient iampapv1.IAMPAPAPI) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		name := i.ServiceName
		if name == "" {
			name = allIAMEnabledServices
		}
		serviceName, err := iamClient.IAMService().GetServiceDispalyName(name)
		if err != nil {
			return result, fmt.Errorf("Error retrieving service : %s", err)
		}
		l := map[string]interface{}{
			"service_name":      serviceName,
			"region":            i.Region,
			"resource_type":     i.ResourceType,
			"resource":          i.Resource,
			"space_guid":        i.SpaceId,
			"organization_guid": i.OrganizationId,
		}
		if i.ServiceInstance != "" {
			l["service_instance"] = []string{i.ServiceInstance}
		}
		result = append(result, l)
	}
	return result, nil
}

func flattenIAMPolicyRoles(list []iampapv1.Roles) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, v := range list {
		l := map[string]interface{}{
			"name": roleIDToName[v.ID],
		}
		result = append(result, l)
	}
	return result
}

func expandProtocols(configured []interface{}) ([]datatypes.Network_LBaaS_LoadBalancerProtocolConfiguration, error) {
	protocols := make([]datatypes.Network_LBaaS_LoadBalancerProtocolConfiguration, 0, len(configured))
	for _, lRaw := range configured {
		data := lRaw.(map[string]interface{})
		p := &datatypes.Network_LBaaS_LoadBalancerProtocolConfiguration{
			FrontendProtocol: sl.String(data["frontend_protocol"].(string)),
			BackendProtocol:  sl.String(data["backend_protocol"].(string)),
			FrontendPort:     sl.Int(data["frontend_port"].(int)),
			BackendPort:      sl.Int(data["backend_port"].(int)),
		}
		if v, ok := data["session_stickiness"]; ok && v.(string) != "" {
			p.SessionType = sl.String(v.(string))
		}
		if v, ok := data["max_conn"]; ok && v.(int) != 0 {
			p.MaxConn = sl.Int(v.(int))
		}
		if v, ok := data["tls_certificate_id"]; ok && v.(int) != 0 {
			p.TlsCertificateId = sl.Int(v.(int))
		}
		if v, ok := data["load_balancing_method"]; ok {
			p.LoadBalancingMethod = sl.String(lbMethodToId[v.(string)])
		}
		if v, ok := data["protocol_id"]; ok && v.(string) != "" {
			p.ListenerUuid = sl.String(v.(string))
		}

		var isValid bool
		if p.TlsCertificateId != nil && *p.TlsCertificateId != 0 {
			// validate the protocol is correct
			if *p.FrontendProtocol == "HTTPS" {
				isValid = true
			}
		} else {
			isValid = true
		}

		if isValid {
			protocols = append(protocols, *p)
		} else {
			return protocols, fmt.Errorf("tls_certificate_id may be set only when frontend protocol is 'HTTPS'")
		}

	}
	return protocols, nil
}

func expandMembers(configured []interface{}) []datatypes.Network_LBaaS_LoadBalancerServerInstanceInfo {
	members := make([]datatypes.Network_LBaaS_LoadBalancerServerInstanceInfo, 0, len(configured))
	for _, lRaw := range configured {
		data := lRaw.(map[string]interface{})
		p := &datatypes.Network_LBaaS_LoadBalancerServerInstanceInfo{}
		if v, ok := data["private_ip_address"]; ok && v.(string) != "" {
			p.PrivateIpAddress = sl.String(v.(string))
		}
		if v, ok := data["weight"]; ok && v.(int) != 0 {
			p.Weight = sl.Int(v.(int))
		}

		members = append(members, *p)
	}
	return members
}

func flattenServerInstances(list []datatypes.Network_LBaaS_Member) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"private_ip_address": *i.Address,
			"member_id":          *i.Uuid,
		}
		if i.Weight != nil {
			l["weight"] = *i.Weight
		}
		result = append(result, l)
	}
	return result
}

func flattenProtocols(list []datatypes.Network_LBaaS_Listener) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"frontend_protocol":     *i.Protocol,
			"frontend_port":         *i.ProtocolPort,
			"backend_protocol":      *i.DefaultPool.Protocol,
			"backend_port":          *i.DefaultPool.ProtocolPort,
			"load_balancing_method": lbIdToMethod[*i.DefaultPool.LoadBalancingAlgorithm],
			"protocol_id":           *i.Uuid,
		}
		if i.DefaultPool.SessionAffinity != nil && i.DefaultPool.SessionAffinity.Type != nil && *i.DefaultPool.SessionAffinity.Type != "" {
			l["session_stickiness"] = *i.DefaultPool.SessionAffinity.Type
		}
		if i.ConnectionLimit != nil && *i.ConnectionLimit != 0 {
			l["max_conn"] = *i.ConnectionLimit
		}
		if i.TlsCertificateId != nil && *i.TlsCertificateId != 0 {
			l["tls_certificate_id"] = *i.TlsCertificateId
		}
		result = append(result, l)
	}
	return result
}

func flattenVlans(list []containerv1.Vlan) []map[string]interface{} {
	vlans := make([]map[string]interface{}, len(list))
	for i, vlanR := range list {
		subnets := make([]map[string]interface{}, len(vlanR.Subnets))
		for j, subnetR := range vlanR.Subnets {
			subnet := make(map[string]interface{})
			subnet["id"] = subnetR.ID
			subnet["cidr"] = subnetR.Cidr
			subnet["is_byoip"] = subnetR.IsByOIP
			subnet["is_public"] = subnetR.IsPublic
			ips := make([]string, len(subnetR.Ips))
			for k, ip := range subnetR.Ips {
				ips[k] = ip
			}
			subnet["ips"] = ips
			subnets[j] = subnet
		}
		l := map[string]interface{}{
			"id":      vlanR.ID,
			"subnets": subnets,
		}
		vlans[i] = l
	}
	return vlans
}

func normalizeJSONString(jsonString interface{}) (string, error) {
	var j interface{}
	if jsonString == nil || jsonString.(string) == "" {
		return "", nil
	}
	s := jsonString.(string)
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return s, err
	}
	bytes, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

func expandAnnotations(annotations string) (whisk.KeyValueArr, error) {
	var result whisk.KeyValueArr
	dc := json.NewDecoder(strings.NewReader(annotations))
	dc.UseNumber()
	err := dc.Decode(&result)
	return result, err
}

func flattenAnnotations(in whisk.KeyValueArr) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func expandParameters(annotations string) (whisk.KeyValueArr, error) {
	var result whisk.KeyValueArr
	dc := json.NewDecoder(strings.NewReader(annotations))
	dc.UseNumber()
	err := dc.Decode(&result)
	return result, err
}

func flattenParameters(in whisk.KeyValueArr) (string, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}

func expandLimits(l []interface{}) *whisk.Limits {
	if len(l) == 0 || l[0] == nil {
		return &whisk.Limits{}
	}
	in := l[0].(map[string]interface{})
	obj := &whisk.Limits{
		Timeout: ptrToInt(in["timeout"].(int)),
		Memory:  ptrToInt(in["memory"].(int)),
		Logsize: ptrToInt(in["log_size"].(int)),
	}
	return obj
}

func flattenLimits(in *whisk.Limits) []interface{} {
	att := make(map[string]interface{})
	if in.Timeout != nil {
		att["timeout"] = *in.Timeout
	}
	if in.Memory != nil {
		att["memory"] = *in.Memory
	}
	if in.Memory != nil {
		att["log_size"] = *in.Logsize
	}
	return []interface{}{att}
}

func expandExec(execs []interface{}) *whisk.Exec {
	for _, exec := range execs {
		e, _ := exec.(map[string]interface{})
		obj := &whisk.Exec{
			Image:      e["image"].(string),
			Init:       e["init"].(string),
			Code:       ptrToString(e["code"].(string)),
			Kind:       e["kind"].(string),
			Main:       e["main"].(string),
			Components: expandStringList(e["components"].([]interface{})),
		}
		return obj
	}

	return &whisk.Exec{}
}

func flattenExec(in *whisk.Exec) []interface{} {
	att := make(map[string]interface{})
	if in.Image != "" {
		att["image"] = in.Image
	}
	if in.Init != "" {
		att["init"] = in.Init
	}
	if in.Code != nil {
		att["code"] = *in.Code
	}
	if in.Kind != "" {
		att["kind"] = in.Kind
	}
	if in.Main != "" {
		att["main"] = in.Main
	}

	if len(in.Components) > 0 {
		att["components"] = flattenStringList(in.Components)
	}

	return []interface{}{att}
}

func ptrToInt(i int) *int {
	return &i
}

func ptrToString(s string) *string {
	return &s
}

func filterActionAnnotations(in whisk.KeyValueArr) (string, error) {
	noExec := make(whisk.KeyValueArr, 0, len(in))
	for _, v := range in {
		if v.Key == "exec" {
			continue
		}
		noExec = append(noExec, v)
	}

	return flattenAnnotations(noExec)
}

func filterActionParameters(in whisk.KeyValueArr) (string, error) {
	noAction := make(whisk.KeyValueArr, 0, len(in))
	for _, v := range in {
		if v.Key == "_actions" {
			continue
		}
		noAction = append(noAction, v)
	}
	return flattenParameters(noAction)
}

func filterInheritedAnnotations(inheritedAnnotations, annotations whisk.KeyValueArr) whisk.KeyValueArr {
	userDefinedAnnotations := make(whisk.KeyValueArr, 0)
	for _, a := range annotations {
		insert := false
		if a.Key == "binding" || a.Key == "exec" {
			insert = false
			break
		}
		for _, b := range inheritedAnnotations {
			if a.Key == b.Key && reflect.DeepEqual(a.Value, b.Value) {
				insert = false
				break
			}
			insert = true
		}
		if insert {
			userDefinedAnnotations = append(userDefinedAnnotations, a)
		}
	}
	return userDefinedAnnotations
}

func filterInheritedParameters(inheritedParameters, parameters whisk.KeyValueArr) whisk.KeyValueArr {
	userDefinedParameters := make(whisk.KeyValueArr, 0)
	for _, p := range parameters {
		insert := false
		if p.Key == "_actions" {
			insert = false
			break
		}
		for _, b := range inheritedParameters {
			if p.Key == b.Key && reflect.DeepEqual(p.Value, b.Value) {
				insert = false
				break
			}
			insert = true
		}
		if insert {
			userDefinedParameters = append(userDefinedParameters, p)
		}

	}
	return userDefinedParameters
}

func isEmpty(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}

func filterTriggerAnnotations(in whisk.KeyValueArr) (string, error) {
	noFeed := make(whisk.KeyValueArr, 0, len(in))
	for _, v := range in {
		if v.Key == "feed" {
			continue
		}
		noFeed = append(noFeed, v)
	}
	return flattenParameters(noFeed)
}

func flattenFeed(feedName string) []interface{} {
	att := make(map[string]interface{})
	att["name"] = feedName
	att["parameters"] = "[]"
	return []interface{}{att}
}

func flattenGatewayVlans(list []datatypes.Network_Gateway_Vlan) []map[string]interface{} {
	vlans := make([]map[string]interface{}, len(list))
	for i, ele := range list {
		vlan := make(map[string]interface{})
		vlan["bypass"] = *ele.BypassFlag
		vlan["network_vlan_id"] = *ele.NetworkVlanId
		vlan["vlan_id"] = *ele.Id
		vlans[i] = vlan
	}
	return vlans
}

func flattenGatewayMembers(d *schema.ResourceData, list []datatypes.Network_Gateway_Member) []map[string]interface{} {
	members := make([]map[string]interface{}, len(list))
	for i, ele := range list {
		hardware := *ele.Hardware
		member := make(map[string]interface{})
		member["member_id"] = *ele.HardwareId
		member["hostname"] = *hardware.Hostname
		member["domain"] = *hardware.Domain
		member["priority"] = *ele.Priority
		if hardware.OperatingSystem != nil && hardware.OperatingSystem.Passwords != nil {
			passwords := make([]map[string]string, len(hardware.OperatingSystem.Passwords))
			OperatingSystem := *hardware.OperatingSystem
			for index := range OperatingSystem.Passwords {
				creds := make(map[string]string)
				creds["username"] = *OperatingSystem.Passwords[index].Username
				creds["password"] = *OperatingSystem.Passwords[index].Password
				passwords[index] = creds
			}
			member["passwords"] = passwords
		}
		if hardware.Notes != nil {
			member["notes"] = *hardware.Notes
		}
		if hardware.Datacenter != nil {
			member["datacenter"] = *hardware.Datacenter.Name
		}
		if hardware.PrimaryNetworkComponent.MaxSpeed != nil {
			member["network_speed"] = *hardware.PrimaryNetworkComponent.MaxSpeed
		}
		member["redundant_network"] = false
		member["unbonded_network"] = false
		backendNetworkComponent := ele.Hardware.BackendNetworkComponents

		if len(backendNetworkComponent) > 2 && ele.Hardware.PrimaryBackendNetworkComponent != nil {
			if *hardware.PrimaryBackendNetworkComponent.RedundancyEnabledFlag {
				member["redundant_network"] = true
			} else {
				member["unbonded_network"] = true
			}
		}
		tagReferences := ele.Hardware.TagReferences
		tagReferencesLen := len(tagReferences)
		if tagReferencesLen > 0 {
			tags := make([]interface{}, 0, tagReferencesLen)
			for _, tagRef := range tagReferences {
				tags = append(tags, *tagRef.Tag.Name)
			}
			member["tags"] = schema.NewSet(schema.HashString, tags)
		}

		member["redundant_power_supply"] = false

		if *hardware.PowerSupplyCount == 2 {
			member["redundant_power_supply"] = true
		}
		member["memory"] = *hardware.MemoryCapacity
		if !(*hardware.PrivateNetworkOnlyFlag) && hardware.NetworkVlans[1].Id != nil {
			member["public_vlan_id"] = *hardware.NetworkVlans[1].Id
		}
		if hardware.NetworkVlans[0].Id != nil {
			member["private_vlan_id"] = *hardware.NetworkVlans[0].Id
		}
		if hardware.PrimaryIpAddress != nil {
			member["public_ipv4_address"] = *hardware.PrimaryIpAddress
		}
		if hardware.PrimaryBackendIpAddress != nil {
			member["private_ipv4_address"] = *hardware.PrimaryBackendIpAddress
		}
		member["ipv6_enabled"] = false
		if ele.Hardware.PrimaryNetworkComponent.PrimaryVersion6IpAddressRecord != nil {
			member["ipv6_enabled"] = true
			member["ipv6_address"] = *hardware.PrimaryNetworkComponent.PrimaryVersion6IpAddressRecord.IpAddress
		}

		member["private_network_only"] = *hardware.PrivateNetworkOnlyFlag
		userData := hardware.UserData
		if len(userData) > 0 && userData[0].Value != nil {
			member["user_metadata"] = *userData[0].Value
		}
		members[i] = member
	}
	return members
}

func flattenDisks(result datatypes.Virtual_Guest, d *schema.ResourceData) []int {
	var out = make([]int, 0)

	for _, v := range result.BlockDevices {
		// skip 1,7 which is reserved for the swap disk and metadata
		if _, ok := d.GetOk("flavor_key_name"); ok {
			if *v.Device != "1" && *v.Device != "7" && *v.Device != "0" {
				out = append(out, *v.DiskImage.Capacity)
			}
		} else {
			if *v.Device != "1" && *v.Device != "7" {
				out = append(out, *v.DiskImage.Capacity)
			}
		}
	}

	return out
}

func filterResourceKeyParameters(params map[string]interface{}) map[string]interface{} {
	delete(params, "role_crn")
	return params
}
