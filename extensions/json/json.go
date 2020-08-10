package json

import "encoding/json"

func ToJSON(t interface{}) ([]byte, error) {
	result, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func FromJSON(t interface{}, data []byte) error {
	err := json.Unmarshal(data, &t)

	if err != nil {
		return err
	}

	return nil
}
