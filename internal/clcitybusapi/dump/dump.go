package dump

import (
	"encoding/json"
	"io/ioutil"
)

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
