package ibm

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

func resourceIBMDNSREVERSERecord() *schema.Resource {
	return &schema.Resource{
		Exists:   resourceIBMDNSREVERSERecordExists,
		Create:   resourceIBMDNSREVERSERecordCreate,
		Read:     resourceIBMDNSREVERSERecordRead,
		Update:   resourceIBMDNSREVERSERecordUpdate,
		Delete:   resourceIBMDNSREVERSERecordDelete,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{
			"dns_ipaddress": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dns_hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dns_ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return 604800, nil
				},
			},
		},
	}
}

//  Creates DNS Domain Reverse Record
//  https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain/CreatePtrRecord
func resourceIBMDNSREVERSERecordCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetDnsDomainService(sess.SetRetries(0))
	Data := sl.String(d.Get("dns_hostname").(string))
	Ttl := sl.Int(d.Get("dns_ttl").(int))
	Ipaddress := sl.String(d.Get("dns_ipaddress").(string))
	var err error
	var id int
	var record datatypes.Dns_Domain_ResourceRecord
	record, err = service.CreatePtrRecord(Ipaddress, Data, Ttl)
	if record.Id != nil {
		log.Printf("record Id is not null")
		id = *record.Id
	}
	log.Printf("id is %d", id)
	if err != nil {
		return fmt.Errorf("Error creating DNS Reverse %s", err)
	}
	d.SetId(fmt.Sprintf("%d", id))
	log.Printf("[INFO] Dns Reverse %s ", d.Id())
	return resourceIBMDNSREVERSERecordRead(d, meta)
}

//  Reads DNS Domain Reverse Record from SL system
//  https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain_ResourceRecord/getObject
func resourceIBMDNSREVERSERecordRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetDnsDomainResourceRecordService(sess)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	log.Printf("id of service inside read %d", id)
	result, err := service.Id(id).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving DNS Reverse Record: %s", err)
	}
	fmt.Println(result.Id)
	return nil
}

//  Updates DNS Domain Reverse Record in SL system
//  https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain_ResourceRecord/editObject
func resourceIBMDNSREVERSERecordUpdate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetDnsDomainResourceRecordService(sess)
	serviceNoRetry := services.GetDnsDomainResourceRecordService(sess.SetRetries(0))
	recordId, _ := strconv.Atoi(d.Id())
	record, err := service.Id(recordId).GetObject()
	if err != nil {
		return fmt.Errorf("Error retrieving DNS Reverse Record: %s", err)
	}
	if data, ok := d.GetOk("dns_hostname"); ok && d.HasChange("dns_hostname") {
		record.Data = sl.String(data.(string))
	}
	if ttl, ok := d.GetOk("dns_ttl"); ok && d.HasChange("dns_ttl") {
		record.Ttl = sl.Int(ttl.(int))
	}
	record.IsGatewayAddress = nil
	_, err = serviceNoRetry.Id(recordId).EditObject(&record)
	if err != nil {
		return fmt.Errorf("Error editing DNS Reverse  Record %d: %s", recordId, err)
	}
	return nil
}

//  Deletes DNS Domain Reverse Record in SL system
//  https://sldn.softlayer.com/reference/services/SoftLayer_Dns_Domain_ResourceRecord/deleteObject
func resourceIBMDNSREVERSERecordDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetDnsDomainResourceRecordService(sess)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	_, err = service.Id(id).DeleteObject()

	if err != nil {
		return fmt.Errorf("Error deleting DNS Reverse Record: %s", err)
	}
	return nil
}

// Exists function is called by refresh
// if the entity is absent - it is deleted from the .tfstate file
func resourceIBMDNSREVERSERecordExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetDnsDomainResourceRecordService(sess)
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, fmt.Errorf("Not a valid ID, must be an integer: %s", err)
	}
	record, err := service.Id(id).GetObject()
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error retrieving domain reverse record info: %s", err)
	}
	return record.Id != nil && *record.Id == id, nil
}
