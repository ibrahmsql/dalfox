package cmd

import (
	"github.com/hahwul/dalfox/v2/internal/printing"
	"github.com/hahwul/dalfox/v2/pkg/server"
	"github.com/spf13/cobra"
)

// Command-line flags for server configuration
var port int                 // Port to bind the server to
var host, server_type string // Host address and server type

// serverCmd represents the server command for starting API servers
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start API Server",
	Run:   runServerCmd,
}

// runServerCmd handles execution of the server command
// It starts either a REST API or MCP server based on the configured type
func runServerCmd(cmd *cobra.Command, args []string) {
	printing.Banner(options)
	options.ServerHost = host
	options.ServerPort = port

	switch server_type {
	case "mcp":
		printing.DalLog("SYSTEM", "Starting MCP Server", options)
		printing.Summary(options, "MCP Server Mode")
		server.RunMCPServer(options)
	default:
		printing.DalLog("SYSTEM", "Starting REST API Server", options)
		printing.Summary(options, "REST API Mode")
		server.RunAPIServer(options)
	}
}

// init registers the server command and its flags
func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVar(&port, "port", 6664, "Specify the port to bind the server to. Example: --port 6664")
	serverCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Specify the address to bind the server to. Example: --host '0.0.0.0'")
	serverCmd.Flags().StringVar(&server_type, "type", "rest", "Specify the server type. Example: --type 'rest' or --type 'mcp'")

	// Apply custom help format to this subcommand
	ApplySubCommandCustomHelp(serverCmd)
}
