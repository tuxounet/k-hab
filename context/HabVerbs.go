package context

type HabVerbs string

const (
	ProvisionVerb   HabVerbs = "Provision"
	UpVerb          HabVerbs = "Up"
	DeployVerb      HabVerbs = "Deploy"
	ShellVerb       HabVerbs = "Shell"
	UndeployVerb    HabVerbs = "Undeploy"
	DownVerb        HabVerbs = "Down"
	RmVerb          HabVerbs = "Rm"
	UnprovisionVerb HabVerbs = "Unprovision"
	NukeVerb        HabVerbs = "Nuke"
)
