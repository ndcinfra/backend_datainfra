package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/YoungsoonLee/backend_datainfra/libs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Controller ...
type S3Controller struct {
	BaseController
}

const (
	S3REGION = "ap-northeast-2"
	S3BUCKET = "ndc-pm-resources"
)

// UploadImage ...
func (s *S3Controller) UploadImage() {
	//var Buf bytes.Buffer
	fileHeader, err := s.GetFiles("file")
	if err != nil {
		fmt.Println("fileheader error: ", err)
	}

	f, err := fileHeader[0].Open()
	creds := credentials.NewSharedCredentials("", "default")

	s3, err := session.NewSession(&aws.Config{
		Region:      aws.String(S3REGION),
		Credentials: creds,
	})
	if err != nil {
		s.ResponseError(libs.ErrS3Session, err)
	}

	err = AddFilesToS3(s3, fileHeader[0].Filename, fileHeader[0].Size, f)
	if err != nil {
		//
		s.ResponseError(libs.ErrS3AddFile, err)
	}

	// return url
	s.ResponseSuccess("", "https://s3.ap-northeast-2.amazonaws.com/ndc-pm-resources/"+fileHeader[0].Filename)

}

// AddFilesToS3 ...
func AddFilesToS3(s *session.Session, fileName string, size int64, r io.Reader) error {

	buffer := make([]byte, size)
	r.Read(buffer)

	// rObj, err := s3.New(s).PutObject(&s3.PutObjectInput{
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket: aws.String(S3BUCKET),
		Key:    aws.String(fileName),
		ACL:    aws.String("public-read"),
		Body:   bytes.NewReader(buffer),
		//ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err

}
