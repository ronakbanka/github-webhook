// This file was generated by counterfeiter
package fake_translator

import (
	"os"
	"sync"

	"github.com/concourse/baggageclaim/uidjunk"
)

type FakeTranslator struct {
	CacheKeyStub        func() string
	cacheKeyMutex       sync.RWMutex
	cacheKeyArgsForCall []struct{}
	cacheKeyReturns     struct {
		result1 string
	}
	TranslateStub        func(path string, info os.FileInfo, err error) error
	translateMutex       sync.RWMutex
	translateArgsForCall []struct {
		path string
		info os.FileInfo
		err  error
	}
	translateReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTranslator) CacheKey() string {
	fake.cacheKeyMutex.Lock()
	fake.cacheKeyArgsForCall = append(fake.cacheKeyArgsForCall, struct{}{})
	fake.recordInvocation("CacheKey", []interface{}{})
	fake.cacheKeyMutex.Unlock()
	if fake.CacheKeyStub != nil {
		return fake.CacheKeyStub()
	} else {
		return fake.cacheKeyReturns.result1
	}
}

func (fake *FakeTranslator) CacheKeyCallCount() int {
	fake.cacheKeyMutex.RLock()
	defer fake.cacheKeyMutex.RUnlock()
	return len(fake.cacheKeyArgsForCall)
}

func (fake *FakeTranslator) CacheKeyReturns(result1 string) {
	fake.CacheKeyStub = nil
	fake.cacheKeyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTranslator) Translate(path string, info os.FileInfo, err error) error {
	fake.translateMutex.Lock()
	fake.translateArgsForCall = append(fake.translateArgsForCall, struct {
		path string
		info os.FileInfo
		err  error
	}{path, info, err})
	fake.recordInvocation("Translate", []interface{}{path, info, err})
	fake.translateMutex.Unlock()
	if fake.TranslateStub != nil {
		return fake.TranslateStub(path, info, err)
	} else {
		return fake.translateReturns.result1
	}
}

func (fake *FakeTranslator) TranslateCallCount() int {
	fake.translateMutex.RLock()
	defer fake.translateMutex.RUnlock()
	return len(fake.translateArgsForCall)
}

func (fake *FakeTranslator) TranslateArgsForCall(i int) (string, os.FileInfo, error) {
	fake.translateMutex.RLock()
	defer fake.translateMutex.RUnlock()
	return fake.translateArgsForCall[i].path, fake.translateArgsForCall[i].info, fake.translateArgsForCall[i].err
}

func (fake *FakeTranslator) TranslateReturns(result1 error) {
	fake.TranslateStub = nil
	fake.translateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTranslator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cacheKeyMutex.RLock()
	defer fake.cacheKeyMutex.RUnlock()
	fake.translateMutex.RLock()
	defer fake.translateMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTranslator) recordInvocation(key string, args []interface{}) {
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

var _ uidjunk.Translator = new(FakeTranslator)