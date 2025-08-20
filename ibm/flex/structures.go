// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

//nolint:deadcode,unused
package flex

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	//ResourceControllerURL ...
	ResourceControllerURL = "resource_controller_url"
	//ResourceName ...
	ResourceName = "resource_name"
	//ResourceCRN ...
	ResourceCRN = "resource_crn"
	//ResourceStatus ...
	ResourceStatus = "resource_status"
	//ResourceGroupName ...
	ResourceGroupName = "resource_group_name"
	//RelatedCRN ...
	RelatedCRN = "related_crn"
)

// HashInt ...
func HashInt(v interface{}) int { return v.(int) }

func ExpandStringList(input []interface{}) []string {
	vs := make([]string, len(input))
	for i, v := range input {
		vs[i] = v.(string)
	}
	return vs
}

func FlattenStringList(list []string) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func ExpandIntList(input []interface{}) []int {
	vs := make([]int, len(input))
	for i, v := range input {
		vs[i] = v.(int)
	}
	return vs
}

func FlattenIntList(list []int) []interface{} {
	vs := make([]interface{}, len(list))
	for i, v := range list {
		vs[i] = v
	}
	return vs
}

func FlattenMapInterfaceVal(m map[string]interface{}) map[string]string {
	out := make(map[string]string)
	for k, v := range m {
		out[k] = fmt.Sprintf("%v", v)
	}
	return out
}

func NewStringSet(f schema.SchemaSetFunc, in []string) *schema.Set {
	var out = make([]interface{}, len(in), len(in))
	for i, v := range in {
		out[i] = v
	}
	return schema.NewSet(f, out)
}

