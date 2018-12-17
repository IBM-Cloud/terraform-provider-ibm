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

    var zone_id string
    flag.StringVar(&zone_id, "zone_id", "", "zone_id for zone")

    flag.Parse()

    if zone_id == "" || cis_id == "" {
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
    settingsAPI := cisClient.Settings()

    log.Println(">>>>>>>>>  Zone Settings read")
    mySettingsPtr, err := settingsAPI.GetSetting(cis_id, zone_id, "min_tls_version")

    if err != nil {
        log.Fatal(err)
    }

    mySettings := *mySettingsPtr
                
    log.Println("Zone Settings by ID:", mySettings)

    log.Println(">>>>>>>>>  Zone Settings Update")
    params := cisv1.SettingsBody{
                        Value: "1.2",
                    }
    mySettingsPtr, err = settingsAPI.UpdateSetting(cis_id, zone_id, "min_tls_version", params)
    if err != nil {
        log.Fatal(err)
    }
    
    mySettings = *mySettingsPtr
                
    log.Println("Zone Setting on update:", mySettings)


}