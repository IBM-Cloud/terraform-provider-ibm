package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMCISDnsRecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISDnsRecordCreate,
		Read:     resourceIBMCISDnsRecordRead,
		Update:   resourceIBMCISDnsRecordUpdate,
		Delete:   resourceIBMCISDnsRecordDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"cis_id": {
				Type:        schema.TypeString,
				Description: "CIS object id",
				Required:    true,
			},
			"domain_id": {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				StateFunc: func(i interface{}) string {
					return strings.ToLower(i.(string))
				},
				DiffSuppressFunc: suppressNameDiff,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"data"},
			},
			"data": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"content"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"algorithm": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"key_tag": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"flags": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"service": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"certificate": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"type": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"usage": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"selector": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"matching_type": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
						},

						// SRV record properties
						"proto": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"target": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// LOC record properties
						"size": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},
						"altitude": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},
						"long_degrees": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lat_degrees": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"precision_horz": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},
						"precision_vert": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},
						"long_direction": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"long_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"long_seconds": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},
						"lat_direction": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"lat_minutes": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"lat_seconds": {
							Type:             schema.TypeFloat,
							Computed:         true,
							Optional:         true,
							DiffSuppressFunc: suppressDataDiff,
						},

						// DNSKEY record properties
						"protocol": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"public_key": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// DS record properties
						"digest_type": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"digest": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// NAPTR record properties
						"order": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"preference": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"replacement": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// SSHFP record properties
						"fingerprint": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// URI record properties
						"content": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"priority": {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: suppressPriority,
			},

			"proxied": {
				Default:  false,
				Optional: true,
				Type:     schema.TypeBool,
			},

			"created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"modified_on": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"proxiable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceIBMCISDnsRecordCreate(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	cisId := d.Get("cis_id").(string)
	zoneId, _, err := convertTftoCisTwoVar(d.Get("domain_id").(string))
	if err != nil {
		return err
	}

	newRecord := v1.DnsBody{
		DnsType: d.Get("type").(string),
		Name:    d.Get("name").(string),
	}

	content, contentOk := d.GetOk("content")
	if contentOk {
		newRecord.Content = content.(string)
	}

	data, dataOk := d.GetOk("data")

	newDataMap := make(map[string]interface{})

	if dataOk {
		for id, content := range data.(map[string]interface{}) {
			newData, err := transformToIBMCISDnsData(newRecord.DnsType, id, content)
			if err != nil {
				return err
			} else if newData == nil {
				continue
			}
			newDataMap[id] = newData
		}

		newRecord.Data = newDataMap
	}

	if contentOk == dataOk {
		return fmt.Errorf(
			"either 'content' (present: %t) or 'data' (present: %t) must be provided",
			contentOk, dataOk)
	}

	if priority, ok := d.GetOk("priority"); ok {
		newRecord.Priority = priority.(int)
	}

	if err := validateRecordName(newRecord.DnsType, newRecord.Content); err != nil {
		return fmt.Errorf("Error validating record name %q: %s", newRecord.Name, err)
	}

	var recordPtr *v1.DnsRecord
	recordPtr, err = cisClient.Dns().CreateDns(cisId, zoneId, newRecord)
	if err != nil {
		return fmt.Errorf("Failed to create record: %s", err)
	}

	// In the Event that the API returns an empty DNS Record, we verify that the
	// ID returned is not the default ""

	record := *recordPtr

	if record.Id == "" {
		return fmt.Errorf("Failed to find record in Create response; Record was empty")
	}

	d.SetId(convertCisToTfThreeVar(record.Id, zoneId, cisId))

	return resourceIBMCISDnsRecordRead(d, meta)
}

func resourceIBMCISDnsRecordRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	recordId, zoneId, cisId, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}

	var recordPtr *v1.DnsRecord
	recordPtr, err = cisClient.Dns().GetDns(cisId, zoneId, recordId)
	if err != nil {
		if checkCisRecordDeleted(d, meta, err, recordPtr) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during DNS Record Read %v\n", err)
		return err
	}

	record := *recordPtr
	d.Set("cis_id", cisId)
	d.Set("domain_id", convertCisToTfTwoVar(zoneId, cisId))
	d.Set("name", record.Name)
	d.Set("type", record.DnsType)
	d.Set("content", record.Content)
	d.Set("ttl", record.Ttl)
	d.Set("priority", record.Priority)
	d.Set("proxied", record.Proxied)
	d.Set("created_on", record.CreatedOn.Format(time.RFC3339Nano))
	d.Set("data", expandStringMap(record.Data))
	d.Set("modified_on", record.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("proxiable", record.Proxiable)

	return nil
}

func resourceIBMCISDnsRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMCISDnsRecordRead(d, meta)
}

func resourceIBMCISDnsRecordDelete(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(ClientSession).CisAPI()
	if err != nil {
		return err
	}
	recordId, zoneId, cisId, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}

	var recordPtr *v1.DnsRecord
	recordPtr, err = cisClient.Dns().GetDns(cisId, zoneId, recordId)
	if err != nil {
		if checkCisRecordDeleted(d, meta, err, recordPtr) {
			d.SetId("")
			return nil
		}
		log.Printf("[WARN] Error getting zone during DNS Record Read %v\n", err)
		return err
	}
	err = cisClient.Dns().DeleteDns(cisId, zoneId, recordId)
	if err != nil {
		return fmt.Errorf("Error deleting IBMCISDNS Record: %s", err)
	}

	return nil
}

var dnsTypeIntFields = []string{
	"algorithm",
	"key_tag",
	"type",
	"usage",
	"selector",
	"matching_type",
	"weight",
	"priority",
	"port",
	"long_degrees",
	"lat_degrees",
	"long_minutes",
	"lat_minutes",
	"protocol",
	"digest_type",
	"order",
	"preference",
}

var dnsTypeFloatFields = []string{
	"size",
	"altitude",
	"precision_horz",
	"precision_vert",
	"long_seconds",
	"lat_seconds",
}

func suppressPriority(k, old, new string, d *schema.ResourceData) bool {
	recordType := d.Get("type").(string)
	if recordType != "MX" && recordType != "URI" {
		return true
	}
	return false
}

func suppressNameDiff(k, old, new string, d *schema.ResourceData) bool {
	// CIS concantenates name with domain. So just check name is the same
	if strings.SplitN(old, ".", 2)[0] == strings.SplitN(new, ".", 2)[0] {
		return true
	}
	// If name is @, its replaced by the domain name. So ignore check.
	if new == "@" {
		return true
	}

	return false
}

func suppressDataDiff(k, old, new string, d *schema.ResourceData) bool {
	// Tuncate after .
	if strings.SplitN(old, ".", 2)[0] == strings.SplitN(new, ".", 2)[0] {
		return true
	}
	return false
}

func checkCisRecordDeleted(d *schema.ResourceData, meta interface{}, errCheck error, record *v1.DnsRecord) bool {
	// Check if error is due to removal of Cis resource and hence all subresources
	if strings.Contains(errCheck.Error(), "Object not found") ||
		strings.Contains(errCheck.Error(), "status code: 404") ||
		strings.Contains(errCheck.Error(), "status code: 400") ||
		strings.Contains(errCheck.Error(), "Invalid dns record identifier") {
		log.Printf("[WARN] Removing resource from state because it's not found via the CIS API")
		return true
	}
	_, _, cisId, _ := convertTfToCisThreeVar(d.Id())
	exists, errNew := rcInstanceExists(cisId, "ibm_cis", meta)
	if errNew != nil {
		log.Printf("resourceCISDnsRecordRead - Failure validating service exists %s\n", errNew)
		return false
	}
	if !exists {
		log.Printf("[WARN] Removing Dns Record from state because parent cis instance is in removed state")
		return true
	}
	return false
}
