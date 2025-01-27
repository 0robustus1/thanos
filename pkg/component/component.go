package component

import (
	"strings"

	"github.com/improbable-eng/thanos/pkg/store/storepb"
)

// Component is a generic component interface.
type Component interface {
	String() string
}

// StoreAPI is a component that implements Thanos' gRPC StoreAPI.
type StoreAPI interface {
	implementsStoreAPI()
	String() string
	ToProto() storepb.StoreType
}

// Source is a Thanos component that produce blocks of metrics.
type Source interface {
	producesBlocks()
	String() string
}

// SourceStoreAPI is a component that implements Thanos' gRPC StoreAPI
// and produce blocks of metrics.
type SourceStoreAPI interface {
	implementsStoreAPI()
	producesBlocks()
	String() string
	ToProto() storepb.StoreType
}

type component struct {
	name string
}

func (c component) String() string { return c.name }

type storeAPI struct {
	component
}

func (storeAPI) implementsStoreAPI() {}

func (s sourceStoreAPI) ToProto() storepb.StoreType {
	return storepb.StoreType(storepb.StoreType_value[strings.ToUpper(s.String())])
}

func (s storeAPI) ToProto() storepb.StoreType {
	return storepb.StoreType(storepb.StoreType_value[strings.ToUpper(s.String())])
}

type source struct {
	component
}

func (source) producesBlocks() {}

type sourceStoreAPI struct {
	component
	source
	storeAPI
}

// FromProto converts from a gRPC StoreType to StoreAPI.
func FromProto(storeType storepb.StoreType) StoreAPI {
	switch storeType {
	case storepb.StoreType_QUERY:
		return Query
	case storepb.StoreType_RULE:
		return Rule
	case storepb.StoreType_SIDECAR:
		return Sidecar
	case storepb.StoreType_STORE:
		return Store
	case storepb.StoreType_RECEIVE:
		return Receive
	default:
		return nil
	}
}

var (
	Bucket     = source{component: component{name: "bucket"}}
	Compact    = source{component: component{name: "compact"}}
	Downsample = source{component: component{name: "downsample"}}
	Query      = sourceStoreAPI{component: component{name: "query"}}
	Rule       = sourceStoreAPI{component: component{name: "rule"}}
	Sidecar    = sourceStoreAPI{component: component{name: "sidecar"}}
	Store      = sourceStoreAPI{component: component{name: "store"}}
	Receive    = sourceStoreAPI{component: component{name: "receive"}}
)
