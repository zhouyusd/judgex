package problemset

import (
	"fmt"

	"entgo.io/ent/dialect/entsql"
)

// RuleType defines the rule_type for the "rule_type" enum field.
type RuleType string

// RuleType values.
const (
	RuleTypeCourse RuleType = "COURSE" // 用作按课程章节关联题目
	RuleTypeCustom RuleType = "CUSTOM"
	RuleTypeExam   RuleType = "EXAM" // 用作考试
	RuleTypeICPC   RuleType = "ICPC"
	RuleTypeIOI    RuleType = "IOI"
	RuleTypeOI     RuleType = "OI"
)

func (_type RuleType) String() string {
	return string(_type)
}

func (RuleType) Values() []string {
	return []string{
		string(RuleTypeCourse),
		string(RuleTypeCustom),
		string(RuleTypeExam),
		string(RuleTypeICPC),
		string(RuleTypeIOI),
		string(RuleTypeOI),
	}
}

func (_type RuleType) Check() *entsql.Annotation {
	return entsql.Check(fmt.Sprintf("`rule_type`='%v'", _type))
}

// RuleTypeValidator is a validator for the "rule_type" field enum values. It is called by the builders before save.
func RuleTypeValidator(_type RuleType) error {
	switch _type {
	case RuleTypeCourse, RuleTypeCustom, RuleTypeExam, RuleTypeICPC, RuleTypeIOI, RuleTypeOI:
		return nil
	default:
		return fmt.Errorf("problemset: invalid enum value for rule_type field: %q", _type)
	}
}
