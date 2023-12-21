package s3

import (
	"io"
	"errors"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	// "github.com/aws/aws-sdk-go/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

const accessKey = "6muhsu0u51ad6rf9"
const secretKey = "8295474f-f94e-41cf-81ae-2dddbd7b6cdf"
const apiUrlValue = "https://storage.iran.liara.space"


var s3Client *s3.S3
var bucket = aws.String("thlearn")

func Init() error {
	// svc := s3.New(session.New(), &aws.Config{Region: aws.String("iran")})
	// bucket := aws.String("thlearn")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint: aws.String(apiUrlValue),
		Region: aws.String("ir"),
		DisableSSL: aws.Bool(false),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession := session.New(s3Config)

	s3Client = s3.New(newSession)

	if s3Client == nil{
		return errors.New("can't create s3 client")
	}

	return nil
	// cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	// cnf, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(cred))
	// if err!=nil{
	// 	return err
	// }

	// s3Client = s3.NewFromConfig(cnf, s3.WithEndpointResolver(s3.EndpointResolverFromURL(apiUrlValue)),
	// 		func (opts *s3.Options){
	// 			opts.UsePathStyle = true
	// 		})

	// // if err != nil{
	// 	// return errors.New("can't create s3Client.")
	// // }

	// fmt.Println(s3Client)


	// return nil
}

func PutObject(body io.ReadSeeker, key string) error{
	_, err := s3Client.PutObject(&s3.PutObjectInput{
		Body: body,
		Bucket: bucket,
		Key: &key,
	})

	return err
}

func DeleteObject(key string) error {
	_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: bucket,
		Key: aws.String(key),
	})

	return err
}