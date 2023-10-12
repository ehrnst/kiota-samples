package graphjavav4.utilities.models;

import com.microsoft.kiota.serialization.AdditionalDataHolder;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.SerializationWriter;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
@jakarta.annotation.Generated("com.microsoft.kiota")
public class InferenceClassificationOverrideCollectionResponse implements AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    private Map<String, Object> additionalData;
    /**
     * The OdataNextLink property
     */
    private String odataNextLink;
    /**
     * The value property
     */
    private java.util.List<InferenceClassificationOverride> value;
    /**
     * Instantiates a new InferenceClassificationOverrideCollectionResponse and sets the default values.
     */
    public InferenceClassificationOverrideCollectionResponse() {
        this.setAdditionalData(new HashMap<>());
    }
    /**
     * Creates a new instance of the appropriate class based on discriminator value
     * @param parseNode The parse node to use to read the discriminator value and create the object
     * @return a InferenceClassificationOverrideCollectionResponse
     */
    @jakarta.annotation.Nonnull
    public static InferenceClassificationOverrideCollectionResponse createFromDiscriminatorValue(@jakarta.annotation.Nonnull final ParseNode parseNode) {
        Objects.requireNonNull(parseNode);
        return new InferenceClassificationOverrideCollectionResponse();
    }
    /**
     * Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @return a Map<String, Object>
     */
    @jakarta.annotation.Nonnull
    public Map<String, Object> getAdditionalData() {
        return this.additionalData;
    }
    /**
     * The deserialization information for the current model
     * @return a Map<String, java.util.function.Consumer<ParseNode>>
     */
    @jakarta.annotation.Nonnull
    public Map<String, java.util.function.Consumer<ParseNode>> getFieldDeserializers() {
        final HashMap<String, java.util.function.Consumer<ParseNode>> deserializerMap = new HashMap<String, java.util.function.Consumer<ParseNode>>(2);
        deserializerMap.put("@odata.nextLink", (n) -> { this.setOdataNextLink(n.getStringValue()); });
        deserializerMap.put("value", (n) -> { this.setValue(n.getCollectionOfObjectValues(InferenceClassificationOverride::createFromDiscriminatorValue)); });
        return deserializerMap;
    }
    /**
     * Gets the @odata.nextLink property value. The OdataNextLink property
     * @return a String
     */
    @jakarta.annotation.Nullable
    public String getOdataNextLink() {
        return this.odataNextLink;
    }
    /**
     * Gets the value property value. The value property
     * @return a java.util.List<InferenceClassificationOverride>
     */
    @jakarta.annotation.Nullable
    public java.util.List<InferenceClassificationOverride> getValue() {
        return this.value;
    }
    /**
     * Serializes information the current object
     * @param writer Serialization writer to use to serialize this model
     */
    public void serialize(@jakarta.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeStringValue("@odata.nextLink", this.getOdataNextLink());
        writer.writeCollectionOfObjectValues("value", this.getValue());
        writer.writeAdditionalData(this.getAdditionalData());
    }
    /**
     * Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     * @param value Value to set for the AdditionalData property.
     */
    public void setAdditionalData(@jakarta.annotation.Nullable final Map<String, Object> value) {
        this.additionalData = value;
    }
    /**
     * Sets the @odata.nextLink property value. The OdataNextLink property
     * @param value Value to set for the @odata.nextLink property.
     */
    public void setOdataNextLink(@jakarta.annotation.Nullable final String value) {
        this.odataNextLink = value;
    }
    /**
     * Sets the value property value. The value property
     * @param value Value to set for the value property.
     */
    public void setValue(@jakarta.annotation.Nullable final java.util.List<InferenceClassificationOverride> value) {
        this.value = value;
    }
}
