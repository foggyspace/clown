package plugins

// java plugins interface
type JavaPlugin interface {
	Init() Plugin
	Verify(target string) bool
	GetResult() []Plugin
}

// register java plugin
func Register(target string, plugin JavaPlugin) {
	JavaPlugins[target] = append(JavaPlugins[target], plugin)
}

// java plugin map
var JavaPlugins map[string][]JavaPlugin

func runJavaPlugin() {}

func RunJava() {}

