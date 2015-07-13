package datastore

import (
	//"github.com/lilakurse/learning-golang/datastore/mocks"
	"github.com/lilakurse/learning-golang/mock"
	//"github.com/stretchr/testify/mock"
	"testing"
)

//var mockStore = DataStore(mocks.NewMockDataStore())

func TestCreateUser(t *testing.T) {
	err := CreateUser(ctx, mock.GoodUser)
	if err != nil {
		t.Error("no error")
	}

	err = CreateUser(ctx, mock.EmptyUser)
	if err == nil {
		t.Error("no errors")
	}
	err = CreateUser(ctx, mock.BadUser)
	if err == nil {
		t.Error("Shoulb be an error")
	}
}
