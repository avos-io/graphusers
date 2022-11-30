package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81CompliancePolicy 
type Windows81CompliancePolicy struct {
    DeviceCompliancePolicy
    // Maximum Windows 8.1 version.
    osMaximumVersion *string
    // Minimum Windows 8.1 version.
    osMinimumVersion *string
    // Indicates whether or not to block simple password.
    passwordBlockSimple *bool
    // Password expiration in days.
    passwordExpirationDays *int32
    // The number of character sets required in the password.
    passwordMinimumCharacterSetCount *int32
    // The minimum password length.
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // The number of previous passwords to prevent re-use of. Valid values 0 to 24
    passwordPreviousPasswordBlockCount *int32
    // Require a password to unlock Windows device.
    passwordRequired *bool
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // Indicates whether or not to require encryption on a windows 8.1 device.
    storageRequireEncryption *bool
}
// NewWindows81CompliancePolicy instantiates a new Windows81CompliancePolicy and sets the default values.
func NewWindows81CompliancePolicy()(*Windows81CompliancePolicy) {
    m := &Windows81CompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.windows81CompliancePolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateWindows81CompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81CompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81CompliancePolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81CompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["osMaximumVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsMaximumVersion)
    res["osMinimumVersion"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetOsMinimumVersion)
    res["passwordBlockSimple"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordBlockSimple)
    res["passwordExpirationDays"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordExpirationDays)
    res["passwordMinimumCharacterSetCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinimumCharacterSetCount)
    res["passwordMinimumLength"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinimumLength)
    res["passwordMinutesOfInactivityBeforeLock"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordMinutesOfInactivityBeforeLock)
    res["passwordPreviousPasswordBlockCount"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetInt32Value(m.SetPasswordPreviousPasswordBlockCount)
    res["passwordRequired"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetPasswordRequired)
    res["passwordRequiredType"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseRequiredPasswordType , m.SetPasswordRequiredType)
    res["storageRequireEncryption"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetBoolValue(m.SetStorageRequireEncryption)
    return res
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Windows 8.1 version.
func (m *Windows81CompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Windows 8.1 version.
func (m *Windows81CompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Indicates whether or not to block simple password.
func (m *Windows81CompliancePolicy) GetPasswordBlockSimple()(*bool) {
    return m.passwordBlockSimple
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Password expiration in days.
func (m *Windows81CompliancePolicy) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows81CompliancePolicy) GetPasswordMinimumCharacterSetCount()(*int32) {
    return m.passwordMinimumCharacterSetCount
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. The minimum password length.
func (m *Windows81CompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *Windows81CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of. Valid values 0 to 24
func (m *Windows81CompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock Windows device.
func (m *Windows81CompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows81CompliancePolicy) GetPasswordRequiredType()(*RequiredPasswordType) {
    return m.passwordRequiredType
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Indicates whether or not to require encryption on a windows 8.1 device.
func (m *Windows81CompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *Windows81CompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockSimple", m.GetPasswordBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumCharacterSetCount", m.GetPasswordMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Windows 8.1 version.
func (m *Windows81CompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Windows 8.1 version.
func (m *Windows81CompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Indicates whether or not to block simple password.
func (m *Windows81CompliancePolicy) SetPasswordBlockSimple(value *bool)() {
    m.passwordBlockSimple = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Password expiration in days.
func (m *Windows81CompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows81CompliancePolicy) SetPasswordMinimumCharacterSetCount(value *int32)() {
    m.passwordMinimumCharacterSetCount = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. The minimum password length.
func (m *Windows81CompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *Windows81CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of. Valid values 0 to 24
func (m *Windows81CompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock Windows device.
func (m *Windows81CompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows81CompliancePolicy) SetPasswordRequiredType(value *RequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Indicates whether or not to require encryption on a windows 8.1 device.
func (m *Windows81CompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
