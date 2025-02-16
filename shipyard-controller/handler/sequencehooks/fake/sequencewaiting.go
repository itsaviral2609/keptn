// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// ISequenceWaitingHookMock is a mock implementation of sequencehooks.ISequenceWaitingHook.
//
// 	func TestSomethingThatUsesISequenceWaitingHook(t *testing.T) {
//
// 		// make and configure a mocked sequencehooks.ISequenceWaitingHook
// 		mockedISequenceWaitingHook := &ISequenceWaitingHookMock{
// 			OnSequenceWaitingFunc: func(event apimodels.KeptnContextExtendedCE)  {
// 				panic("mock out the OnSequenceWaiting method")
// 			},
// 		}
//
// 		// use mockedISequenceWaitingHook in code that requires sequencehooks.ISequenceWaitingHook
// 		// and then make assertions.
//
// 	}
type ISequenceWaitingHookMock struct {
	// OnSequenceWaitingFunc mocks the OnSequenceWaiting method.
	OnSequenceWaitingFunc func(event apimodels.KeptnContextExtendedCE)

	// calls tracks calls to the methods.
	calls struct {
		// OnSequenceWaiting holds details about calls to the OnSequenceWaiting method.
		OnSequenceWaiting []struct {
			//models.KeptnContextExtendedCEis the event argument value.
			Event apimodels.KeptnContextExtendedCE
		}
	}
	lockOnSequenceWaiting sync.RWMutex
}

// OnSequenceWaiting calls OnSequenceWaitingFunc.
func (mock *ISequenceWaitingHookMock) OnSequenceWaiting(event apimodels.KeptnContextExtendedCE) {
	if mock.OnSequenceWaitingFunc == nil {
		panic("ISequenceWaitingHookMock.OnSequenceWaitingFunc: method is nil but ISequenceWaitingHook.OnSequenceWaiting was just called")
	}
	callInfo := struct {
		Event apimodels.KeptnContextExtendedCE
	}{
		Event: event,
	}
	mock.lockOnSequenceWaiting.Lock()
	mock.calls.OnSequenceWaiting = append(mock.calls.OnSequenceWaiting, callInfo)
	mock.lockOnSequenceWaiting.Unlock()
	mock.OnSequenceWaitingFunc(event)
}

// OnSequenceWaitingCalls gets all the calls that were made to OnSequenceWaiting.
// Check the length with:
//     len(mockedISequenceWaitingHook.OnSequenceWaitingCalls())
func (mock *ISequenceWaitingHookMock) OnSequenceWaitingCalls() []struct {
	Event apimodels.KeptnContextExtendedCE
} {
	var calls []struct {
		Event apimodels.KeptnContextExtendedCE
	}
	mock.lockOnSequenceWaiting.RLock()
	calls = mock.calls.OnSequenceWaiting
	mock.lockOnSequenceWaiting.RUnlock()
	return calls
}
