package ibm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv2"
	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/schematics"
	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"
	"github.com/apache/openwhisk-client-go/whisk"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/sl"
)

const (
	prodBaseController  = "https://cloud.ibm.com"
	stageBaseController = "https://test.cloud.ibm.com"
	//ResourceControllerURL ...
	ResourceControllerURL = "resource_controller_url"
	//ResourceName ...
	ResourceName = "resource_name"
	//ResourceCRN ...
	ResourceCRN = "resource_crn"
	//ResourceStatus ...
	ResourceStatus = "resource_status"
	//ResourceGroupName ...
	ResourceGroupName = "resource_group_name"
	//RelatedCRN ...
	RelatedCRN = "related_crn"
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
		m["credentials"] = Flatten(k.Entity.Credentials)
		out[i] = m
	}
	return out
}

func flattenUsersSet(userList *schema.Set) []string {
	users := make([]string, 0)
	for _, user := range userList.List() {
		users = append(users, user.(string))
	}
	return users
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

func flattenVpcWorkerPools(list []containerv2.GetWorkerPoolResponse) []map[string]interface{} {
	workerPools := make([]map[string]interface{}, len(list))
	for i, workerPool := range list {
		l := map[string]interface{}{
			"id":           workerPool.ID,
			"name":         workerPool.PoolName,
			"flavor":       workerPool.Flavor,
			"worker_count": workerPool.WorkerCount,
			"isolation":    workerPool.Isolation,
			"labels":       workerPool.Labels,
			"state":        workerPool.Lifecycle.ActualState,
		}
		zones := workerPool.Zones
		zonesConfig := make([]map[string]interface{}, len(zones))
		for j, zone := range zones {
			z := map[string]interface{}{
				"zone":         zone.ID,
				"worker_count": zone.WorkerCount,
			}
			subnets := zone.Subnets
			subnetConfig := make([]map[string]interface{}, len(subnets))
			for k, subnet := range subnets {
				s := map[string]interface{}{
					"id":      subnet.ID,
					"primary": subnet.Primary,
				}
				subnetConfig[k] = s
			}
			z["subnets"] = subnetConfig
			zonesConfig[j] = z
		}
		l["zones"] = zonesConfig
		workerPools[i] = l
	}

	return workerPools
}

func flattenVpcZones(list []containerv2.ZoneResp) []map[string]interface{} {
	zones := make([]map[string]interface{}, len(list))
	for i, zone := range list {
		l := map[string]interface{}{
			"id":           zone.ID,
			"subnet_id":    flattenSubnets(zone.Subnets),
			"worker_count": zone.WorkerCount,
		}
		zones[i] = l
	}
	return zones
}
func flattenConditions(list []iamuumv2.Condition) []map[string]interface{} {
	conditions := make([]map[string]interface{}, len(list))
	for i, cond := range list {
		l := map[string]interface{}{
			"claim":    cond.Claim,
			"operator": cond.Operator,
			"value":    strings.ReplaceAll(cond.Value, "\"", ""),
		}
		conditions[i] = l
	}
	return conditions
}
func flattenAccessGroupRules(list []iamuumv2.CreateRuleResponse) []map[string]interface{} {
	rules := make([]map[string]interface{}, len(list))
	for i, item := range list {
		l := map[string]interface{}{
			"name":              item.Name,
			"expiration":        item.Expiration,
			"identity_provider": item.RealmName,
			"conditions":        flattenConditions(item.Conditions),
		}
		rules[i] = l
	}
	return rules
}

func flattenSubnets(list []containerv2.Subnet) []map[string]interface{} {
	subs := make([]map[string]interface{}, len(list))
	for i, sub := range list {
		l := map[string]interface{}{
			"id":           sub.ID,
			"worker_count": sub.Primary,
		}
		subs[i] = l
	}
	return subs
}

func flattenZones(list []containerv1.WorkerPoolZoneResponse) []map[string]interface{} {
	zones := make([]map[string]interface{}, len(list))
	for i, zone := range list {
		l := map[string]interface{}{
			"zone":         zone.WorkerPoolZone.ID,
			"private_vlan": zone.WorkerPoolZone.WorkerPoolZoneNetwork.PrivateVLAN,
			"public_vlan":  zone.WorkerPoolZone.WorkerPoolZoneNetwork.PublicVLAN,
			"worker_count": zone.WorkerCount,
		}
		zones[i] = l
	}
	return zones
}

func flattenWorkerPools(list []containerv1.WorkerPoolResponse) []map[string]interface{} {
	workerPools := make([]map[string]interface{}, len(list))
	for i, workerPool := range list {
		l := map[string]interface{}{
			"id":            workerPool.ID,
			"hardware":      workerPool.Isolation,
			"name":          workerPool.Name,
			"machine_type":  workerPool.MachineType,
			"size_per_zone": workerPool.Size,
			"state":         workerPool.State,
			"labels":        workerPool.Labels,
		}
		zones := workerPool.Zones
		zonesConfig := make([]map[string]interface{}, len(zones))
		for j, zone := range zones {
			z := map[string]interface{}{
				"zone":         zone.ID,
				"private_vlan": zone.PrivateVLAN,
				"public_vlan":  zone.PublicVLAN,
				"worker_count": zone.WorkerCount,
			}
			zonesConfig[j] = z
		}
		l["zones"] = zonesConfig
		workerPools[i] = l
	}

	return workerPools
}

func flattenAlbs(list []containerv1.ALBConfig, filterType string) []map[string]interface{} {
	albs := make([]map[string]interface{}, 0)
	for _, alb := range list {
		if alb.ALBType == filterType || filterType == "all" {
			l := map[string]interface{}{
				"id":                 alb.ALBID,
				"name":               alb.Name,
				"alb_type":           alb.ALBType,
				"enable":             alb.Enable,
				"state":              alb.State,
				"num_of_instances":   alb.NumOfInstances,
				"alb_ip":             alb.ALBIP,
				"resize":             alb.Resize,
				"disable_deployment": alb.DisableDeployment,
			}
			albs = append(albs, l)
		}
	}
	return albs
}

func flattenVpcAlbs(list []containerv2.AlbConfig, filterType string) []map[string]interface{} {
	albs := make([]map[string]interface{}, 0)
	for _, alb := range list {
		if alb.AlbType == filterType || filterType == "all" {
			l := map[string]interface{}{
				"id":                     alb.AlbID,
				"name":                   alb.Name,
				"alb_type":               alb.AlbType,
				"enable":                 alb.Enable,
				"state":                  alb.State,
				"resize":                 alb.Resize,
				"disable_deployment":     alb.DisableDeployment,
				"load_balancer_hostname": alb.LoadBalancerHostname,
			}
			albs = append(albs, l)
		}
	}
	return albs
}

func flattenNetworkInterfaces(list []containerv2.Network) []map[string]interface{} {
	nwInterfaces := make([]map[string]interface{}, len(list))
	for i, nw := range list {
		l := map[string]interface{}{
			"cidr":       nw.Cidr,
			"ip_address": nw.IpAddress,
			"subnet_id":  nw.SubnetID,
		}
		nwInterfaces[i] = l
	}
	return nwInterfaces
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

func flattenIcdGroups(grouplist icdv4.GroupList) []map[string]interface{} {
	groups := make([]map[string]interface{}, len(grouplist.Groups))
	for i, group := range grouplist.Groups {
		memorys := make([]map[string]interface{}, 1)
		memory := make(map[string]interface{})
		memory["units"] = group.Memory.Units
		memory["allocation_mb"] = group.Memory.AllocationMb
		memory["minimum_mb"] = group.Memory.MinimumMb
		memory["step_size_mb"] = group.Memory.StepSizeMb
		memory["is_adjustable"] = group.Memory.IsAdjustable
		memory["can_scale_down"] = group.Memory.CanScaleDown
		memorys[0] = memory

		cpus := make([]map[string]interface{}, 1)
		cpu := make(map[string]interface{})
		cpu["units"] = group.Cpu.Units
		cpu["allocation_count"] = group.Cpu.AllocationCount
		cpu["minimum_count"] = group.Cpu.MinimumCount
		cpu["step_size_count"] = group.Cpu.StepSizeCount
		cpu["is_adjustable"] = group.Cpu.IsAdjustable
		cpu["can_scale_down"] = group.Cpu.CanScaleDown
		cpus[0] = cpu

		disks := make([]map[string]interface{}, 1)
		disk := make(map[string]interface{})
		disk["units"] = group.Disk.Units
		disk["allocation_mb"] = group.Disk.AllocationMb
		disk["minimum_mb"] = group.Disk.MinimumMb
		disk["step_size_mb"] = group.Disk.StepSizeMb
		disk["is_adjustable"] = group.Disk.IsAdjustable
		disk["can_scale_down"] = group.Disk.CanScaleDown
		disks[0] = disk

		l := map[string]interface{}{
			"group_id": group.Id,
			"count":    group.Count,
			"memory":   memorys,
			"cpu":      cpus,
			"disk":     disks,
		}
		groups[i] = l
	}
	return groups
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

func flattenActivityTrack(in *resourceconfigurationv1.ActivityTracking) []interface{} {

	att := make(map[string]interface{})
	if in != nil {
		if in.ReadDataEvents != nil {
			att["read_data_events"] = *in.ReadDataEvents
		}
		if in.WriteDataEvents != nil {
			att["write_data_events"] = *in.WriteDataEvents
		}
		if in.ActivityTrackerCrn != nil {
			att["activity_tracker_crn"] = *in.ActivityTrackerCrn
		}
	}
	return []interface{}{att}
}

func flattenMetricsMonitor(in *resourceconfigurationv1.MetricsMonitoring) []interface{} {
	att := make(map[string]interface{})
	if in != nil {
		if in.UsageMetricsEnabled != nil {
			att["usage_metrics_enabled"] = *in.UsageMetricsEnabled
		}
		if in.MetricsMonitoringCrn != nil {
			att["metrics_monitoring_crn"] = *in.MetricsMonitoringCrn
		}
	}
	return []interface{}{att}
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
		if !(*hardware.PrivateNetworkOnlyFlag) {
			member["public_vlan_id"] = *hardware.NetworkVlans[1].Id
		}
		member["private_vlan_id"] = *hardware.NetworkVlans[0].Id

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

func flattenDisks(result datatypes.Virtual_Guest) []int {
	var out = make([]int, 0)

	for _, v := range result.BlockDevices {
		// skip 1,7 which is reserved for the swap disk and metadata
		_, ok := sl.GrabOk(result, "BillingItem.OrderItem.Preset")
		if ok {
			if *v.Device != "1" && *v.Device != "7" && *v.Device != "0" {
				capacity, ok := sl.GrabOk(v, "DiskImage.Capacity")

				if ok {
					out = append(out, capacity.(int))
				}

			}
		} else {
			if *v.Device != "1" && *v.Device != "7" {
				capacity, ok := sl.GrabOk(v, "DiskImage.Capacity")

				if ok {
					out = append(out, capacity.(int))
				}
			}
		}
	}

	return out
}

func flattenDisksForWindows(result datatypes.Virtual_Guest) []int {
	var out = make([]int, 0)

	for _, v := range result.BlockDevices {
		// skip 1,7 which is reserved for the swap disk and metadata
		_, ok := sl.GrabOk(result, "BillingItem.OrderItem.Preset")
		if ok {
			if *v.Device != "1" && *v.Device != "7" && *v.Device != "0" && *v.Device != "3" {
				capacity, ok := sl.GrabOk(v, "DiskImage.Capacity")

				if ok {
					out = append(out, capacity.(int))
				}
			}
		} else {
			if *v.Device != "1" && *v.Device != "7" && *v.Device != "3" {
				capacity, ok := sl.GrabOk(v, "DiskImage.Capacity")

				if ok {
					out = append(out, capacity.(int))
				}
			}
		}
	}

	return out
}

func filterResourceKeyParameters(params map[string]interface{}) map[string]interface{} {
	delete(params, "role_crn")
	return params
}

func idParts(id string) ([]string, error) {
	if strings.Contains(id, "/") {
		parts := strings.Split(id, "/")
		return parts, nil
	}
	return []string{}, fmt.Errorf("The given id %s does not contain / please check documentation on how to provider id during import command", id)
}

func vmIdParts(id string) ([]string, error) {
	parts := strings.Split(id, "/")
	return parts, nil
}

func cfIdParts(id string) ([]string, error) {
	parts := strings.Split(id, ":")
	return parts, nil
}

func flattenPolicyResource(list []iampapv1.Resource) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, i := range list {
		l := map[string]interface{}{
			"service":              i.GetAttribute("serviceName"),
			"resource_instance_id": i.GetAttribute("serviceInstance"),
			"region":               i.GetAttribute("region"),
			"resource_type":        i.GetAttribute("resourceType"),
			"resource":             i.GetAttribute("resource"),
			"resource_group_id":    i.GetAttribute("resourceGroupId"),
		}
		customAttributes := i.CustomAttributes()
		if len(customAttributes) > 0 {
			out := make(map[string]string)
			for _, a := range customAttributes {
				out[a.Name] = a.Value
			}
			l["attributes"] = out
		}

		result = append(result, l)
	}
	return result
}

// Cloud Internet Services
func flattenHealthMonitors(list []datatypes.Network_LBaaS_Listener) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	ports := make([]int, 0, 0)
	for _, i := range list {
		l := map[string]interface{}{
			"protocol":    *i.DefaultPool.Protocol,
			"port":        *i.DefaultPool.ProtocolPort,
			"interval":    *i.DefaultPool.HealthMonitor.Interval,
			"max_retries": *i.DefaultPool.HealthMonitor.MaxRetries,
			"timeout":     *i.DefaultPool.HealthMonitor.Timeout,
			"monitor_id":  *i.DefaultPool.HealthMonitor.Uuid,
		}

		if i.DefaultPool.HealthMonitor.UrlPath != nil {
			l["url_path"] = *i.DefaultPool.HealthMonitor.UrlPath
		}

		if !contains(ports, *i.DefaultPool.ProtocolPort) {
			result = append(result, l)
		}

		ports = append(ports, *i.DefaultPool.ProtocolPort)
	}
	return result
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func flattenMembersData(list []models.AccessGroupMemberV2, users []usermanagementv2.UserInfo, serviceids []models.ServiceID) ([]string, []string) {
	var ibmid []string
	var serviceid []string
	for _, m := range list {
		if m.Type == iamuumv2.AccessGroupMemberUser {
			for _, user := range users {
				if user.IamID == m.ID {
					ibmid = append(ibmid, user.Email)
					break
				}
			}
		} else {

			for _, srid := range serviceids {
				if srid.IAMID == m.ID {
					serviceid = append(serviceid, srid.UUID)
					break
				}
			}

		}

	}
	return ibmid, serviceid
}

func flattenAccessGroupMembers(list []models.AccessGroupMemberV2, users []usermanagementv2.UserInfo, serviceids []models.ServiceID) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(list))
	for _, m := range list {
		var value, vtype string
		if m.Type == iamuumv2.AccessGroupMemberUser {
			vtype = iamuumv2.AccessGroupMemberUser
			for _, user := range users {
				if user.IamID == m.ID {
					value = user.Email
					break
				}
			}
		} else {

			vtype = iamuumv1.AccessGroupMemberService
			for _, srid := range serviceids {
				if srid.IAMID == m.ID {
					value = srid.UUID
					break
				}
			}

		}
		l := map[string]interface{}{
			"iam_id": value,
			"type":   vtype,
		}
		result = append(result, l)
	}
	return result
}

func flattenUserIds(accountID string, users []string, meta interface{}) ([]string, error) {
	userids := make([]string, len(users))
	for i, name := range users {
		iamID, err := getIBMUniqueId(accountID, name, meta)
		if err != nil {
			return nil, err
		}
		userids[i] = iamID
	}
	return userids, nil
}

func flattenServiceIds(services []string, meta interface{}) ([]string, error) {
	serviceids := make([]string, len(services))
	for i, id := range services {
		serviceID, err := getServiceID(id, meta)
		if err != nil {
			return nil, err
		}
		serviceids[i] = serviceID.IAMID
	}
	return serviceids, nil
}

func expandUsers(userList *schema.Set) (users []icdv4.User) {
	for _, iface := range userList.List() {
		userEl := iface.(map[string]interface{})
		user := icdv4.User{
			UserName: userEl["name"].(string),
			Password: userEl["password"].(string),
		}
		users = append(users, user)
	}
	return
}

// IBM Cloud Databases
func flattenConnectionStrings(cs []CsEntry) []map[string]interface{} {
	entries := make([]map[string]interface{}, len(cs), len(cs))
	for i, csEntry := range cs {
		l := map[string]interface{}{
			"name":         csEntry.Name,
			"password":     csEntry.Password,
			"composed":     csEntry.Composed,
			"certname":     csEntry.CertName,
			"certbase64":   csEntry.CertBase64,
			"queryoptions": csEntry.QueryOptions,
			"scheme":       csEntry.Scheme,
			"path":         csEntry.Path,
			"database":     csEntry.Database,
		}
		hosts := csEntry.Hosts
		hostsList := make([]map[string]interface{}, len(hosts), len(hosts))
		for j, host := range hosts {
			z := map[string]interface{}{
				"hostname": host.HostName,
				"port":     strconv.Itoa(host.Port),
			}
			hostsList[j] = z
		}
		l["hosts"] = hostsList
		var queryOpts string
		if len(csEntry.QueryOptions) != 0 {
			queryOpts = "?"
			count := 0
			for k, v := range csEntry.QueryOptions {
				if count >= 1 {
					queryOpts = queryOpts + "&"
				}
				queryOpts = queryOpts + fmt.Sprintf("%v", k) + "=" + fmt.Sprintf("%v", v)
				count++
			}
		} else {
			queryOpts = ""
		}
		l["queryoptions"] = queryOpts
		entries[i] = l
	}

	return entries
}

func flattenPhaseOneAttributes(vpn *datatypes.Network_Tunnel_Module_Context) []map[string]interface{} {
	phaseoneAttributesMap := make([]map[string]interface{}, 0, 1)
	phaseoneAttributes := make(map[string]interface{})
	phaseoneAttributes["authentication"] = *vpn.PhaseOneAuthentication
	phaseoneAttributes["encryption"] = *vpn.PhaseOneEncryption
	phaseoneAttributes["diffie_hellman_group"] = *vpn.PhaseOneDiffieHellmanGroup
	phaseoneAttributes["keylife"] = *vpn.PhaseOneKeylife
	phaseoneAttributesMap = append(phaseoneAttributesMap, phaseoneAttributes)
	return phaseoneAttributesMap
}

func flattenPhaseTwoAttributes(vpn *datatypes.Network_Tunnel_Module_Context) []map[string]interface{} {
	phasetwoAttributesMap := make([]map[string]interface{}, 0, 1)
	phasetwoAttributes := make(map[string]interface{})
	phasetwoAttributes["authentication"] = *vpn.PhaseTwoAuthentication
	phasetwoAttributes["encryption"] = *vpn.PhaseTwoEncryption
	phasetwoAttributes["diffie_hellman_group"] = *vpn.PhaseTwoDiffieHellmanGroup
	phasetwoAttributes["keylife"] = *vpn.PhaseTwoKeylife
	phasetwoAttributesMap = append(phasetwoAttributesMap, phasetwoAttributes)
	return phasetwoAttributesMap
}

func flattenaddressTranslation(vpn *datatypes.Network_Tunnel_Module_Context, fwID int) []map[string]interface{} {
	addressTranslationMap := make([]map[string]interface{}, 0, 1)
	addressTranslationAttributes := make(map[string]interface{})
	for _, networkAddressTranslation := range vpn.AddressTranslations {
		if *networkAddressTranslation.NetworkTunnelContext.Id == fwID {
			addressTranslationAttributes["remote_ip_adress"] = *networkAddressTranslation.CustomerIpAddress
			addressTranslationAttributes["internal_ip_adress"] = *networkAddressTranslation.InternalIpAddress
			addressTranslationAttributes["notes"] = *networkAddressTranslation.Notes
		}
	}
	addressTranslationMap = append(addressTranslationMap, addressTranslationAttributes)
	return addressTranslationMap
}

func flattenremoteSubnet(vpn *datatypes.Network_Tunnel_Module_Context) []map[string]interface{} {
	remoteSubnetMap := make([]map[string]interface{}, 0, 1)
	remoteSubnetAttributes := make(map[string]interface{})
	for _, customerSubnet := range vpn.CustomerSubnets {
		remoteSubnetAttributes["remote_ip_adress"] = customerSubnet.NetworkIdentifier
		remoteSubnetAttributes["remote_ip_cidr"] = customerSubnet.Cidr
		remoteSubnetAttributes["account_id"] = customerSubnet.AccountId
	}
	remoteSubnetMap = append(remoteSubnetMap, remoteSubnetAttributes)
	return remoteSubnetMap
}

// IBM Cloud Databases
func expandWhitelist(whiteList *schema.Set) (whitelist []icdv4.WhitelistEntry) {
	for _, iface := range whiteList.List() {
		wlItem := iface.(map[string]interface{})
		wlEntry := icdv4.WhitelistEntry{
			Address:     wlItem["address"].(string),
			Description: wlItem["description"].(string),
		}
		whitelist = append(whitelist, wlEntry)
	}
	return
}

// Cloud Internet Services
func flattenWhitelist(whitelist icdv4.Whitelist) []map[string]interface{} {
	entries := make([]map[string]interface{}, len(whitelist.WhitelistEntrys), len(whitelist.WhitelistEntrys))
	for i, whitelistEntry := range whitelist.WhitelistEntrys {
		l := map[string]interface{}{
			"address":     whitelistEntry.Address,
			"description": whitelistEntry.Description,
		}
		entries[i] = l
	}
	return entries
}

func expandStringMap(inVal interface{}) map[string]string {
	outVal := make(map[string]string)
	if inVal == nil {
		return outVal
	}
	for k, v := range inVal.(map[string]interface{}) {
		strValue := fmt.Sprintf("%v", v)
		outVal[k] = strValue
	}
	return outVal
}

// Cloud Internet Services
func convertTfToCisThreeVar(glbTfId string) (glbId string, zoneId string, cisId string, err error) {
	g := strings.SplitN(glbTfId, ":", 3)
	glbId = g[0]
	if len(g) > 2 {
		zoneId = g[1]
		cisId = g[2]
	} else {
		err = errors.New("cis_id or zone_id not passed")
		return
	}
	return
}
func convertCisToTfFourVar(firewallType string, ID string, ID2 string, cisID string) (buildID string) {
	if ID != "" {
		buildID = firewallType + ":" + ID + ":" + ID2 + ":" + cisID
	} else {
		buildID = ""
	}
	return
}

func convertCisToTfFiveVar(etag string, scriptFile string, ID string, ID2 string, cisID string) (buildID string) {
	if ID != "" {
		buildID = etag + ":" + scriptFile + ":" + ID + ":" + ID2 + ":" + cisID
	} else {
		buildID = ""
	}
	return
}

func convertTfToCisFiveVar(TfID string) (eTag string, scriptFile string, ID string, zoneID string, cisID string, err error) {
	g := strings.SplitN(TfID, ":", 5)
	eTag = g[0]
	if len(g) > 3 {
		scriptFile = g[1]
		ID = g[2]
		zoneID = g[3]
		cisID = g[4]

	} else {
		err = errors.New("Id or cis_id or zone_id or script file or etag not passed")
		return
	}
	return
}
func convertTfToCisFourVar(TfID string) (firewallType string, ID string, zoneID string, cisID string, err error) {
	g := strings.SplitN(TfID, ":", 4)
	firewallType = g[0]
	if len(g) > 3 {
		ID = g[1]
		zoneID = g[2]
		cisID = g[3]
	} else {
		err = errors.New("Id or cis_id or zone_id not passed")
		return
	}
	return
}

// Cloud Internet Services
func convertCisToTfThreeVar(Id string, Id2 string, cisId string) (buildId string) {
	if Id != "" {
		buildId = Id + ":" + Id2 + ":" + cisId
	} else {
		buildId = ""
	}
	return
}

// Cloud Internet Services
func convertTfToCisTwoVarSlice(tfIds []string) (Ids []string, cisId string, err error) {
	for _, item := range tfIds {
		Id := strings.SplitN(item, ":", 2)
		if len(Id) < 2 {
			err = errors.New("cis_id not passed")
			return
		}
		Ids = append(Ids, Id[0])
		cisId = Id[1]
	}
	return
}

// Cloud Internet Services
func convertCisToTfTwoVarSlice(Ids []string, cisId string) (buildIds []string) {
	for _, Id := range Ids {
		buildIds = append(buildIds, Id+":"+cisId)
	}
	return
}

// Cloud Internet Services
func convertCisToTfTwoVar(Id string, cisId string) (buildId string) {
	if Id != "" {
		buildId = Id + ":" + cisId
	} else {
		buildId = ""
	}
	return
}

// Cloud Internet Services
func convertTftoCisTwoVar(tfId string) (Id string, cisId string, err error) {
	g := strings.SplitN(tfId, ":", 2)
	Id = g[0]
	if len(g) > 1 {
		cisId = g[1]
	} else {
		err = errors.New(" cis_id or zone_id not passed")
		return
	}
	return
}

// Cloud Internet Services
func transformToIBMCISDnsData(recordType string, id string, value interface{}) (newValue interface{}, err error) {
	switch {
	case id == "flags":
		switch {
		case strings.ToUpper(recordType) == "SRV",
			strings.ToUpper(recordType) == "CAA",
			strings.ToUpper(recordType) == "DNSKEY":
			newValue, err = strconv.Atoi(value.(string))
		case strings.ToUpper(recordType) == "NAPTR":
			newValue, err = value.(string), nil
		}
	case stringInSlice(id, dnsTypeIntFields):
		newValue, err = strconv.Atoi(value.(string))
	case stringInSlice(id, dnsTypeFloatFields):
		newValue, err = strconv.ParseFloat(value.(string), 32)
	default:
		newValue, err = value.(string), nil
	}

	return
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func rcInstanceExists(resourceId string, resourceType string, meta interface{}) (bool, error) {
	// Check to see if Resource Manager instance exists
	rsConClient, err := meta.(ClientSession).ResourceControllerAPI()
	if err != nil {
		return true, nil
	}
	exists := true
	instance, err := rsConClient.ResourceServiceInstance().GetInstance(resourceId)
	if err != nil {
		if strings.Contains(err.Error(), "Object not found") ||
			strings.Contains(err.Error(), "status code: 404") {
			exists = false
		} else {
			return true, fmt.Errorf("Error checking resource instance exists: %s", err)
		}
	} else {
		if strings.Contains(instance.State, "removed") {
			exists = false
		}
	}
	if exists {
		return true, nil
	}
	// Implement when pointer to terraform.State available
	// If rcInstance is now in removed state, set TF state to removed
	// s := *terraform.State
	// for _, r := range s.RootModule().Resources {
	//  if r.Type != resourceType {
	//      continue
	//  }
	//  if r.Primary.ID == resourceId {
	//      r.Primary.Set("status", "removed")
	//  }
	// }
	return false, nil
}

// Implement when pointer to terraform.State available
// func resourceInstanceExistsTf(resourceId string, resourceType string) bool {
//  // Check TF state to see if Cloud resource instance has already been removed
//  s := *terraform.State
//  for _, r := range s.RootModule().Resources {
//      if r.Type != resourceType {
//          continue
//      }
//      if r.Primary.ID == resourceId {
//          if strings.Contains(r.Primary.Attributes["status"], "removed") {
//              return false
//          }
//      }
//  }
//  return true
// }

// convert CRN to be url safe
func EscapeUrlParm(urlParm string) string {
	if strings.Contains(urlParm, "/") {
		newUrlParm := url.PathEscape(urlParm)
		return newUrlParm
	}
	return urlParm
}

func GetTags(d *schema.ResourceData, meta interface{}) error {
	resourceID := d.Id()
	gtClient, err := meta.(ClientSession).GlobalTaggingAPI()
	if err != nil {
		return fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	taggingResult, err := gtClient.Tags().GetTags(resourceID)
	if err != nil {
		return err
	}
	var taglist []string
	for _, item := range taggingResult.Items {
		taglist = append(taglist, item.Name)
	}
	d.Set("tags", flattenStringList(taglist))
	return nil
}

func UpdateTags(d *schema.ResourceData, meta interface{}) error {
	resourceID := d.Id()
	gtClient, err := meta.(ClientSession).GlobalTaggingAPI()
	if err != nil {
		return fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	oldList, newList := d.GetChange("tags")
	if oldList == nil {
		oldList = new(schema.Set)
	}
	if newList == nil {
		newList = new(schema.Set)
	}
	olds := oldList.(*schema.Set)
	news := newList.(*schema.Set)
	removeInt := olds.Difference(news).List()
	addInt := news.Difference(olds).List()
	add := make([]string, len(addInt))
	for i, v := range addInt {
		add[i] = fmt.Sprint(v)
	}
	add = append(add, "mytag")
	remove := make([]string, len(removeInt))
	for i, v := range removeInt {
		remove[i] = fmt.Sprint(v)
	}

	if len(add) > 0 {
		_, err := gtClient.Tags().AttachTags(resourceID, add)
		if err != nil {
			return fmt.Errorf("Error updating database tags %v : %s", add, err)
		}
	}
	if len(remove) > 0 {
		_, err := gtClient.Tags().DetachTags(resourceID, remove)
		if err != nil {
			return fmt.Errorf("Error detaching database tags %v: %s", remove, err)
		}
		for _, v := range remove {
			_, err := gtClient.Tags().DeleteTag(v)
			if err != nil {
				return fmt.Errorf("Error deleting database tag %v: %s", v, err)
			}
		}
	}
	return nil
}

func GetTagsUsingCRN(meta interface{}, resourceCRN string) (*schema.Set, error) {

	gtClient, err := meta.(ClientSession).GlobalTaggingAPI()
	if err != nil {
		return nil, fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	taggingResult, err := gtClient.Tags().GetTags(resourceCRN)
	if err != nil {
		return nil, err
	}
	var taglist []string
	for _, item := range taggingResult.Items {
		taglist = append(taglist, item.Name)
	}
	log.Println("tagList: ", taglist)
	return newStringSet(resourceIBMVPCHash, taglist), nil
}

func UpdateTagsUsingCRN(oldList, newList interface{}, meta interface{}, resourceCRN string) error {
	gtClient, err := meta.(ClientSession).GlobalTaggingAPI()
	if err != nil {
		return fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	if oldList == nil {
		oldList = new(schema.Set)
	}
	if newList == nil {
		newList = new(schema.Set)
	}
	olds := oldList.(*schema.Set)
	news := newList.(*schema.Set)
	removeInt := olds.Difference(news).List()
	addInt := news.Difference(olds).List()
	add := make([]string, len(addInt))
	for i, v := range addInt {
		add[i] = fmt.Sprint(v)
	}
	remove := make([]string, len(removeInt))
	for i, v := range removeInt {
		remove[i] = fmt.Sprint(v)
	}

	schematicTags := os.Getenv("IC_ENV_TAGS")
	var envTags []string
	if schematicTags != "" {
		envTags = strings.Split(schematicTags, ",")
		add = append(add, envTags...)
	}

	if len(remove) > 0 {
		_, err := gtClient.Tags().DetachTags(resourceCRN, remove)
		if err != nil {
			return fmt.Errorf("Error detaching database tags %v: %s", remove, err)
		}
		for _, v := range remove {
			_, err := gtClient.Tags().DeleteTag(v)
			if err != nil {
				return fmt.Errorf("Error deleting database tag %v: %s", v, err)
			}
		}
	}

	if len(add) > 0 {
		_, err := gtClient.Tags().AttachTags(resourceCRN, add)
		if err != nil {
			return fmt.Errorf("Error updating database tags %v : %s", add, err)
		}
	}

	return nil
}

func getBaseController(meta interface{}) (string, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return "", err
	}
	if userDetails != nil && userDetails.cloudName == "staging" {
		return stageBaseController, nil
	}
	return prodBaseController, nil
}

func flattenSSLCiphers(ciphers []datatypes.Network_LBaaS_SSLCipher) *schema.Set {
	c := make([]string, len(ciphers))
	for i, v := range ciphers {
		c[i] = *v.Name
	}
	return newStringSet(schema.HashString, c)
}

func resourceTagsCustomizeDiff(diff *schema.ResourceDiff) error {

	if diff.Id() != "" && diff.HasChange("tags") {
		o, n := diff.GetChange("tags")
		oldSet := o.(*schema.Set)
		newSet := n.(*schema.Set)
		removeInt := oldSet.Difference(newSet).List()
		addInt := newSet.Difference(oldSet).List()
		if v := os.Getenv("IC_ENV_TAGS"); v != "" {
			s := strings.Split(v, ",")
			if len(removeInt) == len(s) && len(addInt) == 0 {
				fmt.Println("Suppresing the TAG diff ")
				return diff.Clear("tags")
			}
		}
	}
	return nil
}

func flattenRoleData(object []iampapv2.Role, roleType string) []map[string]string {
	var roles []map[string]string

	for _, item := range object {
		role := make(map[string]string)
		role["name"] = item.DisplayName
		role["type"] = roleType
		role["description"] = item.Description
		roles = append(roles, role)
	}
	return roles
}

func flattenActions(object []iampapv2.Role) map[string]interface{} {
	actions := map[string]interface{}{
		"reader":      flattenActionbyDisplayName("Reader", object),
		"manager":     flattenActionbyDisplayName("Manager", object),
		"reader_plus": flattenActionbyDisplayName("ReaderPlus", object),
		"writer":      flattenActionbyDisplayName("Writer", object),
	}
	return actions
}

func flattenActionbyDisplayName(displayName string, object []iampapv2.Role) []string {
	var actionIDs []string
	for _, role := range object {
		if role.DisplayName == displayName {
			actionIDs = role.Actions
		}
	}
	return actionIDs
}

func flattenCatalogRef(object schematics.CatalogInfo) map[string]interface{} {
	catalogRef := map[string]interface{}{
		"item_id":          object.ItemID,
		"item_name":        object.ItemName,
		"item_url":         object.ItemURL,
		"offering_version": object.OfferingVersion,
	}
	return catalogRef
}

// GetNext ...
func GetNext(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}
