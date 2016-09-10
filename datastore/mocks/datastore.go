package mocks

import (
	"timezones_mc/revel_app/app/models"

	"github.com/stretchr/testify/mock"
)

// Datastore is an autogenerated mock type for the Datastore type
type Datastore struct {
	mock.Mock
}

// AddDocument provides a mock function with given fields: doc
func (_m *Datastore) AddDocument(doc models.Document) error {
	ret := _m.Called(doc)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Document) error); ok {
		r0 = rf(doc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDocumentById provides a mock function with given fields: id
func (_m *Datastore) FindDocumentById(id string) (models.Document, error) {
	ret := _m.Called(id)

	var r0 models.Document
	if rf, ok := ret.Get(0).(func(string) models.Document); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestDocuments provides a mock function with given fields: suggesterName, text, field, payloadKey
func (_m *Datastore) SuggestDocuments(suggesterName string, text string, field string, payloadKey string) ([]models.Document, error) {
	ret := _m.Called(suggesterName, text, field, payloadKey)

	var r0 []models.Document
	if rf, ok := ret.Get(0).(func(string, string, string, string) []models.Document); ok {
		r0 = rf(suggesterName, text, field, payloadKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Document)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(suggesterName, text, field, payloadKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}