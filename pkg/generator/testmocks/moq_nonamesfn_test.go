// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockNoNamesFn holds the state of a mock of the NoNamesFn type
type mockNoNamesFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams map[mockNoNamesFn_params]*mockNoNamesFn_resultMgr
}

// mockNoNamesFn_mock isolates the mock interface of the NoNamesFn type
type mockNoNamesFn_mock struct {
	mock *mockNoNamesFn
}

// mockNoNamesFn_recorder isolates the recorder interface of the NoNamesFn type
type mockNoNamesFn_recorder struct {
	mock *mockNoNamesFn
}

// mockNoNamesFn_params holds the params of the NoNamesFn type
type mockNoNamesFn_params struct {
	sParam string
	bParam bool
}

// mockNoNamesFn_resultMgr manages multiple results and the state of the NoNamesFn type
type mockNoNamesFn_resultMgr struct {
	results  []*mockNoNamesFn_results
	index    uint32
	anyTimes bool
}

// mockNoNamesFn_results holds the results of the NoNamesFn type
type mockNoNamesFn_results struct {
	sResult string
	err     error
}

// mockNoNamesFn_fnRecorder routes recorded function calls to the mockNoNamesFn mock
type mockNoNamesFn_fnRecorder struct {
	params  mockNoNamesFn_params
	results *mockNoNamesFn_resultMgr
	mock    *mockNoNamesFn
}

// newMockNoNamesFn creates a new mock of the NoNamesFn type
func newMockNoNamesFn(scene *moq.Scene, config *moq.MockConfig) *mockNoNamesFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockNoNamesFn{
		scene:  scene,
		config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the NoNamesFn type
func (m *mockNoNamesFn) mock() testmocks.NoNamesFn {
	return func(sParam string, bParam bool) (sResult string, err error) {
		mock := &mockNoNamesFn_mock{mock: m}
		return mock.fn(sParam, bParam)
	}
}

func (m *mockNoNamesFn_mock) fn(sParam string, bParam bool) (sResult string, err error) {
	params := mockNoNamesFn_params{
		sParam: sParam,
		bParam: bParam,
	}
	results, ok := m.mock.resultsByParams[params]
	if !ok {
		if m.mock.config.Expectation == moq.Strict {
			m.mock.scene.MoqT.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.index, 1)) - 1
	if i >= len(results.results) {
		if !results.anyTimes {
			if m.mock.config.Expectation == moq.Strict {
				m.mock.scene.MoqT.Fatalf("Too many calls to mock with parameters %#v", params)
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

func (m *mockNoNamesFn) onCall(sParam string, bParam bool) *mockNoNamesFn_fnRecorder {
	return &mockNoNamesFn_fnRecorder{
		params: mockNoNamesFn_params{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m,
	}
}

func (r *mockNoNamesFn_fnRecorder) returnResults(sResult string, err error) *mockNoNamesFn_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams[r.params]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockNoNamesFn_resultMgr{results: []*mockNoNamesFn_results{}, index: 0, anyTimes: false}
		r.mock.resultsByParams[r.params] = r.results
	}
	r.results.results = append(r.results.results, &mockNoNamesFn_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockNoNamesFn_fnRecorder) times(count int) *mockNoNamesFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockNoNamesFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockNoNamesFn) Reset() {
	m.resultsByParams = map[mockNoNamesFn_params]*mockNoNamesFn_resultMgr{}
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockNoNamesFn) AssertExpectationsMet() {
	for params, results := range m.resultsByParams {
		missing := len(results.results) - int(atomic.LoadUint32(&results.index))
		if missing == 1 && results.anyTimes == true {
			continue
		}
		if missing > 0 {
			m.scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, params)
		}
	}
}
