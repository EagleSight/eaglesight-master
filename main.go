package main

import "github.com/eaglesight/eaglesight-master/master"
import manager "github.com/eaglesight/eaglesight-master/demo-manager"

func main() {

	manager := manager.CreateManager()

	master.Start(manager)

}
