package compare

import (
	"fmt"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
	"gopkg.in/yaml.v3"
)

var cfgPath = "compare/cfg.yml"

type file struct {
	Type string `yaml:"type"`
}
type FileCfg interface {
	GetType() string
}

func (file xmlCfg) GetType() string  { return file.Type }
func (file xlsxCfg) GetType() string { return file.Type }

type xmlCfg struct {
	Type    string `yaml:"type"`
	Path    string `yaml:"path"`
	Keyword string `yaml:"keyword"`
}

type xlsxCfg struct {
	Type  string `yaml:"type"`
	Path  string `yaml:"path"`
	Sheet string `yaml:"sheet"`
	Col   int    `yaml:"col"`
}

type XlsxFiles struct {
	files map[string]xlsxCfg `yaml:",inline"`
}

func unmarshalCfg(filePath string) (XlsxFiles, error) {
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

func getXlsxColData(xlsxCfg xlsxCfg) ([]string, error) {
	f, err := excelize.OpenFile(xlsxCfg.Path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", xlsxCfg.Path, err)
	}
	defer f.Close()

	rows, err := f.GetRows(xlsxCfg.Sheet)
	if err != nil {
		return nil, fmt.Errorf("failed to get rows from sheet %s: %w", xlsxCfg.Sheet, err)
	}
	var columnData []string
	for _, row := range rows {
		if len(row) > xlsxCfg.Col {
			columnData = append(columnData, row[xlsxCfg.Col])
		}
	}
	return columnData, err
}

func InitCompare() error {
	cfg, err := unmarshalCfg(cfgPath)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", cfg.files)

	file1, err := getXlsxColData(cfg.files["xlsxFile1"])
	if err != nil {
		return err
	}

	file2, err := getXlsxColData(cfg.files["xlsxFile2"])
	if err != nil {
		return err
	}

	limit := 10
	if len(file1) < limit {
		limit = len(file1)
	}
	fmt.Println(file1[:limit])

	if len(file2) < limit {
		limit = len(file2)
	}
	fmt.Println(file2[:limit])

	return nil
}

func main() {
	if err := InitCompare(); err != nil {
		log.Fatal(err)
	}
}
