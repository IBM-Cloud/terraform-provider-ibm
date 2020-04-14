package vpcintegration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcclassicv1"
)

/**
 * REST methods
 *
 */
const (
	POST   = http.MethodPost
	GET    = http.MethodGet
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
)

// Print - Marshal JSON and print
func Print(printObject interface{}) {
	p, _ := json.MarshalIndent(printObject, "", "\t")
	fmt.Println(string(p))
}

// ToString - Marshal a JSON and return a string
func ToString(printObject interface{}) string {
	p, _ := json.MarshalIndent(printObject, "", "\t")
	return string(p)
}

// PollInstance - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollInstance(vpcService *vpcclassicv1.VpcClassicV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetInstance(vpcService, ID)
			fmt.Println("Current status of VSI - ", *res.Status)
			fmt.Println("Expected status of VSI - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving instance ID %s with error message: %s", ID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollSubnet - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollSubnet(vpcService *vpcclassicv1.VpcClassicV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetSubnet(vpcService, ID)
			fmt.Println("Current status of Subnet - ", *res.Status)
			fmt.Println("Expected status of Subnet - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving subnet ID %s with error message: %s", ID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollVolAttachment - poll and check the status of Volume attachment before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVolAttachment(vpcService *vpcclassicv1.VpcClassicV1, vpcID, volAttachmentID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVolumeAttachment(vpcService, vpcID, volAttachmentID)
			fmt.Println("Current status of attachment - ", *res.Status)
			fmt.Println("Expected status of attachment - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving volume attachment ID %s with error message: %s", vpcID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollLB - poll and check the status of LB Listener before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollLB(vpcService *vpcclassicv1.VpcClassicV1, lbID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetLoadBalancer(vpcService, lbID)
			fmt.Println("Current status of load balancer - ", *res.ProvisioningStatus)
			fmt.Println("Expected status of load balancer - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving load balancer ID %s with error message: %s", lbID, err)
				return false
			}
			if *res.ProvisioningStatus == status {
				fmt.Println("Received expected status - ", *res.ProvisioningStatus)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollVpnGateway - poll and check the status of VpnGateway before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVpnGateway(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVpnGateway(vpcService, gatewayID)
			fmt.Println("Current status of VpnGateway - ", *res.Status)
			fmt.Println("Expected status of VpnGateway - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving VpnGateway ID %s with error message: %s", gatewayID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// TestListResponse - format output for test list APIs
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - GET
// detailed *bool - bool to view the detailed response from API
func TestListResponse(t *testing.T, x interface{}, err error, operation string, detailed *bool, increment func()) {
	if err != nil && x == nil {
		fmt.Println("Error: ", err)
		t.Errorf("Error: %s %s", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		fmt.Println("Error: ", err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	if *detailed {
		Print(x)
	}
	increment()
}

// TestResponse - format output for test get and update
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - GET/POST/PATCH/PUT
// detailed *bool - bool to view the detailed response from API
// resourceID string - resource ID
func TestResponse(t *testing.T, x interface{}, err error, operation string, detailed *bool, increment func()) {
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error("Error: ", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		fmt.Println("Error: ", err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	if *detailed {
		Print(x)
	}
	increment()
}

// TestDeleteResponse - format output for test delete
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - DELETE
// detailed *bool - bool to view the detailed response from API
// resourceID string - resource ID
// statusCode int - status code from response
func TestDeleteResponse(t *testing.T, x interface{}, err error, operation string, statusCode int, detailed *bool, increment func()) {
	if err != nil && x == nil {
		fmt.Println("Error: ", err)
		t.Errorf("Error: %s %s", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		fmt.Println("Error: ", err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	t.Log("Status Code:", statusCode)
	if *detailed {
		Print(x)
	}
	increment()
}

// Counter - Count number of test run.
type Counter struct {
	count int
}

func (counter Counter) currentValue() int {
	return counter.count
}
func (counter *Counter) increment() {
	counter.count++
}
