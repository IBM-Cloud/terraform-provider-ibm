package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var allowedPrivateDomainRecordTypes = []string{
	"A", "AAAA", "CNAME", "MX", "PTR", "SRV", "TXT",
}

const (
	pdnsResourceRecordID = "resource_record_id"
	pdnsRecordType       = "type"
	pdnsRecordTTL        = "ttl"
	pdnsRecordName       = "name"
	pdnsRdata            = "rdata"
	pdnsMxPreference     = "preference"
	pdnsSrvPort          = "port"
	pdnsSrvPriority      = "priority"
	pdnsSrvWeight        = "weight"
	pdnsSrvProtocol      = "protocol"
	pdnsSrvService       = "service"
	pdnsRecordCreatedOn  = "created_on"
	pdnsRecordModifiedOn = "modified_on"
)

func caseDiffSuppress(_, old, new string, _ *schema.ResourceData) bool {
	return strings.ToUpper(old) == strings.ToUpper(new)
}

func resourceIBMPrivateDNSResourceRecord() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMPrivateDNSResourceRecordCreate,
		Read:     resourceIBMPrivateDNSResourceRecordRead,
		Update:   resourceIBMPrivateDNSResourceRecordUpdate,
		Delete:   resourceIBMPrivateDNSResourceRecordDelete,
		Exists:   resourceIBMPrivateDNSResourceRecordExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			pdnsResourceRecordID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource record ID",
			},

			pdnsInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID",
			},

			pdnsZoneID: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Zone ID",
			},

			pdnsRecordName: {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: caseDiffSuppress,
				Description:      "DNS record name",
			},

			pdnsRecordType: {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					for _, rtype := range allowedPrivateDomainRecordTypes {
						if value == rtype {
							return
						}
					}

					errors = append(
						errors,
						fmt.Errorf("%s is not one of the valid domain record types: %s",
							value, strings.Join(allowedPrivateDomainRecordTypes, ", "),
						),
					)
					return
				},
				Description: "DNS record Type",
			},

			pdnsRdata: {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: func(val interface{}, field string) (warnings []string, errors []error) {
					value := val.(string)
					if ipv6Regexp.MatchString(value) && upcaseRegexp.MatchString(value) {
						errors = append(
							errors,
							fmt.Errorf(
								"IPv6 addresses in the data property cannot have upper case letters: %s",
								value,
							),
						)
					}
					return
				},
				Description: "DNS record Data",
			},

			pdnsRecordTTL: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  900,
				DefaultFunc: func() (interface{}, error) {
					return 900, nil
				},
				Description: "DNS record TTL",
			},

			pdnsMxPreference: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "DNS maximum preference",
			},

			pdnsSrvPort: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "DNS server Port",
			},

			pdnsSrvPriority: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "DNS server Priority",
			},

			pdnsSrvWeight: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "DNS server weight",
			},

			pdnsSrvService: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service info",
			},

			pdnsSrvProtocol: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protocol",
			},

			pdnsRecordCreatedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation Data",
			},

			pdnsRecordModifiedOn: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Modification date",
			},
		},
	}
}

