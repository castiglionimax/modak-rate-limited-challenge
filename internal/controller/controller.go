package controller

import (
	"errors"
	"modak-rated-limited-challenge/internal/domain"
	"strings"
	"time"
)

type (
	Service interface {
		SendNotification(group domain.GroupName, userID, msg string) error
		SetRule(group domain.GroupName, rule domain.Rule) error
	}
	Controller struct {
		service Service
	}
)

func NewController(service Service) Controller {
	return Controller{service: service}
}
func (c Controller) SendNotification(group, userID, msg string) error {
	return c.service.SendNotification(domain.GroupName(strings.ToLower(group)), userID, msg)
}

func (c Controller) SetRule(group, rangeTime string, qty uint) error {
	if qty == 0 {
		return errors.New("qty must be mayor than 0")
	}
	rangeTimeDuration, err := time.ParseDuration(rangeTime)
	if err != nil {
		return err
	}
	return c.service.SetRule(domain.GroupName(strings.ToLower(group)), domain.Rule{
		Qty:       qty,
		RangeTime: rangeTimeDuration,
	})
}
