package ast

import "github.com/goccy/go-googlesql"

// ASTBinaryExpression enums.
const (
	NotSetOp     = googlesql.ASTBinaryExpressionEnums_OpNotSet
	LikeOp       = googlesql.ASTBinaryExpressionEnums_OpLike
	IsOp         = googlesql.ASTBinaryExpressionEnums_OpIs
	EqOp         = googlesql.ASTBinaryExpressionEnums_OpEq
	NeOp         = googlesql.ASTBinaryExpressionEnums_OpNe
	Ne2Op        = googlesql.ASTBinaryExpressionEnums_OpNe2
	GtOp         = googlesql.ASTBinaryExpressionEnums_OpGt
	LtOp         = googlesql.ASTBinaryExpressionEnums_OpLt
	GeOp         = googlesql.ASTBinaryExpressionEnums_OpGe
	LeOp         = googlesql.ASTBinaryExpressionEnums_OpLe
	BitwiseOrOp  = googlesql.ASTBinaryExpressionEnums_OpBitwiseOr
	BitwiseXorOp = googlesql.ASTBinaryExpressionEnums_OpBitwiseXor
	BitwiseAndOp = googlesql.ASTBinaryExpressionEnums_OpBitwiseAnd
	PlusOp       = googlesql.ASTBinaryExpressionEnums_OpPlus
	MinusOp      = googlesql.ASTBinaryExpressionEnums_OpMinus
	MultiplyOp   = googlesql.ASTBinaryExpressionEnums_OpMultiply
	DivideOp     = googlesql.ASTBinaryExpressionEnums_OpDivide
	ConcatOp     = googlesql.ASTBinaryExpressionEnums_OpConcatOp
	DistinctOp   = googlesql.ASTBinaryExpressionEnums_OpDistinct
)

// ASTUnaryExpression enums.
const (
	NotSetUnaryOp       = googlesql.ASTUnaryExpressionEnums_OpNotSet
	NotUnaryOp          = googlesql.ASTUnaryExpressionEnums_OpNot
	BitwiseNotUnaryOp   = googlesql.ASTUnaryExpressionEnums_OpBitwiseNot
	MinusUnaryOp        = googlesql.ASTUnaryExpressionEnums_OpMinus
	PlusUnaryOp         = googlesql.ASTUnaryExpressionEnums_OpPlus
	IsUnknownUnaryOp    = googlesql.ASTUnaryExpressionEnums_OpIsUnknown
	IsNotUnknownUnaryOp = googlesql.ASTUnaryExpressionEnums_OpIsNotUnknown
)

// ASTJoin enums.
const (
	DefaultJoinType = googlesql.ASTJoinEnums_JoinTypeDefaultJoinType
	CommaJoinType   = googlesql.ASTJoinEnums_JoinTypeComma
	CrossJoinType   = googlesql.ASTJoinEnums_JoinTypeCross
	FullJoinType    = googlesql.ASTJoinEnums_JoinTypeFull
	InnerJoinType   = googlesql.ASTJoinEnums_JoinTypeInner
	LeftJoinType    = googlesql.ASTJoinEnums_JoinTypeLeft
	RightJoinType   = googlesql.ASTJoinEnums_JoinTypeRight
)

// ASTJoinHint enums.
const (
	NoJoinHint     = googlesql.ASTJoinEnums_JoinHintNoJoinHint
	HashJoinHint   = googlesql.ASTJoinEnums_JoinHintHash
	LookupJoinHint = googlesql.ASTJoinEnums_JoinHintLookup
)

// ASTExpressionSubquery enums.
const (
	NoneExpressionSubquery   = googlesql.ASTExpressionSubqueryEnums_ModifierNone
	ArrayExpressionSubquery  = googlesql.ASTExpressionSubqueryEnums_ModifierArray
	ExistsExpressionSubquery = googlesql.ASTExpressionSubqueryEnums_ModifierExists
	ValueExpressionSubquery  = googlesql.ASTExpressionSubqueryEnums_ModifierValue
)

