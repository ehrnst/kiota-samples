package graphjavav4.utilities.users.mailFolders.messages;

import java.util.Objects;
import graphjavav4.utilities.users.MultiValueLegacyExtendedProperty;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
import java.util.List;
public class MultiValueExtendedPropertiesResponse implements Parsable {
    @javax.annotation.Nullable
    public List<MultiValueLegacyExtendedProperty> value;
    @javax.annotation.Nullable
    public String nextLink;
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        writer.writeCollectionOfObjectValues("value", value);
        writer.writeStringValue("nextLink", nextLink);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(2);
        fields.put("value", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).value = n.getCollectionOfObjectValues(MultiValueLegacyExtendedProperty.class); });
        fields.put("nextLink", (o, n) -> { ((MultiValueExtendedPropertiesResponse)o).nextLink = n.getStringValue(); });
        return fields;
    }
}
