package service

import (
	"time"

	"getir-case/storage"
	"getir-case/model"
)

type Service struct {
	dbStorage       *storage.MongoDBStorage
	inMemoryStorage *storage.InMemoryStorage
}

func New(dbStorage *storage.MongoDBStorage, inMemoryStorage *storage.InMemoryStorage) *Service {
	return &Service{
		dbStorage:       dbStorage,
		inMemoryStorage: inMemoryStorage,
	}
}

// logic in service layer is simple and may be seen redundant but it is necessary for separation of concerns
// and helpful if application grows

func (s *Service) GetRecordsFromDB(startDate, endDate time.Time, minCount, maxCount int64) ([]model.DBRecord, error) {
	return s.dbStorage.GetRecords(startDate, endDate, minCount, maxCount)
}

func (s *Service) GetAllRecords() ([]model.DBRecord, error) {
	return s.dbStorage.GetAllRecords()
}

func (s *Service) AddRecordToIM(key, value string) error {
	s.inMemoryStorage.AddItem(
		&model.IMRecord{
			Key:   key,
			Value: value,
	})
	return nil
}

func (s *Service) GetRecordFromIM(key string) (model.IMRecord, bool) {
	return s.inMemoryStorage.GetItem(key)
}

func (s *Service) GetAllRecordsFromIM() []model.IMRecord {
	return s.inMemoryStorage.GetAllItems()
}
