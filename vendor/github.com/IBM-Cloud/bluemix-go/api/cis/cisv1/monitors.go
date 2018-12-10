package cisv1

import (
	"github.com/IBM-Cloud/bluemix-go/client"
    "fmt"
)

type Monitor struct {
      Id string `json:"id"`
      Path string `json:"path"`
      ExpBody string `json:"expected_body"`
      ExpCodes string `json:"expected_codes"`
      // Headers omitted future enhancement
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
      // Headers ommited TBC
      MonType string `json:"type,omitempty"`            
      Method string `json:"method,omitempty"`           
      Timeout int `json:"timeout,omitempty"`          
      Retries int `json:"retries,omitempty"`          
      Interval int `json:"interval,omitempty"`         
      FollowRedirects bool `json:"follow_redirects,omitempty"` 
      AllowInsecure bool `json:"allow_insecure,omitempty"`   
    }


type MonitorDelete  struct {
      Result struct {
        monitorId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

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
  rawURL := fmt.Sprintf("/v1/%s/load_balancers/monitors/", cisId)
  _, err := r.client.Get(rawURL, &monitorResults)
  if err != nil {
		return nil, err
	}
    return &monitorResults.MonitorList, err
}

func (r *monitors)  GetMonitor(cisId string, monitorId string) (*Monitor, error) {
  monitorResult := MonitorResult{}
  rawURL := fmt.Sprintf("/v1/%s/load_balancers/monitors/%s", cisId, monitorId)
	_, err := r.client.Get(rawURL, &monitorResult, nil)
	if err != nil {
		return nil, err
	}
	return &monitorResult.Monitor, nil
}

func  (r *monitors) DeleteMonitor(cisId string, monitorId string) (error) {
    rawURL := fmt.Sprintf("/v1/%s/load_balancers/monitors/%s", cisId, monitorId)
    _, err := r.client.Delete(rawURL)
    if err != nil {
      return err
    }  
    return nil
}

func (r *monitors)  CreateMonitor(cisId string, monitorBody MonitorBody) (*Monitor, error) {
  monitorResult := MonitorResult{}		
	rawURL := fmt.Sprintf("/v1/%s/load_balancers/monitors/", cisId)
      _, err := r.client.Post(rawURL, &monitorBody, &monitorResult)
      if err != nil {
		return nil, err
	}   
	return &monitorResult.Monitor, nil
}


