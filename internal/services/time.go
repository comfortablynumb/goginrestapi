package services

import "time"

// Interfaces

type TimeService interface {
	GetCurrentUtcTime() time.Time
}

// Structs

type timeService struct {
}

func (s *timeService) GetCurrentUtcTime() time.Time {
	return time.Now().UTC()
}

// Static functions

func NewTimeService() TimeService {
	return &timeService{}
}
