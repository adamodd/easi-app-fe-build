package cedarcore

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ld "gopkg.in/launchdarkly/go-server-sdk.v5"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/guregu/null"
	"github.com/guregu/null/zero"
	"go.uber.org/zap"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/apperrors"
	apiclient "github.com/cmsgov/easi-app/pkg/cedar/core/gen/client"
	apideployments "github.com/cmsgov/easi-app/pkg/cedar/core/gen/client/deployment"
	apiroles "github.com/cmsgov/easi-app/pkg/cedar/core/gen/client/role"
	apisystems "github.com/cmsgov/easi-app/pkg/cedar/core/gen/client/system"
	wiremodels "github.com/cmsgov/easi-app/pkg/cedar/core/gen/models"
	"github.com/cmsgov/easi-app/pkg/flags"
	"github.com/cmsgov/easi-app/pkg/models"
)

const (
	cedarCoreEnabledKey     = "cedarCoreEnabled"
	cedarCoreEnabledDefault = false
	cedarRoleApplication    = "alfabet" // used for queries to GET /role endpoint
)

// NewClient builds the type that holds a connection to the CEDAR Core API
func NewClient(cedarHost string, cedarAPIKey string, ldClient *ld.LDClient) *Client {
	fnEmit := func(ctx context.Context) bool {
		lduser := flags.Principal(ctx)
		result, err := ldClient.BoolVariation(cedarCoreEnabledKey, lduser, cedarCoreEnabledDefault)
		if err != nil {
			appcontext.ZLogger(ctx).Info(
				"problem evaluating feature flag",
				zap.Error(err),
				zap.String("flagName", cedarCoreEnabledKey),
				zap.Bool("flagDefault", cedarCoreEnabledDefault),
				zap.Bool("flagResult", result),
			)
		}
		return result
	}

	hc := http.DefaultClient

	return &Client{
		cedarCoreEnabled: fnEmit,
		auth: httptransport.APIKeyAuth(
			"x-Gateway-APIKey",
			"header",
			cedarAPIKey,
		),
		sdk: apiclient.New(
			httptransport.New(
				cedarHost,
				apiclient.DefaultBasePath,
				apiclient.DefaultSchemes,
			),
			strfmt.Default,
		),
		hc: hc,
	}
}

// Client represents a connection to the CEDAR Core API
type Client struct {
	cedarCoreEnabled func(context.Context) bool
	auth             runtime.ClientAuthInfoWriter
	sdk              *apiclient.CEDARCoreAPI
	hc               *http.Client
}

// GetSystemSummary makes a GET call to the /system/summary endpoint
func (c *Client) GetSystemSummary(ctx context.Context) ([]*models.CedarSystem, error) {
	if !c.cedarCoreEnabled(ctx) {
		appcontext.ZLogger(ctx).Info("CEDAR Core is disabled")
		return []*models.CedarSystem{}, nil
	}

	// Construct the parameters
	params := apisystems.NewSystemSummaryFindListParams()
	params.HTTPClient = c.hc

	// Make the API call
	resp, err := c.sdk.System.SystemSummaryFindList(params, c.auth)
	if err != nil {
		return []*models.CedarSystem{}, err
	}

	if resp.Payload == nil {
		return []*models.CedarSystem{}, fmt.Errorf("no body received")
	}

	// Convert the auto-generated struct to our own pkg/models struct
	retVal := []*models.CedarSystem{}

	// Populate the SystemSummary field by converting each item in resp.Payload.SystemSummary
	for _, sys := range resp.Payload.SystemSummary {
		retVal = append(retVal, &models.CedarSystem{
			ID:                      *sys.ID,
			Name:                    *sys.Name,
			Description:             sys.Description,
			Acronym:                 sys.Acronym,
			Status:                  sys.Status,
			BusinessOwnerOrg:        sys.BusinessOwnerOrg,
			BusinessOwnerOrgComp:    sys.BusinessOwnerOrgComp,
			SystemMaintainerOrg:     sys.SystemMaintainerOrg,
			SystemMaintainerOrgComp: sys.SystemMaintainerOrgComp,
		})
	}

	return retVal, nil
}

