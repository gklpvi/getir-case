package service

import (
	"encoding/json"
	"getir-case/model"
	"getir-case/storage"
	"getir-case/utils"
	"log"
	"net/http"
)

type Service struct {
	dbStorage       *storage.MongoDBStorage
	inMemoryStorage *storage.InMemoryStorage
}

func NewService(dbStorage *storage.MongoDBStorage, inMemoryStorage *storage.InMemoryStorage) *Service {
	return &Service{
		dbStorage:       dbStorage,
		inMemoryStorage: inMemoryStorage,
	}
}

func (s *Service) GetRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	var requestBody model.DBHandlerRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Default().Println("Error while decoding request body, value: ", requestBody)
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body, err: "+err.Error())
		return
	}

	startDate, err := utils.ParseDate(requestBody.StartDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid start date, err: "+err.Error())
		return
	}
	endDate, err := utils.ParseDate(requestBody.EndDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid end date, err: "+err.Error())
		return
	}

	records, err := s.dbStorage.GetRecords(startDate, endDate, requestBody.MinCount, requestBody.MaxCount)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while fetching records, err: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, records)
}

func (s *Service) GetAllRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	records, err := s.dbStorage.GetAllRecords()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while fetching records, err: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, records)
}

func (s *Service) AddRecordToIM(w http.ResponseWriter, r *http.Request) {
	//var requestBody inMemoryHandlerRequestBody
	//err := json.NewDecoder(r.Body).Decode(&requestBody)
	//if err != nil {
	//	utils.respondWithError(w, http.StatusBadRequest, "Invalid request body")
	//	return
	//}
	//
	//err = s.inMemoryStorage.AddItem(requestBody)
	//if err != nil {
	//	utils.respondWithError(w, http.StatusInternalServerError, "Error while adding record")
	//	return
	//}

	utils.RespondWithJSON(w, http.StatusOK, "Record added successfully")
}

func (s *Service) GetRecordFromIM(w http.ResponseWriter, r *http.Request) {
	//records, err := s.inMemoryStorage.GetAllItems()
	//if err != nil {
	//	utils.respondWithError(w, http.StatusInternalServerError, "Error while fetching records")
	//	return
	//}

	utils.RespondWithJSON(w, http.StatusOK, "testtest")
}
