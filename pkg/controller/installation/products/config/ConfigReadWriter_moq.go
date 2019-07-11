// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package config

import (
	"github.com/integr8ly/integreatly-operator/pkg/apis/integreatly/v1alpha1"
	"sync"
)

var (
	lockConfigReadWriterMockReadAMQStreams       sync.RWMutex
	lockConfigReadWriterMockReadCodeReady        sync.RWMutex
	lockConfigReadWriterMockReadConfigForProduct sync.RWMutex
	lockConfigReadWriterMockReadRHSSO            sync.RWMutex
	lockConfigReadWriterMockWriteConfig          sync.RWMutex
)

// Ensure, that ConfigReadWriterMock does implement ConfigReadWriter.
// If this is not the case, regenerate this file with moq.
var _ ConfigReadWriter = &ConfigReadWriterMock{}

// ConfigReadWriterMock is a mock implementation of ConfigReadWriter.
//
//     func TestSomethingThatUsesConfigReadWriter(t *testing.T) {
//
//         // make and configure a mocked ConfigReadWriter
//         mockedConfigReadWriter := &ConfigReadWriterMock{
//             ReadAMQStreamsFunc: func() (*AMQStreams, error) {
// 	               panic("mock out the ReadAMQStreams method")
//             },
//             ReadCodeReadyFunc: func() (*CodeReady, error) {
// 	               panic("mock out the ReadCodeReady method")
//             },
//             ReadConfigForProductFunc: func(product v1alpha1.ProductName) (ProductConfig, error) {
// 	               panic("mock out the ReadConfigForProduct method")
//             },
//             ReadRHSSOFunc: func() (*RHSSO, error) {
// 	               panic("mock out the ReadRHSSO method")
//             },
//             WriteConfigFunc: func(config ConfigReadable) error {
// 	               panic("mock out the WriteConfig method")
//             },
//         }
//
//         // use mockedConfigReadWriter in code that requires ConfigReadWriter
//         // and then make assertions.
//
//     }
type ConfigReadWriterMock struct {
	// ReadAMQStreamsFunc mocks the ReadAMQStreams method.
	ReadAMQStreamsFunc func() (*AMQStreams, error)

	// ReadCodeReadyFunc mocks the ReadCodeReady method.
	ReadCodeReadyFunc func() (*CodeReady, error)

	// ReadConfigForProductFunc mocks the ReadConfigForProduct method.
	ReadConfigForProductFunc func(product v1alpha1.ProductName) (ProductConfig, error)

	// ReadRHSSOFunc mocks the ReadRHSSO method.
	ReadRHSSOFunc func() (*RHSSO, error)

	// WriteConfigFunc mocks the WriteConfig method.
	WriteConfigFunc func(config ConfigReadable) error

	// calls tracks calls to the methods.
	calls struct {
		// ReadAMQStreams holds details about calls to the ReadAMQStreams method.
		ReadAMQStreams []struct {
		}
		// ReadCodeReady holds details about calls to the ReadCodeReady method.
		ReadCodeReady []struct {
		}
		// ReadConfigForProduct holds details about calls to the ReadConfigForProduct method.
		ReadConfigForProduct []struct {
			// Product is the product argument value.
			Product v1alpha1.ProductName
		}
		// ReadRHSSO holds details about calls to the ReadRHSSO method.
		ReadRHSSO []struct {
		}
		// WriteConfig holds details about calls to the WriteConfig method.
		WriteConfig []struct {
			// Config is the config argument value.
			Config ConfigReadable
		}
	}
}

