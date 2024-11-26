// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider_test

import (
	"context"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
)

// go test -bench=BenchmarkProtoV6ProviderServerFactory -benchtime 1x -benchmem -run=Bench -v ./internal/provider
func BenchmarkProtoV6ProviderServerFactory(b *testing.B) {
	_, p, err := provider.ProtoV6ProviderServerFactory(context.Background())

	if err != nil {
		b.Fatal(err)
	}

	if b.N == 1 {
		b.Logf("%d resources, %d data sources", len(p.ResourcesMap), len(p.DataSourcesMap))
	}
}