func resourceIBMPrivateDNSResourceRecordCreate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	var (
		instanceID string
		zoneID     string
		recordType string
		name       string
		rdata      string
		service    string
		protocol   string
		ttl        int
		preference int
		port       int
		priority   int
		weight     int
	)

	instanceID = d.Get(pdnsInstanceID).(string)
	zoneID = d.Get(pdnsZoneID).(string)
	recordType = d.Get(pdnsRecordType).(string)
	name = d.Get(pdnsRecordName).(string)
	rdata = d.Get(pdnsRdata).(string)

	if v, ok := d.GetOk(pdnsRecordTTL); ok {
		ttl = v.(int)
	}

	createResourceRecordOptions := sess.NewCreateResourceRecordOptions(instanceID, zoneID)
	createResourceRecordOptions.SetName(name)
	createResourceRecordOptions.SetType(recordType)
	createResourceRecordOptions.SetTTL(int64(ttl))

	switch recordType {
	case "A":
		resourceRecordAData, err := sess.NewResourceRecordInputRdataRdataARecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordAData)
	case "AAAA":
		resourceRecordAaaaData, err := sess.NewResourceRecordInputRdataRdataAaaaRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordAaaaData)
	case "CNAME":
		resourceRecordCnameData, err := sess.NewResourceRecordInputRdataRdataCnameRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordCnameData)
	case "PTR":
		resourceRecordPtrData, err := sess.NewResourceRecordInputRdataRdataPtrRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordPtrData)
	case "TXT":
		resourceRecordTxtData, err := sess.NewResourceRecordInputRdataRdataTxtRecord(rdata)
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordTxtData)
	case "MX":
		if v, ok := d.GetOk(pdnsMxPreference); ok {
			preference = v.(int)
		}
		resourceRecordMxData, err := sess.NewResourceRecordInputRdataRdataMxRecord(rdata, int64(preference))
		if err != nil {
			return err
		}
		createResourceRecordOptions.SetRdata(resourceRecordMxData)
	case "SRV":
		if v, ok := d.GetOk(pdnsSrvPort); ok {
			port = v.(int)
		}
		if v, ok := d.GetOk(pdnsSrvPriority); ok {
			priority = v.(int)
		}
		if v, ok := d.GetOk(pdnsSrvWeight); ok {
			weight = v.(int)
		}
		resourceRecordSrvData, err := sess.NewResourceRecordInputRdataRdataSrvRecord(int64(port), int64(priority), rdata, int64(weight))
		if err != nil {
			return err
		}
		if v, ok := d.GetOk(pdnsSrvService); ok {
			service = v.(string)
		}
		if v, ok := d.GetOk(pdnsSrvProtocol); ok {
			protocol = v.(string)
		}
		createResourceRecordOptions.SetRdata(resourceRecordSrvData)
		createResourceRecordOptions.SetService(service)
		createResourceRecordOptions.SetProtocol(protocol)
	}
	mk := "private_dns_resource_record_" + instanceID + zoneID
	ibmMutexKV.Lock(mk)
	defer ibmMutexKV.Unlock(mk)
	response, detail, err := sess.CreateResourceRecord(createResourceRecordOptions)
	if err != nil {
		log.Printf("Error creating dns record:%s", detail)
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", instanceID, zoneID, *response.ID))
	d.Set(pdnsResourceRecordID, *response.ID)

	return resourceIBMPrivateDNSResourceRecordRead(d, meta)
}

func resourceIBMPrivateDNSResourceRecordRead(d *schema.ResourceData, meta interface{}) error {
	id_set := strings.Split(d.Id(), "/")
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	response, detail, err := sess.GetResourceRecord(getResourceRecordOptions)
	if err != nil {
		log.Printf("Error reading dns record:%s", detail)
		return err
	}

	// extract the record name by removing zone details
	var recordName string
	zone := strings.Split(id_set[1], ":")
	name := strings.Split(*response.Name, zone[0])
	name[0] = strings.Trim(name[0], ".")
	recordName = name[0]

	if *response.Type == "SRV" {
		// "_sip._udp.testsrv"
		temp := strings.Split(name[0], ".")
		recordName = temp[2]
	}

	d.Set("id", response.ID)
	d.Set(pdnsResourceRecordID, response.ID)
	d.Set(pdnsInstanceID, id_set[0])
	d.Set(pdnsZoneID, id_set[1])
	d.Set(pdnsRecordName, recordName)
	d.Set(pdnsRdata, response.Rdata)
	d.Set(pdnsRecordType, response.Type)
	d.Set(pdnsRecordTTL, response.TTL)
	d.Set(pdnsRecordCreatedOn, response.CreatedOn)
	d.Set(pdnsRecordModifiedOn, response.ModifiedOn)

	if *response.Type == "SRV" {
		data := response.Rdata.(map[string]interface{})
		d.Set(pdnsSrvPort, data["port"])
		d.Set(pdnsSrvPriority, data["priority"])
		d.Set(pdnsSrvWeight, data["weight"])
		d.Set(pdnsRdata, data["target"])
		d.Set(pdnsSrvService, response.Service)
		d.Set(pdnsSrvProtocol, response.Protocol)
	}

	if *response.Type == "MX" {
		data := response.Rdata.(map[string]interface{})
		d.Set(pdnsMxPreference, data["preference"])
		d.Set(pdnsRdata, data["exchange"])
	}

	return nil
}

