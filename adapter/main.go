package main

func main() {
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorComputer(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{windowMachine: windowsMachine}
	client.insertLightningConnectorComputer(windowsMachineAdapter)
}