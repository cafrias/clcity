package dump

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Read reads from a dump file, returns false if it can read it doesn't matter what reason.
func Read(v interface{}, outFile string) bool {
	if _, err := os.Stat(outFile); os.IsNotExist(err) {
		return false
	}

	c, err := ioutil.ReadFile(outFile)
	if err != nil {
		return false
	}

	err = json.Unmarshal(c, &v)
	if err != nil {
		return false
	}

	return true
}

// Write writes a dump file.
func Write(v interface{}, outFile string) error {
	fcontent, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outFile, fcontent, 0644)
	if err != nil {
		return err
	}

	return nil
}
