package vpcintegration

import "os"

// URL - URL for the environment
var URL = "https://dev.console.test.cloud.ibm.com/rias-mock/v1"

// APIKey - for your VPC account
var APIKey = os.Getenv("IBM_CLOUD_KEY")

// IAMURL - IAM url
var IAMURL = "https://iam.test.cloud.ibm.com/identity/token"
