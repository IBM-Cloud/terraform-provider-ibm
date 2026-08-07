// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestTransformGen2ConnectionToSchema_PostgreSQL tests PostgreSQL connection transformation
func TestTransformGen2ConnectionToSchema_PostgreSQL(t *testing.T) {
	input := map[string]interface{}{
		"type":     "uri",
		"scheme":   "postgres",
		"database": "postgres",
		"path":     "/postgres",
		"composed": []interface{}{
			"postgres://user:pass@host.example.com:5432/postgres?sslmode=verify-full",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(5432),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
		"certificate": map[string]interface{}{
			"name":               "ca-certificate",
			"certificate_base64": "LS0tLS1CRUdJTi...",
		},
		"query_options": map[string]interface{}{
			"sslmode": "verify-full",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "postgres", result["scheme"])
	assert.Equal(t, "postgres", result["database"])
	assert.Equal(t, "/postgres", result["path"])
	assert.NotNil(t, result["composed"])
	assert.NotNil(t, result["hosts"])
	assert.NotNil(t, result["authentication"])
	assert.NotNil(t, result["certificate"])
	assert.NotNil(t, result["query_options"])

	// Verify hosts transformation
	hosts := result["hosts"].([]map[string]interface{})
	assert.Len(t, hosts, 1)
	assert.Equal(t, "host.example.com", hosts[0]["hostname"])
	assert.Equal(t, int64(5432), hosts[0]["port"])

	// Verify authentication transformation
	auth := result["authentication"].([]map[string]interface{})
	assert.Len(t, auth, 1)
	assert.Equal(t, "direct", auth[0]["method"])
	assert.Equal(t, "user", auth[0]["username"])
	assert.Equal(t, "pass", auth[0]["password"])

	// Verify certificate transformation
	cert := result["certificate"].([]map[string]interface{})
	assert.Len(t, cert, 1)
	assert.Equal(t, "ca-certificate", cert[0]["name"])
	assert.Equal(t, "LS0tLS1CRUdJTi...", cert[0]["certificate_base64"])

	// Verify query_options transformation
	queryOpts := result["query_options"].(map[string]interface{})
	assert.Equal(t, "verify-full", queryOpts["sslmode"])
}

// TestTransformGen2ConnectionToSchema_MongoDB tests MongoDB connection transformation
func TestTransformGen2ConnectionToSchema_MongoDB(t *testing.T) {
	input := map[string]interface{}{
		"type":     "uri",
		"scheme":   "mongodb",
		"database": "admin",
		"path":     "/admin",
		"composed": []interface{}{
			"mongodb://user:pass@host1.example.com:27017,host2.example.com:27017/admin?authSource=admin&replicaSet=replset",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host1.example.com",
				"port":     float64(27017),
			},
			map[string]interface{}{
				"hostname": "host2.example.com",
				"port":     float64(27017),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
		"query_options": map[string]interface{}{
			"authSource": "admin",
			"replicaSet": "replset",
			"tls":        true, // Boolean value
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "mongodb", result["scheme"])
	assert.Equal(t, "admin", result["database"])

	// Verify multiple hosts
	hosts := result["hosts"].([]map[string]interface{})
	assert.Len(t, hosts, 2)
	assert.Equal(t, "host1.example.com", hosts[0]["hostname"])
	assert.Equal(t, int64(27017), hosts[0]["port"])
	assert.Equal(t, "host2.example.com", hosts[1]["hostname"])
	assert.Equal(t, int64(27017), hosts[1]["port"])

	// Verify query_options with boolean conversion
	queryOpts := result["query_options"].(map[string]interface{})
	assert.Equal(t, "admin", queryOpts["authSource"])
	assert.Equal(t, "replset", queryOpts["replicaSet"])
	assert.Equal(t, "true", queryOpts["tls"]) // Boolean converted to string
}

// TestTransformGen2ConnectionToSchema_Redis tests Redis connection transformation
func TestTransformGen2ConnectionToSchema_Redis(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "rediss",
		"composed": []interface{}{
			"rediss://:password@host.example.com:6379/0",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(6379),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"password": "password",
		},
		"query_options": map[string]interface{}{
			"ssl": true,
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "rediss", result["scheme"])

	// Verify query_options boolean conversion
	queryOpts := result["query_options"].(map[string]interface{})
	assert.Equal(t, "true", queryOpts["ssl"])
}

// TestTransformGen2ConnectionToSchema_RabbitMQ_AMQPS tests RabbitMQ AMQPS connection
func TestTransformGen2ConnectionToSchema_RabbitMQ_AMQPS(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "amqps",
		"composed": []interface{}{
			"amqps://user:pass@host.example.com:5671",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(5671),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "amqps", result["scheme"])
}

// TestTransformGen2ConnectionToSchema_RabbitMQ_MQTTS tests RabbitMQ MQTTS connection
func TestTransformGen2ConnectionToSchema_RabbitMQ_MQTTS(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "mqtts",
		"composed": []interface{}{
			"mqtts://user:pass@host.example.com:8883",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(8883),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "mqtts", result["scheme"])
}

// TestTransformGen2ConnectionToSchema_RabbitMQ_StompSSL tests RabbitMQ STOMP SSL connection
func TestTransformGen2ConnectionToSchema_RabbitMQ_StompSSL(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "stomp+ssl",
		"composed": []interface{}{
			"stomp+ssl://user:pass@host.example.com:61614",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(61614),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "stomp+ssl", result["scheme"])
}

// TestTransformGen2ConnectionToSchema_MySQL tests MySQL connection transformation
func TestTransformGen2ConnectionToSchema_MySQL(t *testing.T) {
	input := map[string]interface{}{
		"type":     "uri",
		"scheme":   "mysql",
		"database": "ibmclouddb",
		"composed": []interface{}{
			"mysql://user:pass@host.example.com:3306/ibmclouddb?ssl-mode=REQUIRED",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(3306),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
		"query_options": map[string]interface{}{
			"ssl-mode": "REQUIRED",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "mysql", result["scheme"])
	assert.Equal(t, "ibmclouddb", result["database"])
}

// TestTransformGen2ConnectionToSchema_Elasticsearch_HTTPS tests Elasticsearch HTTPS connection
func TestTransformGen2ConnectionToSchema_Elasticsearch_HTTPS(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "https",
		"composed": []interface{}{
			"https://user:pass@host.example.com:9200",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(9200),
			},
		},
		"authentication": map[string]interface{}{
			"method":   "direct",
			"username": "user",
			"password": "pass",
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "https", result["scheme"])
}

// TestTransformGen2ConnectionToSchema_GRPC tests gRPC connection (etcd, Elasticsearch)
func TestTransformGen2ConnectionToSchema_GRPC(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "grpc",
		"composed": []interface{}{
			"grpc://host.example.com:2379",
		},
		"hosts": []interface{}{
			map[string]interface{}{
				"hostname": "host.example.com",
				"port":     float64(2379),
			},
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "grpc", result["scheme"])
}

// TestTransformGen2ConnectionToSchema_DataStax tests DataStax connection types
func TestTransformGen2ConnectionToSchema_DataStax(t *testing.T) {
	testCases := []struct {
		name   string
		scheme string
	}{
		{"Analytics", "analytics"},
		{"BIConnector", "bi_connector"},
		{"OpsManager", "ops_manager"},
		{"EMP", "emp"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := map[string]interface{}{
				"type":   "uri",
				"scheme": tc.scheme,
				"composed": []interface{}{
					tc.scheme + "://host.example.com:27017",
				},
				"hosts": []interface{}{
					map[string]interface{}{
						"hostname": "host.example.com",
						"port":     float64(27017),
					},
				},
			}

			result, err := transformGen2ConnectionToSchema(input)

			assert.NoError(t, err)
			assert.Equal(t, "uri", result["type"])
			assert.Equal(t, tc.scheme, result["scheme"])
		})
	}
}

// TestTransformGen2ConnectionToSchema_QueryOptionsTypes tests various query option types
func TestTransformGen2ConnectionToSchema_QueryOptionsTypes(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "test",
		"query_options": map[string]interface{}{
			"string_value":   "value",
			"bool_true":      true,
			"bool_false":     false,
			"int_value":      42,
			"float_value":    3.14,
			"int64_value":    int64(9999),
			"negative_int":   -10,
			"negative_float": -2.5,
		},
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	queryOpts := result["query_options"].(map[string]interface{})

	// String should remain string
	assert.Equal(t, "value", queryOpts["string_value"])

	// Booleans should be converted to strings
	assert.Equal(t, "true", queryOpts["bool_true"])
	assert.Equal(t, "false", queryOpts["bool_false"])

	// Numbers should be converted to strings
	assert.Equal(t, "42", queryOpts["int_value"])
	assert.Equal(t, "3.14", queryOpts["float_value"])
	assert.Equal(t, "9999", queryOpts["int64_value"])
	assert.Equal(t, "-10", queryOpts["negative_int"])
	assert.Equal(t, "-2.5", queryOpts["negative_float"])
}

// TestTransformGen2ConnectionToSchema_PortTypes tests various port number types
func TestTransformGen2ConnectionToSchema_PortTypes(t *testing.T) {
	testCases := []struct {
		name     string
		portType interface{}
		expected int64
	}{
		{"Float64", float64(5432), int64(5432)},
		{"Int", int(5432), int64(5432)},
		{"Int64", int64(5432), int64(5432)},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := map[string]interface{}{
				"type": "uri",
				"hosts": []interface{}{
					map[string]interface{}{
						"hostname": "host.example.com",
						"port":     tc.portType,
					},
				},
			}

			result, err := transformGen2ConnectionToSchema(input)

			assert.NoError(t, err)
			hosts := result["hosts"].([]map[string]interface{})
			assert.Len(t, hosts, 1)
			assert.Equal(t, tc.expected, hosts[0]["port"])
		})
	}
}

// TestTransformGen2CliToSchema tests CLI connection transformation
func TestTransformGen2CliToSchema(t *testing.T) {
	input := map[string]interface{}{
		"type": "cli",
		"bin":  "psql",
		"arguments": []interface{}{
			"host=host.example.com",
			"port=5432",
			"dbname=postgres",
			"user=user",
			"password=pass",
			"sslmode=verify-full",
		},
		"composed": []interface{}{
			"PGUSER=user PGPASSWORD=pass PGSSLMODE=verify-full psql 'host=host.example.com port=5432 dbname=postgres'",
		},
		"environment": map[string]interface{}{
			"PGUSER":        "user",
			"PGPASSWORD":    "pass",
			"PGSSLMODE":     "verify-full",
			"PGSSLROOTCERT": "system",
		},
		"certificate": map[string]interface{}{
			"name":               "ca-certificate",
			"certificate_base64": "LS0tLS1CRUdJTi...",
		},
	}

	result, err := transformGen2CliToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "cli", result["type"])
	assert.Equal(t, "psql", result["bin"])
	assert.NotNil(t, result["arguments"])
	assert.NotNil(t, result["composed"])
	assert.NotNil(t, result["environment"])
	assert.NotNil(t, result["certificate"])

	// Verify environment variables
	env := result["environment"].(map[string]interface{})
	assert.Equal(t, "user", env["PGUSER"])
	assert.Equal(t, "pass", env["PGPASSWORD"])
	assert.Equal(t, "verify-full", env["PGSSLMODE"])
	assert.Equal(t, "system", env["PGSSLROOTCERT"])

	// Verify certificate
	cert := result["certificate"].([]map[string]interface{})
	assert.Len(t, cert, 1)
	assert.Equal(t, "ca-certificate", cert[0]["name"])
	assert.Equal(t, "LS0tLS1CRUdJTi...", cert[0]["certificate_base64"])
}

// TestTransformGen2CliToSchema_MongoDB tests MongoDB CLI transformation
func TestTransformGen2CliToSchema_MongoDB(t *testing.T) {
	input := map[string]interface{}{
		"type": "cli",
		"bin":  "mongosh",
		"arguments": []interface{}{
			"mongodb://user:pass@host.example.com:27017/admin?authSource=admin&replicaSet=replset",
		},
		"composed": []interface{}{
			"mongosh 'mongodb://user:pass@host.example.com:27017/admin?authSource=admin&replicaSet=replset'",
		},
		"environment": map[string]interface{}{},
	}

	result, err := transformGen2CliToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "cli", result["type"])
	assert.Equal(t, "mongosh", result["bin"])
	assert.NotNil(t, result["arguments"])
	assert.NotNil(t, result["composed"])
}

