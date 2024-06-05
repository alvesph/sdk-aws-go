package s3

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ListS3(bucketName string) ([]string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)

	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})

	if err != nil {
		log.Fatal(err)
	}

	var formattedResults []string
	for _, obj := range output.Contents {
		formattedResult := fmt.Sprintf("key=%s size=%d", aws.ToString(obj.Key), obj.Size)
		formattedResults = append(formattedResults, formattedResult)
	}

	return formattedResults, nil
}
