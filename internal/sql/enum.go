package sql

import "github.com/goccy/go-googlesql"

// SkipTargetType enums.
type SkipTargetType = googlesql.ASTAfterMatchSkipClauseEnums_AfterMatchSkipTargetType

const (
	UnspecifiedSkipTarget SkipTargetType = googlesql.ASTAfterMatchSkipClauseEnums_AfterMatchSkipTargetTypeAfterMatchSkipTargetUnspecified
	PastLastRow           SkipTargetType = googlesql.ASTAfterMatchSkipClauseEnums_AfterMatchSkipTargetTypePastLastRow
	ToNextRow             SkipTargetType = googlesql.ASTAfterMatchSkipClauseEnums_AfterMatchSkipTargetTypeToNextRow
)

// IndexType enums.
type IndexType = googlesql.ASTAlterIndexStatementEnums_IndexType

const (
	DefaultIndex IndexType = googlesql.ASTAlterIndexStatementEnums_IndexTypeIndexDefault
	SearchIndex  IndexType = googlesql.ASTAlterIndexStatementEnums_IndexTypeIndexSearch
	IndexVector  IndexType = googlesql.ASTAlterIndexStatementEnums_IndexTypeIndexVector
)

// AnySomeAllOpType enums.
type AnySomeAllOpType = googlesql.ASTAnySomeAllOpEnums_Op

const (
	UninitializedOp AnySomeAllOpType = googlesql.ASTAnySomeAllOpEnums_OpKUninitialized
	AnyOp           AnySomeAllOpType = googlesql.ASTAnySomeAllOpEnums_OpKAny
	SomeOp          AnySomeAllOpType = googlesql.ASTAnySomeAllOpEnums_OpKSome
	AllOp           AnySomeAllOpType = googlesql.ASTAnySomeAllOpEnums_OpKAll
)

// InsertionMode enums.
type InsertionMode = googlesql.ASTAuxLoadDataStatementEnums_InsertionMode

const (
	NotSetInsertionMode InsertionMode = googlesql.ASTAuxLoadDataStatementEnums_InsertionModeNotSet
	Append              InsertionMode = googlesql.ASTAuxLoadDataStatementEnums_InsertionModeAppend
	Overwrite           InsertionMode = googlesql.ASTAuxLoadDataStatementEnums_InsertionModeOverwrite
)

// BinaryOp enums.
type BinaryOp = googlesql.ASTBinaryExpressionEnums_Op

const (
	NotSetBinaryOp BinaryOp = googlesql.ASTBinaryExpressionEnums_OpNotSet
	LikeOp         BinaryOp = googlesql.ASTBinaryExpressionEnums_OpLike
	IsOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpIs
	EqOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpEq
	NEOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpNe
	NE2Op          BinaryOp = googlesql.ASTBinaryExpressionEnums_OpNe2
	GTOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpGt
	LTOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpLt
	GEOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpGe
	LEOp           BinaryOp = googlesql.ASTBinaryExpressionEnums_OpLe
	BitwiseOrOp    BinaryOp = googlesql.ASTBinaryExpressionEnums_OpBitwiseOr
	BitwiseXorOp   BinaryOp = googlesql.ASTBinaryExpressionEnums_OpBitwiseXor
	BitwiseAndOp   BinaryOp = googlesql.ASTBinaryExpressionEnums_OpBitwiseAnd
	PlusBinaryOp   BinaryOp = googlesql.ASTBinaryExpressionEnums_OpPlus
	MinusBinaryOp  BinaryOp = googlesql.ASTBinaryExpressionEnums_OpMinus
	MultiplyOp     BinaryOp = googlesql.ASTBinaryExpressionEnums_OpMultiply
	DivideOp       BinaryOp = googlesql.ASTBinaryExpressionEnums_OpDivide
	ConcatOpOp     BinaryOp = googlesql.ASTBinaryExpressionEnums_OpConcatOp
	DistinctOp     BinaryOp = googlesql.ASTBinaryExpressionEnums_OpDistinct
	IsSourceNodeOp BinaryOp = googlesql.ASTBinaryExpressionEnums_OpIsSourceNode
	IsDestNodeOp   BinaryOp = googlesql.ASTBinaryExpressionEnums_OpIsDestNode
)

