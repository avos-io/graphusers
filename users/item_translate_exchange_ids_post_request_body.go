package users

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemTranslateExchangeIdsPostRequestBody 
type ItemTranslateExchangeIdsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The InputIds property
    inputIds []string
    // The SourceIdType property
    sourceIdType *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat
    // The TargetIdType property
    targetIdType *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat
}
// NewItemTranslateExchangeIdsPostRequestBody instantiates a new ItemTranslateExchangeIdsPostRequestBody and sets the default values.
func NewItemTranslateExchangeIdsPostRequestBody()(*ItemTranslateExchangeIdsPostRequestBody) {
    m := &ItemTranslateExchangeIdsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemTranslateExchangeIdsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemTranslateExchangeIdsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTranslateExchangeIdsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTranslateExchangeIdsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemTranslateExchangeIdsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["inputIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetInputIds(res)
        }
        return nil
    }
    res["sourceIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdType(val.(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat))
        }
        return nil
    }
    res["targetIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ParseExchangeIdFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdType(val.(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat))
        }
        return nil
    }
    return res
}
// GetInputIds gets the inputIds property value. The InputIds property
func (m *ItemTranslateExchangeIdsPostRequestBody) GetInputIds()([]string) {
    return m.inputIds
}
// GetSourceIdType gets the sourceIdType property value. The SourceIdType property
func (m *ItemTranslateExchangeIdsPostRequestBody) GetSourceIdType()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat) {
    return m.sourceIdType
}
// GetTargetIdType gets the targetIdType property value. The TargetIdType property
func (m *ItemTranslateExchangeIdsPostRequestBody) GetTargetIdType()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat) {
    return m.targetIdType
}
// Serialize serializes information the current object
func (m *ItemTranslateExchangeIdsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetInputIds() != nil {
        err := writer.WriteCollectionOfStringValues("inputIds", m.GetInputIds())
        if err != nil {
            return err
        }
    }
    if m.GetSourceIdType() != nil {
        cast := (*m.GetSourceIdType()).String()
        err := writer.WriteStringValue("sourceIdType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetIdType() != nil {
        cast := (*m.GetTargetIdType()).String()
        err := writer.WriteStringValue("targetIdType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemTranslateExchangeIdsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetInputIds sets the inputIds property value. The InputIds property
func (m *ItemTranslateExchangeIdsPostRequestBody) SetInputIds(value []string)() {
    m.inputIds = value
}
// SetSourceIdType sets the sourceIdType property value. The SourceIdType property
func (m *ItemTranslateExchangeIdsPostRequestBody) SetSourceIdType(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)() {
    m.sourceIdType = value
}
// SetTargetIdType sets the targetIdType property value. The TargetIdType property
func (m *ItemTranslateExchangeIdsPostRequestBody) SetTargetIdType(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)() {
    m.targetIdType = value
}
// ItemTranslateExchangeIdsPostRequestBodyable 
type ItemTranslateExchangeIdsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInputIds()([]string)
    GetSourceIdType()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)
    GetTargetIdType()(*i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)
    SetInputIds(value []string)()
    SetSourceIdType(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)()
    SetTargetIdType(value *i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ExchangeIdFormat)()
}
