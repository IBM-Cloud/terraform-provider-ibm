/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */
package ibm

import (
	"log"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMCISCacheSetting() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCISCacheSettingsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Description:      "Associated CIS domain",
				Required:         true,
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisCacheSettingsCachingLevel: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cache Level Setting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			cisCacheServeStaleContent: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Serve Stale Content ",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			cisCacheSettingsBrowserExpiration: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Browser Expiration setting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			cisCacheSettingsDevelopmentMode: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Development mode setting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			cisCacheSettingsQueryStringSort: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Query String sort setting",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"editable": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceCISCacheSettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisCacheClientSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	zoneID, _, _ := convertTftoCisTwoVar(d.Get(cisDomainID).(string))
	cisClient.Crn = core.StringPtr(crn)
	log.Println("*********** crn ", crn)
	cisClient.ZoneID = core.StringPtr(zoneID)
	log.Println("*********** Zone ID ", zoneID)

	// Cache Level Setting
	cacheLevel_result, resp, err := cisClient.GetCacheLevel(cisClient.NewGetCacheLevelOptions())
	log.Println("*********** cache level ", resp)

	if err != nil {
		log.Printf("Get Cache Level  setting failed : %v\n", resp)
		return err
	}
	if cacheLevel_result != nil || cacheLevel_result.Result != nil {
		log.Println("*************cacheLevel_result.Result", cacheLevel_result.Result)
		cacheLevels := make([]map[string]interface{}, 0)
		cacheLevel := make(map[string]interface{})

		if cacheLevel_result.Result.ID != nil {
			cacheLevel["id"] = cacheLevel_result.Result.ID
		}
		if cacheLevel_result.Result.Value != nil {
			cacheLevel["value"] = cacheLevel_result.Result.Value
		}
		if cacheLevel_result.Result.Editable != nil {
			cacheLevel["editable"] = cacheLevel_result.Result.Editable
		}
		if cacheLevel_result.Result.ModifiedOn != nil {
			cacheLevel["modified_on"] = cacheLevel_result.Result.ModifiedOn
		}
		cacheLevels = append(cacheLevels, cacheLevel)
		d.Set(cisCacheSettingsCachingLevel, cacheLevels)
		log.Println("$$$$$$$$$$$$$$$$ cache level setting $$$$$", d)
	}
	// Serve Stale Content setting
	servestaleContent_result, resp, err := cisClient.GetServeStaleContent(cisClient.NewGetServeStaleContentOptions())
	log.Println("*********** servestaleContent_result ", resp)
	if err != nil {
		log.Printf("Get Serve Stale Content setting failed : %v\n", resp)
		return err
	}
	if servestaleContent_result != nil || servestaleContent_result.Result != nil {

		servestalecontents := make([]map[string]interface{}, 0)
		servestalecontent := make(map[string]interface{})

		if servestaleContent_result.Result.ID != nil {
			servestalecontent["id"] = servestaleContent_result.Result.ID
		}
		if servestaleContent_result.Result.Value != nil {
			servestalecontent["value"] = servestaleContent_result.Result.Value
		}
		if servestaleContent_result.Result.Editable != nil {
			servestalecontent["editable"] = servestaleContent_result.Result.Editable
		}
		if servestaleContent_result.Result.ModifiedOn != nil {
			servestalecontent["modified_on"] = servestaleContent_result.Result.ModifiedOn
		}
		servestalecontents = append(servestalecontents, servestalecontent)
		d.Set(cisCacheServeStaleContent, servestalecontents)
		log.Println("$$$$$$$$$$$$$$$$ serve stale Setting$$$$$", d)
	}

	// Browser Expiration setting
	browserCacheTTL_result, resp, err := cisClient.GetBrowserCacheTTL(cisClient.NewGetBrowserCacheTtlOptions())
	log.Println("*********** browserCacheTTL_result ", resp)
	if err != nil {
		log.Printf("Get browser expiration setting failed : %v\n", resp)
		return err
	}
	if browserCacheTTL_result != nil || browserCacheTTL_result.Result != nil {

		browserCacheTTLs := make([]map[string]interface{}, 0)
		browserCacheTTL := make(map[string]interface{})

		if browserCacheTTL_result.Result.ID != nil {
			browserCacheTTL["id"] = browserCacheTTL_result.Result.ID
		}
		if browserCacheTTL_result.Result.Value != nil {
			browserCacheTTL["value"] = browserCacheTTL_result.Result.Value
		}
		if browserCacheTTL_result.Result.Editable != nil {
			browserCacheTTL["editable"] = browserCacheTTL_result.Result.Editable
		}
		if browserCacheTTL_result.Result.ModifiedOn != nil {
			browserCacheTTL["modified_on"] = browserCacheTTL_result.Result.ModifiedOn
		}
		browserCacheTTLs = append(browserCacheTTLs, browserCacheTTL)
		d.Set(cisCacheSettingsBrowserExpiration, browserCacheTTLs)
		log.Println("$$$$$$$$$$$$$$$$ browswe Setting$$$$$", d)
	}
	// development mode setting
	devMode_result, resp, err := cisClient.GetDevelopmentMode(cisClient.NewGetDevelopmentModeOptions())
	log.Println("*********** devMode_result ", resp)
	if err != nil {
		log.Printf("Get development mode setting failed : %v", resp)
		return err
	}
	if devMode_result != nil || devMode_result.Result != nil {

		devModes := make([]map[string]interface{}, 0)
		devMode := make(map[string]interface{})

		if devMode_result.Result.ID != nil {
			devMode["id"] = devMode_result.Result.ID
		}
		if devMode_result.Result.Value != nil {
			devMode["value"] = devMode_result.Result.Value
		}
		if devMode_result.Result.Editable != nil {
			devMode["editable"] = devMode_result.Result.Editable
		}
		if devMode_result.Result.ModifiedOn != nil {
			devMode["modified_on"] = devMode_result.Result.ModifiedOn
		}
		devModes = append(devModes, devMode)
		d.Set(cisCacheSettingsDevelopmentMode, devModes)
		log.Println("$$$$$$$$$$$$$$$$ dev mod Setting$$$$$", d)
	}

	// Query string sort setting
	queryStringSort_result, resp, err := cisClient.GetQueryStringSort(cisClient.NewGetQueryStringSortOptions())
	log.Println("*********** queryStringSort_result ", resp)
	if err != nil {
		log.Printf("Get query string sort setting failed : %v", resp)
		return err
	}
	if queryStringSort_result != nil || queryStringSort_result.Result != nil {

		queryStringSorts := make([]map[string]interface{}, 0)
		queryStringSort := make(map[string]interface{})

		if queryStringSort_result.Result.ID != nil {
			queryStringSort["id"] = queryStringSort_result.Result.ID
		}
		if queryStringSort_result.Result.Value != nil {
			queryStringSort["value"] = queryStringSort_result.Result.Value
		}
		if queryStringSort_result.Result.Editable != nil {
			queryStringSort["editable"] = queryStringSort_result.Result.Editable
		}
		if queryStringSort_result.Result.ModifiedOn != nil {
			queryStringSort["modified_on"] = queryStringSort_result.Result.ModifiedOn
		}
		queryStringSorts = append(queryStringSorts, queryStringSort)
		d.Set(cisCacheSettingsQueryStringSort, queryStringSorts)
		log.Println("$$$$$$$$$$$$$$$$ query Setting$$$$$", d)
	}
	d.SetId(dataSourceIBMCISCertificatesID(d))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)
	return nil
}
