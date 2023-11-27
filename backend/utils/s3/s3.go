package s3

import (
	"context"
	// "errors"
	 "fmt"
	
	//"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const accessKey = "6muhsu0u51ad6rf9"
const secretKey = "8295474f-f94e-41cf-81ae-2dddbd7b6cdf"
const apiUrlValue = "storage.iran.liara.space"

var s3Client *s3.Client

func Init() error {
	// svc := s3.New(session.New(), &aws.Config{Region: aws.String("iran")})

	cred := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cnf, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(cred))
	if err!=nil{
		return err
	}

	s3Client = s3.NewFromConfig(cnf, s3.WithEndpointResolver(s3.EndpointResolverFromURL(apiUrlValue)),
			func (opts *s3.Options){
				opts.UsePathStyle = true
			})

	// if err != nil{
		// return errors.New("can't create s3Client.")
	// }

	fmt.Println(s3Client)


	return nil
}