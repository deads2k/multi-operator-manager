package libraryapplyconfiguration

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericiooptions"
)

type ApplyConfigurationFunc func(ctx context.Context, inputDirectory string, now time.Time, streams genericiooptions.IOStreams) (AllDesiredMutationsGetter, error)

func NewSampleOperatorApplyConfigurationCommand(applyConfigurationFn ApplyConfigurationFunc, streams genericiooptions.IOStreams) *cobra.Command {
	return newSampleOperatorApplyConfigurationCommand(applyConfigurationFn, streams)
}

type sampleOperatorApplyConfigurationFlags struct {
	applyConfigurationFn ApplyConfigurationFunc

	// InputDirectory is a directory that contains the must-gather formatted inputs
	inputDirectory string

	// OutputDirectory is the directory to where output should be stored
	outputDirectory string

	streams genericiooptions.IOStreams
}

func newSampleOperatorApplyConfigurationFlags(streams genericiooptions.IOStreams) *sampleOperatorApplyConfigurationFlags {
	return &sampleOperatorApplyConfigurationFlags{
		streams: streams,
	}
}

func newSampleOperatorApplyConfigurationCommand(applyConfigurationFn ApplyConfigurationFunc, streams genericiooptions.IOStreams) *cobra.Command {
	f := newSampleOperatorApplyConfigurationFlags(streams)
	f.applyConfigurationFn = applyConfigurationFn

	cmd := &cobra.Command{
		Use:   "apply-configuration",
		Short: "Sample operator the apply-configuration command.",

		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(f.streams.ErrOut, "stderr output here\n")
			fmt.Fprintf(f.streams.Out, "stdout output here\n")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			if err := f.Validate(); err != nil {
				return err
			}
			o, err := f.ToOptions(ctx)
			if err != nil {
				return err
			}
			if err := o.Run(ctx); err != nil {
				return err
			}
			return nil
		},
	}

	f.BindFlags(cmd.Flags())

	return cmd
}

func (f *sampleOperatorApplyConfigurationFlags) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(&f.inputDirectory, "input-dir", f.inputDirectory, "The directory where the resource input is stored.")
	flags.StringVar(&f.outputDirectory, "output-dir", f.outputDirectory, "The directory where the output is stored.")
}

func (f *sampleOperatorApplyConfigurationFlags) Validate() error {
	if len(f.inputDirectory) == 0 {
		return fmt.Errorf("--input-dir is required")
	}
	if len(f.outputDirectory) == 0 {
		return fmt.Errorf("--output-dir is required")
	}
	return nil
}

func (f *sampleOperatorApplyConfigurationFlags) ToOptions(ctx context.Context) (*ApplyConfigurationOptions, error) {
	return NewApplyConfigurationOptions(
			f.applyConfigurationFn,
			f.inputDirectory,
			f.outputDirectory,
			time.Now(),
			f.streams,
		),
		nil
}
