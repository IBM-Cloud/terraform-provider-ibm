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

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vendor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"origin_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"origin_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"bucketname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HTTP",
			},
			"httpport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"httpsport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  false,
			},
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceIBMCDNCreate(d *schema.ResourceData, meta interface{}) error {
	///create  session
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("ordering cdn service...")
	str := ".cdnedge.bluemix.net"
	///get the value of all the parameters
	domain := d.Get("hostname").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("httpport").(int)
	httpsport := d.Get("httpsport").(int)
	bucketname := d.Get("bucketname").(string)
	path := d.Get("path").(string)
	cname := d.Get("cname").(string)
	cname = cname + str
	///creat an object of CDN service
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	//////pass the parameters to create domain mapping
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		receipt1, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt1)
		if err != nil {
			log.Print("Error creating : ")
			log.Println(err)
		}
		d.SetId(*receipt1[0].UniqueId)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		receipt2, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt2)
		if err != nil {
			log.Println("error creating")
			log.Println(err)
		}
		d.SetId(*receipt2[0].UniqueId)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		receipt3, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
			BucketName: sl.String(bucketname),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt3)
		if err != nil {
			log.Println("error creating")
			log.Println(err)
		}
		d.SetId(*receipt3[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		receipt4, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			OriginType: sl.String(origintype),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt4)
		if err != nil {
			log.Println("error creating")
			log.Println(err)
		}
		d.SetId(*receipt4[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		receipt5, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt5)
		if err != nil {
			log.Println("error creating")
			log.Println(err)
		}
		d.SetId(*receipt5[0].UniqueId)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		receipt6, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:     sl.String(originaddress),
			VendorName: sl.String(vendorname),
			Domain:     sl.String(domain),
			Path:       sl.String(path),
			Cname:      sl.String(cname),
			Protocol:   sl.String(protocol),
			HttpPort:   sl.Int(httpport),
			HttpsPort:  sl.Int(httpsport),
			OriginType: sl.String(origintype),
		})
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt6)
		if err != nil {
			log.Println("error creating")
			log.Println(err)
		}
		d.SetId(*receipt6[0].UniqueId)
	}
	return nil
}

func resourceIBMCDNRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("reading cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId, err := strconv.Atoi(d.Id())
	///read the changes in the remote resource and update in the local resource.
	read, err := service.VerifyDomainMapping(sl.Int(cdnId))
	///Print the response of the requested the service.
	log.Print("Response for cdn verification: ")
	log.Println(read)
	if err != nil {
		log.Println("error Reading")
		log.Println(err)
	}
	return nil
}

func resourceIBMCDNUpdate(d *schema.ResourceData, meta interface{}) error {
	/// Nothing to update for now. Not supported.
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Updating cdn service...")
	domain := d.Get("hostname").(string)
	vendorname := d.Get("vendor_name").(string)
	origintype := d.Get("origin_type").(string)
	originaddress := d.Get("origin_address").(string)
	protocol := d.Get("protocol").(string)
	httpport := d.Get("httpport").(int)
	uniqueId := d.Id()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	///pass the changed as well as unchanged parameters to update the resource.
	update, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
		Origin:     sl.String(originaddress),
		VendorName: sl.String(vendorname),
		Domain:     sl.String(domain),
		Protocol:   sl.String(protocol),
		HttpPort:   sl.Int(httpport),
		OriginType: sl.String(origintype),
		UniqueId:   sl.String(uniqueId),
	})
	///Print the response of the requested service.
	log.Print("Response for cdn update: ")
	log.Println(update)
	if err != nil {
		log.Println("error updating")
		log.Println(err)
	}
	return nil
}

func resourceIBMCDNDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Deleting cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)

	log.Printf("[INFO] Deleting Domain Mapping:")
	cdnId := sl.String(d.Id())
	///pass the id to delete the resource.
	delete, err := service.DeleteDomainMapping(cdnId)
	if err != nil {
		log.Println("error destroying")
		log.Println(err)
	}
	///print the delete response
	log.Print("Delete response is : ")
	log.Println(delete)
	d.SetId("")
	return nil
}

func resourceIBMCDNExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("Exists cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId := sl.String(d.Id())
	///check if the resource exists with the given id.
	exists, err := service.ListDomainMappingByUniqueId(cdnId)
	///Print the response for exist request.
	log.Print("Exists response is : ")
	log.Println(exists)
	if err != nil {
		if apiErr, ok := err.(sl.Error); ok {
			if apiErr.StatusCode == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error retrieving CDN mapping info: %s", err)
	}
	return true, nil
}
