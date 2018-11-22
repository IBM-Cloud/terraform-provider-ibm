package cisv1

import (
	//"fmt"
	"github.com/IBM-Cloud/bluemix-go/client"
  "errors"
    "log"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "bytes"
)


 type Pool struct {
      Id string `json:"id"`
      Desc string `json:"description"`
      Name string `json:"name"`
      CheckRegions []string `json:"check_regions"`
      Enabled bool `json:"enabled"`
      MinOrigins int `json:"minimum_origins"`
      Monitor string `json:"monitor"`
      NotEmail string `json:"notification_email"`
      Origins []Origin `json:"origins"`
      Healthy bool `json:"healthy"`
      CreatedOn string `json:"created_on"`
      ModifiedOn string `json:"modified_on"`
      }

    type CheckRegion struct {
      Region string `json:"0"`
    }


    type Origin struct { 
      Name string `json:"name"`
      Address string `json:"address"`
      Enabled bool `json:"enabled"`
      Weight int `json:"weight"`
      }  

    type PoolResults  struct {
      PoolList []Pool `json:"result"`
      ResultsInfo ResultsCount `json:"result_info"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }    

    type PoolResult  struct {
      Pool Pool `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }    




    // Only one pool returned as json object rather than array in get case 
    type poolCreate  struct {
      Pool Pool `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }



   
    type PoolBody struct {
      Name string `json:"name"`
      Desc string `json:"description"`
      Notification string `json:"notification_email"` 
      Regions []string `json:"check_regions"`
      Origins []Origin `json:"origins"`
      CheckRegions []string `json:"check_regions"`
      Enabled bool `json:"enabled"`
      MinOrigins int `json:"minimum_origins"`
      Monitor string `json:"monitor"`
      NotEmail string `json:"notification_email"`
      }


 type PoolDelete  struct {
      Result struct {
        poolId string
        } `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }




//Pools interface
type Pools interface {
	getPools(cisId string, iamToken string) (poolList []Pool, err error)
	getPool(cisId string, iamToken string, poolId string) (pool Pool, err error)
	deletePool(cisId string, iamToken string, poolId string) (err error)
	createPool(cisId string, iamToken string, poolBody PoolBody) (Pool, error)
}

type pools struct {
	client *client.Client
}

func newPoolAPI(c *client.Client) Pools {
	return &pools{
		client: c,
	}
}

func (r *pools)  getPools(cisId string, iamToken string) ( poolList []Pool, err error) {

    var getPoolRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/pools"

    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return []Pool{}, errors.New("The HTTP Pool Get request failed")
    } else {
        getPoolRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()



      rp := PoolResults{}
      err = json.Unmarshal(getPoolRes, &rp)
        if err != nil {
            log.Printf("Failed to deserialise GET poolBody Json %s\n", err)  
            return []Pool{}, errors.New("Failed to deserialise poolBody Json")
        }


        if rp.Success != true {
            errMsg := errorsToString(rp.Errors)
            return []Pool{}, errors.New(errMsg)
        }

    poolList = rp.PoolList

    return 
}    

func (r *pools)  getPool(cisId string, iamToken string, poolId string) ( pool Pool, err error) {

    var getPoolRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/pools/" + poolId

    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return Pool{}, errors.New("The HTTP Pool Get request failed")
    } else {
        if response.StatusCode == 404 {
            // Zone not found using ID. Assumed it must have been deleted, so return
            pool = Pool{}
            return 
        }
        getPoolRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()

      rp := PoolResult{}
      err = json.Unmarshal(getPoolRes, &rp)
        if err != nil {
            log.Printf("Failed to deserialise GET poolBody Json %s\n", err)  
            return Pool{}, errors.New("Failed to deserialise poolBody Json")
        }
        if rp.Success != true {
            errMsg := errorsToString(rp.Errors)
            return Pool{}, errors.New(errMsg)
        }

  pool = rp.Pool

    return 
}



func (r *pools)  deletePool(cisId string, iamToken string, poolId string) (err error) {

    var deletePoolRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/pools/" + poolId

    request, _ := http.NewRequest("DELETE", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return errors.New("The HTTP Pool Delete request failed")
    } else {
        deletePoolRes, _ = ioutil.ReadAll(response.Body)
        //os.Stdout.Write(deletePoolRes)
    }
    defer response.Body.Close()
        log.Printf("deletePool  %v\n", response.StatusCode)  
        if response.StatusCode == 404 || response.StatusCode == 403 {
            log.Printf("deletePool - Pool not found so must have already been deleted")  
            return nil
        }    


      rp := PoolDelete{}
      err = json.Unmarshal(deletePoolRes, &rp)
        if err != nil {
            log.Printf("Failed to deserialise DELETE deletePoolRes Json %s\n", err)  
            return errors.New("Failed to deserialise deletePoolRes Json")
        }


        if rp.Success != true {
            errMsg := errorsToString(rp.Errors)
            return errors.New(errMsg)
        }


    return 
}    


func (r *pools) createPool(cisId string, iamToken string, poolBody PoolBody) (pool Pool, err error) { 
    defer func() {
      if r := recover(); r != nil {
         err = r.(error)
      }
   }()

    poolJson, err := json.MarshalIndent(poolBody, "", "\t")

    if err != nil {
            log.Printf("Failed to serialise poolBody Json %s\n", err)  
            return Pool{}, errors.New("Failed to serialise poolBody Json")
            }

// log.Println("error:", poolJson)
// os.Stdout.Write(poolJson)


      var createPoolRes []byte 

      var urlstr string 
      urlstr = "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/pools"
      request, _ := http.NewRequest("POST", urlstr, bytes.NewBuffer(poolJson))
      request.Header.Add("X-Auth-User-Token", iamToken)
      request.Header.Add("Accept", "application/json")
      client := &http.Client{}
      response, err := client.Do(request)
      if err != nil {
          log.Printf("The HTTP request failed with error %s\n", err)
          return Pool{}, errors.New("The HTTP request failed with error")
      } else {
          createPoolRes, _ = ioutil.ReadAll(response.Body)
      }
      defer response.Body.Close()


        rp := poolCreate{}
        err = json.Unmarshal(createPoolRes, &rp)
        if err != nil {
            log.Printf("Failed to deserialise CREATE poolResults Json %s\n", err)  
            return Pool{}, errors.New("Failed to deserialise poolBody Json")
        }


        if rp.Success != true {
            errMsg := errorsToString(rp.Errors)
            return Pool{}, errors.New(errMsg)
        }


        pool = rp.Pool

    return pool, nil
 }   


