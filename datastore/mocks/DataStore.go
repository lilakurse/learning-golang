package mocks

import "github.com/stretchr/testify/mock"

import "github.com/lilakurse/learning-golang/model/user"

type DataStore struct {
	mock.Mock
}

func (m *DataStore) CreateUser(u *user.User) error {
	ret := m.Called(u)

	r0 := ret.Error(0)

	return r0
}
func (m *DataStore) GetUserByID(ID string) (*user.User, error) {
	ret := m.Called(ID)

	var r0 *user.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*user.User)
	}
	r1 := ret.Error(1)

	return r0, r1
}
