// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakePipelineDBFactory struct {
	BuildStub        func(pipeline db.SavedPipeline) db.PipelineDB
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		pipeline db.SavedPipeline
	}
	buildReturns struct {
		result1 db.PipelineDB
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelineDBFactory) Build(pipeline db.SavedPipeline) db.PipelineDB {
	fake.buildMutex.Lock()
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		pipeline db.SavedPipeline
	}{pipeline})
	fake.recordInvocation("Build", []interface{}{pipeline})
	fake.buildMutex.Unlock()
	if fake.BuildStub != nil {
		return fake.BuildStub(pipeline)
	} else {
		return fake.buildReturns.result1
	}
}

func (fake *FakePipelineDBFactory) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakePipelineDBFactory) BuildArgsForCall(i int) db.SavedPipeline {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.buildArgsForCall[i].pipeline
}

func (fake *FakePipelineDBFactory) BuildReturns(result1 db.PipelineDB) {
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 db.PipelineDB
	}{result1}
}

func (fake *FakePipelineDBFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePipelineDBFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.PipelineDBFactory = new(FakePipelineDBFactory)