// ReadAMQStreams calls ReadAMQStreamsFunc.
func (mock *ConfigReadWriterMock) ReadAMQStreams() (*AMQStreams, error) {
	if mock.ReadAMQStreamsFunc == nil {
		panic("ConfigReadWriterMock.ReadAMQStreamsFunc: method is nil but ConfigReadWriter.ReadAMQStreams was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadAMQStreams.Lock()
	mock.calls.ReadAMQStreams = append(mock.calls.ReadAMQStreams, callInfo)
	lockConfigReadWriterMockReadAMQStreams.Unlock()
	return mock.ReadAMQStreamsFunc()
}

// ReadAMQStreamsCalls gets all the calls that were made to ReadAMQStreams.
// Check the length with:
//     len(mockedConfigReadWriter.ReadAMQStreamsCalls())
func (mock *ConfigReadWriterMock) ReadAMQStreamsCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadAMQStreams.RLock()
	calls = mock.calls.ReadAMQStreams
	lockConfigReadWriterMockReadAMQStreams.RUnlock()
	return calls
}

// ReadCodeReady calls ReadCodeReadyFunc.
func (mock *ConfigReadWriterMock) ReadCodeReady() (*CodeReady, error) {
	if mock.ReadCodeReadyFunc == nil {
		panic("ConfigReadWriterMock.ReadCodeReadyFunc: method is nil but ConfigReadWriter.ReadCodeReady was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadCodeReady.Lock()
	mock.calls.ReadCodeReady = append(mock.calls.ReadCodeReady, callInfo)
	lockConfigReadWriterMockReadCodeReady.Unlock()
	return mock.ReadCodeReadyFunc()
}

// ReadCodeReadyCalls gets all the calls that were made to ReadCodeReady.
// Check the length with:
//     len(mockedConfigReadWriter.ReadCodeReadyCalls())
func (mock *ConfigReadWriterMock) ReadCodeReadyCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadCodeReady.RLock()
	calls = mock.calls.ReadCodeReady
	lockConfigReadWriterMockReadCodeReady.RUnlock()
	return calls
}

// ReadConfigForProduct calls ReadConfigForProductFunc.
func (mock *ConfigReadWriterMock) ReadConfigForProduct(product v1alpha1.ProductName) (ProductConfig, error) {
	if mock.ReadConfigForProductFunc == nil {
		panic("ConfigReadWriterMock.ReadConfigForProductFunc: method is nil but ConfigReadWriter.ReadConfigForProduct was just called")
	}
	callInfo := struct {
		Product v1alpha1.ProductName
	}{
		Product: product,
	}
	lockConfigReadWriterMockReadConfigForProduct.Lock()
	mock.calls.ReadConfigForProduct = append(mock.calls.ReadConfigForProduct, callInfo)
	lockConfigReadWriterMockReadConfigForProduct.Unlock()
	return mock.ReadConfigForProductFunc(product)
}

// ReadConfigForProductCalls gets all the calls that were made to ReadConfigForProduct.
// Check the length with:
//     len(mockedConfigReadWriter.ReadConfigForProductCalls())
func (mock *ConfigReadWriterMock) ReadConfigForProductCalls() []struct {
	Product v1alpha1.ProductName
} {
	var calls []struct {
		Product v1alpha1.ProductName
	}
	lockConfigReadWriterMockReadConfigForProduct.RLock()
	calls = mock.calls.ReadConfigForProduct
	lockConfigReadWriterMockReadConfigForProduct.RUnlock()
	return calls
}

// ReadRHSSO calls ReadRHSSOFunc.
func (mock *ConfigReadWriterMock) ReadRHSSO() (*RHSSO, error) {
	if mock.ReadRHSSOFunc == nil {
		panic("ConfigReadWriterMock.ReadRHSSOFunc: method is nil but ConfigReadWriter.ReadRHSSO was just called")
	}
	callInfo := struct {
	}{}
	lockConfigReadWriterMockReadRHSSO.Lock()
	mock.calls.ReadRHSSO = append(mock.calls.ReadRHSSO, callInfo)
	lockConfigReadWriterMockReadRHSSO.Unlock()
	return mock.ReadRHSSOFunc()
}

// ReadRHSSOCalls gets all the calls that were made to ReadRHSSO.
// Check the length with:
//     len(mockedConfigReadWriter.ReadRHSSOCalls())
func (mock *ConfigReadWriterMock) ReadRHSSOCalls() []struct {
} {
	var calls []struct {
	}
	lockConfigReadWriterMockReadRHSSO.RLock()
	calls = mock.calls.ReadRHSSO
	lockConfigReadWriterMockReadRHSSO.RUnlock()
	return calls
}

// WriteConfig calls WriteConfigFunc.
func (mock *ConfigReadWriterMock) WriteConfig(config ConfigReadable) error {
	if mock.WriteConfigFunc == nil {
		panic("ConfigReadWriterMock.WriteConfigFunc: method is nil but ConfigReadWriter.WriteConfig was just called")
	}
	callInfo := struct {
		Config ConfigReadable
	}{
		Config: config,
	}
	lockConfigReadWriterMockWriteConfig.Lock()
	mock.calls.WriteConfig = append(mock.calls.WriteConfig, callInfo)
	lockConfigReadWriterMockWriteConfig.Unlock()
	return mock.WriteConfigFunc(config)
}

// WriteConfigCalls gets all the calls that were made to WriteConfig.
// Check the length with:
//     len(mockedConfigReadWriter.WriteConfigCalls())
func (mock *ConfigReadWriterMock) WriteConfigCalls() []struct {
	Config ConfigReadable
} {
	var calls []struct {
		Config ConfigReadable
	}
	lockConfigReadWriterMockWriteConfig.RLock()
	calls = mock.calls.WriteConfig
	lockConfigReadWriterMockWriteConfig.RUnlock()
	return calls
}
