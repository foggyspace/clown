package commons

import (
	"encoding/json"
	"fmt"
	"os"
	"net"
	"github.com/olekukonko/tablewriter"
	"github.com/c-bata/go-prompt"
)

func parseLocalIP() map[string]string {
	res := make(map[string]string)

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			res[i.Name] = addr.String()
			break
		}
	}
	return res
}

func parseCIDR(s string) (string, error) {
	_, ipv4, err := net.ParseCIDR(s)
	if err != nil {
		return "", err
	}
	return ipv4.String(), nil
}

func StructToMap(obj interface{}) map[string]interface{} {
	jsonObject, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(jsonObject, &result)
	return result
}

func GetTargetSuggestions() (s []prompt.Suggest) {
	s = make([]prompt.Suggest, 1, 5)
	s[0] = prompt.Suggest{
		Text: Const_example_target_cidr,
		Description: Const_example_target_desc,
	}

	localInterfaces := parseLocalIP()

	for eth, ip := range localInterfaces {
		parseIP, err := parseCIDR(ip)
		if err != nil {
			s = append(s, prompt.Suggest{
				Text: parseIP,
				Description: fmt.Sprintf("subnet from interface : %s", eth),
			})
		}
	}

	return s
}

func ShowHelp() {
	data := [][]string{}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Area", "Command", "Syntax"})
	table.SetAlignment(3)
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}
