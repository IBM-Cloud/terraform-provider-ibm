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

const str string = ".cdnedge.bluemix.net"

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
				Optional: true,
				Default:  "HOST_SERVER",
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"httpsport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  443,
			},
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"respectheaders": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"fileextension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"certificatetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"cachekeyqueryrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "include-all",
			},
			"performanceconfiguration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "General web delivery",
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},
		},
	}
}

func resourceIBMCDNCreate(d *schema.ResourceData, meta interface{}) error {
	///create  session
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("ordering cdn service...")
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
	header := d.Get("header").(string)
	cachekeyqueryrule := d.Get("cachekeyqueryrule").(string)
	performanceconfiguration := d.Get("performanceconfiguration").(string)
	respectheaders := d.Get("respectheaders").(bool)
	cname := d.Get("cname").(string)
	certificateType := d.Get("certificatetype").(string)
	if cname != "0" {
		cname = cname + str
	}
	if cname == "0" {
		cname = ""
	}
	///creat an object of CDN service
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	//////pass the parameters to create domain mapping
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		receipt1, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt1)
		d.SetId(*receipt1[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result1, err := service.VerifyDomainMapping(&id)
		log.Println(result1)
		return resourceIBMCDNRead(d, meta)

	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		receipt2, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt2)
		d.SetId(*receipt2[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result2, err := service.VerifyDomainMapping(&id)
		log.Println(result2)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		receipt3, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt3)
		d.SetId(*receipt3[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result3, err := service.VerifyDomainMapping(&id)
		log.Println(result3)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		receipt4, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt4)
		d.SetId(*receipt4[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result4, err := service.VerifyDomainMapping(&id)
		log.Println(result4)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		receipt5, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt5)
		d.SetId(*receipt5[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result5, err := service.VerifyDomainMapping(&id)
		log.Println(result5)
		return resourceIBMCDNRead(d, meta)
	}
	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		receipt6, err := service.CreateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Cname:                    sl.String(cname),
			Protocol:                 sl.String(protocol),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			Header:                   sl.String(header),
			RespectHeaders:           sl.Bool(respectheaders),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		if err != nil {
			return fmt.Errorf("Error creating CDN: %s", err)
		}
		///Print the response of the requested the service.
		log.Print("Response for cdn order")
		log.Println(receipt6)
		d.SetId(*receipt6[0].UniqueId)
		id, err := strconv.Atoi((d.Id()))
		result6, err := service.VerifyDomainMapping(&id)
		log.Println(result6)
		return resourceIBMCDNRead(d, meta)
	}

	return nil
}

func resourceIBMCDNRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	log.Println("reading cdn service...")
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	cdnId := sl.String(d.Id())
	///read the changes in the remote resource and update in the local resource.
	read, err := service.ListDomainMappingByUniqueId(cdnId)
	///Print the response of the requested the service.
	d.Set("id", *read[0].UniqueId)
	d.Set("originaddress", *read[0].OriginHost)
	d.Set("vendorname", *read[0].VendorName)
	d.Set("domain", *read[0].Domain)
	d.Set("header", *read[0].Header)
	d.Set("cname", *read[0].Cname)
	d.Set("origin_type", *read[0].OriginType)
	d.Set("status", *read[0].Status)

	log.Print("Response for cdn verification: ")

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
	httpsport := d.Get("httpsport").(int)
	path := d.Get("path").(string)
	cname := d.Get("cname").(string)
	header := d.Get("header").(string)
	bucketname := d.Get("bucketname").(string)
	fileextension := d.Get("fileextension").(string)
	respectheaders := d.Get("respectheaders").(bool)
	certificateType := d.Get("certificatetype").(string)
	cachekeyqueryrule := d.Get("cachekeyqueryrule").(string)
	performanceconfiguration := d.Get("performanceconfiguration").(string)
	uniqueId := d.Id()
	service := services.GetNetworkCdnMarketplaceConfigurationMappingService(sess)
	///pass the changed as well as unchanged parameters to update the resource.

	if origintype == "HOST_SERVER" && protocol == "HTTP_AND_HTTPS" {
		update1, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update1)

		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "HOST_SERVER" && protocol == "HTTPS" {
		update2, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update2)
		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)

	}

	if origintype == "HOST_SERVER" && protocol == "HTTP" {
		update3, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			Header:                   sl.String(header),
			UniqueId:                 sl.String(uniqueId),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update3)
		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)

	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTP_AND_HTTPS" {
		update4, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update4)
		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTPS" {
		update5, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpsPort:                sl.Int(httpsport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CertificateType:          sl.String(certificateType),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update5)
		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
	}

	if origintype == "OBJECT_STORAGE" && protocol == "HTTP" {
		update6, err := service.UpdateDomainMapping(&datatypes.Container_Network_CdnMarketplace_Configuration_Input{
			Origin:                   sl.String(originaddress),
			VendorName:               sl.String(vendorname),
			Domain:                   sl.String(domain),
			Path:                     sl.String(path),
			Protocol:                 sl.String(protocol),
			Cname:                    sl.String(cname),
			HttpPort:                 sl.Int(httpport),
			OriginType:               sl.String(origintype),
			RespectHeaders:           sl.Bool(respectheaders),
			BucketName:               sl.String(bucketname),
			Header:                   sl.String(header),
			FileExtension:            sl.String(fileextension),
			UniqueId:                 sl.String(uniqueId),
			CacheKeyQueryRule:        sl.String(cachekeyqueryrule),
			PerformanceConfiguration: sl.String(performanceconfiguration),
		})
		///Print the response of the requested service.
		log.Print("Response for cdn update: ")
		log.Println(update6)
		if err != nil {
			log.Println("error updating")
			log.Println(err)
		}
		return resourceIBMCDNRead(d, meta)
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
