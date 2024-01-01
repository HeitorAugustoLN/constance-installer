package system

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/HeitorAugustoLN/constance-installer/styles"
	"github.com/elastic/go-sysinfo"
)

func GetOperatingSystem() (string, string) {
	host, err := sysinfo.Host()
	userRuntime := sysinfo.Go().OS
	userOperatingSystem := host.Info().OS.Name

	if err != nil {
		fmt.Println(err)
	}

	return userRuntime, userOperatingSystem
}

func CheckOperatingSystem() {
	fmt.Printf("%s\n", styles.GreenText.Render("Checking operating system..."))
	switch userRuntime, userOperatingSystem := GetOperatingSystem(); userRuntime {
	case "windows", "darwin", "linux":
		fmt.Printf("%s %s\n", styles.CyanText.Render("Found user runtime:"), styles.PurpleBoldText.Render(userRuntime))
		fmt.Printf("%s %s\n", styles.CyanText.Render("Found user operating system:"), styles.PurpleBoldText.Render(userOperatingSystem))
		fmt.Printf("%s\n", styles.YellowText.Render("Supported operating system."))
	case "freebsd", "netbsd", "openbsd":
		fmt.Println("Not supported yet.")
		os.Exit(1)
	default:
		fmt.Println("Unknown operating system")
		os.Exit(1)
	}
}

func CheckIfCommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
