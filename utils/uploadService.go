package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

func RunUploadServer() {
	r := gin.Default()

	key := os.Getenv("DIGITALOCEAN_SPACE_KEY")
	secret := os.Getenv("DIGITALOCEAN_SPACE_SECRET")
	bucket_name := os.Getenv("DIGITALOCEAN_SPACE_BUCKETNAME")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:    aws.String("https://ams3.digitaloceanspaces.com"),
		Region:      aws.String("us-east-1"),
	}

	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)

	r.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		content, err := file.Open()
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		byteContainer, err := ioutil.ReadAll(content)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		object := s3.PutObjectInput{
			Bucket: aws.String(bucket_name),
			Key:    aws.String(file.Filename),
			Body:   strings.NewReader(string(byteContainer)),
			ACL:    aws.String("private"),
			Metadata: map[string]*string{
				"x-amz-meta-my-key": aws.String("your-value"), //required
			},
		}

		result, err := s3Client.PutObject(&object)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(200, gin.H{
			"message": result,
		})
	})

	r.GET("/upload", func(c *gin.Context) {
		fmt.Println("getFile")
		input := &s3.ListObjectsInput{
			Bucket: aws.String(bucket_name),
		}

		objects, err := s3Client.ListObjects(input)
		if err != nil {
			fmt.Println(err.Error())
		}

		c.JSON(200, gin.H{
			"message": objects.Contents,
		})

	})

	r.Run()
}
