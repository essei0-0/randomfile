package randomfile_test

import (
	"testing"

	"github.com/essei0-0/randomfile"
)

func TestGenerateFile(t *testing.T) {
	config := randomfile.Config{MinSize: 100, MaxSize: 100}
	fileTypes := []randomfile.FileType{randomfile.TXT}

	for _, fileType := range fileTypes {
		fileName, err := randomfile.GenerateFile(fileType, config)
		if err != nil {
			t.Errorf("failed to generate %s file: %s", fileType, err)
		}

		// Check that the generated file has the expected extension
		if fileName[len(fileName)-len(string(fileType))-1:] != "."+string(fileType) {
			t.Errorf("generated file has wrong extension for %s file", fileType)
		}
	}
}
