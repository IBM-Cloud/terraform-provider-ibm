/*
Code to call the IBM IAM Services and get a session object that will be used by the Power Colo Code.


*/

package ibmpisession

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/IBM-Cloud/power-go-client/power/client"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM-Cloud/power-go-client/utils"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	//"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/go-openapi/strfmt"
)

const (
	offering                 = "power-iaas"
	crnString                = "crn"
	version                  = "v1"
	service                  = "bluemix"
	serviceType              = "public"
	serviceInstanceSeparator = "/"
	separator                = ":"

	//var crn = "crn:v1:bluemix:public:power-iaas:us-east:a/ba6042d7f84a4a318f64003e691bf700:d16705bd-7f1a-48c9-9e0e-1c17b71e7331::"
)

// Session ...
type IBMPISession struct {
	IAMToken    string
	IMSToken    string
	Power       *client.PowerIaas
	Timeout     time.Duration
	UserAccount string
	Region      string
	Zone        string
}

func powerJSONConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		/*t := reflect.TypeOf(data)
		if data != nil && t.Kind() == reflect.Ptr {
			v := reflect.Indirect(reflect.ValueOf(data))
			if t.Elem().Kind() == reflect.String {
				buf := new(bytes.Buffer)
				_, err := buf.ReadFrom(reader)
				if err != nil {
					return err
				}
				b := buf.Bytes()
				v.SetString(string(b))
				return nil
			}
		}*/
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(reader)
		if err != nil {
			return err
		}
		b := buf.Bytes()
		if b != nil {
			dec := json.NewDecoder(bytes.NewReader(b))
			dec.UseNumber() // preserve number formats
			err = dec.Decode(data)
		}
		if string(b) == "null" || err != nil {
			errorRecord, _ := data.(*models.Error)
			log.Printf("The errorrecord is %s ", errorRecord)
			return nil
		}
		return err
	})
}

// New ...
/*
The method takes in the following params
iamtoken : this is the token that is passed from the client
region : Obtained from the terraform template. Every template /resource will be required to have this information
timeout:
useraccount:
*/
func New(iamtoken, region string, debug bool, timeout time.Duration, useraccount string, zone string) (*IBMPISession, error) {
	session := &IBMPISession{
		IAMToken:    iamtoken,
		UserAccount: useraccount,
		Region:      region,
		Zone:        zone,
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	apiEndpointURL := utils.GetPowerEndPoint(region)
	log.Printf("the apiendpoint url for power is %s", apiEndpointURL)
	transport := httptransport.New(apiEndpointURL, "/", []string{"https"})
	if debug {
		transport.Debug = debug
	}
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.Power = client.New(transport, nil)
	session.Timeout = timeout
	return session, nil
}

func NewAuth(sess *IBMPISession, PowerInstanceId string) runtime.ClientAuthInfoWriter {
	log.Printf("Calling the New Auth Method in the IBMPower Session Code")
	//var generatedCRN := crn.New(
	//var region :=
	var crndata = crnBuilder(PowerInstanceId, sess.UserAccount, sess.Region, sess.Zone)
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization", sess.IAMToken); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crndata)
	})

}

func BearerTokenAndCRN(session *IBMPISession, crn string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization", session.IAMToken); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crn)
	})
}

func crnBuilder(powerinstance, useraccount, region string, zone string) string {
	log.Printf("Calling the crn constructor that is to be passed back to the caller  %s", useraccount)
	log.Printf("the region is %s", region)
	var crnData string
	//var crnData = crnString + separator + version + separator + service + separator + serviceType + separator + offering + separator + "us-south" + separator + "a" + serviceInstanceSeparator + useraccount + separator + powerinstance + separator + separator
	if zone == "" {
		crnData = crnString + separator + version + separator + service + separator + serviceType + separator + offering + separator + region + separator + "a" + serviceInstanceSeparator + useraccount + separator + powerinstance + separator + separator
	} else {
		crnData = crnString + separator + version + separator + service + separator + serviceType + separator + offering + separator + zone + separator + "a" + serviceInstanceSeparator + useraccount + separator + powerinstance + separator + separator
	}

	log.Printf("the crndata is ... %s ", crnData)
	return crnData
}