// LhsOp enums.
type LhsOp = googlesql.ASTBracedConstructorLhsEnums_Operation

const (
	UpdateSingle           LhsOp = googlesql.ASTBracedConstructorLhsEnums_OperationUpdateSingle
	UpdateManyLhsOp        LhsOp = googlesql.ASTBracedConstructorLhsEnums_OperationUpdateMany
	UpdateSingleNoCreation LhsOp = googlesql.ASTBracedConstructorLhsEnums_OperationUpdateSingleNoCreation
)

// BreakContinueKeyword enums.
type BreakContinueKeyword = googlesql.ASTBreakContinueStatementEnums_BreakContinueKeyword

const (
	BreakKeyword    BreakContinueKeyword = googlesql.ASTBreakContinueStatementEnums_BreakContinueKeywordBreak
	LeaveKeyword    BreakContinueKeyword = googlesql.ASTBreakContinueStatementEnums_BreakContinueKeywordLeave
	ContinueKeyword BreakContinueKeyword = googlesql.ASTBreakContinueStatementEnums_BreakContinueKeywordContinue
	IterateKeyword  BreakContinueKeyword = googlesql.ASTBreakContinueStatementEnums_BreakContinueKeywordIterate
)

// RelativePosition enums.
type RelativePosition = googlesql.ASTColumnPositionEnums_RelativePositionType

const (
	Preceding RelativePosition = googlesql.ASTColumnPositionEnums_RelativePositionTypePreceding
	Following RelativePosition = googlesql.ASTColumnPositionEnums_RelativePositionTypeFollowing
)

// DeterminismLevel enums.
type DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevel

const (
	UnspecifiedDeterminism DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelDeterminismUnspecified
	Deterministic          DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelDeterministic
	NotDeterministic       DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelNotDeterministic
	Immutable              DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelImmutable
	Stable                 DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelStable
	Volatile               DeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelVolatile
)

// Scope enums.
type Scope = googlesql.ASTCreateStatementEnums_Scope

const (
	DefaultScope Scope = googlesql.ASTCreateStatementEnums_ScopeDefaultScope
	Private      Scope = googlesql.ASTCreateStatementEnums_ScopePrivate
	Public       Scope = googlesql.ASTCreateStatementEnums_ScopePublic
	Temporary    Scope = googlesql.ASTCreateStatementEnums_ScopeTemporary
)

// SqlSecurity enums.
type SqlSecurity = googlesql.ASTCreateStatementEnums_SqlSecurity

const (
	SQLSecurityUnspecifiedSecurity SqlSecurity = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityUnspecified
	SQLSecurityDefiner             SqlSecurity = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityDefiner
	SQLSecurityInvoker             SqlSecurity = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityInvoker
)

// DropMode enums.
type DropMode = googlesql.ASTDropStatementEnums_DropMode

const (
	UnspecifiedDropMode DropMode = googlesql.ASTDropStatementEnums_DropModeDropModeUnspecified
	RestrictDropMode    DropMode = googlesql.ASTDropStatementEnums_DropModeRestrict
	CascadeDropMode     DropMode = googlesql.ASTDropStatementEnums_DropModeCascade
)

// SubqueryModifier enums.
type SubqueryModifier = googlesql.ASTExpressionSubqueryEnums_Modifier

const (
	NoneModifier          SubqueryModifier = googlesql.ASTExpressionSubqueryEnums_ModifierNone
	Array                 SubqueryModifier = googlesql.ASTExpressionSubqueryEnums_ModifierArray
	Exists                SubqueryModifier = googlesql.ASTExpressionSubqueryEnums_ModifierExists
	ValueSubqueryModifier SubqueryModifier = googlesql.ASTExpressionSubqueryEnums_ModifierValue
)

