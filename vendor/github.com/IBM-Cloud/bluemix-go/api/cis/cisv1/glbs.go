package cisv1

import (
	//"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
    "fmt"
    //"log"
    "time"
)


 // unresolved question as to whether codes always come back as ints
// glb and glb currently return 'code' as a string - issue raised




type Glb struct {
      Id string `json:"id"`
      Name string `json:"name"`
      Desc string `json:"description"`
      FallbackPool string `json:"fallback_pool"`
      DefaultPools []string `json:"default_pools"`
      Ttl int `json:"ttl"`
      Proxied bool `json:"proxied"`
      CreatedOn    *time.Time  `json:"created_on,omitempty"`
      ModifiedOn   *time.Time  `json:"modified_on,omitempty"`
      SessionAffinity string `json:"session_affinity"`
      // RegionPools  map[string][]string `json:"region_pools"`
      // PopPools     map[string][]string `json:"pop_pools"`
    }

type GlbResults  struct {
      GlbList []Glb `json:"result"`
      ResultsInfo ResultsCount `json:"result_info"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      }

  // Only one glb returned as json object rather than array in get case 
type GlbResult  struct {
      Glb Glb `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

type GlbBody struct {
      Desc string `json:"description,omitempty"`
      Proxied bool `json:"proxied,omitempty"`
      Name string `json:"name"`
      FallbackPool string `json:"fallback_pool"`
      DefaultPools []string `json:"default_pools"`
      SessionAffinity string `json:"session_affinity,omitempty"`             
      }


type GlbDelete  struct {
      Result struct {
        glbId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }




//Glbs interface
type Glbs interface {
	ListGlbs(cisId string, zoneId string) (*[]Glb, error)
	GetGlb(cisId string, zoneId string, glbId string) (*Glb, error)
	CreateGlb(cisId string, zoneId string, glbBody GlbBody) (*Glb, error)
	DeleteGlb(cisId string, zoneId string, glbId string) (error)
	
}

type glbs struct {
	client *client.Client
}

func newGlbAPI(c *client.Client) Glbs {
	return &glbs{
		client: c,
	}
}

func (r *glbs)  ListGlbs(cisId string, zoneId string) (*[]Glb, error) {   
  glbResults := GlbResults{}
  rawURL := fmt.Sprintf("/v1/%s/zones/%s/load_balancers", cisId, zoneId)
  _, err := r.client.Get(rawURL, &glbResults)
  if err != nil {
		return nil, err
	}
    return &glbResults.GlbList, err
}


func (r *glbs)  GetGlb(cisId string, zoneId string, glbId string) (*Glb, error) {
  glbResult := GlbResult{}
  rawURL := fmt.Sprintf("/v1/%s/zones/%s/load_balancers/%s", cisId, zoneId, glbId)
	_, err := r.client.Get(rawURL, &glbResult, nil)
	if err != nil {
		return nil, err
	}
	return &glbResult.Glb, nil
}



func  (r *glbs) DeleteGlb(cisId string, zoneId string, glbId string) (error) {

	// Only call if the glb exists first. Otherwise API errors with a 
  // 403 if glb does not exist. Should return 404. Issue reported against CIS
  // DeleteGlb is only called by resourceCISdomainDelete if glb exists. 

    rawURL := fmt.Sprintf("/v1/%s/zones/%s/load_balancers/%s", cisId, zoneId, glbId)
    _, err := r.client.Delete(rawURL)
    if err != nil {
      return err
    }  
    return nil
}


func (r *glbs)  CreateGlb(cisId string, zoneId string, glbBody GlbBody) (*Glb, error) {
  glbResult := GlbResult{}		
	rawURL := fmt.Sprintf("/v1/%s/zones/%s/load_balancers", cisId, zoneId)
      _, err := r.client.Post(rawURL, &glbBody, &glbResult)
      if err != nil {
		return nil, err
	}   
	return &glbResult.Glb, nil
}


