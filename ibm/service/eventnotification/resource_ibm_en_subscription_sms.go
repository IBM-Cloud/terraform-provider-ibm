// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func ResourceIBMEnSMSSubscription() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMEnSMSSubscriptionCreate,
		ReadContext:   resourceIBMEnSMSSubscriptionRead,
		UpdateContext: resourceIBMEnSMSSubscriptionUpdate,
		DeleteContext: resourceIBMEnSMSSubscriptionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription name.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subscription description.",
			},
			"destination_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Destination ID.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Topic ID.",
			},
			"attributes": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"invited": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "The phone numbers to invite. Add a number by adding it to this list; remove a number by removing it from this list.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"subscribed": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Phone numbers that have accepted the invitation and are currently subscribed. Populated by the service; read-only.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"unsubscribed": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Phone numbers that have unsubscribed. Populated by the service; read-only.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"add": {
							Type:       schema.TypeList,
							Optional:   true,
							Computed:   true,
							Deprecated: "Use invited to manage phone numbers.",
							Elem:       &schema.Schema{Type: schema.TypeString},
						},
						"remove": {
							Type:       schema.TypeList,
							Optional:   true,
							Computed:   true,
							Deprecated: "Use invited to manage phone numbers.",
							Elem:       &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"subscription_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Subscription ID.",
			},
			"destination_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of Destination.",
			},
			"destination_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Destination name.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the topic.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last updated time.",
			},
		},
	}
}

func resourceIBMEnSMSSubscriptionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.CreateSubscriptionOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))

	options.SetName(d.Get("name").(string))
	options.SetTopicID(d.Get("topic_id").(string))
	options.SetDestinationID(d.Get("destination_id").(string))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}

	attributes, _ := SMSattributesMapToAttributes(d.Get("attributes.0").(map[string]interface{}))
	options.SetAttributes(&attributes)

	result, _, err := enClient.CreateSubscriptionWithContext(context, options)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateSubscriptionWithContext failed: %s", err.Error()), "ibm_en_subscription_sms", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *options.InstanceID, *result.ID))

	return resourceIBMEnSMSSubscriptionRead(context, d, meta)
}

func resourceIBMEnSMSSubscriptionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.GetSubscriptionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "read")
		return tfErr.GetDiag()
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	result, response, err := enClient.GetSubscriptionWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSubscriptionWithContext failed: %s", err.Error()), "ibm_en_subscription_sms", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_guid", options.InstanceID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting instance_guid: %s", err))
	}

	if err = d.Set("subscription_id", result.ID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting instance_guid: %s", err))
	}

	if err = d.Set("name", result.Name); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting name: %s", err))
	}

	if result.Description != nil {
		if err = d.Set("description", result.Description); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting description: %s", err))
		}
	}

	if result.From != nil {
		if err = d.Set("from", result.From); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting from: %s", err))
		}
	}

	if err = d.Set("destination_id", result.DestinationID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting destination_id: %s", err))
	}

	if err = d.Set("destination_type", result.DestinationType); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting destination_type: %s", err))
	}

	if result.DestinationName != nil {
		if err = d.Set("destination_name", result.DestinationName); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting destination_name: %s", err))
		}
	}

	if err = d.Set("topic_id", result.TopicID); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting topic_id: %s", err))
	}

	if result.TopicName != nil {
		if err = d.Set("topic_name", result.TopicName); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting topic_name: %s", err))
		}
	}

	if err = d.Set("updated_at", result.UpdatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error setting updated_at: %s", err))
	}

	if result.Attributes != nil {
		if err = d.Set("attributes", enSMSSubscriptionResourceFlattenAttributes(result.Attributes, d)); err != nil {
			return diag.FromErr(fmt.Errorf("[ERROR] Error setting attributes: %s", err))
		}
	}

	return nil
}

func resourceIBMEnSMSSubscriptionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.UpdateSubscriptionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "update")
		return tfErr.GetDiag()
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	if ok := d.HasChanges("name", "description", "attributes"); ok {
		options.SetName(d.Get("name").(string))

		if _, ok := d.GetOk("description"); ok {
			options.SetDescription(d.Get("description").(string))
		}

		attributes := SMSattributesUpdateFromResourceData(d)
		options.SetAttributes(&attributes)

		_, _, err := enClient.UpdateSubscriptionWithContext(context, options)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateSubscriptionWithContext failed: %s", err.Error()), "ibm_en_subscription_sms", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		return resourceIBMEnSMSSubscriptionRead(context, d, meta)
	}

	return nil
}

func resourceIBMEnSMSSubscriptionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	options := &en.DeleteSubscriptionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_en_subscription_sms", "delete")
		return tfErr.GetDiag()
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	response, err := enClient.DeleteSubscriptionWithContext(context, options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteSubscriptionWithContext: failed: %s", err.Error()), "ibm_en_subscription_sms", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func SMSattributesMapToAttributes(attributeMap map[string]interface{}) (en.SubscriptionCreateAttributes, en.SubscriptionUpdateAttributesSmsUpdateAttributes) {
	attributesCreate := en.SubscriptionCreateAttributes{}
	attributesUpdate := en.SubscriptionUpdateAttributesSmsUpdateAttributes{}

	if attributeMap["invited"] != nil {
		to := []string{}
		for _, toItem := range attributeMap["invited"].([]interface{}) {
			to = append(to, toItem.(string))
		}
		attributesCreate.Invited = to
	}

	return attributesCreate, attributesUpdate
}

// SMSattributesUpdateFromResourceData computes the update attributes for SMS subscriptions
// using a set-diff on the invited list so that order changes never trigger a spurious API call.
func SMSattributesUpdateFromResourceData(d *schema.ResourceData) en.SubscriptionUpdateAttributesSmsUpdateAttributes {
	attributesUpdate := en.SubscriptionUpdateAttributesSmsUpdateAttributes{}

	oldInvitedRaw, newInvitedRaw := d.GetChange("attributes.0.invited")
	oldInvited := oldInvitedRaw.([]interface{})
	newInvited := newInvitedRaw.([]interface{})

	oldSet := toStringSet(oldInvited)
	newSet := toStringSet(newInvited)

	var invitedAdd, invitedRemove []string
	for _, v := range newInvited {
		if _, exists := oldSet[v.(string)]; !exists {
			invitedAdd = append(invitedAdd, v.(string))
		}
	}
	for _, v := range oldInvited {
		if _, exists := newSet[v.(string)]; !exists {
			invitedRemove = append(invitedRemove, v.(string))
		}
	}

	if len(invitedAdd) > 0 || len(invitedRemove) > 0 {
		invited := &en.UpdateAttributesInvited{}
		if len(invitedAdd) > 0 {
			invited.Add = invitedAdd
		}
		if len(invitedRemove) > 0 {
			invited.Remove = invitedRemove
		}
		attributesUpdate.Invited = invited
	}

	return attributesUpdate
}

// add and remove are always written as empty slices — they are deprecated transient fields.
func enSMSSubscriptionResourceFlattenAttributes(result en.SubscriptionAttributesIntf, d *schema.ResourceData) []map[string]interface{} {
	attributes := result.(*en.SubscriptionAttributes)

	attrMap := map[string]interface{}{
		"add":    []string{},
		"remove": []string{},
	}

	// Build a set of all phone numbers known on the service side.
	knownOnService := map[string]struct{}{}
	for _, item := range attributes.Invited {
		if item.PhoneNumber != nil {
			knownOnService[*item.PhoneNumber] = struct{}{}
		}
	}
	for _, item := range attributes.Subscribed {
		if item.PhoneNumber != nil {
			knownOnService[*item.PhoneNumber] = struct{}{}
		}
	}
	for _, item := range attributes.Unsubscribed {
		if item.PhoneNumber != nil {
			knownOnService[*item.PhoneNumber] = struct{}{}
		}
	}

	// Write invited in config order to keep state == config order for TypeList.
	seen := map[string]struct{}{}
	invited := []string{}

	if v, ok := d.GetOk("attributes.0.invited"); ok {
		for _, e := range v.([]interface{}) {
			phone := e.(string)
			if _, onService := knownOnService[phone]; onService {
				if _, already := seen[phone]; !already {
					seen[phone] = struct{}{}
					invited = append(invited, phone)
				}
			}
		}
	}

	// Append pending-invite numbers not in config (e.g. invited outside Terraform).
	for _, item := range attributes.Invited {
		if item.PhoneNumber != nil {
			if _, already := seen[*item.PhoneNumber]; !already {
				seen[*item.PhoneNumber] = struct{}{}
				invited = append(invited, *item.PhoneNumber)
			}
		}
	}

	attrMap["invited"] = invited

	// Populate subscribed — read-only, directly from the API response.
	subscribed := []string{}
	for _, item := range attributes.Subscribed {
		if item.PhoneNumber != nil {
			subscribed = append(subscribed, *item.PhoneNumber)
		}
	}
	attrMap["subscribed"] = subscribed

	// Populate unsubscribed — read-only, directly from the API response.
	unsubscribed := []string{}
	for _, item := range attributes.Unsubscribed {
		if item.PhoneNumber != nil {
			unsubscribed = append(unsubscribed, *item.PhoneNumber)
		}
	}
	attrMap["unsubscribed"] = unsubscribed

	return []map[string]interface{}{attrMap}
}
