package plugin

import (
	"github.com/Fiber-Man/funplugin"

	"github.com/graphql-go/graphql"
)

var pls funplugin.PluginManger

//set pls
func New(pls1 funplugin.PluginManger) {
	pls = pls1
}

func GetObject(objname string) (*graphql.Object, bool) {
	return pls.GetObject(objname)
}

func AutoField(names string) (*graphql.Field, error) {
	return pls.AutoField(names)
}

func Go(query string, VariableValues map[string]interface{}, dstStructPtr interface{}) (interface{}, error) {
	return pls.Go(query, VariableValues, dstStructPtr)
}
