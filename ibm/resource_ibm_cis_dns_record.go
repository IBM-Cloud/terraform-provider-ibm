package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	cisID                  = "cis_id"
	cisDomainID            = "domain_id"
	cisZoneName            = "zone_name"
	cisDNSRecordID         = "record_id"
	cisDNSRecordCreatedOn  = "created_on"
	cisDNSRecordModifiedOn = "modified_on"
	cisDNSRecordName       = "name"
	cisDNSRecordType       = "type"
	cisDNSRecordContent    = "content"
	cisDNSRecordProxiable  = "proxiable"
	cisDNSRecordProxied    = "proxied"
	cisDNSRecordTTL        = "ttl"
	cisDNSRecordPriority   = "priority"
	cisDNSRecordData       = "data"
)

// Constants associated with the DNS Record Type property.
// dns record type.
const (
	cisDNSRecordTypeA     = "A"
	cisDNSRecordTypeAAAA  = "AAAA"
	cisDNSRecordTypeCAA   = "CAA"
	cisDNSRecordTypeCNAME = "CNAME"
	cisDNSRecordTypeLOC   = "LOC"
	cisDNSRecordTypeMX    = "MX"
	cisDNSRecordTypeNS    = "NS"
	cisDNSRecordTypeSPF   = "SPF"
	cisDNSRecordTypeSRV   = "SRV"
	cisDNSRecordTypeTXT   = "TXT"
)

func resourceIBMCISDnsRecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMCISDnsRecordCreate,
		Read:     resourceIBMCISDnsRecordRead,
		Update:   resourceIBMCISDnsRecordUpdate,
		Delete:   resourceIBMCISDnsRecordDelete,
		Exists:   resourceIBMCISDnsRecordExist,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Description: "CIS object id or CRN",
				Required:    true,
			},
			cisDomainID: {
				Type:        schema.TypeString,
				Description: "Associated CIS domain",
				Required:    true,
			},
			cisDNSRecordName: {
				Type:     schema.TypeString,
				Optional: true,
				StateFunc: func(i interface{}) string {
					return strings.ToLower(i.(string))
				},
				DiffSuppressFunc: suppressNameDiff,
				Description:      "DNS record name",
			},
			cisDNSRecordType: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Record type",
			},
			cisDNSRecordContent: {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{cisDNSRecordData},
				Description:   "DNS record content",
			},
			cisDNSRecordData: {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{cisDNSRecordContent},
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

						// CAA record properties
						"tag": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
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
							Type:        schema.TypeString,
							Optional:    true,
							Description: "SSH fingerprint properties",
						},

						// URI record properties
						"content": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Content info",
						},
					},
				},
			},
			cisDNSRecordPriority: {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: suppressPriority,
				Description:      "Priority Value",
			},
			cisDNSRecordProxied: {
				Default:     false,
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Boolean value true if proxied else flase",
			},
			cisDNSRecordTTL: {
				Optional:    true,
				Type:        schema.TypeInt,
				Default:     1,
				Description: "TTL value",
			},
			cisDNSRecordCreatedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},

			cisDNSRecordModifiedOn: {
				Type:     schema.TypeString,
				Computed: true,
			},
			cisDNSRecordProxiable: {
				Type:     schema.TypeBool,
				Computed: true,
			},
			cisDNSRecordID: {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMCISDnsRecordCreate(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	var (
		crn            string
		zoneID         string
		recordName     string
		recordType     string
		recordContent  string
		recordPriority int
		ttl            int
		ok             interface{}
		data           interface{}
		v              interface{}
		recordData     map[string]interface{}
	)
	// session options
	crn = d.Get(cisID).(string)
	zoneID, _, err = convertTftoCisTwoVar(d.Get(cisDomainID).(string))
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	// Input options
	opt := sess.NewCreateDnsRecordOptions()

	// set record type
	recordType = d.Get(cisDNSRecordType).(string)
	opt.SetType(recordType)
	// set ttl value
	ttl = d.Get(cisDNSRecordTTL).(int)
	opt.SetTTL(int64(ttl))

	switch recordType {
	// A, AAAA, CNAME, SPF, TXT & NS records inputs
	case cisDNSRecordTypeA,
		cisDNSRecordTypeAAAA,
		cisDNSRecordTypeCNAME,
		cisDNSRecordTypeSPF,
		cisDNSRecordTypeTXT,
		cisDNSRecordTypeNS:
		// set record name & content
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		recordContent = d.Get(cisDNSRecordContent).(string)
		opt.SetContent(recordContent)

	// MX Record inputs
	case cisDNSRecordTypeMX:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		recordContent = d.Get(cisDNSRecordContent).(string)
		opt.SetContent(recordContent)
		recordPriority = d.Get(cisDNSRecordPriority).(int)
		opt.SetPriority(int64(recordPriority))

	// LOC Record inputs
	case cisDNSRecordTypeLOC:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// altitude
		v, ok = strconv.ParseFloat(dataMap["altitude"].(string), 64)
		if ok != nil {
			return fmt.Errorf("data input error")
		}
		recordData["altitude"] = v

		// lat_degrees
		v, ok = strconv.Atoi(dataMap["lat_degrees"].(string))
		if ok != nil {
			return fmt.Errorf("data input error")
		}
		recordData["lat_degrees"] = v

		// lat_direction
		recordData["lat_direction"] = dataMap["lat_direction"].(string)

		// long_direction
		recordData["long_direction"] = dataMap["long_direction"].(string)

		// lat_minutes
		v, ok = strconv.Atoi(dataMap["lat_minutes"].(string))
		if ok != nil {
			return fmt.Errorf("data input error")
		}
		recordData["lat_minutes"] = v

		// lat_seconds
		v, ok = strconv.ParseFloat(dataMap["lat_seconds"].(string), 64)
		if ok != nil {
			return fmt.Errorf("data input error")

		}
		recordData["lat_seconds"] = v

		// long_degrees
		v, ok := strconv.Atoi(dataMap["long_degrees"].(string))
		if ok != nil {
			return ok
		}
		recordData["long_degrees"] = v

		// long_minutes
		v, ok = strconv.Atoi(dataMap["long_minutes"].(string))
		if ok != nil {
			return ok
		}
		recordData["long_minutes"] = v

		// long_seconds
		i, ok := strconv.ParseFloat(dataMap["long_seconds"].(string), 64)
		if ok != nil {
			return ok
		}
		recordData["long_seconds"] = i

		// percision_horz
		i, ok = strconv.ParseFloat(dataMap["precision_horz"].(string), 64)
		if ok != nil {
			return ok
		}
		recordData["precision_horz"] = v

		// precision_vert
		i, ok = strconv.ParseFloat(dataMap["precision_vert"].(string), 64)
		if ok != nil {
			return ok
		}
		recordData["precision_vert"] = i

		// size
		i, ok = strconv.ParseFloat(dataMap["size"].(string), 64)
		if ok != nil {
			return ok
		}
		recordData["size"] = i

		opt.SetData(recordData)

	// CAA Record inputs
	case cisDNSRecordTypeCAA:

		// set record name
		recordName = d.Get(cisDNSRecordName).(string)
		opt.SetName(recordName)
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// tag
		v := dataMap["tag"].(string)
		recordData["tag"] = v

		// value
		v = dataMap["value"].(string)
		recordData["value"] = v

		opt.SetData(recordData)

	// SRV record input
	case cisDNSRecordTypeSRV:
		data, ok = d.GetOk(cisDNSRecordData)
		if ok == false {
			log.Printf("Error in getting data")
			return err
		}
		recordData = make(map[string]interface{}, 0)
		var dataMap map[string]interface{} = data.(map[string]interface{})

		// name
		v := dataMap["name"].(string)
		recordData["name"] = v

		// target
		v = dataMap["target"].(string)
		recordData["target"] = v

		// proto
		v = dataMap["proto"].(string)
		recordData["proto"] = v

		// service
		v = dataMap["service"].(string)
		recordData["service"] = v
		opt.SetData(recordData)

		// port
		s, ok := strconv.Atoi(dataMap["port"].(string))
		if ok != nil {
			return ok
		}
		recordData["port"] = s

		// priority
		s, ok = strconv.Atoi(dataMap["priority"].(string))
		if ok != nil {
			return ok
		}
		recordData["priority"] = s

		// weight
		s, ok = strconv.Atoi(dataMap["weight"].(string))
		if ok != nil {
			return ok
		}
		recordData["weight"] = s
		opt.SetData(recordData)

	default:
		name, nameOk := d.GetOk("name")
		if nameOk {
			opt.SetName(name.(string))
		}
		content, contentOk := d.GetOk("content")
		if contentOk {
			opt.SetContent(content.(string))
		}

		data, dataOk := d.GetOk("data")

		newDataMap := make(map[string]interface{})

		if dataOk {
			for id, content := range data.(map[string]interface{}) {
				newData, err := transformToIBMCISDnsData(recordType, id, content)
				if err != nil {
					return err
				} else if newData == nil {
					continue
				}
				newDataMap[id] = newData
			}

			opt.SetData(newDataMap)
		}

		if contentOk == dataOk {
			return fmt.Errorf(
				"either 'content' (present: %t) or 'data' (present: %t) must be provided",
				contentOk, dataOk)
		}

		if priority, ok := d.GetOk("priority"); ok {
			opt.SetPriority(priority.(int64))
		}
		if ttl, ok := d.GetOk("ttl"); ok {
			opt.SetTTL(ttl.(int64))
		}

	}

	result, response, err := sess.CreateDnsRecord(opt)
	if err != nil {
		log.Printf("Error creating dns record: %s, error %s", response, err)
		return err
	}

	d.SetId(convertCisToTfThreeVar(*result.Result.ID, zoneID, crn))
	return resourceIBMCISDnsRecordUpdate(d, meta)

}

func resourceIBMCISDnsRecordRead(d *schema.ResourceData, meta interface{}) error {
	var (
		crn      string
		zoneID   string
		recordID string
	)
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}

	recordID, zoneID, crn, _ = convertTfToCisThreeVar(d.Id())
	if err != nil {
		return err
	}
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewGetDnsRecordOptions(recordID)
	result, response, err := sess.GetDnsRecord(opt)
	if err != nil {
		log.Printf("Error reading dns record: %s", response)
		return err
	}

	d.Set(cisDNSRecordID, *result.Result.ID)
	d.Set(cisZoneName, *result.Result.ZoneName)
	d.Set(cisDNSRecordCreatedOn, *result.Result.CreatedOn)
	d.Set(cisDNSRecordModifiedOn, *result.Result.ModifiedOn)
	d.Set(cisDNSRecordName, *result.Result.Name)
	d.Set(cisDNSRecordType, *result.Result.Type)
	d.Set(cisDNSRecordContent, *result.Result.Content)
	d.Set(cisDomainID, *result.Result.ZoneID)
	d.Set(cisDNSRecordProxiable, *result.Result.Proxiable)
	d.Set(cisDNSRecordProxied, *result.Result.Proxied)
	d.Set(cisDNSRecordTTL, *result.Result.TTL)

	switch *result.Result.Type {
	// for MX & SRV records ouptut
	case cisDNSRecordTypeMX, cisDNSRecordTypeSRV:
		d.Set(cisDNSRecordPriority, *result.Result.Priority)
	// for LOC & CAA records output
	case cisDNSRecordTypeLOC, cisDNSRecordTypeCAA:
		d.Set(cisDNSRecordData, result.Result.Data)
	}
	return nil
}

func resourceIBMCISDnsRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		log.Printf("Error: %s", err)
		return err
	}
	var (
		recordID       string
		crn            string
		zoneID         string
		recordName     string
		recordType     string
		recordContent  string
		recordPriority int
		ttl            int
		ok             bool
		proxied        bool
		data           interface{}
		recordData     map[string]interface{}
	)
	// session options
	recordID, zoneID, crn, err = convertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading input")
		return err
	}
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	// Input options
	opt := sess.NewUpdateDnsRecordOptions(recordID)

	if d.HasChange(cisDNSRecordName) ||
		d.HasChange(cisDNSRecordType) ||
		d.HasChange(cisDNSRecordContent) ||
		d.HasChange(cisDNSRecordProxiable) ||
		d.HasChange(cisDNSRecordProxied) ||
		d.HasChange(cisDNSRecordTTL) ||
		d.HasChange(cisDNSRecordPriority) ||
		d.HasChange(cisDNSRecordData) {

		// set record type
		recordType = d.Get(cisDNSRecordType).(string)
		opt.SetType(recordType)
		// set ttl value
		ttl = d.Get(cisDNSRecordTTL).(int)
		opt.SetTTL(int64(ttl))

		// set proxied
		proxied = d.Get(cisDNSRecordProxied).(bool)
		opt.SetProxied(proxied)

		switch recordType {
		// A, AAAA, CNAME, SPF, TXT & NS records inputs
		case cisDNSRecordTypeA,
			cisDNSRecordTypeAAAA,
			cisDNSRecordTypeCNAME,
			cisDNSRecordTypeSPF,
			cisDNSRecordTypeTXT,
			cisDNSRecordTypeNS:
			// set record name & content
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
			recordContent = d.Get(cisDNSRecordContent).(string)
			opt.SetContent(recordContent)

		// MX Record inputs
		case cisDNSRecordTypeMX:

			// set record name
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)

			// set content
			recordContent = d.Get(cisDNSRecordContent).(string)
			opt.SetContent(recordContent)

			// set priority
			recordPriority = d.Get(cisDNSRecordPriority).(int)
			opt.SetPriority(int64(recordPriority))

		// LOC Record inputs
		case cisDNSRecordTypeLOC:

			// set record name
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)

			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// altitude
			v, ok := strconv.ParseFloat(dataMap["altitude"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["altitude"] = v

			// lat_degrees
			i, ok := strconv.Atoi(dataMap["lat_degrees"].(string))
			if ok != nil {
				return ok
			}
			recordData["lat_degrees"] = i

			// lat_direction
			recordData["lat_direction"] = dataMap["lat_direction"].(string)

			// long_direction
			recordData["long_direction"] = dataMap["long_direction"].(string)

			// lat_minutes
			i, ok = strconv.Atoi(dataMap["lat_minutes"].(string))
			if ok != nil {
				return ok
			}
			recordData["lat_minutes"] = i

			// lat_seconds
			v, ok = strconv.ParseFloat(dataMap["lat_seconds"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["lat_seconds"] = v

			// long_degrees
			i, ok = strconv.Atoi(dataMap["long_degrees"].(string))
			if ok != nil {
				return ok
			}
			recordData["long_degrees"] = i

			// long_minutes
			i, ok = strconv.Atoi(dataMap["long_minutes"].(string))
			if ok != nil {
				return ok
			}
			recordData["long_minutes"] = i

			// long_seconds
			v, ok = strconv.ParseFloat(dataMap["long_seconds"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["long_seconds"] = v

			// percision_horz
			v, ok = strconv.ParseFloat(dataMap["precision_horz"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["precision_horz"] = v

			// precision_vert
			v, ok = strconv.ParseFloat(dataMap["precision_vert"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["precision_vert"] = v

			// size
			v, ok = strconv.ParseFloat(dataMap["size"].(string), 64)
			if ok != nil {
				return ok
			}
			recordData["size"] = v

			opt.SetData(recordData)

		// CAA Record inputs
		case cisDNSRecordTypeCAA:

			// set record name
			recordName = d.Get(cisDNSRecordName).(string)
			opt.SetName(recordName)
			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// tag
			v := dataMap["tag"].(string)
			recordData["tag"] = v

			// value
			v = dataMap["value"].(string)
			recordData["value"] = v

			opt.SetData(recordData)

		// SRV record input
		case cisDNSRecordTypeSRV:
			data, ok = d.GetOk(cisDNSRecordData)
			if ok == false {
				log.Printf("Error in getting data")
				return err
			}
			recordData = make(map[string]interface{}, 0)
			var dataMap map[string]interface{} = data.(map[string]interface{})

			// name
			v := dataMap["name"].(string)
			recordData["name"] = v

			// target
			v = dataMap["target"].(string)
			recordData["target"] = v

			// proto
			v = dataMap["proto"].(string)
			recordData["proto"] = v

			// service
			v = dataMap["service"].(string)
			recordData["service"] = v
			opt.SetData(recordData)

			// port
			s, ok := strconv.Atoi(dataMap["port"].(string))
			if ok != nil {
				return ok
			}
			recordData["port"] = s

			// priority
			s, ok = strconv.Atoi(dataMap["priority"].(string))
			if ok != nil {
				return ok
			}
			recordData["priority"] = s

			// weight
			s, ok = strconv.Atoi(dataMap["weight"].(string))
			if ok != nil {
				return ok
			}
			recordData["weight"] = s
			opt.SetData(recordData)
		default:
			if d.HasChange(cisDNSRecordName) ||
				d.HasChange(cisDNSRecordContent) ||
				d.HasChange(cisDNSRecordProxied) ||
				d.HasChange(cisDNSRecordTTL) ||
				d.HasChange(cisDNSRecordPriority) ||
				d.HasChange(cisDNSRecordData) {

				if name, ok := d.Get(cisDNSRecordName).(string); ok {
					opt.SetName(name)
				}
				content, contentOk := d.GetOk(cisDNSRecordContent)
				if contentOk {
					opt.SetContent(content.(string))
				}
				proxied, proxiedOk := d.GetOk(cisDNSRecordProxied)
				ttl, ttlOK := d.GetOk(cisDNSRecordTTL)
				if proxiedOk {
					opt.SetProxied(proxied.(bool))
				}
				if ttlOK {
					opt.SetTTL(ttl.(int64))
				}
				if ttl != 1 && proxied == true {
					return fmt.Errorf("To enable proxy TTL should be Automatic %s",
						"i.e it should be set to 1. For the the values other than Automatic, proxy should be disabled.")
				}
				priority, priorityOk := d.GetOk(cisDNSRecordPriority)
				if priorityOk {
					opt.SetPriority(priority.(int64))
				}

				data, dataOk := d.GetOk(cisDNSRecordData)
				newDataMap := make(map[string]interface{})
				if dataOk {
					for id, content := range data.(map[string]interface{}) {
						newData, err := transformToIBMCISDnsData(recordType, id, content)
						if err != nil {
							return err
						} else if newData == nil {
							continue
						}
						newDataMap[id] = newData
					}

					opt.SetData(newDataMap)
				}
				if contentOk == dataOk {
					return fmt.Errorf(
						"either 'content' (present: %t) or 'data' (present: %t) must be provided",
						contentOk, dataOk)
				}
			}
		}

		result, response, err := sess.UpdateDnsRecord(opt)
		if err != nil {
			log.Printf("Error updating dns record: %s, error %s", response, err)
			return err
		}
		log.Printf("record id: %s", *result.Result.ID)
	}
	return resourceIBMCISDnsRecordRead(d, meta)
}

func resourceIBMCISDnsRecordDelete(d *schema.ResourceData, meta interface{}) error {
	var (
		crn      string
		zoneID   string
		recordID string
	)
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		return err
	}

	// session options
	recordID, zoneID, crn, _ = convertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading input")
		return err
	}
	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewDeleteDnsRecordOptions(recordID)
	result, response, err := sess.DeleteDnsRecord(opt)
	if err != nil {
		log.Printf("Error deleting dns record: %s", response)
		return err
	}
	log.Printf("record id: %s", *result.Result.ID)
	return err
}

func resourceIBMCISDnsRecordExist(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).CisDNSRecordClientSession()
	if err != nil {
		log.Printf("session creation failed: %s", err)
		return false, err
	}

	// session options
	recordID, zoneID, crn, _ := convertTfToCisThreeVar(d.Id())
	if err != nil {
		log.Println("Error in reading input")
		return false, err
	}

	sess.Crn = core.StringPtr(crn)
	sess.ZoneIdentifier = core.StringPtr(zoneID)

	opt := sess.NewGetDnsRecordOptions(recordID)
	_, response, err := sess.GetDnsRecord(opt)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			log.Printf("DNS record is not found")
			return false, nil
		}
		log.Printf("get DNS record failed")
		return false, err
	}
	return true, nil
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
