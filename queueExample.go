// This package demonstrates the usage of Azure Queue services using Go.
package main

import (
	"fmt"
	
	// "github.com/Azure/azure-sdk-for-go/storage"
	"azure-sdk-for-go/storage" // referring to local one (dev) for now.
)

// queueSamples creates a queue, pushes messages, pops messages
func queueSamples(queueName string) {

	fmt.Println("Create queue")
	queueRef := queueCli.GetQueueReference(queueName)

	// delete the queue first, incas it exists	
	queueRef.Delete(nil)

	// create new queue
	err := queueRef.Create(nil)
	if err != nil {
		onErrorFail(err, "Create queue failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = addMessageToQueue( queueRef )	
	if err != nil {
		onErrorFail(err, "addMessageToQueue failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}


	err = getMessageFromQueue( queueRef )	
	if err != nil {
		onErrorFail(err, "getMessageFromQueue failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}

	err = deleteMessageFromQueue( queueRef )	
	if err != nil {
		onErrorFail(err, "deleteMessageFromQueue failed: If you are running with the emulator credentials, plaase make sure you have started the storage emmulator. Press the Windows key and type Azure Storage to select and run it from the list of applications - then restart the sample")
	}


	fmt.Println("Done")
}

func addMessageToQueue( queue *storage.Queue) error {
	fmt.Println("Add a message to the queue...")
	m := queue.GetMessageReference("my message data")
	err := m.Put(nil)
	if err != nil {
		return err
	}
	
	return nil
}

func getMessageFromQueue( queue *storage.Queue) error {
	fmt.Println("Get a message from the queue...")

	// empty queue.
	queue.ClearMessages(nil)
	
	// add message
	m := queue.GetMessageReference("my message data")
	err := m.Put(nil)
	if err != nil {
		return err
	}
	
	options := storage.GetMessagesOptions{
		NumOfMessages: 1,
	}
	list, err := queue.GetMessages(&options)
	if err != nil {
		return err
	}
	
	fmt.Printf("Have message %s\n", list[0].Text)
	return nil
}

func deleteMessageFromQueue( queue *storage.Queue) error {
	fmt.Println("Delete a message from the queue...")

	// empty queue.
	queue.ClearMessages(nil)
	
	// add message
	m := queue.GetMessageReference("my message data")
	err := m.Put(nil)
	if err != nil {
		return err
	}
	
	options := storage.GetMessagesOptions{
		NumOfMessages: 1,
	}
	list, err := queue.GetMessages(&options)
	if err != nil {
		return err
	}
	
	fmt.Printf("Message ID %s\n", list[0].ID)

	err = list[0].Delete(nil)
	if err != nil {
		return err
	}
	
	return nil
}







