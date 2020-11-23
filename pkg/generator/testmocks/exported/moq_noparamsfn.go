// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// MockNoParamsFn holds the state of a mock of the NoParamsFn type
type MockNoParamsFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams []MockNoParamsFn_resultsByParams
}

// MockNoParamsFn_mock isolates the mock interface of the NoParamsFn type
type MockNoParamsFn_mock struct {
	Mock *MockNoParamsFn
}

// MockNoParamsFn_params holds the params of the NoParamsFn type
type MockNoParamsFn_params struct{}

// MockNoParamsFn_paramsKey holds the map key params of the NoParamsFn type
type MockNoParamsFn_paramsKey struct{}

// MockNoParamsFn_resultsByParams contains the results for a given set of parameters for the NoParamsFn type
type MockNoParamsFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MockNoParamsFn_paramsKey]*MockNoParamsFn_resultMgr
}

// MockNoParamsFn_resultMgr manages multiple results and the state of the NoParamsFn type
type MockNoParamsFn_resultMgr struct {
	Params   MockNoParamsFn_params
	Results  []*MockNoParamsFn_results
	Index    uint32
	AnyTimes bool
}

// MockNoParamsFn_results holds the results of the NoParamsFn type
type MockNoParamsFn_results struct {
	SResult      string
	Err          error
	Moq_Sequence uint32
}

// MockNoParamsFn_fnRecorder routes recorded function calls to the MockNoParamsFn mock
type MockNoParamsFn_fnRecorder struct {
	Params    MockNoParamsFn_params
	ParamsKey MockNoParamsFn_paramsKey
	AnyParams uint64
	Sequence  bool
	Results   *MockNoParamsFn_resultMgr
	Mock      *MockNoParamsFn
}

// NewMockNoParamsFn creates a new mock of the NoParamsFn type
func NewMockNoParamsFn(scene *moq.Scene, config *moq.MockConfig) *MockNoParamsFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockNoParamsFn{
		Scene:  scene,
		Config: *config,
	}
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the NoParamsFn type
func (m *MockNoParamsFn) Mock() testmocks.NoParamsFn {
	return func() (sResult string, err error) { mock := &MockNoParamsFn_mock{Mock: m}; return mock.Fn() }
}

func (m *MockNoParamsFn_mock) Fn() (sResult string, err error) {
	params := MockNoParamsFn_params{}
	var results *MockNoParamsFn_resultMgr
	for _, resultsByParams := range m.Mock.ResultsByParams {
		paramsKey := MockNoParamsFn_paramsKey{}
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
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
	if result.Moq_Sequence != 0 {
		sequence := m.Mock.Scene.NextMockSequence()
		if result.Moq_Sequence != sequence {
			m.Mock.Scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	sResult = result.SResult
	err = result.Err
	return
}

func (m *MockNoParamsFn) OnCall() *MockNoParamsFn_fnRecorder {
	return &MockNoParamsFn_fnRecorder{
		Params:    MockNoParamsFn_params{},
		ParamsKey: MockNoParamsFn_paramsKey{},
		Sequence:  m.Config.Sequence == moq.SeqDefaultOn,
		Mock:      m,
	}
}

func (r *MockNoParamsFn_fnRecorder) Seq() *MockNoParamsFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Seq must be called prior to returning results, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MockNoParamsFn_fnRecorder) NoSeq() *MockNoParamsFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("NoSeq must be called prior to returning results, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MockNoParamsFn_fnRecorder) ReturnResults(sResult string, err error) *MockNoParamsFn_fnRecorder {
	if r.Results == nil {
		anyCount := bits.OnesCount64(r.AnyParams)
		insertAt := -1
		var results *MockNoParamsFn_resultsByParams
		for n, res := range r.Mock.ResultsByParams {
			if res.AnyParams == r.AnyParams {
				results = &res
				break
			}
			if res.AnyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &MockNoParamsFn_resultsByParams{
				AnyCount:  anyCount,
				AnyParams: r.AnyParams,
				Results:   map[MockNoParamsFn_paramsKey]*MockNoParamsFn_resultMgr{},
			}
			r.Mock.ResultsByParams = append(r.Mock.ResultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.Mock.ResultsByParams) {
				copy(r.Mock.ResultsByParams[insertAt+1:], r.Mock.ResultsByParams[insertAt:0])
				r.Mock.ResultsByParams[insertAt] = *results
			}
		}

		paramsKey := MockNoParamsFn_paramsKey{}

		if _, ok := results.Results[paramsKey]; ok {
			r.Mock.Scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.Params)
			return nil
		}

		r.Results = &MockNoParamsFn_resultMgr{
			Params:   r.Params,
			Results:  []*MockNoParamsFn_results{},
			Index:    0,
			AnyTimes: false,
		}
		results.Results[paramsKey] = r.Results
	}

	var sequence uint32
	if r.Sequence {
		sequence = r.Mock.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, &MockNoParamsFn_results{
		SResult:      sResult,
		Err:          err,
		Moq_Sequence: sequence,
	})
	return r
}

func (r *MockNoParamsFn_fnRecorder) Times(count int) *MockNoParamsFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < count-1; n++ {
		if last.Moq_Sequence != 0 {
			last = &MockNoParamsFn_results{
				SResult:      last.SResult,
				Err:          last.Err,
				Moq_Sequence: r.Mock.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (r *MockNoParamsFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockNoParamsFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockNoParamsFn) AssertExpectationsMet() {
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := len(results.Results) - int(atomic.LoadUint32(&results.Index))
			if missing == 1 && results.AnyTimes == true {
				continue
			}
			if missing > 0 {
				m.Scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.Params)
			}
		}
	}
}
