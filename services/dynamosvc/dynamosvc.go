package dynamosvc

import (
	"fmt"

	"github.com/yongchengchen/sysgo/contract"
	"github.com/yongchengchen/sysgo/services/container"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mitchellh/mapstructure"
)

// Config ...
type DYNAMODB_CONFIG struct {
	REGION string `json:"region"`
	ID     string `json:"id"`
	SECRET string `json:"secret"`
	TOKEN  string `json:"token"`
}

var awsConfig *aws.Config
var dyConfig DYNAMODB_CONFIG

func loadConfig() *aws.Config {
	if awsConfig == nil {
		if conf, ok := container.Get("config").(contract.IConfig); ok {
			configs := conf.Get("dynamodb.config")

			if err := mapstructure.Decode(configs, &dyConfig); err != nil {
				fmt.Printf("dynamodb config is not correct %#v\n", configs)
			}

			awsConfig = &aws.Config{
				Region: aws.String(dyConfig.REGION),
				Credentials: credentials.NewStaticCredentials(dyConfig.ID,
					dyConfig.SECRET, dyConfig.TOKEN),
			}
		}
	}

	return awsConfig
}

func PutRecord(table string, record interface{}) {

	sess, _ := session.NewSession(loadConfig())
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(record)
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}
	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