// FilterType enums.
type FilterType = googlesql.ASTFilterFieldsArgEnums_FilterType

const (
	NotSetFilterType  FilterType = googlesql.ASTFilterFieldsArgEnums_FilterTypeNotSet
	IncludeFilterType FilterType = googlesql.ASTFilterFieldsArgEnums_FilterTypeInclude
	ExcludeFilterType FilterType = googlesql.ASTFilterFieldsArgEnums_FilterTypeExclude
)

// ForeignKeyAction enums.
type ForeignKeyAction = googlesql.ASTForeignKeyActionsEnums_Action

const (
	NoAction       ForeignKeyAction = googlesql.ASTForeignKeyActionsEnums_ActionNoAction
	RestrictAction ForeignKeyAction = googlesql.ASTForeignKeyActionsEnums_ActionRestrict
	CascadeAction  ForeignKeyAction = googlesql.ASTForeignKeyActionsEnums_ActionCascade
	SetNullAction  ForeignKeyAction = googlesql.ASTForeignKeyActionsEnums_ActionSetNull
)

// ForeignKeyMatch enums.
type ForeignKeyMatch = googlesql.ASTForeignKeyReferenceEnums_Match

const (
	SimpleMatch      ForeignKeyMatch = googlesql.ASTForeignKeyReferenceEnums_MatchSimple
	FullMatch        ForeignKeyMatch = googlesql.ASTForeignKeyReferenceEnums_MatchFull
	NotDistinctMatch ForeignKeyMatch = googlesql.ASTForeignKeyReferenceEnums_MatchNotDistinct
)

// NullHandlingModifier enums.
type NullHandlingModifier = googlesql.ASTFunctionCallEnums_NullHandlingModifier

const (
	DefaultNullHandling NullHandlingModifier = googlesql.ASTFunctionCallEnums_NullHandlingModifierDefaultNullHandling
	IgnoreNulls         NullHandlingModifier = googlesql.ASTFunctionCallEnums_NullHandlingModifierIgnoreNulls
	RespectNulls        NullHandlingModifier = googlesql.ASTFunctionCallEnums_NullHandlingModifierRespectNulls
)

// ParameterMode enums.
type ParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterMode

const (
	NotSetParameterMode ParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeNotSet
	InParameterMode     ParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeIn
	Out                 ParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeOut
	InOut               ParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeInout
)

// GeneratedMode enums.
type GeneratedMode = googlesql.ASTGeneratedColumnInfoEnums_GeneratedMode

const (
	Always        GeneratedMode = googlesql.ASTGeneratedColumnInfoEnums_GeneratedModeAlways
	ByDefaultMode GeneratedMode = googlesql.ASTGeneratedColumnInfoEnums_GeneratedModeByDefault
)

// StoredMode enums.
type StoredMode = googlesql.ASTGeneratedColumnInfoEnums_StoredMode

const (
	NonStored      StoredMode = googlesql.ASTGeneratedColumnInfoEnums_StoredModeNonStored
	Stored         StoredMode = googlesql.ASTGeneratedColumnInfoEnums_StoredModeStored
	StoredVolatile StoredMode = googlesql.ASTGeneratedColumnInfoEnums_StoredModeStoredVolatile
)

// EdgeOrientation enums.
type EdgeOrientation = googlesql.ASTGraphEdgePatternEnums_EdgeOrientation

const (
	NotSetEdgeOrientation EdgeOrientation = googlesql.ASTGraphEdgePatternEnums_EdgeOrientationEdgeOrientationNotSet
	AnyOrientation        EdgeOrientation = googlesql.ASTGraphEdgePatternEnums_EdgeOrientationAny
	LeftOrientation       EdgeOrientation = googlesql.ASTGraphEdgePatternEnums_EdgeOrientationLeft
	RightOrientation      EdgeOrientation = googlesql.ASTGraphEdgePatternEnums_EdgeOrientationRight
)

// GraphLabelOp enums.
type GraphLabelOp = googlesql.ASTGraphLabelOperationEnums_OperationType

