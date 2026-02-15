// Copyright © 2026 Timothy E. Peoples

package constr

import (
	"errors"
	"testing"
)

type begetType int

const (
	begetNil begetType = iota
	begetError
	begetSentinel
)

const (
	errTestVal  = Error("a test error value")
	sentinelVal = Sentinel("a test sentinel value")
)

func beget(bt begetType) error {
	switch bt {
	case begetNil:
		return nil
	case begetError:
		return errTestVal
	case begetSentinel:
		return sentinelVal
	default:
		return errors.New("<BROKEN>")
	}
}

//╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴╶╴

type testcase struct {
	what begetType
	want error
}

func (tc testcase) test(t *testing.T) {
	if got := beget(tc.what); got != tc.want {
		t.Errorf("beget(%v) == (%v); wanted (%v)", tc.what, got, tc.want)
	}
}

func TestAll(t *testing.T) {
	cases := map[string]testcase{
		"nil":      {begetNil, nil},
		"error":    {begetError, errTestVal},
		"sentinel": {begetSentinel, sentinelVal},
	}

	for name, tc := range cases {
		t.Run(name, tc.test)
	}
}
