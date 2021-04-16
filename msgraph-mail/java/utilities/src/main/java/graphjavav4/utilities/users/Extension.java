package graphjavav4.utilities.users;

import java.util.Objects;
import com.microsoft.kiota.serialization.SerializationWriter;
import com.microsoft.kiota.serialization.ParseNode;
import com.microsoft.kiota.serialization.Parsable;
import java.util.function.BiConsumer;
import java.util.Map;
import java.util.HashMap;
public class Extension extends Entity implements Parsable {
    public void serialize(@javax.annotation.Nonnull final SerializationWriter writer) {
        Objects.requireNonNull(writer);
        super.serialize(writer);
    }
    @javax.annotation.Nonnull
    public <T> Map<String, BiConsumer<T, ParseNode>> getDeserializeFields() {
        final Map<String, BiConsumer<T, ParseNode>> fields = new HashMap<>(super.getDeserializeFields());
        return fields;
    }
}
