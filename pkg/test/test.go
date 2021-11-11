// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package test

import (
	"encoding/json"
	"errors"

	"github.com/insomniacslk/xjson"
	"github.com/linuxboot/contest/pkg/target"
)

// RetryParameters describes optional parameters for retry
type RetryParameters struct {
	NumRetries    uint32
	RetryInterval *xjson.Duration
}

// Test describes a test definition.
type Test struct {
	Name                string
	TestStepsBundles    []TestStepBundle
	TargetManagerBundle *target.TargetManagerBundle
	TestFetcherBundle   *TestFetcherBundle
	RetryParameters     RetryParameters
}

// TestDescriptor models the JSON encoded blob which is given as input to the
// job creation request. The test descriptors are part of the main JobDescriptor
// JSON document.
type TestDescriptor struct {
	// Disabled allows to disable the test
	Disabled bool

	RetryParameters RetryParameters

	// TargetManager-related parameters
	TargetManagerName              string
	TargetManagerAcquireParameters json.RawMessage
	TargetManagerReleaseParameters json.RawMessage

	// TestFetcher-related parameters
	TestFetcherName            string
	TestFetcherFetchParameters json.RawMessage
}

// Validate performs sanity checks on the Descriptor
func (d *TestDescriptor) Validate() error {
	if d.TargetManagerName == "" {
		return errors.New("target manager name cannot be empty")
	}
	if d.TestFetcherName == "" {
		return errors.New("test fetcher name cannot be empty")
	}
	return nil
}
