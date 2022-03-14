package plugins

import (
	"fmt"
	"net/http"

	"github.com/seaung/clown/pkg/utils"
)

type JbossCve201712149 struct {
	pluginInfo Plugins
}

func (j *JbossCve201712149) InitPlugin() Plugins {
	j.pluginInfo = Plugins{
		PluginName: "Jboss反序列化漏洞",
		Author:     "seaung",
		Url:        "https://github.com/seaung",
		VulnsInfo: VulnsInfo{
			RefererUrl: "http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2017-12149",
			CVE:        "CVE-2017-12149",
			Level:      "Hight",
		},
	}

	return j.pluginInfo
}

func (j *JbossCve201712149) Audit(target string) {
	vulnUrl := target + "/invoker/readonly"

	info := j.InitPlugin()

	client := &http.Client{}

	request, err := http.NewRequest("POST", vulnUrl, nil)
	if err != nil {
		utils.NewLogger().ErrorLog(fmt.Sprintf("Request Cread Error : %v\n", err))
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:63.0) Gecko/20100101 Firefox/63.0")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Connection", "close")
	request.Header.Set("Cache-Control", "no-cache")

	response, err := client.Do(request)
	if err != nil {
		utils.NewLogger().ErrorLog(fmt.Sprintf("Send Request Error : %v\n", err))
	}

	defer response.Body.Close()

	if response.StatusCode == 500 {
		utils.NewLogger().Warnning("Found Vulnerable !")
		utils.NewLogger().Warnning(fmt.Sprintf("Vulnerable Name   : %s\n", info.PluginName))
		utils.NewLogger().Warnning(fmt.Sprintf("Vulnerable Level  : %s\n", info.VulnsInfo.Level))
	} else {
		utils.NewLogger().Info("Not Found Vulnerable")
	}
}
