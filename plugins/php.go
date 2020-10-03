package plugins

// php plugins interface
type PHPPlugin interface {
	Init() Plugin
	Verify(target string) bool
	GetResult() []Plugin
}

// register php plugin
func RegiserPHPPlugin(target string, plugin PHPPlugin) {
	PHPPlugins[target] = append(PHPPlugins[target], plugin)
}

var PHPPlugins map[string][]PHPPlugin

func runPHPPLugin() {}

func RunPHP() {}