// ASTHavingModifier enums.
const (
	NotSetHavingModifier = googlesql.ASTHavingModifierEnums_ModifierKindNotSet
	MinHavingModifier    = googlesql.ASTHavingModifierEnums_ModifierKindMin
	MaxHavingModifier    = googlesql.ASTHavingModifierEnums_ModifierKindMax
)

// ASTSetOperation enums.
const (
	NotSetSetOperation    = googlesql.ASTSetOperationEnums_OperationTypeNotSet
	UnionSetOperation     = googlesql.ASTSetOperationEnums_OperationTypeUnion
	ExceptSetOperation    = googlesql.ASTSetOperationEnums_OperationTypeExcept
	IntersectSetOperation = googlesql.ASTSetOperationEnums_OperationTypeIntersect
)

// AllOrDistinct enums.
const (
	UnspecifiedSetOperation = googlesql.ASTSetOperationEnums_AllOrDistinctAllOrDistinctNotSet
	AllSetOperation         = googlesql.ASTSetOperationEnums_AllOrDistinctAll
	DistinctSetOperation    = googlesql.ASTSetOperationEnums_AllOrDistinctDistinct
)

// ASTSelectAs enums.
const (
	NotSetAsMode   = googlesql.ASTSelectAsEnums_AsModeNotSet
	StructAsMode   = googlesql.ASTSelectAsEnums_AsModeStruct
	ValueAsMode    = googlesql.ASTSelectAsEnums_AsModeValue
	TypeNameAsMode = googlesql.ASTSelectAsEnums_AsModeTypeName
)

// ASTSampleSize enums.
const (
	NotSetUnit  = googlesql.ASTSampleSizeEnums_UnitNotSet
	RowsUnit    = googlesql.ASTSampleSizeEnums_UnitRows
	PercentUnit = googlesql.ASTSampleSizeEnums_UnitPercent
)

// ASTTemplatedParameterType enums.
const (
	UninitializedTemplatedTypeKind = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindUninitialized
	AnyTypeTemplatedTypeKind       = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyType
	AnyProtoTemplatedTypeKind      = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyProto
	AnyEnumTemplatedTypeKind       = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyEnum
	AnyStructTemplatedTypeKind     = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyStruct
	AnyArrayTemplatedTypeKind      = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyArray
	AnyTableTemplatedTypeKind      = googlesql.ASTTemplatedParameterTypeEnums_TemplatedTypeKindAnyTable
)

// ASTOrderingExpression enums.
const (
	NotSetOrderingSpec = googlesql.ASTOrderingExpressionEnums_OrderingSpecNotSet
	AscOrderingSpec    = googlesql.ASTOrderingExpressionEnums_OrderingSpecAsc
	DescOrderingSpec   = googlesql.ASTOrderingExpressionEnums_OrderingSpecDesc
)

// DeterminismLevel enums.
const (
	UnspecifiedDeterminismLevel      = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelDeterminismUnspecified
	DeterministicDeterminismLevel    = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelDeterministic
	NotDeterministicDeterminismLevel = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelNotDeterministic
	ImmutableDeterminismLevel        = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelImmutable
	StableDeterminismLevel           = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelStable
	VolatileDeterminismLevel         = googlesql.ASTCreateFunctionStmtBaseEnums_DeterminismLevelVolatile
)

