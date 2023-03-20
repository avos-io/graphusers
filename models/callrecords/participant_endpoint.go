package callrecords

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ParticipantEndpoint 
type ParticipantEndpoint struct {
    Endpoint
    // The feedback provided by the user of this endpoint about the quality of the session.
    feedback UserFeedbackable
    // Identity associated with the endpoint.
    identity i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable
}
// NewParticipantEndpoint instantiates a new ParticipantEndpoint and sets the default values.
func NewParticipantEndpoint()(*ParticipantEndpoint) {
    m := &ParticipantEndpoint{
        Endpoint: *NewEndpoint(),
    }
    odataTypeValue := "#microsoft.graph.callRecords.participantEndpoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateParticipantEndpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateParticipantEndpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewParticipantEndpoint(), nil
}
// GetFeedback gets the feedback property value. The feedback provided by the user of this endpoint about the quality of the session.
func (m *ParticipantEndpoint) GetFeedback()(UserFeedbackable) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ParticipantEndpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Endpoint.GetFieldDeserializers()
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFeedbackFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(UserFeedbackable))
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable))
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity associated with the endpoint.
func (m *ParticipantEndpoint) GetIdentity()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable) {
    return m.identity
}
// Serialize serializes information the current object
func (m *ParticipantEndpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Endpoint.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("feedback", m.GetFeedback())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFeedback sets the feedback property value. The feedback provided by the user of this endpoint about the quality of the session.
func (m *ParticipantEndpoint) SetFeedback(value UserFeedbackable)() {
    m.feedback = value
}
// SetIdentity sets the identity property value. Identity associated with the endpoint.
func (m *ParticipantEndpoint) SetIdentity(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable)() {
    m.identity = value
}
// ParticipantEndpointable 
type ParticipantEndpointable interface {
    Endpointable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFeedback()(UserFeedbackable)
    GetIdentity()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable)
    SetFeedback(value UserFeedbackable)()
    SetIdentity(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.IdentitySetable)()
}
