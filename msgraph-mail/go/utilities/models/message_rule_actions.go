package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

// MessageRuleActions 
type MessageRuleActions struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewMessageRuleActions instantiates a new messageRuleActions and sets the default values.
func NewMessageRuleActions()(*MessageRuleActions) {
    m := &MessageRuleActions{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any));
    return m
}
// CreateMessageRuleActionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRuleActionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageRuleActions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetAssignCategories gets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) GetAssignCategories()([]string) {
    val, err := m.GetBackingStore().Get("assignCategories")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *MessageRuleActions) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetCopyToFolder gets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) GetCopyToFolder()(*string) {
    val, err := m.GetBackingStore().Get("copyToFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDelete gets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) GetDelete()(*bool) {
    val, err := m.GetBackingStore().Get("delete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageRuleActions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAssignCategories(res)
        }
        return nil
    }
    res["copyToFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyToFolder(val)
        }
        return nil
    }
    res["delete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelete(val)
        }
        return nil
    }
    res["forwardAsAttachmentTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetForwardAsAttachmentTo(res)
        }
        return nil
    }
    res["forwardTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetForwardTo(res)
        }
        return nil
    }
    res["markAsRead"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkAsRead(val)
        }
        return nil
    }
    res["markImportance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMarkImportance(val.(*Importance))
        }
        return nil
    }
    res["moveToFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoveToFolder(val)
        }
        return nil
    }
    res["permanentDelete"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermanentDelete(val)
        }
        return nil
    }
    res["redirectTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetRedirectTo(res)
        }
        return nil
    }
    res["stopProcessingRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStopProcessingRules(val)
        }
        return nil
    }
    return res
}
// GetForwardAsAttachmentTo gets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) GetForwardAsAttachmentTo()([]Recipientable) {
    val, err := m.GetBackingStore().Get("forwardAsAttachmentTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Recipientable)
    }
    return nil
}
// GetForwardTo gets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) GetForwardTo()([]Recipientable) {
    val, err := m.GetBackingStore().Get("forwardTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Recipientable)
    }
    return nil
}
// GetMarkAsRead gets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) GetMarkAsRead()(*bool) {
    val, err := m.GetBackingStore().Get("markAsRead")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetMarkImportance gets the markImportance property value. 
func (m *MessageRuleActions) GetMarkImportance()(*Importance) {
    val, err := m.GetBackingStore().Get("markImportance")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*Importance)
    }
    return nil
}
// GetMoveToFolder gets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) GetMoveToFolder()(*string) {
    val, err := m.GetBackingStore().Get("moveToFolder")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetPermanentDelete gets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) GetPermanentDelete()(*bool) {
    val, err := m.GetBackingStore().Get("permanentDelete")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetRedirectTo gets the redirectTo property value. The email addresses to which a message should be redirected.
func (m *MessageRuleActions) GetRedirectTo()([]Recipientable) {
    val, err := m.GetBackingStore().Get("redirectTo")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Recipientable)
    }
    return nil
}
// GetStopProcessingRules gets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) GetStopProcessingRules()(*bool) {
    val, err := m.GetBackingStore().Get("stopProcessingRules")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// Serialize serializes information the current object
func (m *MessageRuleActions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAssignCategories() != nil {
        err := writer.WriteCollectionOfStringValues("assignCategories", m.GetAssignCategories())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("copyToFolder", m.GetCopyToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("delete", m.GetDelete())
        if err != nil {
            return err
        }
    }
    if m.GetForwardAsAttachmentTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardAsAttachmentTo()))
        for i, v := range m.GetForwardAsAttachmentTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("forwardAsAttachmentTo", cast)
        if err != nil {
            return err
        }
    }
    if m.GetForwardTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetForwardTo()))
        for i, v := range m.GetForwardTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("forwardTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("markAsRead", m.GetMarkAsRead())
        if err != nil {
            return err
        }
    }
    if m.GetMarkImportance() != nil {
        cast := (*m.GetMarkImportance()).String()
        err := writer.WriteStringValue("markImportance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("moveToFolder", m.GetMoveToFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("permanentDelete", m.GetPermanentDelete())
        if err != nil {
            return err
        }
    }
    if m.GetRedirectTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRedirectTo()))
        for i, v := range m.GetRedirectTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("redirectTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("stopProcessingRules", m.GetStopProcessingRules())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRuleActions) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetAssignCategories sets the assignCategories property value. A list of categories to be assigned to a message.
func (m *MessageRuleActions) SetAssignCategories(value []string)() {
    err := m.GetBackingStore().Set("assignCategories", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *MessageRuleActions) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    err := m.GetBackingStore().Set("backingStore", value)
    if err != nil {
        panic(err)
    }
}
// SetCopyToFolder sets the copyToFolder property value. The ID of a folder that a message is to be copied to.
func (m *MessageRuleActions) SetCopyToFolder(value *string)() {
    err := m.GetBackingStore().Set("copyToFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetDelete sets the delete property value. Indicates whether a message should be moved to the Deleted Items folder.
func (m *MessageRuleActions) SetDelete(value *bool)() {
    err := m.GetBackingStore().Set("delete", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardAsAttachmentTo sets the forwardAsAttachmentTo property value. The email addresses of the recipients to which a message should be forwarded as an attachment.
func (m *MessageRuleActions) SetForwardAsAttachmentTo(value []Recipientable)() {
    err := m.GetBackingStore().Set("forwardAsAttachmentTo", value)
    if err != nil {
        panic(err)
    }
}
// SetForwardTo sets the forwardTo property value. The email addresses of the recipients to which a message should be forwarded.
func (m *MessageRuleActions) SetForwardTo(value []Recipientable)() {
    err := m.GetBackingStore().Set("forwardTo", value)
    if err != nil {
        panic(err)
    }
}
// SetMarkAsRead sets the markAsRead property value. Indicates whether a message should be marked as read.
func (m *MessageRuleActions) SetMarkAsRead(value *bool)() {
    err := m.GetBackingStore().Set("markAsRead", value)
    if err != nil {
        panic(err)
    }
}
// SetMarkImportance sets the markImportance property value. 
func (m *MessageRuleActions) SetMarkImportance(value *Importance)() {
    err := m.GetBackingStore().Set("markImportance", value)
    if err != nil {
        panic(err)
    }
}
// SetMoveToFolder sets the moveToFolder property value. The ID of the folder that a message will be moved to.
func (m *MessageRuleActions) SetMoveToFolder(value *string)() {
    err := m.GetBackingStore().Set("moveToFolder", value)
    if err != nil {
        panic(err)
    }
}
// SetPermanentDelete sets the permanentDelete property value. Indicates whether a message should be permanently deleted and not saved to the Deleted Items folder.
func (m *MessageRuleActions) SetPermanentDelete(value *bool)() {
    err := m.GetBackingStore().Set("permanentDelete", value)
    if err != nil {
        panic(err)
    }
}
// SetRedirectTo sets the redirectTo property value. The email addresses to which a message should be redirected.
func (m *MessageRuleActions) SetRedirectTo(value []Recipientable)() {
    err := m.GetBackingStore().Set("redirectTo", value)
    if err != nil {
        panic(err)
    }
}
// SetStopProcessingRules sets the stopProcessingRules property value. Indicates whether subsequent rules should be evaluated.
func (m *MessageRuleActions) SetStopProcessingRules(value *bool)() {
    err := m.GetBackingStore().Set("stopProcessingRules", value)
    if err != nil {
        panic(err)
    }
}
