package filemanager

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Filemanager struct {
	InputFileName  string
	OutputFileName string
}

func (fm *Filemanager) ReadFile() ([]string, error) {

	file, err := os.ReadFile(fm.InputFileName)

	if err != nil {
		fmt.Println("error")
		return nil, fmt.Errorf("unable to read")
	}
	str := string(file)

	lines := strings.Split(str, "\n")
	return lines, nil

}

func (fm *Filemanager) WriteJSON(data interface{}) error {
	arr, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err

	}
	time.Sleep(time.Second * 3)
	err = os.WriteFile(fm.OutputFileName, arr, os.FileMode(0644))
	if err != nil {
		return err

	}
	return nil
}

func New(inputFileName string, outputFileName string) Filemanager {
	return Filemanager{
		InputFileName:  inputFileName,
		OutputFileName: outputFileName,
	}
}
