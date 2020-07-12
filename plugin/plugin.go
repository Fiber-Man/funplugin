package plugin

import (
	"github.com/Fiber-Man/funplugin"

	"github.com/graphql-go/graphql"
)

var pls funplugin.PluginManger

func New(pls1 funplugin.PluginManger) {
	pls = pls1
}

func Get(appname string) (funplugin.Plugin, error) {
	return pls.Get(appname)
}

func GetObject(objname string) (*graphql.Object, bool) {
	return pls.GetObject(objname)
}

func GetUnion(objname string) (*graphql.Union, bool) {
	return pls.GetUnion(objname)
}

func PluginList() map[string]funplugin.Plugin {
	return pls.PluginList()
}

func AutoField(names string) (*graphql.Field, error) {
	return pls.AutoField(names)
}
