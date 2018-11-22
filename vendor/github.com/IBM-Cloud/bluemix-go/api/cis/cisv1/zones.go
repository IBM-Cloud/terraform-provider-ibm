package cisv1

import (
	//"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
    "fmt"
    "log"
)


type ResultsCount struct {
      Count int `json:"count"`  
      }


// unresolved question as to whether codes always come back as ints
// zone and glb currently return 'code' as a string - issue raised

type Error struct {
      Code int `json:"code"`
      Msg string `json:"message"` 
      }  


type NameServer struct {
      NameS int64 `json:"0"`
      }

type Zone struct {
      Id string `json:"id"`
      Name string `json:"name"`
      Status string `json:"status"`
      Paused bool `json:"paused"`
      NameServers []string `json:"name_servers"`
      OriginalNameServer []string `json:"original_name_servers"`
      } 

type ZoneResults  struct {
      ZoneList []Zone `json:"result"`
      ResultsInfo ResultsCount `json:"result_info"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      }

  // Only one zone returned as json object rather than array in get case 
type ZoneResult  struct {
      Zone Zone `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

type ZoneBody struct {
      Name string `json:"name"`
    }


type ZoneDelete  struct {
      Result struct {
        zoneId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }




//Zones interface
type Zones interface {
	ListZones(cisId string) (*[]Zone, error)
	GetZone(cisId string, zoneId string) (*Zone, error)
	CreateZone(cisId string, zoneBody ZoneBody) (*Zone, error)
	DeleteZone(cisId string, zoneId string) (error)
	
}

type zones struct {
	client *client.Client
}

func newZoneAPI(c *client.Client) Zones {
	return &zones{
		client: c,
	}
}

func (r *zones)  ListZones(cisId string) (*[]Zone, error) {   
  log.Printf(">>>>>> Reached ListZones with %s\n", cisId)
  zoneResults := ZoneResults{}
  rawURL := fmt.Sprintf("/v1/%s/zones/", cisId)
  _, err := r.client.Get(rawURL, &zoneResults)
  if err != nil {
		return nil, err
	}
    return &zoneResults.ZoneList, err
}


func (r *zones)  GetZone(cisId string, zoneId string) (*Zone, error) {
  zoneResult := ZoneResult{}
  rawURL := fmt.Sprintf("/v1/" + cisId + "/zones/" + zoneId)
	_, err := r.client.Get(rawURL, &zoneResult, nil)
	if err != nil {
		return nil, err
	}
	return &zoneResult.Zone, nil
}



func  (r *zones) DeleteZone(cisId string, zoneId string) (error) {

	// Only call if the zone exists first. Otherwise API errors with a 
  // 403 if zone does not exist. Should return 404. Issue reported against CIS
  // DeleteZone is only called by resourceCISdomainDelete if zone exists. 

    rawURL := fmt.Sprintf("/v1/" + cisId + "/zones/" + zoneId)
    _, err := r.client.Delete(rawURL)
    if err != nil {
      return err
    }  
    return nil
}


func (r *zones)  CreateZone(cisId string, zoneBody ZoneBody) (*Zone, error) {
  zoneResult := ZoneResult{}		
	rawURL := "/v1/" + cisId + "/zones/" 
      _, err := r.client.Post(rawURL, &zoneBody, &zoneResult)
      if err != nil {
		return nil, err
	}   
	return &zoneResult.Zone, nil
}


