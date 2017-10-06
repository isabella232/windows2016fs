// Code generated by counterfeiter. DO NOT EDIT.
package hydratorfakes

import (
	"sync"

	"code.cloudfoundry.org/windows2016fs/hydrator"

	"github.com/opencontainers/image-spec/specs-go/v1"
)

type FakeRegistry struct {
	DownloadManifestStub        func(string) (v1.Manifest, error)
	downloadManifestMutex       sync.RWMutex
	downloadManifestArgsForCall []struct {
		arg1 string
	}
	downloadManifestReturns struct {
		result1 v1.Manifest
		result2 error
	}
	downloadManifestReturnsOnCall map[int]struct {
		result1 v1.Manifest
		result2 error
	}
	DownloadLayerStub        func(v1.Descriptor, string) error
	downloadLayerMutex       sync.RWMutex
	downloadLayerArgsForCall []struct {
		arg1 v1.Descriptor
		arg2 string
	}
	downloadLayerReturns struct {
		result1 error
	}
	downloadLayerReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRegistry) DownloadManifest(arg1 string) (v1.Manifest, error) {
	fake.downloadManifestMutex.Lock()
	ret, specificReturn := fake.downloadManifestReturnsOnCall[len(fake.downloadManifestArgsForCall)]
	fake.downloadManifestArgsForCall = append(fake.downloadManifestArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DownloadManifest", []interface{}{arg1})
	fake.downloadManifestMutex.Unlock()
	if fake.DownloadManifestStub != nil {
		return fake.DownloadManifestStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.downloadManifestReturns.result1, fake.downloadManifestReturns.result2
}

func (fake *FakeRegistry) DownloadManifestCallCount() int {
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	return len(fake.downloadManifestArgsForCall)
}

func (fake *FakeRegistry) DownloadManifestArgsForCall(i int) string {
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	return fake.downloadManifestArgsForCall[i].arg1
}

func (fake *FakeRegistry) DownloadManifestReturns(result1 v1.Manifest, result2 error) {
	fake.DownloadManifestStub = nil
	fake.downloadManifestReturns = struct {
		result1 v1.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) DownloadManifestReturnsOnCall(i int, result1 v1.Manifest, result2 error) {
	fake.DownloadManifestStub = nil
	if fake.downloadManifestReturnsOnCall == nil {
		fake.downloadManifestReturnsOnCall = make(map[int]struct {
			result1 v1.Manifest
			result2 error
		})
	}
	fake.downloadManifestReturnsOnCall[i] = struct {
		result1 v1.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistry) DownloadLayer(arg1 v1.Descriptor, arg2 string) error {
	fake.downloadLayerMutex.Lock()
	ret, specificReturn := fake.downloadLayerReturnsOnCall[len(fake.downloadLayerArgsForCall)]
	fake.downloadLayerArgsForCall = append(fake.downloadLayerArgsForCall, struct {
		arg1 v1.Descriptor
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("DownloadLayer", []interface{}{arg1, arg2})
	fake.downloadLayerMutex.Unlock()
	if fake.DownloadLayerStub != nil {
		return fake.DownloadLayerStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.downloadLayerReturns.result1
}

func (fake *FakeRegistry) DownloadLayerCallCount() int {
	fake.downloadLayerMutex.RLock()
	defer fake.downloadLayerMutex.RUnlock()
	return len(fake.downloadLayerArgsForCall)
}

func (fake *FakeRegistry) DownloadLayerArgsForCall(i int) (v1.Descriptor, string) {
	fake.downloadLayerMutex.RLock()
	defer fake.downloadLayerMutex.RUnlock()
	return fake.downloadLayerArgsForCall[i].arg1, fake.downloadLayerArgsForCall[i].arg2
}

func (fake *FakeRegistry) DownloadLayerReturns(result1 error) {
	fake.DownloadLayerStub = nil
	fake.downloadLayerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) DownloadLayerReturnsOnCall(i int, result1 error) {
	fake.DownloadLayerStub = nil
	if fake.downloadLayerReturnsOnCall == nil {
		fake.downloadLayerReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadLayerReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downloadManifestMutex.RLock()
	defer fake.downloadManifestMutex.RUnlock()
	fake.downloadLayerMutex.RLock()
	defer fake.downloadLayerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRegistry) recordInvocation(key string, args []interface{}) {
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

var _ hydrator.Registry = new(FakeRegistry)
