package query

import (
	"fmt"
	"time"
)

// See https://developers.google.com/drive/api/v3/search-parameters

const (
	operatorDropSecond    = 0
	operatorAnd           = iota
	operatorOr            = iota
	operatorContains      = iota
	operatorEqual         = iota
	operatorNotEqual      = iota
	operatorAfter         = iota
	operatorAfterOrEqual  = iota
	operatorBefore        = iota
	operatorBeforeOrEqual = iota
)

var operators = initOperators()

func initOperators() map[int]string {
	return map[int]string{
		operatorAfter:         ">",
		operatorBefore:        "<",
		operatorAfterOrEqual:  ">=",
		operatorBeforeOrEqual: "<=",
		operatorAnd:           "and",
		operatorContains:      "contains",
		operatorEqual:         "=",
		operatorNotEqual:      "!=",
		operatorOr:            "or",
	}
}

type query struct {
	operator int
	left     Statement
	right    Statement
}

func (e *query) toString() string {
	switch e.operator {
	case operatorOr, operatorAnd:
		return fmt.Sprintf("(%s) %s (%s)", e.left.toString(), operators[e.operator], e.right.toString())
	case operatorDropSecond:
		return e.left.toString()
	default:
		panic("unsupported operator")
	}
}

func (e *query) And(right Statement) QueryStatement {
	return &query{operatorAnd, e, right}

}
func (e *query) Or(right Statement) QueryStatement {
	return &query{operatorOr, e, right}
}

type propertiesHasStatement struct {
	key   string
	value string
}

func (p *propertiesHasStatement) toString() string {
	return fmt.Sprintf("properties has { key=\"%s\" and value=\"%s\" }", p.key, p.value)
}

type textAttributeStatement struct {
	operator  int
	attribute string
	value     string
}

func (s *textAttributeStatement) toString() string {
	switch s.operator {
	case operatorContains, operatorEqual, operatorNotEqual:
		return fmt.Sprintf("%s %s \"%s\"", s.attribute, operators[s.operator], s.value)
	default:
		panic("textAttributeStatement: operator not supported")
	}
}

type fullTextAttribute struct {
}

func (a *fullTextAttribute) Contains(v string) Statement {
	return &textAttributeStatement{operatorContains, "fullText", v}
}

type textAttribute struct {
	name string
}

func (a *textAttribute) Contains(v string) Statement {
	return &textAttributeStatement{operatorContains, a.name, v}
}

func (a *textAttribute) Equal(v string) Statement {
	return &textAttributeStatement{operatorEqual, a.name, v}
}
func (a *textAttribute) NotEqual(v string) Statement {
	return &textAttributeStatement{operatorNotEqual, a.name, v}
}

type timeAttributeStatement struct {
	operator  int
	attribute string
	value     time.Time
}

func (s *timeAttributeStatement) toString() string {
	switch s.operator {
	case operatorAfter, operatorAfterOrEqual, operatorBefore, operatorBeforeOrEqual, operatorEqual, operatorNotEqual:
		return fmt.Sprintf("%s %s \"%s\"", s.attribute, operators[s.operator], s.value.Format(time.RFC3339))
	default:
		panic("textAttributeStatement: operator not supported")
	}

}

type timeAttribute struct {
	name string
}

func (a *timeAttribute) After(v time.Time) Statement {
	return &timeAttributeStatement{operatorAfter, a.name, v}
}
func (a *timeAttribute) AfterOrEqual(v time.Time) Statement {
	return &timeAttributeStatement{operatorAfterOrEqual, a.name, v}
}

func (a *timeAttribute) Before(v time.Time) Statement {
	return &timeAttributeStatement{operatorBefore, a.name, v}
}
func (a *timeAttribute) BeforeOrEqual(v time.Time) Statement {
	return &timeAttributeStatement{operatorBeforeOrEqual, a.name, v}
}

func (a *timeAttribute) Equal(v time.Time) Statement {
	return &timeAttributeStatement{operatorEqual, a.name, v}
}
func (a *timeAttribute) NotEqual(v time.Time) Statement {
	return &timeAttributeStatement{operatorNotEqual, a.name, v}
}

type rawStatement struct {
	raw string
}

func (s *rawStatement) toString() string {
	return s.raw
}

type collectionAttributeStatement struct {
	collection string
	value      string
}

func (s *collectionAttributeStatement) toString() string {
	return fmt.Sprintf("\"%s\" in %s", s.value, s.collection)

}

type collectionAttributeAttribute struct {
	name string
}

func (a *collectionAttributeAttribute) Includes(v string) Statement {
	return &collectionAttributeStatement{a.name, v}
}
