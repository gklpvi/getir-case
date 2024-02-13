package model

import "time"

// {
//"key":"TAKwGc6Jr4i8Z487",
//"createdAt":"2017-01-28T01:22:14.398Z",
//"totalCount":2800
//},

type DBRecord struct {
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalCount int       `json:"totalCount"`
}

//{
//"key": "active-tabs",
//"value": "getir"
//}

type IMRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DBHandlerRequestBody struct {
	StartDate string `json:"startDate" required:"true"`
	EndDate   string `json:"endDate" required:"true"`
	MinCount  int64  `json:"minCount"`
	MaxCount  int64  `json:"maxCount"`
}

type IMHandlerRequestBody struct {
	Key   string `json:"key" required:"true" omitempty:"true"`
	Value string `json:"value" required:"true" omitempty:"true"`
}
