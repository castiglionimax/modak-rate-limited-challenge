package service_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"modak-rated-limited-challenge/internal/domain"
	"modak-rated-limited-challenge/internal/service"
)

type RepositoryMock struct {
	mock.Mock
}

func (s *RepositoryMock) GetRule(group domain.GroupName) (*domain.Rule, error) {
	args := s.Called(group)
	return args.Get(0).(*domain.Rule), args.Error(1)
}

func (s *RepositoryMock) GetLatestNotification(group domain.GroupName, qty uint) ([]domain.Notification, error) {
	args := s.Called(group)
	return args.Get(0).([]domain.Notification), args.Error(1)
}

func (s *RepositoryMock) AddNotification(group domain.GroupName, notification domain.Notification, qty uint) error {
	args := s.Called(group)
	return args.Error(0)
}

func (s *RepositoryMock) SetRule(group domain.GroupName, rule domain.Rule) error {
	args := s.Called(group)
	return args.Error(0)
}

func TestSendNotification(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		clientMock := new(RepositoryMock)
		underTest := service.NewService(clientMock)

		clientMock.On("GetRule", mock.Anything).Return(&domain.Rule{
			Qty:       2,
			RangeTime: time.Minute,
		}, nil)
		clientMock.On("GetLatestNotification", mock.Anything).Return([]domain.Notification{{
			UserID:      "user",
			CreatedDate: time.Now().Add(-time.Minute*2 - time.Second),
		}, {
			UserID:      "user",
			CreatedDate: time.Now().Add(-time.Minute),
		}}, nil)
		clientMock.On("AddNotification", mock.Anything).Return(nil)

		err := underTest.SendNotification("Status", "user", "hello world")
		assert.NoError(t, err)
	})
	t.Run("err when CreateAuthorizationsBulk fails", func(t *testing.T) {
		clientMock := new(RepositoryMock)
		underTest := service.NewService(clientMock)

		clientMock.On("GetRule", mock.Anything).Return(&domain.Rule{
			Qty:       2,
			RangeTime: time.Minute,
		}, nil)
		clientMock.On("GetLatestNotification", mock.Anything).Return([]domain.Notification{{
			UserID:      "user",
			CreatedDate: time.Now().Add(-time.Second),
		}, {
			UserID:      "user",
			CreatedDate: time.Now().Add(-time.Second * 20),
		}}, nil)
		clientMock.On("AddNotification", mock.Anything).Return(nil)
		err := underTest.SendNotification("Status", "user", "hello world")
		assert.Error(t, err)
	})
}
