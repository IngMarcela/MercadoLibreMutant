package repository

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type NotificationRepository struct {
	SnsAws snsiface.SNSAPI
}

func (n NotificationRepository) SendNotification(mutant string) error {
	pubInput := &sns.PublishInput{
		Message:  aws.String(mutant),
		TopicArn: aws.String("arn:aws:sns:us-east-1:548546023079:Mutant"),
	}

	_, err := n.SnsAws.Publish(pubInput)
	if err != nil {
		log.Println(err.Error())

		return err
	}
	return nil
}

func NewNotificationRepository(
	SnsAws snsiface.SNSAPI,
) *NotificationRepository {
	return &NotificationRepository{
		SnsAws: SnsAws,
	}
}
