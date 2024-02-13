package service

import (
	"encoding/json"
	"getir-case/model"
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
	if r.Method == http.MethodPost {
		// we use pointer to not mistake us on missing fields in request body by filling with default values
		// thus use declared tags in model
		var requestBody *model.DBHandlerRequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{
				"code":    "1",
				"msg":     "Invalid request body, err: " + err.Error(),
				"records": "",
			})
			return
		}

		startDate, err := utils.ParseDate(requestBody.StartDate)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{
				"code":    "1",
				"msg":     "Invalid start date, err: " + err.Error(),
				"records": "",
			})
			return
		}
		endDate, err := utils.ParseDate(requestBody.EndDate)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, map[string]string{
				"code":    "1",
				"msg":     "Invalid end date, err: " + err.Error(),
				"records": "",
			})
			return
		}

		records, err := s.dbStorage.GetRecords(startDate, endDate, requestBody.MinCount, requestBody.MaxCount)
		if err != nil {
			utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
				"code":    "1",
				"msg":     "Error while fetching records, err: " + err.Error(),
				"records": "",
			})
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
			"code":    0,
			"msg":     "Success",
			"records": records,
		})
	} else {
		utils.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"code":    "2",
			"msg":     "Method not allowed",
			"records": "",
		})
	}
}

func (s *Service) GetAllRecordsFromDB(w http.ResponseWriter, r *http.Request) {
	records, err := s.dbStorage.GetAllRecords()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error while fetching records, err: "+err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"code":    0,
		"msg":     "Success",
		"records": records,
	})
}

func (s *Service) InMemoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.addRecordToIM(w, r)
	} else if r.Method == http.MethodGet {
		s.getRecordFromIM(w, r)
	} else {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (s *Service) addRecordToIM(w http.ResponseWriter, r *http.Request) {
	var requestBody *model.IMHandlerRequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	s.inMemoryStorage.AddItem(
		&model.IMRecord{
			Key:   requestBody.Key,
			Value: requestBody.Value,
		})

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{
		"key": requestBody.Key,
		"msg": "Record added successfully",
	})
}

func (s *Service) getRecordFromIM(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Key is required")
		return
	}

	record, ok := s.inMemoryStorage.GetItem(key)
	if !ok {
		utils.RespondWithError(w, http.StatusNotFound, "Record not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, record)
}

func (s *Service) GetAllRecordsFromIM(w http.ResponseWriter, r *http.Request) {
	records := s.inMemoryStorage.GetAllItems()

	utils.RespondWithJSON(w, http.StatusOK, records)
}
