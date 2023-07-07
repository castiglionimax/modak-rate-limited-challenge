package service

import (
	"errors"
	"fmt"
	"time"

	"modak-rated-limited-challenge/internal/domain"
	pkgErr "modak-rated-limited-challenge/pkg/error"
)

type (
	Repository interface {
		GetRule(domain.GroupName) (*domain.Rule, error)
		GetLatestNotification(domain.GroupName, uint) ([]domain.Notification, error)
		AddNotification(domain.GroupName, domain.Notification, uint) error
		SetRule(domain.GroupName, domain.Rule) error
	}
	Service struct {
		repository Repository
	}
)

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s Service) SendNotification(group domain.GroupName, userID, msg string) error {
	rule, err := s.repository.GetRule(group)
	if err != nil {
		return err
	}
	notification := domain.Notification{
		UserID:      userID,
		CreatedDate: time.Now().UTC(),
	}
	notifications, err := s.repository.GetLatestNotification(group, rule.Qty)
	if err != nil {
		if errors.As(err, &pkgErr.NotFoundError{}) {
			if err = s.repository.AddNotification(group, notification, rule.Qty); err != nil {
				return err
			}
			return s.sendNotification(notification.UserID)
		}
		return err
	}

	limiter := calculation(notification.CreatedDate, rule.RangeTime, rule.Qty)

	if len(notifications) >= int(rule.Qty) {
		if limiter.Before(notifications[0].CreatedDate) {
			return pkgErr.ErrRateLimitedReach
		}
	}

	if err = s.repository.AddNotification(group, notification, rule.Qty); err != nil {
		return err
	}

	if err = s.sendNotification(notification.UserID); err != nil {
		return err
	}
	return nil
}

func calculation(notificationTime time.Time, rangeTime time.Duration, qty uint) time.Time {
	return notificationTime.Add(-rangeTime * time.Duration(qty))
}

func (s Service) SetRule(group domain.GroupName, rule domain.Rule) error {
	return s.repository.SetRule(group, rule)
}

func (s Service) sendNotification(user string) error {
	fmt.Printf("sending message to user %s, %s \n", user, time.Now().String())
	return nil
}
