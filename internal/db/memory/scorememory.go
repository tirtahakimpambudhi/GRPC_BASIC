package memory

import (
	"grpc_course/internal/domain/model"
	"sync"
)

type ScoreStore struct {
	mutex sync.RWMutex
	db    map[string]*model.Ratings
}

func NewScoreStore() model.ScoreStore {
	return &ScoreStore{db: make(map[string]*model.Ratings)}
}

func (s *ScoreStore) Add(laptopId string, score float64) (*model.Ratings, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	rating := s.db[laptopId]
	if rating == nil {
		rating = &model.Ratings{
			Count: 1,
			Sum:   score,
		}
	} else {
		rating.Count++
		rating.Sum += score
	}
	s.db[laptopId] = rating
	return rating, nil
}
