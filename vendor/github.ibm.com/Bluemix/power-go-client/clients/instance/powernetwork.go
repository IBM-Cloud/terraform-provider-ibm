package instance


import

(
"github.ibm.com/Bluemix/power-go-client/session"
"github.ibm.com/Bluemix/power-go-client/power/models"
"github.ibm.com/Bluemix/power-go-client/errors"
"github.ibm.com/Bluemix/power-go-client/power/client/p_cloud_networks"
"log"
)

type PowerNetworkClient struct {

	session *session.Session
}


// NewPowerNetworkClient ...
func NewPowerNetworkClient(sess *session.Session) *PowerNetworkClient {
	return &PowerNetworkClient{
		sess,
	}
}


func (f *PowerNetworkClient) Get(id string) (*models.Network, error) {

	params := p_cloud_networks.NewPcloudNetworksGetParams().WithCloudInstanceID(f.session.PowerServiceInstance).WithNetworkID(id)
	resp,err := f.session.Power.PCloudNetworks.PcloudNetworksGet(params,session.NewAuth(f.session))

	if err != nil || resp.Payload == nil  {
		log.Printf("Failed to perform the operation... %v",err)
		return nil, errors.ToError(err)
	}
	return resp.Payload, nil
}


func (f *PowerNetworkClient) Create(name string,networktype string,cidr string,dnsservers []string, gateway string,startip string,endip string) (*models.Network,*models.Network,error){


	var ipbody = []*models.IPAddressRange{
		{&endip, &startip,},
	}

	var body = models.NetworkCreate{}
	body.Cidr = cidr
	body.Gateway = gateway
	body.IPAddressRanges = ipbody
	body.Name = name
	body.DNSServers = dnsservers
	body.Type = &networktype

	log.Printf("Printing the body %+v",body)
	params := p_cloud_networks.NewPcloudNetworksPostParamsWithTimeout(f.session.Timeout).WithCloudInstanceID(f.session.PowerServiceInstance).WithBody(&body)
	_,resp,err := f.session.Power.PCloudNetworks.PcloudNetworksPost(params,session.NewAuth(f.session))

	//log.Printf("The error is %d ",resp.Payload.VlanID)
	if err != nil {
		return nil,nil, errors.ToError(err)
	}

	if resp != nil{
		log.Printf("Failed to create the network ")
	}
	//if postok != nil {
	//	log.Print("Request failed ")
	//}

	return resp.Payload, nil, nil
}