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

// MultipleChoiceScoringMethodType defines the multiple_choice_scoring_method for the "multiple_choice_scoring_method" enum field.
type MultipleChoiceScoringMethodType string

// MultipleChoiceScoringMethodType values.
const (
	AllOrNothing  MultipleChoiceScoringMethodType = "ALL_OR_NOTHING"
	HalfForMissed MultipleChoiceScoringMethodType = "HALF_FOR_MISSED"
)

func (_type MultipleChoiceScoringMethodType) String() string {
	return string(_type)
}

func (MultipleChoiceScoringMethodType) Values() []string {
	return []string{
		string(AllOrNothing),
		string(HalfForMissed),
	}
}

// JudgementType defines the judgement_type for the "judgement_type" enum field.
type JudgementType string

// ObjectiveProblemJudgementType values.
const (
	BatchJudgeAfterEnd JudgementType = "BATCH_JUDGE_AFTER_END"
	JudgeInTime        JudgementType = "JUDGE_IN_TIME"
)

func (_type JudgementType) String() string {
	return string(_type)
}

func (JudgementType) Values() []string {
	return []string{
		string(BatchJudgeAfterEnd),
		string(JudgeInTime),
	}
}

type TypeOrderItem struct {
	Name string `json:"name"`
	Type Type   `json:"type"`
	Min  int    `json:"min"` // 最少题目数
	Max  int    `json:"max"` // 最多题目数
}
