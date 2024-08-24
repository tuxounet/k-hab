package dependencies

// import (
// 	"testing"

// 	"github.com/tuxounet/k-hab/bases"
// 	"github.com/tuxounet/k-hab/tests"

// 	"github.com/tuxounet/k-hab/utils"
// )

// func TestTTSnapPackage(t *testing.T) {

// 	fakeContext := tests.NewTestContext()
// 	fakeConfigYaml := `{
// 		"snap": {
// 			"command":{
// 				"prefix": "sudo",
// 				"name": "snap"
// 			}
// 		}
// 	}`
// 	testSnap := "jq"

// 	config, err := utils.LoadJSONFromString[bases.HabConfig](fakeConfigYaml)
// 	if err != nil {
// 		t.Fatalf("Error loading config: %v", err)
// 	}
// 	fakeContext.SetHabConfig(config)

// 	deps := NewDependenciesController(fakeContext)

// 	installted, err := deps.InstalledSnap(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error checking snap %s: %v", testSnap, err)
// 	}
// 	if !installted {
// 		err := deps.RemoveSnap(testSnap)
// 		if err != nil {
// 			t.Fatalf("Error removing snap %s: %v", testSnap, err)
// 		}
// 	}

// 	err = deps.InstallSnap(testSnap, "")
// 	if err != nil {
// 		t.Fatalf("Error installing snap %s: %v", testSnap, err)
// 	}

// 	installted, err = deps.InstalledSnap(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error checking snap %s: %v", testSnap, err)
// 	}

// 	if !installted {
// 		t.Fatalf("Error installing snap %s: %v", testSnap, err)
// 	}

// 	err = deps.TakeSnapSnapshots(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error taking snapshot snap %s: %v", testSnap, err)
// 	}

// 	err = deps.RemoveSnap(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error removing snap %s: %v", testSnap, err)
// 	}

// 	err = deps.RemoveSnapSnapshots(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error removing snap snapshots %s: %v", testSnap, err)
// 	}

// 	out, err := deps.ListSnapshots(testSnap)
// 	if err != nil {
// 		t.Fatalf("Error listing snapshots %s: %v", testSnap, err)
// 	}

// 	if len(out) != 0 {
// 		t.Fatalf("Error listing snapshots %s: %v", testSnap, err)
// 	}

// }
