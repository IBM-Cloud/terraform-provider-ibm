// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"log"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// matchNetworkAttachments pairs state rows (old) with plan rows (new) and
// returns old index -> new index. Unpaired old rows were removed, unpaired
// new rows were added.
//
// Terraform plans list blocks by position, so after a removal from the
// middle a plan row carries the id of the row that used to sit there. Rows
// are therefore paired by name first and by id only for the rest, which
// still lets a rename be an update.
func matchNetworkAttachments(oldList, newList []interface{}) map[int]int {
	matches := map[int]int{}
	newTaken := map[int]bool{}

	key := func(item interface{}, field string) string {
		m, ok := item.(map[string]interface{})
		if !ok {
			return ""
		}
		s, _ := m[field].(string)
		return s
	}

	pair := func(field string) {
		byKey := map[string]int{}
		for j, item := range newList {
			if newTaken[j] {
				continue
			}
			if k := key(item, field); k != "" {
				if _, dup := byKey[k]; !dup {
					byKey[k] = j
				}
			}
		}
		for i, item := range oldList {
			if _, done := matches[i]; done {
				continue
			}
			k := key(item, field)
			if k == "" {
				continue
			}
			if j, ok := byKey[k]; ok && !newTaken[j] {
				matches[i] = j
				newTaken[j] = true
			}
		}
	}

	pair("name")
	pair("id")
	return matches
}

// networkAttachmentShifted reports whether a plan row paired by name carries
// a different attachment's id.
func networkAttachmentShifted(oldItem, newItem interface{}) bool {
	oldMap, ok1 := oldItem.(map[string]interface{})
	newMap, ok2 := newItem.(map[string]interface{})
	if !ok1 || !ok2 {
		return false
	}
	oldID, _ := oldMap["id"].(string)
	newID, _ := newMap["id"].(string)
	return newID != "" && oldID != newID
}

// networkAttachmentFromConfig returns the network_attachments row at idx with
// only the values the configuration sets.
func networkAttachmentFromConfig(d *schema.ResourceData, idx int) map[string]interface{} {
	out, _ := networkAttachmentFromRawConfig(d.GetRawConfig(), idx)
	return out
}

// configuredNetworkAttachmentCount returns the number of network_attachments
// rows in the configuration, or -1 when unknown.
func configuredNetworkAttachmentCount(raw cty.Value) int {
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute("network_attachments") {
		return -1
	}
	list := raw.GetAttr("network_attachments")
	if list.IsNull() {
		return 0
	}
	if !list.IsKnown() || !list.CanIterateElements() {
		return -1
	}
	return list.LengthInt()
}

