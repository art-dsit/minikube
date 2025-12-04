/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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
	"log"
	"os"
	"path/filepath"

	"github.com/docker/machine/libmachine/mcnutils"
	"github.com/pkg/errors"
)

// Disable reverts containerd config files and restarts containerd
func Disable() error {
	log.Print("Disabling gvisor...")
	if err := os.Remove(filepath.Join(nodeDir, containerdConfigPath)); err != nil {
		return errors.Wrapf(err, "removing %s", containerdConfigPath)
	}
	log.Printf("Restoring default config.toml at %s", containerdConfigPath)
	if err := mcnutils.CopyFile(filepath.Join(nodeDir, containerdConfigBackupPath), filepath.Join(nodeDir, containerdConfigPath)); err != nil {
		return errors.Wrap(err, "reverting back to default config.toml")
	}

	// Remove runsc.toml if it exists (created for cgroupv2 support)
	runscConfigFullPath := filepath.Join(nodeDir, runscConfigPath)
	if _, err := os.Stat(runscConfigFullPath); err == nil {
		log.Printf("Removing %s", runscConfigPath)
		if err := os.Remove(runscConfigFullPath); err != nil {
			log.Printf("Warning: failed to remove %s: %v", runscConfigPath, err)
		}
	}

	// restart containerd
	if err := restartContainerd(); err != nil {
		return errors.Wrap(err, "restarting containerd")
	}
	log.Print("Successfully disabled gvisor")
	return nil
}
