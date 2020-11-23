package testmocks_test

import (
	. "github.com/onsi/gomega"

	"github.com/myshkin5/moqueries/pkg/generator/testmocks/exported"
	"github.com/myshkin5/moqueries/pkg/moq"
)

type usualAdaptor struct{ m *mockUsual }

func (a *usualAdaptor) exported() bool { return false }

func (a *usualAdaptor) tracksParams() bool { return true }

func (a *usualAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &usualRecorder{r: a.m.onCall().Usual(sParams[0], bParam)}
}

func (a *usualAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.mock().Usual(sParams[0], bParam)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *usualAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return mockUsual_Usual_params{sParam: sParams[0], bParam: bParam}
}

func (a *usualAdaptor) sceneMock() moq.Mock {
	return a.m
}

type usualRecorder struct{ r *mockUsual_Usual_fnRecorder }

func (r *usualRecorder) anySParam() {
	r.r = r.r.anySParam()
}

func (r *usualRecorder) anyBParam() {
	r.r = r.r.anyBParam()
}

func (r *usualRecorder) seq() {
	r.r = r.r.seq()
}

func (r *usualRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *usualRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.returnResults(sResults[0], err)
}

func (r *usualRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *usualRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *usualRecorder) isNil() bool {
	return r.r == nil
}

type exportedUsualAdaptor struct{ m *exported.MockUsual }

func (a *exportedUsualAdaptor) exported() bool { return true }

func (a *exportedUsualAdaptor) tracksParams() bool { return true }

func (a *exportedUsualAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &exportedUsualRecorder{r: a.m.OnCall().Usual(sParams[0], bParam)}
}

func (a *exportedUsualAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.Mock().Usual(sParams[0], bParam)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *exportedUsualAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return exported.MockUsual_Usual_params{SParam: sParams[0], BParam: bParam}
}

