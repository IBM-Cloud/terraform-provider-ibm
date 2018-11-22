package cisv1

import (
	//"fmt"
  "log"
	"github.com/IBM-Cloud/bluemix-go/client"
   "errors"
    "io/ioutil"
    "net/http"
     "encoding/json"
    "bytes"
    //"strconv"
)


type Glb struct {
      Id string `json:"id"`
      Name string `json:"name"`
      Desc string `json:"description"`
      FallbackPool string `json:"fallback_pool"`
      DefaultPools []string `json:"default_pools"`
      Ttl int `json:"ttl"`
      Proxied bool `json:"proxied"`
      }
    type GlbResults  struct {
      GlbList []Glb `json:"result"`
      ResultsInfo ResultsCount `json:"result_info"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

    type GlbResult  struct {
      Glb Glb `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }

    type glbCreate  struct {
      Glb Glb `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
    }


    type GlbBody struct {
      Desc string `json:"description"`
      Proxied bool `json:"proxied"`
      Name string `json:"name"`
      FallbackPool string `json:"fallback_pool"`
      DefaultPools []string `json:"default_pools"`             
      }

  type GlbDelete  struct {
       Result struct {
        glbId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
    }



//Zones interface
type Glbs interface {
	getGlbs(cisId string, iamToken string, zoneId string) (glbList []Glb, err error)
	getGlb(cisId string, iamToken string, zoneId string, glbId string) (glb Glb, err error)
	deleteGlb(cisId string, iamToken string, zoneId string, glbId string) (err error)
	createGlb(cisId string, iamToken string, zoneId string, glbBody GlbBody) (glb Glb, err error)
}

type glbs struct {
	client *client.Client
}

func newGlbAPI(c *client.Client) Glbs {
	return &glbs{
		client: c,
	}
}

func (r *glbs)  getGlbs(cisId string, iamToken string, zoneId string) ( glbList []Glb, err error) {

    var getGlbRes []byte 
    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/zones/" + zoneId + "/load_balancers"

    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
        if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return []Glb{}, errors.New("The HTTP Get request failed ")
    } else {
        getGlbRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()


      rg := GlbResults{}
      err = json.Unmarshal(getGlbRes, &rg)
        if err != nil {
            log.Printf("Failed to deserialise GET glbResults Json %s\n", err)    
            return []Glb{}, errors.New("Failed to deserialise glbResults Json")
        }


        if rg.Success != true {
            errMsg := errorsToString(rg.Errors)
            return []Glb{}, errors.New(errMsg)
        }
   
    glbList = rg.GlbList

 
    

    return

 }   

func (r *glbs)  getGlb(cisId string, iamToken string, zoneId string, glbId string) ( glb Glb, err error) {

    var getGlbRes []byte 
    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/zones/" + zoneId + "/load_balancers/" + glbId

    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
        if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return Glb{}, errors.New("The HTTP Get request failed ")
    } else {
        if response.StatusCode == 404 {
            // Zone not found using ID. Assumed it must have been deleted, so return
            glb = Glb{}
            return 
        }
        getGlbRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()


      rg := GlbResult{}
      err = json.Unmarshal(getGlbRes, &rg)
        if err != nil {
            log.Printf("Failed to deserialise GET glbResults Json %s\n", err)    
            return Glb{}, errors.New("Failed to deserialise glbResults Json")
        }


        if rg.Success != true {
            errMsg := errorsToString(rg.Errors)
            return Glb{}, errors.New(errMsg)
        }
   
    glb = rg.Glb

 
    

    return

 }   


func (r *glbs)  deleteGlb(cisId string, iamToken string, zoneId string, glbId string)(err error) {

    var deleteGlbRes []byte 

    ///////////////////////////////////////////////////////////////
    //  Only supports look up of one zone. Standard Plan
    ///////////////////////////////////////////////////////////////
    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/zones/" + zoneId + "/load_balancers/" + glbId

    request, _ := http.NewRequest("DELETE", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
        if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return errors.New("The HTTP Delete request failed ")
    } else {
        deleteGlbRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()
    log.Printf("deleteGlb  %v\n", response.StatusCode)  
        // Work around for erroneous Glb Delete respomse when zome already deleted. 
        if response.StatusCode == 404 || response.StatusCode == 403 {
            log.Printf("deleteGlb - Gone not found so must have already been deleted")  
            return nil
        }    

      rg := GlbDelete{}
      err = json.Unmarshal(deleteGlbRes, &rg)
        if err != nil {
            log.Printf("Failed to deserialise DELETE glbResults Json %s\n", err)    
            return errors.New("Failed to deserialise glbResults Json")
        }


        if rg.Success != true {
            errMsg := errorsToString(rg.Errors)
            return errors.New(errMsg)
        }

 
    

    return

 }   


func (r *glbs)  createGlb(cisId string, iamToken string, zoneId string, glbBody GlbBody) (glb Glb, err error) {


    glbJson, err := json.Marshal(glbBody)

    if err != nil {
            log.Printf("Failed to serialise glbBody Json %s\n", err)    
            return Glb{}, errors.New("Failed to serialise glbBody Json")
          }

// log.Println("error:", glbJson)
// os.Stdout.Write(glbJson)
   
      //poolJson, err := json.Marshal(pool{Name: domain})
      var createGlbRes []byte 

      var urlstr string 
      urlstr = "https://api.cis.cloud.ibm.com/v1/" + cisId + "/zones/" + zoneId + "/load_balancers"
      request, _ := http.NewRequest("POST", urlstr, bytes.NewBuffer(glbJson))
      request.Header.Add("X-Auth-User-Token", iamToken)
      request.Header.Add("Accept", "application/json")
      client := &http.Client{}
      response, err := client.Do(request)
      if err != nil {
          log.Printf("The HTTP request failed with error %s\n", err)
          return Glb{}, errors.New("The HTTP request failed with error")
      } else {
          createGlbRes, _ = ioutil.ReadAll(response.Body)
      }
      defer response.Body.Close()
      // log.Println(string("createGlbRes"))
      // log.Println(string(createGlbRes))
      

        ri := glbCreate{}
        err = json.Unmarshal(createGlbRes, &ri)
        if err != nil {
            log.Printf("Failed to deserialise CREATE glbResults Json %s\n", err)    
            return Glb{}, errors.New("Failed to deserialise glbResults Json")
        }


        if ri.Success != true {
            errMsg := errorsToString(ri.Errors)
            return Glb{}, errors.New(errMsg)
        }

        glb = ri.Glb

    return 
}


