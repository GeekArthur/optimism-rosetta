// Code generated by mockery v2.28.2. DO NOT EDIT.

package optimism

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/coinbase/rosetta-sdk-go/types"
)

// CurrencyFetcher is an autogenerated mock type for the CurrencyFetcher type
type CurrencyFetcher struct {
	mock.Mock
}

// FetchCurrency provides a mock function with given fields: ctx, blockNum, contractAddress
func (_m *CurrencyFetcher) FetchCurrency(ctx context.Context, blockNum uint64, contractAddress string) (*types.Currency, error) {
	ret := _m.Called(ctx, blockNum, contractAddress)

	var r0 *types.Currency
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) (*types.Currency, error)); ok {
		return rf(ctx, blockNum, contractAddress)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64, string) *types.Currency); ok {
		r0 = rf(ctx, blockNum, contractAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Currency)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64, string) error); ok {
		r1 = rf(ctx, blockNum, contractAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCurrencyFetcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewCurrencyFetcher creates a new instance of CurrencyFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCurrencyFetcher(t mockConstructorTestingTNewCurrencyFetcher) *CurrencyFetcher {
	mock := &CurrencyFetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