func (a *exportedUsualAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedUsualRecorder struct {
	r *exported.MockUsual_Usual_fnRecorder
}

func (r *exportedUsualRecorder) anySParam() {
	r.r = r.r.AnySParam()
}

func (r *exportedUsualRecorder) anyBParam() {
	r.r = r.r.AnyBParam()
}

func (r *exportedUsualRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedUsualRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedUsualRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.ReturnResults(sResults[0], err)
}

func (r *exportedUsualRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedUsualRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedUsualRecorder) isNil() bool {
	return r.r == nil
}

type noNamesAdaptor struct{ m *mockUsual }

func (a *noNamesAdaptor) exported() bool { return false }

func (a *noNamesAdaptor) tracksParams() bool { return true }

func (a *noNamesAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &noNamesRecorder{r: a.m.onCall().NoNames(sParams[0], bParam)}
}

func (a *noNamesAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.mock().NoNames(sParams[0], bParam)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *noNamesAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return mockUsual_NoNames_params{param1: sParams[0], param2: bParam}
}

func (a *noNamesAdaptor) sceneMock() moq.Mock {
	return a.m
}

type noNamesRecorder struct{ r *mockUsual_NoNames_fnRecorder }

func (r *noNamesRecorder) anySParam() {
	r.r = r.r.anyParam1()
}

func (r *noNamesRecorder) anyBParam() {
	r.r = r.r.anyParam2()
}

func (r *noNamesRecorder) seq() {
	r.r = r.r.seq()
}

func (r *noNamesRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *noNamesRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.returnResults(sResults[0], err)
}

func (r *noNamesRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *noNamesRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *noNamesRecorder) isNil() bool {
	return r.r == nil
}

type exportedNoNamesAdaptor struct{ m *exported.MockUsual }

func (a *exportedNoNamesAdaptor) exported() bool { return true }

func (a *exportedNoNamesAdaptor) tracksParams() bool { return true }

func (a *exportedNoNamesAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &exportedNoNamesRecorder{r: a.m.OnCall().NoNames(sParams[0], bParam)}
}

func (a *exportedNoNamesAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.Mock().NoNames(sParams[0], bParam)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *exportedNoNamesAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return exported.MockUsual_NoNames_params{Param1: sParams[0], Param2: bParam}
}

func (a *exportedNoNamesAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedNoNamesRecorder struct {
	r *exported.MockUsual_NoNames_fnRecorder
}

func (r *exportedNoNamesRecorder) anySParam() {
	r.r = r.r.AnyParam1()
}

func (r *exportedNoNamesRecorder) anyBParam() {
	r.r = r.r.AnyParam2()
}

func (r *exportedNoNamesRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedNoNamesRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedNoNamesRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.ReturnResults(sResults[0], err)
}

func (r *exportedNoNamesRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedNoNamesRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedNoNamesRecorder) isNil() bool {
	return r.r == nil
}

type noResultsAdaptor struct{ m *mockUsual }

func (a *noResultsAdaptor) exported() bool { return false }

func (a *noResultsAdaptor) tracksParams() bool { return true }

func (a *noResultsAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &noResultsRecorder{r: a.m.onCall().NoResults(sParams[0], bParam)}
}

func (a *noResultsAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, _ results) {
	a.m.mock().NoResults(sParams[0], bParam)
}

func (a *noResultsAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return mockUsual_NoResults_params{sParam: sParams[0], bParam: bParam}
}

func (a *noResultsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type noResultsRecorder struct {
	r *mockUsual_NoResults_fnRecorder
}

func (r *noResultsRecorder) anySParam() {
	r.r = r.r.anySParam()
}

func (r *noResultsRecorder) anyBParam() {
	r.r = r.r.anyBParam()
}

func (r *noResultsRecorder) seq() {
	r.r = r.r.seq()
}

func (r *noResultsRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *noResultsRecorder) returnResults([]string, error) {
	r.r = r.r.returnResults()
}

func (r *noResultsRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *noResultsRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *noResultsRecorder) isNil() bool {
	return r.r == nil
}

type exportedNoResultsAdaptor struct{ m *exported.MockUsual }

func (a *exportedNoResultsAdaptor) exported() bool { return true }

func (a *exportedNoResultsAdaptor) tracksParams() bool { return true }

func (a *exportedNoResultsAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &exportedNoResultsRecorder{r: a.m.OnCall().NoResults(sParams[0], bParam)}
}

func (a *exportedNoResultsAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, _ results) {
	a.m.Mock().NoResults(sParams[0], bParam)
}

func (a *exportedNoResultsAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return exported.MockUsual_NoResults_params{SParam: sParams[0], BParam: bParam}
}

func (a *exportedNoResultsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedNoResultsRecorder struct {
	r *exported.MockUsual_NoResults_fnRecorder
}

func (r *exportedNoResultsRecorder) anySParam() {
	r.r = r.r.AnySParam()
}

func (r *exportedNoResultsRecorder) anyBParam() {
	r.r = r.r.AnyBParam()
}

func (r *exportedNoResultsRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedNoResultsRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedNoResultsRecorder) returnResults([]string, error) {
	r.r = r.r.ReturnResults()
}

func (r *exportedNoResultsRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedNoResultsRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedNoResultsRecorder) isNil() bool {
	return r.r == nil
}

type noParamsAdaptor struct{ m *mockUsual }

func (a *noParamsAdaptor) exported() bool { return false }

func (a *noParamsAdaptor) tracksParams() bool { return false }

func (a *noParamsAdaptor) newRecorder([]string, bool) recorder {
	return &noParamsRecorder{r: a.m.onCall().NoParams()}
}

func (a *noParamsAdaptor) invokeMockAndExpectResults(_ []string, _ bool, res results) {
	sResult, err := a.m.mock().NoParams()
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *noParamsAdaptor) bundleParams([]string, bool) interface{} {
	return mockUsual_NoParams_params{}
}

func (a *noParamsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type noParamsRecorder struct {
	r *mockUsual_NoParams_fnRecorder
}

func (r *noParamsRecorder) anySParam() {}

func (r *noParamsRecorder) anyBParam() {}

func (r *noParamsRecorder) seq() {
	r.r = r.r.seq()
}

func (r *noParamsRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *noParamsRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.returnResults(sResults[0], err)
}

func (r *noParamsRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *noParamsRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *noParamsRecorder) isNil() bool {
	return r.r == nil
}

type exportedNoParamsAdaptor struct{ m *exported.MockUsual }

func (a *exportedNoParamsAdaptor) exported() bool { return true }

func (a *exportedNoParamsAdaptor) tracksParams() bool { return false }

func (a *exportedNoParamsAdaptor) newRecorder([]string, bool) recorder {
	return &exportedNoParamsRecorder{r: a.m.OnCall().NoParams()}
}

func (a *exportedNoParamsAdaptor) invokeMockAndExpectResults(_ []string, _ bool, res results) {
	sResult, err := a.m.Mock().NoParams()
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *exportedNoParamsAdaptor) bundleParams([]string, bool) interface{} {
	return exported.MockUsual_NoParams_params{}
}

func (a *exportedNoParamsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedNoParamsRecorder struct {
	r *exported.MockUsual_NoParams_fnRecorder
}

func (r *exportedNoParamsRecorder) anySParam() {}

func (r *exportedNoParamsRecorder) anyBParam() {}

func (r *exportedNoParamsRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedNoParamsRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedNoParamsRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.ReturnResults(sResults[0], err)
}

func (r *exportedNoParamsRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedNoParamsRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedNoParamsRecorder) isNil() bool {
	return r.r == nil
}

type nothingAdaptor struct{ m *mockUsual }

func (a *nothingAdaptor) exported() bool { return false }

func (a *nothingAdaptor) tracksParams() bool { return false }

func (a *nothingAdaptor) newRecorder([]string, bool) recorder {
	return &nothingRecorder{r: a.m.onCall().Nothing()}
}

func (a *nothingAdaptor) invokeMockAndExpectResults([]string, bool, results) {
	a.m.mock().Nothing()
}

func (a *nothingAdaptor) bundleParams([]string, bool) interface{} {
	return mockUsual_Nothing_params{}
}

func (a *nothingAdaptor) sceneMock() moq.Mock {
	return a.m
}

type nothingRecorder struct{ r *mockUsual_Nothing_fnRecorder }

func (r *nothingRecorder) anySParam() {}

func (r *nothingRecorder) anyBParam() {}

func (r *nothingRecorder) seq() {
	r.r = r.r.seq()
}

func (r *nothingRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *nothingRecorder) returnResults([]string, error) {
	r.r = r.r.returnResults()
}

func (r *nothingRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *nothingRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *nothingRecorder) isNil() bool {
	return r.r == nil
}

type exportedNothingAdaptor struct{ m *exported.MockUsual }

func (a *exportedNothingAdaptor) exported() bool { return true }

func (a *exportedNothingAdaptor) tracksParams() bool { return false }

func (a *exportedNothingAdaptor) newRecorder([]string, bool) recorder {
	return &exportedNothingRecorder{r: a.m.OnCall().Nothing()}
}

func (a *exportedNothingAdaptor) invokeMockAndExpectResults([]string, bool, results) {
	a.m.Mock().Nothing()
}

func (a *exportedNothingAdaptor) bundleParams([]string, bool) interface{} {
	return exported.MockUsual_Nothing_params{}
}

func (a *exportedNothingAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedNothingRecorder struct {
	r *exported.MockUsual_Nothing_fnRecorder
}

func (r *exportedNothingRecorder) anySParam() {}

func (r *exportedNothingRecorder) anyBParam() {}

func (r *exportedNothingRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedNothingRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedNothingRecorder) returnResults([]string, error) {
	r.r = r.r.ReturnResults()
}

func (r *exportedNothingRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedNothingRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedNothingRecorder) isNil() bool {
	return r.r == nil
}

type variadicAdaptor struct{ m *mockUsual }

func (a *variadicAdaptor) exported() bool { return false }

func (a *variadicAdaptor) tracksParams() bool { return true }

func (a *variadicAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &variadicRecorder{r: a.m.onCall().Variadic(bParam, sParams...)}
}

func (a *variadicAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.mock().Variadic(bParam, sParams...)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *variadicAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return mockUsual_Variadic_params{args: sParams, other: bParam}
}

func (a *variadicAdaptor) sceneMock() moq.Mock {
	return a.m
}

type variadicRecorder struct {
	r *mockUsual_Variadic_fnRecorder
}

func (r *variadicRecorder) anySParam() {
	r.r = r.r.anyArgs()
}

func (r *variadicRecorder) anyBParam() {
	r.r = r.r.anyOther()
}

func (r *variadicRecorder) seq() {
	r.r = r.r.seq()
}

func (r *variadicRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *variadicRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.returnResults(sResults[0], err)
}

func (r *variadicRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *variadicRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *variadicRecorder) isNil() bool {
	return r.r == nil
}

type exportedVariadicAdaptor struct{ m *exported.MockUsual }

func (a *exportedVariadicAdaptor) exported() bool { return true }

func (a *exportedVariadicAdaptor) tracksParams() bool { return true }

func (a *exportedVariadicAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &exportedVariadicRecorder{r: a.m.OnCall().Variadic(bParam, sParams...)}
}

func (a *exportedVariadicAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult, err := a.m.Mock().Variadic(bParam, sParams...)
	Expect(sResult).To(Equal(res.sResults[0]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *exportedVariadicAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return exported.MockUsual_Variadic_params{Args: sParams, Other: bParam}
}

func (a *exportedVariadicAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedVariadicRecorder struct {
	r *exported.MockUsual_Variadic_fnRecorder
}

func (r *exportedVariadicRecorder) anySParam() {
	r.r = r.r.AnyArgs()
}

func (r *exportedVariadicRecorder) anyBParam() {
	r.r = r.r.AnyOther()
}

func (r *exportedVariadicRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedVariadicRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedVariadicRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.ReturnResults(sResults[0], err)
}

func (r *exportedVariadicRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedVariadicRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedVariadicRecorder) isNil() bool {
	return r.r == nil
}

type repeatedIdsAdaptor struct{ m *mockUsual }

func (a *repeatedIdsAdaptor) exported() bool { return false }

func (a *repeatedIdsAdaptor) tracksParams() bool { return true }

func (a *repeatedIdsAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &repeatedIdsRecorder{r: a.m.onCall().RepeatedIds(sParams[0], sParams[1], bParam)}
}

func (a *repeatedIdsAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult1, sResult2, err := a.m.mock().RepeatedIds(sParams[0], sParams[1], bParam)
	Expect(sResult1).To(Equal(res.sResults[0]))
	Expect(sResult2).To(Equal(res.sResults[1]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *repeatedIdsAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return mockUsual_RepeatedIds_params{sParam1: sParams[0], sParam2: sParams[1], bParam: bParam}
}

func (a *repeatedIdsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type repeatedIdsRecorder struct {
	r *mockUsual_RepeatedIds_fnRecorder
}

func (r *repeatedIdsRecorder) anySParam() {
	r.r = r.r.anySParam1()
}

func (r *repeatedIdsRecorder) anyBParam() {
	r.r = r.r.anyBParam()
}

func (r *repeatedIdsRecorder) seq() {
	r.r = r.r.seq()
}

func (r *repeatedIdsRecorder) noSeq() {
	r.r = r.r.noSeq()
}

func (r *repeatedIdsRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.returnResults(sResults[0], sResults[1], err)
}

func (r *repeatedIdsRecorder) times(count int) {
	r.r = r.r.times(count)
}

func (r *repeatedIdsRecorder) anyTimes() {
	r.r.anyTimes()
	r.r = nil
}

func (r *repeatedIdsRecorder) isNil() bool {
	return r.r == nil
}

type exportedRepeatedIdsAdaptor struct{ m *exported.MockUsual }

func (a *exportedRepeatedIdsAdaptor) exported() bool { return true }

func (a *exportedRepeatedIdsAdaptor) tracksParams() bool { return true }

func (a *exportedRepeatedIdsAdaptor) newRecorder(sParams []string, bParam bool) recorder {
	return &exportedRepeatedIdsRecorder{r: a.m.OnCall().RepeatedIds(sParams[0], sParams[1], bParam)}
}

func (a *exportedRepeatedIdsAdaptor) invokeMockAndExpectResults(sParams []string, bParam bool, res results) {
	sResult1, sResult2, err := a.m.Mock().RepeatedIds(sParams[0], sParams[1], bParam)
	Expect(sResult1).To(Equal(res.sResults[0]))
	Expect(sResult2).To(Equal(res.sResults[1]))
	if res.err == nil {
		Expect(err).To(BeNil())
	} else {
		Expect(err).To(Equal(res.err))
	}
}

func (a *exportedRepeatedIdsAdaptor) bundleParams(sParams []string, bParam bool) interface{} {
	return exported.MockUsual_RepeatedIds_params{SParam1: sParams[0], SParam2: sParams[1], BParam: bParam}
}

func (a *exportedRepeatedIdsAdaptor) sceneMock() moq.Mock {
	return a.m
}

type exportedRepeatedIdsRecorder struct {
	r *exported.MockUsual_RepeatedIds_fnRecorder
}

func (r *exportedRepeatedIdsRecorder) anySParam() {
	r.r = r.r.AnySParam1()
}

func (r *exportedRepeatedIdsRecorder) anyBParam() {
	r.r = r.r.AnyBParam()
}

func (r *exportedRepeatedIdsRecorder) seq() {
	r.r = r.r.Seq()
}

func (r *exportedRepeatedIdsRecorder) noSeq() {
	r.r = r.r.NoSeq()
}

func (r *exportedRepeatedIdsRecorder) returnResults(sResults []string, err error) {
	r.r = r.r.ReturnResults(sResults[0], sResults[1], err)
}

func (r *exportedRepeatedIdsRecorder) times(count int) {
	r.r = r.r.Times(count)
}

func (r *exportedRepeatedIdsRecorder) anyTimes() {
	r.r.AnyTimes()
	r.r = nil
}

func (r *exportedRepeatedIdsRecorder) isNil() bool {
	return r.r == nil
}
