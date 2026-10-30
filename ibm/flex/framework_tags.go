// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"context"
	"fmt"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Tag helpers for Plugin Framework resources. They are the counterparts of
// GetGlobalTagsUsingCRN, UpdateGlobalTagsUsingCRN, ResourceTagsCustomizeDiff
// and ResourceValidateAccessTags for SDKv2 resources.

// FrameworkSetToStrings returns the elements of a known set, or nil.
func FrameworkSetToStrings(ctx context.Context, s types.Set) ([]string, diag.Diagnostics) {
	if s.IsNull() || s.IsUnknown() {
		return nil, nil
	}
	out := []string{}
	d := s.ElementsAs(ctx, &out, false)
	return out, d
}

// FrameworkReadTagsUsingCRN reads the tags of a type attached to a CRN as a set
// for the state. User tags come without the IC_ENV_TAGS tags, so the state
// matches the config. A read error is a warning and gives an empty set.
func FrameworkReadTagsUsingCRN(ctx context.Context, meta interface{}, resourceCRN, resourceType, tagType string) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics
	tags, err := ReadGlobalTagsUsingCRN(meta, resourceCRN, resourceType, tagType)
	if err != nil {
		diags.AddWarning(fmt.Sprintf("Error reading the %s tags of %s", tagTypeName(tagType), resourceCRN), err.Error())
	}
	if tags == nil {
		tags = []string{}
	}
	set, d := types.SetValueFrom(ctx, types.StringType, tags)
	diags.Append(d...)
	return set, diags
}

// FrameworkUpdateTagsUsingCRN attaches and detaches tags so the CRN ends up
// with the new set. For user tags the IC_ENV_TAGS tags are attached as well.
// Call it from Create with a null old set, and from Update when the set changed.
func FrameworkUpdateTagsUsingCRN(ctx context.Context, meta interface{}, resourceCRN, resourceType, tagType string, oldTags, newTags types.Set) error {
	oldList, d := FrameworkSetToStrings(ctx, oldTags)
	if d.HasError() {
		return fmt.Errorf("reading the old %s tags: %s", tagTypeName(tagType), d)
	}
	newList, d := FrameworkSetToStrings(ctx, newTags)
	if d.HasError() {
		return fmt.Errorf("reading the new %s tags: %s", tagTypeName(tagType), d)
	}
	return UpdateGlobalTagsUsingCRN(
		NewStringSet(ResourceIBMVPCHash, oldList),
		NewStringSet(ResourceIBMVPCHash, newList),
		meta, resourceCRN, resourceType, tagType)
}

// FrameworkValidateAccessTags fails the plan when a configured access tag does
// not exist in the account. Call it from ModifyPlan with the config value.
func FrameworkValidateAccessTags(ctx context.Context, meta interface{}, accessTags types.Set, attrPath path.Path) diag.Diagnostics {
	var diags diag.Diagnostics
	tags, d := FrameworkSetToStrings(ctx, accessTags)
	diags.Append(d...)
	if len(tags) == 0 {
		return diags
	}
	missing, err := missingAccessTags(ctx, meta, tags)
	if err != nil {
		diags.AddAttributeError(attrPath, "Error checking access tags", err.Error())
	} else if len(missing) > 0 {
		diags.AddAttributeError(attrPath, "Access tags do not exist",
			fmt.Sprintf("Access tag(s) %s do not exist in the account. Create them first, for example with ibm_iam_access_tag.", strings.Join(missing, ", ")))
	}
	return diags
}

func missingAccessTags(ctx context.Context, meta interface{}, tags []string) ([]string, error) {
	gtClient, err := meta.(conns.ClientSession).GlobalTaggingAPIv1()
	if err != nil {
		return nil, fmt.Errorf("error getting the Global Tagging client: %s", err)
	}
	existing := map[string]struct{}{}
	var offset int64
	const limit int64 = 1000
	for {
		result, _, err := gtClient.ListTagsWithContext(ctx, &globaltaggingv1.ListTagsOptions{
			TagType: core.StringPtr("access"),
			Offset:  &offset,
			Limit:   core.Int64Ptr(limit),
		})
		if err != nil {
			return nil, fmt.Errorf("error listing access tags: %s", err)
		}
		for _, item := range result.Items {
			if item.Name != nil {
				existing[strings.ToLower(*item.Name)] = struct{}{}
			}
		}
		if int64(len(result.Items)) < limit {
			break
		}
		offset += limit
	}
	missing := []string{}
	for _, t := range tags {
		if _, ok := existing[strings.ToLower(t)]; !ok {
			missing = append(missing, t)
		}
	}
	return missing, nil
}

func tagTypeName(tagType string) string {
	if strings.TrimSpace(tagType) == "" {
		return "user"
	}
	return tagType
}