const (
	UnspecifiedGraphLabelOperationType GraphLabelOp = googlesql.ASTGraphLabelOperationEnums_OperationTypeOperationTypeUnspecified
	NotGraphLabelOp                    GraphLabelOp = googlesql.ASTGraphLabelOperationEnums_OperationTypeNot
	AndOp                              GraphLabelOp = googlesql.ASTGraphLabelOperationEnums_OperationTypeAnd
	OrOp                               GraphLabelOp = googlesql.ASTGraphLabelOperationEnums_OperationTypeOr
)

// NodeReferenceType enums.
type NodeReferenceType = googlesql.ASTGraphNodeTableReferenceEnums_NodeReferenceType

const (
	UnspecifiedRefType NodeReferenceType = googlesql.ASTGraphNodeTableReferenceEnums_NodeReferenceTypeNodeReferenceTypeUnspecified
	Source             NodeReferenceType = googlesql.ASTGraphNodeTableReferenceEnums_NodeReferenceTypeSource
	Destination        NodeReferenceType = googlesql.ASTGraphNodeTableReferenceEnums_NodeReferenceTypeDestination
)

// PathMode enums.
type PathMode = googlesql.ASTGraphPathModeEnums_PathMode

const (
	UnspecifiedPathMode PathMode = googlesql.ASTGraphPathModeEnums_PathModePathModeUnspecified
	Walk                PathMode = googlesql.ASTGraphPathModeEnums_PathModeWalk
	Trail               PathMode = googlesql.ASTGraphPathModeEnums_PathModeTrail
	SimplePathMode      PathMode = googlesql.ASTGraphPathModeEnums_PathModeSimple
	Acyclic             PathMode = googlesql.ASTGraphPathModeEnums_PathModeAcyclic
)

// PathSearchPrefix enums.
type PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixType

const (
	UnspecifiedPrefix PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypePathSearchPrefixTypeUnspecified
	AnyPrefix         PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeAny
	Shortest          PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeShortest
	AllPrefix         PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeAll
	AllShortest       PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeAllShortest
	Cheapest          PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeCheapest
	AllCheapest       PathSearchPrefix = googlesql.ASTGraphPathSearchPrefixEnums_PathSearchPrefixTypeAllCheapest
)

// HavingModifierType enums.
type HavingModifierType = googlesql.ASTHavingModifierEnums_ModifierKind

const (
	NotSetModifier HavingModifierType = googlesql.ASTHavingModifierEnums_ModifierKindNotSet
	Min            HavingModifierType = googlesql.ASTHavingModifierEnums_ModifierKindMin
	Max            HavingModifierType = googlesql.ASTHavingModifierEnums_ModifierKindMax
)

// ImportKind enums.
type ImportKind = googlesql.ASTImportStatementEnums_ImportKind

const (
	Module ImportKind = googlesql.ASTImportStatementEnums_ImportKindModule
	Proto  ImportKind = googlesql.ASTImportStatementEnums_ImportKindProto
)

// InsertMode enums.
type InsertMode = googlesql.ASTInsertStatementEnums_InsertMode

const (
	DefaultMode      InsertMode = googlesql.ASTInsertStatementEnums_InsertModeDefaultMode
	Replace          InsertMode = googlesql.ASTInsertStatementEnums_InsertModeReplace
	UpdateInsertMode InsertMode = googlesql.ASTInsertStatementEnums_InsertModeUpdate
	Ignore           InsertMode = googlesql.ASTInsertStatementEnums_InsertModeIgnore
)

// InsertParseProgress enums.
type InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgress

const (
	Initial                   InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgressKInitial
	SeenOrIgnoreReplaceUpdate InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgressKSeenOrIgnoreReplaceUpdate
	SeenTargetPath            InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgressKSeenTargetPath
	SeenColumnList            InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgressKSeenColumnList
	SeenValuesList            InsertParseProgress = googlesql.ASTInsertStatementEnums_ParseProgressKSeenValuesList
)

