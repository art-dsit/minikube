/*
Copyright 2025 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gvisor

import (
	"strings"
	"testing"
)

func TestConfigFragment(t *testing.T) {
	// Note: configFragment() now depends on detectCgroupVersion() which checks
	// the actual filesystem at /node/sys/fs/cgroup. In a real test environment,
	// this will detect the actual cgroup version of the host system.
	// This test just verifies the base structure is always present.

	got := configFragment()

	// These should always be present regardless of cgroup version
	wantContains := []string{
		`[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runsc]`,
		`runtime_type = "io.containerd.runsc.v1"`,
		`pod_annotations = [ "dev.gvisor.*" ]`,
	}

	for _, want := range wantContains {
		if !strings.Contains(got, want) {
			t.Errorf("configFragment() missing expected substring:\ngot:\n%s\n\nwant to contain:\n%s", got, want)
		}
	}

}
