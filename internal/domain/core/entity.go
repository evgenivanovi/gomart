package core

/* __________________________________________________ */

type Entity[ID any, DATA any] interface {
	Identity() Identity[ID]
	Data() DATA
	Metadata() Metadata
}

type IntEntity[DATA any] interface {
	Entity[int, DATA]
}

type StringEntity[DATA any] interface {
	Entity[string, DATA]
}

type VersionEntity[ID any, DATA any] interface {
	Identity() VersionIdentity[ID]
	Data() DATA
}

type IntVersionEntity[DATA any] interface {
	VersionEntity[int, DATA]
}

type StringVersionEntity[DATA any] interface {
	VersionEntity[string, DATA]
}

/* __________________________________________________ */

type PartialEntity[ID any, DATA any] interface {
	Identity() Identity[ID]
	PartData() DATA
}

type IntPartialEntity[DATA any] interface {
	PartialEntity[int, DATA]
}

type StringPartialEntity[DATA any] interface {
	PartialEntity[string, DATA]
}

type PartialVersionEntity[ID any, DATA any] interface {
	Identity() VersionIdentity[ID]
	PartData() DATA
}

type IntVersionPartialEntity[DATA any] interface {
	PartialVersionEntity[int, DATA]
}

type StringVersionPartialEntity[DATA any] interface {
	PartialVersionEntity[string, DATA]
}

/* __________________________________________________ */
