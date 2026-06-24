// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"crypto/hmac"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/crypto/sha3"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SuppressEquivalentJSON(k, old, new string, d *schema.ResourceData) bool {

	if old == "" {
		return false
	}
	var oldObj, newObj []map[string]interface{}
	err := json.Unmarshal([]byte(old), &oldObj)
	if err != nil {
		log.Printf("Error unmarshalling old json :: %s", err.Error())
		return false
	}
	err = json.Unmarshal([]byte(new), &newObj)
	if err != nil {
		log.Printf("Error unmarshalling new json :: %s", err.Error())
		return false
	}

	oldm := make(map[interface{}]interface{})
	newm := make(map[interface{}]interface{})

	for _, m := range oldObj {
		oldm[m["key"]] = m["value"]
	}
	for _, m := range newObj {
		newm[m["key"]] = m["value"]
	}
	return reflect.DeepEqual(oldm, newm)
}

func SuppressHashedRawSecret(k, old, new string, d *schema.ResourceData) bool {
	if len(d.Id()) == 0 {
		return false
	}
	isSecretRef, _ := regexp.MatchString("[{]{1}(\\b(vault)\\b[:]{2}([ a-zA-Z0-9_-]*)[.]{0,1}(.*))[}]{1}", new)
	if isSecretRef {
		return false
	}
	parts, _ := SepIdParts(d.Id(), "/")
	secret := parts[1]
	mac := hmac.New(sha3.New512, []byte(secret))
	mac.Write([]byte(new))
	secureHmac := hex.EncodeToString(mac.Sum(nil))
	return cmp.Equal(strings.Join([]string{"hash", "SHA3-512", secureHmac}, ":"), old)
}

func SuppressPipelinePropertyRawSecret(k, old, new string, d *schema.ResourceData) bool {
	// ResourceIBMCdTektonPipelineProperty
	if d.Get("type").(string) == "secure" {
		segs := []string{d.Get("pipeline_id").(string), d.Get("name").(string)}
		secret := strings.Join(segs, ".")
		mac := hmac.New(sha3.New512, []byte(secret))
		mac.Write([]byte(new))
		secureHmac := hex.EncodeToString(mac.Sum(nil))
		return cmp.Equal(strings.Join([]string{"hash", "SHA3-512", secureHmac}, ":"), old)
	} else {
		return old == new
	}
}

func SuppressTriggerPropertyRawSecret(k, old, new string, d *schema.ResourceData) bool {
	// ResourceIBMCdTektonPipelineTriggerProperty
	if d.Get("type").(string) == "secure" {
		segs := []string{d.Get("pipeline_id").(string), d.Get("trigger_id").(string), d.Get("name").(string)}
		secret := strings.Join(segs, ".")
		mac := hmac.New(sha3.New512, []byte(secret))
		mac.Write([]byte(new))
		secureHmac := hex.EncodeToString(mac.Sum(nil))
		return cmp.Equal(strings.Join([]string{"hash", "SHA3-512", secureHmac}, ":"), old)
	} else {
		return old == new
	}
}

func SuppressGenericWebhookRawSecret(k, old, new string, d *schema.ResourceData) bool {
	// ResourceIBMCdTektonPipelineTrigger
	segs := []string{d.Get("pipeline_id").(string), d.Get("trigger_id").(string)}
	secret := strings.Join(segs, ".")
	mac := hmac.New(sha3.New512, []byte(secret))
	mac.Write([]byte(new))
	secureHmac := hex.EncodeToString(mac.Sum(nil))
	return cmp.Equal(strings.Join([]string{"hash", "SHA3-512", secureHmac}, ":"), old)
}

func SuppressTriggerEvents(key, oldValue, newValue string, d *schema.ResourceData) bool {
	// The key is a path not the list itself, e.g. "events.0"
	lastDotIndex := strings.LastIndex(key, ".")
	if lastDotIndex != -1 {
		key = string(key[:lastDotIndex])
	}
	oldData, newData := d.GetChange(key)
	if oldData == nil || newData == nil {
		return false
	}
	oldArray := oldData.([]interface{})
	newArray := newData.([]interface{})
	if len(oldArray) != len(newArray) {
		// Items added or removed, always detect as changed
		return false
	}

	// Convert data to string arrays
	oldEvents := make([]string, len(oldArray))
	newEvents := make([]string, len(newArray))
	for i, oldEvt := range oldArray {
		oldEvents[i] = fmt.Sprint(oldEvt)
	}
	for j, newEvt := range newArray {
		newEvents[j] = fmt.Sprint(newEvt)
	}
	// Ensure consistent sorting before comparison, to suppress unnecessary change detections
	sort.Strings(oldEvents)
	sort.Strings(newEvents)
	return reflect.DeepEqual(oldEvents, newEvents)
}

func SuppressAllowBlank(k, old, new string, d *schema.ResourceData) bool {
	if new == "" && old != "" {
		return true
	}
	return false
}

func SuppressRestrictUserDomains(key, old, new string, d *schema.ResourceData) bool {
	// Special case: force diff when restrict_invitation or when invitation_email_allow_patterns are removed

	// Extract the field from the key (e.g. "restrict_user_domains.0.restrict_invitation" -> "restrict_invitation")
	parts := strings.Split(key, ".")

	if len(parts) > 2 {
		field := parts[2]
		if field == "restrict_invitation" {
			// Edge case when removing restrict_invitation, since d.HasChange returns false.
			// Force a diff when old and new dont match
			if old == "true" && new == "false" {
				return false
			}
		}
		if field == "invitation_email_allow_patterns" && len(parts) > 3 {
			// Get the hash/index from the key: "restrict_user_domains.0.invitation_email_allow_patterns.#"
			hash := parts[3]
			if hash == "#" && new == "0" {
				// # indicates we are assessing the size of the list
				// If the new patterns list has size 0 then it has been removed, so force a diff
				return false
			}
		}
	}

	// default HasChange handling for all other cases
	return !d.HasChange(key)
}
