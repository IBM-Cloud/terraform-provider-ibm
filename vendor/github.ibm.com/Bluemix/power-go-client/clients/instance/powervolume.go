package instance


import

(
	"github.ibm.com/Bluemix/power-go-client/session"
	"github.ibm.com/Bluemix/power-go-client/power/models"
	"github.ibm.com/Bluemix/power-go-client/errors"
	"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_volumes"
	"log"
)

type PowerVolumeClient struct {

	session *session.Session

}


// NewPowerVolumeClient ...
func NewPowerVolumeClient(sess *session.Session) *PowerVolumeClient {
	return &PowerVolumeClient{
		sess,
	}
}


//Get information about a single volume only
func (f *PowerVolumeClient) Get(id string) (*models.Volume, error) {



	//var cloudinstanceid = f.session.PowerServiceInstance

	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesGetParams().WithCloudInstanceID(f.session.PowerServiceInstance).WithVolumeID(id)
	resp,err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesGet(params,session.NewAuth(f.session))

	if err != nil || resp.Payload == nil  {
		log.Printf("Failed to perform the operation... %v",err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}

//Create

func( f *PowerVolumeClient) Create(volumename string,volumesize float64,volumetype string,volumeshareable bool) (*models.Volume, error) {

	log.Printf("calling the PowerVolume Create Method")

	var body = models.CreateDataVolume{
		Name: &volumename,
		Size: &volumesize,
		DiskType: &volumetype,
		Shareable: &volumeshareable,

	}


	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithBody(&body)
	resp, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPost(params,session.NewAuth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	return resp.Payload, nil
}



// Delete ...
func (f *PowerVolumeClient) Delete(id string) error {
	//var cloudinstanceid = f.session.PowerServiceInstance
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesDeleteParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithVolumeID(id)
	_, err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesDelete(params, session.NewAuth(f.session))
	if err != nil {
		return errors.ToError(err)
	}
	return nil
}

// Update..
func(f *PowerVolumeClient) Update(id, volumename string,volumesize float64,volumeshare bool) (*models.Volume, error){

	var cloudinstanceid = f.session.PowerServiceInstance
	var patchbody = models.UpdateVolume{}
	patchbody.Name=&volumename
	patchbody.Size=volumesize
	patchbody.Shareable =&volumeshare
	params := p_cloud_volumes.NewPcloudCloudinstancesVolumesPutParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(cloudinstanceid).WithVolumeID(id).WithBody(&patchbody)

	resp,err := f.session.Power.PCloudVolumes.PcloudCloudinstancesVolumesPut(params,session.NewAuth(f.session))

	if err != nil {
		return nil, errors.ToError(err)
	}

	log.Print("Printing the response data .. %+v",resp.Payload)
	return resp.Payload, nil
}