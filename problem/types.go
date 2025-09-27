package problem

import (
	"fmt"

	"entgo.io/ent/dialect/entsql"
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeFillInBlank    Type = "FILL_IN_BLANK"
	TypeMultipleChoice Type = "MULTIPLE_CHOICE"
	TypeProgramming    Type = "PROGRAMMING"
	TypeSingleChoice   Type = "SINGLE_CHOICE"
	TypeSubjective     Type = "SUBJECTIVE"
	TypeTrueOrFalse    Type = "TRUE_OR_FALSE"
)

func (_type Type) String() string {
	return string(_type)
}

func (Type) Values() []string {
	return []string{
		string(TypeFillInBlank),
		string(TypeMultipleChoice),
		string(TypeProgramming),
		string(TypeSingleChoice),
		string(TypeSubjective),
		string(TypeTrueOrFalse),
	}
}

func (_type Type) Check() *entsql.Annotation {
	return entsql.Check(fmt.Sprintf("`type`='%v'", _type))
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeFillInBlank, TypeMultipleChoice, TypeProgramming, TypeSingleChoice, TypeSubjective, TypeTrueOrFalse:
		return nil
	default:
		return fmt.Errorf("problem: invalid enum value for type field: %q", _type)
	}
}
