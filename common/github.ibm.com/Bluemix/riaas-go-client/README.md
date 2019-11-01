# GO client for the RIaaS api

The swagger spec is at [spec/swagger.yaml](spec/swagger.yaml). This is a v2.0 spec built to mirror the original riaas spec at https://pages.github.ibm.com/riaas/api-spec/. It has the same properties, but is missing validation rules.

If any changes are made to the real spec, they need to be added into this spec file too. If the doc is modified, install [go-swagger](https://github.com/go-swagger/go-swagger) and run `sh gen.sh` to generate the stubs. The generated stubs are in the [generated](generated) folder.

Helper clients are available in the [clients](clients) directory.

example usage:

```
package main

import (
        "fmt"

        "github.ibm.com/Bluemix/riaas-go-client/clients/compute"
        "github.ibm.com/Bluemix/riaas-go-client/session"
)

func main() {

        IAM_TOKEN := "eyJraWQ............................."
        sess, err := session.New(IAM_TOKEN)

        flavorC := compute.NewFlavorClient(sess)

        x, err := flavorC.List()
        fmt.Println(*x[0].Name)
        fmt.Println(*x[0].Href)
        fmt.Println(*x[0].MaxBandwidth)

        y, err := flavorC.Get(*x[0].Name)
        fmt.Println(y.CPU.Frequency)
}
```
