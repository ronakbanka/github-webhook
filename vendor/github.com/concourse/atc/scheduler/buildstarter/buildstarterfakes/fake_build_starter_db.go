// This file was generated by counterfeiter
package buildstarterfakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/scheduler/buildstarter"
)

type FakeBuildStarterDB struct {
	GetNextBuildInputsStub        func(jobName string) ([]db.BuildInput, bool, error)
	getNextBuildInputsMutex       sync.RWMutex
	getNextBuildInputsArgsForCall []struct {
		jobName string
	}
	getNextBuildInputsReturns struct {
		result1 []db.BuildInput
		result2 bool
		result3 error
	}
	IsPausedStub        func() (bool, error)
	isPausedMutex       sync.RWMutex
	isPausedArgsForCall []struct{}
	isPausedReturns     struct {
		result1 bool
		result2 error
	}
	GetJobStub        func(job string) (db.SavedJob, bool, error)
	getJobMutex       sync.RWMutex
	getJobArgsForCall []struct {
		job string
	}
	getJobReturns struct {
		result1 db.SavedJob
		result2 bool
		result3 error
	}
	UpdateBuildToScheduledStub        func(int) (bool, error)
	updateBuildToScheduledMutex       sync.RWMutex
	updateBuildToScheduledArgsForCall []struct {
		arg1 int
	}
	updateBuildToScheduledReturns struct {
		result1 bool
		result2 error
	}
	UseInputsForBuildStub        func(buildID int, inputs []db.BuildInput) error
	useInputsForBuildMutex       sync.RWMutex
	useInputsForBuildArgsForCall []struct {
		buildID int
		inputs  []db.BuildInput
	}
	useInputsForBuildReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuildStarterDB) GetNextBuildInputs(jobName string) ([]db.BuildInput, bool, error) {
	fake.getNextBuildInputsMutex.Lock()
	fake.getNextBuildInputsArgsForCall = append(fake.getNextBuildInputsArgsForCall, struct {
		jobName string
	}{jobName})
	fake.recordInvocation("GetNextBuildInputs", []interface{}{jobName})
	fake.getNextBuildInputsMutex.Unlock()
	if fake.GetNextBuildInputsStub != nil {
		return fake.GetNextBuildInputsStub(jobName)
	} else {
		return fake.getNextBuildInputsReturns.result1, fake.getNextBuildInputsReturns.result2, fake.getNextBuildInputsReturns.result3
	}
}

func (fake *FakeBuildStarterDB) GetNextBuildInputsCallCount() int {
	fake.getNextBuildInputsMutex.RLock()
	defer fake.getNextBuildInputsMutex.RUnlock()
	return len(fake.getNextBuildInputsArgsForCall)
}

func (fake *FakeBuildStarterDB) GetNextBuildInputsArgsForCall(i int) string {
	fake.getNextBuildInputsMutex.RLock()
	defer fake.getNextBuildInputsMutex.RUnlock()
	return fake.getNextBuildInputsArgsForCall[i].jobName
}

