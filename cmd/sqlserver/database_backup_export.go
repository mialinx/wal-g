package sqlserver

import (
	"github.com/spf13/cobra"
	"github.com/wal-g/wal-g/internal/databases/sqlserver"
)

const databaseBackupExportShortDescription = "Exports one database backup to file or URL"

var databaseBackupExportCmd = &cobra.Command{
	Use:   "database-backup-export",
	Short: databaseBackupExportShortDescription,
	Run: func(cmd *cobra.Command, args []string) {
		sqlserver.HandleDatabaseBackupExport(args[0])
	},
}

func init() {
	cmd.AddCommand(databaseBackupExportCmd)
}
