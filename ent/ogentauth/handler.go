

// Code generated by ogent-auth, DO NOT EDIT.
package ogentauth

import (
  "context"
  "errors"
  "github.com/orbit-ops/launchpad-core/ent/ogent"

  sessions "github.com/tiagoposse/go-auth/sessions"
  authz "github.com/tiagoposse/go-auth/authorization"
)

type ISessionHandler interface {
  ValidateApiKeyAuth(context.Context, string) (authz.ScopedSession, error)
  ValidateCookieAuth(context.Context, string) (authz.ScopedSession, error)

  CreateSessionToken(context.Context, any) (string, error)
  ValidateSessionToken(context.Context, string) (authz.ScopedSession, error)
}

type SecurityHandler struct {
  ISessionHandler
}

func NewSecurityHandler(sess ISessionHandler) *SecurityHandler {
  return &SecurityHandler{
    ISessionHandler: sess,
  }
}


func (h *SecurityHandler) HandleApiKeyAuth(c context.Context, operationName string, t ogent.ApiKeyAuth) (context.Context, error) {
  session, err := h.ValidateApiKeyAuth(c, t.APIKey)
	if err != nil {
		return c, err
	}

  ctx := context.WithValue(c, sessions.ContextSessionKey{}, session)
  if err := ValidateScopes(operationName, session.GetScopes()); err != nil {
    return ctx, err
  }

	return ctx, nil
}

func (h *SecurityHandler) HandleCookieAuth(c context.Context, operationName string, t ogent.CookieAuth) (context.Context, error) {
  session, err := h.ValidateCookieAuth(c, t.APIKey)
	if err != nil {
		return c, err
	}

  ctx := context.WithValue(c, sessions.ContextSessionKey{}, session)
  if err := ValidateScopes(operationName, session.GetScopes()); err != nil {
    return ctx, err
  }

	return ctx, nil
}


func ValidateScopes(operationName string, reqScopes authz.Scopes) error {
  if opScopes, ok := scopes[operationName]; ok {
    found := false
    for _, reqScope := range reqScopes {
      for _, opScope := range opScopes {
        if reqScope == opScope {
          found = true
          break
        }
      }
      if found {
        break
      }
    }

    if !found {
      return errors.New("no scopes match")
    }
  }

  return nil
}

var scopes map[string]authz.Scopes = map[string]authz.Scopes{
  "createApiKey": {
    "admin",
  },
  "createMission": {
    "admin",
  },
  "createMissionRequests": {
    "admin",
  },
  "createMissionRockets": {
    "admin",
  },
  "createRocket": {
    "admin",
  },
  "deleteApiKey": {
    "admin",
  },
  "deleteMission": {
    "admin",
  },
  "deleteMissionRequests": {
    "admin",
  },
  "deleteMissionRockets": {
    "admin",
  },
  "deleteRocket": {
    "admin",
  },
  "listApiKey": {
    "admin",
  },
  "listMission": {
    "admin",
  },
  "listMissionRequests": {
    "admin",
  },
  "listMissionRockets": {
    "admin",
  },
  "listRocket": {
    "admin",
  },
  "readApiKey": {
    "admin",
  },
  "readMission": {
    "admin",
  },
  "readMissionRequests": {
    "admin",
  },
  "readMissionRockets": {
    "admin",
  },
  "readRocket": {
    "admin",
  },
  "updateApiKey": {
    "admin",
  },
  "updateMission": {
    "admin",
  },
  "updateMissionRequests": {
    "admin",
  },
  "updateMissionRockets": {
    "admin",
  },
  "updateRocket": {
    "admin",
  },
}
