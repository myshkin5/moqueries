// Code generated by Moqueries - https://github.com/myshkin5/moqueries - DO NOT EDIT!

package demo_test

import (
	"math/bits"
	"sync/atomic"

	"github.com/myshkin5/moqueries/demo"
	"github.com/myshkin5/moqueries/pkg/moq"
)

// mockStore holds the state of a mock of the Store type
type mockStore struct {
	scene                                  *moq.Scene
	config                                 moq.MockConfig
	resultsByParams_AllWidgetsIds          []mockStore_AllWidgetsIds_resultsByParams
	resultsByParams_GadgetsByWidgetId      []mockStore_GadgetsByWidgetId_resultsByParams
	resultsByParams_LightGadgetsByWidgetId []mockStore_LightGadgetsByWidgetId_resultsByParams
}

// mockStore_mock isolates the mock interface of the Store type
type mockStore_mock struct {
	mock *mockStore
}

// mockStore_recorder isolates the recorder interface of the Store type
type mockStore_recorder struct {
	mock *mockStore
}

// mockStore_AllWidgetsIds_params holds the params of the Store type
type mockStore_AllWidgetsIds_params struct{}

// mockStore_AllWidgetsIds_paramsKey holds the map key params of the Store type
type mockStore_AllWidgetsIds_paramsKey struct{}

// mockStore_AllWidgetsIds_resultsByParams contains the results for a given set of parameters for the Store type
type mockStore_AllWidgetsIds_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockStore_AllWidgetsIds_paramsKey]*mockStore_AllWidgetsIds_resultMgr
}

// mockStore_AllWidgetsIds_resultMgr manages multiple results and the state of the Store type
type mockStore_AllWidgetsIds_resultMgr struct {
	params   mockStore_AllWidgetsIds_params
	results  []*mockStore_AllWidgetsIds_results
	index    uint32
	anyTimes bool
}

// mockStore_AllWidgetsIds_results holds the results of the Store type
type mockStore_AllWidgetsIds_results struct {
	result1      []int
	result2      error
	moq_sequence uint32
}

// mockStore_AllWidgetsIds_fnRecorder routes recorded function calls to the mockStore mock
type mockStore_AllWidgetsIds_fnRecorder struct {
	params    mockStore_AllWidgetsIds_params
	paramsKey mockStore_AllWidgetsIds_paramsKey
	anyParams uint64
	sequence  bool
	results   *mockStore_AllWidgetsIds_resultMgr
	mock      *mockStore
}

// mockStore_GadgetsByWidgetId_params holds the params of the Store type
type mockStore_GadgetsByWidgetId_params struct{ widgetId int }

// mockStore_GadgetsByWidgetId_paramsKey holds the map key params of the Store type
type mockStore_GadgetsByWidgetId_paramsKey struct{ widgetId int }

// mockStore_GadgetsByWidgetId_resultsByParams contains the results for a given set of parameters for the Store type
type mockStore_GadgetsByWidgetId_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockStore_GadgetsByWidgetId_paramsKey]*mockStore_GadgetsByWidgetId_resultMgr
}

// mockStore_GadgetsByWidgetId_resultMgr manages multiple results and the state of the Store type
type mockStore_GadgetsByWidgetId_resultMgr struct {
	params   mockStore_GadgetsByWidgetId_params
	results  []*mockStore_GadgetsByWidgetId_results
	index    uint32
	anyTimes bool
}

// mockStore_GadgetsByWidgetId_results holds the results of the Store type
type mockStore_GadgetsByWidgetId_results struct {
	result1      []demo.Gadget
	result2      error
	moq_sequence uint32
}

// mockStore_GadgetsByWidgetId_fnRecorder routes recorded function calls to the mockStore mock
type mockStore_GadgetsByWidgetId_fnRecorder struct {
	params    mockStore_GadgetsByWidgetId_params
	paramsKey mockStore_GadgetsByWidgetId_paramsKey
	anyParams uint64
	sequence  bool
	results   *mockStore_GadgetsByWidgetId_resultMgr
	mock      *mockStore
}

// mockStore_LightGadgetsByWidgetId_params holds the params of the Store type
type mockStore_LightGadgetsByWidgetId_params struct {
	widgetId  int
	maxWeight uint32
}

// mockStore_LightGadgetsByWidgetId_paramsKey holds the map key params of the Store type
type mockStore_LightGadgetsByWidgetId_paramsKey struct {
	widgetId  int
	maxWeight uint32
}

