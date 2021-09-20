package tfstate

import (
	"encoding/json"
	"io/ioutil"
)

func ParseTerraformStateFile(path string) (*stateV4, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return readStateV4(bytes)
}

func readStateV4(src []byte) (*stateV4, error) {
	sV4 := &stateV4{}
	err := json.Unmarshal(src, sV4)
	if err != nil {
		return nil, err
	}

	return sV4, nil
}

