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