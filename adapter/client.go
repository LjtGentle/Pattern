package main

import "fmt"

type client struct {

}

func (c *client) insertLightningConnectorComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer")
	com.insertIntoLightningPort()
}
