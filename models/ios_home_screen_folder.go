package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosHomeScreenFolder 
type IosHomeScreenFolder struct {
    IosHomeScreenItem
    // Pages of Home Screen Layout Icons which must be applications or web clips. This collection can contain a maximum of 500 elements.
    pages []IosHomeScreenFolderPageable
}
// NewIosHomeScreenFolder instantiates a new IosHomeScreenFolder and sets the default values.
func NewIosHomeScreenFolder()(*IosHomeScreenFolder) {
    m := &IosHomeScreenFolder{
        IosHomeScreenItem: *NewIosHomeScreenItem(),
    }
    odataTypeValue := "#microsoft.graph.iosHomeScreenFolder";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateIosHomeScreenFolderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosHomeScreenFolderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosHomeScreenFolder(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosHomeScreenFolder) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IosHomeScreenItem.GetFieldDeserializers()
    res["pages"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateIosHomeScreenFolderPageFromDiscriminatorValue , m.SetPages)
    return res
}
// GetPages gets the pages property value. Pages of Home Screen Layout Icons which must be applications or web clips. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenFolder) GetPages()([]IosHomeScreenFolderPageable) {
    return m.pages
}
// Serialize serializes information the current object
func (m *IosHomeScreenFolder) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IosHomeScreenItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPages() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetPages())
        err = writer.WriteCollectionOfObjectValues("pages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPages sets the pages property value. Pages of Home Screen Layout Icons which must be applications or web clips. This collection can contain a maximum of 500 elements.
func (m *IosHomeScreenFolder) SetPages(value []IosHomeScreenFolderPageable)() {
    m.pages = value
}
