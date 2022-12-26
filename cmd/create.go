package cmd

import (
	"bytes"
	"fmt"
	"github.com/ledboot/argo-man/config"
	"github.com/ledboot/logger"
	"net/http"
	"time"
)

func createApp() {
	_, err := getArgoToken()
	if err != nil {
		panic(err)
	}
	if len(config.Cfg.AppList) == 0 {
		panic("app list is empty")
	}
	for name, info := range config.Cfg.AppList {
		logger.Infof("create app: %s,info=%v", name, info)
		if code := getApp(name); code == -1 {
			doCreate(name, info)
		}
	}
}

func doCreate(name string, info config.AppInfo) {
	client := http.Client{Timeout: 5 * time.Second}

	playLoad := fmt.Sprintf(`{"metadata":{"name":"%s"},"spec":{"destination":{"name":"%s","namespace":"%s","server":"https://kubernetes.default.svc"},"project":"%s","source":{"path":"%s","repoURL":"%s","targetRevision":"HEAD"},"syncPolicy":{"automated":{"prune":true}}}}`,
		name, name, info.Namespace, info.Namespace, info.SourcePath, config.Cfg.RepoUrl)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/applications", config.Cfg.ArgoConfig.Host), bytes.NewBuffer([]byte(playLoad)))
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
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("create app %s failed, status code: %d", name, resp.StatusCode))
	}
}
