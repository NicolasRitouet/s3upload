package command

import (
  "fmt"
  "github.com/goamz/goamz/aws"
  "github.com/goamz/goamz/s3"
  "log"
	"github.com/codegangsta/cli"
	"io/ioutil"
	//"os"
//	"net/http"
//	"bufio"
)

var bucketname = "mycs-operations"
var foldername = "test/"

func CmdUpload(c *cli.Context) {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	// uploadFile("READsME.md")

  auth, err := aws.EnvAuth()
  euwest := aws.EUWest
	fmt.Println(auth)

  connection := s3.New(auth, euwest)
  mybucket := connection.Bucket(bucketname)
  res, err := mybucket.List(foldername, "", "", 1000)
  if err != nil {
      log.Fatal(err)
  }
  for _, v := range res.Contents {
      fmt.Println("https://s3-eu-west-1.amazonaws.com/" + bucketname + "/" + v.Key)
  }
}
/**
func uploadFile(filename string) {

	auth, err := aws.EnvAuth()
	if err != nil {
		panic(err.Error())
	}

	// Open Bucket
	s := s3.New(auth, aws.APNortheast)
	bucket := s.Bucket(bucketname)

	data := []byte("Hello, Goamz!!")
	err = bucket.Put("sample.txt", data, "text/plain", s3.BucketOwnerFull)
	if err != nil {
		panic(err.Error())
	}
}
**/
