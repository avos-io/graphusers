package groups

import (
    i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390 "github.com/avos-io/graphusers/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody 
type ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The commands property
    commands []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable
}
// NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody instantiates a new ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody and sets the default values.
func NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody()(*ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) {
    m := &ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommands gets the commands property value. The commands property
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetCommands()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable) {
    return m.commands
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.CreateOnenotePatchContentCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable, len(val))
            for i, v := range val {
                res[i] = v.(i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable)
            }
            m.SetCommands(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("commands", cast)
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
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommands sets the commands property value. The commands property
func (m *ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetCommands(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable)() {
    m.commands = value
}
// ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyable 
type ItemSitesItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommands()([]i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable)
    SetCommands(value []i43734bed85aefb0f6a3d313be76230963d1e26491f666899a105a0936ec1d390.OnenotePatchContentCommandable)()
}
