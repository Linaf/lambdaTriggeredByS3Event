package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"os"
)


type  Record struct {
	eventVersion string `json:"event_version"`
	eventSource string `json:"event_source"`
	awsRegion string `json:"aws_region"`
	eventTime string `json:"event_time"`
	eventName string `json:"event_name"`
	userIdentity UserIdentity `json:"user_identity"`
	requestParameters RequestParameters `json:"requestParameters"`
	responseElements ResponseElements `json:"responseElements"`
	s3 S3 `json:"s3"`
	object Object `json:"object"`

}


type UserIdentity struct{
	principalId string `json:"principalId"`
}
type RequestParameters struct{
	sourceIPAddress string `json:"sourceIPAddress"`
}

type ResponseElements struct{
	xAmzRequestId string `json:"x_amz_request_id"`
	xAmzId2 string `json:"x-amz-id-2"`

}

type S3 struct{
	s3SchemaVersion string `json:"s3SchemaVersion"`
	configurationId string `json:"configurationId"`
	bucket Bucket `json:"bucket"`
	arn string `json:"arn"`
}
type Bucket struct {
	name string `json:"name"`
	ownerIdentity OwnerIdentity `json:"ownerIdentity"`
}
type OwnerIdentity struct{
	principalId string `json:"principalId"`
}
type Object struct{
	key string `json:"key"`
	size int64 `json:"size"`
	eTag string `json:"eTag"`
	sequencer string `json:"sequencer"`

}
func handler(c context.Context, e events.SNSEvent) {
	for _, record := range e.Records {
		snsRecord := record.SNS
		snsMessage := snsRecord.Message
		log.Infof("Passed message is %v", snsMessage)
		//message := snsRecord.Message
		 var messageRecord = Record{}
		err := json.Unmarshal([]byte(snsMessage),&messageRecord)
		if err != nil {
			log.Errorf("Error unmarshalling the SNS record ")
		}

	}
}
func main() {
	cloudenv := os.Getenv("CLOUD_ENVIRONMENT")
	log.Infof("Started executing  lambda triggered by SNS published due to S3 event in %s", cloudenv)
	lambda.Start(handler)
}
