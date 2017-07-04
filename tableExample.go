// This package demonstrates the usage of Azure Table services using Go.
package main

import (
	"fmt"
	"time"

	// "github.com/Azure/azure-sdk-for-go/storage"
	"azure-sdk-for-go/storage" // referring to local one (dev) for now.
)

// tableSamples creates a table, populates rows, retrieved based on partition and row key,
// updates a row and deletes a row.
func tableSamples(tableName string) {
	
	fmt.Println("Create table")
	tableRef := tableCli.GetTableReference(tableName)
	err := tableRef.Create(30, storage.EmptyPayload, nil)
	if err != nil {
		onErrorFail(err, "Create table failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = addEntityToTable( tableRef )	
	if err != nil {
		onErrorFail(err, "AddRowToTable failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	fmt.Println("Done")
}

func addEntityToTable( table *storage.Table) error {

	entity := table.GetEntityReference("partitionkey1", "rowkey1")
	props := map[string]interface{}{
		"SomeNumber":      123,
		"SomeString":  "some string",
		"SomeDate":  time.Date(1992, time.December, 20, 21, 55, 0, 0, time.UTC),
		"IsActive":       true,
	}
	entity.Properties = props
	err := entity.Insert(storage.EmptyPayload, nil)
	if err != nil {
		return err
	}

	return nil
}
