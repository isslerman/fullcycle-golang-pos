package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"ASDFASDF",
				"asdfaASDFASF",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100) // control channel
	errorFileUpload := make(chan string, 10)  // upload with errors channel

	go func() {
		for {
			select {
			case filename := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1) // read dir and return a slice
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading directory: %s\n", err)
			continue
		}
		// we can't just put go here, because it will run and die.
		// we need waitgroups! Calm down papaleguas!
		wg.Add(1)
		uploadControl <- struct{}{} // vamos enchendo o channel com um struct vazio. quando estiver cheio fica travado.
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {
	defer wg.Done()
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", completeFileName, s3Bucket)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("error opening file %s.\n", completeFileName)
		<-uploadControl                     // esvazia o canal em 1 unidade
		errorFileUpload <- completeFileName // joga o arquivo no channel de erros
		return
	}
	defer f.Close()
	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("error uploading file %s.\n", completeFileName)
		<-uploadControl // esvazia o canal em 1 unidade
		return
	}
	<-uploadControl // esvazia o canal em 1 unidade
	fmt.Printf("File %s uploaded sucessfully\n", completeFileName)
}
