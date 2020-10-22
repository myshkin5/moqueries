// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// MockRepeatedIdsFn holds the state of a mock of the RepeatedIdsFn type
type MockRepeatedIdsFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams map[MockRepeatedIdsFn_paramsKey]*MockRepeatedIdsFn_resultMgr
}

// MockRepeatedIdsFn_mock isolates the mock interface of the RepeatedIdsFn type
type MockRepeatedIdsFn_mock struct {
	Mock *MockRepeatedIdsFn
}

// MockRepeatedIdsFn_recorder isolates the recorder interface of the RepeatedIdsFn type
type MockRepeatedIdsFn_recorder struct {
	Mock *MockRepeatedIdsFn
}

// MockRepeatedIdsFn_params holds the params of the RepeatedIdsFn type
type MockRepeatedIdsFn_params struct {
	SParam1, SParam2 string
	BParam           bool
}

// MockRepeatedIdsFn_paramsKey holds the map key params of the RepeatedIdsFn type
type MockRepeatedIdsFn_paramsKey struct {
	SParam1, SParam2 string
	BParam           bool
}

// MockRepeatedIdsFn_resultMgr manages multiple results and the state of the RepeatedIdsFn type
type MockRepeatedIdsFn_resultMgr struct {
	Params   MockRepeatedIdsFn_params
	Results  []*MockRepeatedIdsFn_results
	Index    uint32
	AnyTimes bool
}

// MockRepeatedIdsFn_results holds the results of the RepeatedIdsFn type
type MockRepeatedIdsFn_results struct {
	SResult1, SResult2 string
	Err                error
}

// MockRepeatedIdsFn_fnRecorder routes recorded function calls to the MockRepeatedIdsFn mock
type MockRepeatedIdsFn_fnRecorder struct {
	Params    MockRepeatedIdsFn_params
	ParamsKey MockRepeatedIdsFn_paramsKey
	Results   *MockRepeatedIdsFn_resultMgr
	Mock      *MockRepeatedIdsFn
}

// NewMockRepeatedIdsFn creates a new mock of the RepeatedIdsFn type
func NewMockRepeatedIdsFn(scene *moq.Scene, config *moq.MockConfig) *MockRepeatedIdsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockRepeatedIdsFn{
		Scene:  scene,
		Config: *config,
	}
	m.Reset()
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the RepeatedIdsFn type
func (m *MockRepeatedIdsFn) Mock() testmocks.RepeatedIdsFn {
	return func(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
		mock := &MockRepeatedIdsFn_mock{Mock: m}
		return mock.Fn(sParam1, sParam2, bParam)
	}
}

func (m *MockRepeatedIdsFn_mock) Fn(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
	params := MockRepeatedIdsFn_params{
		SParam1: sParam1,
		SParam2: sParam2,
		BParam:  bParam,
	}
	paramsKey := MockRepeatedIdsFn_paramsKey{
		SParam1: sParam1,
		SParam2: sParam2,
		BParam:  bParam,
	}
	results, ok := m.Mock.ResultsByParams[paramsKey]
	if !ok {
		if m.Mock.Config.Expectation == moq.Strict {
			m.Mock.Scene.MoqT.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= len(results.Results) {
		if !results.AnyTimes {
			if m.Mock.Config.Expectation == moq.Strict {
				m.Mock.Scene.MoqT.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.Results) - 1
	}
	result := results.Results[i]
	sResult1 = result.SResult1
	sResult2 = result.SResult2
	err = result.Err
	return
}

func (m *MockRepeatedIdsFn) OnCall(sParam1, sParam2 string, bParam bool) *MockRepeatedIdsFn_fnRecorder {
	return &MockRepeatedIdsFn_fnRecorder{
		Params: MockRepeatedIdsFn_params{
			SParam1: sParam1,
			SParam2: sParam2,
			BParam:  bParam,
		},
		ParamsKey: MockRepeatedIdsFn_paramsKey{
			SParam1: sParam1,
			SParam2: sParam2,
			BParam:  bParam,
		},
		Mock: m,
	}
}

func (r *MockRepeatedIdsFn_fnRecorder) ReturnResults(sResult1, sResult2 string, err error) *MockRepeatedIdsFn_fnRecorder {
	if r.Results == nil {
		if _, ok := r.Mock.ResultsByParams[r.ParamsKey]; ok {
			r.Mock.Scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.Params)
			return nil
		}

		r.Results = &MockRepeatedIdsFn_resultMgr{
			Params:   r.Params,
			Results:  []*MockRepeatedIdsFn_results{},
			Index:    0,
			AnyTimes: false,
		}
		r.Mock.ResultsByParams[r.ParamsKey] = r.Results
	}
	r.Results.Results = append(r.Results.Results, &MockRepeatedIdsFn_results{
		SResult1: sResult1,
		SResult2: sResult2,
		Err:      err,
	})
	return r
}

func (r *MockRepeatedIdsFn_fnRecorder) Times(count int) *MockRepeatedIdsFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < count-1; n++ {
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (r *MockRepeatedIdsFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockRepeatedIdsFn) Reset() {
	m.ResultsByParams = map[MockRepeatedIdsFn_paramsKey]*MockRepeatedIdsFn_resultMgr{}
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockRepeatedIdsFn) AssertExpectationsMet() {
	for _, results := range m.ResultsByParams {
		missing := len(results.Results) - int(atomic.LoadUint32(&results.Index))
		if missing == 1 && results.AnyTimes == true {
			continue
		}
		if missing > 0 {
			m.Scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.Params)
		}
	}
}
