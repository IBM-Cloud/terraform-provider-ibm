// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package flex

import (
	"reflect"
	"testing"
)

func TestEnvTags(t *testing.T) {
	t.Setenv(EnvTagsVar, "")
	if got := EnvTags(); got != nil {
		t.Errorf("EnvTags() unset = %v, want nil", got)
	}

	t.Setenv(EnvTagsVar, "Team:Platform, env:dev,,")
	if got := EnvTags(); !reflect.DeepEqual(got, []string{"Team:Platform", "env:dev"}) {
		t.Errorf("EnvTags() = %v", got)
	}
}

func TestIsEnvTag(t *testing.T) {
	t.Setenv(EnvTagsVar, "Team:Platform")
	// The tagging service stores the tag in lower case.
	if !IsEnvTag("team:platform") {
		t.Error("IsEnvTag(team:platform) = false, want true")
	}
	if IsEnvTag("team:other") {
		t.Error("IsEnvTag(team:other) = true, want false")
	}
}

func TestTagsWithoutEnvTags(t *testing.T) {
	t.Setenv(EnvTagsVar, "Team:Platform,env:ci")
	got := TagsWithoutEnvTags([]string{"team:platform", "rhaii-da", "env:ci", "test"})
	if !reflect.DeepEqual(got, []string{"rhaii-da", "test"}) {
		t.Errorf("TagsWithoutEnvTags() = %v", got)
	}

	t.Setenv(EnvTagsVar, "")
	got = TagsWithoutEnvTags([]string{"team:platform", "rhaii-da"})
	if !reflect.DeepEqual(got, []string{"team:platform", "rhaii-da"}) {
		t.Errorf("TagsWithoutEnvTags() without IC_ENV_TAGS = %v", got)
	}
}
