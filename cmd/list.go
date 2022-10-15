/*
Copyright © 2022 Jan Feddern
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/spf13/cobra"
)

var bucketName string
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List objects",
	Long:  `This command lists all objects for a given S3 Bucket`,
	Run: func(cmd *cobra.Command, args []string) {
		flagBucketName, _ := cmd.Flags().GetString("bucket")
		err := listObjectsInS3Bucket(flagBucketName)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&bucketName, "bucket", "b", "", "Name of the S3 Bucket")
	listCmd.MarkFlagRequired("bucket")
	rootCmd.AddCommand(listCmd)
}

func initS3Client() (*s3.Client, error) {

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-central-1"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return client, nil
}

func listObjectsInS3Bucket(bucketName string) error {
	client, err := initS3Client()
	if err != nil {
		panic(err)
	}

	resp, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Number of objects inside the Bucket: %v \n", resp.KeyCount)
	i := 0
	for i < int(resp.KeyCount) {
		fmt.Printf("Object Name %v: %s \n", i, *resp.Contents[i].Key)
		i++
	}
	return nil
}
