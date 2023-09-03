package main

import (
	"log"
	"time"

	"github.com/AkashGit21/typeface-assignment/utils"
	"github.com/go-co-op/gocron"
	"github.com/spf13/cobra"
)

func init() {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Starts running the application server",
		Run: func(cmd *cobra.Command, args []string) {
			srv, err := NewServer()
			if err != nil {
				utils.ErrorLog("Error getting new server:", err)
				return
			}

			s := gocron.NewScheduler(time.Local)
			_, _ = s.Cron("30 1 * * *").Do(func() {
				log.Println("Cron runs at 1:30 AM every night asynchronously")
				metaOps, err := utils.NewPersistenceDBLayer()
				if err != nil {
					utils.ErrorLog("Error getting new persistence db layer:", err)
					return
				}
				metaOps.DeleteRecords()
			}) // every day at 1:30 am
			s.StartAsync()

			StartServer(srv)
		},
	}

	rootCmd.AddCommand(runCmd)
}
