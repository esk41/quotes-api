package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

type QuoteStore struct {
	sync.Mutex
	quotes []Quote
	nextID int
}

func NewStore() *QuoteStore {
	return &QuoteStore{
		quotes: []Quote{},
		nextID: 1,
	}
}

func (s *QuoteStore) AddQuote(q Quote) Quote {
	s.Lock()
	defer s.Unlock()
	q.ID = s.nextID
	q.CreatedAT = time.Now()
	s.nextID++
	s.quotes = append(s.quotes, q)
	return q
}

func (s *QuoteStore) GetAll(author string) []Quote {
	s.Lock()
	defer s.Unlock()
	if author == "" {
		return s.quotes
	}
	filtered := []Quote{}
	for _, q := range s.quotes {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}
	return filtered
}

func (s *QuoteStore) GetRandom() (Quote, error) {
	s.Lock()
	defer s.Unlock()
	if len(s.quotes) == 0 {
		return Quote{}, errors.New("no quotes available")
	}
	return s.quotes[rand.Intn(len(s.quotes))], nil
}

func (s *QuoteStore) DeleteByID(id int) bool {
	s.Lock()
	defer s.Unlock()
	for i, q := range s.quotes {
		if q.ID == id {
			s.quotes = append(s.quotes[:i], s.quotes[i+1:]...)
			return true
		}
	}
	return false
}
