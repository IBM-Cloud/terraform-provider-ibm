package ibm

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
)

func resourceIBMCISSettings() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"waf": {
				Type:         schema.TypeString,
				Description:  "WAF setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "on"}),
			},
			"ssl": {
				Type:         schema.TypeString,
				Description:  "SSL/TLS setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"off", "flexible", "full", "strict", "origin_pull"}),
			},
			"certificate_status": {
				Type:        schema.TypeString,
				Description: "Certificate status",
				Computed:    true,
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Description:  "Minimum version of TLS required",
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"1.1", "1.2", "1.3", "1.4"}),
				Default:      "1.1",
			},
			"cname_flattening": {
				Type:         schema.TypeString,
				Description:  "cname_flattening setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"flatten_at_root", "flatten_all", "flatten_none"}),
			},
			"opportunistic_encryption": {
				Type:         schema.TypeString,
				Description:  "opportunistic_encryption setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"automatic_https_rewrites": {
				Type:         schema.TypeString,
				Description:  "automatic_https_rewrites setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"always_use_https": {
				Type:         schema.TypeString,
				Description:  "always_use_https setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"ipv6": {
				Type:         schema.TypeString,
				Description:  "ipv6 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"browser_check": {
				Type:         schema.TypeString,
				Description:  "browser_check setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"hotlink_protection": {
				Type:         schema.TypeString,
				Description:  "hotlink_protection setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"http2": {
				Type:         schema.TypeString,
				Description:  "http2 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"image_load_optimization": {
				Type:         schema.TypeString,
				Description:  "image_load_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"image_size_optimization": {
				Type:         schema.TypeString,
				Description:  "image_size_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"lossless", "off", "lossy"}),
			},
			"ip_geolocation": {
				Type:         schema.TypeString,
				Description:  "ip_geolocation setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"origin_error_page_pass_thru": {
				Type:         schema.TypeString,
				Description:  "origin_error_page_pass_thru setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"brotli": {
				Type:         schema.TypeString,
				Description:  "brotli setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"pseudo_ipv4": {
				Type:         schema.TypeString,
				Description:  "pseudo_ipv4 setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"overwrite_header", "off", "add_header"}),
			},
			"prefetch_preload": {
				Type:         schema.TypeString,
				Description:  "prefetch_preload setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"response_buffering": {
				Type:         schema.TypeString,
				Description:  "response_buffering setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"script_load_optimization": {
				Type:         schema.TypeString,
				Description:  "script_load_optimization setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"server_side_exclude": {
				Type:         schema.TypeString,
				Description:  "server_side_exclude setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"tls_client_auth": {
				Type:         schema.TypeString,
				Description:  "tls_client_auth setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"true_client_ip_header": {
				Type:         schema.TypeString,
				Description:  "true_client_ip_header setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
			"websockets": {
				Type:         schema.TypeString,
				Description:  "websockets setting",
				Optional:     true,
				Computed:     true,
				ValidateFunc: validateAllowedStringValue([]string{"on", "off"}),
			},
		},

		Create:   resourceCISSettingsUpdate,
		Read:     resourceCISSettingsRead,
		Update:   resourceCISSettingsUpdate,
		Delete:   resourceCISSettingsDelete,
		Importer: &schema.ResourceImporter{},
	}
}

var settingsList = [...]string{"waf", "ssl", "min_tls_version", "automatic_https_rewrites", "opportunistic_encryption", "cname_flattening", "always_use_https", "ipv6", "browser_check", "hotlink_protection", "http2", "image_load_optimization", "image_size_optimization", "ip_geolocation", "origin_error_page_pass_thru", "brotli", "pseudo_ipv4", "prefetch_preload", "response_buffering", "script_load_optimization", "server_side_exclude", "tls_client_auth", "true_client_ip_header", "websockets"}

func resourceCISSettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	zoneId, cisId, _ := convertTftoCisTwoVar(d.Get("domain_id").(string))

	type Setting struct {
		Name  string
		Value string
	}

	for _, item := range settingsList {
		if d.HasChange(item) {
			if v, ok := d.GetOk(item); ok {
				value := v.(string)
				settingsNew := v1.SettingsBody{Value: value}
				_, err = cisClient.Settings().UpdateSetting(cisId, zoneId, item, settingsNew)
				if err != nil {
					log.Printf("Update settings Failed on %s, %s\n", item, err)
					return err
				}
			}
		}
	}

	d.SetId(convertCisToTfTwoVar(zoneId, cisId))

	return resourceCISSettingsRead(d, meta)
}

func resourceCISSettingsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}

	settingsId, cisId, _ := convertTftoCisTwoVar(d.Id())
	for _, item := range settingsList {
		settingsResult, err := cisClient.Settings().GetSetting(cisId, settingsId, item)
		if err != nil {
			if checkCisSettingsDeleted(d, meta, err, cisId) {
				d.SetId("")
				return nil
			}
			if strings.Contains(err.Error(), "Request failed with status code: 405") {
				log.Printf("[WARN] This action is unavailable for the current plan")
				return nil
			}
			log.Printf("[WARN] Error getting zone during DomainRead %v\n", err)
			return err
		} else {
			settingsObj := *settingsResult
			d.Set(item, settingsObj.Value)
		}
	}
	return nil
}

func resourceCISSettingsDelete(d *schema.ResourceData, meta interface{}) error {
	// Nothing to delete on CIS resource
	d.SetId("")
	return nil
}

func checkCisSettingsDeleted(d *schema.ResourceData, meta interface{}, errCheck error, cisId string) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "Invalid zone identifier") { //code 400
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("resourceCISdomainRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing domain settings from state because parent cis instance is in removed state")
		return true
	}
	return false
}
