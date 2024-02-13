# Getir Case

* This is a simple REST API that provides a couple endpoints, one to fetch records from a MongoDB collection, and to insert and read from in-memory db.
* The API is written in Go, and is containerized using Docker and uses docker-compose to run the application and the MongoDB instance (if needed).


### Prerequisites
* Docker
* Docker Compose
* Go 1.16 or later
* MongoDB

### Installation

Clone the repository

```
git clone github.com/gklpvi/getir-case
```

### Running
```
cd getir-case
go test ./... -coverprofile=coverage.out 
go tool cover -html=coverage.out

make dev-run
```

### Usage

* To fetch records from the MongoDB collection, make a GET request to `http://localhost:8080/records` with the following query parameters:
  * `startDate` (required) - The start date of the record
  * `endDate` (required) - The end date of the record
  * `minCount` - The minimum count of the record
  * `maxCount` - The maximum count of the record

* To insert and read from in-memory db, make a POST request to `http://localhost:8080/in-memory` with the following request body:
* 
```
{
    "key": "key1",
    "value": "value1"
}
```

* To fetch records from the in-memory db, make a GET request to `http://localhost:8080/in-memory` with the following query parameters:
  * `key` (required) - The key of the record

