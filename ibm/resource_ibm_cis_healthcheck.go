package ibm

import (
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
)

func resourceIBMCISHealthCheck() *schema.Resource {
	return &schema.Resource{

		Create:   resourceCIShealthCheckCreate,
		Read:     resourceCIShealthCheckRead,
		Update:   resourceCIShealthCheckUpdate,
		Delete:   resourceCIShealthCheckDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS instance crn",
				Required:    true,
			},
			"path": {
				Type:        schema.TypeString,
				Description: "path",
				Optional:    true,
			},
			"expected_body": {
				Type:        schema.TypeString,
				Description: "expected_body",
				Optional:    true,
			},
			"expected_codes": {
				Type:        schema.TypeString,
				Description: "expected_codes",
				Optional:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "description",
				Optional:    true,
			},
			"type": {
				Type:         schema.TypeString,
				Description:  "type",
				Optional:     true,
				Default:      "http",
				ValidateFunc: validateAllowedStringValue([]string{"http", "https", "tcp"}),
			},
			"method": {
				Type:         schema.TypeString,
				Description:  "method",
				Optional:     true,
				ValidateFunc: validateAllowedStringValue([]string{"GET", "HEAD"}),
			},
			"timeout": {
				Type:        schema.TypeInt,
				Description: "timeout",
				Optional:    true,
				Default:     5,
				//ValidateFunc: validation.IntBetween(1, 10),
			},
			"retries": {
				Type:        schema.TypeInt,
				Description: "retries",
				Optional:    true,
				Default:     2,
				//ValidateFunc: validation.IntBetween(1, 5),
			},
			"interval": {
				Type:        schema.TypeInt,
				Description: "interval",
				Optional:    true,
				Default:     60,
			},
			"follow_redirects": {
				Type:        schema.TypeBool,
				Description: "follow_redirects",
				Optional:    true,
			},
			"allow_insecure": {
				Type:        schema.TypeBool,
				Description: "allow_insecure",
				Optional:    true,
				Default:     false,
			},
			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// "header": {
			// 	Type:     schema.TypeMap,
			// 	Optional: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"header": {
			// 				Type:     schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"values": {
			// 				Type:     schema.TypeSet,
			// 				Required: true,
			// 				Elem: &schema.Schema{
			// 					Type: schema.TypeString,
			// 				},
			// 			},
			// 		},
			// 	},
			// 	Set: HashByMapKey("header"),
			// },
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceCIShealthCheckCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	cisId := d.Get("cis_id").(string)

	monitorNew := v1.MonitorBody{
		MonType: d.Get("type").(string),
	}
	if expCodes, ok := d.GetOk("expected_codes"); ok {
		monitorNew.ExpCodes = expCodes.(string)
	}
	if expBody, ok := d.GetOk("expected_body"); ok {
		monitorNew.ExpBody = expBody.(string)
	}
	if monPath, ok := d.GetOk("path"); ok {
		monitorNew.Path = monPath.(string)
	}
	if description, ok := d.GetOk("description"); ok {
		monitorNew.Description = description.(string)
	}
	if method, ok := d.GetOk("method"); ok {
		monitorNew.Method = method.(string)
	}
	if timeout, ok := d.GetOk("timeout"); ok {
		monitorNew.Timeout = timeout.(int)
	}
	if retries, ok := d.GetOk("retries"); ok {
		monitorNew.Retries = retries.(int)
	}
	if interval, ok := d.GetOk("interval"); ok {
		monitorNew.Interval = interval.(int)
	}
	if follow_redirects, ok := d.GetOk("follow_redirects"); ok {
		monitorNew.FollowRedirects = follow_redirects.(bool)
	}
	if allow_insecure, ok := d.GetOk("allow_insecure"); ok {
		monitorNew.AllowInsecure = allow_insecure.(bool)
	}
	if port, ok := d.GetOk("port"); ok {
		monitorNew.Port = port.(int)
	}
	var monitor *v1.Monitor
	var monitorObj v1.Monitor

	monitor, err = cisClient.Monitors().CreateMonitor(cisId, monitorNew)
	if err != nil {
		log.Printf("CreateMonitors Failed %s\n", err)
		return err
	}
	monitorObj = *monitor
	//Set unique TF Id from concatenated CIS Ids
	d.SetId(monitorObj.Id + ":" + cisId)
	d.Set("path", monitorObj.Path)

	return resourceCIShealthCheckRead(d, meta)
}

func resourceCIShealthCheckRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	monitorId, cisId, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	var monitor *v1.Monitor
	monitor, err = cisClient.Monitors().GetMonitor(cisId, monitorId)
	if err != nil {
		if checkCisMonitorDeleted(d, meta, err, monitor) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during MonitorRead %v\n", err)
		return err
	}
	monitorObj := *monitor
	d.Set("cis_id", cisId)
	d.Set("description", monitorObj.Description)
	d.Set("path", monitorObj.Path)
	d.Set("expected_body", monitorObj.ExpBody)
	d.Set("expected_codes", monitorObj.ExpCodes)
	d.Set("type", monitorObj.MonType)
	d.Set("method", monitorObj.Method)
	d.Set("timeout", monitorObj.Timeout)
	d.Set("retries", monitorObj.Retries)
	d.Set("interval", monitorObj.Interval)
	d.Set("follow_redirects", monitorObj.FollowRedirects)
	d.Set("allow_insecure", monitorObj.AllowInsecure)
	d.Set("port", monitorObj.Port)
	return nil
}

func resourceCIShealthCheckUpdate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	monitorId, cisId, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	monitorUpdate := v1.MonitorBody{}
	if d.HasChange("type") || d.HasChange("description") || d.HasChange("path") || d.HasChange("expected_body") || d.HasChange("expected_codes") || d.HasChange("method") || d.HasChange("timeout") || d.HasChange("retries") || d.HasChange("interval") || d.HasChange("follow_redirects") || d.HasChange("allow_insecure") || d.HasChange("port") {
		if monType, ok := d.GetOk("type"); ok {
			monitorUpdate.MonType = monType.(string)
		}
		if expCodes, ok := d.GetOk("expected_codes"); ok {
			monitorUpdate.ExpCodes = expCodes.(string)
		}
		if expBody, ok := d.GetOk("expected_body"); ok {
			monitorUpdate.ExpBody = expBody.(string)
		}
		if monPath, ok := d.GetOk("path"); ok {
			monitorUpdate.Path = monPath.(string)
		}
		if description, ok := d.GetOk("description"); ok {
			monitorUpdate.Description = description.(string)
		}
		if method, ok := d.GetOk("method"); ok {
			monitorUpdate.Method = method.(string)
		}
		if timeout, ok := d.GetOk("timeout"); ok {
			monitorUpdate.Timeout = timeout.(int)
		}
		if retries, ok := d.GetOk("retries"); ok {
			monitorUpdate.Retries = retries.(int)
		}
		if interval, ok := d.GetOk("interval"); ok {
			monitorUpdate.Interval = interval.(int)
		}
		if follow_redirects, ok := d.GetOk("follow_redirects"); ok {
			monitorUpdate.FollowRedirects = follow_redirects.(bool)
		}
		if allow_insecure, ok := d.GetOk("allow_insecure"); ok {
			monitorUpdate.AllowInsecure = allow_insecure.(bool)
		}
		if port, ok := d.GetOk("port"); ok {
			monitorUpdate.Port = port.(int)
		}
		_, err = cisClient.Monitors().UpdateMonitor(cisId, monitorId, monitorUpdate)
		if err != nil {
			log.Printf("[WARN] Error getting zone during MonitorUpdate %v \n", err)
			return err
		}
	}
	return resourceCIShealthCheckRead(d, meta)
}

func resourceCIShealthCheckDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	monitorId, cisId, err := convertTftoCisTwoVar(d.Id())
	if err != nil {
		return err
	}
	var monitor *v1.Monitor
	emptyMonitor := new(v1.Monitor)
	monitor, err = cisClient.Monitors().GetMonitor(cisId, monitorId)
	if err != nil {
		if checkCisMonitorDeleted(d, meta, err, monitor) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during MonitorRead %v\n", err)
		return err
	}

	monitorObj := *monitor
	if !reflect.DeepEqual(emptyMonitor, monitorObj) {
		err = cisClient.Monitors().DeleteMonitor(cisId, monitorId)
		if err != nil {
			log.Printf("[WARN] DeleteMonitor Failed %s\n", err)
			return err
		}
	}

	d.SetId("")
	return nil
}

func checkCisMonitorDeleted(d *schema.ResourceData, meta interface{}, errCheck error, monitor *v1.Monitor) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "Invalid zone identifier") { //code 400
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	_, cisId, _ := convertTftoCisTwoVar(d.Id())
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("resourceCISmonitorRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing monitor from state because parent cis instance is in removed state")
		return true
	}
	return false
}
