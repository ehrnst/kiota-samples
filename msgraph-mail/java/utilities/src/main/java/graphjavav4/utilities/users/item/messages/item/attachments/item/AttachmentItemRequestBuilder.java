package graphjavav4.utilities.users.item.messages.item.attachments.item;

import com.microsoft.kiota.HttpMethod;
import com.microsoft.kiota.QueryParameter;
import com.microsoft.kiota.QueryParametersBase;
import com.microsoft.kiota.RequestAdapter;
import com.microsoft.kiota.RequestInformation;
import com.microsoft.kiota.RequestOption;
import com.microsoft.kiota.ResponseHandler;
import com.microsoft.kiota.serialization.Parsable;
import com.microsoft.kiota.serialization.ParsableFactory;
import graphjavav4.utilities.models.microsoft.graph.Attachment;
import java.net.URISyntaxException;
import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;
/** Builds and executes requests for operations under /users/{user-id}/messages/{message-id}/attachments/{attachment-id}  */
public class AttachmentItemRequestBuilder {
    /** Path parameters for the request  */
    private final HashMap<String, Object> pathParameters;
    /** The request adapter to use to execute the requests.  */
    private final RequestAdapter requestAdapter;
    /** Url template to use to build the URL for the current request builder  */
    private final String urlTemplate;
    /**
     * Instantiates a new AttachmentItemRequestBuilder and sets the default values.
     * @param pathParameters Path parameters for the request
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public AttachmentItemRequestBuilder(@javax.annotation.Nonnull final HashMap<String, Object> pathParameters, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        Objects.requireNonNull(pathParameters);
        Objects.requireNonNull(requestAdapter);
        this.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}{?%24select,%24expand}";
        var urlTplParams = new HashMap<String, Object>(pathParameters);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * Instantiates a new AttachmentItemRequestBuilder and sets the default values.
     * @param rawUrl The raw URL to use for the request builder.
     * @param requestAdapter The request adapter to use to execute the requests.
     * @return a void
     */
    public AttachmentItemRequestBuilder(@javax.annotation.Nonnull final String rawUrl, @javax.annotation.Nonnull final RequestAdapter requestAdapter) {
        this.urlTemplate = "{+baseurl}/users/{user%2Did}/messages/{message%2Did}/attachments/{attachment%2Did}{?%24select,%24expand}";
        var urlTplParams = new HashMap<String, Object>();
        urlTplParams.put("request-raw-url", rawUrl);
        this.pathParameters = urlTplParams;
        this.requestAdapter = requestAdapter;
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation() throws URISyntaxException {
        return createDeleteRequestInformation(null, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createDeleteRequestInformation(headers, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createDeleteRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.DELETE;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation() throws URISyntaxException {
        return createGetRequestInformation(null, null, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters) throws URISyntaxException {
        return createGetRequestInformation(queryParameters, null, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createGetRequestInformation(queryParameters, headers, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createGetRequestInformation(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.GET;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        if (queryParameters != null) {
            final GetQueryParameters qParams = new GetQueryParameters();
            queryParameters.accept(qParams);
            qParams.AddQueryParameters(requestInfo.queryParameters);
        }
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final Attachment body) throws URISyntaxException {
        return createPatchRequestInformation(body, null, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @param headers Request headers
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final Attachment body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) throws URISyntaxException {
        return createPatchRequestInformation(body, headers, null);
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @return a RequestInformation
     */
    @javax.annotation.Nonnull
    public RequestInformation createPatchRequestInformation(@javax.annotation.Nonnull final Attachment body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) throws URISyntaxException {
        Objects.requireNonNull(body);
        final RequestInformation requestInfo = new RequestInformation() {{
            httpMethod = HttpMethod.PATCH;
        }};
        requestInfo.urlTemplate = urlTemplate;
        requestInfo.pathParameters = pathParameters;
        requestInfo.setContentFromParsable(requestAdapter, "application/json", body);
        if (headers != null) {
            headers.accept(requestInfo.headers);
        }
        if (options != null) {
            requestInfo.addRequestOptions(options.toArray(new RequestOption[0]));
        }
        return requestInfo;
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete() {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(null, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> delete(@javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createDeleteRequestInformation(headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @return a CompletableFuture of attachment
     */
    public java.util.concurrent.CompletableFuture<Attachment> get() {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(null, null, null);
            return this.requestAdapter.sendAsync(requestInfo, Attachment::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of attachment
     */
    public java.util.concurrent.CompletableFuture<Attachment> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, null, null);
            return this.requestAdapter.sendAsync(requestInfo, Attachment::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of attachment
     */
    public java.util.concurrent.CompletableFuture<Attachment> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, null);
            return this.requestAdapter.sendAsync(requestInfo, Attachment::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @return a CompletableFuture of attachment
     */
    public java.util.concurrent.CompletableFuture<Attachment> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, options);
            return this.requestAdapter.sendAsync(requestInfo, Attachment::createFromDiscriminatorValue, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param headers Request headers
     * @param options Request options
     * @param queryParameters Request query parameters
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of attachment
     */
    public java.util.concurrent.CompletableFuture<Attachment> get(@javax.annotation.Nullable final java.util.function.Consumer<GetQueryParameters> queryParameters, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        try {
            final RequestInformation requestInfo = createGetRequestInformation(queryParameters, headers, options);
            return this.requestAdapter.sendAsync(requestInfo, Attachment::createFromDiscriminatorValue, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Attachment body) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, null, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @param headers Request headers
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Attachment body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, null);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Attachment body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options) {
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, null, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /**
     * The fileAttachment and itemAttachment attachments for the message.
     * @param body 
     * @param headers Request headers
     * @param options Request options
     * @param responseHandler Response handler to use in place of the default response handling provided by the core service
     * @return a CompletableFuture of void
     */
    public java.util.concurrent.CompletableFuture<Void> patch(@javax.annotation.Nonnull final Attachment body, @javax.annotation.Nullable final java.util.function.Consumer<Map<String, String>> headers, @javax.annotation.Nullable final Collection<RequestOption> options, @javax.annotation.Nullable final ResponseHandler responseHandler) {
        Objects.requireNonNull(body);
        try {
            final RequestInformation requestInfo = createPatchRequestInformation(body, headers, options);
            return this.requestAdapter.sendPrimitiveAsync(requestInfo, Void.class, responseHandler, null);
        } catch (URISyntaxException ex) {
            return java.util.concurrent.CompletableFuture.failedFuture(ex);
        }
    }
    /** The fileAttachment and itemAttachment attachments for the message.  */
    public class GetQueryParameters extends QueryParametersBase {
        /** Expand related entities  */
        @QueryParameter(name = "%24expand")
        @javax.annotation.Nullable
        public String[] expand;
        /** Select properties to be returned  */
        @QueryParameter(name = "%24select")
        @javax.annotation.Nullable
        public String[] select;
    }
}
