// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmocks_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockNoResultsFn holds the state of a mock of the NoResultsFn type
type mockNoResultsFn struct {
	scene           *moq.Scene
	config          moq.MockConfig
	resultsByParams []mockNoResultsFn_resultsByParams
}

// mockNoResultsFn_mock isolates the mock interface of the NoResultsFn type
type mockNoResultsFn_mock struct {
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

// mockNoResultsFn_resultsByParams contains the results for a given set of parameters for the NoResultsFn type
type mockNoResultsFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockNoResultsFn_paramsKey]*mockNoResultsFn_resultMgr
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
	moq_sequence uint32
}

// mockNoResultsFn_fnRecorder routes recorded function calls to the mockNoResultsFn mock
type mockNoResultsFn_fnRecorder struct {
	params    mockNoResultsFn_params
	paramsKey mockNoResultsFn_paramsKey
	anyParams uint64
	sequence  bool
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
	var results *mockNoResultsFn_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams {
		var sParamUsed string
		if resultsByParams.anyParams&(1<<0) == 0 {
			sParamUsed = sParam
		}
		var bParamUsed bool
		if resultsByParams.anyParams&(1<<1) == 0 {
			bParamUsed = bParam
		}
		paramsKey := mockNoResultsFn_paramsKey{
			sParam: sParamUsed,
			bParam: bParamUsed,
		}
		var ok bool
		results, ok = resultsByParams.results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
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
	if result.moq_sequence != 0 {
		sequence := m.mock.scene.NextMockSequence()
		if result.moq_sequence != sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
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
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		mock:     m,
	}
}

func (r *mockNoResultsFn_fnRecorder) anySParam() *mockNoResultsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockNoResultsFn_fnRecorder) anyBParam() *mockNoResultsFn_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockNoResultsFn_fnRecorder) seq() *mockNoResultsFn_fnRecorder {
	r.sequence = true
	return r
}

func (r *mockNoResultsFn_fnRecorder) noSeq() *mockNoResultsFn_fnRecorder {
	r.sequence = false
	return r
}

func (r *mockNoResultsFn_fnRecorder) returnResults() *mockNoResultsFn_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockNoResultsFn_resultsByParams
		for n, res := range r.mock.resultsByParams {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockNoResultsFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockNoResultsFn_paramsKey]*mockNoResultsFn_resultMgr{},
			}
			r.mock.resultsByParams = append(r.mock.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams) {
				copy(r.mock.resultsByParams[insertAt+1:], r.mock.resultsByParams[insertAt:0])
				r.mock.resultsByParams[insertAt] = *results
			}
		}

		var sParamUsed string
		if r.anyParams&(1<<0) == 0 {
			sParamUsed = r.paramsKey.sParam
		}
		var bParamUsed bool
		if r.anyParams&(1<<1) == 0 {
			bParamUsed = r.paramsKey.bParam
		}
		paramsKey := mockNoResultsFn_paramsKey{
			sParam: sParamUsed,
			bParam: bParamUsed,
		}

		if _, ok := results.results[paramsKey]; ok {
			r.mock.scene.MoqT.Fatalf("Expectations already recorded for mock with parameters %#v", r.params)
			return nil
		}

		r.results = &mockNoResultsFn_resultMgr{
			params:   r.params,
			results:  []*mockNoResultsFn_results{},
			index:    0,
			anyTimes: false,
		}
		results.results[paramsKey] = r.results
	}

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, &mockNoResultsFn_results{moq_sequence: sequence})
	return r
}

func (r *mockNoResultsFn_fnRecorder) times(count int) *mockNoResultsFn_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.moq_sequence != 0 {
			last = &mockNoResultsFn_results{moq_sequence: r.mock.scene.NextRecorderSequence()}
		}
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
func (m *mockNoResultsFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockNoResultsFn) AssertExpectationsMet() {
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := len(results.results) - int(atomic.LoadUint32(&results.index))
			if missing == 1 && results.anyTimes == true {
				continue
			}
			if missing > 0 {
				m.scene.MoqT.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}
