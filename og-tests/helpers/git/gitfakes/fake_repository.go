// Code generated by counterfeiter. DO NOT EDIT.
package gitfakes

import (
	"og/helpers/git"
	"sync"

	gita "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

type FakeRepository struct {
	CreateRemoteStub        func(*config.RemoteConfig) (*gita.Remote, error)
	createRemoteMutex       sync.RWMutex
	createRemoteArgsForCall []struct {
		arg1 *config.RemoteConfig
	}
	createRemoteReturns struct {
		result1 *gita.Remote
		result2 error
	}
	createRemoteReturnsOnCall map[int]struct {
		result1 *gita.Remote
		result2 error
	}
	LogStub        func(*gita.LogOptions) (object.CommitIter, error)
	logMutex       sync.RWMutex
	logArgsForCall []struct {
		arg1 *gita.LogOptions
	}
	logReturns struct {
		result1 object.CommitIter
		result2 error
	}
	logReturnsOnCall map[int]struct {
		result1 object.CommitIter
		result2 error
	}
	TagsStub        func() (storer.ReferenceIter, error)
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct {
	}
	tagsReturns struct {
		result1 storer.ReferenceIter
		result2 error
	}
	tagsReturnsOnCall map[int]struct {
		result1 storer.ReferenceIter
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) CreateRemote(arg1 *config.RemoteConfig) (*gita.Remote, error) {
	fake.createRemoteMutex.Lock()
	ret, specificReturn := fake.createRemoteReturnsOnCall[len(fake.createRemoteArgsForCall)]
	fake.createRemoteArgsForCall = append(fake.createRemoteArgsForCall, struct {
		arg1 *config.RemoteConfig
	}{arg1})
	fake.recordInvocation("CreateRemote", []interface{}{arg1})
	fake.createRemoteMutex.Unlock()
	if fake.CreateRemoteStub != nil {
		return fake.CreateRemoteStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createRemoteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) CreateRemoteCallCount() int {
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	return len(fake.createRemoteArgsForCall)
}

func (fake *FakeRepository) CreateRemoteCalls(stub func(*config.RemoteConfig) (*gita.Remote, error)) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = stub
}

func (fake *FakeRepository) CreateRemoteArgsForCall(i int) *config.RemoteConfig {
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	argsForCall := fake.createRemoteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) CreateRemoteReturns(result1 *gita.Remote, result2 error) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = nil
	fake.createRemoteReturns = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) CreateRemoteReturnsOnCall(i int, result1 *gita.Remote, result2 error) {
	fake.createRemoteMutex.Lock()
	defer fake.createRemoteMutex.Unlock()
	fake.CreateRemoteStub = nil
	if fake.createRemoteReturnsOnCall == nil {
		fake.createRemoteReturnsOnCall = make(map[int]struct {
			result1 *gita.Remote
			result2 error
		})
	}
	fake.createRemoteReturnsOnCall[i] = struct {
		result1 *gita.Remote
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Log(arg1 *gita.LogOptions) (object.CommitIter, error) {
	fake.logMutex.Lock()
	ret, specificReturn := fake.logReturnsOnCall[len(fake.logArgsForCall)]
	fake.logArgsForCall = append(fake.logArgsForCall, struct {
		arg1 *gita.LogOptions
	}{arg1})
	fake.recordInvocation("Log", []interface{}{arg1})
	fake.logMutex.Unlock()
	if fake.LogStub != nil {
		return fake.LogStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.logReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) LogCallCount() int {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	return len(fake.logArgsForCall)
}

func (fake *FakeRepository) LogCalls(stub func(*gita.LogOptions) (object.CommitIter, error)) {
	fake.logMutex.Lock()
	defer fake.logMutex.Unlock()
	fake.LogStub = stub
}

func (fake *FakeRepository) LogArgsForCall(i int) *gita.LogOptions {
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	argsForCall := fake.logArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) LogReturns(result1 object.CommitIter, result2 error) {
	fake.logMutex.Lock()
	defer fake.logMutex.Unlock()
	fake.LogStub = nil
	fake.logReturns = struct {
		result1 object.CommitIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) LogReturnsOnCall(i int, result1 object.CommitIter, result2 error) {
	fake.logMutex.Lock()
	defer fake.logMutex.Unlock()
	fake.LogStub = nil
	if fake.logReturnsOnCall == nil {
		fake.logReturnsOnCall = make(map[int]struct {
			result1 object.CommitIter
			result2 error
		})
	}
	fake.logReturnsOnCall[i] = struct {
		result1 object.CommitIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Tags() (storer.ReferenceIter, error) {
	fake.tagsMutex.Lock()
	ret, specificReturn := fake.tagsReturnsOnCall[len(fake.tagsArgsForCall)]
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct {
	}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.tagsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeRepository) TagsCalls(stub func() (storer.ReferenceIter, error)) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = stub
}

func (fake *FakeRepository) TagsReturns(result1 storer.ReferenceIter, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) TagsReturnsOnCall(i int, result1 storer.ReferenceIter, result2 error) {
	fake.tagsMutex.Lock()
	defer fake.tagsMutex.Unlock()
	fake.TagsStub = nil
	if fake.tagsReturnsOnCall == nil {
		fake.tagsReturnsOnCall = make(map[int]struct {
			result1 storer.ReferenceIter
			result2 error
		})
	}
	fake.tagsReturnsOnCall[i] = struct {
		result1 storer.ReferenceIter
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createRemoteMutex.RLock()
	defer fake.createRemoteMutex.RUnlock()
	fake.logMutex.RLock()
	defer fake.logMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ git.Repository = new(FakeRepository)
