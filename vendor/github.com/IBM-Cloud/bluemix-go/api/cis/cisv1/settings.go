package cisv1

import (
	//"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
    "fmt"
    )



type SettingsResults struct {
      Waf SettingsResObj 
      Ssl SettingsResObj
      Min_tls_version SettingsResObj
      Tls_1_3 SettingsResObj
      }


type SettingsResult  struct {
      Result SettingsResObj`json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

type SettingsResObj struct {
      Id string `json:"id"`
      Value string `json:"value"`
      Editable bool `json:"editable"`
      ModifiedDate string `json:"modified_on"`
      CertificateStatus string `json:"certificate_status"`

}

type SettingsBody struct {
      Value string `json:"value"`
    }


//Settingss interface
type Settings interface {
	GetSettings(cisId string, settingsId string) (*SettingsResults, error)
	UpdateSettings(cisId string, zoneId string, setting string, settingsBody SettingsBody) (*SettingsResObj, error)
	
}

type settings struct {
	client *client.Client
}

func newSettingsAPI(c *client.Client) Settings {
	return &settings{
		client: c,
	}
}

func (r *settings)  GetSettings(cisId string, zoneId string) (*SettingsResults, error) {   
  settingsResult := SettingsResult{}
  settingsResults := SettingsResults{}
  rawURL := fmt.Sprintf("/v1/%s/zones/%s/settings/waf", cisId, zoneId)
  _, err := r.client.Get(rawURL, &settingsResult)
  if err != nil {
		return nil, err
	}
  settingsResults.Waf = settingsResult.Result

  rawURL = fmt.Sprintf("/v1/%s/zones/%s/settings/ssl", cisId, zoneId)
  _, err = r.client.Get(rawURL, &settingsResult)
  if err != nil {
    return nil, err
  }
  settingsResults.Ssl = settingsResult.Result
  rawURL = fmt.Sprintf("/v1/%s/zones/%s/settings/min_tls_version", cisId, zoneId)
  _, err = r.client.Get(rawURL, &settingsResult)
  if err != nil {
    return nil, err
  }
  settingsResults.Min_tls_version = settingsResult.Result
  // rawURL = fmt.Sprintf("/v1/%s/zones/%s/settings/tls_1_3_only", cisId, zoneId)
  // _, err = r.client.Get(rawURL, &settingsResult)
  // if err != nil {
  //   return nil, err
  // }
  // settingsResults.Tls_1_3 = settingsResult.Result

  return &settingsResults, err
}



func (r *settings)  UpdateSettings(cisId string, zoneId string, setting string, settingsBody SettingsBody) (*SettingsResObj, error) {
  settingsResult := SettingsResult{}		
	rawURL := fmt.Sprintf("/v1/%s/zones/%s/settings/%s", cisId, zoneId, setting)
      _, err := r.client.Patch(rawURL, &settingsBody, &settingsResult)
      if err != nil {
		return nil, err
	}   
	return &settingsResult.Result, nil
}


