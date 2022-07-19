package repository

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockSNSClient struct {
	snsiface.SNSAPI
}

func (m *mockSNSClient) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	return nil, nil
}

func Test(t *testing.T) {
	notificationRepository := NewNotificationRepository(&mockSNSClient{})

	err := notificationRepository.SendNotification("ATGCA CCGTA TAAAG AAAAT CTTAC TCAGG")

	assert.Equal(t, nil, err)
}
