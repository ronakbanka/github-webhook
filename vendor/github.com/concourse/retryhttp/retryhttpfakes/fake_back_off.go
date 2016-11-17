// This file was generated by counterfeiter
package retryhttpfakes

import (
	"sync"
	"time"

	"github.com/concourse/retryhttp"
)

type FakeBackOff struct {
	NextBackOffStub        func() time.Duration
	nextBackOffMutex       sync.RWMutex
	nextBackOffArgsForCall []struct{}
	nextBackOffReturns     struct {
		result1 time.Duration
	}
	GetElapsedTimeStub        func() time.Duration
	getElapsedTimeMutex       sync.RWMutex
	getElapsedTimeArgsForCall []struct{}
	getElapsedTimeReturns     struct {
		result1 time.Duration
	}
	ResetStub        func()
	resetMutex       sync.RWMutex
	resetArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackOff) NextBackOff() time.Duration {
	fake.nextBackOffMutex.Lock()
	fake.nextBackOffArgsForCall = append(fake.nextBackOffArgsForCall, struct{}{})
	fake.recordInvocation("NextBackOff", []interface{}{})
	fake.nextBackOffMutex.Unlock()
	if fake.NextBackOffStub != nil {
		return fake.NextBackOffStub()
	} else {
		return fake.nextBackOffReturns.result1
	}
}

func (fake *FakeBackOff) NextBackOffCallCount() int {
	fake.nextBackOffMutex.RLock()
	defer fake.nextBackOffMutex.RUnlock()
	return len(fake.nextBackOffArgsForCall)
}

func (fake *FakeBackOff) NextBackOffReturns(result1 time.Duration) {
	fake.NextBackOffStub = nil
	fake.nextBackOffReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeBackOff) GetElapsedTime() time.Duration {
	fake.getElapsedTimeMutex.Lock()
	fake.getElapsedTimeArgsForCall = append(fake.getElapsedTimeArgsForCall, struct{}{})
	fake.recordInvocation("GetElapsedTime", []interface{}{})
	fake.getElapsedTimeMutex.Unlock()
	if fake.GetElapsedTimeStub != nil {
		return fake.GetElapsedTimeStub()
	} else {
		return fake.getElapsedTimeReturns.result1
	}
}

func (fake *FakeBackOff) GetElapsedTimeCallCount() int {
	fake.getElapsedTimeMutex.RLock()
	defer fake.getElapsedTimeMutex.RUnlock()
	return len(fake.getElapsedTimeArgsForCall)
}

func (fake *FakeBackOff) GetElapsedTimeReturns(result1 time.Duration) {
	fake.GetElapsedTimeStub = nil
	fake.getElapsedTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeBackOff) Reset() {
	fake.resetMutex.Lock()
	fake.resetArgsForCall = append(fake.resetArgsForCall, struct{}{})
	fake.recordInvocation("Reset", []interface{}{})
	fake.resetMutex.Unlock()
	if fake.ResetStub != nil {
		fake.ResetStub()
	}
}

func (fake *FakeBackOff) ResetCallCount() int {
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return len(fake.resetArgsForCall)
}

func (fake *FakeBackOff) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nextBackOffMutex.RLock()
	defer fake.nextBackOffMutex.RUnlock()
	fake.getElapsedTimeMutex.RLock()
	defer fake.getElapsedTimeMutex.RUnlock()
	fake.resetMutex.RLock()
	defer fake.resetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBackOff) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ retryhttp.BackOff = new(FakeBackOff)