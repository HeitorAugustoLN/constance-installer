package packages

import (
	"fmt"
	"os/exec"

	"github.com/HeitorAugustoLN/constance-installer/styles"
)

type PackageManagerHandler func() []string

var PackageManagerHandlers = map[string]PackageManagerHandler{
	"scoop":  getScoopPackages,
	"choco":  getChocoPackages,
	"winget": getWingetPackages,

	"brew": getBrewPackages,

	"apt":          getAptPackages,
	"dnf":          getDnfPackages,
	"zypper":       getZypperPackages,
	"pacman":       getPacmanPackages,
	"xbps-install": getXbpsPackages,

	"aura":   getAurPackages,
	"aurman": getAurPackages,
	"pacaur": getAurPackages,
	"pakku":  getAurPackages,
	"paru":   getAurPackages,
	"pikaur": getAurPackages,
	"trizen": getAurPackages,
	"yay":    getAurPackages,

	"snap":    getSnapPackages,
	"flatpak": getFlatpakPackages,
}

func getScoopPackages() []string {
	command := []string{"scoop", "list"}
	regex := `^(?!Installed|Name|----)\w+([-.]\w+)*`

	fmt.Printf("%s\n", styles.GreenText.Render("Checking for packages in scoop..."))
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getChocoPackages() []string {
	command := []string{"choco", "list"}
	regex := `^[\w.+-]+(?=\s+\d+\.\d+\.\d+)`

	fmt.Printf("%s\n", styles.GreenText.Render("Checking for packages in choco..."))
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getWingetPackages() []string {
	commands := [][]string{
		{"powershell", "-Command", "(winget list) -match ' winget$'"},
		{"powershell", "-Command", "(winget list) -match ' msstore$'"},
	}
	regexes := []string{
		`(?<=\s)[\w+-.]*\.[\w+-.]+(?=\s+\d)`,
		`(?<=\s)[\w+-.]+(?=\s+\d+\.\d+\.\d+)`,
	}
	var packages []string

	fmt.Printf("%s\n", styles.GreenText.Render("Checking for packages in winget..."))
	for i, command := range commands {
		regex := regexes[i]
		cmd := exec.Command(command[0], command[1:]...)
		cmd.Stdout = nil
		cmd.Stderr = nil
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Couldn't get installed packages: %v\n", err)
			return nil
		}

		packages = append(packages, parseInstalledPackages(string(output), regex)...)
	}

	return packages
}

func getBrewPackages() []string {
	command := []string{"brew", "list", "--formulae", "-1"}
	regex := `^[\w.+-]+$`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getAptPackages() []string {
	command := []string{"apt-mark", "showmanual"}
	regex := `^[\w.+-]+$`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getDnfPackages() []string {
	command := []string{"dnf", "repoquery", "--userinstalled"}
	regex := `^([\w+.-]+?)(?=-\d)`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getZypperPackages() []string {
	/*
		Zypper support is kinda tricky since it display the installed packages in a table format.
		It is super hard to handle this with regex, because the regex matches some versions too
		The way I found to handle this is making match have a letter in the beginning of the package name with [A-Za-z]{1}
		But this is not a good solution, since it will not match packages that don't have a letter in the beginning of the name
		But I found only 3 packages in the openSUSE repositories that don't have a letter in the beginning of the name
		So it is not a big deal, but it is not a good solution either

		If you have a better solution, please open a pull request
	*/
	command := []string{"zypper", "pa", "-i", "|", "grep", "i+"}
	regex := `(?<=\|\s)[A-Za-z]{1}[\w+-._]+(?=\s+\|)`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getPacmanPackages() []string {
	command := []string{"pacman", "-Qen"}
	regex := `^[\w.+-]+`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getXbpsPackages() []string {
	command := []string{"xbps-query", "-m"}
	regex := `^[\w.+-]+(?=-\d)`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getAurPackages() []string {
	command := []string{"pacman", "-Qem"}
	regex := `^[\w.+-]+`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getSnapPackages() []string {
	command := []string{"snap", "list"}
	regex := `^(?!Name\s)\S+`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}

func getFlatpakPackages() []string {
	command := []string{"flatpak", "list", "--app", "--columns=application"}
	regex := `^[\w+-.]*\.[\w+-.]+`

	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Couldn't get installed packages: %v\n", err)
		return nil
	}

	return parseInstalledPackages(string(output), regex)
}
