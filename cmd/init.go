package cmd

import (
	"fmt"
	"os"

	"github.com/HeitorAugustoLN/constance-installer/styles"
	"github.com/HeitorAugustoLN/constance-installer/utils/packages"
	"github.com/HeitorAugustoLN/constance-installer/utils/system"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes constance installer for usage",
	Long:  `Initializes constance installer for usage`,

	Run: setup,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func setup(cmd *cobra.Command, args []string) {
	if !system.CheckIfProgramAlreadySetup() {
		runtime, operatingSystem := system.GetOperatingSystem()

		system.CheckOperatingSystem()
		packageManagers := packages.CheckPackageManager(runtime, operatingSystem)
		installedPackages := packages.GetInstalledPackages(packageManagers)
		system.BootstrapConstanceFiles()
		system.ParseJSONToFiles(installedPackages)
		fmt.Printf("%s\n", styles.YellowText.Render("Done! Constance is now set up."))
		packages.CreateGeneration()
	} else {
		fmt.Println("Constance is already set up.")
		os.Exit(1)
	}
}
