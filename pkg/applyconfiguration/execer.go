package applyconfiguration

import (
	"context"
	"errors"
	"fmt"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"os"
	"os/exec"
	"path/filepath"
)

func ApplyConfiguration(ctx context.Context, binaryPath, inputDirectory, outputDirectory string) (ApplyConfigurationResult, error) {
	// the cmd.Wait() closes these output files.
	stdoutFilename := filepath.Join(outputDirectory, "stdout.log")
	stdoutFile, err := os.OpenFile(stdoutFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("unable to open stdout.log: %w", err)
	}
	stderrFilename := filepath.Join(outputDirectory, "stderr.log")
	stderrFile, err := os.OpenFile(stderrFilename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("unable to open stderr.log: %w", err)
	}

	args := []string{
		"apply-configuration",
		"--input-dir", inputDirectory,
		"--output-dir", outputDirectory,
	}

	cmd := exec.CommandContext(ctx, binaryPath, args...)
	cmd.Stdout = stdoutFile
	cmd.Stderr = stderrFile
	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start process: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			if err := stdoutFile.Close(); err != nil {
				utilruntime.HandleError(err)
			}
			if err := stderrFile.Close(); err != nil {
				utilruntime.HandleError(err)
			}
			return NewApplyConfigurationResult(outputDirectory,
				fmt.Errorf("failed to wait for process %v: %w stderr: %v", cmd, err, string(exitErr.Stderr)))
		}
		return NewApplyConfigurationResult(outputDirectory,
			fmt.Errorf("failed to wait for process: %w", err))
	}

	return NewApplyConfigurationResult(outputDirectory, nil)
}
