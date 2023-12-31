// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/actiontokens"
	"github.com/orbit-ops/launchpad-core/ent/apikey"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/audit"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/request"
	"github.com/orbit-ops/launchpad-core/ent/rocket"
	"github.com/orbit-ops/launchpad-core/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accessFields := schema.Access{}.Fields()
	_ = accessFields
	// accessDescRolledBack is the schema descriptor for rolled_back field.
	accessDescRolledBack := accessFields[3].Descriptor()
	// access.DefaultRolledBack holds the default value on creation for the rolled_back field.
	access.DefaultRolledBack = accessDescRolledBack.Default.(bool)
	// accessDescID is the schema descriptor for id field.
	accessDescID := accessFields[0].Descriptor()
	// access.DefaultID holds the default value on creation for the id field.
	access.DefaultID = accessDescID.Default.(func() uuid.UUID)
	actiontokensFields := schema.ActionTokens{}.Fields()
	_ = actiontokensFields
	// actiontokensDescID is the schema descriptor for id field.
	actiontokensDescID := actiontokensFields[0].Descriptor()
	// actiontokens.DefaultID holds the default value on creation for the id field.
	actiontokens.DefaultID = actiontokensDescID.Default.(func() uuid.UUID)
	apikeyFields := schema.ApiKey{}.Fields()
	_ = apikeyFields
	// apikeyDescName is the schema descriptor for name field.
	apikeyDescName := apikeyFields[0].Descriptor()
	// apikey.NameValidator is a validator for the "name" field. It is called by the builders before save.
	apikey.NameValidator = apikeyDescName.Validators[0].(func(string) error)
	approvalFields := schema.Approval{}.Fields()
	_ = approvalFields
	// approvalDescRevoked is the schema descriptor for revoked field.
	approvalDescRevoked := approvalFields[4].Descriptor()
	// approval.DefaultRevoked holds the default value on creation for the revoked field.
	approval.DefaultRevoked = approvalDescRevoked.Default.(bool)
	// approvalDescID is the schema descriptor for id field.
	approvalDescID := approvalFields[0].Descriptor()
	// approval.DefaultID holds the default value on creation for the id field.
	approval.DefaultID = approvalDescID.Default.(func() uuid.UUID)
	auditFields := schema.Audit{}.Fields()
	_ = auditFields
	// auditDescAuthor is the schema descriptor for author field.
	auditDescAuthor := auditFields[2].Descriptor()
	// audit.AuthorValidator is a validator for the "author" field. It is called by the builders before save.
	audit.AuthorValidator = auditDescAuthor.Validators[0].(func(string) error)
	// auditDescID is the schema descriptor for id field.
	auditDescID := auditFields[0].Descriptor()
	// audit.DefaultID holds the default value on creation for the id field.
	audit.DefaultID = auditDescID.Default.(func() uuid.UUID)
	missionFields := schema.Mission{}.Fields()
	_ = missionFields
	// missionDescName is the schema descriptor for name field.
	missionDescName := missionFields[1].Descriptor()
	// mission.NameValidator is a validator for the "name" field. It is called by the builders before save.
	mission.NameValidator = missionDescName.Validators[0].(func(string) error)
	// missionDescMinApprovers is the schema descriptor for min_approvers field.
	missionDescMinApprovers := missionFields[3].Descriptor()
	// mission.MinApproversValidator is a validator for the "min_approvers" field. It is called by the builders before save.
	mission.MinApproversValidator = missionDescMinApprovers.Validators[0].(func(int) error)
	// missionDescID is the schema descriptor for id field.
	missionDescID := missionFields[0].Descriptor()
	// mission.DefaultID holds the default value on creation for the id field.
	mission.DefaultID = missionDescID.Default.(func() uuid.UUID)
	requestFields := schema.Request{}.Fields()
	_ = requestFields
	// requestDescID is the schema descriptor for id field.
	requestDescID := requestFields[0].Descriptor()
	// request.DefaultID holds the default value on creation for the id field.
	request.DefaultID = requestDescID.Default.(func() uuid.UUID)
	rocketFields := schema.Rocket{}.Fields()
	_ = rocketFields
	// rocketDescName is the schema descriptor for name field.
	rocketDescName := rocketFields[1].Descriptor()
	// rocket.NameValidator is a validator for the "name" field. It is called by the builders before save.
	rocket.NameValidator = rocketDescName.Validators[0].(func(string) error)
	// rocketDescCode is the schema descriptor for code field.
	rocketDescCode := rocketFields[3].Descriptor()
	// rocket.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	rocket.CodeValidator = rocketDescCode.Validators[0].(func(string) error)
	// rocketDescConfig is the schema descriptor for config field.
	rocketDescConfig := rocketFields[4].Descriptor()
	// rocket.DefaultConfig holds the default value on creation for the config field.
	rocket.DefaultConfig = rocketDescConfig.Default.(map[string]string)
	// rocketDescID is the schema descriptor for id field.
	rocketDescID := rocketFields[0].Descriptor()
	// rocket.DefaultID holds the default value on creation for the id field.
	rocket.DefaultID = rocketDescID.Default.(func() uuid.UUID)
}
