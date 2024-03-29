package app

import (
	"bitscale/buildnet/lib/common"
)

type Blockchain struct {
	common.HasNameAndId
	Modules          []Module
	Network          Network
	HasProjectModule bool
	TokenFactory     string
	Frontends        []string
}

type Network struct {
}

type Module struct {
	common.HasNameAndId
	Dependencies []string
	Entities     []Artifact
	Messages     []Artifact
	Queries      []Artifact
	IsIBCEnabled bool
}

type Artifact struct {
	common.HasNameAndId
	Type   string
	Fields []Field
}

type Field struct {
	common.HasNameAndId
	Type string
}
