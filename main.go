// This package demonstrates the usage of Azure Storage services using Go.
package main

import (
	"flag"
	"fmt"
	
	// "github.com/Azure/azure-sdk-for-go/storage"
	"azure-sdk-for-go/storage" // referring to local one (dev) for now.
)

var (
	accountName string
	accountKey  string
	emulator    *bool
	blobCli     storage.BlobStorageClient
	tableCli     storage.TableServiceClient

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
}

func main() {
	fmt.Println("Azure Storage Blob Sample")
	blobSamples("demoblobcontainer", "demoPageBlob", "demoAppendBlob", "demoBlockBlob")
}
