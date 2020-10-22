// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockUsualFn holds the state of a mock of the UsualFn type
type mockUsualFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams map[mockUsualFn_paramsKey]*mockUsualFn_resultMgr
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

// mockUsualFn_paramsKey holds the map key params of the UsualFn type
type mockUsualFn_paramsKey struct {
	sParam string
	bParam bool
}

// mockUsualFn_resultMgr manages multiple results and the state of the UsualFn type
type mockUsualFn_resultMgr struct {
	params   mockUsualFn_params
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
	params    mockUsualFn_params
	paramsKey mockUsualFn_paramsKey
	results   *mockUsualFn_resultMgr
	mock      *mockUsualFn
}

// newMockUsualFn creates a new mock of the UsualFn type
func newMockUsualFn(scene *moq.Scene, config *moq.MockConfig) *mockUsualFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockUsualFn{
		scene:  scene,
		config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
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
	paramsKey := mockUsualFn_paramsKey{
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
		paramsKey: mockUsualFn_paramsKey{
			sParam: sParam,
			bParam: bParam,
		},
		mock: m,
	}
}

func (r *mockUsualFn_fnRecorder) returnResults(sResult string, err error) *mockUsualFn_fnRecorder {
	if r.results == nil {
		if _, ok := r.mock.resultsByParams[r.paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockUsualFn_resultMgr{
			params:   r.params,
			results:  []*mockUsualFn_results{},
			index:    0,
			anyTimes: false,
		}
		r.mock.resultsByParams[r.paramsKey] = r.results
	}
	r.results.results = append(r.results.results, &mockUsualFn_results{
		sResult: sResult,
		err:     err,
	})
	return r
}

func (r *mockUsualFn_fnRecorder) times(count int) *mockUsualFn_fnRecorder {
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

func (r *mockUsualFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockUsualFn) Reset() { m.resultsByParams = map[mockUsualFn_paramsKey]*mockUsualFn_resultMgr{} }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockUsualFn) AssertExpectationsMet() {
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
