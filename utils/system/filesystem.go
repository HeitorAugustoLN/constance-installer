package system

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HeitorAugustoLN/constance-installer/styles"
)

func CreateFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, 0755)
	}
}

func CreateFile(fileName string) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		os.Create(fileName)
	}
}

func BootstrapConstanceFiles() {
	operatingSystem, _ := GetOperatingSystem()
	home, _ := os.UserHomeDir()

	constanceFolder := filepath.Join(home, ".constance")
	subfolders := []string{"imports", "profiles", "generations"}
	files := []string{"config.json", fmt.Sprintf("profiles/%s.json", operatingSystem), "generations/info.json", fmt.Sprintf("imports/core_%s.json", operatingSystem)}

	fmt.Printf("%s %s\n", styles.CyanText.Render("Creating folder:"), styles.PurpleBoldText.Render(constanceFolder))
	CreateFolder(constanceFolder)

	for _, subfolder := range subfolders {
		subfolderPath := filepath.Join(constanceFolder, subfolder)
		fmt.Printf("%s %s\n", styles.CyanText.Render("Creating folder:"), styles.PurpleBoldText.Render(subfolderPath))

		CreateFolder(subfolderPath)
	}
	for _, file := range files {
		filePath := filepath.Join(constanceFolder, file)
		fmt.Printf("%s %s\n", styles.CyanText.Render("Creating file:"), styles.PurpleBoldText.Render(filePath))

		CreateFile(filePath)
	}
}

func CheckIfProgramAlreadySetup() bool {
	home, _ := os.UserHomeDir()
	constanceFolder := filepath.Join(home, ".constance")
	_, err := os.Stat(constanceFolder)
	return !os.IsNotExist(err)
}

type PackageManager struct {
	Winget  *[]string `json:"winget,omitempty"`
	Scoop   *[]string `json:"scoop,omitempty"`
	Choco   *[]string `json:"choco,omitempty"`
	Apt     *[]string `json:"apt,omitempty"`
	Snap    *[]string `json:"snap,omitempty"`
	Flatpak *[]string `json:"flatpak,omitempty"`
	Pacman  *[]string `json:"pacman,omitempty"`
	Dnf     *[]string `json:"dnf,omitempty"`
	Xbps    *[]string `json:"xbps,omitempty"`
	Zypper  *[]string `json:"zypper,omitempty"`
	Brew    *[]string `json:"brew,omitempty"`
	Aura    *[]string `json:"aura,omitempty"`
	Aurman  *[]string `json:"aurman,omitempty"`
	Pacaur  *[]string `json:"pacaur,omitempty"`
	Pakku   *[]string `json:"pakku,omitempty"`
	Yay     *[]string `json:"yay,omitempty"`
	Trizen  *[]string `json:"trizen,omitempty"`
	Paru    *[]string `json:"paru,omitempty"`
	Pikaur  *[]string `json:"pikaur,omitempty"`
}

type PackageData struct {
	Imports  *[]string       `json:"imports,omitempty"`
	Packages *PackageManager `json:"packages,omitempty"`
}

type GenInfo struct {
	Current int `json:"current"`
	Built   int `json:"built"`
}

type Config struct {
	Profile string `json:"profile"`
}

func writeJSONToFile(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ") // 4 spacing indent

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func ParseJSONToFiles(packages map[string][]string) {
	runtime, _ := GetOperatingSystem()
	home, _ := os.UserHomeDir()
	constanceFolder := filepath.Join(home, ".constance")

	packageData := PackageData{
		Packages: &PackageManager{},
	}

	mainFileData := PackageData{
		Imports: &[]string{fmt.Sprintf("core_%s", runtime)},
	}

	genInfo := GenInfo{
		Current: 0,
		Built:   0,
	}

	config := Config{
		Profile: runtime,
	}

	for packageManager, packageList := range packages {
		currentPackageList := make([]string, len(packageList))
		copy(currentPackageList, packageList)

		switch packageManager {
		case "winget":
			packageData.Packages.Winget = &currentPackageList
		case "scoop":
			packageData.Packages.Scoop = &currentPackageList
		case "choco":
			packageData.Packages.Choco = &currentPackageList
		case "apt":
			packageData.Packages.Apt = &currentPackageList
		case "snap":
			packageData.Packages.Snap = &currentPackageList
		case "flatpak":
			packageData.Packages.Flatpak = &currentPackageList
		case "pacman":
			packageData.Packages.Pacman = &currentPackageList
		case "dnf":
			packageData.Packages.Dnf = &currentPackageList
		case "xbps":
			packageData.Packages.Xbps = &currentPackageList
		case "zypper":
			packageData.Packages.Zypper = &currentPackageList
		case "brew":
			packageData.Packages.Brew = &currentPackageList
		case "aura":
			packageData.Packages.Aura = &currentPackageList
		case "aurman":
			packageData.Packages.Aurman = &currentPackageList
		case "pacaur":
			packageData.Packages.Pacaur = &currentPackageList
		case "pakku":
			packageData.Packages.Pakku = &currentPackageList
		case "yay":
			packageData.Packages.Yay = &currentPackageList
		case "trizen":
			packageData.Packages.Trizen = &currentPackageList
		case "paru":
			packageData.Packages.Paru = &currentPackageList
		case "pikaur":
			packageData.Packages.Pikaur = &currentPackageList
		}
	}
	coreSysPath := filepath.Join(constanceFolder, fmt.Sprintf("imports/core_%s.json", runtime))
	mainPath := filepath.Join(constanceFolder, fmt.Sprintf("profiles/%s.json", runtime))
	genInfoPath := filepath.Join(constanceFolder, "generations/info.json")
	configPath := filepath.Join(constanceFolder, "config.json")

	fmt.Printf("%s %s\n", styles.CyanText.Render("Writing JSON to file:"), styles.PurpleBoldText.Render(coreSysPath))
	err := writeJSONToFile(coreSysPath, packageData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}
	fmt.Printf("%s %s\n", styles.CyanText.Render("Writing JSON to file:"), styles.PurpleBoldText.Render(mainPath))
	err = writeJSONToFile(mainPath, mainFileData)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}
	fmt.Printf("%s %s\n", styles.CyanText.Render("Writing JSON to file:"), styles.PurpleBoldText.Render(genInfoPath))
	err = writeJSONToFile(genInfoPath, genInfo)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}
	fmt.Printf("%s %s\n", styles.CyanText.Render("Writing JSON to file:"), styles.PurpleBoldText.Render(configPath))
	err = writeJSONToFile(configPath, config)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
	}
}

func ReadJSONFromFile(filePath string) map[string]interface{} {
	var result map[string]interface{}
	byteValue, _ := os.ReadFile(filePath)
	json.Unmarshal(byteValue, &result)

	return result
}
