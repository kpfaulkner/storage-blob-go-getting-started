// This package demonstrates the usage of Azure Storage services using Go.
package main

import (
	"flag"
	"fmt"
	"os"	
	// "github.com/Azure/azure-sdk-for-go/storage"
	"azure-sdk-for-go/storage" // referring to local one (dev) for now.
)

var (
	accountName string
	accountKey  string
	emulator    *bool
	blobCli     storage.BlobStorageClient
	tableCli    storage.TableServiceClient
	queueCli    storage.QueueServiceClient
)

func init() {
	emulator = flag.Bool("emulator", false, "use the Azure Storage Emulator")
	flag.Parse()
	if *emulator {
		accountName = storage.StorageEmulatorAccountName
		accountKey = storage.StorageEmulatorAccountKey
	} else {
		accountName = getEnvVarOrExit("ACCOUNT_NAME")
		accountKey = getEnvVarOrExit("ACCOUNT_KEY")
	}
	client, err := storage.NewBasicClient(accountName, accountKey)
	onErrorFail(err, "Create client failed")

	blobCli = client.GetBlobService()
	tableCli = client.GetTableService()
	queueCli = client.GetQueueService()
}

func main() {
	fmt.Println("Azure Storage Blob Sample")
	blobSamples("demoblobcontainer", "demoPageBlob", "demoAppendBlob", "demoBlockBlob")
	tableSamples("demotable")
	queueSamples("demoqueue")
}

// getEnvVarOrExit returns the value of specified environment variable or terminates if it's not defined.
func getEnvVarOrExit(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		fmt.Printf("Missing environment variable %s\n", varName)
		os.Exit(1)
	}

	return value
}

// onErrorFail prints a failure message and exits the program if err is not nil.
func onErrorFail(err error, message string) {
	if err != nil {
		fmt.Printf("%s: %s\n", message, err)
		os.Exit(1)
	}
}

