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

type InMemoryRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DBHandlerRequestBody struct {
	// “startDate” and “endDate” fields will contain the date in a “YYYY-MM-DD” format.
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	// “minCount” and “maxCount” are for filtering the data. Sum of the “count” array in the documents should be between “minCount” and “maxCount”.
	MinCount int64 `json:"minCount"`
	MaxCount int64 `json:"maxCount"`
}