// JoinHint enums.
type JoinHint = googlesql.ASTJoinEnums_JoinHint

const (
	NoJoinHint JoinHint = googlesql.ASTJoinEnums_JoinHintNoJoinHint
	Hash       JoinHint = googlesql.ASTJoinEnums_JoinHintHash
	Lookup     JoinHint = googlesql.ASTJoinEnums_JoinHintLookup
)

// JoinType enums.
type JoinType = googlesql.ASTJoinEnums_JoinType

const (
	DefaultJoinType JoinType = googlesql.ASTJoinEnums_JoinTypeDefaultJoinType
	Comma           JoinType = googlesql.ASTJoinEnums_JoinTypeComma
	Cross           JoinType = googlesql.ASTJoinEnums_JoinTypeCross
	FullJoin        JoinType = googlesql.ASTJoinEnums_JoinTypeFull
	InnerJoin       JoinType = googlesql.ASTJoinEnums_JoinTypeInner
	LeftJoin        JoinType = googlesql.ASTJoinEnums_JoinTypeLeft
	RightJoin       JoinType = googlesql.ASTJoinEnums_JoinTypeRight
)

// LockStrength enums.
type LockStrength = googlesql.ASTLockModeEnums_LockStrengthSpec

const (
	NotSetLockStrength LockStrength = googlesql.ASTLockModeEnums_LockStrengthSpecNotSet
	UpdateLockStrength LockStrength = googlesql.ASTLockModeEnums_LockStrengthSpecUpdate
)

// MergeActionType enums.
type MergeActionType = googlesql.ASTMergeActionEnums_ActionType

const (
	NotSetMergeAction MergeActionType = googlesql.ASTMergeActionEnums_ActionTypeNotSet
	InsertAction      MergeActionType = googlesql.ASTMergeActionEnums_ActionTypeInsert
	UpdateMergeAction MergeActionType = googlesql.ASTMergeActionEnums_ActionTypeUpdate
	DeleteAction      MergeActionType = googlesql.ASTMergeActionEnums_ActionTypeDelete
)

// MergeMatchType enums.
type MergeMatchType = googlesql.ASTMergeWhenClauseEnums_MatchType

const (
	NotSetMatchType    MergeMatchType = googlesql.ASTMergeWhenClauseEnums_MatchTypeNotSet
	Matched            MergeMatchType = googlesql.ASTMergeWhenClauseEnums_MatchTypeMatched
	NotMatchedBySource MergeMatchType = googlesql.ASTMergeWhenClauseEnums_MatchTypeNotMatchedBySource
	NotMatchedByTarget MergeMatchType = googlesql.ASTMergeWhenClauseEnums_MatchTypeNotMatchedByTarget
)

// ConflictAction enums.
type ConflictAction = googlesql.ASTOnConflictClauseEnums_ConflictAction

const (
	NotSetConflictAction  ConflictAction = googlesql.ASTOnConflictClauseEnums_ConflictActionNotSet
	NothingConflictAction ConflictAction = googlesql.ASTOnConflictClauseEnums_ConflictActionNothing
	UpdateConflictAction  ConflictAction = googlesql.ASTOnConflictClauseEnums_ConflictActionUpdate
)

// AssignmentOp enums.
type AssignmentOp = googlesql.ASTOptionsEntryEnums_AssignmentOp

const (
	NotSetAssignmentOp AssignmentOp = googlesql.ASTOptionsEntryEnums_AssignmentOpNotSet
	AssignOp           AssignmentOp = googlesql.ASTOptionsEntryEnums_AssignmentOpAssign
	AddAssignOp        AssignmentOp = googlesql.ASTOptionsEntryEnums_AssignmentOpAddAssign
	SubAssignOp        AssignmentOp = googlesql.ASTOptionsEntryEnums_AssignmentOpSubAssign
)

// OrderingSpec enums.
type OrderingSpec = googlesql.ASTOrderingExpressionEnums_OrderingSpec

