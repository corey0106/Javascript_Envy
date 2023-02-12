package safe

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// BoxMock implements Box
type BoxMock struct {
	t minimock.Tester

	funcGet          func(s1 string) (np1 *Namespace, err error)
	inspectFuncGet   func(s1 string)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mBoxMockGet

	funcList          func() (sa1 []string, err error)
	inspectFuncList   func()
	afterListCounter  uint64
	beforeListCounter uint64
	ListMock          mBoxMockList

	funcPurge          func(s1 string) (err error)
	inspectFuncPurge   func(s1 string)
	afterPurgeCounter  uint64
	beforePurgeCounter uint64
	PurgeMock          mBoxMockPurge

	funcSet          func(np1 *Namespace) (err error)
	inspectFuncSet   func(np1 *Namespace)
	afterSetCounter  uint64
	beforeSetCounter uint64
	SetMock          mBoxMockSet

	funcUpdate          func(np1 *Namespace) (err error)
	inspectFuncUpdate   func(np1 *Namespace)
	afterUpdateCounter  uint64
	beforeUpdateCounter uint64
	UpdateMock          mBoxMockUpdate
}

// NewBoxMock returns a mock for Box
func NewBoxMock(t minimock.Tester) *BoxMock {
	m := &BoxMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetMock = mBoxMockGet{mock: m}
	m.GetMock.callArgs = []*BoxMockGetParams{}

	m.ListMock = mBoxMockList{mock: m}

	m.PurgeMock = mBoxMockPurge{mock: m}
	m.PurgeMock.callArgs = []*BoxMockPurgeParams{}

	m.SetMock = mBoxMockSet{mock: m}
	m.SetMock.callArgs = []*BoxMockSetParams{}

	m.UpdateMock = mBoxMockUpdate{mock: m}
	m.UpdateMock.callArgs = []*BoxMockUpdateParams{}

	return m
}

type mBoxMockGet struct {
	mock               *BoxMock
	defaultExpectation *BoxMockGetExpectation
	expectations       []*BoxMockGetExpectation

	callArgs []*BoxMockGetParams
	mutex    sync.RWMutex
}

// BoxMockGetExpectation specifies expectation struct of the Box.Get
type BoxMockGetExpectation struct {
	mock    *BoxMock
	params  *BoxMockGetParams
	results *BoxMockGetResults
	Counter uint64
}

// BoxMockGetParams contains parameters of the Box.Get
type BoxMockGetParams struct {
	s1 string
}

// BoxMockGetResults contains results of the Box.Get
type BoxMockGetResults struct {
	np1 *Namespace
	err error
}

// Expect sets up expected params for Box.Get
func (mmGet *mBoxMockGet) Expect(s1 string) *mBoxMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BoxMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &BoxMockGetParams{s1}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the Box.Get
func (mmGet *mBoxMockGet) Inspect(f func(s1 string)) *mBoxMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for BoxMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by Box.Get
func (mmGet *mBoxMockGet) Return(np1 *Namespace, err error) *BoxMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &BoxMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &BoxMockGetResults{np1, err}
	return mmGet.mock
}

// Set uses given function f to mock the Box.Get method
func (mmGet *mBoxMockGet) Set(f func(s1 string) (np1 *Namespace, err error)) *BoxMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the Box.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the Box.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the Box.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mBoxMockGet) When(s1 string) *BoxMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("BoxMock.Get mock is already set by Set")
	}

	expectation := &BoxMockGetExpectation{
		mock:   mmGet.mock,
		params: &BoxMockGetParams{s1},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up Box.Get return parameters for the expectation previously defined by the When method
func (e *BoxMockGetExpectation) Then(np1 *Namespace, err error) *BoxMock {
	e.results = &BoxMockGetResults{np1, err}
	return e.mock
}

// Get implements Box
func (mmGet *BoxMock) Get(s1 string) (np1 *Namespace, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(s1)
	}

	mm_params := &BoxMockGetParams{s1}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.np1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := BoxMockGetParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("BoxMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the BoxMock.Get")
		}
		return (*mm_results).np1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(s1)
	}
	mmGet.t.Fatalf("Unexpected call to BoxMock.Get. %v", s1)
	return
}

// GetAfterCounter returns a count of finished BoxMock.Get invocations
func (mmGet *BoxMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of BoxMock.Get invocations
func (mmGet *BoxMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mBoxMockGet) Calls() []*BoxMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*BoxMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *BoxMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Get")
		} else {
			m.t.Errorf("Expected call to BoxMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Get")
	}
}

