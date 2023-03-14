package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody 
type ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The destinationFileName property
    destinationFileName *string
    // The sourceFile property
    sourceFile i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable
}
// NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody instantiates a new ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody and sets the default values.
func NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody()(*ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) {
    m := &ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestinationFileName gets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetDestinationFileName()(*string) {
    return m.destinationFileName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationFileName(val)
        }
        return nil
    }
    res["sourceFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceFile(val.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable))
        }
        return nil
    }
    return res
}
// GetSourceFile gets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) GetSourceFile()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable) {
    return m.sourceFile
}
// Serialize serializes information the current object
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationFileName", m.GetDestinationFileName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sourceFile", m.GetSourceFile())
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
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestinationFileName sets the destinationFileName property value. The destinationFileName property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetDestinationFileName(value *string)() {
    m.destinationFileName = value
}
// SetSourceFile sets the sourceFile property value. The sourceFile property
func (m *ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBody) SetSourceFile(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable)() {
    m.sourceFile = value
}
// ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable 
type ItemSitesItemContentTypesItemCopyToDefaultContentLocationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDestinationFileName()(*string)
    GetSourceFile()(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable)
    SetDestinationFileName(value *string)()
    SetSourceFile(value i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.ItemReferenceable)()
}