const (
	NotSetSpec OrderingSpec = googlesql.ASTOrderingExpressionEnums_OrderingSpecNotSet
	Asc        OrderingSpec = googlesql.ASTOrderingExpressionEnums_OrderingSpecAsc
	Desc       OrderingSpec = googlesql.ASTOrderingExpressionEnums_OrderingSpecDesc
)

// RowPatternAnchorType enums.
type RowPatternAnchorType = googlesql.ASTRowPatternAnchorEnums_Anchor

const (
	UnspecifiedAnchor RowPatternAnchorType = googlesql.ASTRowPatternAnchorEnums_AnchorAnchorUnspecified
	Start             RowPatternAnchorType = googlesql.ASTRowPatternAnchorEnums_AnchorStart
	End               RowPatternAnchorType = googlesql.ASTRowPatternAnchorEnums_AnchorEnd
)

// RowPatternOp enums.
type RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationType

const (
	UnspecifiedRowPatternOp RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationTypeOperationTypeUnspecified
	ConcatOp                RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationTypeConcat
	AlternateOp             RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationTypeAlternate
	PermuteOp               RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationTypePermute
	ExcludeOp               RowPatternOp = googlesql.ASTRowPatternOperationEnums_OperationTypeExclude
)

// SampleSizeUnit enums.
type SampleSizeUnit = googlesql.ASTSampleSizeEnums_Unit

const (
	NotSetUnit     SampleSizeUnit = googlesql.ASTSampleSizeEnums_UnitNotSet
	RowsSampleSize SampleSizeUnit = googlesql.ASTSampleSizeEnums_UnitRows
	Percent        SampleSizeUnit = googlesql.ASTSampleSizeEnums_UnitPercent
)

// AsMode enums.
type AsMode = googlesql.ASTSelectAsEnums_AsMode

const (
	NotSetAsMode AsMode = googlesql.ASTSelectAsEnums_AsModeNotSet
	Struct       AsMode = googlesql.ASTSelectAsEnums_AsModeStruct
	ValueAsMode  AsMode = googlesql.ASTSelectAsEnums_AsModeValue
	TypeName     AsMode = googlesql.ASTSelectAsEnums_AsModeTypeName
)

// AllOrDistinct enums.
type AllOrDistinct = googlesql.ASTSetOperationEnums_AllOrDistinct

const (
	NotSetAllOrDistinct AllOrDistinct = googlesql.ASTSetOperationEnums_AllOrDistinctAllOrDistinctNotSet
	All                 AllOrDistinct = googlesql.ASTSetOperationEnums_AllOrDistinctAll
	Distinct            AllOrDistinct = googlesql.ASTSetOperationEnums_AllOrDistinctDistinct
)

// ColumnMatchMode enums.
type ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchMode

const (
	ByPosition      ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchModeByPosition
	Corresponding   ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchModeCorresponding
	CorrespondingBy ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchModeCorrespondingBy
	ByName          ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchModeByName
	ByNameOn        ColumnMatchMode = googlesql.ASTSetOperationEnums_ColumnMatchModeByNameOn
)

// ColumnPropagationMode enums.
type ColumnPropagationMode = googlesql.ASTSetOperationEnums_ColumnPropagationMode

const (
	Strict           ColumnPropagationMode = googlesql.ASTSetOperationEnums_ColumnPropagationModeStrict
	InnerPropagation ColumnPropagationMode = googlesql.ASTSetOperationEnums_ColumnPropagationModeInner
	LeftPropagation  ColumnPropagationMode = googlesql.ASTSetOperationEnums_ColumnPropagationModeLeft
	FullPropagation  ColumnPropagationMode = googlesql.ASTSetOperationEnums_ColumnPropagationModeFull
)

// SetOp enums.
type SetOp = googlesql.ASTSetOperationEnums_OperationType

const (
	NotSetSetOp SetOp = googlesql.ASTSetOperationEnums_OperationTypeNotSet
	UnionOp     SetOp = googlesql.ASTSetOperationEnums_OperationTypeUnion
	ExceptOp    SetOp = googlesql.ASTSetOperationEnums_OperationTypeExcept
	IntersectOp SetOp = googlesql.ASTSetOperationEnums_OperationTypeIntersect
)

