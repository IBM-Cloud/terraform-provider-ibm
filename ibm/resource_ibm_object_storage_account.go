package ibm

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/softlayer/softlayer-go/datatypes"
	"github.com/softlayer/softlayer-go/filter"
	"github.com/softlayer/softlayer-go/helpers/order"
	"github.com/softlayer/softlayer-go/helpers/product"
	"github.com/softlayer/softlayer-go/services"
	"github.com/softlayer/softlayer-go/sl"
)

const (
	filterPath                  = "networkStorage.billingItem.orderItemId"
	objectMask                  = "id,username,billingItem[id,orderItemId]"
	packageObjectMask           = "id,name,description,isActive,type[keyName]"
	objectStoragePackageKeyname = "OBJECT_STORAGE"
	objectStorageMask           = "id,capacity,description,units,keyName,prices[id,categories[id,name,categoryCode]]"
	s3                          = "CLOUD_OBJECT_STORAGE"
	swift                       = "OBJECT_STORAGE_PAY_AS_YOU_GO"
)

func resourceIBMObjectStorageAccount() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMObjectStorageAccountCreate,
		Read:     resourceIBMObjectStorageAccountRead,
		Update:   resourceIBMObjectStorageAccountUpdate,
		Delete:   resourceIBMObjectStorageAccountDelete,
		Exists:   resourceIBMObjectStorageAccountExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accountType": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "SWIFT",
				Optional: true,
			},
			"local_note": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func resourceIBMObjectStorageAccountCreate(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	service := services.GetProductPackageService(sess)
	packages, err := service.
		Mask(packageObjectMask).
		Filter(
			filter.Build(
				filter.Path("keyName").Eq(objectStoragePackageKeyname),
			),
		).
		Limit(1).
		GetAllObjects()
	if err != nil {
		return fmt.Errorf(
			"resource_ibm_object_storage_account: Error ordering object storage account: %s", err)
	}
	ObjectStorages, err := product.GetPackageProducts(sess, *packages[0].Id, objectStorageMask)
	if err != nil {
		return fmt.Errorf(
			"resource_ibm_object_storage_account: Error ordering object storage account: %s", err)
	}

	if len(ObjectStorages) == 0 {
		return fmt.Errorf(
			"resource_ibm_object_storage_account: Error getting Package Data: %s", err)
	}
	// Order the account
	productOrderService := services.GetProductOrderService(sess.SetRetries(0))

	itemPriceID := sl.Int(*ObjectStorages[0].Prices[0].Id)
	accountType := d.Get("accountType").(string)
	keyName := swift
	switch accountType {
	case "SWIFT":
		keyName = swift
	case "S3":
		keyName = s3
	default:
		return fmt.Errorf("Error during creation of storage: Invalid accountType %s", accountType)
	}

	for _, ObjectStorage := range ObjectStorages {
		if *ObjectStorage.KeyName == keyName {
			itemPriceID = sl.Int(*ObjectStorage.Prices[0].Id)
		}
	}

	productOrderContainer := datatypes.Container_Product_Order{
		PackageId: sl.Int(*packages[0].Id),
		Prices: []datatypes.Product_Item_Price{
			{
				Id: itemPriceID,
			},
		},
		Quantity: sl.Int(1),
	}
	receipt, err := productOrderService.PlaceOrder(&productOrderContainer, sl.Bool(false))
	if err != nil {
		return fmt.Errorf(
			"resource_ibm_object_storage_account: Error ordering object storage: %s", err)
	}

	// Wait for the object storage account order to complete.
	_, err = WaitForOrderCompletion(&receipt, meta)
	if err != nil {
		return fmt.Errorf(
			"Error waiting for object storage account order (%d) to complete: %s", receipt.OrderId, err)
	}

	objectStorageAccounts, err := services.GetAccountService(sess).
		Filter(filter.Build(
			filter.Path(filterPath).
				Eq(strconv.Itoa(*receipt.PlacedOrder.Items[0].Id)))).
		Mask(objectMask).GetNetworkStorage()

	if err != nil {
		return fmt.Errorf("resource_ibm_object_storage_account: Error on retrieving new: %s", err)
	}

	if len(objectStorageAccounts) != 1 {
		return fmt.Errorf("resource_ibm_object_storage_account: Expected one object storage account.")
	}

	d.SetId(fmt.Sprintf("%d", *objectStorageAccounts[0].Id))
	d.Set("username", *objectStorageAccounts[0].Username)
	log.Printf("[INFO] Storage Account ID: %s", d.Id())

	return resourceIBMObjectStorageAccountRead(d, meta)
}

func WaitForOrderCompletion(
	receipt *datatypes.Container_Product_Order_Receipt, meta interface{}) (datatypes.Billing_Order_Item, error) {

	log.Printf("Waiting for billing order %d to have zero active transactions", receipt.OrderId)
	var billingOrderItem *datatypes.Billing_Order_Item

	stateConf := &resource.StateChangeConf{
		Pending: []string{"", "in progress"},
		Target:  []string{"complete"},
		Refresh: func() (interface{}, string, error) {
			var err error
			var completed bool

			completed, billingOrderItem, err = order.CheckBillingOrderComplete(meta.(ClientSession).SoftLayerSession(), receipt)
			if err != nil {
				return nil, "", err
			}

			if completed {
				return billingOrderItem, "complete", nil
			} else {
				return billingOrderItem, "in progress", nil
			}
		},
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	_, err := stateConf.WaitForState()
	return *billingOrderItem, err
}

func resourceIBMObjectStorageAccountRead(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	storageAccountID, _ := strconv.Atoi(d.Id())

	objectStorageAccounts, err := services.GetAccountService(sess).
		Filter(filter.Build(
			filter.Path("id").Eq(storageAccountID))).
		Mask(objectMask).GetNetworkStorage()

	if err != nil {
		return fmt.Errorf("resource_ibm_object_storage_account: Error on Read: %s", err)
	}

	for _, objectStorageAccount := range objectStorageAccounts {
		if *objectStorageAccount.Id == storageAccountID {
			log.Printf("[INFO] Found Storage Account ID: %s", d.Id())
			d.Set("username", *objectStorageAccount.Username)
			return nil
		}
	}

	return fmt.Errorf("resource_ibm_object_storage_account: Could not find account %s", d.Id())
}

func resourceIBMObjectStorageAccountUpdate(d *schema.ResourceData, meta interface{}) error {
	// Nothing to update for now. Not supported.
	return nil
}

func resourceIBMObjectStorageAccountDelete(d *schema.ResourceData, meta interface{}) error {
	sess := meta.(ClientSession).SoftLayerSession()
	storageService := services.GetNetworkStorageService(sess)
	storageID, _ := strconv.Atoi(d.Id())

	// Get billing item associated with the object storage account
	billingItem, err := storageService.Id(storageID).GetBillingItem()

	if err != nil {
		return fmt.Errorf("Error while looking up billing item associated with the object storage account: %s", err)
	}

	if billingItem.Id == nil {
		return fmt.Errorf("Error while looking up billing item associated with the object storage account: No billing item for ID:%d", storageID)
	}

	success, err := services.GetBillingItemService(sess).Id(*billingItem.Id).CancelService()
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("SoftLayer reported an unsuccessful cancellation")
	}
	return nil
}

func resourceIBMObjectStorageAccountExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	err := resourceIBMObjectStorageAccountRead(d, meta)
	if err != nil {
		if strings.Contains(err.Error(), "Could not find account") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
