// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package testmoqs_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/generator/testmoqs"
	"github.com/myshkin5/moqueries/moq"
)

// moqRepeatedIdsFn holds the state of a moq of the RepeatedIdsFn type
type moqRepeatedIdsFn struct {
	scene           *moq.Scene
	config          moq.Config
	resultsByParams []moqRepeatedIdsFn_resultsByParams
}

// moqRepeatedIdsFn_mock isolates the mock interface of the RepeatedIdsFn type
type moqRepeatedIdsFn_mock struct {
	moq *moqRepeatedIdsFn
}

// moqRepeatedIdsFn_params holds the params of the RepeatedIdsFn type
type moqRepeatedIdsFn_params struct {
	sParam1, sParam2 string
	bParam           bool
}

// moqRepeatedIdsFn_paramsKey holds the map key params of the RepeatedIdsFn type
type moqRepeatedIdsFn_paramsKey struct {
	sParam1, sParam2 string
	bParam           bool
}

// moqRepeatedIdsFn_resultsByParams contains the results for a given set of parameters for the RepeatedIdsFn type
type moqRepeatedIdsFn_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[moqRepeatedIdsFn_paramsKey]*moqRepeatedIdsFn_results
}

// moqRepeatedIdsFn_doFn defines the type of function needed when calling andDo for the RepeatedIdsFn type
type moqRepeatedIdsFn_doFn func(sParam1, sParam2 string, bParam bool)

// moqRepeatedIdsFn_doReturnFn defines the type of function needed when calling doReturnResults for the RepeatedIdsFn type
type moqRepeatedIdsFn_doReturnFn func(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error)

// moqRepeatedIdsFn_results holds the results of the RepeatedIdsFn type
type moqRepeatedIdsFn_results struct {
	params  moqRepeatedIdsFn_params
	results []struct {
		values *struct {
			sResult1, sResult2 string
			err                error
		}
		sequence   uint32
		doFn       moqRepeatedIdsFn_doFn
		doReturnFn moqRepeatedIdsFn_doReturnFn
	}
	index    uint32
	anyTimes bool
}

// moqRepeatedIdsFn_fnRecorder routes recorded function calls to the moqRepeatedIdsFn moq
type moqRepeatedIdsFn_fnRecorder struct {
	params    moqRepeatedIdsFn_params
	paramsKey moqRepeatedIdsFn_paramsKey
	anyParams uint64
	sequence  bool
	results   *moqRepeatedIdsFn_results
	moq       *moqRepeatedIdsFn
}

// moqRepeatedIdsFn_anyParams isolates the any params functions of the RepeatedIdsFn type
type moqRepeatedIdsFn_anyParams struct {
	recorder *moqRepeatedIdsFn_fnRecorder
}

// newMoqRepeatedIdsFn creates a new moq of the RepeatedIdsFn type
func newMoqRepeatedIdsFn(scene *moq.Scene, config *moq.Config) *moqRepeatedIdsFn {
	if config == nil {
		config = &moq.Config{}
	}
	m := &moqRepeatedIdsFn{
		scene:  scene,
		config: *config,
	}
	scene.AddMoq(m)
	return m
}

// mock returns the moq implementation of the RepeatedIdsFn type
func (m *moqRepeatedIdsFn) mock() testmoqs.RepeatedIdsFn {
	return func(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
		moq := &moqRepeatedIdsFn_mock{moq: m}
		return moq.fn(sParam1, sParam2, bParam)
	}
}

func (m *moqRepeatedIdsFn_mock) fn(sParam1, sParam2 string, bParam bool) (sResult1, sResult2 string, err error) {
	params := moqRepeatedIdsFn_params{
		sParam1: sParam1,
		sParam2: sParam2,
		bParam:  bParam,
	}
	var results *moqRepeatedIdsFn_results
	for _, resultsByParams := range m.moq.resultsByParams {
		var sParam1Used string
		if resultsByParams.anyParams&(1<<0) == 0 {
			sParam1Used = sParam1
		}
		var sParam2Used string
		if resultsByParams.anyParams&(1<<1) == 0 {
			sParam2Used = sParam2
		}
		var bParamUsed bool
		if resultsByParams.anyParams&(1<<2) == 0 {
			bParamUsed = bParam
		}
		paramsKey := moqRepeatedIdsFn_paramsKey{
			sParam1: sParam1Used,
			sParam2: sParam2Used,
			bParam:  bParamUsed,
		}
		var ok bool
		results, ok = resultsByParams.results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.moq.config.Expectation == moq.Strict {
			m.moq.scene.T.Fatalf("Unexpected call with parameters %#v", params)
		}
		return
	}

	i := int(atomic.AddUint32(&results.index, 1)) - 1
	if i >= len(results.results) {
		if !results.anyTimes {
			if m.moq.config.Expectation == moq.Strict {
				m.moq.scene.T.Fatalf("Too many calls to mock with parameters %#v", params)
			}
			return
		}
		i = len(results.results) - 1
	}

	result := results.results[i]
	if result.sequence != 0 {
		sequence := m.moq.scene.NextMockSequence()
		if (!results.anyTimes && result.sequence != sequence) || result.sequence > sequence {
			m.moq.scene.T.Fatalf("Call sequence does not match %#v", params)
		}
	}

	if result.doFn != nil {
		result.doFn(sParam1, sParam2, bParam)
	}

	if result.values != nil {
		sResult1 = result.values.sResult1
		sResult2 = result.values.sResult2
		err = result.values.err
	}
	if result.doReturnFn != nil {
		sResult1, sResult2, err = result.doReturnFn(sParam1, sParam2, bParam)
	}
	return
}

