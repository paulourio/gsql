package sql

import (
	"github.com/goccy/go-googlesql"
)

// AlterAllRowAccessPoliciesStatement wraps *googlesql.ASTAlterAllRowAccessPoliciesStatement.
type AlterAllRowAccessPoliciesStatement struct {
	baseNode[*googlesql.ASTAlterAllRowAccessPoliciesStatement]
}

func newAlterAllRowAccessPoliciesStatement(r *googlesql.ASTAlterAllRowAccessPoliciesStatement) *AlterAllRowAccessPoliciesStatement {
	if r == nil {
		return nil
	}
	return &AlterAllRowAccessPoliciesStatement{baseNode[*googlesql.ASTAlterAllRowAccessPoliciesStatement]{raw: r}}
}

func (n *AlterAllRowAccessPoliciesStatement) isStatement() {}

// AlterDatabaseStatement wraps *googlesql.ASTAlterDatabaseStatement.
type AlterDatabaseStatement struct {
	baseNode[*googlesql.ASTAlterDatabaseStatement]
}

func newAlterDatabaseStatement(r *googlesql.ASTAlterDatabaseStatement) *AlterDatabaseStatement {
	if r == nil {
		return nil
	}
	return &AlterDatabaseStatement{baseNode[*googlesql.ASTAlterDatabaseStatement]{raw: r}}
}

func (n *AlterDatabaseStatement) isStatement() {}

func (n *AlterDatabaseStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterDatabaseStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterEntityStatement wraps *googlesql.ASTAlterEntityStatement.
type AlterEntityStatement struct {
	baseNode[*googlesql.ASTAlterEntityStatement]
}

func newAlterEntityStatement(r *googlesql.ASTAlterEntityStatement) *AlterEntityStatement {
	if r == nil {
		return nil
	}
	return &AlterEntityStatement{baseNode[*googlesql.ASTAlterEntityStatement]{raw: r}}
}

func (n *AlterEntityStatement) isStatement() {}

func (n *AlterEntityStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterEntityStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterMaterializedViewStatement wraps *googlesql.ASTAlterMaterializedViewStatement.
type AlterMaterializedViewStatement struct {
	baseNode[*googlesql.ASTAlterMaterializedViewStatement]
}

func newAlterMaterializedViewStatement(r *googlesql.ASTAlterMaterializedViewStatement) *AlterMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &AlterMaterializedViewStatement{baseNode[*googlesql.ASTAlterMaterializedViewStatement]{raw: r}}
}

func (n *AlterMaterializedViewStatement) isStatement() {}

func (n *AlterMaterializedViewStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterMaterializedViewStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterPrivilegeRestrictionStatement wraps *googlesql.ASTAlterPrivilegeRestrictionStatement.
type AlterPrivilegeRestrictionStatement struct {
	baseNode[*googlesql.ASTAlterPrivilegeRestrictionStatement]
}

func newAlterPrivilegeRestrictionStatement(r *googlesql.ASTAlterPrivilegeRestrictionStatement) *AlterPrivilegeRestrictionStatement {
	if r == nil {
		return nil
	}
	return &AlterPrivilegeRestrictionStatement{baseNode[*googlesql.ASTAlterPrivilegeRestrictionStatement]{raw: r}}
}

func (n *AlterPrivilegeRestrictionStatement) isStatement() {}

// AlterRowAccessPolicyStatement wraps *googlesql.ASTAlterRowAccessPolicyStatement.
type AlterRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTAlterRowAccessPolicyStatement]
}

func newAlterRowAccessPolicyStatement(r *googlesql.ASTAlterRowAccessPolicyStatement) *AlterRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &AlterRowAccessPolicyStatement{baseNode[*googlesql.ASTAlterRowAccessPolicyStatement]{raw: r}}
}

func (n *AlterRowAccessPolicyStatement) isStatement() {}

func (n *AlterRowAccessPolicyStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterRowAccessPolicyStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterSchemaStatement wraps *googlesql.ASTAlterSchemaStatement.
type AlterSchemaStatement struct {
	baseNode[*googlesql.ASTAlterSchemaStatement]
}

func newAlterSchemaStatement(r *googlesql.ASTAlterSchemaStatement) *AlterSchemaStatement {
	if r == nil {
		return nil
	}
	return &AlterSchemaStatement{baseNode[*googlesql.ASTAlterSchemaStatement]{raw: r}}
}

func (n *AlterSchemaStatement) isStatement() {}

func (n *AlterSchemaStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterSchemaStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterTableStatement wraps *googlesql.ASTAlterTableStatement.
type AlterTableStatement struct {
	baseNode[*googlesql.ASTAlterTableStatement]
}

func newAlterTableStatement(r *googlesql.ASTAlterTableStatement) *AlterTableStatement {
	if r == nil {
		return nil
	}
	return &AlterTableStatement{baseNode[*googlesql.ASTAlterTableStatement]{raw: r}}
}

func (n *AlterTableStatement) isStatement() {}

func (n *AlterTableStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterTableStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterViewStatement wraps *googlesql.ASTAlterViewStatement.
type AlterViewStatement struct {
	baseNode[*googlesql.ASTAlterViewStatement]
}

func newAlterViewStatement(r *googlesql.ASTAlterViewStatement) *AlterViewStatement {
	if r == nil {
		return nil
	}
	return &AlterViewStatement{baseNode[*googlesql.ASTAlterViewStatement]{raw: r}}
}

func (n *AlterViewStatement) isStatement() {}

func (n *AlterViewStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterViewStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

type AssertStatement struct {
	baseNode[*googlesql.ASTAssertStatement]
}

func newAssertStatement(r *googlesql.ASTAssertStatement) *AssertStatement {
	if r == nil {
		return nil
	}
	return &AssertStatement{baseNode[*googlesql.ASTAssertStatement]{raw: r}}
}

func (n *AssertStatement) isStatement() {}

func (n *AssertStatement) Expr() ExpressionNode {
	return wrapExpr(must(n.raw.Expr()))
}

func (n *AssertStatement) Description() *StringLiteral {
	return newStringLiteral(must(n.raw.Description()))
}

// AssignmentFromStruct wraps *googlesql.ASTAssignmentFromStruct.
type AssignmentFromStruct struct {
	baseNode[*googlesql.ASTAssignmentFromStruct]
}

func newAssignmentFromStruct(r *googlesql.ASTAssignmentFromStruct) *AssignmentFromStruct {
	if r == nil {
		return nil
	}
	return &AssignmentFromStruct{baseNode[*googlesql.ASTAssignmentFromStruct]{raw: r}}
}

func (n *AssignmentFromStruct) isStatement() {}

func (n *AssignmentFromStruct) Variables() *IdentifierList {
	return newIdentifierList(must(n.raw.Variables()))
}

func (n *AssignmentFromStruct) StructExpression() ExpressionNode {
	return wrapExpr(must(n.raw.StructExpression()))
}

// BeginEndBlock wraps *googlesql.ASTBeginEndBlock.
type BeginEndBlock struct {
	baseNode[*googlesql.ASTBeginEndBlock]
}

func newBeginEndBlock(r *googlesql.ASTBeginEndBlock) *BeginEndBlock {
	if r == nil {
		return nil
	}
	return &BeginEndBlock{baseNode[*googlesql.ASTBeginEndBlock]{raw: r}}
}

func (n *BeginEndBlock) isStatement() {}

func (n *BeginEndBlock) Label() *Label {
	return newLabel(must(n.raw.Label()))
}

func (n *BeginEndBlock) StatementListNode() *StatementList {
	return newStatementList(must(n.raw.StatementListNode()))
}

func (n *BeginEndBlock) HasExceptionHandler() bool { return must(n.raw.HasExceptionHandler()) }

func (n *BeginEndBlock) HandlerList() *ExceptionHandlerList {
	return newExceptionHandlerList(must(n.raw.HandlerList()))
}

// BeginStatement wraps *googlesql.ASTBeginStatement.
type BeginStatement struct {
	baseNode[*googlesql.ASTBeginStatement]
}

func newBeginStatement(r *googlesql.ASTBeginStatement) *BeginStatement {
	if r == nil {
		return nil
	}
	return &BeginStatement{baseNode[*googlesql.ASTBeginStatement]{raw: r}}
}

func (n *BeginStatement) isStatement() {}

// BreakStatement wraps *googlesql.ASTBreakStatement.
type BreakStatement struct {
	baseNode[*googlesql.ASTBreakStatement]
}

func newBreakStatement(r *googlesql.ASTBreakStatement) *BreakStatement {
	if r == nil {
		return nil
	}
	return &BreakStatement{baseNode[*googlesql.ASTBreakStatement]{raw: r}}
}

func (n *BreakStatement) isStatement() {}
func (n *BreakStatement) Keyword() BreakContinueKeyword {
	return must(n.raw.Keyword())
}

func (n *BreakStatement) GetKeywordText() string {
	return must(n.raw.GetKeywordText())
}

func (n *BreakStatement) Label() *Label {
	return newLabel(must(n.raw.Label()))
}

// CallStatement wraps *googlesql.ASTCallStatement.
type CallStatement struct {
	baseNode[*googlesql.ASTCallStatement]
}

func newCallStatement(r *googlesql.ASTCallStatement) *CallStatement {
	if r == nil {
		return nil
	}
	return &CallStatement{baseNode[*googlesql.ASTCallStatement]{raw: r}}
}

func (n *CallStatement) isStatement() {}

func (n *CallStatement) ProcedureName() *PathExpression {
	return newPathExpression(must(n.raw.ProcedureName()))
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

func newCommitStatement(r *googlesql.ASTCommitStatement) *CommitStatement {
	if r == nil {
		return nil
	}
	return &CommitStatement{baseNode[*googlesql.ASTCommitStatement]{raw: r}}
}

func (n *CommitStatement) isStatement() {}

// ContinueStatement wraps *googlesql.ASTContinueStatement.
type ContinueStatement struct {
	baseNode[*googlesql.ASTContinueStatement]
}

func newContinueStatement(r *googlesql.ASTContinueStatement) *ContinueStatement {
	if r == nil {
		return nil
	}
	return &ContinueStatement{baseNode[*googlesql.ASTContinueStatement]{raw: r}}
}

func (n *ContinueStatement) isStatement() {}
func (n *ContinueStatement) Keyword() BreakContinueKeyword {
	return must(n.raw.Keyword())
}

func (n *ContinueStatement) GetKeywordText() string {
	return must(n.raw.GetKeywordText())
}

func (n *ContinueStatement) Label() *Label {
	return newLabel(must(n.raw.Label()))
}

// CreateExternalTableStatement wraps *googlesql.ASTCreateExternalTableStatement.
type CreateExternalTableStatement struct {
	baseNode[*googlesql.ASTCreateExternalTableStatement]
}

func newCreateExternalTableStatement(r *googlesql.ASTCreateExternalTableStatement) *CreateExternalTableStatement {
	if r == nil {
		return nil
	}
	return &CreateExternalTableStatement{baseNode[*googlesql.ASTCreateExternalTableStatement]{raw: r}}
}

func (n *CreateExternalTableStatement) isStatement() {}

func (n *CreateExternalTableStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateExternalTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateExternalTableStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateExternalTableStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateExternalTableStatement) TableElementList() *TableElementList {
	return newTableElementList(must(n.raw.TableElementList()))
}

func (n *CreateExternalTableStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

func (n *CreateExternalTableStatement) WithPartitionColumnsClause() *WithPartitionColumnsClause {
	return newWithPartitionColumnsClause(must(n.raw.WithPartitionColumnsClause()))
}

func (n *CreateExternalTableStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateFunctionStatement wraps *googlesql.ASTCreateFunctionStatement.
type CreateFunctionStatement struct {
	baseNode[*googlesql.ASTCreateFunctionStatement]
}

func newCreateFunctionStatement(r *googlesql.ASTCreateFunctionStatement) *CreateFunctionStatement {
	if r == nil {
		return nil
	}
	return &CreateFunctionStatement{baseNode[*googlesql.ASTCreateFunctionStatement]{raw: r}}
}

func (n *CreateFunctionStatement) isStatement() {}

func (n *CreateFunctionStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateFunctionStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateFunctionStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateFunctionStatement) IsAggregate() bool { return must(n.raw.IsAggregate()) }

func (n *CreateFunctionStatement) IsRemote() bool { return must(n.raw.IsRemote()) }

func (n *CreateFunctionStatement) FunctionDeclaration() *FunctionDeclaration {
	return newFunctionDeclaration(must(n.raw.FunctionDeclaration()))
}

func (n *CreateFunctionStatement) ReturnType() TypeNode { return wrapType(must(n.raw.ReturnType())) }

func (n *CreateFunctionStatement) DeterminismLevel() DeterminismLevel {
	return must(n.raw.DeterminismLevel())
}

func (n *CreateFunctionStatement) Language() *Identifier {
	return newIdentifier(must(n.raw.Language()))
}

func (n *CreateFunctionStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateFunctionStatement) Code() *StringLiteral {
	return newStringLiteral(must(n.raw.Code()))
}

func (n *CreateFunctionStatement) SQLFunctionBody() *SQLFunctionBody {
	return newSQLFunctionBody(must(n.raw.SqlFunctionBody()))
}

func (n *CreateFunctionStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// CreateMaterializedViewStatement wraps *googlesql.ASTCreateMaterializedViewStatement.
type CreateMaterializedViewStatement struct {
	baseNode[*googlesql.ASTCreateMaterializedViewStatement]
}

func newCreateMaterializedViewStatement(r *googlesql.ASTCreateMaterializedViewStatement) *CreateMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &CreateMaterializedViewStatement{baseNode[*googlesql.ASTCreateMaterializedViewStatement]{raw: r}}
}

func (n *CreateMaterializedViewStatement) isStatement() {}

func (n *CreateMaterializedViewStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateMaterializedViewStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateMaterializedViewStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateMaterializedViewStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateMaterializedViewStatement) Recursive() bool { return must(n.raw.Recursive()) }

func (n *CreateMaterializedViewStatement) PartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.PartitionBy()))
}

func (n *CreateMaterializedViewStatement) ClusterBy() *ClusterBy {
	return newClusterBy(must(n.raw.ClusterBy()))
}

func (n *CreateMaterializedViewStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateMaterializedViewStatement) Query() *Query {
	return newQuery(must(n.raw.Query()))
}

func (n *CreateMaterializedViewStatement) ReplicaSource() *PathExpression {
	return newPathExpression(must(n.raw.ReplicaSource()))
}

// CreateProcedureStatement wraps *googlesql.ASTCreateProcedureStatement.
type CreateProcedureStatement struct {
	baseNode[*googlesql.ASTCreateProcedureStatement]
}

func newCreateProcedureStatement(r *googlesql.ASTCreateProcedureStatement) *CreateProcedureStatement {
	if r == nil {
		return nil
	}
	return &CreateProcedureStatement{baseNode[*googlesql.ASTCreateProcedureStatement]{raw: r}}
}

func (n *CreateProcedureStatement) isStatement() {}

func (n *CreateProcedureStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateProcedureStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateProcedureStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateProcedureStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateProcedureStatement) Parameters() *FunctionParameters {
	return newFunctionParameters(must(n.raw.Parameters()))
}

func (n *CreateProcedureStatement) ExternalSecurity() SQLSecurity {
	return must(n.raw.ExternalSecurity())
}

func (n *CreateProcedureStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateProcedureStatement) Body() *Script {
	return newScript(must(n.raw.Body()))
}

// CreateRowAccessPolicyStatement wraps *googlesql.ASTCreateRowAccessPolicyStatement.
type CreateRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTCreateRowAccessPolicyStatement]
}

func newCreateRowAccessPolicyStatement(r *googlesql.ASTCreateRowAccessPolicyStatement) *CreateRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &CreateRowAccessPolicyStatement{baseNode[*googlesql.ASTCreateRowAccessPolicyStatement]{raw: r}}
}

func (n *CreateRowAccessPolicyStatement) isStatement() {}

func (n *CreateRowAccessPolicyStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateRowAccessPolicyStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateRowAccessPolicyStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateRowAccessPolicyStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateRowAccessPolicyStatement) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

func (n *CreateRowAccessPolicyStatement) GrantTo() *GrantToClause {
	return newGrantToClause(must(n.raw.GrantTo()))
}

func (n *CreateRowAccessPolicyStatement) FilterUsing() *FilterUsingClause {
	return newFilterUsingClause(must(n.raw.FilterUsing()))
}

// CreateSchemaStatement wraps *googlesql.ASTCreateSchemaStatement.
type CreateSchemaStatement struct {
	baseNode[*googlesql.ASTCreateSchemaStatement]
}

func newCreateSchemaStatement(r *googlesql.ASTCreateSchemaStatement) *CreateSchemaStatement {
	if r == nil {
		return nil
	}
	return &CreateSchemaStatement{baseNode[*googlesql.ASTCreateSchemaStatement]{raw: r}}
}

func (n *CreateSchemaStatement) isStatement() {}

func (n *CreateSchemaStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateSchemaStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateSchemaStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateSchemaStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateSchemaStatement) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *CreateSchemaStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateSnapshotTableStatement wraps *googlesql.ASTCreateSnapshotTableStatement.
type CreateSnapshotTableStatement struct {
	baseNode[*googlesql.ASTCreateSnapshotTableStatement]
}

func newCreateSnapshotTableStatement(r *googlesql.ASTCreateSnapshotTableStatement) *CreateSnapshotTableStatement {
	if r == nil {
		return nil
	}
	return &CreateSnapshotTableStatement{baseNode[*googlesql.ASTCreateSnapshotTableStatement]{raw: r}}
}

func (n *CreateSnapshotTableStatement) isStatement() {}

func (n *CreateSnapshotTableStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateSnapshotTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateSnapshotTableStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateSnapshotTableStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *CreateSnapshotTableStatement) CloneDataSource() *CloneDataSource {
	return newCloneDataSource(must(n.raw.CloneDataSource()))
}

func (n *CreateSnapshotTableStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateTableFunctionStatement wraps *googlesql.ASTCreateTableFunctionStatement.
type CreateTableFunctionStatement struct {
	baseNode[*googlesql.ASTCreateTableFunctionStatement]
}

func newCreateTableFunctionStatement(r *googlesql.ASTCreateTableFunctionStatement) *CreateTableFunctionStatement {
	if r == nil {
		return nil
	}
	return &CreateTableFunctionStatement{baseNode[*googlesql.ASTCreateTableFunctionStatement]{raw: r}}
}

func (n *CreateTableFunctionStatement) isStatement() {}

func (n *CreateTableFunctionStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateTableFunctionStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateTableFunctionStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateTableFunctionStatement) FunctionDeclaration() *FunctionDeclaration {
	return newFunctionDeclaration(must(n.raw.FunctionDeclaration()))
}

func (n *CreateTableFunctionStatement) ReturnTvfSchema() Node {
	return Wrap(must(n.raw.ReturnTvfSchema()))
}

func (n *CreateTableFunctionStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateTableFunctionStatement) Query() *Query {
	return newQuery(must(n.raw.Query()))
}

// CreateTableStatement wraps *googlesql.ASTCreateTableStatement.
type CreateTableStatement struct {
	baseNode[*googlesql.ASTCreateTableStatement]
}

func newCreateTableStatement(r *googlesql.ASTCreateTableStatement) *CreateTableStatement {
	if r == nil {
		return nil
	}
	return &CreateTableStatement{baseNode[*googlesql.ASTCreateTableStatement]{raw: r}}
}

func (n *CreateTableStatement) isStatement() {}

func (n *CreateTableStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateTableStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateTableStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateTableStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateTableStatement) TableElementList() *TableElementList {
	return newTableElementList(must(n.raw.TableElementList()))
}

func (n *CreateTableStatement) CopyDataSource() *CopyDataSource {
	return newCopyDataSource(must(n.raw.CopyDataSource()))
}

func (n *CreateTableStatement) CloneDataSource() *CloneDataSource {
	return newCloneDataSource(must(n.raw.CloneDataSource()))
}

func (n *CreateTableStatement) LikeTableName() *PathExpression {
	return newPathExpression(must(n.raw.LikeTableName()))
}

func (n *CreateTableStatement) Collate() *Collate {
	return newCollate(must(n.raw.Collate()))
}

func (n *CreateTableStatement) PartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.PartitionBy()))
}