// TestTransformGen2CliToSchema_Redis tests Redis CLI transformation
func TestTransformGen2CliToSchema_Redis(t *testing.T) {
	input := map[string]interface{}{
		"type": "cli",
		"bin":  "redis-cli",
		"arguments": []interface{}{
			"-h", "host.example.com",
			"-p", "6379",
			"-a", "password",
			"--tls",
		},
		"composed": []interface{}{
			"redis-cli -h host.example.com -p 6379 -a password --tls",
		},
		"environment": map[string]interface{}{
			"REDISCLI_AUTH": "password",
		},
	}

	result, err := transformGen2CliToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "cli", result["type"])
	assert.Equal(t, "redis-cli", result["bin"])

	// Verify environment
	env := result["environment"].(map[string]interface{})
	assert.Equal(t, "password", env["REDISCLI_AUTH"])
}

// TestTransformGen2ConnectionToSchema_EmptyInput tests handling of empty input
func TestTransformGen2ConnectionToSchema_EmptyInput(t *testing.T) {
	input := map[string]interface{}{}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result)
}

// TestTransformGen2CliToSchema_EmptyInput tests handling of empty CLI input
func TestTransformGen2CliToSchema_EmptyInput(t *testing.T) {
	input := map[string]interface{}{}

	result, err := transformGen2CliToSchema(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result)
}

// TestTransformGen2ConnectionToSchema_MissingOptionalFields tests handling of missing optional fields
func TestTransformGen2ConnectionToSchema_MissingOptionalFields(t *testing.T) {
	input := map[string]interface{}{
		"type":   "uri",
		"scheme": "postgres",
		// Missing: database, path, hosts, authentication, certificate, query_options
	}

	result, err := transformGen2ConnectionToSchema(input)

	assert.NoError(t, err)
	assert.Equal(t, "uri", result["type"])
	assert.Equal(t, "postgres", result["scheme"])
	assert.Nil(t, result["database"])
	assert.Nil(t, result["path"])
	assert.Nil(t, result["hosts"])
	assert.Nil(t, result["authentication"])
	assert.Nil(t, result["certificate"])
	assert.Nil(t, result["query_options"])
}

// Made with Bob
