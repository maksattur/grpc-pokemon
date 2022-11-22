// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pokemon.proto

package pokemonpc

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

// Validate checks the field values on Pokemon with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Pokemon) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pokemon with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PokemonMultiError, or nil if none found.
func (m *Pokemon) ValidateAll() error {
	return m.validate(true)
}

func (m *Pokemon) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Pid

	// no validation rules for Name

	// no validation rules for Power

	// no validation rules for Description

	if len(errors) > 0 {
		return PokemonMultiError(errors)
	}

	return nil
}

// PokemonMultiError is an error wrapping multiple validation errors returned
// by Pokemon.ValidateAll() if the designated constraints aren't met.
type PokemonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PokemonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PokemonMultiError) AllErrors() []error { return m }

// PokemonValidationError is the validation error returned by Pokemon.Validate
// if the designated constraints aren't met.
type PokemonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PokemonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PokemonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PokemonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PokemonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PokemonValidationError) ErrorName() string { return "PokemonValidationError" }

// Error satisfies the builtin error interface
func (e PokemonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPokemon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PokemonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PokemonValidationError{}

// Validate checks the field values on CreatePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePokemonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePokemonRequestMultiError, or nil if none found.
func (m *CreatePokemonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePokemonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePokemonRequestValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePokemonRequestValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePokemonRequestValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePokemonRequestMultiError(errors)
	}

	return nil
}

// CreatePokemonRequestMultiError is an error wrapping multiple validation
// errors returned by CreatePokemonRequest.ValidateAll() if the designated
// constraints aren't met.
type CreatePokemonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePokemonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePokemonRequestMultiError) AllErrors() []error { return m }

// CreatePokemonRequestValidationError is the validation error returned by
// CreatePokemonRequest.Validate if the designated constraints aren't met.
type CreatePokemonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePokemonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePokemonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePokemonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePokemonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePokemonRequestValidationError) ErrorName() string {
	return "CreatePokemonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePokemonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePokemonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePokemonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePokemonRequestValidationError{}

// Validate checks the field values on CreatePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePokemonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePokemonResponseMultiError, or nil if none found.
func (m *CreatePokemonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePokemonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreatePokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreatePokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePokemonResponseValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreatePokemonResponseMultiError(errors)
	}

	return nil
}

// CreatePokemonResponseMultiError is an error wrapping multiple validation
// errors returned by CreatePokemonResponse.ValidateAll() if the designated
// constraints aren't met.
type CreatePokemonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePokemonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePokemonResponseMultiError) AllErrors() []error { return m }

// CreatePokemonResponseValidationError is the validation error returned by
// CreatePokemonResponse.Validate if the designated constraints aren't met.
type CreatePokemonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePokemonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePokemonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePokemonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePokemonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePokemonResponseValidationError) ErrorName() string {
	return "CreatePokemonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePokemonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePokemonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePokemonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePokemonResponseValidationError{}

// Validate checks the field values on ReadPokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadPokemonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadPokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadPokemonRequestMultiError, or nil if none found.
func (m *ReadPokemonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadPokemonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pid

	if len(errors) > 0 {
		return ReadPokemonRequestMultiError(errors)
	}

	return nil
}

// ReadPokemonRequestMultiError is an error wrapping multiple validation errors
// returned by ReadPokemonRequest.ValidateAll() if the designated constraints
// aren't met.
type ReadPokemonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadPokemonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadPokemonRequestMultiError) AllErrors() []error { return m }

// ReadPokemonRequestValidationError is the validation error returned by
// ReadPokemonRequest.Validate if the designated constraints aren't met.
type ReadPokemonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadPokemonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadPokemonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadPokemonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadPokemonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadPokemonRequestValidationError) ErrorName() string {
	return "ReadPokemonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReadPokemonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadPokemonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadPokemonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadPokemonRequestValidationError{}

// Validate checks the field values on ReadPokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadPokemonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadPokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadPokemonResponseMultiError, or nil if none found.
func (m *ReadPokemonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadPokemonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadPokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadPokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadPokemonResponseValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReadPokemonResponseMultiError(errors)
	}

	return nil
}

// ReadPokemonResponseMultiError is an error wrapping multiple validation
// errors returned by ReadPokemonResponse.ValidateAll() if the designated
// constraints aren't met.
type ReadPokemonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadPokemonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadPokemonResponseMultiError) AllErrors() []error { return m }

// ReadPokemonResponseValidationError is the validation error returned by
// ReadPokemonResponse.Validate if the designated constraints aren't met.
type ReadPokemonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadPokemonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadPokemonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadPokemonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadPokemonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadPokemonResponseValidationError) ErrorName() string {
	return "ReadPokemonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReadPokemonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadPokemonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadPokemonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadPokemonResponseValidationError{}

// Validate checks the field values on UpdatePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePokemonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePokemonRequestMultiError, or nil if none found.
func (m *UpdatePokemonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePokemonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePokemonRequestValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePokemonRequestValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePokemonRequestValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePokemonRequestMultiError(errors)
	}

	return nil
}