// GetSystem makes a GET call to the /system/summary/{id} endpoint
func (c *Client) GetSystem(ctx context.Context, id string) (*models.CedarSystem, error) {
	if !c.cedarCoreEnabled(ctx) {
		appcontext.ZLogger(ctx).Info("CEDAR Core is disabled")
		return &models.CedarSystem{}, nil
	}

	// Construct the parameters
	params := apisystems.NewSystemSummaryFindByIDParams()
	params.SetID(id)
	params.HTTPClient = c.hc

	// Make the API call
	resp, err := c.sdk.System.SystemSummaryFindByID(params, c.auth)
	if err != nil {
		return &models.CedarSystem{}, err
	}

	if resp.Payload == nil {
		return &models.CedarSystem{}, fmt.Errorf("no body received")
	}

	responseArray := resp.Payload.SystemSummary

	if len(responseArray) == 0 {
		return nil, &apperrors.ResourceNotFoundError{Err: fmt.Errorf("no system found"), Resource: models.CedarSystem{}}
	}

	// Convert the auto-generated struct to our own pkg/models struct
	return &models.CedarSystem{
		ID:                      *responseArray[0].ID,
		Name:                    *responseArray[0].Name,
		Description:             responseArray[0].Description,
		Acronym:                 responseArray[0].Acronym,
		Status:                  responseArray[0].Status,
		BusinessOwnerOrg:        responseArray[0].BusinessOwnerOrg,
		BusinessOwnerOrgComp:    responseArray[0].BusinessOwnerOrgComp,
		SystemMaintainerOrg:     responseArray[0].SystemMaintainerOrg,
		SystemMaintainerOrgComp: responseArray[0].SystemMaintainerOrgComp,
	}, nil
}

// GetDeploymentsOptionalParams represents the optional parameters that can be used to filter deployments when searching through the CEDAR API
type GetDeploymentsOptionalParams struct {
	DeploymentType null.String
	State          null.String
	Status         null.String
}

