package main

import (
	"flag"
	"log"
	"os"

	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func main() {

	var icdId string
	var userId string
	flag.StringVar(&icdId, "icdId", "", "CRN of the IBM Cloud Database service instance")
	flag.StringVar(&userId, "userId", "", "Userid for IBM Cloud Database connection")
	flag.Parse()

	if icdId == "" || userId == "" {
		flag.Usage()
		os.Exit(1)
	}

	trace.Logger = trace.NewLogger("true")
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}

	icdClient, err := icdv4.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	connectionAPI := icdClient.Connections()
	connection, err := connectionAPI.GetConnection(icdId, userId)

	if err != nil {
		log.Fatal(err)
	}
	//log.Println("Connection :",connection)
	if connection.Postgres.Type != "" {
		for _, l := range connection.Postgres.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type ", connection.Postgres.Type)
		log.Println("Scheme: ", connection.Postgres.Scheme)
		log.Println("Path: ", connection.Postgres.Path)
		log.Println("Hosts.Name: ", connection.Postgres.Hosts[0].HostName)
		log.Println("Hosts.Ports: ", connection.Postgres.Hosts[0].Port)
		log.Println("Query Options: ", connection.Postgres.QueryOptions)

		log.Println("Method: ", connection.Postgres.Authentication.Method)
		log.Println("UserName: ", connection.Postgres.Authentication.UserName)
		log.Println("Password: ", connection.Postgres.Authentication.Password)
		log.Println("Name: ", connection.Postgres.Certificate.Name)
		log.Println("CertificateBase64: ", connection.Postgres.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Postgres.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}

	if connection.Grpc.Type != "" {
		for _, l := range connection.Postgres.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type ", connection.Grpc.Type)
		log.Println("Scheme: ", connection.Grpc.Scheme)
		log.Println("Path: ", connection.Grpc.Path)
		log.Println("Hosts.Name: ", connection.Grpc.Hosts[0].HostName)
		log.Println("Hosts.Ports: ", connection.Grpc.Hosts[0].Port)
		log.Println("Query Options: ", connection.Grpc.QueryOptions)

		log.Println("Method: ", connection.Grpc.Authentication.Method)
		log.Println("UserName: ", connection.Grpc.Authentication.UserName)
		log.Println("Password: ", connection.Grpc.Authentication.Password)
		log.Println("Name: ", connection.Grpc.Certificate.Name)
		//log.Println("CertificateBase64: ", connection.Grpc.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Grpc.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}

	if connection.Rediss.Type != "" {
		for _, l := range connection.Postgres.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type: ", connection.Rediss.Type)
		log.Println("Scheme: ", connection.Rediss.Scheme)
		log.Println("Path: ", connection.Rediss.Path)
		log.Println("Hosts.Name: ", connection.Rediss.Hosts[0].HostName)
		log.Println("Hosts.Ports: ", connection.Rediss.Hosts[0].Port)
		log.Println("Query Options: ", connection.Rediss.QueryOptions)

		log.Println("Method: ", connection.Rediss.Authentication.Method)
		log.Println("UserName: ", connection.Rediss.Authentication.UserName)
		log.Println("Password: ", connection.Rediss.Authentication.Password)
		log.Println("Name: ", connection.Rediss.Certificate.Name)
		//log.Println("CertificateBase64: ", connection.Rediss.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Rediss.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}
	if connection.Https.Type != "" {
		for _, l := range connection.Postgres.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type: ", connection.Https.Type)
		log.Println("Scheme: ", connection.Https.Scheme)
		log.Println("Path: ", connection.Https.Path)
		log.Println("Hosts.Name: ", connection.Https.Hosts[0].HostName)
		log.Println("Hosts.Ports: ", connection.Https.Hosts[0].Port)
		log.Println("Query Options: ", connection.Https.QueryOptions)

		log.Println("Method: ", connection.Https.Authentication.Method)
		log.Println("UserName: ", connection.Https.Authentication.UserName)
		log.Println("Password: ", connection.Https.Authentication.Password)
		log.Println("Name: ", connection.Https.Certificate.Name)
		//log.Println("CertificateBase64: ", connection.Https.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Https.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}
	if connection.Amqps.Type != "" {
		for _, l := range connection.Postgres.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type: ", connection.Amqps.Type)
		log.Println("Scheme: ", connection.Amqps.Scheme)
		log.Println("Path: ", connection.Amqps.Path)
		log.Println("Hosts.Name: ", connection.Amqps.Hosts[0].HostName)
		log.Println("Hosts.Ports: ", connection.Amqps.Hosts[0].Port)
		log.Println("Query Options: ", connection.Amqps.QueryOptions)

		log.Println("Method: ", connection.Amqps.Authentication.Method)
		log.Println("UserName: ", connection.Amqps.Authentication.UserName)
		log.Println("Password: ", connection.Amqps.Authentication.Password)
		log.Println("Name: ", connection.Amqps.Certificate.Name)
		//log.Println("CertificateBase64: ", connection.Amqps.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Amqps.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}
	if connection.Mongo.Type != "" {
		for _, l := range connection.Mongo.Composed {
			log.Println("Composed: ", l)
		}
		log.Println("Type ", connection.Mongo.Type)
		log.Println("Scheme: ", connection.Mongo.Scheme)
		log.Println("Path: ", connection.Mongo.Path)
		for _, l := range connection.Mongo.Hosts {
			log.Println("Hosts.Name: ", l.HostName)
			log.Println("Hosts.Ports: ", l.Port)
		}
		log.Println("Query Options: ", connection.Mongo.QueryOptions)

		log.Println("Method: ", connection.Mongo.Authentication.Method)
		log.Println("UserName: ", connection.Mongo.Authentication.UserName)
		log.Println("Password: ", connection.Mongo.Authentication.Password)
		log.Println("Name: ", connection.Mongo.Certificate.Name)
		log.Println("CertificateBase64: ", connection.Mongo.Certificate.CertificateBase64)
		log.Println("Database: ", connection.Mongo.Database)
		log.Println("Composed: ", connection.Cli.Composed[0])
		log.Println("Type: ", connection.Cli.Type)
		log.Println("Bin: ", connection.Cli.Bin)
		log.Println("Arguments: ", connection.Cli.Arguments[0][0])
		//log.Println("CertificateBase64: ",connection.Cli.Certificate.CertificateBase64)
	}
}
