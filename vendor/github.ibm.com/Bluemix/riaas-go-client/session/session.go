package session

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"

	"github.ibm.com/riaas/rias-api/riaas/client"
	"github.ibm.com/riaas/rias-api/riaas/models"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

//const apiEndpointURL = "rias.wrig.me:5000"

// Session ...
type Session struct {
	IAMToken string
	IMSToken string
	Riaas    *client.Riaas
}

func riaasJSONConsumer() runtime.Consumer {
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
			errorRecord, isError := data.(*models.Riaaserror)
			if isError {
				errorRecord.Errors = make([]*models.RiaaserrorErrorsItems, 1, 1)
				errorRecord.Errors[0] = &models.RiaaserrorErrorsItems{
					Message:  string(b),
					Code:     "unexpected_return_value",
					MoreInfo: "",
					Target: &models.RiaaserrorErrorsItemsTarget{
						Name: "",
						Type: "",
					},
				}
			}
			return nil
		}
		return err
	})
}

// New ...
func New(iamtoken, apiEndpointURL string, debug bool) (*Session, error) {
	session := &Session{
		IAMToken: iamtoken,
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	transport := httptransport.New(apiEndpointURL, "/v1", []string{"https"})
	transport.Debug = debug
	transport.Consumers[runtime.JSONMime] = riaasJSONConsumer()
	session.Riaas = client.New(transport, nil)
	return session, nil
}

// Auth ...
func Auth(sess *Session) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", sess.IAMToken)
	})
}
