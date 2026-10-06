// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams

import (
	"strings"
	"testing"

	"github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatObject(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]interface{}
		expected string
	}{
		{
			name:     "empty map",
			input:    map[string]interface{}{},
			expected: "",
		},
		{
			name: "flat map with primitive types",
			input: map[string]interface{}{
				"string_val": "foo",
				"int_val":    123,
				"bool_val":   true,
				"float_val":  45.67,
			},
			expected: `"bool_val":bool,"float_val":float64,"int_val":int,"string_val":string`,
		},
		{
			name: "nested map",
			input: map[string]interface{}{
				"a": "hello",
				"nested": map[string]interface{}{
					"count": 10,
					"flag":  false,
				},
			},
			expected: `"a":string,"nested":{"count":int,"flag":bool}`,
		},
		{
			name: "deeply nested map",
			input: map[string]interface{}{
				"level1": map[string]interface{}{
					"level2": map[string]interface{}{
						"leaf": "value",
					},
				},
			},
			expected: `"level1":{"level2":{"leaf":string}}`,
		},
		{
			name: "map with slice and nil values",
			input: map[string]interface{}{
				"slice_val": []string{"a", "b"},
				"nil_val":   nil,
			},
			expected: `"nil_val":<nil>,"slice_val":[]string`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := formatObject(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestParseInstanceExtensions(t *testing.T) {
	t.Run("gen1: valid kafka_http_url and kafka_brokers_sasl", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{
				"kafka_http_url": "https://tenant1.example.com",
				"kafka_brokers_sasl": []interface{}{
					"broker-c.example.com:9093",
					"broker-a.example.com:9093",
					"broker-b.example.com:9093",
				},
			},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.NoError(t, err)
		assert.Equal(t, 1, ext.platformGeneration)
		assert.Equal(t, "https://tenant1.example.com", ext.adminURL)
		wantServers := []string{
			"broker-a.example.com:9093",
			"broker-b.example.com:9093",
			"broker-c.example.com:9093",
		}
		require.Len(t, ext.bootstrapServers, len(wantServers))
		assert.Equal(t, wantServers, ext.bootstrapServers)
	})

	t.Run("gen2: valid dataservices connection block", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{
				"dataservices": map[string]any{
					"connection": map[string]any{
						"bootstrap_servers": "broker-1.example.com:9093,broker-2.example.com:9093",
						"rest_url":          "https://admin.example.com",
					},
				},
			},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.NoError(t, err)
		assert.Equal(t, 2, ext.platformGeneration)
		assert.Equal(t, "https://admin.example.com", ext.adminURL)
		wantServers := strings.Split("broker-1.example.com:9093,broker-2.example.com:9093", ",")
		require.Len(t, ext.bootstrapServers, len(wantServers))
		assert.Equal(t, wantServers, ext.bootstrapServers)
	})

	t.Run("error: empty extensions", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.Error(t, err)
		assert.Nil(t, ext)
	})

	t.Run("error: kafka_http_url present but kafka_brokers_sasl missing", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{
				"kafka_http_url": "https://tenant1.example.com",
			},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.Error(t, err)
		assert.Nil(t, ext)
	})

	t.Run("error: dataservices present but rest_url missing", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{
				"dataservices": map[string]any{
					"connection": map[string]any{
						"bootstrap_servers": "broker-1.example.com:9093",
						// rest_url intentionally omitted
					},
				},
			},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.Error(t, err)
		assert.Nil(t, ext)
	})

	t.Run("error: dataservices present but bootstrap_servers missing", func(t *testing.T) {
		instance := &resourcecontrollerv2.ResourceInstance{
			Extensions: map[string]interface{}{
				"dataservices": map[string]any{
					"connection": map[string]any{
						// bootstrap_servers intentionally omitted
						"rest_url": "https://admin.example.com",
					},
				},
			},
		}
		ext, err := parseInstanceExtensions(instance, nil)
		require.Error(t, err)
		assert.Nil(t, ext)
	})
}