type mBoxMockList struct {
	mock               *BoxMock
	defaultExpectation *BoxMockListExpectation
	expectations       []*BoxMockListExpectation
}

// BoxMockListExpectation specifies expectation struct of the Box.List
type BoxMockListExpectation struct {
	mock *BoxMock

	results *BoxMockListResults
	Counter uint64
}

// BoxMockListResults contains results of the Box.List
type BoxMockListResults struct {
	sa1 []string
	err error
}

// Expect sets up expected params for Box.List
func (mmList *mBoxMockList) Expect() *mBoxMockList {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("BoxMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &BoxMockListExpectation{}
	}

	return mmList
}

// Inspect accepts an inspector function that has same arguments as the Box.List
func (mmList *mBoxMockList) Inspect(f func()) *mBoxMockList {
	if mmList.mock.inspectFuncList != nil {
		mmList.mock.t.Fatalf("Inspect function is already set for BoxMock.List")
	}

	mmList.mock.inspectFuncList = f

	return mmList
}

// Return sets up results that will be returned by Box.List
func (mmList *mBoxMockList) Return(sa1 []string, err error) *BoxMock {
	if mmList.mock.funcList != nil {
		mmList.mock.t.Fatalf("BoxMock.List mock is already set by Set")
	}

	if mmList.defaultExpectation == nil {
		mmList.defaultExpectation = &BoxMockListExpectation{mock: mmList.mock}
	}
	mmList.defaultExpectation.results = &BoxMockListResults{sa1, err}
	return mmList.mock
}

// Set uses given function f to mock the Box.List method
func (mmList *mBoxMockList) Set(f func() (sa1 []string, err error)) *BoxMock {
	if mmList.defaultExpectation != nil {
		mmList.mock.t.Fatalf("Default expectation is already set for the Box.List method")
	}

	if len(mmList.expectations) > 0 {
		mmList.mock.t.Fatalf("Some expectations are already set for the Box.List method")
	}

	mmList.mock.funcList = f
	return mmList.mock
}

// List implements Box
func (mmList *BoxMock) List() (sa1 []string, err error) {
	mm_atomic.AddUint64(&mmList.beforeListCounter, 1)
	defer mm_atomic.AddUint64(&mmList.afterListCounter, 1)

	if mmList.inspectFuncList != nil {
		mmList.inspectFuncList()
	}

	if mmList.ListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmList.ListMock.defaultExpectation.Counter, 1)

		mm_results := mmList.ListMock.defaultExpectation.results
		if mm_results == nil {
			mmList.t.Fatal("No results are set for the BoxMock.List")
		}
		return (*mm_results).sa1, (*mm_results).err
	}
	if mmList.funcList != nil {
		return mmList.funcList()
	}
	mmList.t.Fatalf("Unexpected call to BoxMock.List.")
	return
}

// ListAfterCounter returns a count of finished BoxMock.List invocations
func (mmList *BoxMock) ListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.afterListCounter)
}

// ListBeforeCounter returns a count of BoxMock.List invocations
func (mmList *BoxMock) ListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmList.beforeListCounter)
}

// MinimockListDone returns true if the count of the List invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockListDone() bool {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		return false
	}
	return true
}

// MinimockListInspect logs each unmet expectation
func (m *BoxMock) MinimockListInspect() {
	for _, e := range m.ListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to BoxMock.List")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to BoxMock.List")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcList != nil && mm_atomic.LoadUint64(&m.afterListCounter) < 1 {
		m.t.Error("Expected call to BoxMock.List")
	}
}

type mBoxMockPurge struct {
	mock               *BoxMock
	defaultExpectation *BoxMockPurgeExpectation
	expectations       []*BoxMockPurgeExpectation

	callArgs []*BoxMockPurgeParams
	mutex    sync.RWMutex
}

// BoxMockPurgeExpectation specifies expectation struct of the Box.Purge
type BoxMockPurgeExpectation struct {
	mock    *BoxMock
	params  *BoxMockPurgeParams
	results *BoxMockPurgeResults
	Counter uint64
}

// BoxMockPurgeParams contains parameters of the Box.Purge
type BoxMockPurgeParams struct {
	s1 string
}

// BoxMockPurgeResults contains results of the Box.Purge
type BoxMockPurgeResults struct {
	err error
}

