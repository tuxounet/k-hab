package hab

import "testing"

func TestTTHabLifecycle(t *testing.T) {

	hab := NewHab(true)

	err := hab.Provision()
	if err != nil {
		t.Fatalf("Error provisioning: %v", err)
	}

	err = hab.Start()
	if err != nil {
		t.Fatalf("Error starting: %v", err)
	}

	err = hab.Stop()
	if err != nil {
		t.Fatalf("Error stopping: %v", err)
	}

	err = hab.Unprovision()
	if err != nil {
		t.Fatalf("Error unprovisioning: %v", err)
	}
	err = hab.Nuke()
	if err != nil {
		t.Fatalf("Error nuking: %v", err)
	}

}
