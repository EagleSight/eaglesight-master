package manager

import (
	"testing"
	"time"
)

func TestCreateManager(t *testing.T) {

	manager := CreateManager()

	if manager.list == nil {
		t.Error(manager)
	}

}

func TestStartSlave(t *testing.T) {
	closer := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		close(closer)
	}()

	startSlave(closer)
}

func TestSpawnAndKill(t *testing.T) {

	manager := CreateManager()

	id, err := manager.Spawn()

	if err != nil {
		t.Error(err)
	}

	// Check if the id is in the list
	if _, ok := manager.list[id]; !ok {
		t.Errorf("%s not in manager.list", id)
	}

	err = manager.Kill(id)

	if err != nil {
		t.Error(err)
	}

	if _, ok := manager.list[id]; ok {
		t.Errorf("%s still in manager.list", id)
	}

}

func TestKill(t *testing.T) {
	manager := CreateManager()

	err := manager.Kill("doen't exists")

	if err == nil {
		t.Fail()
	}

}
