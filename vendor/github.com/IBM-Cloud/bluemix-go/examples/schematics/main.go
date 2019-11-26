package main

import (
	"fmt"
	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/session"

	sch "github.com/IBM-Cloud/bluemix-go/api/schematics"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {
	c := new(bluemix.Config)

	var workspaceID string
	var templateid string
	fmt.Print("Enter Workspace ID: ")
	fmt.Scanln(&workspaceID)
	fmt.Print("Enter Template ID: ")
	fmt.Scanln(&templateid)

	trace.Logger = trace.NewLogger("true")
	var payload = sch.Payload{
		Name:        "bbb",
		Type:        []string{"terraform-v1.0"},
		Description: "terraform workspace",
		Tags:        []string{"department:HR", "application:compensation", "environment:staging"},
		WorkspaceStatus: sch.WorkspaceStatus{
			Frozen: true,
		},
		TemplateRepo: sch.TemplateRepo{
			URL: "https://github.com/ptaube/tf_cloudless_sleepy",
		},
		TemplateRef: "ibm-open-liberty-2ae855ce3ca4",
		TemplateData: []sch.TemplateData{
			{
				Folder: ".",
				Type:   "terraform-v1.0",

				Variablestore: []sch.Variablestore{
					{
						Name:        "sample_var",
						Secure:      false,
						Value:       "THIS IS IBM CLOUD TERRAFORM CLI DEMO",
						Description: "Description of sample_var",
					},
					{
						Name:  "sleepy_time",
						Value: "15",
					},
				},
			},
		},
	}

	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}
	schClient, err := sch.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	schAPI := schClient.Workspaces()
	//Get the workspace
	works, err := schAPI.GetWorkspaceByID(workspaceID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nThe workspace info= ", works)

	state, err := schAPI.GetStateStore(workspaceID, templateid)
	if err != nil {
		log.Fatal(err)
	}
	statestr := fmt.Sprintf("%v", state)
	fmt.Println("The state info= ", statestr)

	out, err := schAPI.GetOutputValues(workspaceID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The output info= ", out)
	items := make(map[string]interface{})

	for _, feilds := range out {
		if feilds.TemplateID == "653f60a4-f64f-41" {
			output := feilds.Output

			for _, value := range output {
				for key, val := range value {
					val2 := val.Value
					items[key] = val2

				}
			}
		}
	}

	fmt.Println("\nThe array info= ", items)

	cdata, err := schAPI.CreateWorkspace(payload)

	fmt.Println("\ndata=", cdata)
	fmt.Println("\nid=", cdata.ID)

}
