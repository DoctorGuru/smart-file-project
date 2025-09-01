package app

import (
	"log"
	"os"
	"path/filepath"
	"smart-file-organizer/internal/repository"
	"smart-file-organizer/internal/service"
	"smart-file-organizer/pkg/config"
	"time"

	"github.com/fsnotify/fsnotify"
)

func Run(cfg *config.Config) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	repo := repository.NewFileRepo(cfg.WatchDir)
	ruleService := service.NewRuleService(cfg.Rules, repo)

	err = watcher.Add(cfg.WatchDir)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Smart File Organizer watching:", cfg.WatchDir)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				go func(f string) {
					time.Sleep(1 * time.Second) // wait for write
					ruleService.ApplyRules(f)
				}(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

// OrganizeOnce → goes through all existing files and organizes
func OrganizeOnce(cfg *config.Config) int {
	repo := repository.NewFileRepo(cfg.WatchDir)
	ruleService := service.NewRuleService(cfg.Rules, repo)

	files, _ := os.ReadDir(cfg.WatchDir)
	count := 0
	for _, f := range files {
		if !f.IsDir() {
			path := filepath.Join(cfg.WatchDir, f.Name())
			ruleService.ApplyRules(path)
			count++
		}
	}
	return count
}

// DryRun → shows what *would* be moved
func DryRun(cfg *config.Config) int {
	repo := repository.NewFileRepo(cfg.WatchDir)
	ruleService := service.NewRuleService(cfg.Rules, repo)

	files, _ := os.ReadDir(cfg.WatchDir)
	count := 0
	for _, f := range files {
		if !f.IsDir() {
			path := filepath.Join(cfg.WatchDir, f.Name())
			ext := filepath.Ext(path)
			category := ruleService.FindCategory(ext)
			if category != "" {
				log.Printf("Would move %s → %s", f.Name(), category)
				count++
			}
		}
	}
	return count
}
