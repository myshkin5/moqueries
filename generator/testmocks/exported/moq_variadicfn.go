// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package exported

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmocks"
	"github.com/myshkin5/moqueries/hash"
	"github.com/myshkin5/moqueries/moq"
)

// MockVariadicFn holds the state of a mock of the VariadicFn type
type MockVariadicFn struct {
	Scene           *moq.Scene
	Config          moq.MockConfig
	ResultsByParams []MockVariadicFn_resultsByParams
}

// MockVariadicFn_mock isolates the mock interface of the VariadicFn type
type MockVariadicFn_mock struct {
	Mock *MockVariadicFn
}

// MockVariadicFn_params holds the params of the VariadicFn type
type MockVariadicFn_params struct {
	Other bool
	Args  []string
}

// MockVariadicFn_paramsKey holds the map key params of the VariadicFn type
type MockVariadicFn_paramsKey struct {
	Other bool
	Args  hash.Hash
}

// MockVariadicFn_resultsByParams contains the results for a given set of parameters for the VariadicFn type
type MockVariadicFn_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MockVariadicFn_paramsKey]*MockVariadicFn_results
}

// MockVariadicFn_doFn defines the type of function needed when calling AndDo for the VariadicFn type
type MockVariadicFn_doFn func(other bool, args ...string)

// MockVariadicFn_doReturnFn defines the type of function needed when calling DoReturnResults for the VariadicFn type
type MockVariadicFn_doReturnFn func(other bool, args ...string) (sResult string, err error)

// MockVariadicFn_results holds the results of the VariadicFn type
type MockVariadicFn_results struct {
	Params  MockVariadicFn_params
	Results []struct {
		Values *struct {
			SResult string
			Err     error
		}
		Sequence   uint32
		DoFn       MockVariadicFn_doFn
		DoReturnFn MockVariadicFn_doReturnFn
	}
	Index    uint32
	AnyTimes bool
}

// MockVariadicFn_fnRecorder routes recorded function calls to the MockVariadicFn mock
type MockVariadicFn_fnRecorder struct {
	Params    MockVariadicFn_params
	ParamsKey MockVariadicFn_paramsKey
	AnyParams uint64
	Sequence  bool
	Results   *MockVariadicFn_results
	Mock      *MockVariadicFn
}

// NewMockVariadicFn creates a new mock of the VariadicFn type
func NewMockVariadicFn(scene *moq.Scene, config *moq.MockConfig) *MockVariadicFn {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &MockVariadicFn{
		Scene:  scene,
		Config: *config,
	}
	scene.AddMock(m)
	return m
}

// Mock returns the mock implementation of the VariadicFn type
func (m *MockVariadicFn) Mock() testmocks.VariadicFn {
	return func(other bool, args ...string) (sResult string, err error) {
		mock := &MockVariadicFn_mock{Mock: m}
		return mock.Fn(other, args...)
	}
}

func (m *MockVariadicFn_mock) Fn(other bool, args ...string) (sResult string, err error) {
	params := MockVariadicFn_params{
		Other: other,
		Args:  args,
	}
	var results *MockVariadicFn_results
	for _, resultsByParams := range m.Mock.ResultsByParams {
		var otherUsed bool
		if resultsByParams.AnyParams&(1<<0) == 0 {
			otherUsed = other
		}
		var argsUsed hash.Hash
		if resultsByParams.AnyParams&(1<<1) == 0 {
			argsUsed = hash.DeepHash(args)
		}
		paramsKey := MockVariadicFn_paramsKey{
			Other: otherUsed,
			Args:  argsUsed,
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
		result.DoFn(other, args...)
	}

	if result.Values != nil {
		sResult = result.Values.SResult
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		sResult, err = result.DoReturnFn(other, args...)
	}
	return
}

func (m *MockVariadicFn) OnCall(other bool, args ...string) *MockVariadicFn_fnRecorder {
	return &MockVariadicFn_fnRecorder{
		Params: MockVariadicFn_params{
			Other: other,
			Args:  args,
		},
		ParamsKey: MockVariadicFn_paramsKey{
			Other: other,
			Args:  hash.DeepHash(args),
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Mock:     m,
	}
}

func (r *MockVariadicFn_fnRecorder) AnyOther() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 0
	return r
}

func (r *MockVariadicFn_fnRecorder) AnyArgs() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.AnyParams |= 1 << 1
	return r
}

