package cmd

import (
	"github.com/spf13/cobra"

	"github.com/myshkin5/moqueries/logs"
)

const (
	debugFlag  = "debug"
	exportFlag = "export"

	destinationFlag = "destination"

	packageFlag    = "package"
	importFlag     = "import"
	testImportFlag = "test-import"

	shallowPointerCompareFlag   = "shallow-pointer-compare"
	shallowInterfaceCompareFlag = "shallow-interface-compare"
)

var rootCmd = &cobra.Command{
	Use:   "moqueries [interfaces and/or function types to mock]",
	Short: "Moqueries generates simple but thread-safe mocks",
	Args:  cobra.MinimumNArgs(1),
	Run:   generate,
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().Bool(debugFlag, false,
		"If true, debugging output will be logged")
	rootCmd.PersistentFlags().Bool(exportFlag, false,
		"If true, generated mocks will be exported and accessible from other packages")

	rootCmd.PersistentFlags().String(destinationFlag, "",
		"The file path where mocks are generated relative to the directory "+
			"containing the generate directive (or relative to the current directory)")
	err := rootCmd.MarkPersistentFlagRequired(destinationFlag)
	if err != nil {
		logs.Panic("Error configuring required flag", err)
	}

	rootCmd.PersistentFlags().String(packageFlag, "",
		"The package to generate code into (defaults to the test package of the "+
			"destination directory when --export=false or the package of the "+
			"destination directory when --export=true)")
	rootCmd.PersistentFlags().String(importFlag, ".",
		"The package containing the type (interface or function type) to be "+
			"mocked (defaults to the directory containing generate directive)")
	rootCmd.PersistentFlags().Bool(testImportFlag, false,
		"Look for the types to be mocked in the test package")

	rootCmd.PersistentFlags().Bool(shallowPointerCompareFlag, false,
		"Compare pointer by value rather than a deep comparison")
	rootCmd.PersistentFlags().Bool(shallowInterfaceCompareFlag, false,
		"Compare interfaces by value rather than a deep comparison (may cause a panic)")
}

// Execute generates one or more mocks
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logs.Panic("Error executing command", err)
	}
}
