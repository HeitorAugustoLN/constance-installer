package packages

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/HeitorAugustoLN/constance-installer/styles"
	"github.com/HeitorAugustoLN/constance-installer/utils/system"
)

type GenInfo struct {
	Current int `json:"current"`
	Built   int `json:"built"`
}

func CreateGeneration() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}
	constanceFolder := filepath.Join(home, ".constance")
	profilesFolder := filepath.Join(constanceFolder, "profiles")
	importsFolder := filepath.Join(constanceFolder, "imports")
	generationsFolder := filepath.Join(constanceFolder, "generations")
	genInfoPath := filepath.Join(generationsFolder, "info.json")

	genInfo := system.ReadJSONFromFile(genInfoPath)
	currentGeneration := genInfo["current"].(float64) + 1

	currentGenerationFolder := filepath.Join(constanceFolder, "generations", fmt.Sprintf("%d", int(currentGeneration)))

	fmt.Printf("%s %s\n", styles.CyanText.Render("Creating generation:"), styles.PurpleBoldText.Render(fmt.Sprintf("%d", int(currentGeneration))))
	err = system.CopyDirectory(profilesFolder, fmt.Sprintf("%s/profiles", currentGenerationFolder))
	if err != nil {
		fmt.Println("Error copying directory:", err)
		return
	}
	err = system.CopyDirectory(importsFolder, fmt.Sprintf("%s/imports", currentGenerationFolder))
	if err != nil {
		fmt.Println("Error copying directory:", err)
		return
	}

	genInfoData := GenInfo{
		Current: int(currentGeneration),
		Built:   int(currentGeneration),
	}

	err = system.WriteJSONToFile(genInfoPath, genInfoData)

	if err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}
	fmt.Printf("%s\n", styles.YellowText.Render("Successfully created generation."))
}
