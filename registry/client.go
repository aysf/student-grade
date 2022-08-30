package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "json/application", buf)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. responded with code %v", res.StatusCode)
	}
	return err
}