// Expect sets up expected params for Box.Purge
func (mmPurge *mBoxMockPurge) Expect(s1 string) *mBoxMockPurge {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	if mmPurge.defaultExpectation == nil {
		mmPurge.defaultExpectation = &BoxMockPurgeExpectation{}
	}

	mmPurge.defaultExpectation.params = &BoxMockPurgeParams{s1}
	for _, e := range mmPurge.expectations {
		if minimock.Equal(e.params, mmPurge.defaultExpectation.params) {
			mmPurge.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmPurge.defaultExpectation.params)
		}
	}

	return mmPurge
}

// Inspect accepts an inspector function that has same arguments as the Box.Purge
func (mmPurge *mBoxMockPurge) Inspect(f func(s1 string)) *mBoxMockPurge {
	if mmPurge.mock.inspectFuncPurge != nil {
		mmPurge.mock.t.Fatalf("Inspect function is already set for BoxMock.Purge")
	}

	mmPurge.mock.inspectFuncPurge = f

	return mmPurge
}

// Return sets up results that will be returned by Box.Purge
func (mmPurge *mBoxMockPurge) Return(err error) *BoxMock {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	if mmPurge.defaultExpectation == nil {
		mmPurge.defaultExpectation = &BoxMockPurgeExpectation{mock: mmPurge.mock}
	}
	mmPurge.defaultExpectation.results = &BoxMockPurgeResults{err}
	return mmPurge.mock
}

// Set uses given function f to mock the Box.Purge method
func (mmPurge *mBoxMockPurge) Set(f func(s1 string) (err error)) *BoxMock {
	if mmPurge.defaultExpectation != nil {
		mmPurge.mock.t.Fatalf("Default expectation is already set for the Box.Purge method")
	}

	if len(mmPurge.expectations) > 0 {
		mmPurge.mock.t.Fatalf("Some expectations are already set for the Box.Purge method")
	}

	mmPurge.mock.funcPurge = f
	return mmPurge.mock
}

// When sets expectation for the Box.Purge which will trigger the result defined by the following
// Then helper
func (mmPurge *mBoxMockPurge) When(s1 string) *BoxMockPurgeExpectation {
	if mmPurge.mock.funcPurge != nil {
		mmPurge.mock.t.Fatalf("BoxMock.Purge mock is already set by Set")
	}

	expectation := &BoxMockPurgeExpectation{
		mock:   mmPurge.mock,
		params: &BoxMockPurgeParams{s1},
	}
	mmPurge.expectations = append(mmPurge.expectations, expectation)
	return expectation
}

// Then sets up Box.Purge return parameters for the expectation previously defined by the When method
func (e *BoxMockPurgeExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockPurgeResults{err}
	return e.mock
}

// Purge implements Box
func (mmPurge *BoxMock) Purge(s1 string) (err error) {
	mm_atomic.AddUint64(&mmPurge.beforePurgeCounter, 1)
	defer mm_atomic.AddUint64(&mmPurge.afterPurgeCounter, 1)

	if mmPurge.inspectFuncPurge != nil {
		mmPurge.inspectFuncPurge(s1)
	}

	mm_params := &BoxMockPurgeParams{s1}

	// Record call args
	mmPurge.PurgeMock.mutex.Lock()
	mmPurge.PurgeMock.callArgs = append(mmPurge.PurgeMock.callArgs, mm_params)
	mmPurge.PurgeMock.mutex.Unlock()

	for _, e := range mmPurge.PurgeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmPurge.PurgeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmPurge.PurgeMock.defaultExpectation.Counter, 1)
		mm_want := mmPurge.PurgeMock.defaultExpectation.params
		mm_got := BoxMockPurgeParams{s1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmPurge.t.Errorf("BoxMock.Purge got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmPurge.PurgeMock.defaultExpectation.results
		if mm_results == nil {
			mmPurge.t.Fatal("No results are set for the BoxMock.Purge")
		}
		return (*mm_results).err
	}
	if mmPurge.funcPurge != nil {
		return mmPurge.funcPurge(s1)
	}
	mmPurge.t.Fatalf("Unexpected call to BoxMock.Purge. %v", s1)
	return
}

// PurgeAfterCounter returns a count of finished BoxMock.Purge invocations
func (mmPurge *BoxMock) PurgeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurge.afterPurgeCounter)
}

