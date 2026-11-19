// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package accountmanagement

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func TestAccountTraitsSchemasValid(t *testing.T) {
	ctx := context.Background()
	rr := &resource.SchemaResponse{}
	NewAccountTraitsResource().Schema(ctx, resource.SchemaRequest{}, rr)
	if d := rr.Schema.ValidateImplementation(ctx); d.HasError() {
		t.Errorf("resource schema invalid: %v", d)
	}
	dr := &datasource.SchemaResponse{}
	NewAccountTraitsDataSource().Schema(ctx, datasource.SchemaRequest{}, dr)
	if d := dr.Schema.ValidateImplementation(ctx); d.HasError() {
		t.Errorf("data source schema invalid: %v", d)
	}
}
