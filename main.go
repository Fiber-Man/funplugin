package funplugin

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

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

//o2g
type ObjectSchema struct {
	GraphQLType *graphql.Object
	Object      interface{}
	Query       graphql.Fields
	Mutation    graphql.Fields
}

//o2g
type EnumValue struct {
	Value       interface{}
	Description string
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
	// Install error
	// Dependencies map[string]string // map[name]version
}

// PluginManger Interface
type PluginManger interface {
	GetObject(objname string) (*graphql.Object, bool)
	AutoField(names string) (*graphql.Field, error)
	Go(query string, params map[string]interface{}, dstStructPtr interface{}) (interface{}, error)
	NewSchemaBuilder(object interface{}) (*ObjectSchema, error)
}

//ID2id is string to uint
func ID2id(ID interface{}) (uint, error) {
	if ID == nil || reflect.TypeOf(ID).String() != "string" {
		return 0, errors.New("ID Type error")
	}

	ID2 := ID.(string)
	if p1 := strings.Index(ID2, "-"); p1 > -1 {
		ID2 = ID2[p1+1:]
	}
	id, err := strconv.ParseUint(ID2, 10, 64)
	return uint(id), err
}
