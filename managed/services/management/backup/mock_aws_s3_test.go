// Code generated by mockery v1.0.0. DO NOT EDIT.

package backup

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// mockAwsS3 is an autogenerated mock type for the awsS3 type
type mockAwsS3 struct {
	mock.Mock
}

// BucketExists provides a mock function with given fields: ctx, host, accessKey, secretKey, name
func (_m *mockAwsS3) BucketExists(ctx context.Context, host string, accessKey string, secretKey string, name string) (bool, error) {
	ret := _m.Called(ctx, host, accessKey, secretKey, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) bool); ok {
		r0 = rf(ctx, host, accessKey, secretKey, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, host, accessKey, secretKey, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLocation provides a mock function with given fields: ctx, host, accessKey, secretKey, name
func (_m *mockAwsS3) GetBucketLocation(ctx context.Context, host string, accessKey string, secretKey string, name string) (string, error) {
	ret := _m.Called(ctx, host, accessKey, secretKey, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) string); ok {
		r0 = rf(ctx, host, accessKey, secretKey, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, host, accessKey, secretKey, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRecursive provides a mock function with given fields: ctx, endpoint, accessKey, secretKey, bucketName, prefix
func (_m *mockAwsS3) RemoveRecursive(ctx context.Context, endpoint string, accessKey string, secretKey string, bucketName string, prefix string) error {
	ret := _m.Called(ctx, endpoint, accessKey, secretKey, bucketName, prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string, string) error); ok {
		r0 = rf(ctx, endpoint, accessKey, secretKey, bucketName, prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}