// GetDeployments makes a GET call to the /deployment endpoint
func (c *Client) GetDeployments(ctx context.Context, systemID string, optionalParams *GetDeploymentsOptionalParams) ([]*models.CedarDeployment, error) {
	if !c.cedarCoreEnabled(ctx) {
		appcontext.ZLogger(ctx).Info("CEDAR Core is disabled")
		return []*models.CedarDeployment{}, nil
	}

	// Construct the parameters
	params := apideployments.NewDeploymentFindListParams()
	params.SetSystemID(systemID)
	params.HTTPClient = c.hc

	if optionalParams != nil {
		if optionalParams.DeploymentType.Ptr() != nil {
			params.SetDeploymentType(optionalParams.DeploymentType.Ptr())
		}

		if optionalParams.State.Ptr() != nil {
			params.SetState(optionalParams.State.Ptr())
		}

		if optionalParams.Status.Ptr() != nil {
			params.SetStatus(optionalParams.Status.Ptr())
		}
	}

	// Make the API call
	resp, err := c.sdk.Deployment.DeploymentFindList(params, c.auth)
	if err != nil {
		return []*models.CedarDeployment{}, err
	}

	if resp.Payload == nil {
		return []*models.CedarDeployment{}, fmt.Errorf("no body received")
	}

	responseArray := resp.Payload.Deployments

	if len(responseArray) == 0 {
		return nil, &apperrors.ResourceNotFoundError{Err: fmt.Errorf("no deployments found"), Resource: []*models.CedarDeployment{}}
	}

	// Convert the auto-generated struct to our own pkg/models struct
	retVal := []*models.CedarDeployment{}

	// Populate the Deployment field by converting each item in resp.Payload.Deployments
	// generated swagger client turns JSON nulls into Go zero values, so use null/zero package to convert them back to nullable values
	for _, deployment := range resp.Payload.Deployments {
		if deployment.ID == nil {
			appcontext.ZLogger(ctx).Error("Error decoding deployment; deployment ID was null", zap.String("systemID", systemID))
			continue
		}

		if deployment.Name == nil {
			appcontext.ZLogger(ctx).Error("Error decoding deployment; deployment name was null", zap.String("systemID", systemID))
			continue
		}

		if deployment.SystemID == nil {
			appcontext.ZLogger(ctx).Error("Error decoding deployment; deployment system ID was null", zap.String("systemID", systemID))
			continue
		}

		retDeployment := &models.CedarDeployment{
			ID:                *deployment.ID,
			Name:              *deployment.Name,
			SystemID:          *deployment.SystemID,
			StartDate:         zero.TimeFrom(time.Time(deployment.StartDate)),
			EndDate:           zero.TimeFrom(time.Time(deployment.EndDate)),
			IsHotSite:         zero.StringFrom(deployment.IsHotSite),
			Description:       zero.StringFrom(deployment.Description),
			ContractorName:    zero.StringFrom(deployment.ContractorName),
			SystemVersion:     zero.StringFrom(deployment.ContractorName),
			HasProductionData: zero.StringFrom(deployment.HasProductionData),

			// TODO - assumes no nulls in array returned from query
			ReplicatedSystemElements: deployment.ReplicatedSystemElements,

			DeploymentType:      zero.StringFrom(deployment.DeploymentType),
			SystemName:          zero.StringFrom(deployment.SystemName),
			DeploymentElementID: zero.StringFrom(deployment.DeploymentElementID),
			State:               zero.StringFrom(deployment.State),
			Status:              zero.StringFrom(deployment.Status),
			WanType:             zero.StringFrom(deployment.WanType),
		}

		if deployment.DataCenter != nil {
			retDataCenter := &models.CedarDataCenter{
				ID:           zero.StringFrom(deployment.DataCenter.ID),
				Name:         zero.StringFrom(deployment.DataCenter.Name),
				Version:      zero.StringFrom(deployment.DataCenter.Version),
				Description:  zero.StringFrom(deployment.DataCenter.Description),
				State:        zero.StringFrom(deployment.DataCenter.State),
				Status:       zero.StringFrom(deployment.DataCenter.Status),
				StartDate:    zero.TimeFrom(time.Time(deployment.DataCenter.StartDate)),
				EndDate:      zero.TimeFrom(time.Time(deployment.DataCenter.EndDate)),
				Address1:     zero.StringFrom(deployment.DataCenter.Address1),
				Address2:     zero.StringFrom(deployment.DataCenter.Address2),
				City:         zero.StringFrom(deployment.DataCenter.City),
				AddressState: zero.StringFrom(deployment.DataCenter.AddressState),
				Zip:          zero.StringFrom(deployment.DataCenter.Zip),
			}
			retDeployment.DataCenter = retDataCenter
		}

		retVal = append(retVal, retDeployment)
	}

	return retVal, nil
}

// GetRolesBySystem makes a GET call to the /role endpoint using a system ID and an optional role type ID
func (c *Client) GetRolesBySystem(ctx context.Context, systemID string, roleTypeID null.String) ([]*models.CedarRole, error) {
	if !c.cedarCoreEnabled(ctx) {
		appcontext.ZLogger(ctx).Info("CEDAR Core is disabled")
		return []*models.CedarRole{}, nil
	}

	// Construct the parameters
	params := apiroles.NewRoleFindByIDParams()
	params.SetApplication(cedarRoleApplication)
	params.SetObjectID(&systemID)
	params.HTTPClient = c.hc

	if roleTypeID.Ptr() != nil {
		params.SetRoleTypeID(roleTypeID.Ptr())
	}

	// Make the API call
	resp, err := c.sdk.Role.RoleFindByID(params, c.auth)
	if err != nil {
		return []*models.CedarRole{}, err
	}

	if resp.Payload == nil {
		return []*models.CedarRole{}, fmt.Errorf("no body received")
	}

	if len(resp.Payload.Roles) == 0 {
		return nil, &apperrors.ResourceNotFoundError{Err: fmt.Errorf("no deployments found"), Resource: []*models.CedarRole{}}
	}

	// Convert the auto-generated struct to our own pkg/models struct
	return convertWireRolesToModel(ctx, resp.Payload.Roles), nil
}