// Type Kinds
const (
	SwitchMustHaveADefaultTypeKind = googlesql.TypeKindTypeKindSwitchMustHaveADefault
	UnknownTypeKind                = googlesql.TypeKindTypeUnknown
	Int32                          = googlesql.TypeKindTypeInt32
	Int64                          = googlesql.TypeKindTypeInt64
	UInt32                         = googlesql.TypeKindTypeUint32
	UInt64                         = googlesql.TypeKindTypeUint64
	Bool                           = googlesql.TypeKindTypeBool
	Float                          = googlesql.TypeKindTypeFloat
	Double                         = googlesql.TypeKindTypeDouble
	String                         = googlesql.TypeKindTypeString
	Bytes                          = googlesql.TypeKindTypeBytes
	Date                           = googlesql.TypeKindTypeDate
	Enum                           = googlesql.TypeKindTypeEnum
	Array                          = googlesql.TypeKindTypeArray
	Struct                         = googlesql.TypeKindTypeStruct
	Proto                          = googlesql.TypeKindTypeProto
	Timestamp                      = googlesql.TypeKindTypeTimestamp
	Time                           = googlesql.TypeKindTypeTime
	Datetime                       = googlesql.TypeKindTypeDatetime
	Geography                      = googlesql.TypeKindTypeGeography
	Numeric                        = googlesql.TypeKindTypeNumeric
	Bignumeric                     = googlesql.TypeKindTypeBignumeric
	Extended                       = googlesql.TypeKindTypeExtended
	JSON                           = googlesql.TypeKindTypeJson
	Interval                       = googlesql.TypeKindTypeInterval
	Tokenlist                      = googlesql.TypeKindTypeTokenlist
	Range                          = googlesql.TypeKindTypeRange
	GraphElement                   = googlesql.TypeKindTypeGraphElement
	Map                            = googlesql.TypeKindTypeMap
	UUID                           = googlesql.TypeKindTypeUuid
	GraphPath                      = googlesql.TypeKindTypeGraphPath
	Measure                        = googlesql.TypeKindTypeMeasure
	Row                            = googlesql.TypeKindTypeRow
)

// Function parameter modes.
const (
	NotSetParameterMode = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeNotSet
	InParameterMode     = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeIn
	OutParameterMode    = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeOut
	InoutParameterMode  = googlesql.ASTFunctionParameterEnums_ProcedureParameterModeInout
)

// FrameUnit enums.
const (
	RowsFrameUnit  = googlesql.ASTWindowFrameEnums_FrameUnitRows
	RangeFrameUnit = googlesql.ASTWindowFrameEnums_FrameUnitRange
)

// WindowFrameExpr enums.
const (
	UnboundedPrecedingBoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeUnboundedPreceding
	OffsetPrecedingBoundaryType    = googlesql.ASTWindowFrameExprEnums_BoundaryTypeOffsetPreceding
	CurrentRowBoundaryType         = googlesql.ASTWindowFrameExprEnums_BoundaryTypeCurrentRow
	OffsetFollowingBoundaryType    = googlesql.ASTWindowFrameExprEnums_BoundaryTypeOffsetFollowing
	UnboundedFollowingBoundaryType = googlesql.ASTWindowFrameExprEnums_BoundaryTypeUnboundedFollowing
)

// Sql Security enums.
const (
	UnspecifiedSQLSecurity = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityUnspecified
	DefinerSQLSecurity     = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityDefiner
	InvokerSQLSecurity     = googlesql.ASTCreateStatementEnums_SqlSecuritySqlSecurityInvoker
)

// Scope enums.
const (
	DefaultScope   = googlesql.ASTCreateStatementEnums_ScopeDefaultScope
	PrivateScope   = googlesql.ASTCreateStatementEnums_ScopePrivate
	PublicScope    = googlesql.ASTCreateStatementEnums_ScopePublic
	TemporaryScope = googlesql.ASTCreateStatementEnums_ScopeTemporary
)

// Drop Statement modes.
const (
	UnspecifiedDropMode = googlesql.ASTDropStatementEnums_DropModeDropModeUnspecified
	RestrictDropMode    = googlesql.ASTDropStatementEnums_DropModeRestrict
	CascadeDropMode     = googlesql.ASTDropStatementEnums_DropModeCascade
)

// NullFilter enums.
const (
	UnspecifiedNullFilter = googlesql.ASTUnpivotClauseEnums_NullFilterKUnspecified
	IncludeNullFilter     = googlesql.ASTUnpivotClauseEnums_NullFilterKInclude
	ExcludeNullFilter     = googlesql.ASTUnpivotClauseEnums_NullFilterKExclude
)
