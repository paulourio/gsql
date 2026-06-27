package sql

// IsInsideOfWhereClause returns true when n is inside a WHERE clause,
// possibly through AND/OR chains.
func IsInsideOfWhereClause(n Node) bool {
	for p := n.Parent(); p != nil; p = p.Parent() {
		switch p.Kind() {
		case WhereClauseKind:
			return true
		case AndExprKind, OrExprKind:
			continue
		default:
			return false
		}
	}
	return false
}

// IsInsideOfOnClause returns true when n is inside an ON clause,
// possibly through AND/OR chains.
func IsInsideOfOnClause(n Node) bool {
	for p := n.Parent(); p != nil; p = p.Parent() {
		k := p.Kind()
		if k == OnClauseKind {
			return true
		}
		if k != AndExprKind && k != OrExprKind {
			return false
		}
	}
	return false
}

// IsInsideOfMergeStatement returns true when n's immediate parent is a
// MERGE statement.
func IsInsideOfMergeStatement(n Node) bool {
	p := n.Parent()
	if p == nil {
		return false
	}
	return p.Kind() == MergeStatementKind
}
