package test

import (
	"bytes"
	"encoding/json"
	"getir-case/model"
	"getir-case/service"
	"getir-case/storage"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetRecordsFromDB(t *testing.T) { //load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	// create a new service
	dbStorage, err := storage.NewMongoDBStorage(dbUri, dbName, collectionName)
	inMemoryStorage := storage.NewInMemoryStorage()
	s := service.NewService(dbStorage, inMemoryStorage)

	// create a new request
	requestBody := model.DBHandlerRequestBody{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  3000,
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/records", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// create a new response recorder
	rr := httptest.NewRecorder()

	// call GetRecordsFromDB handler
	handler := http.HandlerFunc(s.GetRecordsFromDB)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body
	expected := `{"code":"0","msg":"Success","records":[{"key":"TAKwGc6Jr4i8Z487","createdAt":"2017-01-28T01:22:14.398Z","totalCount":2800},{"key":"NAeQ8eX7e5TEg7oH","createdAt":"2017-01-27T08:19:14.135Z","totalCount":2900}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetRecordsFromInMemory(t *testing.T) {
	// create a new service
	dbStorage, err := storage.NewMongoDBStorage("mongodb://localhost:27017", "getir-case", "records")
	inMemoryStorage := storage.NewInMemoryStorage()
	s := service.NewService(dbStorage, inMemoryStorage)

	// create a new request
	req, err := http.NewRequest("GET", "/in-memory-records", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new response recorder
	rr := httptest.NewRecorder()

	// call GetRecordsFromInMemory handler
	handler := http.HandlerFunc(s.GetAllRecordsFromIM)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body
	expected := `{"code":"0","msg":"Success","records":[]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateRecordInMemory(t *testing.T) {
	// create a new service
	dbStorage, err := storage.NewMongoDBStorage("mongodb://localhost:27017", "getir-case", "records")
	inMemoryStorage := storage.NewInMemoryStorage()
	s := service.NewService(dbStorage, inMemoryStorage)

	// create a new request
	requestBody := model.IMHandlerRequestBody{
		Key:   "active-tabs",
		Value: "getir",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", "/in-memory", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	// create a new response recorder
	rr := httptest.NewRecorder()

	// call CreateRecordInMemory handler
	handler := http.HandlerFunc(s.InMemoryHandler)
	handler.ServeHTTP(rr, req)

	// check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body
	expected := `{"code":"0","msg":"Success","record":{"key":"active-tabs","value":"getir"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
