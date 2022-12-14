// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/Javanshir-SH/query-monitoring/internal/query/service"
	"github.com/Javanshir-SH/query-monitoring/internal/query/storage"
	"sync"
)

// Ensure, that MockQueryStatementsRepository does implement service.QueryStatementsRepository.
// If this is not the case, regenerate this file with moq.
var _ service.QueryStatementsRepository = &MockQueryStatementsRepository{}

// MockQueryStatementsRepository is a mock implementation of service.QueryStatementsRepository.
//
// 	func TestSomethingThatUsesQueryStatementsRepository(t *testing.T) {
//
// 		// make and configure a mocked service.QueryStatementsRepository
// 		mockedQueryStatementsRepository := &MockQueryStatementsRepository{
// 			ListFunc: func(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error) {
// 				panic("mock out the List method")
// 			},
// 		}
//
// 		// use mockedQueryStatementsRepository in code that requires service.QueryStatementsRepository
// 		// and then make assertions.
//
// 	}
type MockQueryStatementsRepository struct {
	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error)

	// calls tracks calls to the methods.
	calls struct {
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter storage.FilterOption
		}
	}
	lockList sync.RWMutex
}

// List calls ListFunc.
func (mock *MockQueryStatementsRepository) List(ctx context.Context, filter storage.FilterOption) (storage.ListOfQuery, error) {
	if mock.ListFunc == nil {
		panic("MockQueryStatementsRepository.ListFunc: method is nil but QueryStatementsRepository.List was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter storage.FilterOption
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx, filter)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedQueryStatementsRepository.ListCalls())
func (mock *MockQueryStatementsRepository) ListCalls() []struct {
	Ctx    context.Context
	Filter storage.FilterOption
} {
	var calls []struct {
		Ctx    context.Context
		Filter storage.FilterOption
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}
