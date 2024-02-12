package service

import (
	"getir-case/storage"
	"getir-case/utils"
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
	//var requestBody mongoHandlerRequestBody
	//err := json.NewDecoder(r.Body).Decode(&requestBody)
	//if err != nil {
	//	utils.respondWithError(w, http.StatusBadRequest, "Invalid request body")
	//	return
	//}
	//
	//records, err := s.dbStorage.GetRecords(requestBody.StartDate, requestBody.EndDate, requestBody.MinCount, requestBody.MaxCount)
	//if err != nil {
	//	utils.respondWithError(w, http.StatusInternalServerError, "Error while fetching records")
	//	return
	//}
	//
	//utils.respondWithJSON(w, http.StatusOK, records)
	utils.RespondWithJSON(w, http.StatusOK, "records")
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
