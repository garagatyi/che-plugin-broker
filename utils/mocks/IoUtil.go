// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"

// IoUtil is an autogenerated mock type for the IoUtil type
type IoUtil struct {
	mock.Mock
}

// CopyFile provides a mock function with given fields: src, dest
func (_m *IoUtil) CopyFile(src string, dest string) error {
	ret := _m.Called(src, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(src, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CopyResource provides a mock function with given fields: src, dest
func (_m *IoUtil) CopyResource(src string, dest string) error {
	ret := _m.Called(src, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(src, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateFile provides a mock function with given fields: file, tr
func (_m *IoUtil) CreateFile(file string, tr io.Reader) error {
	ret := _m.Called(file, tr)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(file, tr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Download provides a mock function with given fields: URL, destPath
func (_m *IoUtil) Download(URL string, destPath string) error {
	ret := _m.Called(URL, destPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(URL, destPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResolveDestPath provides a mock function with given fields: filePath, destDir
func (_m *IoUtil) ResolveDestPath(filePath string, destDir string) string {
	ret := _m.Called(filePath, destDir)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(filePath, destDir)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResolveDestPathFromURL provides a mock function with given fields: url, destDir
func (_m *IoUtil) ResolveDestPathFromURL(url string, destDir string) string {
	ret := _m.Called(url, destDir)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(url, destDir)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TempDir provides a mock function with given fields: _a0, _a1
func (_m *IoUtil) TempDir(_a0 string, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Untar provides a mock function with given fields: tarPath, dest
func (_m *IoUtil) Untar(tarPath string, dest string) error {
	ret := _m.Called(tarPath, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(tarPath, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unzip provides a mock function with given fields: arch, dest
func (_m *IoUtil) Unzip(arch string, dest string) error {
	ret := _m.Called(arch, dest)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(arch, dest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