// UpdatePokemonRequestMultiError is an error wrapping multiple validation
// errors returned by UpdatePokemonRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdatePokemonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePokemonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePokemonRequestMultiError) AllErrors() []error { return m }

// UpdatePokemonRequestValidationError is the validation error returned by
// UpdatePokemonRequest.Validate if the designated constraints aren't met.
type UpdatePokemonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePokemonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePokemonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePokemonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePokemonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePokemonRequestValidationError) ErrorName() string {
	return "UpdatePokemonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePokemonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePokemonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePokemonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePokemonRequestValidationError{}

// Validate checks the field values on UpdatePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdatePokemonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdatePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdatePokemonResponseMultiError, or nil if none found.
func (m *UpdatePokemonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdatePokemonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdatePokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdatePokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePokemonResponseValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdatePokemonResponseMultiError(errors)
	}

	return nil
}

// UpdatePokemonResponseMultiError is an error wrapping multiple validation
// errors returned by UpdatePokemonResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdatePokemonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdatePokemonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdatePokemonResponseMultiError) AllErrors() []error { return m }

// UpdatePokemonResponseValidationError is the validation error returned by
// UpdatePokemonResponse.Validate if the designated constraints aren't met.
type UpdatePokemonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePokemonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePokemonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePokemonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePokemonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePokemonResponseValidationError) ErrorName() string {
	return "UpdatePokemonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePokemonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePokemonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePokemonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePokemonResponseValidationError{}

// Validate checks the field values on DeletePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePokemonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePokemonRequestMultiError, or nil if none found.
func (m *DeletePokemonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePokemonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pid

	if len(errors) > 0 {
		return DeletePokemonRequestMultiError(errors)
	}

	return nil
}

// DeletePokemonRequestMultiError is an error wrapping multiple validation
// errors returned by DeletePokemonRequest.ValidateAll() if the designated
// constraints aren't met.
type DeletePokemonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePokemonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePokemonRequestMultiError) AllErrors() []error { return m }

// DeletePokemonRequestValidationError is the validation error returned by
// DeletePokemonRequest.Validate if the designated constraints aren't met.
type DeletePokemonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePokemonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePokemonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePokemonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePokemonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePokemonRequestValidationError) ErrorName() string {
	return "DeletePokemonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePokemonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePokemonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePokemonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePokemonRequestValidationError{}

// Validate checks the field values on DeletePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeletePokemonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeletePokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeletePokemonResponseMultiError, or nil if none found.
func (m *DeletePokemonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeletePokemonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Pid

	if len(errors) > 0 {
		return DeletePokemonResponseMultiError(errors)
	}

	return nil
}

// DeletePokemonResponseMultiError is an error wrapping multiple validation
// errors returned by DeletePokemonResponse.ValidateAll() if the designated
// constraints aren't met.
type DeletePokemonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeletePokemonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeletePokemonResponseMultiError) AllErrors() []error { return m }

// DeletePokemonResponseValidationError is the validation error returned by
// DeletePokemonResponse.Validate if the designated constraints aren't met.
type DeletePokemonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeletePokemonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeletePokemonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeletePokemonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeletePokemonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeletePokemonResponseValidationError) ErrorName() string {
	return "DeletePokemonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeletePokemonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeletePokemonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeletePokemonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeletePokemonResponseValidationError{}

// Validate checks the field values on ListPokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPokemonRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPokemonRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPokemonRequestMultiError, or nil if none found.
func (m *ListPokemonRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPokemonRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListPokemonRequestMultiError(errors)
	}

	return nil
}

// ListPokemonRequestMultiError is an error wrapping multiple validation errors
// returned by ListPokemonRequest.ValidateAll() if the designated constraints
// aren't met.
type ListPokemonRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPokemonRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPokemonRequestMultiError) AllErrors() []error { return m }

// ListPokemonRequestValidationError is the validation error returned by
// ListPokemonRequest.Validate if the designated constraints aren't met.
type ListPokemonRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPokemonRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPokemonRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPokemonRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPokemonRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPokemonRequestValidationError) ErrorName() string {
	return "ListPokemonRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPokemonRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPokemonRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPokemonRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPokemonRequestValidationError{}

// Validate checks the field values on ListPokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListPokemonResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListPokemonResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListPokemonResponseMultiError, or nil if none found.
func (m *ListPokemonResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListPokemonResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPokemon()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListPokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListPokemonResponseValidationError{
					field:  "Pokemon",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPokemon()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListPokemonResponseValidationError{
				field:  "Pokemon",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListPokemonResponseMultiError(errors)
	}

	return nil
}

// ListPokemonResponseMultiError is an error wrapping multiple validation
// errors returned by ListPokemonResponse.ValidateAll() if the designated
// constraints aren't met.
type ListPokemonResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListPokemonResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListPokemonResponseMultiError) AllErrors() []error { return m }

// ListPokemonResponseValidationError is the validation error returned by
// ListPokemonResponse.Validate if the designated constraints aren't met.
type ListPokemonResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPokemonResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPokemonResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPokemonResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPokemonResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPokemonResponseValidationError) ErrorName() string {
	return "ListPokemonResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPokemonResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPokemonResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPokemonResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPokemonResponseValidationError{}