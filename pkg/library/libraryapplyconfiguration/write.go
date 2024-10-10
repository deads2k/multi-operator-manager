package libraryapplyconfiguration

import (
	"errors"
	"fmt"
	"github.com/openshift/library-go/pkg/manifestclient"
	"path/filepath"
)

func WriteApplyConfiguration(desiredApplyConfiguration *ApplyConfiguration, outputDirectory string) error {
	errs := []error{}

	clusterTypeDir := filepath.Join(outputDirectory, string(desiredApplyConfiguration.DesiredConfigurationCluster.GetClusterType()))
	if err := WriteClusterApplyResult(desiredApplyConfiguration.DesiredConfigurationCluster, clusterTypeDir); err != nil {
		errs = append(errs, err)
	}

	clusterTypeDir = filepath.Join(outputDirectory, string(desiredApplyConfiguration.DesiredManagementCluster.GetClusterType()))
	if err := WriteClusterApplyResult(desiredApplyConfiguration.DesiredManagementCluster, clusterTypeDir); err != nil {
		errs = append(errs, err)
	}

	clusterTypeDir = filepath.Join(outputDirectory, string(desiredApplyConfiguration.DesiredUserWorkloadCluster.GetClusterType()))
	if err := WriteClusterApplyResult(desiredApplyConfiguration.DesiredUserWorkloadCluster, clusterTypeDir); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func WriteClusterApplyResult(desiredApplyConfiguration ClusterApplyResult, outputDirectory string) error {
	mutationRequests, err := desiredApplyConfiguration.Requests()
	if err != nil {
		return fmt.Errorf("failed getting requests: %w", err)
	}
	err = manifestclient.WriteMutationDirectory(outputDirectory, mutationRequests.AllRequests()...)
	if err != nil {
		return fmt.Errorf("failed writing requests: %w", err)
	}

	return nil
}
