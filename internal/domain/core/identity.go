package core

import (
	"github.com/evgenivanovi/gomart/pkg/std/cmp"
)

/* __________________________________________________ */

type Identity[ID any] interface {
	ID() ID
	cmp.Equaler
}

type IntID interface {
	Identity[int]
	cmp.Comparer
}

type StringID interface {
	Identity[string]
}

/* __________________________________________________ */

type VersionIdentity[ID any] interface {
	Identity[ID]
	Version() int
}

type IntVersionID interface {
	VersionIdentity[int]
	cmp.Comparer
}

type StringVersionID interface {
	cmp.Equaler
	VersionIdentity[string]
}

/* __________________________________________________ */
