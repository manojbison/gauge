package TektonHub

import (
	"fmt"
	"os"
	"sync"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/k8s"
	"github.com/openshift-pipelines/release-tests/pkg/operator"
	"github.com/openshift-pipelines/release-tests/pkg/store"
)

var once sync.Once
var _ = gauge.Step("Validate Operator should be installed", func() {
	once.Do(func() {
		operator.ValidateOperatorInstallStatus(store.Clients(), store.GetCRNames())
	})
})

var _ = gauge.Step("Verify ServiceAccount <sa> exist", func(sa string) {
	k8s.VerifyServiceAccountExists(store.Clients().Ctx, store.Clients().KubeClient, sa, store.Namespace())
})

var _ = gauge.Step("Create TektonHub CR yaml file", func() {
	data := []byte(`apiVersion: operator.tekton.dev/v1alpha1
	kind: TektonChain
	metadata:
	  name: chain
	spec:
	  targetNamespace: openshift-pipelines`)
	fileName := "./Tektonhub.yaml"
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Println("Error creating YAML file:", err)
		return
	}

	fmt.Println("YAML file created successfully:", fileName)
})