// mockStore_LightGadgetsByWidgetId_resultsByParams contains the results for a given set of parameters for the Store type
type mockStore_LightGadgetsByWidgetId_resultsByParams struct {
	anyCount  int
	anyParams uint64
	results   map[mockStore_LightGadgetsByWidgetId_paramsKey]*mockStore_LightGadgetsByWidgetId_resultMgr
}

// mockStore_LightGadgetsByWidgetId_resultMgr manages multiple results and the state of the Store type
type mockStore_LightGadgetsByWidgetId_resultMgr struct {
	params   mockStore_LightGadgetsByWidgetId_params
	results  []*mockStore_LightGadgetsByWidgetId_results
	index    uint32
	anyTimes bool
}

// mockStore_LightGadgetsByWidgetId_results holds the results of the Store type
type mockStore_LightGadgetsByWidgetId_results struct {
	result1      []demo.Gadget
	result2      error
	moq_sequence uint32
}

// mockStore_LightGadgetsByWidgetId_fnRecorder routes recorded function calls to the mockStore mock
type mockStore_LightGadgetsByWidgetId_fnRecorder struct {
	params    mockStore_LightGadgetsByWidgetId_params
	paramsKey mockStore_LightGadgetsByWidgetId_paramsKey
	anyParams uint64
	sequence  bool
	results   *mockStore_LightGadgetsByWidgetId_resultMgr
	mock      *mockStore
}

// newMockStore creates a new mock of the Store type
func newMockStore(scene *moq.Scene, config *moq.MockConfig) *mockStore {
	if config == nil {
		config = &moq.MockConfig{}
	}
	m := &mockStore{
		scene:  scene,
		config: *config,
	}
	scene.AddMock(m)
	return m
}

// mock returns the mock implementation of the Store type
func (m *mockStore) mock() *mockStore_mock {
	return &mockStore_mock{
		mock: m,
	}
}

func (m *mockStore_mock) AllWidgetsIds() (result1 []int, result2 error) {
	params := mockStore_AllWidgetsIds_params{}
	var results *mockStore_AllWidgetsIds_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams_AllWidgetsIds {
		paramsKey := mockStore_AllWidgetsIds_paramsKey{}
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
		if (!results.anyTimes && result.moq_sequence != sequence) || result.moq_sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	result1 = result.result1
	result2 = result.result2
	return
}

func (m *mockStore_mock) GadgetsByWidgetId(widgetId int) (result1 []demo.Gadget, result2 error) {
	params := mockStore_GadgetsByWidgetId_params{
		widgetId: widgetId,
	}
	var results *mockStore_GadgetsByWidgetId_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams_GadgetsByWidgetId {
		var widgetIdUsed int
		if resultsByParams.anyParams&(1<<0) == 0 {
			widgetIdUsed = widgetId
		}
		paramsKey := mockStore_GadgetsByWidgetId_paramsKey{
			widgetId: widgetIdUsed,
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
		if (!results.anyTimes && result.moq_sequence != sequence) || result.moq_sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	result1 = result.result1
	result2 = result.result2
	return
}

func (m *mockStore_mock) LightGadgetsByWidgetId(widgetId int, maxWeight uint32) (result1 []demo.Gadget, result2 error) {
	params := mockStore_LightGadgetsByWidgetId_params{
		widgetId:  widgetId,
		maxWeight: maxWeight,
	}
	var results *mockStore_LightGadgetsByWidgetId_resultMgr
	for _, resultsByParams := range m.mock.resultsByParams_LightGadgetsByWidgetId {
		var widgetIdUsed int
		if resultsByParams.anyParams&(1<<0) == 0 {
			widgetIdUsed = widgetId
		}
		var maxWeightUsed uint32
		if resultsByParams.anyParams&(1<<1) == 0 {
			maxWeightUsed = maxWeight
		}
		paramsKey := mockStore_LightGadgetsByWidgetId_paramsKey{
			widgetId:  widgetIdUsed,
			maxWeight: maxWeightUsed,
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
		if (!results.anyTimes && result.moq_sequence != sequence) || result.moq_sequence > sequence {
			m.mock.scene.MoqT.Fatalf("Call sequence does not match %#v", params)
		}
	}

	result1 = result.result1
	result2 = result.result2
	return
}

// onCall returns the recorder implementation of the Store type
func (m *mockStore) onCall() *mockStore_recorder {
	return &mockStore_recorder{
		mock: m,
	}
}

func (m *mockStore_recorder) AllWidgetsIds() *mockStore_AllWidgetsIds_fnRecorder {
	return &mockStore_AllWidgetsIds_fnRecorder{
		params:    mockStore_AllWidgetsIds_params{},
		paramsKey: mockStore_AllWidgetsIds_paramsKey{},
		sequence:  m.mock.config.Sequence == moq.SeqDefaultOn,
		mock:      m.mock,
	}
}

func (r *mockStore_AllWidgetsIds_fnRecorder) seq() *mockStore_AllWidgetsIds_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockStore_AllWidgetsIds_fnRecorder) noSeq() *mockStore_AllWidgetsIds_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockStore_AllWidgetsIds_fnRecorder) returnResults(result1 []int, result2 error) *mockStore_AllWidgetsIds_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockStore_AllWidgetsIds_resultsByParams
		for n, res := range r.mock.resultsByParams_AllWidgetsIds {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockStore_AllWidgetsIds_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockStore_AllWidgetsIds_paramsKey]*mockStore_AllWidgetsIds_resultMgr{},
			}
			r.mock.resultsByParams_AllWidgetsIds = append(r.mock.resultsByParams_AllWidgetsIds, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams_AllWidgetsIds) {
				copy(r.mock.resultsByParams_AllWidgetsIds[insertAt+1:], r.mock.resultsByParams_AllWidgetsIds[insertAt:0])
				r.mock.resultsByParams_AllWidgetsIds[insertAt] = *results
			}
		}

		paramsKey := mockStore_AllWidgetsIds_paramsKey{}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &mockStore_AllWidgetsIds_resultMgr{
				params:   r.params,
				results:  []*mockStore_AllWidgetsIds_results{},
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, &mockStore_AllWidgetsIds_results{
		result1:      result1,
		result2:      result2,
		moq_sequence: sequence,
	})
	return r
}

