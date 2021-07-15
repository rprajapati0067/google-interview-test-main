package services

import (
	"github.com/rprajapati0067/golang-interview-test-main/api/domain"
	"github.com/rprajapati0067/golang-interview-test-main/api/errors"
	"github.com/rprajapati0067/golang-interview-test-main/api/repository"
)

var (
	SubscriberService subServiceInterface
)

func init() {
	SubscriberService = &subService{}
}

type subService struct {
}
type subServiceInterface interface {
	CreateSubscriber(domain.Subscriber) (*domain.Subscriber, *errors.RestErr)
}

func (s *subService) CreateSubscriber(input domain.Subscriber) (*domain.Subscriber, *errors.RestErr) {
	id := 0
	var savedData map[int]domain.Subscriber

	if input.Name() != "" && input.Url() != "" {
		id++
		input.SetId(id)
		savedData = repository.InMemoryDb(input, id)
	}
	val := input.Id()
	actualData := savedData[val]
	return &actualData, nil

}
