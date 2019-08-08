package session

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.ibm.com/Bluemix/power-go-client/power/client"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/utils"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

//const apiEndpointURL = "rias.wrig.me:5000"

const(

	offering="power-iaas"
	crnString="crn"
	version = "v1"
	service="bluemix"
	serviceType="public"
	serviceInstanceSeparator="/"
	separator=":"


//var crn = "crn:v1:bluemix:public:power-iaas:us-east:a/ba6042d7f84a4a318f64003e691bf700:d16705bd-7f1a-48c9-9e0e-1c17b71e7331::"
)

// Session ...
type Session struct {
	IAMToken   string
	IMSToken   string
	Power 	   *client.PowerIaas
	Timeout    time.Duration
	Generation int64
	PowerServiceInstance string
	UserAccount string
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
			//if isError {
			//	errorRecord.Errors = make([]*models.RiaaserrorErrorsItems, 1, 1)
			//	errorRecord.Errors[0] = &models.RiaaserrorErrorsItems{
			///		Message:  string(b),
				//	Code:     "unexpected_return_value",
				//	MoreInfo: "",
				//	Target: &models.RiaaserrorErrorsItemsTarget{
				//		Name: "",
				//		Type: "",
				//	},
				//}
		//	}
			return nil
		}
		return err
	})
}

// New ...
func New(iamtoken, region , powerinstance string, generation int, debug bool, timeout time.Duration,useraccount string) (*Session, error) {
	session := &Session{
		IAMToken: iamtoken,
		PowerServiceInstance: powerinstance,
		UserAccount: useraccount,
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	apiEndpointURL := utils.GetEndpoint(generation, region)
	log.Printf("the apiendpoint url for power is %s",apiEndpointURL)
	transport := httptransport.New(apiEndpointURL, "/", []string{"https"})
	transport.Debug = debug
	transport.Consumers[runtime.JSONMime] = powerJSONConsumer()
	session.Power = client.New(transport, nil)
	session.Timeout = timeout
	session.Generation = int64(generation)
	return session, nil
}



func NewAuth(sess *Session) runtime.ClientAuthInfoWriter {
	log.Printf("Calling the New Auth Method")
	var crndata = crnBuilder(sess.PowerServiceInstance,sess.UserAccount)
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization",sess.IAMToken); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crndata)
	})


}



func BearerTokenAndCRN(session *Session, crn string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		if err := r.SetHeaderParam("Authorization", session.IAMToken); err != nil {
			return err
		}
		return r.SetHeaderParam("CRN", crn)
	})
}

func crnBuilder(powerinstance ,useraccount string) string{
	//log.Printf ("Calling the crn constructor that is to be passed back to the caller  %s",useraccount)
	var crnData=crnString+separator+version+separator+service+separator+serviceType+separator+offering+separator+"us-east"+separator+"a"+serviceInstanceSeparator+useraccount+separator+powerinstance+separator+separator
	//log.Printf("the crndata is ... %s ",crnData)
	return crnData
}