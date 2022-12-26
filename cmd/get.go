package cmd

import (
	"fmt"
	"github.com/ledboot/argo-man/config"
	"net/http"
	"time"
)

func getApp(name string) int {
	client := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/applications/%s", config.Cfg.ArgoConfig.Host, name), nil)
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
	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("app %s not found\n", name)
		return -1
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("get app %s failed, status code: %d\n", name, resp.StatusCode)
		return 1
	}
	return 0
}
