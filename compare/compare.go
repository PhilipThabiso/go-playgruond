package compare

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var cfgPath = "compare/cfg.yml"

type XlsxFile struct {
	Path  string `yaml:"path"`
	Sheet string `yaml:"sheet"`
	Col   int    `yaml:"col"`
}

type XlsxFiles struct {
	files map[string]XlsxFile `yaml:",inline"`
}

func unmarshalXlsx(filePath string) (XlsxFiles, error) {
	var files XlsxFiles

	data, err := os.ReadFile(filePath)
	if err != nil {
		return files, fmt.Errorf("could not read file %s: %w", filePath, err)
	}

	err = yaml.Unmarshal(data, &files.files)
	if err != nil {
		return files, fmt.Errorf("could not Unmarshal: %w", err)
	}
	return files, err
}

func InitCompare() {
	cfg, err := unmarshalXlsx(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", cfg.files)
}

func main() {
	InitCompare()
}
