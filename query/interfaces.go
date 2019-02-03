package query

import "time"

type Statement interface {
	toString() string
}

type QueryStatement interface {
	Statement
	And(Statement) QueryStatement
	Or(Statement) QueryStatement
}

type FreeTextAttribute interface {
	Contains(string) Statement
}

type TextAttribute interface {
	FreeTextAttribute
	Equal(string) Statement
	NotEqual(string) Statement
}

type TimeAttribute interface {
	Equal(time.Time) Statement
	After(time.Time) Statement
	AfterOrEqual(time.Time) Statement
	Before(time.Time) Statement
	BeforeOrEqual(time.Time) Statement
}

type CollectionAttribute interface {
	Includes(string) Statement
}