// PurgeBeforeCounter returns a count of BoxMock.Purge invocations
func (mmPurge *BoxMock) PurgeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmPurge.beforePurgeCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Purge.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmPurge *mBoxMockPurge) Calls() []*BoxMockPurgeParams {
	mmPurge.mutex.RLock()

	argCopy := make([]*BoxMockPurgeParams, len(mmPurge.callArgs))
	copy(argCopy, mmPurge.callArgs)

	mmPurge.mutex.RUnlock()

	return argCopy
}

// MinimockPurgeDone returns true if the count of the Purge invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockPurgeDone() bool {
	for _, e := range m.PurgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurge != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		return false
	}
	return true
}

// MinimockPurgeInspect logs each unmet expectation
func (m *BoxMock) MinimockPurgeInspect() {
	for _, e := range m.PurgeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Purge with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.PurgeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		if m.PurgeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Purge")
		} else {
			m.t.Errorf("Expected call to BoxMock.Purge with params: %#v", *m.PurgeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcPurge != nil && mm_atomic.LoadUint64(&m.afterPurgeCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Purge")
	}
}

type mBoxMockSet struct {
	mock               *BoxMock
	defaultExpectation *BoxMockSetExpectation
	expectations       []*BoxMockSetExpectation

	callArgs []*BoxMockSetParams
	mutex    sync.RWMutex
}

// BoxMockSetExpectation specifies expectation struct of the Box.Set
type BoxMockSetExpectation struct {
	mock    *BoxMock
	params  *BoxMockSetParams
	results *BoxMockSetResults
	Counter uint64
}

// BoxMockSetParams contains parameters of the Box.Set
type BoxMockSetParams struct {
	np1 *Namespace
}

// BoxMockSetResults contains results of the Box.Set
type BoxMockSetResults struct {
	err error
}

// Expect sets up expected params for Box.Set
func (mmSet *mBoxMockSet) Expect(np1 *Namespace) *mBoxMockSet {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &BoxMockSetExpectation{}
	}

	mmSet.defaultExpectation.params = &BoxMockSetParams{np1}
	for _, e := range mmSet.expectations {
		if minimock.Equal(e.params, mmSet.defaultExpectation.params) {
			mmSet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSet.defaultExpectation.params)
		}
	}

	return mmSet
}

// Inspect accepts an inspector function that has same arguments as the Box.Set
func (mmSet *mBoxMockSet) Inspect(f func(np1 *Namespace)) *mBoxMockSet {
	if mmSet.mock.inspectFuncSet != nil {
		mmSet.mock.t.Fatalf("Inspect function is already set for BoxMock.Set")
	}

	mmSet.mock.inspectFuncSet = f

	return mmSet
}

// Return sets up results that will be returned by Box.Set
func (mmSet *mBoxMockSet) Return(err error) *BoxMock {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	if mmSet.defaultExpectation == nil {
		mmSet.defaultExpectation = &BoxMockSetExpectation{mock: mmSet.mock}
	}
	mmSet.defaultExpectation.results = &BoxMockSetResults{err}
	return mmSet.mock
}

// Set uses given function f to mock the Box.Set method
func (mmSet *mBoxMockSet) Set(f func(np1 *Namespace) (err error)) *BoxMock {
	if mmSet.defaultExpectation != nil {
		mmSet.mock.t.Fatalf("Default expectation is already set for the Box.Set method")
	}

	if len(mmSet.expectations) > 0 {
		mmSet.mock.t.Fatalf("Some expectations are already set for the Box.Set method")
	}

	mmSet.mock.funcSet = f
	return mmSet.mock
}

// When sets expectation for the Box.Set which will trigger the result defined by the following
// Then helper
func (mmSet *mBoxMockSet) When(np1 *Namespace) *BoxMockSetExpectation {
	if mmSet.mock.funcSet != nil {
		mmSet.mock.t.Fatalf("BoxMock.Set mock is already set by Set")
	}

	expectation := &BoxMockSetExpectation{
		mock:   mmSet.mock,
		params: &BoxMockSetParams{np1},
	}
	mmSet.expectations = append(mmSet.expectations, expectation)
	return expectation
}

// Then sets up Box.Set return parameters for the expectation previously defined by the When method
func (e *BoxMockSetExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockSetResults{err}
	return e.mock
}

