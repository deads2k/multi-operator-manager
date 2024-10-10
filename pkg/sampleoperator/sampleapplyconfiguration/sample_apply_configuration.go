package sampleapplyconfiguration

import (
	"context"
	"time"

	"github.com/openshift/library-go/pkg/manifestclient"
	"github.com/openshift/multi-operator-manager/pkg/library/libraryapplyconfiguration"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

func SampleRunApplyConfiguration(ctx context.Context, inputDirectory string, now time.Time, streams genericiooptions.IOStreams) (libraryapplyconfiguration.AllDesiredMutationsGetter, error) {
	momClient := manifestclient.NewHTTPClient(inputDirectory)

	return libraryapplyconfiguration.NewApplyConfigurationFromClient(momClient.GetMutations()), nil
}