func (n *CreateTableStatement) ClusterBy() *ClusterBy {
	return newClusterBy(must(n.raw.ClusterBy()))
}

func (n *CreateTableStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

func (n *CreateTableStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateTableStatement) Query() *Query { return newQuery(must(n.raw.Query())) }

// CreateViewStatement wraps *googlesql.ASTCreateViewStatement.
type CreateViewStatement struct {
	baseNode[*googlesql.ASTCreateViewStatement]
}

func newCreateViewStatement(r *googlesql.ASTCreateViewStatement) *CreateViewStatement {
	if r == nil {
		return nil
	}
	return &CreateViewStatement{baseNode[*googlesql.ASTCreateViewStatement]{raw: r}}
}

func (n *CreateViewStatement) isStatement() {}

func (n *CreateViewStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }

func (n *CreateViewStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }

func (n *CreateViewStatement) Scope() Scope { return must(n.raw.Scope()) }

func (n *CreateViewStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateViewStatement) ColumnWithOptionsList() *ColumnWithOptionsList {
	return newColumnWithOptionsList(must(n.raw.ColumnWithOptionsList()))
}

func (n *CreateViewStatement) SQLSecurity() SQLSecurity { return must(n.raw.SqlSecurity()) }

func (n *CreateViewStatement) Recursive() bool { return must(n.raw.Recursive()) }

func (n *CreateViewStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateViewStatement) Query() *Query { return newQuery(must(n.raw.Query())) }

type DeleteStatement struct {
	baseNode[*googlesql.ASTDeleteStatement]
}

func newDeleteStatement(r *googlesql.ASTDeleteStatement) *DeleteStatement {
	if r == nil {
		return nil
	}
	return &DeleteStatement{baseNode[*googlesql.ASTDeleteStatement]{raw: r}}
}

func (n *DeleteStatement) isStatement() {}

func (n *DeleteStatement) TargetPath() Node {
	return Wrap(must(n.raw.TargetPath()))
}

func (n *DeleteStatement) Alias() *Alias {
	return newAlias(must(n.raw.Alias()))
}

func (n *DeleteStatement) Hint() *Hint {
	return newHint(must(n.raw.Hint()))
}

func (n *DeleteStatement) Offset() *WithOffset {
	return newWithOffset(must(n.raw.Offset()))
}

func (n *DeleteStatement) Where() ExpressionNode {
	return wrapExpr(must(n.raw.Where()))
}

func (n *DeleteStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *DeleteStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

// DropAllRowAccessPoliciesStatement wraps *googlesql.ASTDropAllRowAccessPoliciesStatement.
type DropAllRowAccessPoliciesStatement struct {
	baseNode[*googlesql.ASTDropAllRowAccessPoliciesStatement]
}

func newDropAllRowAccessPoliciesStatement(r *googlesql.ASTDropAllRowAccessPoliciesStatement) *DropAllRowAccessPoliciesStatement {
	if r == nil {
		return nil
	}
	return &DropAllRowAccessPoliciesStatement{baseNode[*googlesql.ASTDropAllRowAccessPoliciesStatement]{raw: r}}
}

func (n *DropAllRowAccessPoliciesStatement) isStatement() {}

func (n *DropAllRowAccessPoliciesStatement) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}

// DropEntityStatement wraps *googlesql.ASTDropEntityStatement.
type DropEntityStatement struct {
	baseNode[*googlesql.ASTDropEntityStatement]
}

func newDropEntityStatement(r *googlesql.ASTDropEntityStatement) *DropEntityStatement {
	if r == nil {
		return nil
	}
	return &DropEntityStatement{baseNode[*googlesql.ASTDropEntityStatement]{raw: r}}
}

func (n *DropEntityStatement) isStatement() {}

func (n *DropEntityStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropEntityStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

// DropFunctionStatement wraps *googlesql.ASTDropFunctionStatement.
type DropFunctionStatement struct {
	baseNode[*googlesql.ASTDropFunctionStatement]
}

func newDropFunctionStatement(r *googlesql.ASTDropFunctionStatement) *DropFunctionStatement {
	if r == nil {
		return nil
	}
	return &DropFunctionStatement{baseNode[*googlesql.ASTDropFunctionStatement]{raw: r}}
}

func (n *DropFunctionStatement) isStatement() {}

func (n *DropFunctionStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropFunctionStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

// DropMaterializedViewStatement wraps *googlesql.ASTDropMaterializedViewStatement.
type DropMaterializedViewStatement struct {
	baseNode[*googlesql.ASTDropMaterializedViewStatement]
}

func newDropMaterializedViewStatement(r *googlesql.ASTDropMaterializedViewStatement) *DropMaterializedViewStatement {
	if r == nil {
		return nil
	}
	return &DropMaterializedViewStatement{baseNode[*googlesql.ASTDropMaterializedViewStatement]{raw: r}}
}

func (n *DropMaterializedViewStatement) isStatement() {}

func (n *DropMaterializedViewStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropMaterializedViewStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

// DropPrivilegeRestrictionStatement wraps *googlesql.ASTDropPrivilegeRestrictionStatement.
type DropPrivilegeRestrictionStatement struct {
	baseNode[*googlesql.ASTDropPrivilegeRestrictionStatement]
}

func newDropPrivilegeRestrictionStatement(r *googlesql.ASTDropPrivilegeRestrictionStatement) *DropPrivilegeRestrictionStatement {
	if r == nil {
		return nil
	}
	return &DropPrivilegeRestrictionStatement{baseNode[*googlesql.ASTDropPrivilegeRestrictionStatement]{raw: r}}
}

func (n *DropPrivilegeRestrictionStatement) isStatement() {}

// DropRowAccessPolicyStatement wraps *googlesql.ASTDropRowAccessPolicyStatement.
type DropRowAccessPolicyStatement struct {
	baseNode[*googlesql.ASTDropRowAccessPolicyStatement]
}

func newDropRowAccessPolicyStatement(r *googlesql.ASTDropRowAccessPolicyStatement) *DropRowAccessPolicyStatement {
	if r == nil {
		return nil
	}
	return &DropRowAccessPolicyStatement{baseNode[*googlesql.ASTDropRowAccessPolicyStatement]{raw: r}}
}

func (n *DropRowAccessPolicyStatement) isStatement() {}

func (n *DropRowAccessPolicyStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropRowAccessPolicyStatement) Name() *Identifier {
	return newIdentifier(must(n.raw.Name()))
}

func (n *DropRowAccessPolicyStatement) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}

// DropSearchIndexStatement wraps *googlesql.ASTDropSearchIndexStatement.
type DropSearchIndexStatement struct {
	baseNode[*googlesql.ASTDropSearchIndexStatement]
}

func newDropSearchIndexStatement(r *googlesql.ASTDropSearchIndexStatement) *DropSearchIndexStatement {
	if r == nil {
		return nil
	}
	return &DropSearchIndexStatement{baseNode[*googlesql.ASTDropSearchIndexStatement]{raw: r}}
}

func (n *DropSearchIndexStatement) isStatement() {}

func (n *DropSearchIndexStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropSearchIndexStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

// DropSnapshotTableStatement wraps *googlesql.ASTDropSnapshotTableStatement.
type DropSnapshotTableStatement struct {
	baseNode[*googlesql.ASTDropSnapshotTableStatement]
}

func newDropSnapshotTableStatement(r *googlesql.ASTDropSnapshotTableStatement) *DropSnapshotTableStatement {
	if r == nil {
		return nil
	}
	return &DropSnapshotTableStatement{baseNode[*googlesql.ASTDropSnapshotTableStatement]{raw: r}}
}

func (n *DropSnapshotTableStatement) isStatement() {}

func (n *DropSnapshotTableStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropSnapshotTableStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

// DropStatement wraps *googlesql.ASTDropStatement.
type DropStatement struct {
	baseNode[*googlesql.ASTDropStatement]
}

func newDropStatement(r *googlesql.ASTDropStatement) *DropStatement {
	if r == nil {
		return nil
	}
	return &DropStatement{baseNode[*googlesql.ASTDropStatement]{raw: r}}
}

func (n *DropStatement) isStatement() {}

func (n *DropStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *DropStatement) DropMode() DropMode { return must(n.raw.DropMode()) }

func (n *DropStatement) SchemaObjectKind() SchemaObjectKind {
	return must(n.raw.SchemaObjectKind())
}

// DropTableFunctionStatement wraps *googlesql.ASTDropTableFunctionStatement.
type DropTableFunctionStatement struct {
	baseNode[*googlesql.ASTDropTableFunctionStatement]
}

func newDropTableFunctionStatement(r *googlesql.ASTDropTableFunctionStatement) *DropTableFunctionStatement {
	if r == nil {
		return nil
	}
	return &DropTableFunctionStatement{baseNode[*googlesql.ASTDropTableFunctionStatement]{raw: r}}
}

func (n *DropTableFunctionStatement) isStatement() {}

func (n *DropTableFunctionStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }

func (n *DropTableFunctionStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

// ExecuteImmediateStatement wraps *googlesql.ASTExecuteImmediateStatement.
type ExecuteImmediateStatement struct {
	baseNode[*googlesql.ASTExecuteImmediateStatement]
}

func newExecuteImmediateStatement(r *googlesql.ASTExecuteImmediateStatement) *ExecuteImmediateStatement {
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
	return newExecuteIntoClause(must(n.raw.IntoClause()))
}

func (n *ExecuteImmediateStatement) UsingClause() *ExecuteUsingClause {
	return newExecuteUsingClause(must(n.raw.UsingClause()))
}

// HintedStatement wraps *googlesql.ASTHintedStatement.
type HintedStatement struct {
	baseNode[*googlesql.ASTHintedStatement]
}

func newHintedStatement(r *googlesql.ASTHintedStatement) *HintedStatement {
	if r == nil {
		return nil
	}
	return &HintedStatement{baseNode[*googlesql.ASTHintedStatement]{raw: r}}
}

func (n *HintedStatement) isStatement() {}

func (n *HintedStatement) Hint() *Hint { return newHint(must(n.raw.Hint())) }

func (n *HintedStatement) Statement() StatementNode {
	return wrapStmt(must(n.raw.Statement()))
}

// IfStatement wraps *googlesql.ASTIfStatement.
type IfStatement struct {
	baseNode[*googlesql.ASTIfStatement]
}

func newIfStatement(r *googlesql.ASTIfStatement) *IfStatement {
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
	return newStatementList(must(n.raw.ThenList()))
}

func (n *IfStatement) ElseifClauses() *ElseifClauseList {
	return newElseifClauseList(must(n.raw.ElseifClauses()))
}

func (n *IfStatement) ElseList() *StatementList {
	return newStatementList(must(n.raw.ElseList()))
}

// InsertStatement wraps *googlesql.ASTInsertStatement.
type InsertStatement struct {
	baseNode[*googlesql.ASTInsertStatement]
}

func newInsertStatement(r *googlesql.ASTInsertStatement) *InsertStatement {
	if r == nil {
		return nil
	}
	return &InsertStatement{baseNode[*googlesql.ASTInsertStatement]{raw: r}}
}

func (n *InsertStatement) isStatement() {}

func (n *InsertStatement) TargetPath() Node { return Wrap(must(n.raw.TargetPath())) }

func (n *InsertStatement) ColumnList() *ColumnList {
	return newColumnList(must(n.raw.ColumnList()))
}

func (n *InsertStatement) Query() *Query { return newQuery(must(n.raw.Query())) }

func (n *InsertStatement) Rows() *InsertValuesRowList {
	return newInsertValuesRowList(must(n.raw.Rows()))
}

func (n *InsertStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *InsertStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

// ForInStatement wraps *googlesql.ASTForInStatement.
type ForInStatement struct {
	baseNode[*googlesql.ASTForInStatement]
}

func newForInStatement(r *googlesql.ASTForInStatement) *ForInStatement {
	if r == nil {
		return nil
	}
	return &ForInStatement{baseNode[*googlesql.ASTForInStatement]{raw: r}}
}

func (n *ForInStatement) isStatement()          {}
func (n *ForInStatement) Label() *Label         { return newLabel(must(n.raw.Label())) }
func (n *ForInStatement) Variable() *Identifier { return newIdentifier(must(n.raw.Variable())) }
func (n *ForInStatement) Query() *Query         { return newQuery(must(n.raw.Query())) }
func (n *ForInStatement) Body() *StatementList  { return newStatementList(must(n.raw.Body())) }

// MergeStatement wraps *googlesql.ASTMergeStatement.
type MergeStatement struct {
	baseNode[*googlesql.ASTMergeStatement]
}

func newMergeStatement(r *googlesql.ASTMergeStatement) *MergeStatement {
	if r == nil {
		return nil
	}
	return &MergeStatement{baseNode[*googlesql.ASTMergeStatement]{raw: r}}
}

func (n *MergeStatement) isStatement() {}

func (n *MergeStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}

func (n *MergeStatement) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *MergeStatement) TableExpression() TableExpressionNode {
	return wrapTableExpr(must(n.raw.TableExpression()))
}

func (n *MergeStatement) MergeCondition() ExpressionNode {
	return wrapExpr(must(n.raw.MergeCondition()))
}

func (n *MergeStatement) WhenClauses() *MergeWhenClauseList {
	return newMergeWhenClauseList(must(n.raw.WhenClauses()))
}

// ParameterAssignment wraps *googlesql.ASTParameterAssignment.
type ParameterAssignment struct {
	baseNode[*googlesql.ASTParameterAssignment]
}

func newParameterAssignment(r *googlesql.ASTParameterAssignment) *ParameterAssignment {
	if r == nil {
		return nil
	}
	return &ParameterAssignment{baseNode[*googlesql.ASTParameterAssignment]{raw: r}}
}

func (n *ParameterAssignment) isStatement() {}

func (n *ParameterAssignment) Parameter() *ParameterExpr {
	return newParameterExpr(must(n.raw.Parameter()))
}

func (n *ParameterAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// QueryStatement wraps *googlesql.ASTQueryStatement.
type QueryStatement struct {
	baseNode[*googlesql.ASTQueryStatement]
}

func newQueryStatement(r *googlesql.ASTQueryStatement) *QueryStatement {
	if r == nil {
		return nil
	}
	return &QueryStatement{baseNode[*googlesql.ASTQueryStatement]{raw: r}}
}

func (n *QueryStatement) isStatement() {}

func (n *QueryStatement) Query() *Query {
	return newQuery(must(n.raw.Query()))
}

type RaiseStatement struct {
	baseNode[*googlesql.ASTRaiseStatement]
}

func newRaiseStatement(r *googlesql.ASTRaiseStatement) *RaiseStatement {
	if r == nil {
		return nil
	}
	return &RaiseStatement{baseNode[*googlesql.ASTRaiseStatement]{raw: r}}
}

func (n *RaiseStatement) isStatement() {}

func (n *RaiseStatement) Message() ExpressionNode {
	return wrapExpr(must(n.raw.Message()))
}

// RenameStatement wraps *googlesql.ASTRenameStatement.
type RenameStatement struct {
	baseNode[*googlesql.ASTRenameStatement]
}

func newRenameStatement(r *googlesql.ASTRenameStatement) *RenameStatement {
	if r == nil {
		return nil
	}
	return &RenameStatement{baseNode[*googlesql.ASTRenameStatement]{raw: r}}
}

func (n *RenameStatement) isStatement()             {}
func (n *RenameStatement) Identifier() *Identifier  { return newIdentifier(must(n.raw.Identifier())) }
func (n *RenameStatement) OldName() *PathExpression { return newPathExpression(must(n.raw.OldName())) }
func (n *RenameStatement) NewName() *PathExpression { return newPathExpression(must(n.raw.NewName())) }

// RepeatStatement wraps *googlesql.ASTRepeatStatement.
type RepeatStatement struct {
	baseNode[*googlesql.ASTRepeatStatement]
}

func newRepeatStatement(r *googlesql.ASTRepeatStatement) *RepeatStatement {
	if r == nil {
		return nil
	}
	return &RepeatStatement{baseNode[*googlesql.ASTRepeatStatement]{raw: r}}
}

func (n *RepeatStatement) isStatement() {}

func (n *RepeatStatement) Label() *Label {
	return newLabel(must(n.raw.Label()))
}

func (n *RepeatStatement) Body() *StatementList {
	return newStatementList(must(n.raw.Body()))
}

func (n *RepeatStatement) UntilClause() *UntilClause {
	return newUntilClause(must(n.raw.UntilClause()))
}

// ReturnStatement wraps *googlesql.ASTReturnStatement.
type ReturnStatement struct {
	baseNode[*googlesql.ASTReturnStatement]
}

func newReturnStatement(r *googlesql.ASTReturnStatement) *ReturnStatement {
	if r == nil {
		return nil
	}
	return &ReturnStatement{baseNode[*googlesql.ASTReturnStatement]{raw: r}}
}

func (n *ReturnStatement) isStatement() {}

// RollbackStatement wraps *googlesql.ASTRollbackStatement.
type RollbackStatement struct {
	baseNode[*googlesql.ASTRollbackStatement]
}

func newRollbackStatement(r *googlesql.ASTRollbackStatement) *RollbackStatement {
	if r == nil {
		return nil
	}
	return &RollbackStatement{baseNode[*googlesql.ASTRollbackStatement]{raw: r}}
}

func (n *RollbackStatement) isStatement() {}

// SingleAssignment wraps *googlesql.ASTSingleAssignment.
type SingleAssignment struct {
	baseNode[*googlesql.ASTSingleAssignment]
}

func newSingleAssignment(r *googlesql.ASTSingleAssignment) *SingleAssignment {
	if r == nil {
		return nil
	}
	return &SingleAssignment{baseNode[*googlesql.ASTSingleAssignment]{raw: r}}
}

func (n *SingleAssignment) isStatement() {}

func (n *SingleAssignment) Variable() *Identifier {
	return newIdentifier(must(n.raw.Variable()))
}

func (n *SingleAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// SystemVariableAssignment wraps *googlesql.ASTSystemVariableAssignment.
type SystemVariableAssignment struct {
	baseNode[*googlesql.ASTSystemVariableAssignment]
}

func newSystemVariableAssignment(r *googlesql.ASTSystemVariableAssignment) *SystemVariableAssignment {
	if r == nil {
		return nil
	}
	return &SystemVariableAssignment{baseNode[*googlesql.ASTSystemVariableAssignment]{raw: r}}
}

func (n *SystemVariableAssignment) isStatement() {}

func (n *SystemVariableAssignment) SystemVariable() *SystemVariableExpr {
	return newSystemVariableExpr(must(n.raw.SystemVariable()))
}

func (n *SystemVariableAssignment) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

// TruncateStatement wraps *googlesql.ASTTruncateStatement.
type TruncateStatement struct {
	baseNode[*googlesql.ASTTruncateStatement]
}

func newTruncateStatement(r *googlesql.ASTTruncateStatement) *TruncateStatement {
	if r == nil {
		return nil
	}
	return &TruncateStatement{baseNode[*googlesql.ASTTruncateStatement]{raw: r}}
}

func (n *TruncateStatement) isStatement() {}

func (n *TruncateStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}

func (n *TruncateStatement) Where() ExpressionNode { return wrapExpr(must(n.raw.Where())) }

type UpdateStatement struct {
	baseNode[*googlesql.ASTUpdateStatement]
}

func newUpdateStatement(r *googlesql.ASTUpdateStatement) *UpdateStatement {
	if r == nil {
		return nil
	}
	return &UpdateStatement{baseNode[*googlesql.ASTUpdateStatement]{raw: r}}
}

func (n *UpdateStatement) isStatement() {}

func (n *UpdateStatement) TargetPath() Node { return Wrap(must(n.raw.TargetPath())) }

func (n *UpdateStatement) Alias() *Alias { return newAlias(must(n.raw.Alias())) }

func (n *UpdateStatement) Offset() *WithOffset { return newWithOffset(must(n.raw.Offset())) }

func (n *UpdateStatement) UpdateItemList() *UpdateItemList {
	return newUpdateItemList(must(n.raw.UpdateItemList()))
}

func (n *UpdateStatement) FromClause() *FromClause { return newFromClause(must(n.raw.FromClause())) }

func (n *UpdateStatement) Where() ExpressionNode { return wrapExpr(must(n.raw.Where())) }

func (n *UpdateStatement) AssertRowsModified() *AssertRowsModified {
	return newAssertRowsModified(must(n.raw.AssertRowsModified()))
}

func (n *UpdateStatement) Returning() *ReturningClause {
	return newReturningClause(must(n.raw.Returning()))
}

// VariableDeclaration wraps *googlesql.ASTVariableDeclaration.
type VariableDeclaration struct {
	baseNode[*googlesql.ASTVariableDeclaration]
}

func newVariableDeclaration(r *googlesql.ASTVariableDeclaration) *VariableDeclaration {
	if r == nil {
		return nil
	}
	return &VariableDeclaration{baseNode[*googlesql.ASTVariableDeclaration]{raw: r}}
}

func (n *VariableDeclaration) isStatement() {}

func (n *VariableDeclaration) VariableList() *IdentifierList {
	return newIdentifierList(must(n.raw.VariableList()))
}

func (n *VariableDeclaration) Type() TypeNode { return wrapType(must(n.raw.Type())) }

func (n *VariableDeclaration) DefaultValue() ExpressionNode {
	return wrapExpr(must(n.raw.DefaultValue()))
}

// WhileStatement wraps *googlesql.ASTWhileStatement.
type WhileStatement struct {
	baseNode[*googlesql.ASTWhileStatement]
}

func newWhileStatement(r *googlesql.ASTWhileStatement) *WhileStatement {
	if r == nil {
		return nil
	}
	return &WhileStatement{baseNode[*googlesql.ASTWhileStatement]{raw: r}}
}

func (n *WhileStatement) isStatement()              {}
func (n *WhileStatement) IsLoopStatement() bool     { return must(n.raw.IsLoopStatement()) }
func (n *WhileStatement) Label() *Label             { return newLabel(must(n.raw.Label())) }
func (n *WhileStatement) Condition() ExpressionNode { return wrapExpr(must(n.raw.Condition())) }
func (n *WhileStatement) Body() *StatementList      { return newStatementList(must(n.raw.Body())) }

// StatementWithPipeOperators wraps *googlesql.ASTStatementWithPipeOperators.
type StatementWithPipeOperators struct {
	baseNode[*googlesql.ASTStatementWithPipeOperators]
}

func newStatementWithPipeOperators(r *googlesql.ASTStatementWithPipeOperators) *StatementWithPipeOperators {
	if r == nil {
		return nil
	}
	return &StatementWithPipeOperators{baseNode[*googlesql.ASTStatementWithPipeOperators]{raw: r}}
}

func (n *StatementWithPipeOperators) isStatement() {}

func (n *StatementWithPipeOperators) Statement() StatementNode {
	return wrapStmt(must(n.raw.Statement()))
}

func (n *StatementWithPipeOperators) PipeOperatorSuffix() *SubpipelineStatement {
	return newSubpipelineStatement(must(n.raw.PipeOperatorSuffix()))
}

// SubpipelineStatement wraps *googlesql.ASTSubpipelineStatement.
type SubpipelineStatement struct {
	baseNode[*googlesql.ASTSubpipelineStatement]
}

func newSubpipelineStatement(r *googlesql.ASTSubpipelineStatement) *SubpipelineStatement {
	if r == nil {
		return nil
	}
	return &SubpipelineStatement{baseNode[*googlesql.ASTSubpipelineStatement]{raw: r}}
}

func (n *SubpipelineStatement) isStatement() {}

func (n *SubpipelineStatement) Subpipeline() *Subpipeline {
	return newSubpipeline(must(n.raw.Subpipeline()))
}

// CaseStatement wraps *googlesql.ASTCaseStatement.
type CaseStatement struct {
	baseNode[*googlesql.ASTCaseStatement]
}

func newCaseStatement(r *googlesql.ASTCaseStatement) *CaseStatement {
	if r == nil {
		return nil
	}
	return &CaseStatement{baseNode[*googlesql.ASTCaseStatement]{raw: r}}
}

func (n *CaseStatement) isStatement() {}

func (n *CaseStatement) Expression() ExpressionNode {
	return wrapExpr(must(n.raw.Expression()))
}

func (n *CaseStatement) WhenThenClauses() *WhenThenClauseList {
	return newWhenThenClauseList(must(n.raw.WhenThenClauses()))
}

func (n *CaseStatement) ElseList() *StatementList {
	return newStatementList(must(n.raw.ElseList()))
}

// CreateIndexStatement wraps *googlesql.ASTCreateIndexStatement.
type CreateIndexStatement struct {
	baseNode[*googlesql.ASTCreateIndexStatement]
}

func newCreateIndexStatement(r *googlesql.ASTCreateIndexStatement) *CreateIndexStatement {
	if r == nil {
		return nil
	}
	return &CreateIndexStatement{baseNode[*googlesql.ASTCreateIndexStatement]{raw: r}}
}
func (n *CreateIndexStatement) isStatement()          {}
func (n *CreateIndexStatement) IsOrReplace() bool     { return must(n.raw.IsOrReplace()) }
func (n *CreateIndexStatement) IsIfNotExists() bool   { return must(n.raw.IsIfNotExists()) }
func (n *CreateIndexStatement) Scope() Scope          { return must(n.raw.Scope()) }
func (n *CreateIndexStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *CreateIndexStatement) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}

func (n *CreateIndexStatement) OptionalTableAlias() *Alias {
	return newAlias(must(n.raw.OptionalTableAlias()))
}
func (n *CreateIndexStatement) IndexItemList() Node { return Wrap(must(n.raw.IndexItemList())) }
func (n *CreateIndexStatement) IsSearch() bool      { return must(n.raw.IsSearch()) }
func (n *CreateIndexStatement) IsUnique() bool      { return must(n.raw.IsUnique()) }
func (n *CreateIndexStatement) IsVector() bool      { return must(n.raw.IsVector()) }
func (n *CreateIndexStatement) OptionalPartitionBy() *PartitionBy {
	return newPartitionBy(must(n.raw.OptionalPartitionBy()))
}

func (n *CreateIndexStatement) OptionalIndexStoringExpressions() Node {
	return Wrap(must(n.raw.OptionalIndexStoringExpressions()))
}

func (n *CreateIndexStatement) OptionalIndexUnnestExpressionList() Node {
	return Wrap(must(n.raw.OptionalIndexUnnestExpressionList()))
}

func (n *CreateIndexStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateModelStatement wraps *googlesql.ASTCreateModelStatement.
type CreateModelStatement struct {
	baseNode[*googlesql.ASTCreateModelStatement]
}

func newCreateModelStatement(r *googlesql.ASTCreateModelStatement) *CreateModelStatement {
	if r == nil {
		return nil
	}
	return &CreateModelStatement{baseNode[*googlesql.ASTCreateModelStatement]{raw: r}}
}
func (n *CreateModelStatement) isStatement()          {}
func (n *CreateModelStatement) IsOrReplace() bool     { return must(n.raw.IsOrReplace()) }
func (n *CreateModelStatement) IsIfNotExists() bool   { return must(n.raw.IsIfNotExists()) }
func (n *CreateModelStatement) Scope() Scope          { return must(n.raw.Scope()) }
func (n *CreateModelStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *CreateModelStatement) Query() *Query         { return newQuery(must(n.raw.Query())) }
func (n *CreateModelStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}
func (n *CreateModelStatement) IsRemote() bool          { return must(n.raw.IsRemote()) }
func (n *CreateModelStatement) TransformClause() Node   { return Wrap(must(n.raw.TransformClause())) }
func (n *CreateModelStatement) InputOutputClause() Node { return Wrap(must(n.raw.InputOutputClause())) }
func (n *CreateModelStatement) AliasedQueryList() *AliasedQueryList {
	return newAliasedQueryList(must(n.raw.AliasedQueryList()))
}

func (n *CreateModelStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// CreateSequenceStatement wraps *googlesql.ASTCreateSequenceStatement.
type CreateSequenceStatement struct {
	baseNode[*googlesql.ASTCreateSequenceStatement]
}

func newCreateSequenceStatement(r *googlesql.ASTCreateSequenceStatement) *CreateSequenceStatement {
	if r == nil {
		return nil
	}
	return &CreateSequenceStatement{baseNode[*googlesql.ASTCreateSequenceStatement]{raw: r}}
}
func (n *CreateSequenceStatement) isStatement()        {}
func (n *CreateSequenceStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateSequenceStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateSequenceStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateSequenceStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateSequenceStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateConnectionStatement wraps *googlesql.ASTCreateConnectionStatement.
type CreateConnectionStatement struct {
	baseNode[*googlesql.ASTCreateConnectionStatement]
}

func newCreateConnectionStatement(r *googlesql.ASTCreateConnectionStatement) *CreateConnectionStatement {
	if r == nil {
		return nil
	}
	return &CreateConnectionStatement{baseNode[*googlesql.ASTCreateConnectionStatement]{raw: r}}
}
func (n *CreateConnectionStatement) isStatement()        {}
func (n *CreateConnectionStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateConnectionStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateConnectionStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateConnectionStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateConnectionStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateDatabaseStatement wraps *googlesql.ASTCreateDatabaseStatement.
type CreateDatabaseStatement struct {
	baseNode[*googlesql.ASTCreateDatabaseStatement]
}

func newCreateDatabaseStatement(r *googlesql.ASTCreateDatabaseStatement) *CreateDatabaseStatement {
	if r == nil {
		return nil
	}
	return &CreateDatabaseStatement{baseNode[*googlesql.ASTCreateDatabaseStatement]{raw: r}}
}
func (n *CreateDatabaseStatement) isStatement() {}
func (n *CreateDatabaseStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateDatabaseStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateConstantStatement wraps *googlesql.ASTCreateConstantStatement.
type CreateConstantStatement struct {
	baseNode[*googlesql.ASTCreateConstantStatement]
}

func newCreateConstantStatement(r *googlesql.ASTCreateConstantStatement) *CreateConstantStatement {
	if r == nil {
		return nil
	}
	return &CreateConstantStatement{baseNode[*googlesql.ASTCreateConstantStatement]{raw: r}}
}
func (n *CreateConstantStatement) isStatement()        {}
func (n *CreateConstantStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateConstantStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateConstantStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateConstantStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}
func (n *CreateConstantStatement) Expr() ExpressionNode { return wrapExpr(must(n.raw.Expr())) }

// CreateSnapshotStatement wraps *googlesql.ASTCreateSnapshotStatement.
type CreateSnapshotStatement struct {
	baseNode[*googlesql.ASTCreateSnapshotStatement]
}

func newCreateSnapshotStatement(r *googlesql.ASTCreateSnapshotStatement) *CreateSnapshotStatement {
	if r == nil {
		return nil
	}
	return &CreateSnapshotStatement{baseNode[*googlesql.ASTCreateSnapshotStatement]{raw: r}}
}
func (n *CreateSnapshotStatement) isStatement()        {}
func (n *CreateSnapshotStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateSnapshotStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateSnapshotStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateSnapshotStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateSnapshotStatement) CloneDataSource() *CloneDataSource {
	return newCloneDataSource(must(n.raw.CloneDataSource()))
}

func (n *CreateSnapshotStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreateEntityStatement wraps *googlesql.ASTCreateEntityStatement.
type CreateEntityStatement struct {
	baseNode[*googlesql.ASTCreateEntityStatement]
}

func newCreateEntityStatement(r *googlesql.ASTCreateEntityStatement) *CreateEntityStatement {
	if r == nil {
		return nil
	}
	return &CreateEntityStatement{baseNode[*googlesql.ASTCreateEntityStatement]{raw: r}}
}
func (n *CreateEntityStatement) isStatement()          {}
func (n *CreateEntityStatement) IsOrReplace() bool     { return must(n.raw.IsOrReplace()) }
func (n *CreateEntityStatement) IsIfNotExists() bool   { return must(n.raw.IsIfNotExists()) }
func (n *CreateEntityStatement) Scope() Scope          { return must(n.raw.Scope()) }
func (n *CreateEntityStatement) Type() *Identifier     { return newIdentifier(must(n.raw.Type())) }
func (n *CreateEntityStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *CreateEntityStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *CreateEntityStatement) JSONBody() *JSONLiteral {
	return newJSONLiteral(must(n.raw.JsonBody()))
}

func (n *CreateEntityStatement) TextBody() *StringLiteral {
	return newStringLiteral(must(n.raw.TextBody()))
}

// CreateApproxViewStatement wraps *googlesql.ASTCreateApproxViewStatement.
type CreateApproxViewStatement struct {
	baseNode[*googlesql.ASTCreateApproxViewStatement]
}

func newCreateApproxViewStatement(r *googlesql.ASTCreateApproxViewStatement) *CreateApproxViewStatement {
	if r == nil {
		return nil
	}
	return &CreateApproxViewStatement{baseNode[*googlesql.ASTCreateApproxViewStatement]{raw: r}}
}
func (n *CreateApproxViewStatement) isStatement()        {}
func (n *CreateApproxViewStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateApproxViewStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateApproxViewStatement) Scope() Scope        { return must(n.raw.Scope()) }

// CreateExternalSchemaStatement wraps *googlesql.ASTCreateExternalSchemaStatement.
type CreateExternalSchemaStatement struct {
	baseNode[*googlesql.ASTCreateExternalSchemaStatement]
}

func newCreateExternalSchemaStatement(r *googlesql.ASTCreateExternalSchemaStatement) *CreateExternalSchemaStatement {
	if r == nil {
		return nil
	}
	return &CreateExternalSchemaStatement{baseNode[*googlesql.ASTCreateExternalSchemaStatement]{raw: r}}
}
func (n *CreateExternalSchemaStatement) isStatement()        {}
func (n *CreateExternalSchemaStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreateExternalSchemaStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreateExternalSchemaStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreateExternalSchemaStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// CreateLocalityGroupStatement wraps *googlesql.ASTCreateLocalityGroupStatement.
type CreateLocalityGroupStatement struct {
	baseNode[*googlesql.ASTCreateLocalityGroupStatement]
}

func newCreateLocalityGroupStatement(r *googlesql.ASTCreateLocalityGroupStatement) *CreateLocalityGroupStatement {
	if r == nil {
		return nil
	}
	return &CreateLocalityGroupStatement{baseNode[*googlesql.ASTCreateLocalityGroupStatement]{raw: r}}
}
func (n *CreateLocalityGroupStatement) isStatement() {}
func (n *CreateLocalityGroupStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreateLocalityGroupStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// CreatePrivilegeRestrictionStatement wraps *googlesql.ASTCreatePrivilegeRestrictionStatement.
type CreatePrivilegeRestrictionStatement struct {
	baseNode[*googlesql.ASTCreatePrivilegeRestrictionStatement]
}

func newCreatePrivilegeRestrictionStatement(r *googlesql.ASTCreatePrivilegeRestrictionStatement) *CreatePrivilegeRestrictionStatement {
	if r == nil {
		return nil
	}
	return &CreatePrivilegeRestrictionStatement{baseNode[*googlesql.ASTCreatePrivilegeRestrictionStatement]{raw: r}}
}
func (n *CreatePrivilegeRestrictionStatement) isStatement()      {}
func (n *CreatePrivilegeRestrictionStatement) IsOrReplace() bool { return must(n.raw.IsOrReplace()) }
func (n *CreatePrivilegeRestrictionStatement) IsIfNotExists() bool {
	return must(n.raw.IsIfNotExists())
}
func (n *CreatePrivilegeRestrictionStatement) Scope() Scope { return must(n.raw.Scope()) }
func (n *CreatePrivilegeRestrictionStatement) Privileges() Node {
	return Wrap(must(n.raw.Privileges()))
}

func (n *CreatePrivilegeRestrictionStatement) ObjectType() *Identifier {
	return newIdentifier(must(n.raw.ObjectType()))
}

func (n *CreatePrivilegeRestrictionStatement) NamePath() *PathExpression {
	return newPathExpression(must(n.raw.NamePath()))
}

func (n *CreatePrivilegeRestrictionStatement) RestrictTo() Node {
	return Wrap(must(n.raw.RestrictTo()))
}

// CreatePropertyGraphStatement wraps *googlesql.ASTCreatePropertyGraphStatement.
type CreatePropertyGraphStatement struct {
	baseNode[*googlesql.ASTCreatePropertyGraphStatement]
}

func newCreatePropertyGraphStatement(r *googlesql.ASTCreatePropertyGraphStatement) *CreatePropertyGraphStatement {
	if r == nil {
		return nil
	}
	return &CreatePropertyGraphStatement{baseNode[*googlesql.ASTCreatePropertyGraphStatement]{raw: r}}
}
func (n *CreatePropertyGraphStatement) isStatement()        {}
func (n *CreatePropertyGraphStatement) IsOrReplace() bool   { return must(n.raw.IsOrReplace()) }
func (n *CreatePropertyGraphStatement) IsIfNotExists() bool { return must(n.raw.IsIfNotExists()) }
func (n *CreatePropertyGraphStatement) Scope() Scope        { return must(n.raw.Scope()) }
func (n *CreatePropertyGraphStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *CreatePropertyGraphStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// AlterApproxViewStatement wraps *googlesql.ASTAlterApproxViewStatement.
type AlterApproxViewStatement struct {
	baseNode[*googlesql.ASTAlterApproxViewStatement]
}

func newAlterApproxViewStatement(r *googlesql.ASTAlterApproxViewStatement) *AlterApproxViewStatement {
	if r == nil {
		return nil
	}
	return &AlterApproxViewStatement{baseNode[*googlesql.ASTAlterApproxViewStatement]{raw: r}}
}
func (n *AlterApproxViewStatement) isStatement() {}
func (n *AlterApproxViewStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterApproxViewStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterConnectionStatement wraps *googlesql.ASTAlterConnectionStatement.
type AlterConnectionStatement struct {
	baseNode[*googlesql.ASTAlterConnectionStatement]
}

func newAlterConnectionStatement(r *googlesql.ASTAlterConnectionStatement) *AlterConnectionStatement {
	if r == nil {
		return nil
	}
	return &AlterConnectionStatement{baseNode[*googlesql.ASTAlterConnectionStatement]{raw: r}}
}
func (n *AlterConnectionStatement) isStatement() {}
func (n *AlterConnectionStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterConnectionStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterExternalSchemaStatement wraps *googlesql.ASTAlterExternalSchemaStatement.
type AlterExternalSchemaStatement struct {
	baseNode[*googlesql.ASTAlterExternalSchemaStatement]
}

func newAlterExternalSchemaStatement(r *googlesql.ASTAlterExternalSchemaStatement) *AlterExternalSchemaStatement {
	if r == nil {
		return nil
	}
	return &AlterExternalSchemaStatement{baseNode[*googlesql.ASTAlterExternalSchemaStatement]{raw: r}}
}
func (n *AlterExternalSchemaStatement) isStatement() {}
func (n *AlterExternalSchemaStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterExternalSchemaStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterIndexStatement wraps *googlesql.ASTAlterIndexStatement.
type AlterIndexStatement struct {
	baseNode[*googlesql.ASTAlterIndexStatement]
}

func newAlterIndexStatement(r *googlesql.ASTAlterIndexStatement) *AlterIndexStatement {
	if r == nil {
		return nil
	}
	return &AlterIndexStatement{baseNode[*googlesql.ASTAlterIndexStatement]{raw: r}}
}
func (n *AlterIndexStatement) isStatement() {}
func (n *AlterIndexStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterIndexStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

func (n *AlterIndexStatement) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}
func (n *AlterIndexStatement) IndexType() IndexType { return must(n.raw.IndexType()) }

// AlterModelStatement wraps *googlesql.ASTAlterModelStatement.
type AlterModelStatement struct {
	baseNode[*googlesql.ASTAlterModelStatement]
}

func newAlterModelStatement(r *googlesql.ASTAlterModelStatement) *AlterModelStatement {
	if r == nil {
		return nil
	}
	return &AlterModelStatement{baseNode[*googlesql.ASTAlterModelStatement]{raw: r}}
}
func (n *AlterModelStatement) isStatement() {}
func (n *AlterModelStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterModelStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// AlterSequenceStatement wraps *googlesql.ASTAlterSequenceStatement.
type AlterSequenceStatement struct {
	baseNode[*googlesql.ASTAlterSequenceStatement]
}

func newAlterSequenceStatement(r *googlesql.ASTAlterSequenceStatement) *AlterSequenceStatement {
	if r == nil {
		return nil
	}
	return &AlterSequenceStatement{baseNode[*googlesql.ASTAlterSequenceStatement]{raw: r}}
}
func (n *AlterSequenceStatement) isStatement() {}
func (n *AlterSequenceStatement) GetDdlTarget() *PathExpression {
	return newPathExpression(must(n.raw.GetDdlTarget()))
}

func (n *AlterSequenceStatement) ActionList() *AlterActionList {
	return newAlterActionList(must(n.raw.ActionList()))
}

// DropVectorIndexStatement wraps *googlesql.ASTDropVectorIndexStatement.
type DropVectorIndexStatement struct {
	baseNode[*googlesql.ASTDropVectorIndexStatement]
}

func newDropVectorIndexStatement(r *googlesql.ASTDropVectorIndexStatement) *DropVectorIndexStatement {
	if r == nil {
		return nil
	}
	return &DropVectorIndexStatement{baseNode[*googlesql.ASTDropVectorIndexStatement]{raw: r}}
}
func (n *DropVectorIndexStatement) isStatement()     {}
func (n *DropVectorIndexStatement) IsIfExists() bool { return must(n.raw.IsIfExists()) }
func (n *DropVectorIndexStatement) Name() *PathExpression {
	return newPathExpression(must(n.raw.Name()))
}

func (n *DropVectorIndexStatement) TableName() *PathExpression {
	return newPathExpression(must(n.raw.TableName()))
}

// UndropStatement wraps *googlesql.ASTUndropStatement.
type UndropStatement struct {
	baseNode[*googlesql.ASTUndropStatement]
}

func newUndropStatement(r *googlesql.ASTUndropStatement) *UndropStatement {
	if r == nil {
		return nil
	}
	return &UndropStatement{baseNode[*googlesql.ASTUndropStatement]{raw: r}}
}
func (n *UndropStatement) isStatement()                       {}
func (n *UndropStatement) IsIfNotExists() bool                { return must(n.raw.IsIfNotExists()) }
func (n *UndropStatement) Name() *PathExpression              { return newPathExpression(must(n.raw.Name())) }
func (n *UndropStatement) SchemaObjectKind() SchemaObjectKind { return must(n.raw.SchemaObjectKind()) }
func (n *UndropStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *UndropStatement) ForSystemTime() *ForSystemTime {
	return newForSystemTime(must(n.raw.ForSystemTime()))
}

// GrantStatement wraps *googlesql.ASTGrantStatement.
type GrantStatement struct {
	baseNode[*googlesql.ASTGrantStatement]
}

func newGrantStatement(r *googlesql.ASTGrantStatement) *GrantStatement {
	if r == nil {
		return nil
	}
	return &GrantStatement{baseNode[*googlesql.ASTGrantStatement]{raw: r}}
}
func (n *GrantStatement) isStatement()            {}
func (n *GrantStatement) Privileges() *Privileges { return newPrivileges(must(n.raw.Privileges())) }
func (n *GrantStatement) TargetType() *Identifier { return newIdentifier(must(n.raw.TargetType())) }
func (n *GrantStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}
func (n *GrantStatement) GranteeList() *GranteeList { return newGranteeList(must(n.raw.GranteeList())) }

// RevokeStatement wraps *googlesql.ASTRevokeStatement.
type RevokeStatement struct {
	baseNode[*googlesql.ASTRevokeStatement]
}

func newRevokeStatement(r *googlesql.ASTRevokeStatement) *RevokeStatement {
	if r == nil {
		return nil
	}
	return &RevokeStatement{baseNode[*googlesql.ASTRevokeStatement]{raw: r}}
}
func (n *RevokeStatement) isStatement()            {}
func (n *RevokeStatement) Privileges() *Privileges { return newPrivileges(must(n.raw.Privileges())) }
func (n *RevokeStatement) TargetType() *Identifier { return newIdentifier(must(n.raw.TargetType())) }
func (n *RevokeStatement) TargetPath() *PathExpression {
	return newPathExpression(must(n.raw.TargetPath()))
}

func (n *RevokeStatement) GranteeList() *GranteeList {
	return newGranteeList(must(n.raw.GranteeList()))
}

// ExportDataStatement wraps *googlesql.ASTExportDataStatement.
type ExportDataStatement struct {
	baseNode[*googlesql.ASTExportDataStatement]
}

func newExportDataStatement(r *googlesql.ASTExportDataStatement) *ExportDataStatement {
	if r == nil {
		return nil
	}
	return &ExportDataStatement{baseNode[*googlesql.ASTExportDataStatement]{raw: r}}
}
func (n *ExportDataStatement) isStatement()  {}
func (n *ExportDataStatement) Query() *Query { return newQuery(must(n.raw.Query())) }
func (n *ExportDataStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *ExportDataStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// ExportModelStatement wraps *googlesql.ASTExportModelStatement.
type ExportModelStatement struct {
	baseNode[*googlesql.ASTExportModelStatement]
}

func newExportModelStatement(r *googlesql.ASTExportModelStatement) *ExportModelStatement {
	if r == nil {
		return nil
	}
	return &ExportModelStatement{baseNode[*googlesql.ASTExportModelStatement]{raw: r}}
}
func (n *ExportModelStatement) isStatement() {}
func (n *ExportModelStatement) ModelNamePath() *PathExpression {
	return newPathExpression(must(n.raw.ModelNamePath()))
}

func (n *ExportModelStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *ExportModelStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

// ExportMetadataStatement wraps *googlesql.ASTExportMetadataStatement.
type ExportMetadataStatement struct {
	baseNode[*googlesql.ASTExportMetadataStatement]
}

func newExportMetadataStatement(r *googlesql.ASTExportMetadataStatement) *ExportMetadataStatement {
	if r == nil {
		return nil
	}
	return &ExportMetadataStatement{baseNode[*googlesql.ASTExportMetadataStatement]{raw: r}}
}
func (n *ExportMetadataStatement) isStatement() {}
func (n *ExportMetadataStatement) NamePath() *PathExpression {
	return newPathExpression(must(n.raw.NamePath()))
}

func (n *ExportMetadataStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

func (n *ExportMetadataStatement) WithConnectionClause() *WithConnectionClause {
	return newWithConnectionClause(must(n.raw.WithConnectionClause()))
}

func (n *ExportMetadataStatement) SchemaObjectKind() SchemaObjectKind {
	return must(n.raw.SchemaObjectKind())
}

// AnalyzeStatement wraps *googlesql.ASTAnalyzeStatement.
type AnalyzeStatement struct {
	baseNode[*googlesql.ASTAnalyzeStatement]
}

func newAnalyzeStatement(r *googlesql.ASTAnalyzeStatement) *AnalyzeStatement {
	if r == nil {
		return nil
	}
	return &AnalyzeStatement{baseNode[*googlesql.ASTAnalyzeStatement]{raw: r}}
}
func (n *AnalyzeStatement) isStatement() {}
func (n *AnalyzeStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// DescribeStatement wraps *googlesql.ASTDescribeStatement.
type DescribeStatement struct {
	baseNode[*googlesql.ASTDescribeStatement]
}

func newDescribeStatement(r *googlesql.ASTDescribeStatement) *DescribeStatement {
	if r == nil {
		return nil
	}
	return &DescribeStatement{baseNode[*googlesql.ASTDescribeStatement]{raw: r}}
}
func (n *DescribeStatement) isStatement()          {}
func (n *DescribeStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *DescribeStatement) OptionalIdentifier() *Identifier {
	return newIdentifier(must(n.raw.OptionalIdentifier()))
}

func (n *DescribeStatement) OptionalFromName() *PathExpression {
	return newPathExpression(must(n.raw.OptionalFromName()))
}

// ExplainStatement wraps *googlesql.ASTExplainStatement.
type ExplainStatement struct {
	baseNode[*googlesql.ASTExplainStatement]
}

func newExplainStatement(r *googlesql.ASTExplainStatement) *ExplainStatement {
	if r == nil {
		return nil
	}
	return &ExplainStatement{baseNode[*googlesql.ASTExplainStatement]{raw: r}}
}
func (n *ExplainStatement) isStatement()             {}
func (n *ExplainStatement) Statement() StatementNode { return wrapStmt(must(n.raw.Statement())) }

// ShowStatement wraps *googlesql.ASTShowStatement.
type ShowStatement struct {
	baseNode[*googlesql.ASTShowStatement]
}

func newShowStatement(r *googlesql.ASTShowStatement) *ShowStatement {
	if r == nil {
		return nil
	}
	return &ShowStatement{baseNode[*googlesql.ASTShowStatement]{raw: r}}
}
func (n *ShowStatement) isStatement()            {}
func (n *ShowStatement) Identifier() *Identifier { return newIdentifier(must(n.raw.Identifier())) }
func (n *ShowStatement) OptionalName() *PathExpression {
	return newPathExpression(must(n.raw.OptionalName()))
}

func (n *ShowStatement) OptionalLikeString() *StringLiteral {
	return newStringLiteral(must(n.raw.OptionalLikeString()))
}

// ImportStatement wraps *googlesql.ASTImportStatement.
type ImportStatement struct {
	baseNode[*googlesql.ASTImportStatement]
}

func newImportStatement(r *googlesql.ASTImportStatement) *ImportStatement {
	if r == nil {
		return nil
	}
	return &ImportStatement{baseNode[*googlesql.ASTImportStatement]{raw: r}}
}
func (n *ImportStatement) isStatement()          {}
func (n *ImportStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *ImportStatement) StringValue() *StringLiteral {
	return newStringLiteral(must(n.raw.StringValue()))
}
func (n *ImportStatement) Alias() *Alias         { return newAlias(must(n.raw.Alias())) }
func (n *ImportStatement) IntoAlias() *IntoAlias { return newIntoAlias(must(n.raw.IntoAlias())) }
func (n *ImportStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}
func (n *ImportStatement) ImportKind() ImportKind { return must(n.raw.ImportKind()) }

// ModuleStatement wraps *googlesql.ASTModuleStatement.
type ModuleStatement struct {
	baseNode[*googlesql.ASTModuleStatement]
}

func newModuleStatement(r *googlesql.ASTModuleStatement) *ModuleStatement {
	if r == nil {
		return nil
	}
	return &ModuleStatement{baseNode[*googlesql.ASTModuleStatement]{raw: r}}
}
func (n *ModuleStatement) isStatement()          {}
func (n *ModuleStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *ModuleStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// DefineMacroStatement wraps *googlesql.ASTDefineMacroStatement.
type DefineMacroStatement struct {
	baseNode[*googlesql.ASTDefineMacroStatement]
}

func newDefineMacroStatement(r *googlesql.ASTDefineMacroStatement) *DefineMacroStatement {
	if r == nil {
		return nil
	}
	return &DefineMacroStatement{baseNode[*googlesql.ASTDefineMacroStatement]{raw: r}}
}
func (n *DefineMacroStatement) isStatement()      {}
func (n *DefineMacroStatement) Name() *Identifier { return newIdentifier(must(n.raw.Name())) }

// DefineTableStatement wraps *googlesql.ASTDefineTableStatement.
type DefineTableStatement struct {
	baseNode[*googlesql.ASTDefineTableStatement]
}

func newDefineTableStatement(r *googlesql.ASTDefineTableStatement) *DefineTableStatement {
	if r == nil {
		return nil
	}
	return &DefineTableStatement{baseNode[*googlesql.ASTDefineTableStatement]{raw: r}}
}
func (n *DefineTableStatement) isStatement()          {}
func (n *DefineTableStatement) Name() *PathExpression { return newPathExpression(must(n.raw.Name())) }
func (n *DefineTableStatement) OptionsList() *OptionsList {
	return newOptionsList(must(n.raw.OptionsList()))
}

// RunStatement wraps *googlesql.ASTRunStatement.
type RunStatement struct {
	baseNode[*googlesql.ASTRunStatement]
}

func newRunStatement(r *googlesql.ASTRunStatement) *RunStatement {
	if r == nil {
		return nil
	}
	return &RunStatement{baseNode[*googlesql.ASTRunStatement]{raw: r}}
}
func (n *RunStatement) isStatement() {}
func (n *RunStatement) TargetPathExpression() *PathExpression {
	return newPathExpression(must(n.raw.TargetPathExpression()))
}

func (n *RunStatement) TargetStringLiteral() *StringLiteral {
	return newStringLiteral(must(n.raw.TargetStringLiteral()))
}

// SetTransactionStatement wraps *googlesql.ASTSetTransactionStatement.
type SetTransactionStatement struct {
	baseNode[*googlesql.ASTSetTransactionStatement]
}

func newSetTransactionStatement(r *googlesql.ASTSetTransactionStatement) *SetTransactionStatement {
	if r == nil {
		return nil
	}
	return &SetTransactionStatement{baseNode[*googlesql.ASTSetTransactionStatement]{raw: r}}
}
func (n *SetTransactionStatement) isStatement() {}

// AbortBatchStatement wraps *googlesql.ASTAbortBatchStatement.
type AbortBatchStatement struct {
	baseNode[*googlesql.ASTAbortBatchStatement]
}

func newAbortBatchStatement(r *googlesql.ASTAbortBatchStatement) *AbortBatchStatement {
	if r == nil {
		return nil
	}
	return &AbortBatchStatement{baseNode[*googlesql.ASTAbortBatchStatement]{raw: r}}
}
func (n *AbortBatchStatement) isStatement() {}

// RunBatchStatement wraps *googlesql.ASTRunBatchStatement.
type RunBatchStatement struct {
	baseNode[*googlesql.ASTRunBatchStatement]
}

func newRunBatchStatement(r *googlesql.ASTRunBatchStatement) *RunBatchStatement {
	if r == nil {
		return nil
	}
	return &RunBatchStatement{baseNode[*googlesql.ASTRunBatchStatement]{raw: r}}
}
func (n *RunBatchStatement) isStatement() {}

// StartBatchStatement wraps *googlesql.ASTStartBatchStatement.
type StartBatchStatement struct {
	baseNode[*googlesql.ASTStartBatchStatement]
}

func newStartBatchStatement(r *googlesql.ASTStartBatchStatement) *StartBatchStatement {
	if r == nil {
		return nil
	}
	return &StartBatchStatement{baseNode[*googlesql.ASTStartBatchStatement]{raw: r}}
}
func (n *StartBatchStatement) isStatement()           {}
func (n *StartBatchStatement) BatchType() *Identifier { return newIdentifier(must(n.raw.BatchType())) }

// CloneDataStatement wraps *googlesql.ASTCloneDataStatement.
type CloneDataStatement struct {
	baseNode[*googlesql.ASTCloneDataStatement]
}

func newCloneDataStatement(r *googlesql.ASTCloneDataStatement) *CloneDataStatement {
	if r == nil {
		return nil
	}
	return &CloneDataStatement{baseNode[*googlesql.ASTCloneDataStatement]{raw: r}}
}
func (n *CloneDataStatement) isStatement() {}

// AuxLoadDataStatement wraps *googlesql.ASTAuxLoadDataStatement.
type AuxLoadDataStatement struct {
	baseNode[*googlesql.ASTAuxLoadDataStatement]
}

func newAuxLoadDataStatement(r *googlesql.ASTAuxLoadDataStatement) *AuxLoadDataStatement {
	if r == nil {
		return nil
	}
	return &AuxLoadDataStatement{baseNode[*googlesql.ASTAuxLoadDataStatement]{raw: r}}
}
func (n *AuxLoadDataStatement) isStatement() {}
