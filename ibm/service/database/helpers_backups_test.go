// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func backupsTestResourceData(t *testing.T, backups []interface{}) *schema.ResourceData {
	t.Helper()
	raw := map[string]interface{}{
		"name":     "test-gen2-db",
		"location": "us-south",
		"service":  "databases-for-postgresql",
		"plan":     "standard-gen2",
	}
	if backups != nil {
		raw["backups"] = backups
	}
	return schema.TestResourceDataRaw(t, ResourceIBMDatabaseInstance().Schema, raw)
}

func TestBuildBackupsConfig(t *testing.T) {
	cases := []struct {
		name    string
		backups []interface{}
		want    map[string]interface{}
	}{
		{
			name:    "no backups block returns nil",
			backups: nil,
			want:    nil,
		},
		{
			name:    "empty backups block returns nil",
			backups: []interface{}{map[string]interface{}{}},
			want:    nil,
		},
		{
			name: "retention only",
			backups: []interface{}{map[string]interface{}{
				"retention": 45,
			}},
			want: map[string]interface{}{"retention": "45d"},
		},
		{
			name: "preserve true",
			backups: []interface{}{map[string]interface{}{
				"preserve": true,
			}},
			want: map[string]interface{}{"preserve": true},
		},
		{
			name: "copy disabled without targets",
			backups: []interface{}{map[string]interface{}{
				"automatic_backups": []interface{}{map[string]interface{}{
					"copy": []interface{}{map[string]interface{}{
						"enabled": false,
					}},
				}},
			}},
			want: map[string]interface{}{
				"automatic_backups": map[string]interface{}{
					"copy": map[string]interface{}{"enabled": false},
				},
			},
		},
		{
			name: "all attributes with two targets",
			backups: []interface{}{map[string]interface{}{
				"preserve":  true,
				"retention": 60,
				"automatic_backups": []interface{}{map[string]interface{}{
					"copy": []interface{}{map[string]interface{}{
						"enabled": true,
						"target": []interface{}{
							map[string]interface{}{"region": "ca-mon"},
							map[string]interface{}{"region": "eu-de"},
						},
					}},
				}},
			}},
			want: map[string]interface{}{
				"preserve":  true,
				"retention": "60d",
				"automatic_backups": map[string]interface{}{
					"copy": map[string]interface{}{
						"enabled": true,
						"target": []interface{}{
							map[string]interface{}{"region": "ca-mon"},
							map[string]interface{}{"region": "eu-de"},
						},
					},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			d := backupsTestResourceData(t, c.backups)
			require.Equal(t, c.want, buildBackupsConfig(d, false))
		})
	}
}

func TestFlattenBackupsFromExtensions(t *testing.T) {
	cases := []struct {
		name       string
		extensions map[string]interface{}
		want       []interface{}
	}{
		{name: "nil extensions", extensions: nil, want: nil},
		{name: "empty extensions", extensions: map[string]interface{}{}, want: nil},
		{
			name:       "dataservices is not a map",
			extensions: map[string]interface{}{"dataservices": "x"},
			want:       nil,
		},
		{
			name: "backups missing",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{"other": "x"},
			},
			want: nil,
		},
		{
			name: "backups nil",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{"backups": nil},
			},
			want: nil,
		},
		{
			name: "full backups with target; account_id and resource_group_id are ignored",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{
						"automatic_backups": map[string]interface{}{
							"copy": map[string]interface{}{
								"enabled": true,
								"target": []interface{}{
									map[string]interface{}{
										"account_id":        "acc",
										"region":            "ca-mon",
										"resource_group_id": "rg",
									},
								},
							},
							"enabled": true,
							"window":  map[string]interface{}{"start_time": "14:32Z"},
						},
						"preserve":  true,
						"retention": "31d",
					},
				},
			},
			want: []interface{}{map[string]interface{}{
				"preserve":  true,
				"retention": 31,
				"automatic_backups": []interface{}{map[string]interface{}{
					"copy": []interface{}{map[string]interface{}{
						"enabled": true,
						"target": []interface{}{
							map[string]interface{}{"region": "ca-mon"},
						},
					}},
				}},
			}},
		},
		{
			name: "copy disabled has no targets",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{
						"automatic_backups": map[string]interface{}{
							"copy": map[string]interface{}{"enabled": false},
						},
						"preserve":  false,
						"retention": "30d",
					},
				},
			},
			want: []interface{}{map[string]interface{}{
				"preserve":  false,
				"retention": 30,
				"automatic_backups": []interface{}{map[string]interface{}{
					"copy": []interface{}{map[string]interface{}{
						"enabled": false,
						"target":  []interface{}{},
					}},
				}},
			}},
		},
		{
			name: "numeric retention is accepted",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{"retention": float64(45)},
				},
			},
			want: []interface{}{map[string]interface{}{"retention": 45}},
		},
		{
			name: "unparseable retention is skipped without panic",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{
						"preserve":  false,
						"retention": "abc",
					},
				},
			},
			want: []interface{}{map[string]interface{}{
				"preserve": false,
			}},
		},
		{
			name: "empty backups map yields empty block",
			extensions: map[string]interface{}{
				"dataservices": map[string]interface{}{
					"backups": map[string]interface{}{},
				},
			},
			want: []interface{}{map[string]interface{}{}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			require.Equal(t, c.want, flattenBackupsFromExtensions(c.extensions))
		})
	}
}
