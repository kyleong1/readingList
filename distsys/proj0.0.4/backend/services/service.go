/* 
	Author: Kyle Ong
	Date: 10/25/2018
	
	backend service for reading list application
*/
package service

import (
	// standard
	"fmt"
	"sync" // user-defined

	"distsys/proj0.0.4/backend/datamodels"
)

type Service interface {
	Get(owner string) []datamodels.Item
	Save(owner string, newItems []datamodels.Item) error
}

type DataService struct {
	items map[string][]datamodels.Item
	mu    sync.RWMutex
}

func NewDataService(source map[string][]datamodels.Item) *DataService {
	return &DataService{
		items: source}
}

func (s *DataService) Get(sessionOwner string) []datamodels.Item {
	s.mu.RLock()
	if _, exists := s.items["test"]; exists {
		testTasks := s.items["test"]
		delete(s.items, "test")
		for i := 0; i < len(testTasks); i++ {
			testTasks[i].SessionID = sessionOwner
			s.items[sessionOwner] = append(s.items[sessionOwner], testTasks[i])
		}
	}
	items := s.items[sessionOwner]
	fmt.Println(items)
	s.mu.RUnlock()
	return items
}

func (s *DataService) Save(sessionOwner string, newItems []datamodels.Item) error {
	var prevID int64
	for i := range newItems {
		if newItems[i].ID == 0 {
			newItems[i].ID = prevID
			prevID++
		}
	}

	s.mu.Lock()
	s.items[sessionOwner] = newItems
	s.mu.Unlock()
	fmt.Println(s.items)
	return nil
}