func resourceIBMPrivateDNSResourceRecordUpdate(d *schema.ResourceData, meta interface{}) error {
	id_set := strings.Split(d.Id(), "/")

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	mk := "private_dns_resource_record_" + id_set[0] + id_set[1]
	ibmMutexKV.Lock(mk)
	defer ibmMutexKV.Unlock(mk)
	response, detail, err := sess.GetResourceRecord(getResourceRecordOptions)
	if err != nil {
		log.Printf("Error updating dns record:%s", detail)
		return err
	}

	updateResourceRecordOptions := sess.NewUpdateResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	recordName := d.Get(pdnsRecordName).(string)
	if *response.Type != "PTR" {
		updateResourceRecordOptions.SetName(recordName)
	}

	//
	var ttl int64
	var rdata string

	temp := d.Get(pdnsRecordTTL).(int)
	ttl = int64(temp)

	recordType := *response.Type
	switch recordType {
	case "A":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordAData, err := sess.NewResourceRecordUpdateInputRdataRdataARecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordAData)
		}
	case "AAAA":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordAaaaData, err := sess.NewResourceRecordUpdateInputRdataRdataAaaaRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordAaaaData)
		}
	case "CNAME":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordCnameData, err := sess.NewResourceRecordUpdateInputRdataRdataCnameRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordCnameData)
		}
	case "PTR":
		if d.HasChange(pdnsRecordTTL) {
			updateResourceRecordOptions.SetTTL(ttl)
		}
	case "TXT":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) {
			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			resourceRecordTxtData, err := sess.NewResourceRecordUpdateInputRdataRdataTxtRecord(rdata)
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordTxtData)
		}
	case "MX":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) ||
			d.HasChange(pdnsMxPreference) {

			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			preference := d.Get(pdnsMxPreference).(int)

			resourceRecordMxData, err := sess.NewResourceRecordUpdateInputRdataRdataMxRecord(rdata, int64(preference))
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordMxData)
		}
	case "SRV":
		if d.HasChange(pdnsRecordTTL) || d.HasChange(pdnsRdata) ||
			d.HasChange(pdnsSrvPort) || d.HasChange(pdnsSrvPriority) ||
			d.HasChange(pdnsSrvWeight) || d.HasChange(pdnsSrvService) ||
			d.HasChange(pdnsSrvProtocol) {

			updateResourceRecordOptions.SetTTL(ttl)
			rdata = d.Get(pdnsRdata).(string)
			port := d.Get(pdnsSrvPort).(int)
			priority := d.Get(pdnsSrvPriority).(int)
			weight := d.Get(pdnsSrvWeight).(int)

			resourceRecordSrvData, err := sess.NewResourceRecordUpdateInputRdataRdataSrvRecord(int64(port), int64(priority), rdata, int64(weight))
			if err != nil {
				return err
			}
			updateResourceRecordOptions.SetRdata(resourceRecordSrvData)

			service := d.Get(pdnsSrvService).(string)
			protocol := d.Get(pdnsSrvProtocol).(string)
			updateResourceRecordOptions.SetService(service)
			updateResourceRecordOptions.SetProtocol(protocol)
		}
	}

	//
	_, detail, err = sess.UpdateResourceRecord(updateResourceRecordOptions)
	if err != nil {
		log.Printf("Error updating dns record:%s", detail)
		return err
	}

	return resourceIBMPrivateDNSResourceRecordRead(d, meta)
}

func resourceIBMPrivateDNSResourceRecordDelete(d *schema.ResourceData, meta interface{}) error {
	id_set := strings.Split(d.Id(), "/")

	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return err
	}

	deleteResourceRecordOptions := sess.NewDeleteResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	mk := "private_dns_resource_record_" + id_set[0] + id_set[1]
	ibmMutexKV.Lock(mk)
	defer ibmMutexKV.Unlock(mk)
	response, err := sess.DeleteResourceRecord(deleteResourceRecordOptions)
	if err != nil {
		log.Printf("Error deleting dns record:%s", response)
		return err
	}

	d.SetId("")
	return nil
}

func resourceIBMPrivateDNSResourceRecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).PrivateDnsClientSession()
	if err != nil {
		return false, err
	}

	id_set := strings.Split(d.Id(), "/")
	getResourceRecordOptions := sess.NewGetResourceRecordOptions(id_set[0], id_set[1], id_set[2])
	mk := "private_dns_resource_record_" + id_set[0] + id_set[1]
	ibmMutexKV.Lock(mk)
	defer ibmMutexKV.Unlock(mk)
	_, response, err := sess.GetResourceRecord(getResourceRecordOptions)

	if err != nil {
		if response != nil && response.StatusCode == 404 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