// networkAttachmentFromRawConfig reads the row at idx from a raw config
// value. The second result is true when a configured value is not known yet.
func networkAttachmentFromRawConfig(raw cty.Value, idx int) (map[string]interface{}, bool) {
	out := map[string]interface{}{}
	unknown := false
	if raw.IsNull() || !raw.IsKnown() || !raw.Type().IsObjectType() || !raw.Type().HasAttribute("network_attachments") {
		return out, false
	}
	list := raw.GetAttr("network_attachments")
	if list.IsNull() || !list.IsKnown() || !list.CanIterateElements() || idx < 0 || idx >= list.LengthInt() {
		return out, false
	}
	row := list.Index(cty.NumberIntVal(int64(idx)))
	if row.IsNull() || !row.IsKnown() || !row.Type().IsObjectType() {
		return out, !row.IsKnown()
	}
	note := func(v cty.Value) bool {
		if !v.IsWhollyKnown() {
			unknown = true
			return true
		}
		return false
	}

	setString := func(dst map[string]interface{}, v cty.Value, name, dstKey string) {
		if !v.Type().HasAttribute(name) {
			return
		}
		a := v.GetAttr(name)
		if a.IsNull() || note(a) || a.Type() != cty.String {
			return
		}
		dst[dstKey] = a.AsString()
	}
	setBool := func(dst map[string]interface{}, v cty.Value, name string) {
		if !v.Type().HasAttribute(name) {
			return
		}
		a := v.GetAttr(name)
		if a.IsNull() || note(a) || a.Type() != cty.Bool {
			return
		}
		dst[name] = a.True()
	}
	setInt := func(dst map[string]interface{}, v cty.Value, name string) {
		if !v.Type().HasAttribute(name) {
			return
		}
		a := v.GetAttr(name)
		if a.IsNull() || note(a) || a.Type() != cty.Number {
			return
		}
		f, _ := a.AsBigFloat().Int64()
		dst[name] = int(f)
	}
	setIntSet := func(dst map[string]interface{}, v cty.Value, name string) {
		if !v.Type().HasAttribute(name) {
			return
		}
		a := v.GetAttr(name)
		if a.IsNull() || note(a) || !a.CanIterateElements() {
			return
		}
		items := []interface{}{}
		for it := a.ElementIterator(); it.Next(); {
			_, ev := it.Element()
			if ev.IsNull() || ev.Type() != cty.Number {
				continue
			}
			f, _ := ev.AsBigFloat().Int64()
			items = append(items, int(f))
		}
		dst[name] = schema.NewSet(schema.HashInt, items)
	}
	setStringSet := func(dst map[string]interface{}, v cty.Value, name string) {
		if !v.Type().HasAttribute(name) {
			return
		}
		a := v.GetAttr(name)
		if a.IsNull() || note(a) || !a.CanIterateElements() {
			return
		}
		items := []interface{}{}
		for it := a.ElementIterator(); it.Next(); {
			_, ev := it.Element()
			if ev.IsNull() || ev.Type() != cty.String {
				continue
			}
			items = append(items, ev.AsString())
		}
		dst[name] = schema.NewSet(schema.HashString, items)
	}

	setString(out, row, "name", "name")
	setString(out, row, "interface_type", "interface_type")
	setInt(out, row, "vlan")
	setBool(out, row, "allow_to_float")
	setIntSet(out, row, "allowed_vlans")

	if row.Type().HasAttribute("virtual_network_interface") {
		vnis := row.GetAttr("virtual_network_interface")
		if !vnis.IsNull() && vnis.IsKnown() && vnis.CanIterateElements() && vnis.LengthInt() > 0 {
			vni := vnis.Index(cty.NumberIntVal(0))
			if !vni.IsNull() && !note(vni) && vni.Type().IsObjectType() {
				vniMap := map[string]interface{}{}
				setString(vniMap, vni, "id", "id")
				setString(vniMap, vni, "name", "name")
				setString(vniMap, vni, "subnet", "subnet")
				setString(vniMap, vni, "protocol_state_filtering_mode", "protocol_state_filtering_mode")
				setBool(vniMap, vni, "allow_ip_spoofing")
				setBool(vniMap, vni, "auto_delete")
				setBool(vniMap, vni, "enable_infrastructure_nat")
				setStringSet(vniMap, vni, "security_groups")
				if vni.Type().HasAttribute("ips") {
					ips := vni.GetAttr("ips")
					if !ips.IsNull() && !note(ips) && ips.CanIterateElements() {
						items := []interface{}{}
						for it := ips.ElementIterator(); it.Next(); {
							_, ev := it.Element()
							if ev.IsNull() || !ev.Type().IsObjectType() || !ev.Type().HasAttribute("reserved_ip") {
								continue
							}
							rip := ev.GetAttr("reserved_ip")
							if rip.IsNull() || rip.Type() != cty.String {
								continue
							}
							items = append(items, map[string]interface{}{"reserved_ip": rip.AsString()})
						}
						vniMap["ips"] = schema.NewSet(hashIpsList, items)
					}
				}
				out["virtual_network_interface"] = []interface{}{vniMap}
			}
		}
	}
	return out, unknown
}

