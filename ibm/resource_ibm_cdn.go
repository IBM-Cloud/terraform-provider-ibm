package ibm

import (
	"fmt"
	"log"
	"strconv"

	"github.com/softlayer/softlayer-go/sl"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/services"
)

func resourceIBMCDN() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMCDNCreate,
		Read:   resourceIBMCDNRead,
		Update: resourceIBMCDNUpdate,
		Delete: resourceIBMCDNDelete,
		Exists: resourceIBMCDNExists,
		//Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},
			"vendor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},

			"origin_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},
			"origin_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
			},
			"bucketname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				//ForceNew: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP",
				//ForceNew: true,
			},
			"httpport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
				//ForceNew: true,
			},
			"httpsport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  false,
				//ForceNew: true,
			},
		},
	}
}

func resourceIBMCDNCreate(d *schema.ResourceData, meta interface{}) error {
	//create  session
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("ordering cdn service...")
	//get the value of all the parameters
	domain := d.Get("hostname").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("httpport").(int)
	httpsport := d.Get("httpsport").(int)
	bucketname := d.Get("bucketname").(string)
	//log.Println(httpport)
	//log.Println(httpsport)
	//productOrderService := services.GetProductOrderService(sess.SetRetries(0))
	///creat an object of CDN service
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	//productOrderService := services.GetProductOrderService(sess.SetRetries(0))
	//////pass the parameters to create domain mapping
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		receipt1, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpsport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		log.Println(receipt1)
		log.Println(err)
		d.SetId(*receipt1[0].UniqueId)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		receipt2, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		log.Println(receipt2)
		log.Println(err)
		d.SetId(*receipt2[0].UniqueId)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		receipt3, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpsport),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		log.Println(receipt3)
		log.Println(err)
		d.SetId(*receipt3[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		receipt4, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			OriginType: sl.String(origintype),
		})
		log.Println(receipt4)
		log.Println(err)
		d.SetId(*receipt4[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		receipt5, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpsPort:  sl.Int(httpport),
			OriginType: sl.String(origintype),
		})
		log.Println(receipt5)
		log.Println(err)
		d.SetId(*receipt5[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		receipt6, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			HttpsPort:  sl.Int(httpport),
			OriginType: sl.String(origintype),
		})
		log.Println(receipt6)
		log.Println(err)
		d.SetId(*receipt6[0].UniqueId)
	}
	//if err != nil {
	//	log.Println("[ERROR] hi...")
	//	log.Println(err)
	//}

	// Get account name and set as the Id
	//d.SetId(*CDNAccounts[0].Username)
	//d.Set("domain", *CDNAccounts[0].Username)
	//d.SetId(*receipt[0].UniqueId)
	e := d.Id()
	log.Println("hi...")
	log.Println(e)
	//return resourceIBMCDNDelete(d, meta)
	return nil
}

func resourceIBMCDNRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("reading cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId, err := strconv.Atoi(d.Id())
	log.Println(cdnId)
	///read the changes in the remote resource and update in the local resource.
	read, err := service.VerifyDomainMapping(sl.Int(cdnId))
	log.Println(read)
	log.Println(err)

	return nil
	//return fmt.Errorf("resource_ibm_object_storage_account: Could not find account %s", accountName)
}

func resourceIBMCDNUpdate(d *schema.ResourceData, meta interface{}) error {
	// Nothing to update for now. Not supported.
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Updating cdn service...")
	domain := d.Get("hostname").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("httpport").(int)
	e := d.Id()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	//pass the changed as well as unchanged parameters to update the resource.
	update, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
		Origin:     sl.String(originaddress),
		VendorName: sl.String(vendorname),
		Domain:     sl.String(domain),
		Protocol:   sl.String(protocol),
		HttpPort:   sl.Int(httpport),
		OriginType: sl.String(origintype),
		UniqueId:   sl.String(e),
	})
	log.Println(update)
	log.Println(err)
	//update := service.UpdateDomainMapping()
	return nil
}

func resourceIBMCDNDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Deleting cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)

	log.Printf("[INFO] Deleting Domain Mapping:")
	cdnId := sl.String(d.Id())
	log.Println(cdnId)
	///pass the id to delete the resource.
	delete, err := service.DeleteDomainMapping(cdnId)
	if err != nil {
		log.Println("error destroying %s", err)
		log.Println(err)
	}
	d.SetId("")

	log.Println(delete)
	//receipt, err := service.DeleteDomainMapping()

	// Delete is not supported for now.
	return nil
}

func resourceIBMCDNExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Exists cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId := sl.String(d.Id())
	//if err != nil {
	//	return false, fmt.Errorf("Not a valid Id, must be an integer: %s", err)
	//}
	log.Println(cdnId)
	///check if the resource exists with the given id.
	exists, err := service.ListDomainMappingByUniqueId(cdnId)
	log.Println(exists)
	//err := resourceIBMCDNRead(d, meta)
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error retrieving CDN mapping info: %s", err)
	}
	log.Println("id found")
	return true, nil
	//return false, nil
}
