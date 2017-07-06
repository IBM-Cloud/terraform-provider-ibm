package ibm

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Bluemix/bluemix-go/api/mccp/mccpv2"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
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

func flattenSpaceRoleUsers(in []mccpv2.SpaceRole) *schema.Set {
	var out = []interface{}{}
	for _, v := range in {
		out = append(out, v.UserName)
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
		m["credentials"] = flattenServiceKeyCredentials(k.Entity.Credentials)
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
