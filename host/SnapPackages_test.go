package host

import (
	"testing"

	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/utils"
)

func TestTTSnapPackage(t *testing.T) {
	ctx := utils.NewTestContext()
	fakeConfigYaml := `{
		"snap": {
			"command":{
				"prefix": "sudo",
				"name": "snap"
			}
		}
	}`
	testSnap := "jq"

	config := utils.LoadJSONFromString[config.HabConfig](ctx, fakeConfigYaml)
	snap := NewSnapPackages(config)

	installted := snap.InstalledSnap(ctx, testSnap)
	if !installted {
		err := snap.RemoveSnap(ctx, testSnap)
		if err != nil {
			t.Fatalf("Error removing snap %s: %v", testSnap, err)
		}
	}

	err := snap.InstallSnap(ctx, testSnap, "")
	if err != nil {
		t.Fatalf("Error installing snap %s: %v", testSnap, err)
	}

	installted = snap.InstalledSnap(ctx, testSnap)
	if !installted {
		t.Fatalf("Error installing snap %s: %v", testSnap, err)
	}

	err = snap.TakeSnapSnapshots(ctx, testSnap)
	if err != nil {
		t.Fatalf("Error taking snapshot snap %s: %v", testSnap, err)
	}

	err = snap.RemoveSnap(ctx, testSnap)
	if err != nil {
		t.Fatalf("Error removing snap %s: %v", testSnap, err)
	}

	err = snap.RemoveSnapSnapshots(ctx, testSnap)
	if err != nil {
		t.Fatalf("Error removing snap snapshots %s: %v", testSnap, err)
	}

	out := snap.ListSnapshots(ctx, testSnap)
	if len(out) != 0 {
		t.Fatalf("Error listing snapshots %s: %v", testSnap, err)
	}

}
