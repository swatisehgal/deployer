/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2021 Red Hat, Inc.
 */

package rte

import (
	"reflect"
	"testing"

	"github.com/k8stopologyawareschedwg/deployer/pkg/deployer/platform"
)

func TestClone(t *testing.T) {
	type testCase struct {
		name string
		mf   Manifests
		plat platform.Platform
	}

	testCases := []testCase{
		{
			name: "kubernetes manifests",
			plat: platform.Kubernetes,
		},
		{
			name: "openshift manifests",
			plat: platform.OpenShift,
		},
	}

	for _, tc := range testCases {
		tc.mf, _ = GetManifests(tc.plat, "")
		cMf := tc.mf.Clone()

		if &cMf == &tc.mf {
			t.Errorf("testcase %q, Clone() should create a pristine copy of Manifests object, thus should have different addresses", tc.name)
		}
	}
}

func TestRender(t *testing.T) {
	type testCase struct {
		name string
		mf   Manifests
		plat platform.Platform
	}

	testCases := []testCase{
		{
			name: "kubernetes manifests",
			plat: platform.Kubernetes,
		},
		{
			name: "openshift manifests",
			plat: platform.OpenShift,
		},
	}

	for _, tc := range testCases {
		tc.mf, _ = GetManifests(tc.plat, "")
		mfBeforeRender := tc.mf.Clone()
		uMf := tc.mf.Render(RenderOptions{})

		if &uMf == &tc.mf {
			t.Errorf("testcase %q, Render() should return a pristine copy of Manifests object, thus should have different addresses", tc.name)
		}
		if !reflect.DeepEqual(mfBeforeRender, tc.mf) {
			t.Errorf("testcase %q, Render() should not modify the original Manifests object", tc.name)
		}
	}
}

func TestGetManifestsOpenShift(t *testing.T) {
	mf, err := GetManifests(platform.OpenShift, "test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if mf.SecurityContextConstraint == nil {
		t.Fatalf("no security context constraint is generated for the OpenShift platform")
	}

	if mf.MachineConfig == nil {
		t.Fatalf("no machine config is generated for the OpenShift platform")
	}

	if mf.DaemonSet == nil {
		t.Fatalf("no daemon set is generated for the OpenShift platform")
	}
}
