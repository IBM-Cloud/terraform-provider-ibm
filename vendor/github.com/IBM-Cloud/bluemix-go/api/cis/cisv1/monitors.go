package cisv1

import (
	//"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
    "fmt"
    //"log"
)




// unresolved question as to whether codes always come back as ints
// monitor and glb currently return 'code' as a string - issue raised




type Monitor struct {
      Id string `json:"id"`
      Path string `json:"path"`
      ExpBody string `json:"expected_body"`
      ExpCodes string `json:"expected_codes"`
      MonType string `json:"type"`
      Method string `json:"method"`
      Timeout int `json:"timeout"`          
      Retries int `json:"retries"`          
      Interval int `json:"interval"`         
      FollowRedirects bool `json:"follow_redirects"` 
      AllowInsecure bool `json:"allow_insecure"`   
      }

type MonitorResults  struct {
      MonitorList []Monitor `json:"result"`
      ResultsInfo ResultsCount `json:"result_info"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      }

  // Only one monitor returned as json object rather than array in get case 
type MonitorResult  struct {
      Monitor Monitor `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

type MonitorBody struct {
        ExpCodes string `json:"expected_codes"`
        ExpBody string `json:"expected_body"`
        Path string `json:"path"`
        // golang objects to the use of type here as it is a reseved keyword. 
        MonType string `json:"type"`            
        Method string `json:"method"`           
        Timeout int `json:"timeout"`          
        Retries int `json:"retries"`          
        Interval int `json:"interval"`         
        FollowRedirects bool `json:"follow_redirects"` 
        AllowInsecure bool `json:"allow_insecure"`   
    }


type MonitorDelete  struct {
      Result struct {
        monitorId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }




//Monitors interface
type Monitors interface {
	ListMonitors(cisId string) (*[]Monitor, error)
	GetMonitor(cisId string, monitorId string) (*Monitor, error)
	CreateMonitor(cisId string, monitorBody MonitorBody) (*Monitor, error)
	DeleteMonitor(cisId string, monitorId string) (error)
	
}

type monitors struct {
	client *client.Client
}

func newMonitorAPI(c *client.Client) Monitors {
	return &monitors{
		client: c,
	}
}

func (r *monitors)  ListMonitors(cisId string) (*[]Monitor, error) {   
  monitorResults := MonitorResults{}
  rawURL := fmt.Sprintf("/v1/%s/monitors/", cisId)
  _, err := r.client.Get(rawURL, &monitorResults)
  if err != nil {
		return nil, err
	}
    return &monitorResults.MonitorList, err
}


func (r *monitors)  GetMonitor(cisId string, monitorId string) (*Monitor, error) {
  monitorResult := MonitorResult{}
  rawURL := fmt.Sprintf("/v1/%s/monitors/%s", cisId, monitorId)
	_, err := r.client.Get(rawURL, &monitorResult, nil)
	if err != nil {
		return nil, err
	}
	return &monitorResult.Monitor, nil
}



func  (r *monitors) DeleteMonitor(cisId string, monitorId string) (error) {

	// Only call if the monitor exists first. Otherwise API errors with a 
  // 403 if monitor does not exist. Should return 404. Issue reported against CIS
  // DeleteMonitor is only called by resourceCISdomainDelete if monitor exists. 

    rawURL := fmt.Sprintf("/v1/%s/monitors/%s", cisId, monitorId)
    _, err := r.client.Delete(rawURL)
    if err != nil {
      return err
    }  
    return nil
}


func (r *monitors)  CreateMonitor(cisId string, monitorBody MonitorBody) (*Monitor, error) {
  monitorResult := MonitorResult{}		
	rawURL := fmt.Sprintf("/v1/%s/monitors/", cisId)
      _, err := r.client.Post(rawURL, &monitorBody, &monitorResult)
      if err != nil {
		return nil, err
	}   
	return &monitorResult.Monitor, nil
}


