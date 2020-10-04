// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/config"
	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/testing"
)

// mockUsualFn holds the state of a mock of the UsualFn type
type mockUsualFn struct {
	t               testing.MoqT
	config          config.MockConfig
	resultsByParams map[mockUsualFn_params]*mockUsualFn_resultMgr
	params          chan mockUsualFn_params
}

// mockUsualFn_mock isolates the mock interface of the UsualFn type
type mockUsualFn_mock struct {
	mock *mockUsualFn
}

// mockUsualFn_recorder isolates the recorder interface of the UsualFn type
type mockUsualFn_recorder struct {
	mock *mockUsualFn
}

// mockUsualFn_params holds the params of the UsualFn type
type mockUsualFn_params struct {
	sParam string
	bParam bool
}

// mockUsualFn_resultMgr manages multiple results and the state of the UsualFn type
type mockUsualFn_resultMgr struct {
	results  []*mockUsualFn_results
	index    uint32
	anyTimes bool
}

// mockUsualFn_results holds the results of the UsualFn type
type mockUsualFn_results struct {
	sResult string
	err     error
}

// mockUsualFn_fnRecorder routes recorded function calls to the mockUsualFn mock
type mockUsualFn_fnRecorder struct {
	params  mockUsualFn_params
	results *mockUsualFn_resultMgr
	mock    *mockUsualFn
}

// newMockUsualFn creates a new mock of the UsualFn type
func newMockUsualFn(t testing.MoqT, c *config.MockConfig) *mockUsualFn {
	if c == nil {
		c = &config.MockConfig{}
	}
	return &mockUsualFn{
		t:               t,
		config:          *c,
		resultsByParams: map[mockUsualFn_params]*mockUsualFn_resultMgr{},
		params:          make(chan mockUsualFn_params, 100),
	}
}

// mock returns the mock implementation of the UsualFn type
func (m *mockUsualFn) mock() testmocks.UsualFn {
	return func(sParam string, bParam bool) (sResult string, err error) {
		mock := &mockUsualFn_mock{mock: m}
		return mock.fn(sParam, bParam)
	}
}

func (m *mockUsualFn_mock) fn(sParam string, bParam bool) (sResult string, err error) {
	params := mockUsualFn_params{
		sParam: sParam,
		bParam: bParam,
	}
	m.mock.params <- params
	results, ok := m.mock.resultsByParams[params]
	if !ok {
		if m.mock.config.Expectation == config.Strict {
			m.mock.t.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.index, 1)) - 1
	if i >= len(results.results) {
		if !results.anyTimes {
			if m.mock.config.Expectation == config.Strict {
				m.mock.t.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.results) - 1
	}
	result := results.results[i]
	sResult = result.sResult
	err = result.err
	return
}

func (m *mockUsualFn) onCall(sParam string, bParam bool) *mockUsualFn_fnRecorder {
	return &mockUsualFn_fnRecorder{
		params: mockUsualFn_params{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m,
	}
}

func (r *mockUsualFn_fnRecorder) returnResults(sResult string, err error) *mockUsualFn_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams[r.params]; ok {
			r.mock.t.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsualFn_resultMgr{results: []*mockUsualFn_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockUsualFn_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockUsualFn_fnRecorder) times(count int) *mockUsualFn_fnRecorder {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockUsualFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.t.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}
