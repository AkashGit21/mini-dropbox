package main

import (
	"github.com/AkashGit21/typeface-assignment/utils"
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

			StartServer(srv)
		},
	}

	rootCmd.AddCommand(runCmd)
}
