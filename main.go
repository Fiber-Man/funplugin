package funplugin

import (
	"github.com/gin-gonic/gin"
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
	Get          map[string]func(c *gin.Context)
	Post         map[string]func(c *gin.Context)
}

// Plugin App Interface
type Plugin interface {
	Schema() Schema
	Init(*gorm.DB) error
	Version() string
	String() string
	Query(arg ...interface{}) (interface{}, error)
	Func(interface{}) (interface{}, error)
	Setup() error
	// Upgrade() error
}

// PluginManger Interface
type PluginManger interface {
	GetObject(objname string) (*graphql.Object, bool)
	AutoField(names string) (*graphql.Field, error)
	Go(query string, params map[string]interface{}) (interface{}, error)
}
