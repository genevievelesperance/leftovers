package fakes

import "sync"

type RouteTables struct {
	DeleteCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			VpcId string
		}
		Returns struct {
			Error error
		}
		Stub func(string) error
	}
}

func (f *RouteTables) Delete(param1 string) error {
	f.DeleteCall.Lock()
	defer f.DeleteCall.Unlock()
	f.DeleteCall.CallCount++
	f.DeleteCall.Receives.VpcId = param1
	if f.DeleteCall.Stub != nil {
		return f.DeleteCall.Stub(param1)
	}
	return f.DeleteCall.Returns.Error
}
