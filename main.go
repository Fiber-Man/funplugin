package funplugin

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// PluginFile PluginInterface
type PluginFile struct {
	Name    string
	Type    string
	Path    string
	Version string
	NewFunc interface{}
}

// Schema struct
type Schema struct {
	Object       map[string]*graphql.Object
	Union        map[string]*graphql.Union
	Query        graphql.Fields
	Mutation     graphql.Fields
	Subscription graphql.Fields
}

// Plugin App Interface
type Plugin interface {
	Schema() Schema
	Init(*gorm.DB) error
	Version() string
	String() string
	AddField(string, graphql.Output, func(graphql.ResolveParams) (interface{}, error)) error
	Query(arg ...interface{}) (interface{}, error)
	Func(interface{}) (interface{}, error)
	Setup() error
}

// PluginManger Interface
type PluginManger interface {
	Register(appname string) (Plugin, error)
	Get(appname string) (Plugin, error)
	GetUnion(objname string) (*graphql.Union, bool)
	GetObject(objname string) (*graphql.Object, bool)
	PluginList() map[string]Plugin
}