// Set implements Box
func (mmSet *BoxMock) Set(np1 *Namespace) (err error) {
	mm_atomic.AddUint64(&mmSet.beforeSetCounter, 1)
	defer mm_atomic.AddUint64(&mmSet.afterSetCounter, 1)

	if mmSet.inspectFuncSet != nil {
		mmSet.inspectFuncSet(np1)
	}

	mm_params := &BoxMockSetParams{np1}

	// Record call args
	mmSet.SetMock.mutex.Lock()
	mmSet.SetMock.callArgs = append(mmSet.SetMock.callArgs, mm_params)
	mmSet.SetMock.mutex.Unlock()

	for _, e := range mmSet.SetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSet.SetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSet.SetMock.defaultExpectation.Counter, 1)
		mm_want := mmSet.SetMock.defaultExpectation.params
		mm_got := BoxMockSetParams{np1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSet.t.Errorf("BoxMock.Set got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSet.SetMock.defaultExpectation.results
		if mm_results == nil {
			mmSet.t.Fatal("No results are set for the BoxMock.Set")
		}
		return (*mm_results).err
	}
	if mmSet.funcSet != nil {
		return mmSet.funcSet(np1)
	}
	mmSet.t.Fatalf("Unexpected call to BoxMock.Set. %v", np1)
	return
}

// SetAfterCounter returns a count of finished BoxMock.Set invocations
func (mmSet *BoxMock) SetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.afterSetCounter)
}

// SetBeforeCounter returns a count of BoxMock.Set invocations
func (mmSet *BoxMock) SetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSet.beforeSetCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Set.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSet *mBoxMockSet) Calls() []*BoxMockSetParams {
	mmSet.mutex.RLock()

	argCopy := make([]*BoxMockSetParams, len(mmSet.callArgs))
	copy(argCopy, mmSet.callArgs)

	mmSet.mutex.RUnlock()

	return argCopy
}

// MinimockSetDone returns true if the count of the Set invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockSetDone() bool {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetInspect logs each unmet expectation
func (m *BoxMock) MinimockSetInspect() {
	for _, e := range m.SetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Set with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		if m.SetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Set")
		} else {
			m.t.Errorf("Expected call to BoxMock.Set with params: %#v", *m.SetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSet != nil && mm_atomic.LoadUint64(&m.afterSetCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Set")
	}
}

type mBoxMockUpdate struct {
	mock               *BoxMock
	defaultExpectation *BoxMockUpdateExpectation
	expectations       []*BoxMockUpdateExpectation

	callArgs []*BoxMockUpdateParams
	mutex    sync.RWMutex
}

// BoxMockUpdateExpectation specifies expectation struct of the Box.Update
type BoxMockUpdateExpectation struct {
	mock    *BoxMock
	params  *BoxMockUpdateParams
	results *BoxMockUpdateResults
	Counter uint64
}

// BoxMockUpdateParams contains parameters of the Box.Update
type BoxMockUpdateParams struct {
	np1 *Namespace
}

// BoxMockUpdateResults contains results of the Box.Update
type BoxMockUpdateResults struct {
	err error
}

// Expect sets up expected params for Box.Update
func (mmUpdate *mBoxMockUpdate) Expect(np1 *Namespace) *mBoxMockUpdate {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("BoxMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &BoxMockUpdateExpectation{}
	}

	mmUpdate.defaultExpectation.params = &BoxMockUpdateParams{np1}
	for _, e := range mmUpdate.expectations {
		if minimock.Equal(e.params, mmUpdate.defaultExpectation.params) {
			mmUpdate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUpdate.defaultExpectation.params)
		}
	}

	return mmUpdate
}

// Inspect accepts an inspector function that has same arguments as the Box.Update
func (mmUpdate *mBoxMockUpdate) Inspect(f func(np1 *Namespace)) *mBoxMockUpdate {
	if mmUpdate.mock.inspectFuncUpdate != nil {
		mmUpdate.mock.t.Fatalf("Inspect function is already set for BoxMock.Update")
	}

	mmUpdate.mock.inspectFuncUpdate = f

	return mmUpdate
}

// Return sets up results that will be returned by Box.Update
func (mmUpdate *mBoxMockUpdate) Return(err error) *BoxMock {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("BoxMock.Update mock is already set by Set")
	}

	if mmUpdate.defaultExpectation == nil {
		mmUpdate.defaultExpectation = &BoxMockUpdateExpectation{mock: mmUpdate.mock}
	}
	mmUpdate.defaultExpectation.results = &BoxMockUpdateResults{err}
	return mmUpdate.mock
}

