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

	// delete the table first, incas it exists	
	tableRef.Delete(30, nil)

	// create new table
	err := tableRef.Create(30, storage.EmptyPayload, nil)
	if err != nil {
		onErrorFail(err, "Create table failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = addEntityToTable( tableRef )	
	if err != nil {
		onErrorFail(err, "addEntityToTable failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = updateEntityForTable( tableRef )	
	if err != nil {
		onErrorFail(err, "updateEntityForTable failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = getEntityFromTable( tableRef )	
	if err != nil {
		onErrorFail(err, "getEntityFromTable failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}


	fmt.Println("Done")
}

func addEntityToTable( table *storage.Table) error {
	fmt.Println("Create an entity to table...")
	
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

// updateEntityForTable creates an entity then updates it.
func updateEntityForTable( table *storage.Table) error {

	fmt.Println("Update an entity to table...")
	
	entity := table.GetEntityReference("partitionkey1", "rowkey2")
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

	// update a number but also introduce a new property
	props = map[string]interface{}{
		"SomeNumber":      234,
		"SomeNewString":  "some new string",
	}
	
	entity.Properties = props
	err = entity.Update( true, nil)
	if err != nil {
		return err
	}

	return nil
}

// getEntityFromTable creates an entity then updates it.
func getEntityFromTable( table *storage.Table) error {

	fmt.Println("Get an entity from table...")
	
	entity := table.GetEntityReference("partitionkey1", "rowkey3")
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

	queryOptions := storage.QueryOptions{
		Filter: "RowKey eq 'rowkey3'",
	}

	entities, err := table.QueryEntities(30, storage.FullMetadata, &queryOptions)
	if err != nil {
		return err
	}

	if len(entities.Entities) != 1 {
		return fmt.Errorf("Could not retrieve entity")
	}

	fmt.Printf("Have entity with number %v\n", entities.Entities[0].Properties["SomeNumber"])
	
	return nil
}



