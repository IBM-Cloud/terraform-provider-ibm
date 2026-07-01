package configurationaggregator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/configuration-aggregator-go-sdk/configurationaggregatorv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIbmConfigAggregatorBatchConfigurations() *schema.Resource {
	log.Printf("test print")
	return &schema.Resource{
		ReadContext: dataSourceIbmConfigAggregatorBatchConfigurationsRead,

		Schema: map[string]*schema.Schema{

			// Path params
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},

			// INPUT
			// renamed to avoid duplicate key issue
			"config": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of resource CRNs for batch query",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_crn": {
							Type:     schema.TypeString,
							Required: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"config_type": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			// OUTPUT
			"configs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configurations returned from API",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"about": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"config": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"config_v2": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Errors
			"errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource_crn": {Type: schema.TypeString, Computed: true},
						"message":      {Type: schema.TypeString, Computed: true},
						"error_code":   {Type: schema.TypeString, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceIbmConfigAggregatorBatchConfigurationsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	configurationAggregatorClient, err := meta.(conns.ClientSession).ConfigurationAggregatorV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_batch_configurations", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	region := getConfigurationInstanceRegion(configurationAggregatorClient, d)
	instanceId := d.Get("instance_id").(string)
	log.Printf("Fetching config for instance_id: %s", instanceId)
	configurationAggregatorClient = getClientWithConfigurationInstanceEndpoint(configurationAggregatorClient, instanceId, region)

	listBatchConfigsOptions := &configurationaggregatorv1.ListBatchConfigsOptions{}

	var configs []configurationaggregatorv1.Requestconfigs
	for _, v := range d.Get("config").([]interface{}) {
		value := v.(map[string]interface{})
		configsItem, err := ResourceMapToRequestconfigs(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_batch_configurations", "read", "parse-configs").GetDiag()
		}
		configs = append(configs, configsItem)
	}
	listBatchConfigsOptions.SetConfigs(configs)

	var listConfigsResponse *configurationaggregatorv1.ListConfigsQueryResponse
	var offset int64
	finalList := []configurationaggregatorv1.Config{}

	for {
		result, _, err := configurationAggregatorClient.ListBatchConfigsWithContext(context, listBatchConfigsOptions)
		listConfigsResponse = result
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListBatchConfigsWithContext failed: %s", err.Error()), "(Data) ibm_config_aggregator_batch_configurations", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
		offset = dataSourceListConfigsResponseGetNext(result.Next)
		finalList = append(finalList, result.Configs...)
		if offset == 0 {
			break
		}
	}

	listConfigsResponse.Configs = finalList

	d.SetId(dataSourceIbmConfigAggregatorBatchConfigurationsID(d))

	if !core.IsNil(listConfigsResponse.Prev) {
		prev := []map[string]interface{}{}
		prevMap, err := DataSourceIbmConfigAggregatorBatchConfigurationsPaginatedPreviousToMap(listConfigsResponse.Prev)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_config_aggregator_batch_configurations", "read", "prev-to-map").GetDiag()
		}
		prev = append(prev, prevMap)
		if err = d.Set("prev", prev); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting prev: %s", err), "(Data) ibm_config_aggregator_batch_configurations", "read", "set-prev").GetDiag()
		}
	}
	configsList := make([]map[string]interface{}, 0)

	for _, item := range listConfigsResponse.Configs {

		configMap := make(map[string]interface{})

		// about
		if item.About != nil {
			aboutBytes, _ := json.Marshal(item.About)
			configMap["about"] = string(aboutBytes)
		} else {
			configMap["about"] = "{}"
		}

		// config
		if item.Config != nil {
			configBytes, _ := json.Marshal(item.Config)
			configMap["config"] = string(configBytes)
		} else {
			configMap["config"] = "{}"
		}

		configsList = append(configsList, configMap)
	}

	log.Printf("[DEBUG] Setting configs: %+v", configsList)

	if err := d.Set("configs", configsList); err != nil {
		return flex.DiscriminatedTerraformErrorf(
			err,
			fmt.Sprintf("Error setting configs: %s", err),
			"(Data) ibm_config_aggregator_batch_configurations",
			"read",
			"set-configs",
		).GetDiag()
	}

	if !core.IsNil(listConfigsResponse.Errors) {

		errList := make([]map[string]interface{}, 0)

		for _, e := range listConfigsResponse.Errors {

			errMap := map[string]interface{}{
				"resource_crn": safeString(e.ResourceCrn),
				"message":      safeString(e.Message),
				"error_code":   safeString(e.ErrorCode),
			}

			errList = append(errList, errMap)
		}

		log.Printf("[DEBUG] Setting errors: %+v", errList)

		if err := d.Set("errors", errList); err != nil {
			return flex.DiscriminatedTerraformErrorf(
				err,
				fmt.Sprintf("Error setting errors: %s", err),
				"(Data) ibm_config_aggregator_batch_configurations",
				"read",
				"set-errors",
			).GetDiag()
		}
	}

	return nil
}

func dataSourceListConfigsResponseGetNext(next interface{}) int64 {
	if reflect.ValueOf(next).IsNil() {
		return 0
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return 0
	}

	q := u.Query()
	var page string

	if q.Get("start") != "" {
		page = q.Get("start")
	} else if q.Get("offset") != "" {
		page = q.Get("offset")
	}

	convertedVal, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return convertedVal
}

// dataSourceIbmConfigAggregatorBatchConfigurationsID returns a reasonable ID for the list.
func dataSourceIbmConfigAggregatorBatchConfigurationsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmConfigAggregatorBatchConfigurationsPaginatedPreviousToMap(model *configurationaggregatorv1.PaginatedPrevious) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.Start != nil {
		modelMap["start"] = *model.Start
	}
	return modelMap, nil
}

func ResourceMapToRequestconfigs(value map[string]interface{}) (configurationaggregatorv1.Requestconfigs, error) {

	config := configurationaggregatorv1.Requestconfigs{}

	// required
	if v, ok := value["resource_crn"]; ok && v.(string) != "" {
		config.ResourceCrn = core.StringPtr(v.(string))
	} else {
		return config, fmt.Errorf("resource_crn is required")
	}

	// optional
	if v, ok := value["service_name"]; ok && v.(string) != "" {
		config.ServiceName = core.StringPtr(v.(string))
	}

	if v, ok := value["type_id"]; ok && v.(string) != "" {
		config.TypeID = core.StringPtr(v.(string))
	}

	if v, ok := value["config_type"]; ok {
		raw := v.([]interface{})
		var list []string
		for _, item := range raw {
			list = append(list, item.(string))
		}
		config.ConfigType = list
	}

	return config, nil
}

func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
