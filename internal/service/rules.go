package service

import (
	"log"
	"path/filepath"
	"smart-file-organizer/internal/repository"
	"strings"
)

type RuleService struct {
	rules map[string][]string
	repo  repository.FileRepo
}

func NewRuleService(rules map[string][]string, repo repository.FileRepo) *RuleService {
	return &RuleService{rules: rules, repo: repo}
}

func (s *RuleService) ApplyRules(path string) {
	ext := strings.ToLower(filepath.Ext(path))
	for category, exts := range s.rules {
		for _, ruleExt := range exts {
			if ext == ruleExt {
				err := s.repo.Move(path, category)
				if err != nil {
					log.Printf("[ERROR] Failed to move %s → %s: %v", path, category, err)
				} else {
					log.Printf("[INFO] Moved %s → %s", path, category)
				}
				return
			}
		}
	}
}

func (s *RuleService) FindCategory(ext string) string {
	ext = strings.ToLower(ext)
	for category, exts := range s.rules {
		for _, ruleExt := range exts {
			if ext == ruleExt {
				return category
			}
		}
	}
	return ""
}
