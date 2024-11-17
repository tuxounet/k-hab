package context

type HabVerbs string

const (
	ProvisionVerb   HabVerbs = "Provision"
	StartVerb       HabVerbs = "Start"
	DeployVerb      HabVerbs = "Deploy"
	UpVerb          HabVerbs = "Up"
	ShellVerb       HabVerbs = "Shell"
	UndeployVerb    HabVerbs = "Undeploy"
	StopVerb        HabVerbs = "Stop"
	DownVerb        HabVerbs = "Down"
	RmVerb          HabVerbs = "Rm"
	UnprovisionVerb HabVerbs = "Unprovision"
	NukeVerb        HabVerbs = "Nuke"
)
