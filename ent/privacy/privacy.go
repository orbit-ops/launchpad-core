// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"context"

	"github.com/orbit-ops/launchpad-core/ent"

	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, ent.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op ent.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op ent.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m ent.Mutation) error {
		return Denyf("ent/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The AccessQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AccessQueryRuleFunc func(context.Context, *ent.AccessQuery) error

// EvalQuery return f(ctx, q).
func (f AccessQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AccessQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AccessQuery", q)
}

// The AccessMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AccessMutationRuleFunc func(context.Context, *ent.AccessMutation) error

// EvalMutation calls f(ctx, m).
func (f AccessMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AccessMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AccessMutation", m)
}

// The ActionTokensQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ActionTokensQueryRuleFunc func(context.Context, *ent.ActionTokensQuery) error

// EvalQuery return f(ctx, q).
func (f ActionTokensQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ActionTokensQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ActionTokensQuery", q)
}

// The ActionTokensMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ActionTokensMutationRuleFunc func(context.Context, *ent.ActionTokensMutation) error

// EvalMutation calls f(ctx, m).
func (f ActionTokensMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ActionTokensMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ActionTokensMutation", m)
}

// The ApiKeyQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ApiKeyQueryRuleFunc func(context.Context, *ent.ApiKeyQuery) error

// EvalQuery return f(ctx, q).
func (f ApiKeyQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ApiKeyQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ApiKeyQuery", q)
}

// The ApiKeyMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ApiKeyMutationRuleFunc func(context.Context, *ent.ApiKeyMutation) error

// EvalMutation calls f(ctx, m).
func (f ApiKeyMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ApiKeyMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ApiKeyMutation", m)
}

// The ApprovalQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type ApprovalQueryRuleFunc func(context.Context, *ent.ApprovalQuery) error

// EvalQuery return f(ctx, q).
func (f ApprovalQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ApprovalQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.ApprovalQuery", q)
}

// The ApprovalMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type ApprovalMutationRuleFunc func(context.Context, *ent.ApprovalMutation) error

// EvalMutation calls f(ctx, m).
func (f ApprovalMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.ApprovalMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.ApprovalMutation", m)
}

// The AuditQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type AuditQueryRuleFunc func(context.Context, *ent.AuditQuery) error

// EvalQuery return f(ctx, q).
func (f AuditQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AuditQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.AuditQuery", q)
}

// The AuditMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type AuditMutationRuleFunc func(context.Context, *ent.AuditMutation) error

// EvalMutation calls f(ctx, m).
func (f AuditMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.AuditMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.AuditMutation", m)
}

// The MissionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type MissionQueryRuleFunc func(context.Context, *ent.MissionQuery) error

// EvalQuery return f(ctx, q).
func (f MissionQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MissionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.MissionQuery", q)
}

// The MissionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type MissionMutationRuleFunc func(context.Context, *ent.MissionMutation) error

// EvalMutation calls f(ctx, m).
func (f MissionMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.MissionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.MissionMutation", m)
}

// The RequestQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RequestQueryRuleFunc func(context.Context, *ent.RequestQuery) error

// EvalQuery return f(ctx, q).
func (f RequestQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RequestQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RequestQuery", q)
}

// The RequestMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RequestMutationRuleFunc func(context.Context, *ent.RequestMutation) error

// EvalMutation calls f(ctx, m).
func (f RequestMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RequestMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RequestMutation", m)
}

// The RocketQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type RocketQueryRuleFunc func(context.Context, *ent.RocketQuery) error

// EvalQuery return f(ctx, q).
func (f RocketQueryRuleFunc) EvalQuery(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RocketQuery); ok {
		return f(ctx, q)
	}
	return Denyf("ent/privacy: unexpected query type %T, expect *ent.RocketQuery", q)
}

// The RocketMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type RocketMutationRuleFunc func(context.Context, *ent.RocketMutation) error

// EvalMutation calls f(ctx, m).
func (f RocketMutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if m, ok := m.(*ent.RocketMutation); ok {
		return f(ctx, m)
	}
	return Denyf("ent/privacy: unexpected mutation type %T, expect *ent.RocketMutation", m)
}