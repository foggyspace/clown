package pocs

import (
	"fmt"
	"strings"

	"Cyker/commons"
)

// auditor interface
type Auditor interface {
	Init() Pocs
	Audit(targetURL string) bool
	GetResult()[]Pocs
}

type ScanTask struct {
	Type      string `json:"type"`
	Target    string `json:"target"`
	Netloc    string `json:"netloc"`
}

// vulnerable description
type VulnDesc struct {
	URL string `json:"url"`
	CVE string `json:"cve"`
}

// pocs description
type Pocs struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	VulnDesc    VulnDesc `json:"vulndesc"`
}

// audit list
var AuditList map[string][]Auditor

// register pocs
func RegisterPoc(target string, audit Auditor) {
	AuditList[target] = append(AuditList[target], audit)
}

// init audit plugins list
func init() {
	AuditList = make(map[string][]Auditor)
	commons.Banner()
}

// function try
func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}

// run pocs plugins
func runPocs(pocsTask ScanTask, audits Auditor) (result []map[string]interface{}) {
	var hasVul bool

	try(func() {
		hasVul = audits.Audit(targetURL)
	},
	func(e interface{}) {
		fmt.Println("error")
	})

	if hasVul == false {
		return
	}

	for _, res := range audits.GetResult() {
		result = append(result, commons.StructToMap(res))
	}
	return result
}

func Start(scanTask ScanTask) (result []map[string]interface{}) {
	for n, auditPlugins := range AuditList {
		if strings.Contains(strings.ToLower(scanTask.Target), "cve-") {
			for _, plugin := range auditPlugins {
				auditPluginInfo := plugin.Init()
				if strings.ToLower(auditPluginInfo.VulnDesc.CVE) != strings.ToLower(scanTask.Target) {
					continue
				}
				resultList := runPocs(scanTask, plugin)
				result = append(result, resultList...)
				break
			}
		} else if strings.Contains(strings.ToLower(scanTask.Target), strings.ToLower(n)) || scanTask.Target == "all" {
			for _, plugin := range auditPlugins {
				auditPluginInfo := plugin.Init()
				fmt.Println(auditPluginInfo.Name)
				resultList := runPocs(scanTask, plugin)
				result = append(result, resultList...)
			}
		}
	}
	return result
}

func GetPluginsInfos() (plugins []map[string]interface{}) {
	for name, pluginList := range AuditList {
		for _, plugin := range pluginList {
			info := plugin.Init()
			pluginMap := commons.StructToMap(info)
			pluginMap["target"] = name
			plugins = append(plugins, pluginMap)
		}
	}
	return
}
