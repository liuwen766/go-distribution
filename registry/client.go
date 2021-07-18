package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/**
 * @desc: 注册服务
 * @data: 2021.7.18 22:46
 */
func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServiceURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to register service. "+
			"Register service responded with code %v \n", res.StatusCode,
		)
	}
	return nil
}

/**
 * @desc: 取消服务
 * @data: 2021.7.18 22:46
 */
func ShutdownService(url string) error {
	request, err := http.NewRequest(http.MethodDelete, ServiceURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	request.Header.Add("Content-type", "text/plain")
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registry Service responed with code:%s", res.StatusCode)
	}
	return nil
}
