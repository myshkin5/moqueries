// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockNoResultsFn holds the state of a mock of the NoResultsFn type
type mockNoResultsFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams map[mockNoResultsFn_paramsKey]*mockNoResultsFn_resultMgr
}

// mockNoResultsFn_mock isolates the mock interface of the NoResultsFn type
type mockNoResultsFn_mock struct {
	mock *mockNoResultsFn
}

// mockNoResultsFn_recorder isolates the recorder interface of the NoResultsFn type
type mockNoResultsFn_recorder struct {
	mock *mockNoResultsFn
}

// mockNoResultsFn_params holds the params of the NoResultsFn type
type mockNoResultsFn_params struct {
	sParam string
	bParam bool
}

// mockNoResultsFn_paramsKey holds the map key params of the NoResultsFn type
type mockNoResultsFn_paramsKey struct {
	sParam string
	bParam bool
}

// mockNoResultsFn_resultMgr manages multiple results and the state of the NoResultsFn type
type mockNoResultsFn_resultMgr struct {
	params   mockNoResultsFn_params
	results  []*mockNoResultsFn_results
	index    uint32
	anyTimes bool
}

// mockNoResultsFn_results holds the results of the NoResultsFn type
type mockNoResultsFn_results struct {
}

// mockNoResultsFn_fnRecorder routes recorded function calls to the mockNoResultsFn mock
type mockNoResultsFn_fnRecorder struct {
	params    mockNoResultsFn_params
	paramsKey mockNoResultsFn_paramsKey
	results   *mockNoResultsFn_resultMgr
	mock      *mockNoResultsFn
}

// newMockNoResultsFn creates a new mock of the NoResultsFn type
func newMockNoResultsFn(scene *moq.Scene, config *moq.MockConfig) *mockNoResultsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockNoResultsFn{
		scene:  scene,
		config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the NoResultsFn type
func (m *mockNoResultsFn) mock() testmocks.NoResultsFn {
	return func(sParam string, bParam bool) { mock := &mockNoResultsFn_mock{mock: m}; mock.fn(sParam, bParam) }
}

func (m *mockNoResultsFn_mock) fn(sParam string, bParam bool) {
	params := mockNoResultsFn_params{
		sParam: sParam,
		bParam: bParam,
	}
	paramsKey := mockNoResultsFn_paramsKey{
		sParam: sParam,
		bParam: bParam,
	}
	results, ok := m.mock.resultsByParams[paramsKey]
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
	return
}

func (m *mockNoResultsFn) onCall(sParam string, bParam bool) *mockNoResultsFn_fnRecorder {
	return &mockNoResultsFn_fnRecorder{
		params: mockNoResultsFn_params{
			sParam: sParam,
			bParam: bParam,
		},
		paramsKey: mockNoResultsFn_paramsKey{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m,
	}
}

func (r *mockNoResultsFn_fnRecorder) returnResults() *mockNoResultsFn_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams[r.paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockNoResultsFn_resultMgr{
			params:   r.params,
			results:  []*mockNoResultsFn_results{},
			index:    0,
			anyTimes: false,
		}
		r.mock.resultsByParams[r.paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockNoResultsFn_results{})
	return r
}

func (r *mockNoResultsFn_fnRecorder) times(count int) *mockNoResultsFn_fnRecorder {
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

func (r *mockNoResultsFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockNoResultsFn) Reset() {
	m.resultsByParams = map[mockNoResultsFn_paramsKey]*mockNoResultsFn_resultMgr{}
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockNoResultsFn) AssertExpectationsMet() {
	for _, results := range m.resultsByParams {
		missing := len(results.results) - int(atomic.LoadUint32(&results.index))
		if missing == 1 && results.anyTimes == true {
			continue
		}
		if missing > 0 {
			m.scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
		}
	}
}
