package external_api

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	uuid "github.com/satori/go.uuid"
)

func Uploader() (string, error) {
	// https://jto.nyc3.digitaloceanspaces.com
	// The session the S3 Uploader will use
	endpoint := "sgp1.digitaloceanspaces.com"
	region := "sgp1"
	key := "G6L46EVCCZBKBPRPRRTW"
	secret := "VoFpZgX8j4Erh7lulOv82P7NcbDfbm9ALBNauI2WWSc"

	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    &endpoint,
		Region:      &region,
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	filename := "kitten123.jpg"
	// filename := "image1.jpg"
	// var content string = `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`
	// fs, err := base64.StdEncoding.DecodeString(content)
	// if err != nil {
	// return fmt.Errorf("failed to open file %q, %v", filename, err)
	// }

	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// original_image, _, err := image.Decode(bytes.NewReader(fs))
	// if err != nil {
	// 	fmt.Println("failed to create buffer", err)
	// }

	// new_image := resize.Resize(160, 0, original_image, resize.Lanczos3)
	// buff := new(bytes.Buffer)

	// // encode image to buffer
	// err = png.Encode(buff, new_image)
	// if err != nil {
	// 	fmt.Println("failed to create buffer", err)
	// }
	// reader := bytes.NewReader(buff.Bytes())
	myuuid := uuid.NewV4().String()

	myBucket := "suumii-dev-space"
	myString := myuuid + ".jpg"
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:         aws.String("public-read"),
		Bucket:      aws.String(myBucket),
		Key:         aws.String(myString),
		ContentType: aws.String("image/jpg"),
		Body:        f,
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
	return aws.StringValue(&result.Location), err
}
