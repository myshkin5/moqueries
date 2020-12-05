// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmocks"
	"github.com/myshkin5/moqueries/moq"
)

// MockNoNamesFn holds the state of a mock of the NoNamesFn type
type MockNoNamesFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams []MockNoNamesFn_resultsByParams
}

// MockNoNamesFn_mock isolates the mock interface of the NoNamesFn type
type MockNoNamesFn_mock struct {
	Mock *MockNoNamesFn
}

// MockNoNamesFn_params holds the params of the NoNamesFn type
type MockNoNamesFn_params struct {
	Param1 string
	Param2 bool
}

// MockNoNamesFn_paramsKey holds the map key params of the NoNamesFn type
type MockNoNamesFn_paramsKey struct {
	Param1 string
	Param2 bool
}

// MockNoNamesFn_resultsByParams contains the results for a given set of parameters for the NoNamesFn type
type MockNoNamesFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MockNoNamesFn_paramsKey]*MockNoNamesFn_results
}

// MockNoNamesFn_doFn defines the type of function needed when calling AndDo for the NoNamesFn type
type MockNoNamesFn_doFn func(string, bool)

// MockNoNamesFn_doReturnFn defines the type of function needed when calling DoReturnResults for the NoNamesFn type
type MockNoNamesFn_doReturnFn func(string, bool) (string, error)

// MockNoNamesFn_results holds the results of the NoNamesFn type
type MockNoNamesFn_results struct {
	Params  MockNoNamesFn_params
	Results []struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MockNoNamesFn_doFn
		DoReturnFn MockNoNamesFn_doReturnFn
	}
	Index    uint32
	AnyTimes bool
}

// MockNoNamesFn_fnRecorder routes recorded function calls to the MockNoNamesFn mock
type MockNoNamesFn_fnRecorder struct {
	Params    MockNoNamesFn_params
	ParamsKey MockNoNamesFn_paramsKey
	AnyParams uint64
	Sequence  bool
	Results   *MockNoNamesFn_results
	Mock      *MockNoNamesFn
}

// NewMockNoNamesFn creates a new mock of the NoNamesFn type
func NewMockNoNamesFn(scene *moq.Scene, config *moq.MockConfig) *MockNoNamesFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockNoNamesFn{
		Scene:  scene,
		Config: *config,
	}
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the NoNamesFn type
func (m *MockNoNamesFn) Mock() testmocks.NoNamesFn {
	return func(param1 string, param2 bool) (string, error) {
		mock := &MockNoNamesFn_mock{Mock: m}
		return mock.Fn(param1, param2)
	}
}

func (m *MockNoNamesFn_mock) Fn(param1 string, param2 bool) (result1 string, result2 error) {
	params := MockNoNamesFn_params{
		Param1: param1,
		Param2: param2,
	}
	var results *MockNoNamesFn_results
	for _, resultsByParams := range m.Mock.ResultsByParams {
		var param1Used string
		if resultsByParams.AnyParams&(1<<0) == 0 {
			param1Used = param1
		}
		var param2Used bool
		if resultsByParams.AnyParams&(1<<1) == 0 {
			param2Used = param2
		}
		paramsKey := MockNoNamesFn_paramsKey{
			Param1: param1Used,
			Param2: param2Used,
		}
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
	if result.Sequence != 0 {
		sequence := m.Mock.Scene.NextMockSequence()
		if (!results.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Mock.Scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1, param2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(param1, param2)
	}
	return
}

func (m *MockNoNamesFn) OnCall(param1 string, param2 bool) *MockNoNamesFn_fnRecorder {
	return &MockNoNamesFn_fnRecorder{
		Params: MockNoNamesFn_params{
			Param1: param1,
			Param2: param2,
		},
		ParamsKey: MockNoNamesFn_paramsKey{
			Param1: param1,
			Param2: param2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Mock:     m,
	}
}

func (r *MockNoNamesFn_fnRecorder) AnyParam1() *MockNoNamesFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 0
	return r
}

func (r *MockNoNamesFn_fnRecorder) AnyParam2() *MockNoNamesFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 1
	return r
}

func (r *MockNoNamesFn_fnRecorder) Seq() *MockNoNamesFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MockNoNamesFn_fnRecorder) NoSeq() *MockNoNamesFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MockNoNamesFn_fnRecorder) ReturnResults(result1 string, result2 error) *MockNoNamesFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Mock.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MockNoNamesFn_doFn
		DoReturnFn MockNoNamesFn_doReturnFn
	}{
		Values: &struct {
			Result1 string
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MockNoNamesFn_fnRecorder) AndDo(fn MockNoNamesFn_doFn) *MockNoNamesFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MockNoNamesFn_fnRecorder) DoReturnResults(fn MockNoNamesFn_doReturnFn) *MockNoNamesFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Mock.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MockNoNamesFn_doFn
		DoReturnFn MockNoNamesFn_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MockNoNamesFn_fnRecorder) FindResults() {
	if r.Results == nil {
		anyCount := bits.OnesCount64(r.AnyParams)
		insertAt := -1
		var results *MockNoNamesFn_resultsByParams
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
			results = &MockNoNamesFn_resultsByParams{
				AnyCount:  anyCount,
				AnyParams: r.AnyParams,
				Results:   map[MockNoNamesFn_paramsKey]*MockNoNamesFn_results{},
			}
			r.Mock.ResultsByParams = append(r.Mock.ResultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.Mock.ResultsByParams) {
				copy(r.Mock.ResultsByParams[insertAt+1:], r.Mock.ResultsByParams[insertAt:0])
				r.Mock.ResultsByParams[insertAt] = *results
			}
		}

		var param1Used string
		if r.AnyParams&(1<<0) == 0 {
			param1Used = r.ParamsKey.Param1
		}
		var param2Used bool
		if r.AnyParams&(1<<1) == 0 {
			param2Used = r.ParamsKey.Param2
		}
		paramsKey := MockNoNamesFn_paramsKey{
			Param1: param1Used,
			Param2: param2Used,
		}

		var ok bool
		r.Results, ok = results.Results[paramsKey]
		if !ok {
			r.Results = &MockNoNamesFn_results{
				Params:   r.Params,
				Results:  nil,
				Index:    0,
				AnyTimes: false,
			}
			results.Results[paramsKey] = r.Results
		}
	}
}

func (r *MockNoNamesFn_fnRecorder) Times(count int) *MockNoNamesFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults or DoReturnResults must be called before calling Times")
		return nil
	}
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < count-1; n++ {
		if last.Sequence != 0 {
			last = struct {
				Values *struct {
					Result1 string
					Result2 error
				}
				Sequence   uint32
				DoFn       MockNoNamesFn_doFn
				DoReturnFn MockNoNamesFn_doReturnFn
			}{
				Values: &struct {
					Result1 string
					Result2 error
				}{
					Result1: last.Values.Result1,
					Result2: last.Values.Result2,
				},
				Sequence: r.Mock.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (r *MockNoNamesFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults or DoReturnResults must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockNoNamesFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockNoNamesFn) AssertExpectationsMet() {
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