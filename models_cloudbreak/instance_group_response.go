package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*InstanceGroupResponse instance group response

swagger:model InstanceGroupResponse
*/
type InstanceGroupResponse struct {

	/* name of the instance group

	Required: true
	*/
	Group string `json:"group"`

	/* id of the resource

	Read Only: true
	*/
	ID *int64 `json:"id,omitempty"`

	/* metadata of instances

	Read Only: true
	Unique: true
	*/
	Metadata []*InstanceMetaData `json:"metadata,omitempty"`

	/* number of nodes

	Required: true
	Maximum: 100000
	Minimum: 0
	*/
	NodeCount int32 `json:"nodeCount"`

	/* cloud specific parameters for instance group
	 */
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	/* instancegroup related securitygroup
	 */
	SecurityGroup *SecurityGroupResponse `json:"securityGroup,omitempty"`

	/* security group resource id for the instance group
	 */
	SecurityGroupID *int64 `json:"securityGroupId,omitempty"`

	/* instancegroup related template
	 */
	Template *TemplateResponse `json:"template,omitempty"`

	/* referenced template id
	 */
	TemplateID *int64 `json:"templateId,omitempty"`

	/* type of the instance group
	 */
	Type *string `json:"type,omitempty"`
}

// Validate validates this instance group response
func (m *InstanceGroupResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupResponse) validateGroup(formats strfmt.Registry) error {

	if err := validate.RequiredString("group", "body", string(m.Group)); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupResponse) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.UniqueItems("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {

		if m.Metadata[i] != nil {

			if err := m.Metadata[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupResponse) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", int32(m.NodeCount)); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

var instanceGroupResponseTypeTypePropEnum []interface{}

func (m *InstanceGroupResponse) validateTypeEnum(path, location string, value string) error {
	if instanceGroupResponseTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["GATEWAY","CORE"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			instanceGroupResponseTypeTypePropEnum = append(instanceGroupResponseTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, instanceGroupResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