// Set uses given function f to mock the Box.Update method
func (mmUpdate *mBoxMockUpdate) Set(f func(np1 *Namespace) (err error)) *BoxMock {
	if mmUpdate.defaultExpectation != nil {
		mmUpdate.mock.t.Fatalf("Default expectation is already set for the Box.Update method")
	}

	if len(mmUpdate.expectations) > 0 {
		mmUpdate.mock.t.Fatalf("Some expectations are already set for the Box.Update method")
	}

	mmUpdate.mock.funcUpdate = f
	return mmUpdate.mock
}

// When sets expectation for the Box.Update which will trigger the result defined by the following
// Then helper
func (mmUpdate *mBoxMockUpdate) When(np1 *Namespace) *BoxMockUpdateExpectation {
	if mmUpdate.mock.funcUpdate != nil {
		mmUpdate.mock.t.Fatalf("BoxMock.Update mock is already set by Set")
	}

	expectation := &BoxMockUpdateExpectation{
		mock:   mmUpdate.mock,
		params: &BoxMockUpdateParams{np1},
	}
	mmUpdate.expectations = append(mmUpdate.expectations, expectation)
	return expectation
}

// Then sets up Box.Update return parameters for the expectation previously defined by the When method
func (e *BoxMockUpdateExpectation) Then(err error) *BoxMock {
	e.results = &BoxMockUpdateResults{err}
	return e.mock
}

// Update implements Box
func (mmUpdate *BoxMock) Update(np1 *Namespace) (err error) {
	mm_atomic.AddUint64(&mmUpdate.beforeUpdateCounter, 1)
	defer mm_atomic.AddUint64(&mmUpdate.afterUpdateCounter, 1)

	if mmUpdate.inspectFuncUpdate != nil {
		mmUpdate.inspectFuncUpdate(np1)
	}

	mm_params := &BoxMockUpdateParams{np1}

	// Record call args
	mmUpdate.UpdateMock.mutex.Lock()
	mmUpdate.UpdateMock.callArgs = append(mmUpdate.UpdateMock.callArgs, mm_params)
	mmUpdate.UpdateMock.mutex.Unlock()

	for _, e := range mmUpdate.UpdateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmUpdate.UpdateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUpdate.UpdateMock.defaultExpectation.Counter, 1)
		mm_want := mmUpdate.UpdateMock.defaultExpectation.params
		mm_got := BoxMockUpdateParams{np1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUpdate.t.Errorf("BoxMock.Update got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmUpdate.UpdateMock.defaultExpectation.results
		if mm_results == nil {
			mmUpdate.t.Fatal("No results are set for the BoxMock.Update")
		}
		return (*mm_results).err
	}
	if mmUpdate.funcUpdate != nil {
		return mmUpdate.funcUpdate(np1)
	}
	mmUpdate.t.Fatalf("Unexpected call to BoxMock.Update. %v", np1)
	return
}

// UpdateAfterCounter returns a count of finished BoxMock.Update invocations
func (mmUpdate *BoxMock) UpdateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.afterUpdateCounter)
}

// UpdateBeforeCounter returns a count of BoxMock.Update invocations
func (mmUpdate *BoxMock) UpdateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUpdate.beforeUpdateCounter)
}

// Calls returns a list of arguments used in each call to BoxMock.Update.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUpdate *mBoxMockUpdate) Calls() []*BoxMockUpdateParams {
	mmUpdate.mutex.RLock()

	argCopy := make([]*BoxMockUpdateParams, len(mmUpdate.callArgs))
	copy(argCopy, mmUpdate.callArgs)

	mmUpdate.mutex.RUnlock()

	return argCopy
}

// MinimockUpdateDone returns true if the count of the Update invocations corresponds
// the number of defined expectations
func (m *BoxMock) MinimockUpdateDone() bool {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		return false
	}
	return true
}

// MinimockUpdateInspect logs each unmet expectation
func (m *BoxMock) MinimockUpdateInspect() {
	for _, e := range m.UpdateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BoxMock.Update with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UpdateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		if m.UpdateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BoxMock.Update")
		} else {
			m.t.Errorf("Expected call to BoxMock.Update with params: %#v", *m.UpdateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUpdate != nil && mm_atomic.LoadUint64(&m.afterUpdateCounter) < 1 {
		m.t.Error("Expected call to BoxMock.Update")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BoxMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetInspect()

		m.MinimockListInspect()

		m.MinimockPurgeInspect()

		m.MinimockSetInspect()

		m.MinimockUpdateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BoxMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *BoxMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetDone() &&
		m.MinimockListDone() &&
		m.MinimockPurgeDone() &&
		m.MinimockSetDone() &&
		m.MinimockUpdateDone()
}
