package context

type HabVerbs string

const (
	ProvisionVerb   HabVerbs = "Provision"
	UpVerb          HabVerbs = "Up"
	ShellVerb       HabVerbs = "Shell"
	DownVerb        HabVerbs = "Down"
	RmVerb          HabVerbs = "Rm"
	UnprovisionVerb HabVerbs = "Unprovision"
	NukeVerb        HabVerbs = "Nuke"
)
