// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/network/redis_proxy/v2/redis_proxy.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on RedisProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RedisProxy) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetStatPrefix()) < 1 {
		return RedisProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetCluster()) < 1 {
		return RedisProxyValidationError{
			field:  "Cluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if m.GetSettings() == nil {
		return RedisProxyValidationError{
			field:  "Settings",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetSettings()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RedisProxyValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for LatencyInMicros

	return nil
}

// RedisProxyValidationError is the validation error returned by
// RedisProxy.Validate if the designated constraints aren't met.
type RedisProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxyValidationError) ErrorName() string { return "RedisProxyValidationError" }

// Error satisfies the builtin error interface
func (e RedisProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxyValidationError{}

// Validate checks the field values on RedisProxy_ConnPoolSettings with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RedisProxy_ConnPoolSettings) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOpTimeout() == nil {
		return RedisProxy_ConnPoolSettingsValidationError{
			field:  "OpTimeout",
			reason: "value is required",
		}
	}

	// no validation rules for EnableHashtagging

	return nil
}

// RedisProxy_ConnPoolSettingsValidationError is the validation error returned
// by RedisProxy_ConnPoolSettings.Validate if the designated constraints
// aren't met.
type RedisProxy_ConnPoolSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisProxy_ConnPoolSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisProxy_ConnPoolSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisProxy_ConnPoolSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisProxy_ConnPoolSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisProxy_ConnPoolSettingsValidationError) ErrorName() string {
	return "RedisProxy_ConnPoolSettingsValidationError"
}

// Error satisfies the builtin error interface
func (e RedisProxy_ConnPoolSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedisProxy_ConnPoolSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisProxy_ConnPoolSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisProxy_ConnPoolSettingsValidationError{}
