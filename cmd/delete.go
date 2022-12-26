package cmd

import (
	"fmt"
	"github.com/ledboot/argo-man/config"
	"github.com/ledboot/logger"
	"net/http"
	"time"
)

func deleteAll() {
	_, err := getArgoToken()
	if err != nil {
		panic(err)
	}
	if len(config.Cfg.AppList) == 0 {
		panic("app list is empty")
	}
	for name, _ := range config.Cfg.AppList {
		logger.Infof("delete app: %s", name)
		doDelete(name)
	}
}

func doDelete(name string) {
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/v1/applications/%s", config.Cfg.ArgoConfig.Host, name), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", argoToken))
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if !(resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusOK) {
		panic(fmt.Sprintf("delete app %s failed, status code: %d", name, resp.StatusCode))
	}
}
