package TektonHub

import (
	"sync"

	"github.com/getgauge-contrib/gauge-go/gauge"
	"github.com/openshift-pipelines/release-tests/pkg/operator"
	"github.com/openshift-pipelines/release-tests/pkg/store"
)

var once sync.Once
var _ = gauge.Step("Validate Operator should be installed", func() {
	once.Do(func() {
		operator.ValidateOperatorInstallStatus(store.Clients(), store.GetCRNames())
	})
})