func (r *mockStore_AllWidgetsIds_fnRecorder) times(count int) *mockStore_AllWidgetsIds_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.moq_sequence != 0 {
			last = &mockStore_AllWidgetsIds_results{
				result1:      last.result1,
				result2:      last.result2,
				moq_sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockStore_AllWidgetsIds_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockStore_recorder) GadgetsByWidgetId(widgetId int) *mockStore_GadgetsByWidgetId_fnRecorder {
	return &mockStore_GadgetsByWidgetId_fnRecorder{
		params: mockStore_GadgetsByWidgetId_params{
			widgetId: widgetId,
		},
		paramsKey: mockStore_GadgetsByWidgetId_paramsKey{
			widgetId: widgetId,
		},
		sequence: m.mock.config.Sequence == moq.SeqDefaultOn,
		mock:     m.mock,
	}
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) anyWidgetId() *mockStore_GadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) seq() *mockStore_GadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) noSeq() *mockStore_GadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) returnResults(result1 []demo.Gadget, result2 error) *mockStore_GadgetsByWidgetId_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockStore_GadgetsByWidgetId_resultsByParams
		for n, res := range r.mock.resultsByParams_GadgetsByWidgetId {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockStore_GadgetsByWidgetId_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockStore_GadgetsByWidgetId_paramsKey]*mockStore_GadgetsByWidgetId_resultMgr{},
			}
			r.mock.resultsByParams_GadgetsByWidgetId = append(r.mock.resultsByParams_GadgetsByWidgetId, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams_GadgetsByWidgetId) {
				copy(r.mock.resultsByParams_GadgetsByWidgetId[insertAt+1:], r.mock.resultsByParams_GadgetsByWidgetId[insertAt:0])
				r.mock.resultsByParams_GadgetsByWidgetId[insertAt] = *results
			}
		}

		var widgetIdUsed int
		if r.anyParams&(1<<0) == 0 {
			widgetIdUsed = r.paramsKey.widgetId
		}
		paramsKey := mockStore_GadgetsByWidgetId_paramsKey{
			widgetId: widgetIdUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &mockStore_GadgetsByWidgetId_resultMgr{
				params:   r.params,
				results:  []*mockStore_GadgetsByWidgetId_results{},
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, &mockStore_GadgetsByWidgetId_results{
		result1:      result1,
		result2:      result2,
		moq_sequence: sequence,
	})
	return r
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) times(count int) *mockStore_GadgetsByWidgetId_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.moq_sequence != 0 {
			last = &mockStore_GadgetsByWidgetId_results{
				result1:      last.result1,
				result2:      last.result2,
				moq_sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockStore_GadgetsByWidgetId_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

func (m *mockStore_recorder) LightGadgetsByWidgetId(widgetId int, maxWeight uint32) *mockStore_LightGadgetsByWidgetId_fnRecorder {
	return &mockStore_LightGadgetsByWidgetId_fnRecorder{
		params: mockStore_LightGadgetsByWidgetId_params{
			widgetId:  widgetId,
			maxWeight: maxWeight,
		},
		paramsKey: mockStore_LightGadgetsByWidgetId_paramsKey{
			widgetId:  widgetId,
			maxWeight: maxWeight,
		},
		sequence: m.mock.config.Sequence == moq.SeqDefaultOn,
		mock:     m.mock,
	}
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) anyWidgetId() *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 0
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) anyMaxWeight() *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("Any functions must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.anyParams |= 1 << 1
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) seq() *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("seq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = true
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) noSeq() *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results != nil {
		r.mock.scene.MoqT.Fatalf("noSeq must be called prior to returning results, parameters: %#v", r.params)
		return nil
	}
	r.sequence = false
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) returnResults(result1 []demo.Gadget, result2 error) *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results == nil {
		anyCount := bits.OnesCount64(r.anyParams)
		insertAt := -1
		var results *mockStore_LightGadgetsByWidgetId_resultsByParams
		for n, res := range r.mock.resultsByParams_LightGadgetsByWidgetId {
			if res.anyParams == r.anyParams {
				results = &res
				break
			}
			if res.anyCount > anyCount {
				insertAt = n
			}
		}
		if results == nil {
			results = &mockStore_LightGadgetsByWidgetId_resultsByParams{
				anyCount:  anyCount,
				anyParams: r.anyParams,
				results:   map[mockStore_LightGadgetsByWidgetId_paramsKey]*mockStore_LightGadgetsByWidgetId_resultMgr{},
			}
			r.mock.resultsByParams_LightGadgetsByWidgetId = append(r.mock.resultsByParams_LightGadgetsByWidgetId, *results)
			if insertAt != -1 && insertAt+1 < len(r.mock.resultsByParams_LightGadgetsByWidgetId) {
				copy(r.mock.resultsByParams_LightGadgetsByWidgetId[insertAt+1:], r.mock.resultsByParams_LightGadgetsByWidgetId[insertAt:0])
				r.mock.resultsByParams_LightGadgetsByWidgetId[insertAt] = *results
			}
		}

		var widgetIdUsed int
		if r.anyParams&(1<<0) == 0 {
			widgetIdUsed = r.paramsKey.widgetId
		}
		var maxWeightUsed uint32
		if r.anyParams&(1<<1) == 0 {
			maxWeightUsed = r.paramsKey.maxWeight
		}
		paramsKey := mockStore_LightGadgetsByWidgetId_paramsKey{
			widgetId:  widgetIdUsed,
			maxWeight: maxWeightUsed,
		}

		var ok bool
		r.results, ok = results.results[paramsKey]
		if !ok {
			r.results = &mockStore_LightGadgetsByWidgetId_resultMgr{
				params:   r.params,
				results:  []*mockStore_LightGadgetsByWidgetId_results{},
				index:    0,
				anyTimes: false,
			}
			results.results[paramsKey] = r.results
		}
	}

	var sequence uint32
	if r.sequence {
		sequence = r.mock.scene.NextRecorderSequence()
	}

	r.results.results = append(r.results.results, &mockStore_LightGadgetsByWidgetId_results{
		result1:      result1,
		result2:      result2,
		moq_sequence: sequence,
	})
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) times(count int) *mockStore_LightGadgetsByWidgetId_fnRecorder {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling Times")
		return nil
	}
	last := r.results.results[len(r.results.results)-1]
	for n := 0; n < count-1; n++ {
		if last.moq_sequence != 0 {
			last = &mockStore_LightGadgetsByWidgetId_results{
				result1:      last.result1,
				result2:      last.result2,
				moq_sequence: r.mock.scene.NextRecorderSequence(),
			}
		}
		r.results.results = append(r.results.results, last)
	}
	return r
}

func (r *mockStore_LightGadgetsByWidgetId_fnRecorder) anyTimes() {
	if r.results == nil {
		r.mock.scene.MoqT.Fatalf("Return must be called before calling AnyTimes")
		return
	}
	r.results.anyTimes = true
}

// Reset resets the state of the mock
func (m *mockStore) Reset() {
	m.resultsByParams_AllWidgetsIds = nil
	m.resultsByParams_GadgetsByWidgetId = nil
	m.resultsByParams_LightGadgetsByWidgetId = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *mockStore) AssertExpectationsMet() {
	for _, res := range m.resultsByParams_AllWidgetsIds {
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
	for _, res := range m.resultsByParams_GadgetsByWidgetId {
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
	for _, res := range m.resultsByParams_LightGadgetsByWidgetId {
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
