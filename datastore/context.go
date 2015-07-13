package datastore

import (
	"github.com/lilakurse/learning-golang/datastore/mocks"
	"golang.org/x/net/context"
)

type Context string

const (
	ContextDataStore = Context("DATASTORE")
)

//создание контекста
var (
	ctx = context.WithValue(context.Background(), ContextDataStore, mocks.NewMockDataStore())
)

func FromContext(ctx context.Context) DataStore {
	return ctx.Value(ContextDataStore).(DataStore)
}
