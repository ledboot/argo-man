package cmd

import (
	"bytes"
	"fmt"
	"github.com/ledboot/argo-man/config"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"time"
)

var argoToken string

func getArgoToken() (string, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	playLoad := fmt.Sprintf(`{"username":"%s","password":"%s"}`, config.Cfg.ArgoConfig.Username, config.Cfg.ArgoConfig.Password)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/session", config.Cfg.ArgoConfig.Host), bytes.NewBuffer([]byte(playLoad)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("getArgoToken, status code: %d", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	argoToken = gjson.GetBytes(data, "token").String()
	return argoToken, nil
}
