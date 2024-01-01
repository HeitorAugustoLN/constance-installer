package packages

import (
	"fmt"
	"os"

	"github.com/HeitorAugustoLN/constance-installer/styles"
	"github.com/HeitorAugustoLN/constance-installer/utils/system"
	"github.com/dlclark/regexp2"
)

func CheckPackageManager(runtime string, operatingSystem string) []string {
	userPackageManagers := []string{}
	packageManagers := map[string][]string{
		"windows": {"scoop", "choco", "winget"},
		"darwin":  {"brew"},
		"linux": {
			"apt",
			"snap",
			"flatpak",
			"pacman",
			"dnf",
			"xbps-install",
			"zypper",

			"aura",
			"aurman",
			"pacaur",
			"pakku",
			"yay",
			"trizen",
			"paru",
			"pikaur",

			"brew",
		},
	}

	fmt.Println(styles.GreenText.Render("Checking for package managers..."))
	for _, packageManager := range packageManagers[runtime] {
		if system.CheckIfCommandExists(packageManager) {
			fmt.Printf("%s %s\n", styles.CyanText.Render("Found:"), styles.PurpleBoldText.Render(packageManager))
			userPackageManagers = append(userPackageManagers, packageManager)
		}
	}

	if len(userPackageManagers) == 0 {
		fmt.Println("Couldn't find any supported package manager.")
		os.Exit(1)
	}

	return userPackageManagers
}

func GetInstalledPackages(packageManagers []string) map[string][]string {
	installedPackages := map[string][]string{}

	for _, packageManager := range packageManagers {
		handler, ok := PackageManagerHandlers[packageManager]

		if ok {
			installedPackages[packageManager] = append(installedPackages[packageManager], handler()...)
		} else {
			fmt.Println("Unsupported package manager.")
			os.Exit(1)
		}
	}

	return installedPackages
}

func parseInstalledPackages(output string, regex string) []string {
	var packages []string

	re := regexp2.MustCompile(regex, regexp2.Multiline)
	m, _ := re.FindStringMatch(output)
	for m != nil {
		packages = append(packages, m.String())
		m, _ = re.FindNextMatch(m)
	}
	return packages
}
