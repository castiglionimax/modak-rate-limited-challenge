package repository

import (
	"container/list"
	"fmt"
	"sync"

	"modak-rated-limited-challenge/internal/domain"
	pkgErr "modak-rated-limited-challenge/pkg/error"
)

type (
	Repository struct {
		rules         map[domain.GroupName]domain.Rule
		notifications map[domain.GroupName]*list.List
		mutex         sync.RWMutex
	}
)

func NewRepository() *Repository {
	return &Repository{
		rules:         make(map[domain.GroupName]domain.Rule),
		notifications: make(map[domain.GroupName]*list.List),
	}
}

func (r *Repository) GetRule(group domain.GroupName) (*domain.Rule, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	mm, ok := r.rules[group]
	if !ok {
		return nil, pkgErr.NotFoundError{Cause: fmt.Errorf("group %s has not found", group)}
	}
	return &mm, nil
}

func (r *Repository) SetRule(group domain.GroupName, rule domain.Rule) error {
	r.rules[group] = rule
	return nil
}

func (r *Repository) GetLatestNotification(group domain.GroupName, qty uint) ([]domain.Notification, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	mm, ok := r.notifications[group]
	if !ok {
		return nil, pkgErr.NotFoundError{Cause: fmt.Errorf("group %s has not found", group)}
	}
	notifications := make([]domain.Notification, 0, qty)
	for element := mm.Front(); element != nil; element = element.Next() {
		notifications = append(notifications, element.Value.(domain.Notification))
	}
	return notifications, nil
}

func (r *Repository) AddNotification(group domain.GroupName, notification domain.Notification, qty uint) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	mm, ok := r.notifications[group]
	if !ok {
		notifications := list.New()
		notifications.PushBack(notification)
		r.notifications[group] = notifications
		return nil
	}
	if mm.Len() >= int(qty) {
		mm.Remove(mm.Front())
	}
	mm.PushBack(notification)
	return nil
}
