/*
Copyright 2018 Pressinfra SRL

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

// nolint: errcheck
package orchestrator

import (
	"path/filepath"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"sigs.k8s.io/controller-runtime/pkg/envtest/printer"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/bitpoke/mysql-operator/pkg/apis"
	"github.com/bitpoke/mysql-operator/pkg/controller/internal/testutil"
)

var cfg *rest.Config
var t *envtest.Environment

func TestMysqlClusterController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Orchestrator Controller Suite", []Reporter{printer.NewlineReporter{}})
}

var _ = BeforeSuite(func() {

	// set reconcile time to 1 second to speed up the tests.
	reconcileTimePeriod = time.Second

	logf.SetLogger(testutil.NewTestLogger(GinkgoWriter))

	var err error

	t = &envtest.Environment{
		CRDDirectoryPaths: []string{filepath.Join("..", "..", "..", "config", "crd", "bases")},
	}

	apis.AddToScheme(scheme.Scheme)

	cfg, err = t.Start()
	Expect(err).NotTo(HaveOccurred())

})

var _ = AfterSuite(func() {
	t.Stop()
})