// networkAttachmentMatchesState reports whether every configured value of a
// row equals the value in its state row.
func networkAttachmentMatchesState(cfg, state map[string]interface{}) bool {
	equalSet := func(want *schema.Set, got interface{}) bool {
		switch g := got.(type) {
		case *schema.Set:
			return want.Equal(g)
		case []interface{}:
			return want.Equal(schema.NewSet(want.F, g))
		case nil:
			return want.Len() == 0
		}
		return false
	}
	equalScalar := func(want, got interface{}) bool {
		switch w := want.(type) {
		case string:
			g, ok := got.(string)
			return ok && g == w
		case bool:
			g, ok := got.(bool)
			return ok && g == w
		case int:
			switch g := got.(type) {
			case int:
				return g == w
			case int64:
				return int(g) == w
			}
			return false
		}
		return false
	}
	for key, want := range cfg {
		switch key {
		case "virtual_network_interface":
			continue
		case "allowed_vlans":
			if !equalSet(want.(*schema.Set), state[key]) {
				return false
			}
		default:
			if !equalScalar(want, state[key]) {
				return false
			}
		}
	}
	cfgVNIs, ok := cfg["virtual_network_interface"].([]interface{})
	if !ok || len(cfgVNIs) == 0 {
		return true
	}
	cfgVNI := cfgVNIs[0].(map[string]interface{})
	stateVNIs, _ := state["virtual_network_interface"].([]interface{})
	if len(stateVNIs) == 0 {
		return len(cfgVNI) == 0
	}
	stateVNI, _ := stateVNIs[0].(map[string]interface{})
	for key, want := range cfgVNI {
		switch key {
		case "security_groups":
			if !equalSet(want.(*schema.Set), stateVNI[key]) {
				return false
			}
		case "ips":
			got := []interface{}{}
			switch g := stateVNI[key].(type) {
			case *schema.Set:
				got = g.List()
			case []interface{}:
				got = g
			}
			gotIDs := []interface{}{}
			for _, item := range got {
				if m, ok := item.(map[string]interface{}); ok {
					if rip, ok := m["reserved_ip"].(string); ok && rip != "" {
						gotIDs = append(gotIDs, map[string]interface{}{"reserved_ip": rip})
					}
				}
			}
			if !want.(*schema.Set).Equal(schema.NewSet(hashIpsList, gotIDs)) {
				return false
			}
		default:
			if !equalScalar(want, stateVNI[key]) {
				return false
			}
		}
	}
	return true
}

// resourceIBMISBareMetalServerNetworkAttachmentsOrderDiff drops the planned
// change to network_attachments when configuration and state hold the same
// rows, by name and configured values, in a different order. A refresh has
// no configuration to order by, so this also covers state written in API
// order by earlier versions.
func resourceIBMISBareMetalServerNetworkAttachmentsOrderDiff(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
	if d.Id() == "" || !d.HasChange("network_attachments") {
		return nil
	}
	raw := d.GetRawConfig()
	oldValue, _ := d.GetChange("network_attachments")
	oldList, _ := oldValue.([]interface{})
	count := configuredNetworkAttachmentCount(raw)
	if count < 0 || count != len(oldList) || count == 0 {
		return nil
	}
	byName := map[string]map[string]interface{}{}
	for _, item := range oldList {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil
		}
		name, _ := m["name"].(string)
		if name == "" {
			return nil
		}
		if _, dup := byName[name]; dup {
			return nil
		}
		byName[name] = m
	}
	for j := 0; j < count; j++ {
		cfg, unknown := networkAttachmentFromRawConfig(raw, j)
		if unknown {
			return nil
		}
		name, _ := cfg["name"].(string)
		state, ok := byName[name]
		if !ok || !networkAttachmentMatchesState(cfg, state) {
			return nil
		}
		delete(byName, name)
	}
	if len(byName) != 0 {
		return nil
	}
	log.Printf("[DEBUG] ibm_is_bare_metal_server %s: network_attachments differ from state only in order, dropping the change", d.Id())
	return d.Clear("network_attachments")
}
