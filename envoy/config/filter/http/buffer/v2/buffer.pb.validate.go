//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/buffer/v2/buffer.proto

package bufferv2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Buffer with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Buffer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Buffer with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BufferMultiError, or nil if none found.
func (m *Buffer) ValidateAll() error {
	return m.validate(true)
}

func (m *Buffer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetMaxRequestBytes(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := BufferValidationError{
				field:  "MaxRequestBytes",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	} else {
		err := BufferValidationError{
			field:  "MaxRequestBytes",
			reason: "value is required and must not be nil.",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return BufferMultiError(errors)
	}

	return nil
}

// BufferMultiError is an error wrapping multiple validation errors returned by
// Buffer.ValidateAll() if the designated constraints aren't met.
type BufferMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferMultiError) AllErrors() []error { return m }

// BufferValidationError is the validation error returned by Buffer.Validate if
// the designated constraints aren't met.
type BufferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferValidationError) ErrorName() string { return "BufferValidationError" }

// Error satisfies the builtin error interface
func (e BufferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferValidationError{}

// Validate checks the field values on BufferPerRoute with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BufferPerRoute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BufferPerRoute with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BufferPerRouteMultiError,
// or nil if none found.
func (m *BufferPerRoute) ValidateAll() error {
	return m.validate(true)
}

func (m *BufferPerRoute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofOverridePresent := false
	switch v := m.Override.(type) {
	case *BufferPerRoute_Disabled:
		if v == nil {
			err := BufferPerRouteValidationError{
				field:  "Override",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofOverridePresent = true

		if m.GetDisabled() != true {
			err := BufferPerRouteValidationError{
				field:  "Disabled",
				reason: "value must equal true",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	case *BufferPerRoute_Buffer:
		if v == nil {
			err := BufferPerRouteValidationError{
				field:  "Override",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofOverridePresent = true

		if m.GetBuffer() == nil {
			err := BufferPerRouteValidationError{
				field:  "Buffer",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetBuffer()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BufferPerRouteValidationError{
						field:  "Buffer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BufferPerRouteValidationError{
						field:  "Buffer",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetBuffer()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BufferPerRouteValidationError{
					field:  "Buffer",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofOverridePresent {
		err := BufferPerRouteValidationError{
			field:  "Override",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return BufferPerRouteMultiError(errors)
	}

	return nil
}

// BufferPerRouteMultiError is an error wrapping multiple validation errors
// returned by BufferPerRoute.ValidateAll() if the designated constraints
// aren't met.
type BufferPerRouteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BufferPerRouteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BufferPerRouteMultiError) AllErrors() []error { return m }

// BufferPerRouteValidationError is the validation error returned by
// BufferPerRoute.Validate if the designated constraints aren't met.
type BufferPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferPerRouteValidationError) ErrorName() string { return "BufferPerRouteValidationError" }

// Error satisfies the builtin error interface
func (e BufferPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferPerRouteValidationError{}
