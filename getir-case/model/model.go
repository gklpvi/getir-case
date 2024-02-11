package model

// {
//"key":"TAKwGc6Jr4i8Z487",
//"createdAt":"2017-01-28T01:22:14.398Z",
//"totalCount":2800
//},

type DBRecord struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}

//{
//"key": "active-tabs",
//"value": "getir"
//}

type InMemoryRecord struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
