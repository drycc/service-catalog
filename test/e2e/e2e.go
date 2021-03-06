/*
Copyright 2015 The Kubernetes Authors.

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

package e2e

import (
	"strings"
	"testing"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/component-base/logs"

	"github.com/kubernetes-sigs/service-catalog/test/e2e/framework"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	"github.com/onsi/gomega"
	"k8s.io/klog/v2"
)

// RunE2ETests checks configuration parameters (specified through flags) and then runs
// E2E tests using the Ginkgo runner.
func RunE2ETests(t *testing.T) {
	logs.InitLogs()
	defer logs.FlushLogs()

	gomega.RegisterFailHandler(ginkgo.Fail)

	// Disable skipped tests unless they are explicitly requested.
	focusString := strings.Join(config.GinkgoConfig.FocusStrings, "|")
	skipString := strings.Join(config.GinkgoConfig.SkipStrings, "|")
	if focusString == "" && skipString == "" {
		config.GinkgoConfig.SkipStrings = []string{`\[Flaky\]`, `\[Feature:.+\]`}
	}

	klog.Infof("Starting e2e run %q on Ginkgo node %d", framework.RunId, config.GinkgoConfig.ParallelNode)
	ginkgo.RunSpecs(t, "Service Catalog e2e suite")
}

