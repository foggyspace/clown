package plugins

// 插件信息
type Plugins struct {
	PluginName string `json:"plugin_name"` // 插件名
	Author     string `json:"author"`      // 插件作者
	Url        string `json:"url"`         // 作者github地址
	Request    string
	Response   string
	VulnsInfo  VulnsInfo
}

// 漏洞基本信息
type VulnsInfo struct {
	RefererUrl string // 漏洞参考链接
	CVE        string // 漏洞CVE编号
	Level      string // 漏洞那个风险等级
}

// 插件接口
type Pluginser interface {
	InitPlugin() Plugins // 初始化插件信息
	Audit(target string) // 检测漏洞,这里只需将漏洞结果进行输出即可,不返回任何内容
}
