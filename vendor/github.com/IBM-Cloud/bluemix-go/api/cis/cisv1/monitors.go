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
      Messages []string `json:"messages"`
      }

    type MonitorResult  struct {
      Monitor Monitor `json:"result"`
      Success bool `json:"success"`
      Errors []Error `json:"errors"`
      Messages []string `json:"messages"`
      }    

// Temporary fix for Monitor Create API returning a string rather than an int

    type MonError struct {
      Code string `json:"code"`
      Msg string `json:"message"` 
      }  

    
    type monitorCreate  struct {
      Monitor Monitor `json:"result"`
      Success bool `json:"success"`
      Errors []MonError `json:"errors"`
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
	getMonitors(cisId string, iamToken string) (monitorList []Monitor, err error)
	getMonitor(cisId string, iamToken string, monitorId string) (monitor Monitor, err error)
	deleteMonitor(cisId string, iamToken string, monitorId string) (err error)
	createMonitor(cisId string, iamToken string, monitorBody MonitorBody) (Monitor, error)
}

type monitors struct {
	client *client.Client
}

func newMonitorAPI(c *client.Client) Monitors {
	return &monitors{
		client: c,
	}
}

func (r *monitors)  getMonitors(cisId string, iamToken string) (monitorList []Monitor, err error) {

    var getMonitorRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/monitors"

    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return []Monitor{}, errors.New("The HTTP Monitor GET request failed ")
    } else {
        getMonitorRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()
    
        rm := MonitorResults{}
        err = json.Unmarshal(getMonitorRes, &rm)
        if err != nil {

            if terr, ok := err.(*json.UnmarshalTypeError); ok {
                log.Printf("Failed to unmarshal field %s \n", terr.Field)
            } else {
                log.Println(err)
            }

            log.Printf("Failed to deserialise GET monitorResults Json %s\n", err)    
            return []Monitor{}, errors.New("Failed to deserialise monitorResults Json")
        }

        if rm.Success != true {
            errMsg := errorsToString(rm.Errors)
            return []Monitor{}, errors.New(errMsg)
        }
        monitorList = rm.MonitorList
        return 
}    


func (r *monitors)  getMonitor(cisId string, iamToken string, monitorId string) (monitor Monitor, err error) {
    log.Println("getMontor - entry")
    var getMonitorRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/monitors/" + monitorId
    //log.Printf("getMonitor URL: %v\n", urlstr)
    request, _ := http.NewRequest("GET", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return Monitor{}, errors.New("The HTTP Monitor GET request failed ")
    } else {
        if response.StatusCode == 404 {
            log.Printf("getMonitor - Monitor not found on Get, 404 code")  
            // Monitor not found using ID. Assumed it must have been deleted, so return with null
            monitor = Monitor{}
            return 
        }
        getMonitorRes, _ = ioutil.ReadAll(response.Body)
    }
    defer response.Body.Close()
    //log.Printf("Get Monitors %v\n", getMonitorRes)

    rm := MonitorResult{}
    err = json.Unmarshal(getMonitorRes, &rm)
    if err != nil {
        log.Printf("Failed to deserialise GET monitorResult  Json %s\n", err)    
        return Monitor{}, errors.New("Failed to deserialise monitorResult on Get Json  ")
    }
    if rm.Success != true {
        errMsg := errorsToString(rm.Errors)
        return Monitor{}, errors.New(errMsg)
    }
    monitor = rm.Monitor
    return 

}




func (r *monitors)  deleteMonitor(cisId string, iamToken string, monitorId string) (err error) {

    var deleteMonitorRes []byte    
    urlstr := "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/monitors/" + monitorId

    request, _ := http.NewRequest("DELETE", urlstr, nil)
    request.Header.Add("X-Auth-User-Token", iamToken)
    request.Header.Add("Accept", "application/json")
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        log.Printf("The HTTP request failed with error %s\n", err)
        return errors.New("The HTTP Monitor DELETE request failed ")
    } else {
        deleteMonitorRes, _ = ioutil.ReadAll(response.Body)
        //os.Stdout.Write(deleteMonitorRes)
    }
    defer response.Body.Close()
         if response.StatusCode == 404 {
            log.Printf("getMonitor - Monitor not found so must have already been deleted")  
            return nil
        }         
  
        rm := MonitorDelete{}
        err = json.Unmarshal(deleteMonitorRes, &rm)
        if err != nil {
            log.Printf("Failed to deserialise DELETE monitorResults Json %s\n", err)    
            return errors.New("Failed to deserialise monitorResults Json")
        }

        if rm.Success != true {
            errMsg := errorsToString(rm.Errors)
            return errors.New(errMsg)
        }

        return nil



}    


func (r *monitors)  createMonitor(cisId string, iamToken string, monitorBody MonitorBody) (Monitor, error) {

      monitorJson, err := json.Marshal(monitorBody) 
      if err != nil {
              log.Printf("createMonitor - Failed to serialise monitorBody Json %s\n", err)    
              return Monitor{}, errors.New("createMonitor - Failed to serialise monitorBody Json")
          }

      //monitorJson, err := json.Marshal(monitor{Name: domain})
      var createMonitorRes []byte 

      var urlstr string 
      urlstr = "https://api.cis.cloud.ibm.com/v1/" + cisId + "/load_balancers/monitors"
      request, _ := http.NewRequest("POST", urlstr, bytes.NewBuffer(monitorJson))
      request.Header.Add("X-Auth-User-Token", iamToken)
      request.Header.Add("Accept", "application/json")
      client := &http.Client{}
      response, err := client.Do(request)
      if err != nil {
          log.Printf("createMonitor - The HTTP request failed with error %s\n", err)
          return Monitor{}, errors.New("createMonitor - The HTTP request failed with error")
      } else {
          createMonitorRes, _ = ioutil.ReadAll(response.Body)
      }
      defer response.Body.Close()


        rm := monitorCreate{}
        err = json.Unmarshal(createMonitorRes, &rm)
        if err != nil {
            if terr, ok := err.(*json.UnmarshalTypeError); ok {
                log.Printf("Failed to unmarshal field %s \n", terr.Field)
            } else {
                log.Println(err)
            }
            log.Printf("createMonitor - Failed to deserialise CREATE monitorResults Json %s\n", err)    
            return Monitor{}, errors.New("createMonitor - Failed to deserialise monitorResults Json")
        }


        if rm.Success != true {
            errMsg := monErrorsToString(rm.Errors)
            return Monitor{}, errors.New(errMsg)
        }



        monitor := rm.Monitor

        log.Printf("createMonitor - Return: %v\n", monitor.Id) 

        return monitor, nil

}
