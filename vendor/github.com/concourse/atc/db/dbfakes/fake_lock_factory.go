// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/db"
)

type FakeLockFactory struct {
	NewLockStub        func(logger lager.Logger, ids db.LockID) db.Lock
	newLockMutex       sync.RWMutex
	newLockArgsForCall []struct {
		logger lager.Logger
		ids    db.LockID
	}
	newLockReturns struct {
		result1 db.Lock
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockFactory) NewLock(logger lager.Logger, ids db.LockID) db.Lock {
	fake.newLockMutex.Lock()
	fake.newLockArgsForCall = append(fake.newLockArgsForCall, struct {
		logger lager.Logger
		ids    db.LockID
	}{logger, ids})
	fake.recordInvocation("NewLock", []interface{}{logger, ids})
	fake.newLockMutex.Unlock()
	if fake.NewLockStub != nil {
		return fake.NewLockStub(logger, ids)
	} else {
		return fake.newLockReturns.result1
	}
}

func (fake *FakeLockFactory) NewLockCallCount() int {
	fake.newLockMutex.RLock()
	defer fake.newLockMutex.RUnlock()
	return len(fake.newLockArgsForCall)
}

func (fake *FakeLockFactory) NewLockArgsForCall(i int) (lager.Logger, db.LockID) {
	fake.newLockMutex.RLock()
	defer fake.newLockMutex.RUnlock()
	return fake.newLockArgsForCall[i].logger, fake.newLockArgsForCall[i].ids
}

func (fake *FakeLockFactory) NewLockReturns(result1 db.Lock) {
	fake.NewLockStub = nil
	fake.newLockReturns = struct {
		result1 db.Lock
	}{result1}
}

func (fake *FakeLockFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newLockMutex.RLock()
	defer fake.newLockMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLockFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.LockFactory = new(FakeLockFactory)
