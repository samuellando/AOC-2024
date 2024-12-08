package common

import (
	"errors"
	"regexp"
	"strings"
)

type SyntaxTree interface {
	Tree() string
	Tag() string
	SetTag(string)
	Value() string
	Find(string, ...bool) []SyntaxTree
	Filter(string) []SyntaxTree
}

type syntaxTree struct {
	tag   string
	value []SyntaxTree
}

func (t *syntaxTree) Find(tag string, recurse ...bool) []SyntaxTree {
	r := false
	if len(recurse) > 0 {
		r = recurse[0]
	}
	matches := make([]SyntaxTree, 0, 5)
	for _, st := range t.value {
		if st != nil {
			if st.Tag() == tag {
				matches = append(matches, st)
				if r {
					matches = append(matches, st.Find(tag)...)
				}
			} else {
				matches = append(matches, st.Find(tag)...)
			}
		}
	}
	return matches
}

func (t *syntaxTree) Filter(tag string) []SyntaxTree {
	matches := make([]SyntaxTree, 0, 5)
	for _, st := range t.value {
		if st != nil {
			if st.Tag() == tag {
				matches = append(matches, st)
			}
		}
	}
	return matches
}

func (t *syntaxTree) Tree() string {
	s := "(" + t.tag + " "
	for _, t := range t.value {
		if t != nil {
			s += t.Tree()
		}
	}
	return s + ")"
}

func (t *syntaxTree) Value() string {
	s := ""
	for _, t := range t.value {
		if t != nil {
			s += t.Value()
		}
	}
	return s
}

func (t *syntaxTree) Tag() string {
	return t.tag
}

func (t *syntaxTree) SetTag(tag string) {
	t.tag = tag
}

type token struct {
	tag   string
	value string
}

func (t *token) Tree() string {
	return "(" + t.tag + " " + t.value + ")"
}

func (t *token) Value() string {
	return t.value
}

func (t *token) Tag() string {
	return t.tag
}

func (t *token) SetTag(tag string) {
	t.tag = tag
}
func (t *token) Find(tag string, recurse ...bool) []SyntaxTree {
	return []SyntaxTree{}
}

func (t *token) Filter(tag string) []SyntaxTree {
	return []SyntaxTree{}
}

type Grammar interface {
	Parse(string) SyntaxTree
}

type grammar struct {
	root  Rule
	rules map[string]Rule
}

func CreateGrammar(root Rule, rules ...Rule) Grammar {
	ruleSet := make(map[string]Rule)
	for _, r := range rules {
		ruleSet[r.Name()] = r
	}
	ruleSet[root.Name()] = root
	return &grammar{root, ruleSet}
}

func (g *grammar) Parse(input string) SyntaxTree {
	t, remainder, err := g.root.Parse(input, g.rules)
	if err != nil {
		panic(err)
	}
	if remainder != "" {
		panic(errors.New("Remaining: " + remainder + " " + t.Tree()))
	}
	return t
}

type Rule interface {
	Name() string
	Parse(string, map[string]Rule) (SyntaxTree, string, error)
}

type rule struct {
	name       string
	expression Expression
}

func CreateRule(name string, ex Expression) Rule {
	return &rule{name, ex}
}

func (r *rule) Name() string {
	return r.name
}

func (r *rule) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	t, remainder, err := r.expression.Parse(input, rules)
	if err != nil {
		return t, remainder, err
	}
	if t != nil {
		t.SetTag(r.name)
	}
	return t, remainder, err
}

type Expression interface {
	Parse(string, map[string]Rule) (SyntaxTree, string, error)
}

type seqExpression struct {
	expressions []Expression
}

func SeqExpression(exs ...Expression) Expression {
	return &seqExpression{exs}
}

func (e *seqExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	sts := make([]SyntaxTree, 0, 5)
	remainder := input
	var sub SyntaxTree
	var err error
	for _, ex := range e.expressions {
		sub, remainder, err = ex.Parse(remainder, rules)
		if err != nil {
			return nil, input, errors.New("Could not parse")
		} else {
			sts = append(sts, sub)
		}
	}
	t := &syntaxTree{"", sts}
	return t, remainder, nil
}

type zeroOrMoreExpression struct {
	expression Expression
}

func ZeroOrMoreExpression(ex Expression) Expression {
	return &zeroOrMoreExpression{ex}
}

func (e *zeroOrMoreExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	sts := make([]SyntaxTree, 0, 5)
	remainder := input
	var sub SyntaxTree
	var err error
	for {
		sub, remainder, err = e.expression.Parse(remainder, rules)
		if err != nil {
			break
		} else {
			sts = append(sts, sub)
		}
	}
	if len(sts) == 0 {
		return nil, remainder, nil
	}
	t := &syntaxTree{"", sts}
	return t, remainder, nil
}

type oneOrMoreExpression struct {
	expression Expression
}

func OneOrMoreExpression(ex Expression) Expression {
	return &oneOrMoreExpression{ex}
}

func (e *oneOrMoreExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	ze := ZeroOrMoreExpression(e.expression)
	t, remainder, err := ze.Parse(input, rules)
	if err != nil {
		return t, remainder, err
	}
	if t == nil {
		return t, remainder, errors.New("One or more expected")
	}
	return t, remainder, err
}

type optionalExpression struct {
	expression Expression
}

func OptionalExpression(ex Expression) Expression {
	return &optionalExpression{ex}
}

func (e *optionalExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	t, remainder, err := e.expression.Parse(input, rules)
	if err != nil {
		return nil, remainder, nil
	}
	return t, remainder, err
}

type orExpression struct {
	expressions []Expression
}

func OrExpression(exs ...Expression) Expression {
	return &orExpression{exs}
}

func (e *orExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	for _, ex := range e.expressions {
		t, remainder, err := ex.Parse(input, rules)
		if err != nil {
			continue
		}
		return t, remainder, err
	}
	return nil, input, errors.New("No option parsed succesfully")
}

type terminalExpression struct {
	pattern string
}

func TerminalExpression(v string) Expression {
	return &terminalExpression{v}
}

func (e *terminalExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	if strings.HasPrefix(input, e.pattern) {
		var remainder string
		if len(input) == len(e.pattern) {
			remainder = ""
		} else {
			remainder = input[len(e.pattern):]
		}
		return &token{"", e.pattern}, remainder, nil
	} else {
		return nil, input, errors.New("Could not match")
	}
}

type referenceExpression struct {
	name string
}

func ReferenceExpression(v string) Expression {
	return &referenceExpression{v}
}

type regexExpression struct {
	pattern string
}

func RegexExpression(v string) Expression {
	return &regexExpression{v}
}

func (e *regexExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	c := regexp.MustCompile("^" + e.pattern)
	match := c.FindString(input)
	if match != "" {
		var remainder string
		if len(input) == len(match) {
			remainder = ""
		} else {
			remainder = input[len(match):]
		}
		return &token{"", match}, remainder, nil
	} else {
		return nil, input, errors.New("Could not match")
	}
}

func (e *referenceExpression) Parse(input string, rules map[string]Rule) (SyntaxTree, string, error) {
	rule, ok := rules[e.name]
	if !ok {
		return nil, input, errors.New("Reference not found")
	}
	return rule.Parse(input, rules)
}
