package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	s3Client s3iface.S3API
)

func init() {
	sess := session.Must(session.NewSession())
	s3Client = s3.New(sess)
}

// UploadObject method for upload elements in S3
func UploadObject(ctx context.Context, bucket, path string, data []byte) error {

	input := &s3.PutObjectInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(bucket),
		Key:    aws.String(path),
	}

	_, err := s3Client.PutObjectWithContext(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