// InterleaveClauseType enums.
type InterleaveClauseType = googlesql.ASTSpannerInterleaveClauseEnums_Type

const (
	NotSetInterleaveType InterleaveClauseType = googlesql.ASTSpannerInterleaveClauseEnums_TypeNotSet
	InInterleaveType     InterleaveClauseType = googlesql.ASTSpannerInterleaveClauseEnums_TypeIn
	InParent             InterleaveClauseType = googlesql.ASTSpannerInterleaveClauseEnums_TypeInParent
)

// QuantifierSymbol enums.
type QuantifierSymbol = googlesql.ASTSymbolQuantifierEnums_Symbol

const (
	UnspecifiedSymbol QuantifierSymbol = googlesql.ASTSymbolQuantifierEnums_SymbolSymbolUnspecified
	QuestionMark      QuantifierSymbol = googlesql.ASTSymbolQuantifierEnums_SymbolQuestionMark
	Plus              QuantifierSymbol = googlesql.ASTSymbolQuantifierEnums_SymbolPlus
	StarSymbol        QuantifierSymbol = googlesql.ASTSymbolQuantifierEnums_SymbolStar
)

// TemplatedTypeKind enums.
type TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKind

const (
	UninitializedTypeKind TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindUninitialized
	AnyType               TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyType
	AnyProto              TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyProto
	AnyEnum               TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyEnum
	AnyStruct             TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyStruct
	AnyArray              TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyArray
	AnyTable              TemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyTable
)

// ReadWriteMode enums.
type ReadWriteMode = googlesql.ASTTransactionReadWriteModeEnums_Mode

const (
	InvalidMode ReadWriteMode = googlesql.ASTTransactionReadWriteModeEnums_ModeInvalid
	ReadOnly    ReadWriteMode = googlesql.ASTTransactionReadWriteModeEnums_ModeReadOnly
	ReadWrite   ReadWriteMode = googlesql.ASTTransactionReadWriteModeEnums_ModeReadWrite
)

// UnaryOp enums.
type UnaryOp = googlesql.ASTUnaryExpressionEnums_Op

const (
	NotSetUnaryOp  UnaryOp = googlesql.ASTUnaryExpressionEnums_OpNotSet
	NotUnaryOp     UnaryOp = googlesql.ASTUnaryExpressionEnums_OpNot
	BitwiseNotOp   UnaryOp = googlesql.ASTUnaryExpressionEnums_OpBitwiseNot
	MinusUnaryOp   UnaryOp = googlesql.ASTUnaryExpressionEnums_OpMinus
	PlusUnaryOp    UnaryOp = googlesql.ASTUnaryExpressionEnums_OpPlus
	IsUnknownOp    UnaryOp = googlesql.ASTUnaryExpressionEnums_OpIsUnknown
	IsNotUnknownOp UnaryOp = googlesql.ASTUnaryExpressionEnums_OpIsNotUnknown
)

// UnpivotNullFilter enums.
type UnpivotNullFilter = googlesql.ASTUnpivotClauseEnums_NullFilter

const (
	UnspecifiedNullFilter UnpivotNullFilter = googlesql.ASTUnpivotClauseEnums_NullFilterKUnspecified
	IncludeNullFilter     UnpivotNullFilter = googlesql.ASTUnpivotClauseEnums_NullFilterKInclude
	ExcludeNullFilter     UnpivotNullFilter = googlesql.ASTUnpivotClauseEnums_NullFilterKExclude
)

// FrameUnit enums.
type FrameUnit = googlesql.ASTWindowFrameEnums_FrameUnit

const (
	RowsFrameUnit FrameUnit = googlesql.ASTWindowFrameEnums_FrameUnitRows
	Range         FrameUnit = googlesql.ASTWindowFrameEnums_FrameUnitRange
)

// BoundaryType enums.
type BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryType

