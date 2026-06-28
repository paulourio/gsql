package sql

import "github.com/goccy/go-googlesql"

// TVFArgument wraps *googlesql.ASTTVFArgument.
// Script is already declared in nodes_query.go.
type TVFArgument struct {
	baseNode[*googlesql.ASTTVFArgument]
}

func newASTTVFArgument(r *googlesql.ASTTVFArgument) *TVFArgument {
	if r == nil {
		return nil
	}
	return &TVFArgument{baseNode[*googlesql.ASTTVFArgument]{raw: r}}
}

func (n *TVFArgument) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *TVFArgument) TableClause() Node {
	return Wrap(must(n.raw.TableClause()))
}

func (n *TVFArgument) ModelClause() Node {
	return Wrap(must(n.raw.ModelClause()))
}

func (n *TVFArgument) ConnectionClause() Node {
	return Wrap(must(n.raw.ConnectionClause()))
}

func (n *TVFArgument) Descriptor() Node {
	return Wrap(must(n.raw.Descriptor()))
}

// ExceptionHandler wraps *googlesql.ASTExceptionHandler.
type ExceptionHandler struct {
	baseNode[*googlesql.ASTExceptionHandler]
}

func newASTExceptionHandler(r *googlesql.ASTExceptionHandler) *ExceptionHandler {
	if r == nil {
		return nil
	}
	return &ExceptionHandler{baseNode[*googlesql.ASTExceptionHandler]{raw: r}}
}

func (n *ExceptionHandler) StatementList() *StatementList {
	return newASTStatementList(must(n.raw.StatementList()))
}

// ExceptionHandlerList wraps *googlesql.ASTExceptionHandlerList.
type ExceptionHandlerList struct {
	baseNode[*googlesql.ASTExceptionHandlerList]
}

func newASTExceptionHandlerList(r *googlesql.ASTExceptionHandlerList) *ExceptionHandlerList {
	if r == nil {
		return nil
	}
	return &ExceptionHandlerList{baseNode[*googlesql.ASTExceptionHandlerList]{raw: r}}
}

// Handlers returns all *ExceptionHandler children.
func (n *ExceptionHandlerList) Handlers() []*ExceptionHandler {
	var result []*ExceptionHandler
	for _, c := range n.Children() {
		if h, ok := c.(*ExceptionHandler); ok {
			result = append(result, h)
		}
	}
	return result
}

// BeginEndBlock wraps *googlesql.ASTBeginEndBlock.
type BeginEndBlock struct {
	baseNode[*googlesql.ASTBeginEndBlock]
}

func newASTBeginEndBlock(r *googlesql.ASTBeginEndBlock) *BeginEndBlock {
	if r == nil {
		return nil
	}
	return &BeginEndBlock{baseNode[*googlesql.ASTBeginEndBlock]{raw: r}}
}
func (n *BeginEndBlock) isStatement() {}

func (n *BeginEndBlock) StatementListNode() *StatementList {
	return newASTStatementList(must(n.raw.StatementListNode()))
}
func (n *BeginEndBlock) HasExceptionHandler() bool { return must(n.raw.HasExceptionHandler()) }
func (n *BeginEndBlock) HandlerList() *ExceptionHandlerList {
	return newASTExceptionHandlerList(must(n.raw.HandlerList()))
}

// BeginStatement wraps *googlesql.ASTBeginStatement.
type BeginStatement struct {
	baseNode[*googlesql.ASTBeginStatement]
}

func newASTBeginStatement(r *googlesql.ASTBeginStatement) *BeginStatement {
	if r == nil {
		return nil
	}
	return &BeginStatement{baseNode[*googlesql.ASTBeginStatement]{raw: r}}
}
func (n *BeginStatement) isStatement() {}

// RollbackStatement wraps *googlesql.ASTRollbackStatement.
type RollbackStatement struct {
	baseNode[*googlesql.ASTRollbackStatement]
}

func newASTRollbackStatement(r *googlesql.ASTRollbackStatement) *RollbackStatement {
	if r == nil {
		return nil
	}
	return &RollbackStatement{baseNode[*googlesql.ASTRollbackStatement]{raw: r}}
}
func (n *RollbackStatement) isStatement() {}

// CallStatement wraps *googlesql.ASTCallStatement.
type CallStatement struct {
	baseNode[*googlesql.ASTCallStatement]
}

func newASTCallStatement(r *googlesql.ASTCallStatement) *CallStatement {
	if r == nil {
		return nil
	}
	return &CallStatement{baseNode[*googlesql.ASTCallStatement]{raw: r}}
}
func (n *CallStatement) isStatement() {}

func (n *CallStatement) ProcedureName() *PathExpression {
	return newASTPathExpression(must(n.raw.ProcedureName()))
}