func (fake *FakeBuildStarterDB) GetNextBuildInputsReturns(result1 []db.BuildInput, result2 bool, result3 error) {
	fake.GetNextBuildInputsStub = nil
	fake.getNextBuildInputsReturns = struct {
		result1 []db.BuildInput
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildStarterDB) IsPaused() (bool, error) {
	fake.isPausedMutex.Lock()
	fake.isPausedArgsForCall = append(fake.isPausedArgsForCall, struct{}{})
	fake.recordInvocation("IsPaused", []interface{}{})
	fake.isPausedMutex.Unlock()
	if fake.IsPausedStub != nil {
		return fake.IsPausedStub()
	} else {
		return fake.isPausedReturns.result1, fake.isPausedReturns.result2
	}
}

func (fake *FakeBuildStarterDB) IsPausedCallCount() int {
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	return len(fake.isPausedArgsForCall)
}

func (fake *FakeBuildStarterDB) IsPausedReturns(result1 bool, result2 error) {
	fake.IsPausedStub = nil
	fake.isPausedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStarterDB) GetJob(job string) (db.SavedJob, bool, error) {
	fake.getJobMutex.Lock()
	fake.getJobArgsForCall = append(fake.getJobArgsForCall, struct {
		job string
	}{job})
	fake.recordInvocation("GetJob", []interface{}{job})
	fake.getJobMutex.Unlock()
	if fake.GetJobStub != nil {
		return fake.GetJobStub(job)
	} else {
		return fake.getJobReturns.result1, fake.getJobReturns.result2, fake.getJobReturns.result3
	}
}

func (fake *FakeBuildStarterDB) GetJobCallCount() int {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return len(fake.getJobArgsForCall)
}

func (fake *FakeBuildStarterDB) GetJobArgsForCall(i int) string {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return fake.getJobArgsForCall[i].job
}

func (fake *FakeBuildStarterDB) GetJobReturns(result1 db.SavedJob, result2 bool, result3 error) {
	fake.GetJobStub = nil
	fake.getJobReturns = struct {
		result1 db.SavedJob
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuildStarterDB) UpdateBuildToScheduled(arg1 int) (bool, error) {
	fake.updateBuildToScheduledMutex.Lock()
	fake.updateBuildToScheduledArgsForCall = append(fake.updateBuildToScheduledArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("UpdateBuildToScheduled", []interface{}{arg1})
	fake.updateBuildToScheduledMutex.Unlock()
	if fake.UpdateBuildToScheduledStub != nil {
		return fake.UpdateBuildToScheduledStub(arg1)
	} else {
		return fake.updateBuildToScheduledReturns.result1, fake.updateBuildToScheduledReturns.result2
	}
}

func (fake *FakeBuildStarterDB) UpdateBuildToScheduledCallCount() int {
	fake.updateBuildToScheduledMutex.RLock()
	defer fake.updateBuildToScheduledMutex.RUnlock()
	return len(fake.updateBuildToScheduledArgsForCall)
}

func (fake *FakeBuildStarterDB) UpdateBuildToScheduledArgsForCall(i int) int {
	fake.updateBuildToScheduledMutex.RLock()
	defer fake.updateBuildToScheduledMutex.RUnlock()
	return fake.updateBuildToScheduledArgsForCall[i].arg1
}

func (fake *FakeBuildStarterDB) UpdateBuildToScheduledReturns(result1 bool, result2 error) {
	fake.UpdateBuildToScheduledStub = nil
	fake.updateBuildToScheduledReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBuildStarterDB) UseInputsForBuild(buildID int, inputs []db.BuildInput) error {
	var inputsCopy []db.BuildInput
	if inputs != nil {
		inputsCopy = make([]db.BuildInput, len(inputs))
		copy(inputsCopy, inputs)
	}
	fake.useInputsForBuildMutex.Lock()
	fake.useInputsForBuildArgsForCall = append(fake.useInputsForBuildArgsForCall, struct {
		buildID int
		inputs  []db.BuildInput
	}{buildID, inputsCopy})
	fake.recordInvocation("UseInputsForBuild", []interface{}{buildID, inputsCopy})
	fake.useInputsForBuildMutex.Unlock()
	if fake.UseInputsForBuildStub != nil {
		return fake.UseInputsForBuildStub(buildID, inputs)
	} else {
		return fake.useInputsForBuildReturns.result1
	}
}

func (fake *FakeBuildStarterDB) UseInputsForBuildCallCount() int {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return len(fake.useInputsForBuildArgsForCall)
}

func (fake *FakeBuildStarterDB) UseInputsForBuildArgsForCall(i int) (int, []db.BuildInput) {
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return fake.useInputsForBuildArgsForCall[i].buildID, fake.useInputsForBuildArgsForCall[i].inputs
}

func (fake *FakeBuildStarterDB) UseInputsForBuildReturns(result1 error) {
	fake.UseInputsForBuildStub = nil
	fake.useInputsForBuildReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuildStarterDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getNextBuildInputsMutex.RLock()
	defer fake.getNextBuildInputsMutex.RUnlock()
	fake.isPausedMutex.RLock()
	defer fake.isPausedMutex.RUnlock()
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	fake.updateBuildToScheduledMutex.RLock()
	defer fake.updateBuildToScheduledMutex.RUnlock()
	fake.useInputsForBuildMutex.RLock()
	defer fake.useInputsForBuildMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBuildStarterDB) recordInvocation(key string, args []interface{}) {
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

var _ buildstarter.BuildStarterDB = new(FakeBuildStarterDB)