func NormalizeJSONString(jsonString interface{}) (string, error) {
	var j interface{}
	if jsonString == nil || jsonString.(string) == "" {
		return "", nil
	}
	s := jsonString.(string)
	err := json.Unmarshal([]byte(s), &j)
	if err != nil {
		return s, err
	}
	bytes, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

func ptrToInt(i int) *int {
	return &i
}

func PtrToString(s string) *string {
	return &s
}

func IntValue(i64 *int64) (i int) {
	if i64 != nil {
		i = int(*i64)
	}
	return
}

func Float64Value(f32 *float32) (f float64) {
	if f32 != nil {
		f = float64(*f32)
	}
	return
}

func DateToString(d *strfmt.Date) (s string) {
	if d != nil {
		s = d.String()
	}
	return
}

func DateTimeToString(dt *strfmt.DateTime) (s string) {
	if dt != nil {
		s = dt.String()
	}
	return
}

func IsEmpty(object interface{}) bool {
	//First check normal definitions of empty
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	//Then see if it's a struct
	if reflect.ValueOf(object).Kind() == reflect.Struct {
		// and create an empty copy of the struct object to compare against
		empty := reflect.New(reflect.TypeOf(object)).Elem().Interface()
		if reflect.DeepEqual(object, empty) {
			return true
		}
	}
	return false
}

func IdParts(id string) ([]string, error) {
	if strings.Contains(id, "/") {
		parts := strings.Split(id, "/")
		return parts, nil
	}
	return []string{}, fmt.Errorf("The given id %s does not contain / please check documentation on how to provider id during import command", id)
}

func SepIdParts(id string, separator string) ([]string, error) {
	if strings.Contains(id, separator) {
		parts := strings.Split(id, separator)
		return parts, nil
	}
	return []string{}, fmt.Errorf("The given id %s does not contain %s please check documentation on how to provider id during import command", id, separator)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func ResourceIBMVPCHash(v interface{}) int {
	var buf bytes.Buffer
	buf.WriteString(strings.ToLower(v.(string)))
	return conns.String(buf.String())
}

func ResourceTagsCustomizeDiff(diff *schema.ResourceDiff) error {

	if diff.Id() != "" && diff.HasChange("tags") {
		o, n := diff.GetChange("tags")
		oldSet := o.(*schema.Set)
		newSet := n.(*schema.Set)
		removeInt := oldSet.Difference(newSet).List()
		addInt := newSet.Difference(oldSet).List()
		if v := os.Getenv("IC_ENV_TAGS"); v != "" {
			s := strings.Split(v, ",")
			if len(removeInt) == len(s) && len(addInt) == 0 {
				fmt.Println("Suppresing the TAG diff ")
				return diff.Clear("tags")
			}
		}
	}
	return nil
}

func GetTagsUsingCRN(meta interface{}, resourceCRN string) (*schema.Set, error) {

	gtClient, err := meta.(conns.ClientSession).GlobalTaggingAPI()
	if err != nil {
		return nil, fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	taggingResult, err := gtClient.Tags().GetTags(resourceCRN)
	if err != nil {
		return nil, err
	}
	var taglist []string
	for _, item := range taggingResult.Items {
		taglist = append(taglist, item.Name)
	}
	log.Println("tagList: ", taglist)
	return NewStringSet(ResourceIBMVPCHash, taglist), nil
}

func UpdateTagsUsingCRN(oldList, newList interface{}, meta interface{}, resourceCRN string) error {
	gtClient, err := meta.(conns.ClientSession).GlobalTaggingAPI()
	if err != nil {
		return fmt.Errorf("Error getting global tagging client settings: %s", err)
	}
	if oldList == nil {
		oldList = new(schema.Set)
	}
	if newList == nil {
		newList = new(schema.Set)
	}
	olds := oldList.(*schema.Set)
	news := newList.(*schema.Set)
	removeInt := olds.Difference(news).List()
	addInt := news.Difference(olds).List()
	add := make([]string, len(addInt))
	for i, v := range addInt {
		add[i] = fmt.Sprint(v)
	}
	remove := make([]string, len(removeInt))
	for i, v := range removeInt {
		remove[i] = fmt.Sprint(v)
	}

	schematicTags := os.Getenv("IC_ENV_TAGS")
	var envTags []string
	if schematicTags != "" {
		envTags = strings.Split(schematicTags, ",")
		add = append(add, envTags...)
	}

	if len(remove) > 0 {
		_, err := gtClient.Tags().DetachTags(resourceCRN, remove)
		if err != nil {
			return fmt.Errorf("Error detaching database tags %v: %s", remove, err)
		}
		for _, v := range remove {
			_, err := gtClient.Tags().DeleteTag(v)
			if err != nil {
				return fmt.Errorf("Error deleting database tag %v: %s", v, err)
			}
		}
	}

	if len(add) > 0 {
		_, err := gtClient.Tags().AttachTags(resourceCRN, add)
		if err != nil {
			return fmt.Errorf("Error updating database tags %v : %s", add, err)
		}
	}

	return nil
}

// GetNext ...
func GetNext(next interface{}) string {
	if reflect.ValueOf(next).IsNil() {
		return ""
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return ""
	}

	q := u.Query()
	return q.Get("start")
}

/* Return the default resource group */
func DefaultResourceGroup(meta interface{}) (string, error) {
	rsMangClient, err := meta.(conns.ClientSession).ResourceManagementAPIv2()
	if err != nil {
		return "", err
	}
	resourceGroupQuery := managementv2.ResourceGroupQuery{
		Default: true,
	}
	grpList, err := rsMangClient.ResourceGroup().List(&resourceGroupQuery)
	if err != nil {
		return "", err
	}
	if len(grpList) <= 0 {
		return "", fmt.Errorf("The targeted resource group could not be found. Make sure you have required permissions to access the resource group.")
	}
	return grpList[0].ID, nil
}

// Use this function for attributes which only should be applied in resource creation time.
func ApplyOnce(k, o, n string, d *schema.ResourceData) bool {
	return len(d.Id()) != 0
}

// Use this function for immutable attributes which should be throw an error on any attempted change.
func Immutable(k, old, new string, d *schema.ResourceData) bool {
	if old != new {
		log.Fatal(fmt.Errorf("Immutable function cannot be modified"))
		return false
	}
	return true
}

func FetchResourceInstanceDetails(d *schema.ResourceData, meta interface{}, instanceID string) error {
	// Get ResourceController from ClientSession
	resourceControllerClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return err
	}

	getResourceOpts := resourcecontrollerv2.GetResourceInstanceOptions{
		ID: &instanceID,
	}

	instance, response, err := resourceControllerClient.GetResourceInstance(&getResourceOpts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving resource instance: %s\n%s", err, response)
		return fmt.Errorf("Error retrieving resource instance: %s\n%s", err, response)
	}
	if strings.Contains(*instance.State, "removed") {
		log.Printf("[WARN] Removing instance from TF state because it's now in removed state")
		d.SetId("")
		return nil
	}

	extensionsMap := Flatten(instance.Extensions)
	if extensionsMap == nil {
		log.Printf("[DEBUG] Error parsing resource instance: Endpoints are missing in instance Extensions map")
		return fmt.Errorf("Error parsing resource instance: Endpoints are missing in instance Extensions map")
	}
	_ = d.Set("extensions", extensionsMap)

	return nil
}

func GetResourceInstanceURL(d *schema.ResourceData, meta interface{}) (*string, error) {

	var endpoint string
	extensions := d.Get("extensions").(map[string]interface{})

	if url, ok := extensions["endpoints.public"]; ok {
		endpoint = "https://" + url.(string)
	}

	if endpoint == "" {
		return nil, fmt.Errorf("[ERROR] Missing endpoints.public in extensions")
	}

	return &endpoint, nil
}

// error object
type ServiceErrorResponse struct {
	Message    string
	StatusCode int
	Result     interface{}
	Error      error
}

func (response *ServiceErrorResponse) String() string {
	output, err := json.MarshalIndent(response, "", "    ")
	if err == nil {
		return fmt.Sprintf("%+v\n", string(output))
	}
	return fmt.Sprintf("Error : %#v", response)
}

// Stringify returns the stringified form of value "v".
// If "v" is a string-based type (string, strfmt.Date, strfmt.DateTime, strfmt.UUID, etc.),
// then it is returned unchanged (e.g. `this is a string`, `foo`, `2025-06-03`).
// Otherwise, json.Marshal() is used to serialze "v" and the resulting string is returned
// (e.g. `32`, `true`, `[true, false, true]`, `{"foo": "bar"}`).
// Note: the backticks in the comments above are not part of the returned strings.
func Stringify(v interface{}) string {
	if !core.IsNil(v) {
		if s, ok := v.(string); ok {
			return s
		} else if s, ok := v.(*string); ok {
			return *s
		} else if s, ok := v.(interface{ String() string }); ok {
			return s.String()
		} else {
			bytes, err := json.Marshal(v)
			if err != nil {
				log.Printf("[ERROR] Error marshaling 'any type' value as string: %s", err.Error())
				return ""
			}
			return string(bytes)
		}
	}
	return ""
}

// NullTagSet is used by generated terraform resource code.
var NullTagSet basetypes.SetValue = basetypes.NewSetNull(types.StringType)

// GetTagsUsingCRNFW is a plugin-framework enabled version of GetTagsUsingCRN().
func GetTagsUsingCRNFW(providerData interface{}, resourceCRN string) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics

	gtClient, err := providerData.(conns.ClientSession).GlobalTaggingAPI()
	if err != nil {
		diags.AddError("error retrieving global tagging service client", err.Error())
		return basetypes.NewSetNull(types.StringType), diags
	}
	taggingResult, err := gtClient.Tags().GetTags(resourceCRN)
	if err != nil {
		diags.AddError(fmt.Sprintf("error retrieving tags for resource CRN '%s'", resourceCRN), err.Error())
		return basetypes.NewSetNull(types.StringType), diags
	}
	var tags []string
	for _, item := range taggingResult.Items {
		tags = append(tags, item.Name)
	}
	tflog.Debug(context.Background(), fmt.Sprintf("Retrieved tags for CRN '%s': %s", resourceCRN, Stringify(tags)))

	tagList, _ := StringSliceToSetValue(tags)

	return tagList, nil
}

// UpdateTagsUsingCRNFW is a plugin-framework enabled version of UpdateTagsUsingCRN().
func UpdateTagsUsingCRNFW(oldTags, newTags types.Set, providerData interface{}, resourceCRN string) diag.Diagnostics {
	var diags diag.Diagnostics

	gtClient, err := providerData.(conns.ClientSession).GlobalTaggingAPI()
	if err != nil {
		diags.AddError("error retrieving global tagging service client", err.Error())
		return diags
	}

	// Store the old tags in a Set.
	oldTagsSlice, _ := SetValueToStringSlice(oldTags)
	if oldTagsSlice == nil {
		oldTagsSlice = []string{}
	}
	oldSet := NewSet()
	oldSet.Add(oldTagsSlice...)

	// Store the new tags in a Set.
	newTagsSlice, _ := SetValueToStringSlice(newTags)
	if newTagsSlice == nil {
		newTagsSlice = []string{}
	}
	newSet := NewSet()
	newSet.Add(newTagsSlice...)

	// The old tags that are not in the new set need to be removed.
	removeSet := oldSet.Difference(newSet)

	// The new tags that are not in the old set need to be added.
	addSet := newSet.Difference(oldSet)

	// We'll also add any tags contained in the IC_ENV_TAGS env variable.
	envTags := os.Getenv("IC_ENV_TAGS")
	if envTags != "" {
		slice := strings.Split(envTags, ",")
		addSet.Add(slice...)
	}

	// Detach and delete any tags that need to be removed.
	if removeSet.Size() > 0 {
		_, err := gtClient.Tags().DetachTags(resourceCRN, removeSet.Values())
		if err != nil {
			diags.AddError(fmt.Sprintf("error detaching tags for resource CRN '%s'", resourceCRN), err.Error())
			return diags
		}
		for _, v := range removeSet.Values() {
			_, err := gtClient.Tags().DeleteTag(v)
			if err != nil {
				diags.AddError(fmt.Sprintf("error deleting tag '%s' for resource CRN '%s'", v, resourceCRN), err.Error())
				return diags
			}
		}
	}

	if addSet.Size() > 0 {
		add := addSet.Values()
		_, err := gtClient.Tags().AttachTags(resourceCRN, add)
		if err != nil {
			diags.AddError(fmt.Sprintf("error attaching tags for resource CRN '%s'", resourceCRN), err.Error())
			return diags
		}
	}

	return diags
}
