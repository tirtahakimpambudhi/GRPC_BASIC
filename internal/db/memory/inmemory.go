package memory

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"grpc_course/internal/domain/model"
	"grpc_course/pb"
	"grpc_course/pkg"
	"sync"
)

var ErrAlreadyExist = errors.New("the laptop already exist")
var ErrNotExist = errors.New("the laptop not exist")

type InMemoryStore struct {
	mutex sync.RWMutex
	db    map[string]*pb.Laptop
}

func NewInMemoryStore() model.InMemoryStore {
	return &InMemoryStore{db: make(map[string]*pb.Laptop)}
}

func (i *InMemoryStore) Search(filter *pb.Filter, found func(laptops *pb.Laptop) error) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	for _, laptop := range i.db {
		if pkg.IsQualified(filter, laptop) {
			other := &pb.Laptop{}
			err := copier.Copy(other, laptop)
			if err != nil {
				return err
			}
			errCallback := found(other)
			if errCallback != nil {
				return errCallback
			}
		}
	}
	return nil
}

func (i *InMemoryStore) Save(laptop *pb.Laptop) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.db[laptop.Id] != nil {
		return ErrAlreadyExist
	}
	other := &pb.Laptop{}
	err := copier.Copy(other, laptop)
	if err != nil {
		return fmt.Errorf("cannot save laptop with id (%s) cuz : %v \n", laptop.Id, err.Error())
	}
	i.db[other.Id] = other
	return nil
}

func (i *InMemoryStore) Find(id string) (error, *pb.Laptop) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.db[id] == nil {
		return ErrNotExist, nil
	}
	return nil, i.db[id]
}

func (i *InMemoryStore) FindAll() (error, []*pb.Laptop) {
	i.mutex.Lock()
	defer i.mutex.Unlock()

	if len(i.db) == 0 {
		return nil, nil
	}
	var result []*pb.Laptop
	for _, laptop := range i.db {
		result = append(result, laptop)
	}
	return nil, result
}

func (i *InMemoryStore) Delete(id string) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	if i.db[id] == nil {
		return ErrNotExist
	}
	delete(i.db, id)
	return nil
}
