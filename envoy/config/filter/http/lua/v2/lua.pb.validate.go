//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/lua/v2/lua.proto

package luav2

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

// Validate checks the field values on Lua with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Lua) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Lua with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in LuaMultiError, or nil if none found.
func (m *Lua) ValidateAll() error {
	return m.validate(true)
}

func (m *Lua) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetInlineCode()) < 1 {
		err := LuaValidationError{
			field:  "InlineCode",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LuaMultiError(errors)
	}

	return nil
}

// LuaMultiError is an error wrapping multiple validation errors returned by
// Lua.ValidateAll() if the designated constraints aren't met.
type LuaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LuaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LuaMultiError) AllErrors() []error { return m }

// LuaValidationError is the validation error returned by Lua.Validate if the
// designated constraints aren't met.
type LuaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LuaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LuaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LuaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LuaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LuaValidationError) ErrorName() string { return "LuaValidationError" }

// Error satisfies the builtin error interface
func (e LuaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLua.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LuaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LuaValidationError{}
