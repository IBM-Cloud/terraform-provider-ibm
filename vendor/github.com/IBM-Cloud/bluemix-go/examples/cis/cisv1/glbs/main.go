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
    var pool_id string
    flag.StringVar(&pool_id, "pool_id", "", "pool_id for zone")
    var domain_name string
    flag.StringVar(&domain_name, "domain_name", "", "domain_name for zone")


    flag.Parse()

    if zone_id == "" || cis_id == "" || pool_id == "" || domain_name == "" {
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
    glbsAPI := cisClient.Glbs()

   
    log.Println(">>>>>>>>>  Glbs create")
    params := cisv1.GlbBody{
                        Name: domain_name,
                        SessionAffinity: "none",
                        DefaultPools: []string {
                           pool_id,
                        },    
                        FallbackPool: pool_id,
                        Proxied: true,
                    }
    myGlbsPtr, err := glbsAPI.CreateGlb(cis_id, zone_id, params)
       
    if err != nil {
        log.Fatal(err)
    }

    myGlbs := *myGlbsPtr
    glbsId := myGlbs.Id
    log.Println("Glbs create :", myGlbs)

    log.Println(">>>>>>>>>  Glbs read")
    myGlbsPtr, err = glbsAPI.GetGlb(cis_id, zone_id, glbsId)

    if err != nil {
        log.Fatal(err)
    }

     myGlbs = *myGlbsPtr
                
    log.Println("Glbs Details by ID:", myGlbs)

    log.Println(">>>>>>>>>  Glbs delete")
    err = glbsAPI.DeleteGlb(cis_id, zone_id, glbsId)
    if err != nil {
        log.Fatal(err)
    }
}