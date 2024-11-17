package context

type HabVerbs string

const (
	InstallVerb   HabVerbs = "Install"
	UninstallVerb HabVerbs = "Uninstall"

	ProvisionVerb   HabVerbs = "Provision"
	StartVerb       HabVerbs = "Start"
	DeployVerb      HabVerbs = "Deploy"
	UpVerb          HabVerbs = "Up"
	ShellVerb       HabVerbs = "Shell"
	RunVerb         HabVerbs = "Run"
	UndeployVerb    HabVerbs = "Undeploy"
	StopVerb        HabVerbs = "Stop"
	DownVerb        HabVerbs = "Down"
	RmVerb          HabVerbs = "Rm"
	UnprovisionVerb HabVerbs = "Unprovision"
	NukeVerb        HabVerbs = "Nuke"
)