// GetRolesByRoleID makes a GET call to the /role endpoint using a role ID
func (c *Client) GetRolesByRoleID(ctx context.Context, roleID string) ([]*models.CedarRole, error) {
	if !c.cedarCoreEnabled(ctx) {
		appcontext.ZLogger(ctx).Info("CEDAR Core is disabled")
		return []*models.CedarRole{}, nil
	}

	// Construct the parameters
	params := apiroles.NewRoleFindByIDParams()
	params.SetApplication(cedarRoleApplication)
	params.SetRoleID(&roleID)
	params.HTTPClient = c.hc

	// Make the API call
	resp, err := c.sdk.Role.RoleFindByID(params, c.auth)
	if err != nil {
		return []*models.CedarRole{}, err
	}

	if resp.Payload == nil {
		return []*models.CedarRole{}, fmt.Errorf("no body received")
	}

	if len(resp.Payload.Roles) == 0 {
		return nil, &apperrors.ResourceNotFoundError{Err: fmt.Errorf("no deployments found"), Resource: []*models.CedarRole{}}
	}

	// Convert the auto-generated struct to our own pkg/models struct
	return convertWireRolesToModel(ctx, resp.Payload.Roles), nil
}

func convertWireRolesToModel(ctx context.Context, wireRoles []*wiremodels.Role) []*models.CedarRole {
	roles := []*models.CedarRole{}

	for _, wireRole := range wireRoles {
		if wireRole.Application == nil {
			appcontext.ZLogger(ctx).Error("Error decoding role; role Application was null")
			continue
		}

		if wireRole.ObjectID == nil {
			appcontext.ZLogger(ctx).Error("Error decoding role; role ObjectID was null")
			continue
		}

		if wireRole.RoleTypeID == nil {
			appcontext.ZLogger(ctx).Error("Error decoding role; role type ID was null")
			continue
		}

		// generated swagger client turns JSON nulls into Go zero values, so use null/zero package to convert them back to nullable values
		role := &models.CedarRole{
			Application: *wireRole.Application,
			ObjectID:    *wireRole.ObjectID,
			RoleTypeID:  *wireRole.RoleTypeID,

			AssigneeUserName:  zero.StringFrom(wireRole.AssigneeUserName),
			AssigneeEmail:     zero.StringFrom(wireRole.AssigneeEmail),
			AssigneeOrgID:     zero.StringFrom(wireRole.AssigneeOrgID),
			AssigneeOrgName:   zero.StringFrom(wireRole.AssigneeOrgName),
			AssigneeFirstName: zero.StringFrom(wireRole.AssigneeFirstName),
			AssigneeLastName:  zero.StringFrom(wireRole.AssigneeLastName),
			AssigneePhone:     zero.StringFrom(wireRole.AssigneePhone),
			AssigneeDesc:      zero.StringFrom(wireRole.AssigneeDesc),

			RoleTypeDesc: zero.StringFrom(wireRole.RoleTypeDesc),
			RoleID:       zero.StringFrom(wireRole.RoleID),
			ObjectType:   zero.StringFrom(wireRole.ObjectType),
		}

		var retAssigneeType models.CedarAssigneeType

		if wireRole.AssigneeType == string(models.PersonAssignee) {
			retAssigneeType = models.PersonAssignee
		} else if wireRole.AssigneeType == string(models.OrganizationAssignee) {
			retAssigneeType = models.OrganizationAssignee
		} else {
			retAssigneeType = models.UndefinedAssignee
		}

		role.AssigneeType = retAssigneeType

		roles = append(roles, role)
	}

	return roles
}
