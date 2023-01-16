package repository

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	SubscriptionPlanID string
	Amount             float64
	SentAt             time.Time
	PaidAt             time.Time
	AccountGroupID     int
	AccountGroup       AccountGroup
}

type Subscription struct {
	gorm.Model
	ID             string
	AccountGroupID int
	AccountGroup   AccountGroup
	ValidFrom      time.Time
	ValidTill      time.Time
	Active         bool
	AutoRenewed    bool
}

type PlanFeature struct {
	gorm.Model
	PlanID   string
	Features []Feature `gorm:"many2many:planfeatures_feature;"`
}

type SubscriptionPlan struct {
	gorm.Model
	ID             string
	SubscriptionID string
	PlanID         string
	ValidFrom      time.Time
	ValidTill      time.Time
	PeriodType     string
	ParentID       string
	IsPaid         bool
	IsRenewable    bool
}

type Feature struct {
	gorm.Model
	ClientCount int
	Cost        float64
}

type Plan struct {
	gorm.Model
	ID           string
	Type         string
	ValidFrom    time.Time
	ValidTill    time.Time
	ParentID     string
	CostPerDay   float64
	CostPerMonth float64
	CostPerYear  float64
	Country      string
	Currency     string
}
