package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func filterIBMDataSources(set *schema.Set, data *[]map[string]interface{}) *[]map[string]interface{} {
	var filteredVals []map[string]interface{}
	for _, v := range set.List() {
		m := v.(map[string]interface{})
		key := m["name"].(string)
		values := m["values"].([]interface{})
		vals := make([]string, len(values))
		for i, v := range values {
			vals[i] = v.(string)
		}

		for _, each := range *data {
			for _, v := range vals {
				if each[key].(string) == v {
					filteredVals = append(filteredVals, each)
				}
			}
		}
	}
	return &filteredVals
}

func dataSourceFiltersSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Computed: false,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},

				"values": {
					Type:     schema.TypeList,
					Required: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}
