package mocks

import (
	"errors"
	"github.com/lilakurse/learning-golang/mock"
)

func NewMockDataStore() *DataStore {
	mockObj := new(DataStore)
	// Tests for CreateUser()
	mockObj.On("CreateUser", mock.GoodUser).Return(nil)
	mockObj.On("CreateUser", mock.EmptyUser).Return(errors.New("short name + invalid contacts"))
	mockObj.On("CreateUser", mock.BadUser).Return(errors.New("bad email"))

	// Tests for GetUserByID()
	mockObj.On("GetUserByID", mock.EmptyUser.ID.Hex()).Return(nil, errors.New("empty ID"))
	mockObj.On("GetUserByID", mock.BadUserID).Return(nil, errors.New("nonexistent ID"))
	mockObj.On("GetUserByID", mock.GoodUser).Return(mock.GoodUser, nil)

	return mockObj
}