func (m *moqRepeatedIdsFn) onCall(sParam1, sParam2 string, bParam bool) *moqRepeatedIdsFn_fnRecorder {
	return &moqRepeatedIdsFn_fnRecorder{
		params: moqRepeatedIdsFn_params{
			sParam1: sParam1,
			sParam2: sParam2,
			bParam:  bParam,
		},
		paramsKey: moqRepeatedIdsFn_paramsKey{
			sParam1: sParam1,
			sParam2: sParam2,
			bParam:  bParam,
		},
		sequence: m.config.Sequence == moq.SeqDefaultOn,
		moq:      m,
	}
}

func (r *moqRepeatedIdsFn_fnRecorder) any() *moqRepeatedIdsFn_anyParams {
	if r.results != nil {
		r.moq.scene.T.Fatalf("Any functions must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	return &moqRepeatedIdsFn_anyParams{recorder: r}
}

func (a *moqRepeatedIdsFn_anyParams) sParam1() *moqRepeatedIdsFn_fnRecorder {
	a.recorder.anyParams |= 1 << 0
	return a.recorder
}

func (a *moqRepeatedIdsFn_anyParams) sParam2() *moqRepeatedIdsFn_fnRecorder {
	a.recorder.anyParams |= 1 << 1
	return a.recorder
}

func (a *moqRepeatedIdsFn_anyParams) bParam() *moqRepeatedIdsFn_fnRecorder {
	a.recorder.anyParams |= 1 << 2
	return a.recorder
}

func (r *moqRepeatedIdsFn_fnRecorder) seq() *moqRepeatedIdsFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("seq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) noSeq() *moqRepeatedIdsFn_fnRecorder {
	if r.results != nil {
		r.moq.scene.T.Fatalf("noSeq must be called before returnResults or doReturnResults calls, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) returnResults(sResult1, sResult2 string, err error) *moqRepeatedIdsFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			sResult1, sResult2 string
			err                error
		}
		sequence   uint32
		doFn       moqRepeatedIdsFn_doFn
		doReturnFn moqRepeatedIdsFn_doReturnFn
	}{
		values: &struct {
			sResult1, sResult2 string
			err                error
		}{
			sResult1: sResult1,
			sResult2: sResult2,
			err:      err,
		},
		sequence: sequence,
	})
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) andDo(fn moqRepeatedIdsFn_doFn) *moqRepeatedIdsFn_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults must be called before calling andDo")
		return nil
	}
	last := &r.results.results[len(r.results.results)-1]
	last.doFn = fn
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) doReturnResults(fn moqRepeatedIdsFn_doReturnFn) *moqRepeatedIdsFn_fnRecorder {
	r.findResults()

	var sequence uint32
	if r.sequence {
		sequence = r.moq.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, struct {
		values *struct {
			sResult1, sResult2 string
			err                error
		}
		sequence   uint32
		doFn       moqRepeatedIdsFn_doFn
		doReturnFn moqRepeatedIdsFn_doReturnFn
	}{sequence: sequence, doReturnFn: fn})
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) findResults() {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *moqRepeatedIdsFn_resultsByParams
		for n, res := range r.moq.resultsByParams {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &moqRepeatedIdsFn_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[moqRepeatedIdsFn_paramsKey]*moqRepeatedIdsFn_results{},
			}
			r.moq.resultsByParams = append(r.moq.resultsByParams, *results)
			if insertAt != -1 && insertAt+1 < len(r.moq.resultsByParams) {
				copy(r.moq.resultsByParams[insertAt+1:], r.moq.resultsByParams[insertAt:0])
				r.moq.resultsByParams[insertAt] = *results
			}
		}

		var sParam1Used string
		if r.anyParams&(1<<0) == 0 {
			sParam1Used = r.paramsKey.sParam1
		}
		var sParam2Used string
		if r.anyParams&(1<<1) == 0 {
			sParam2Used = r.paramsKey.sParam2
		}
		var bParamUsed bool
		if r.anyParams&(1<<2) == 0 {
			bParamUsed = r.paramsKey.bParam
		}
		paramsKey := moqRepeatedIdsFn_paramsKey{
			sParam1: sParam1Used,
			sParam2: sParam2Used,
			bParam:  bParamUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &moqRepeatedIdsFn_results{
				params:   r.params,
				results:  nil,
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}
}

func (r *moqRepeatedIdsFn_fnRecorder) times(count int) *moqRepeatedIdsFn_fnRecorder {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.sequence != 0 {
			last = struct {
				values *struct {
					sResult1, sResult2 string
					err                error
				}
				sequence   uint32
				doFn       moqRepeatedIdsFn_doFn
				doReturnFn moqRepeatedIdsFn_doReturnFn
			}{
				values: &struct {
					sResult1, sResult2 string
					err                error
				}{
					sResult1: last.values.sResult1,
					sResult2: last.values.sResult2,
					err:      last.values.err,
				},
				sequence: r.moq.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *moqRepeatedIdsFn_fnRecorder) anyTimes() {
	if r.results == nil {
		r.moq.scene.T.Fatalf("returnResults or doReturnResults must be called before calling anyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the moq
func (m *moqRepeatedIdsFn) Reset() { m.resultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *moqRepeatedIdsFn) AssertExpectationsMet() {
	for _, res := range m.resultsByParams {
		for _, results := range res.results {
			missing := len(results.results) - int(atomic.LoadUint32(&results.index))
			if missing == 1 && results.anyTimes == true {
				continue
			}
			if missing > 0 {
				m.scene.T.Errorf("Expected %d additional call(s) with parameters %#v", missing, results.params)
			}
		}
	}
}
