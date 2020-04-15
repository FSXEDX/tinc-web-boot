package network

import (
	"bytes"
	"text/template"
)

var (
	tincUpTpl     = template.Must(template.New("").Parse(tincUpTxt))
	tincDownTpl   = template.Must(template.New("").Parse(tincDownText))
	subnetUpTpl   = template.Must(template.New("").Parse(subnetUpText))
	subnetDownTpl = template.Must(template.New("").Parse(subnetDownText))
)

func tincUp(selfNode *Node) string {
	return mustRender(tincUpTpl, selfNode)
}

func tincDown(selfNode *Node) string {
	return mustRender(tincDownTpl, selfNode)
}

func subnetUp(executable string, selfNode *Node) string {
	var params struct {
		Executable string
		Node       *Node
	}
	params.Executable = executable
	params.Node = selfNode
	return mustRender(subnetUpTpl, params)
}

func subnetDown(executable string, selfNode *Node) string {
	var params struct {
		Executable string
		Node       *Node
	}
	params.Executable = executable
	params.Node = selfNode
	return mustRender(subnetDownTpl, params)
}

func mustRender(tpl *template.Template, params interface{}) string {
	var out bytes.Buffer
	err := tpl.Execute(&out, params)
	if err != nil {
		panic(err)
	}
	return out.String()
}
