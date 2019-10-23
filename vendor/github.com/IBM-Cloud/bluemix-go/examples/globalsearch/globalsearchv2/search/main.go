package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/globalsearch/globalsearchv2"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var query string
	flag.StringVar(&query, "query", "", "Query string")
	flag.Parse()

	if query == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	gsClient, err := globalsearchv2.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	searchBody := globalsearchv2.SearchBody{
		Query:  query,
		Fields: []string{"name", "service_name", "tags"},
	}

	searchResult, err := gsClient.Searches().PostQuery(searchBody)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Result: ", searchResult)

	log.Println("Items: ", searchResult.Items)

}