const (
	UnboundedPreceding BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeUnboundedPreceding
	OffsetPreceding    BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeOffsetPreceding
	CurrentRow         BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeCurrentRow
	OffsetFollowing    BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeOffsetFollowing
	UnboundedFollowing BoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeUnboundedFollowing
)

// HavingModifierEnum enums.
type HavingModifierEnum = googlesql.ASTHavingModifierEnums_ModifierKind

const (
	NotSetHavingModifier HavingModifierEnum = googlesql.ASTHavingModifierEnums_ModifierKindNotSet
	MinHavingModifier    HavingModifierEnum = googlesql.ASTHavingModifierEnums_ModifierKindMin
	MaxHavingModifier    HavingModifierEnum = googlesql.ASTHavingModifierEnums_ModifierKindMax
)

// SelectAsMode enums.
type SelectAsMode = googlesql.ASTSelectAsEnums_AsMode

const (
	NotSetSelectAsMode   SelectAsMode = googlesql.ASTSelectAsEnums_AsModeNotSet
	StructSelectAsMode   SelectAsMode = googlesql.ASTSelectAsEnums_AsModeStruct
	ValueSelectAsMode    SelectAsMode = googlesql.ASTSelectAsEnums_AsModeValue
	TypeNameSelectAsMode SelectAsMode = googlesql.ASTSelectAsEnums_AsModeTypeName
)

// NullFilter enums (alias for UnpivotNullFilter).
type NullFilter = UnpivotNullFilter

// SchemaObjectKind enums.
type SchemaObjectKind = googlesql.SchemaObjectKind

const (
	SchemaObjectKindSwitchMustHaveADefault SchemaObjectKind = googlesql.SchemaObjectKindSchemaObjectKindSwitchMustHaveADefault
	InvalidSchemaObjectKind                SchemaObjectKind = googlesql.SchemaObjectKindKInvalidSchemaObjectKind
	AggregateFunction                      SchemaObjectKind = googlesql.SchemaObjectKindKAggregateFunction
	Constant                               SchemaObjectKind = googlesql.SchemaObjectKindKConstant
	Database                               SchemaObjectKind = googlesql.SchemaObjectKindKDatabase
	ExternalTable                          SchemaObjectKind = googlesql.SchemaObjectKindKExternalTable
	FunctionSchemaObject                   SchemaObjectKind = googlesql.SchemaObjectKindKFunction
	IndexSchemaObject                      SchemaObjectKind = googlesql.SchemaObjectKindKIndex
	MaterializedView                       SchemaObjectKind = googlesql.SchemaObjectKindKMaterializedView
	Model                                  SchemaObjectKind = googlesql.SchemaObjectKindKModel
	Procedure                              SchemaObjectKind = googlesql.SchemaObjectKindKProcedure
	Schema                                 SchemaObjectKind = googlesql.SchemaObjectKindKSchema
	TableSchemaObject                      SchemaObjectKind = googlesql.SchemaObjectKindKTable
	TableFunction                          SchemaObjectKind = googlesql.SchemaObjectKindKTableFunction
	View                                   SchemaObjectKind = googlesql.SchemaObjectKindKView
	SnapshotTable                          SchemaObjectKind = googlesql.SchemaObjectKindKSnapshotTable
)

// ProcedureParameterMode is an alias for ParameterMode to satisfy the printer.
type ProcedureParameterMode = ParameterMode

const (
	OutParameterMode   ProcedureParameterMode = Out
	InOutParameterMode ProcedureParameterMode = InOut
)

// DateTimeTypeKind enums for DateOrTimeLiteral.TypeKind()
type DateTimeTypeKind = googlesql.TypeKind

const (
	Date      DateTimeTypeKind = googlesql.TypeKindTypeDate
	Time      DateTimeTypeKind = googlesql.TypeKindTypeTime
	Datetime  DateTimeTypeKind = googlesql.TypeKindTypeDatetime
	Timestamp DateTimeTypeKind = googlesql.TypeKindTypeTimestamp
)
