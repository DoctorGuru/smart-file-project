package transport

import (
	"fmt"
	"log"
	"smart-file-organizer/internal/app"
	"smart-file-organizer/pkg/config"

	"github.com/spf13/cobra"
)

func NewCLI() *cobra.Command {
	var cfgPath string

	rootCmd := &cobra.Command{
		Use:   "organizer",
		Short: "Smart File Organizer - organize your files automatically",
	}

	// Add global flag for config path
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "configs/config.yaml", "Path to config file")

	// `run` command -> starts watching directory
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Start watching directory and organizing files",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				log.Fatal("Failed to load config:", err)
			}
			app.Run(cfg)
		},
	}

	// `once` command -> organize all files once
	onceCmd := &cobra.Command{
		Use:   "once",
		Short: "Organize existing files once (no watching)",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				log.Fatal("Failed to load config:", err)
			}
			count := app.OrganizeOnce(cfg)
			fmt.Printf("✅ Organized %d files\n", count)
		},
	}

	// `dry` command -> simulate organization
	dryCmd := &cobra.Command{
		Use:   "dry",
		Short: "Simulate organizing without moving files",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.Load(cfgPath)
			if err != nil {
				log.Fatal("Failed to load config:", err)
			}
			count := app.DryRun(cfg)
			fmt.Printf("🔎 Dry run complete: %d files would be moved\n", count)
		},
	}

	rootCmd.AddCommand(runCmd, onceCmd, dryCmd)
	return rootCmd
}