// TVFArguments returns []*TVFArgument children via Children().
func (n *CallStatement) TVFArguments() []*TVFArgument {
	var result []*TVFArgument
	for _, c := range n.Children() {
		if a, ok := c.(*TVFArgument); ok {
			result = append(result, a)
		}
	}
	return result
}

// CommitStatement wraps *googlesql.ASTCommitStatement.
type CommitStatement struct {
	baseNode[*googlesql.ASTCommitStatement]
}

func newASTCommitStatement(r *googlesql.ASTCommitStatement) *CommitStatement {
	if r == nil {
		return nil
	}
	return &CommitStatement{baseNode[*googlesql.ASTCommitStatement]{raw: r}}
}
func (n *CommitStatement) isStatement() {}

// ExecuteIntoClause wraps *googlesql.ASTExecuteIntoClause.
type ExecuteIntoClause struct {
	baseNode[*googlesql.ASTExecuteIntoClause]
}

func newASTExecuteIntoClause(r *googlesql.ASTExecuteIntoClause) *ExecuteIntoClause {
	if r == nil {
		return nil
	}
	return &ExecuteIntoClause{baseNode[*googlesql.ASTExecuteIntoClause]{raw: r}}
}

func (n *ExecuteIntoClause) Identifiers() *IdentifierList {
	return newASTIdentifierList(must(n.raw.Identifiers()))
}

// ExecuteUsingArgument wraps *googlesql.ASTExecuteUsingArgument.
type ExecuteUsingArgument struct {
	baseNode[*googlesql.ASTExecuteUsingArgument]
}

func newASTExecuteUsingArgument(r *googlesql.ASTExecuteUsingArgument) *ExecuteUsingArgument {
	if r == nil {
		return nil
	}
	return &ExecuteUsingArgument{baseNode[*googlesql.ASTExecuteUsingArgument]{raw: r}}
}

func (n *ExecuteUsingArgument) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}
func (n *ExecuteUsingArgument) Alias() *Alias { return newASTAlias(must(n.raw.Alias())) }

// ExecuteUsingClause wraps *googlesql.ASTExecuteUsingClause.
type ExecuteUsingClause struct {
	baseNode[*googlesql.ASTExecuteUsingClause]
}

func newASTExecuteUsingClause(r *googlesql.ASTExecuteUsingClause) *ExecuteUsingClause {
	if r == nil {
		return nil
	}
	return &ExecuteUsingClause{baseNode[*googlesql.ASTExecuteUsingClause]{raw: r}}
}

// Arguments returns all *ExecuteUsingArgument children.
func (n *ExecuteUsingClause) Arguments() []*ExecuteUsingArgument {
	var result []*ExecuteUsingArgument
	for _, c := range n.Children() {
		if a, ok := c.(*ExecuteUsingArgument); ok {
			result = append(result, a)
		}
	}
	return result
}

// ExecuteImmediateStatement wraps *googlesql.ASTExecuteImmediateStatement.
type ExecuteImmediateStatement struct {
	baseNode[*googlesql.ASTExecuteImmediateStatement]
}

func newASTExecuteImmediateStatement(r *googlesql.ASTExecuteImmediateStatement) *ExecuteImmediateStatement {
	if r == nil {
		return nil
	}
	return &ExecuteImmediateStatement{baseNode[*googlesql.ASTExecuteImmediateStatement]{raw: r}}
}
func (n *ExecuteImmediateStatement) isStatement() {}

func (n *ExecuteImmediateStatement) SQL() ExpressionNode {
	return wrapExpr(must(n.raw.Sql()))
}

func (n *ExecuteImmediateStatement) IntoClause() *ExecuteIntoClause {
	return newASTExecuteIntoClause(must(n.raw.IntoClause()))
}

func (n *ExecuteImmediateStatement) UsingClause() *ExecuteUsingClause {
	return newASTExecuteUsingClause(must(n.raw.UsingClause()))
}

// ElseifClause wraps *googlesql.ASTElseifClause.
type ElseifClause struct {
	baseNode[*googlesql.ASTElseifClause]
}

func newASTElseifClause(r *googlesql.ASTElseifClause) *ElseifClause {
	if r == nil {
		return nil
	}
	return &ElseifClause{baseNode[*googlesql.ASTElseifClause]{raw: r}}
}

func (n *ElseifClause) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

func (n *ElseifClause) Body() *StatementList {
	return newASTStatementList(must(n.raw.Body()))
}

// ElseifClauseList wraps *googlesql.ASTElseifClauseList (via raw children).
type ElseifClauseList struct {
	baseNode[*googlesql.ASTElseifClauseList]
}

func newASTElseifClauseList(r *googlesql.ASTElseifClauseList) *ElseifClauseList {
	if r == nil {
		return nil
	}
	return &ElseifClauseList{baseNode[*googlesql.ASTElseifClauseList]{raw: r}}
}