func (r *MockVariadicFn_fnRecorder) Seq() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MockVariadicFn_fnRecorder) NoSeq() *MockVariadicFn_fnRecorder {
	if r.Results != nil {
		r.Mock.Scene.MoqT.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, parameters: %#v", r.Params)
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MockVariadicFn_fnRecorder) ReturnResults(sResult string, err error) *MockVariadicFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Mock.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			SResult string
			Err     error
		}
		Sequence   uint32
		DoFn       MockVariadicFn_doFn
		DoReturnFn MockVariadicFn_doReturnFn
	}{
		Values: &struct {
			SResult string
			Err     error
		}{
			SResult: sResult,
			Err:     err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MockVariadicFn_fnRecorder) AndDo(fn MockVariadicFn_doFn) *MockVariadicFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MockVariadicFn_fnRecorder) DoReturnResults(fn MockVariadicFn_doReturnFn) *MockVariadicFn_fnRecorder {
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Mock.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			SResult string
			Err     error
		}
		Sequence   uint32
		DoFn       MockVariadicFn_doFn
		DoReturnFn MockVariadicFn_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MockVariadicFn_fnRecorder) FindResults() {
	if r.Results == nil {
		anyCount := bits.OnesCount64(r.AnyParams)
		insertAt := -1
		var results *MockVariadicFn_resultsByParams
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
			results = &MockVariadicFn_resultsByParams{
				AnyCount:  anyCount,
				AnyParams: r.AnyParams,
				Results:   map[MockVariadicFn_paramsKey]*MockVariadicFn_results{},
			}
			r.Mock.ResultsByParams = append(r.Mock.ResultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.Mock.ResultsByParams) {
				copy(r.Mock.ResultsByParams[insertAt+1:], r.Mock.ResultsByParams[insertAt:0])
				r.Mock.ResultsByParams[insertAt] = *results
			}
		}

		var otherUsed bool
		if r.AnyParams&(1<<0) == 0 {
			otherUsed = r.ParamsKey.Other
		}
		var argsUsed hash.Hash
		if r.AnyParams&(1<<1) == 0 {
			argsUsed = r.ParamsKey.Args
		}
		paramsKey := MockVariadicFn_paramsKey{
			Other: otherUsed,
			Args:  argsUsed,
		}

		var ok bool
		r.Results, ok = results.Results[paramsKey]
		if !ok {
			r.Results = &MockVariadicFn_results{
				Params:   r.Params,
				Results:  nil,
				Index:    0,
				AnyTimes: false,
			}
			results.Results[paramsKey] = r.Results
		}
	}
}

func (r *MockVariadicFn_fnRecorder) Times(count int) *MockVariadicFn_fnRecorder {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults or DoReturnResults must be called before calling Times")
		return nil
	}
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < count-1; n++ {
		if last.Sequence != 0 {
			last = struct {
				Values *struct {
					SResult string
					Err     error
				}
				Sequence   uint32
				DoFn       MockVariadicFn_doFn
				DoReturnFn MockVariadicFn_doReturnFn
			}{
				Values: &struct {
					SResult string
					Err     error
				}{
					SResult: last.Values.SResult,
					Err:     last.Values.Err,
				},
				Sequence: r.Mock.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (r *MockVariadicFn_fnRecorder) AnyTimes() {
	if r.Results == nil {
		r.Mock.Scene.MoqT.Fatalf("ReturnResults or DoReturnResults must be called before calling AnyTimes")
		return
	}
	r.Results.AnyTimes = true
}

// Reset resets the state of the mock
func (m *MockVariadicFn) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MockVariadicFn) AssertExpectationsMet() {
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