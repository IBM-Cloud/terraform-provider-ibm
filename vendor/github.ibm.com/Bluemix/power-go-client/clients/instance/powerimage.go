package instance


import

(
	"github.ibm.com/Bluemix/power-go-client/session"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/errors"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_images"

	//"github.ibm.com/Bluemix/power-go-client/utils"
	"log"
)

type PowerImageClient struct {

	session *session.Session
}


// NewPowerImageClient ...
func NewPowerImageClient(sess *session.Session) *PowerImageClient {
	return &PowerImageClient{
		sess,
	}
}


func (f *PowerImageClient) Get(id string) (*models.Image, error) {

	

	var cloudinstanceid = f.session.PowerServiceInstance
	
	params := p_cloud_images.NewPcloudCloudinstancesImagesGetParams().WithCloudInstanceID(cloudinstanceid).WithImageID(id)
	resp,err := f.session.Power.PCloudImages.PcloudCloudinstancesImagesGet(params,session.NewAuth(f.session))

	if err != nil || resp.Payload == nil  {
		log.Printf("Failed to perform the operation... %v",err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}