package ibm

import (
	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
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
		},

		Create:   resourceCISSettingsUpdate,
		Read:     resourceCISSettingsRead,
		Update:   resourceCISSettingsUpdate,
		Delete:   resourceCISSettingsDelete,
		Importer: &schema.ResourceImporter{},
	}
}

var settingsList = [...]string{"waf", "ssl", "min_tls_version", "automatic_https_rewrites", "opportunistic_encryption", "cname_flattening"}

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

		value := d.Get(item).(string)
		if value != "" {
			settingsNew := v1.SettingsBody{Value: value}
			_, err = cisClient.Settings().UpdateSetting(cisId, zoneId, item, settingsNew)
			if err != nil {
				log.Printf("Update settings Failed on %s, %s\n", item, err)
				return err
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
