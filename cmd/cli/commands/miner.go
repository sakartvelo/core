package commands

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"

	ds "github.com/c2h5oh/datasize"
	pb "github.com/sonm-io/core/proto"
)

func init() {
	minerRootCmd.AddCommand(minersListCmd, minerStatusCmd)
}

func printWorkerList(cmd *cobra.Command, lr *pb.ListReply) {
	if isSimpleFormat() {
		if len(lr.Info) == 0 {
			cmd.Printf("No workers connected\r\n")
			return
		}

		for addr, meta := range lr.Info {
			cmd.Printf("Worker: %s", addr)

			taskCount := len(meta.Values)
			if taskCount == 0 {
				cmd.Printf("\t\tIdle\r\n")
			} else {
				cmd.Printf("\t\t%d active task(s)\r\n", taskCount)
			}
		}
	} else {
		b, _ := json.Marshal(lr)
		cmd.Println(string(b))
	}
}

func printCpuInfo(cmd *cobra.Command, cap *pb.Capabilities) {
	for i, cpu := range cap.Cpu {
		cmd.Printf("    CPU%d: %d x %s\r\n", i, cpu.GetCores(), cpu.GetModelName())
	}
}

func printGpuInfo(cmd *cobra.Command, cap *pb.Capabilities) {
	if len(cap.Gpu) > 0 {
		for i, gpu := range cap.Gpu {
			cmd.Printf("    GPU%d: %s %s\r\n", i, gpu.VendorName, gpu.Name)
		}
	} else {
		cmd.Println("    GPU: None")
	}
}

func printMemInfo(cmd *cobra.Command, cap *pb.Capabilities) {
	cmd.Println("    RAM:")
	cmd.Printf("      Total: %s\r\n", ds.ByteSize(cap.Mem.GetTotal()).HR())
	cmd.Printf("      Used:  %s\r\n", ds.ByteSize(cap.Mem.GetUsed()).HR())
}

func printWorkerStatus(cmd *cobra.Command, workerID string, metrics *pb.InfoReply) {
	if isSimpleFormat() {
		cmd.Printf("Worker \"%s\":\r\n", workerID)

		if metrics.Capabilities != nil {
			cmd.Println("  Hardware:")
			printCpuInfo(cmd, metrics.Capabilities)
			printGpuInfo(cmd, metrics.Capabilities)
			printMemInfo(cmd, metrics.Capabilities)
		}

		if len(metrics.GetUsage()) == 0 {
			cmd.Println("  No active tasks")
		} else {
			cmd.Println("  Tasks:")
			i := 1
			for task := range metrics.Usage {
				cmd.Printf("    %d) %s\r\n", i, task)
				i++
			}
		}
	} else {
		b, _ := json.Marshal(metrics)
		cmd.Println(string(b))
	}
}

var minerRootCmd = &cobra.Command{
	Use:     "miner",
	Short:   "Operations with miners",
	PreRunE: checkHubAddressIsSet,
}

var minersListCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show connected miners",
	PreRunE: minerRootCmd.PreRunE,
	Run: func(cmd *cobra.Command, args []string) {
		itr, err := NewGrpcInteractor(hubAddress, timeout)
		if err != nil {
			showError(cmd, "Cannot connect to hub", err)
			return
		}

		minerListCmdRunner(cmd, itr)
	},
}

var minerStatusCmd = &cobra.Command{
	Use:     "status <miner_addr>",
	Short:   "Miner status",
	PreRunE: checkHubAddressIsSet,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errWorkerIDRequired
		}
		minerID := args[0]

		itr, err := NewGrpcInteractor(hubAddress, timeout)
		if err != nil {
			showError(cmd, "Cannot connect to hub", err)
			return nil
		}

		minerStatusCmdRunner(cmd, minerID, itr)
		return nil
	},
}

func minerListCmdRunner(cmd *cobra.Command, interactor CliInteractor) {
	list, err := interactor.MinerList(context.Background())
	if err != nil {
		showError(cmd, "Cannot get miners list", err)
		return
	}

	printWorkerList(cmd, list)
}

func minerStatusCmdRunner(cmd *cobra.Command, minerID string, interactor CliInteractor) {
	metrics, err := interactor.MinerStatus(minerID, context.Background())
	if err != nil {
		showError(cmd, "Cannot get miner status", err)
		return
	}

	printWorkerStatus(cmd, minerID, metrics)
}
