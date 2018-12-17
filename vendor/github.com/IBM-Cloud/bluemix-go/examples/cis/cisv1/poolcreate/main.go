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
    poolsAPI := cisClient.Pools()

   
    log.Println(">>>>>>>>>  Pool create")
    origins := []cisv1.Origin {
                        { Name: "eu-origin1", Address: "150.0.0.1", Enabled: true, Weight: 1 },
                        { Name: "eu-origin2", Address: "150.0.0.2", Enabled: true, Weight: 1 },
                        }
                checkRegions := []string{"EEU"}
    params := cisv1.PoolBody{
                    Name: "eu-pool",
                    Origins: origins,             
                    CheckRegions: checkRegions,
                    Enabled: true,
                    MinOrigins: 1,
                }
    myPoolPtr, err := poolsAPI.CreatePool(cis_id, params)
       
    if err != nil {
        log.Fatal(err)
    }

    myPool := *myPoolPtr
    poolId := myPool.Id
    log.Println("Pool create :", myPool)

    log.Println(">>>>>>>>>  Pool read")
    myPoolPtr, err = poolsAPI.GetPool(cis_id, poolId)

    if err != nil {
        log.Fatal(err)
    }

}