package main

import (
    "flag"
    "log"
    "os"

    "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
    "github.com/IBM-Cloud/bluemix-go/session"
    "github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

    var cis_id string
    flag.StringVar(&cis_id, "cis_id", "", "CRN of the CIS service instance")

    flag.Parse()

    if cis_id == "" {
        flag.Usage()
        os.Exit(1)
    }

    trace.Logger = trace.NewLogger("true")
    sess, err := session.New()
    if err != nil {
        log.Fatal(err)
    }

    cisClient, err := cisv1.New(sess)
    if err != nil {
        log.Fatal(err)
    }
    monitorsAPI := cisClient.Monitors()

   
    log.Println(">>>>>>>>>  Monitor create")
    params := cisv1.MonitorBody{
                        ExpCodes: "200",
                        ExpBody: "",
                        Path:  "/status",   
                        MonType:  "http",
                        Method: "GET",
                        Timeout:  5,
                        Retries: 2,
                        Interval: 60, 
                        FollowRedirects: true,  
                        AllowInsecure: false,
                    }
    myMonitorPtr, err := monitorsAPI.CreateMonitor(cis_id, params)
       
    if err != nil {
        log.Fatal(err)
    }

    myMonitor := *myMonitorPtr
    monitorId := myMonitor.Id
    log.Println("Monitor create :", myMonitor)

    log.Println(">>>>>>>>>  Monitor read")
    myMonitorPtr, err = monitorsAPI.GetMonitor(cis_id, monitorId)

    if err != nil {
        log.Fatal(err)
    }

     myMonitor = *myMonitorPtr
                
    log.Println("Monitor Details by ID:", myMonitor)

    log.Println(">>>>>>>>>  Monitor delete")
    err = monitorsAPI.DeleteMonitor(cis_id, monitorId)
    if err != nil {
        log.Fatal(err)
    }
}