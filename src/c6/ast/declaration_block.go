package ast

/*
DeclBlock.

.foo {
	property-name: property-value;
}
*/
type DeclBlock struct {
	// The symbol table for storing constant values
	// Only constants can be stored here...
	Stmts *StmtList

	// Nested rulesets
	SubRuleSets []*RuleSet

	Scope *Scope
}

func NewDeclBlock() *DeclBlock {
	return &DeclBlock{
		Stmts: &StmtList{},
	}
}

/**
Append a Declaration
*/
func (self *DeclBlock) Append(decl Stmt) {
	self.Stmts.Append(decl)
}

func (self *DeclBlock) AppendSubRuleSet(ruleset *RuleSet) {
	newRuleSets := append(self.SubRuleSets, ruleset)
	self.SubRuleSets = newRuleSets
}

func (self *DeclBlock) MergeStmts(stmts *StmtList) {
	for _, stm := range *stmts {
		self.Stmts.Append(stm)
	}
}

func (self DeclBlock) String() (out string) {
	out += "{\n"
	for _, decl := range *self.Stmts {
		out += decl.String() + "\n"
	}
	out += "}"
	return out
}
