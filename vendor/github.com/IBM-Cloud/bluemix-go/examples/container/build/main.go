package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/IBM-Cloud/bluemix-go/api/container/registryv1"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/trace"
)

func tarGzContext(context string) (string, error) {
	tarfile, err := ioutil.TempFile("", "docker-*.tar.gz")
	if err != nil {
		return "", err
	}
	defer tarfile.Close()
	gw := gzip.NewWriter(tarfile)
	defer gw.Close()

	tarball := tar.NewWriter(gw)
	defer tarball.Close()

	info, err := os.Stat(context)
	if err != nil {
		return "", err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(context)
	}

	err = filepath.Walk(context,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, context))
			}

			if err := tarball.WriteHeader(header); err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(tarball, file)
			return err
		})
	if err != nil {
		return "", err
	}
	return tarfile.Name(), err
}

//Example: ./build -f Dockerfile -t registry.ng.bluemix.net/ibmcloud-go/imagtest .")
func main() {

	c := new(bluemix.Config)

	var imageTag string
	// should same form of registry.ng.bluemix.net/<namespace>/<imagename>
	flag.StringVar(&imageTag, "t", "registry.ng.bluemix.net/ibmcloud-go/build", "tag")

	var dockerFile string
	flag.StringVar(&dockerFile, "f", "Dockerfile", "Dockerfile")

	var region string
	flag.StringVar(&region, "region", "us-south", "region")

	flag.Parse()

	c.Region = region
	directory := flag.Args()

	fmt.Println(directory)
	trace.Logger = trace.NewLogger("true")
	if len(directory) < 1 || directory[0] == "" {

		flag.Usage()
		fmt.Println("Example: ./build -f Dockerfile -t registry.ng.bluemix.net/ibmcloud-go/imagtest .")
		os.Exit(1)
	}

	session, _ := session.New(c)

	iamAPI, err := iamv1.New(session)
	identityAPI := iamAPI.Identity()
	userInfo, err := identityAPI.UserInfo()
	if err != nil {
		log.Fatal(err)
	}
	registryClient, err := registryv1.New(session)
	if err != nil {
		log.Fatal(err)
	}

	namespaceHeaderStruct := registryv1.NamespaceTargetHeader{
		AccountID: userInfo.Account.Bss,
	}

	namespace := strings.Split(imageTag, "/")
	if len(namespace) != 3 {
		log.Fatal("Image Tag not correct format")
	}

	namespaces, err := registryClient.Namespaces().GetNamespaces(namespaceHeaderStruct)
	found := false
	for _, a := range namespaces {
		if a == namespace[1] {
			found = true
			break
		}
	}
	if !found {
		_, err := registryClient.Namespaces().AddNamespace(namespace[1], namespaceHeaderStruct)
		if err != nil {
			log.Fatal(err)
		}

	}

	headerStruct := registryv1.BuildTargetHeader{
		AccountID: userInfo.Account.Bss,
	}

	requestStruct := registryv1.DefaultImageBuildRequest()
	requestStruct.T = imageTag
	requestStruct.Dockerfile = dockerFile

	tarName, err := tarGzContext(directory[0])
	if err != nil {

	}
	tarFile, err := os.Open(tarName)
	if err != nil {
		log.Fatal(err)
	}

	tarReader := bufio.NewReader(tarFile)
	//Too much binary output
	trace.Logger = trace.NewLogger("false")

	fmt.Println("Building...")
	err = registryClient.Builds().ImageBuild(*requestStruct, tarReader, headerStruct, os.Stdout)

	imageHeaderStruct := registryv1.ImageTargetHeader{
		AccountID: userInfo.Account.Bss,
	}

	fmt.Println("\nInspecting Built Image...")
	image, err := registryClient.Images().InspectImage(imageTag, imageHeaderStruct)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.MarshalIndent(image, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(jsonBytes))

	fmt.Println("\nScanning Built Image...")
	imageVulnerabiliyRequst := registryv1.DefaultImageVulnerabilitiesRequest()
	imageReport, err := registryClient.Images().ImageVulnerabilities(imageTag, *imageVulnerabiliyRequst, imageHeaderStruct)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err = json.MarshalIndent(imageReport, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(jsonBytes))
}
