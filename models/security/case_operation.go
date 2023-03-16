package security

import (
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"

	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CaseOperation
type CaseOperation struct {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entity
	// The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
	action *CaseAction
	// The date and time the operation was completed.
	completedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The user that created the operation.
	createdBy i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable
	// The date and time the operation was created.
	createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The progress of the operation.
	percentProgress *int32
	// Contains success and failure-specific result information.
	resultInfo i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable
	// The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
	status *CaseOperationStatus
}

// NewCaseOperation instantiates a new caseOperation and sets the default values.
func NewCaseOperation() *CaseOperation {
	m := &CaseOperation{
		Entity: *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.NewEntity(),
	}
	return m
}

// CreateCaseOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCaseOperationFromDiscriminatorValue(
	parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode,
) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	if parseNode != nil {
		mappingValueNode, err := parseNode.GetChildNode("@odata.type")
		if err != nil {
			return nil, err
		}
		if mappingValueNode != nil {
			mappingValue, err := mappingValueNode.GetStringValue()
			if err != nil {
				return nil, err
			}
			if mappingValue != nil {
				switch *mappingValue {
				case "#microsoft.graph.security.ediscoveryAddToReviewSetOperation":
					return NewEdiscoveryAddToReviewSetOperation(), nil
				case "#microsoft.graph.security.ediscoveryEstimateOperation":
					return NewEdiscoveryEstimateOperation(), nil
				case "#microsoft.graph.security.ediscoveryHoldOperation":
					return NewEdiscoveryHoldOperation(), nil
				case "#microsoft.graph.security.ediscoveryIndexOperation":
					return NewEdiscoveryIndexOperation(), nil
				case "#microsoft.graph.security.ediscoveryPurgeDataOperation":
					return NewEdiscoveryPurgeDataOperation(), nil
				case "#microsoft.graph.security.ediscoveryTagOperation":
					return NewEdiscoveryTagOperation(), nil
				}
			}
		}
	}
	return NewCaseOperation(), nil
}

// GetAction gets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) GetAction() *CaseAction {
	return m.action
}

// GetCompletedDateTime gets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) GetCompletedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.completedDateTime
}

// GetCreatedBy gets the createdBy property value. The user that created the operation.
func (m *CaseOperation) GetCreatedBy() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable {
	return m.createdBy
}

// GetCreatedDateTime gets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) GetCreatedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.createdDateTime
}

// GetFieldDeserializers the deserialization information for the current model
func (m *CaseOperation) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := m.Entity.GetFieldDeserializers()
	res["action"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseCaseAction)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAction(val.(*CaseAction))
		}
		return nil
	}
	res["completedDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompletedDateTime(val)
		}
		return nil
	}
	res["createdBy"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(
			i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateIdentitySetFromDiscriminatorValue,
		)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedBy(
				val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable),
			)
		}
		return nil
	}
	res["createdDateTime"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedDateTime(val)
		}
		return nil
	}
	res["percentProgress"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPercentProgress(val)
		}
		return nil
	}
	res["resultInfo"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(
			i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateResultInfoFromDiscriminatorValue,
		)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResultInfo(
				val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable),
			)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetEnumValue(ParseCaseOperationStatus)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val.(*CaseOperationStatus))
		}
		return nil
	}
	return res
}

// GetPercentProgress gets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) GetPercentProgress() *int32 {
	return m.percentProgress
}

// GetResultInfo gets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) GetResultInfo() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable {
	return m.resultInfo
}

// GetStatus gets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) GetStatus() *CaseOperationStatus {
	return m.status
}

// Serialize serializes information the current object
func (m *CaseOperation) Serialize(
	writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter,
) error {
	err := m.Entity.Serialize(writer)
	if err != nil {
		return err
	}
	if m.GetAction() != nil {
		cast := (*m.GetAction()).String()
		err = writer.WriteStringValue("action", &cast)
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteTimeValue("completedDateTime", m.GetCompletedDateTime())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteInt32Value("percentProgress", m.GetPercentProgress())
		if err != nil {
			return err
		}
	}
	{
		err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
		if err != nil {
			return err
		}
	}
	if m.GetStatus() != nil {
		cast := (*m.GetStatus()).String()
		err = writer.WriteStringValue("status", &cast)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAction sets the action property value. The type of action the operation represents. Possible values are: addToReviewSet,applyTags,contentExport,convertToPdf,estimateStatistics, purgeData
func (m *CaseOperation) SetAction(value *CaseAction) {
	m.action = value
}

// SetCompletedDateTime sets the completedDateTime property value. The date and time the operation was completed.
func (m *CaseOperation) SetCompletedDateTime(
	value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
) {
	m.completedDateTime = value
}

// SetCreatedBy sets the createdBy property value. The user that created the operation.
func (m *CaseOperation) SetCreatedBy(
	value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable,
) {
	m.createdBy = value
}

// SetCreatedDateTime sets the createdDateTime property value. The date and time the operation was created.
func (m *CaseOperation) SetCreatedDateTime(
	value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
) {
	m.createdDateTime = value
}

// SetPercentProgress sets the percentProgress property value. The progress of the operation.
func (m *CaseOperation) SetPercentProgress(value *int32) {
	m.percentProgress = value
}

// SetResultInfo sets the resultInfo property value. Contains success and failure-specific result information.
func (m *CaseOperation) SetResultInfo(
	value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable,
) {
	m.resultInfo = value
}

// SetStatus sets the status property value. The status of the case operation. Possible values are: notStarted, submissionFailed, running, succeeded, partiallySucceeded, failed.
func (m *CaseOperation) SetStatus(value *CaseOperationStatus) {
	m.status = value
}

// CaseOperationable
type CaseOperationable interface {
	i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.Entityable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAction() *CaseAction
	GetCompletedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetCreatedBy() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable
	GetCreatedDateTime() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPercentProgress() *int32
	GetResultInfo() i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable
	GetStatus() *CaseOperationStatus
	SetAction(value *CaseAction)
	SetCompletedDateTime(
		value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
	)
	SetCreatedBy(
		value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable,
	)
	SetCreatedDateTime(
		value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time,
	)
	SetPercentProgress(value *int32)
	SetResultInfo(
		value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ResultInfoable,
	)
	SetStatus(value *CaseOperationStatus)
}
