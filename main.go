package funplugin

import (
	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
)

// PluginFile 插件文件代码，以后统一由程序生成
type PluginFile struct {
	// Name of the plugin e.g rabbitmq
	Name string
	// Type of the plugin e.g broker
	Type string
	// Path specifies the import path
	Path string
	// Version of the plugin
	Version string
	// NewFunc creates an instance of the plugin
	NewFunc interface{}
}

// Schema 插件 graphql 对象返回值，用于 service 获取视图和外部插件引用本插件视图、方法
type Schema struct {
	Object       map[string]*graphql.Object
	Union        map[string]*graphql.Union
	Query        graphql.Fields
	Mutation     graphql.Fields
	Subscription graphql.Fields
}

// Plugin 应用插件接口定义
type Plugin interface {
	// Global Flags
	// Flags() []cli.Flag
	// Sub-commands
	// Commands() []cli.Command
	// Handle is the middleware handler for HTTP requests. We pass in
	// the existing handler so it can be wrapped to create a call chain.
	Schema() Schema
	// Init called when command line args are parsed.
	// The initialised cli.Context is passed in.
	Init(*gorm.DB) error
	//Version of the plugin
	Version() string
	// Name of the plugin
	String() string
	//
	AddField(string, graphql.Output, func(graphql.ResolveParams) (interface{}, error)) error
	//
	Query(arg ...interface{}) (interface{}, error)
	//
	Func(interface{}) (interface{}, error)
}

// PluginManger service 插件管理器接口定义，各插件需要获取其他插件对象和数据
type PluginManger interface {
	Register(appname string) (Plugin, error)
	Get(appname string) (Plugin, error)
	GetUnion(objname string) (*graphql.Union, bool)
	GetObject(objname string) (*graphql.Object, bool)
	PluginList() map[string]Plugin
}
