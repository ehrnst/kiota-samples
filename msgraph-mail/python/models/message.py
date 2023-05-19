from __future__ import annotations
from datetime import datetime
from kiota_abstractions.serialization import Parsable, ParseNode, SerializationWriter
from typing import Any, Callable, Dict, List, Optional, TYPE_CHECKING, Union

if TYPE_CHECKING:
    from . import attachment, extension, followup_flag, importance, inference_classification_type, internet_message_header, item_body, multi_value_legacy_extended_property, outlook_item, recipient, single_value_legacy_extended_property

from . import outlook_item

class Message(outlook_item.OutlookItem):
    # The fileAttachment and itemAttachment attachments for the message.
    attachments: Optional[List[attachment.Attachment]] = None
    # The Bcc: recipients for the message.
    bcc_recipients: Optional[List[recipient.Recipient]] = None
    # The body property
    body: Optional[item_body.ItemBody] = None
    # The first 255 characters of the message body. It is in text format.
    body_preview: Optional[str] = None
    # The Cc: recipients for the message.
    cc_recipients: Optional[List[recipient.Recipient]] = None
    # The ID of the conversation the email belongs to.
    conversation_id: Optional[str] = None
    # Indicates the position of the message within the conversation.
    conversation_index: Optional[bytes] = None
    # The collection of open extensions defined for the message. Nullable.
    extensions: Optional[List[extension.Extension]] = None
    # The flag property
    flag: Optional[followup_flag.FollowupFlag] = None
    # The from property
    from_: Optional[recipient.Recipient] = None
    # Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    has_attachments: Optional[bool] = None
    # The importance property
    importance: Optional[importance.Importance] = None
    # The inferenceClassification property
    inference_classification: Optional[inference_classification_type.InferenceClassificationType] = None
    # The internetMessageHeaders property
    internet_message_headers: Optional[List[internet_message_header.InternetMessageHeader]] = None
    # The internetMessageId property
    internet_message_id: Optional[str] = None
    # The isDeliveryReceiptRequested property
    is_delivery_receipt_requested: Optional[bool] = None
    # The isDraft property
    is_draft: Optional[bool] = None
    # The isRead property
    is_read: Optional[bool] = None
    # The isReadReceiptRequested property
    is_read_receipt_requested: Optional[bool] = None
    # The collection of multi-value extended properties defined for the message. Nullable.
    multi_value_extended_properties: Optional[List[multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty]] = None
    # The parentFolderId property
    parent_folder_id: Optional[str] = None
    # The receivedDateTime property
    received_date_time: Optional[datetime] = None
    # The replyTo property
    reply_to: Optional[List[recipient.Recipient]] = None
    # The sender property
    sender: Optional[recipient.Recipient] = None
    # The sentDateTime property
    sent_date_time: Optional[datetime] = None
    # The collection of single-value extended properties defined for the message. Nullable.
    single_value_extended_properties: Optional[List[single_value_legacy_extended_property.SingleValueLegacyExtendedProperty]] = None
    # The subject property
    subject: Optional[str] = None
    # The toRecipients property
    to_recipients: Optional[List[recipient.Recipient]] = None
    # The uniqueBody property
    unique_body: Optional[item_body.ItemBody] = None
    # The webLink property
    web_link: Optional[str] = None
    
    @staticmethod
    def create_from_discriminator_value(parse_node: Optional[ParseNode] = None) -> Message:
        """
        Creates a new instance of the appropriate class based on discriminator value
        Args:
            parseNode: The parse node to use to read the discriminator value and create the object
        Returns: Message
        """
        if parse_node is None:
            raise Exception("parse_node cannot be undefined")
        return Message()
    
    def get_field_deserializers(self,) -> Dict[str, Callable[[ParseNode], None]]:
        """
        The deserialization information for the current model
        Returns: Dict[str, Callable[[ParseNode], None]]
        """
        from . import attachment, extension, followup_flag, importance, inference_classification_type, internet_message_header, item_body, multi_value_legacy_extended_property, outlook_item, recipient, single_value_legacy_extended_property

        fields: Dict[str, Callable[[Any], None]] = {
            "attachments": lambda n : setattr(self, 'attachments', n.get_collection_of_object_values(attachment.Attachment)),
            "bccRecipients": lambda n : setattr(self, 'bcc_recipients', n.get_collection_of_object_values(recipient.Recipient)),
            "body": lambda n : setattr(self, 'body', n.get_object_value(item_body.ItemBody)),
            "bodyPreview": lambda n : setattr(self, 'body_preview', n.get_str_value()),
            "ccRecipients": lambda n : setattr(self, 'cc_recipients', n.get_collection_of_object_values(recipient.Recipient)),
            "conversationId": lambda n : setattr(self, 'conversation_id', n.get_str_value()),
            "conversationIndex": lambda n : setattr(self, 'conversation_index', n.get_bytes_value()),
            "extensions": lambda n : setattr(self, 'extensions', n.get_collection_of_object_values(extension.Extension)),
            "flag": lambda n : setattr(self, 'flag', n.get_object_value(followup_flag.FollowupFlag)),
            "from": lambda n : setattr(self, 'from_', n.get_object_value(recipient.Recipient)),
            "hasAttachments": lambda n : setattr(self, 'has_attachments', n.get_bool_value()),
            "importance": lambda n : setattr(self, 'importance', n.get_enum_value(importance.Importance)),
            "inferenceClassification": lambda n : setattr(self, 'inference_classification', n.get_enum_value(inference_classification_type.InferenceClassificationType)),
            "internetMessageHeaders": lambda n : setattr(self, 'internet_message_headers', n.get_collection_of_object_values(internet_message_header.InternetMessageHeader)),
            "internetMessageId": lambda n : setattr(self, 'internet_message_id', n.get_str_value()),
            "isDeliveryReceiptRequested": lambda n : setattr(self, 'is_delivery_receipt_requested', n.get_bool_value()),
            "isDraft": lambda n : setattr(self, 'is_draft', n.get_bool_value()),
            "isRead": lambda n : setattr(self, 'is_read', n.get_bool_value()),
            "isReadReceiptRequested": lambda n : setattr(self, 'is_read_receipt_requested', n.get_bool_value()),
            "multiValueExtendedProperties": lambda n : setattr(self, 'multi_value_extended_properties', n.get_collection_of_object_values(multi_value_legacy_extended_property.MultiValueLegacyExtendedProperty)),
            "parentFolderId": lambda n : setattr(self, 'parent_folder_id', n.get_str_value()),
            "receivedDateTime": lambda n : setattr(self, 'received_date_time', n.get_datetime_value()),
            "replyTo": lambda n : setattr(self, 'reply_to', n.get_collection_of_object_values(recipient.Recipient)),
            "sender": lambda n : setattr(self, 'sender', n.get_object_value(recipient.Recipient)),
            "sentDateTime": lambda n : setattr(self, 'sent_date_time', n.get_datetime_value()),
            "singleValueExtendedProperties": lambda n : setattr(self, 'single_value_extended_properties', n.get_collection_of_object_values(single_value_legacy_extended_property.SingleValueLegacyExtendedProperty)),
            "subject": lambda n : setattr(self, 'subject', n.get_str_value()),
            "toRecipients": lambda n : setattr(self, 'to_recipients', n.get_collection_of_object_values(recipient.Recipient)),
            "uniqueBody": lambda n : setattr(self, 'unique_body', n.get_object_value(item_body.ItemBody)),
            "webLink": lambda n : setattr(self, 'web_link', n.get_str_value()),
        }
        super_fields = super().get_field_deserializers()
        fields.update(super_fields)
        return fields
    
    def serialize(self,writer: SerializationWriter) -> None:
        """
        Serializes information the current object
        Args:
            writer: Serialization writer to use to serialize this model
        """
        if writer is None:
            raise Exception("writer cannot be undefined")
        super().serialize(writer)
        writer.write_collection_of_object_values("attachments", self.attachments)
        writer.write_collection_of_object_values("bccRecipients", self.bcc_recipients)
        writer.write_object_value("body", self.body)
        writer.write_str_value("bodyPreview", self.body_preview)
        writer.write_collection_of_object_values("ccRecipients", self.cc_recipients)
        writer.write_str_value("conversationId", self.conversation_id)
        writer.write_object_value("conversationIndex", self.conversation_index)
        writer.write_collection_of_object_values("extensions", self.extensions)
        writer.write_object_value("flag", self.flag)
        writer.write_object_value("from", self.from_)
        writer.write_bool_value("hasAttachments", self.has_attachments)
        writer.write_enum_value("importance", self.importance)
        writer.write_enum_value("inferenceClassification", self.inference_classification)
        writer.write_collection_of_object_values("internetMessageHeaders", self.internet_message_headers)
        writer.write_str_value("internetMessageId", self.internet_message_id)
        writer.write_bool_value("isDeliveryReceiptRequested", self.is_delivery_receipt_requested)
        writer.write_bool_value("isDraft", self.is_draft)
        writer.write_bool_value("isRead", self.is_read)
        writer.write_bool_value("isReadReceiptRequested", self.is_read_receipt_requested)
        writer.write_collection_of_object_values("multiValueExtendedProperties", self.multi_value_extended_properties)
        writer.write_str_value("parentFolderId", self.parent_folder_id)
        writer.write_datetime_value("receivedDateTime", self.received_date_time)
        writer.write_collection_of_object_values("replyTo", self.reply_to)
        writer.write_object_value("sender", self.sender)
        writer.write_datetime_value("sentDateTime", self.sent_date_time)
        writer.write_collection_of_object_values("singleValueExtendedProperties", self.single_value_extended_properties)
        writer.write_str_value("subject", self.subject)
        writer.write_collection_of_object_values("toRecipients", self.to_recipients)
        writer.write_object_value("uniqueBody", self.unique_body)
        writer.write_str_value("webLink", self.web_link)
    