// Clauses returns all *ElseifClause children.
func (n *ElseifClauseList) Clauses() []*ElseifClause {
	var result []*ElseifClause
	for _, c := range n.Children() {
		if e, ok := c.(*ElseifClause); ok {
			result = append(result, e)
		}
	}
	return result
}

// IfStatement wraps *googlesql.ASTIfStatement.
type IfStatement struct {
	baseNode[*googlesql.ASTIfStatement]
}

func newASTIfStatement(r *googlesql.ASTIfStatement) *IfStatement {
	if r == nil {
		return nil
	}
	return &IfStatement{baseNode[*googlesql.ASTIfStatement]{raw: r}}
}
func (n *IfStatement) isStatement() {}

func (n *IfStatement) Condition() ExpressionNode {
	return wrapExpr(must(n.raw.Condition()))
}

func (n *IfStatement) ThenList() *StatementList {
	return newASTStatementList(must(n.raw.ThenList()))
}

func (n *IfStatement) ElseifClauses() *ElseifClauseList {
	return newASTElseifClauseList(must(n.raw.ElseifClauses()))
}

func (n *IfStatement) ElseList() *StatementList {
	return newASTStatementList(must(n.raw.ElseList()))
}

// ParameterAssignment wraps *googlesql.ASTParameterAssignment.
type ParameterAssignment struct {
	baseNode[*googlesql.ASTParameterAssignment]
}

func newASTParameterAssignment(r *googlesql.ASTParameterAssignment) *ParameterAssignment {
	if r == nil {
		return nil
	}
	return &ParameterAssignment{baseNode[*googlesql.ASTParameterAssignment]{raw: r}}
}
func (n *ParameterAssignment) isStatement() {}

func (n *ParameterAssignment) Parameter() *ParameterExpr {
	return newASTParameterExpr(must(n.raw.Parameter()))
}

func (n *ParameterAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

type RaiseStatement struct {
	baseNode[*googlesql.ASTRaiseStatement]
}

func newASTRaiseStatement(r *googlesql.ASTRaiseStatement) *RaiseStatement {
	if r == nil {
		return nil
	}
	return &RaiseStatement{baseNode[*googlesql.ASTRaiseStatement]{raw: r}}
}
func (n *RaiseStatement) isStatement() {}

func (n *RaiseStatement) Message() ExpressionNode {
	return wrapExpr(must(n.raw.Message()))
}

// ReturnStatement wraps *googlesql.ASTReturnStatement.
type ReturnStatement struct {
	baseNode[*googlesql.ASTReturnStatement]
}

func newASTReturnStatement(r *googlesql.ASTReturnStatement) *ReturnStatement {
	if r == nil {
		return nil
	}
	return &ReturnStatement{baseNode[*googlesql.ASTReturnStatement]{raw: r}}
}
func (n *ReturnStatement) isStatement() {}

// SystemVariableAssignment wraps *googlesql.ASTSystemVariableAssignment.
type SystemVariableAssignment struct {
	baseNode[*googlesql.ASTSystemVariableAssignment]
}

func newASTSystemVariableAssignment(r *googlesql.ASTSystemVariableAssignment) *SystemVariableAssignment {
	if r == nil {
		return nil
	}
	return &SystemVariableAssignment{baseNode[*googlesql.ASTSystemVariableAssignment]{raw: r}}
}
func (n *SystemVariableAssignment) isStatement() {}

func (n *SystemVariableAssignment) SystemVariable() *SystemVariableExpr {
	return newASTSystemVariableExpr(must(n.raw.SystemVariable()))
}

func (n *SystemVariableAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// SingleAssignment wraps *googlesql.ASTSingleAssignment.
type SingleAssignment struct {
	baseNode[*googlesql.ASTSingleAssignment]
}

func newASTSingleAssignment(r *googlesql.ASTSingleAssignment) *SingleAssignment {
	if r == nil {
		return nil
	}
	return &SingleAssignment{baseNode[*googlesql.ASTSingleAssignment]{raw: r}}
}
func (n *SingleAssignment) isStatement() {}

func (n *SingleAssignment) Variable() *Identifier {
	return newASTIdentifier(must(n.raw.Variable()))
}

func (n *SingleAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// VariableDeclaration wraps *googlesql.ASTVariableDeclaration.
type VariableDeclaration struct {
	baseNode[*googlesql.ASTVariableDeclaration]
}

func newASTVariableDeclaration(r *googlesql.ASTVariableDeclaration) *VariableDeclaration {
	if r == nil {
		return nil
	}
	return &VariableDeclaration{baseNode[*googlesql.ASTVariableDeclaration]{raw: r}}
}
func (n *VariableDeclaration) isStatement() {}

func (n *VariableDeclaration) VariableList() *IdentifierList {
	return newASTIdentifierList(must(n.raw.VariableList()))
}
func (n *VariableDeclaration) Type() TypeNode { return wrapType(must(n.raw.Type())) }
func (n *VariableDeclaration) DefaultValue() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultValue()))
}
