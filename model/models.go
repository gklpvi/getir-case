package model

import "time"


type DBRecord struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

type IMRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DBHandlerRequestBody struct {
	StartDate string `json:"startDate" validate:"required"`
	EndDate   string `json:"endDate" validate:"required"`
	MinCount  int64  `json:"minCount"`
	MaxCount  int64  `json:"maxCount"`
}

type IMHandlerRequestBody struct {
	Key   string `json:"key" validate:"required" omitempty:"true"`
	Value string `json:"value" validate:"required" omitempty:"true"`
}
