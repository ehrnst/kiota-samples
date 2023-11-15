// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { type AttachmentCollectionResponse } from '../../../../../../../../../models/';
import { createAttachmentFromDiscriminatorValue, deserializeIntoAttachment, serializeAttachment, type Attachment } from '../../../../../../../../../models/attachment';
import { createAttachmentCollectionResponseFromDiscriminatorValue } from '../../../../../../../../../models/attachmentCollectionResponse';
import { type ODataError } from '../../../../../../../../../models/oDataErrors/';
import { createODataErrorFromDiscriminatorValue, deserializeIntoODataError, serializeODataError } from '../../../../../../../../../models/oDataErrors/oDataError';
import { CountRequestBuilder } from './count/countRequestBuilder';
import { AttachmentItemRequestBuilder } from './item/attachmentItemRequestBuilder';
import { BaseRequestBuilder, getPathParameters, HttpMethod, RequestInformation, type Parsable, type ParsableFactory, type RequestAdapter, type RequestConfiguration, type RequestOption } from '@microsoft/kiota-abstractions';

export interface AttachmentsRequestBuilderGetQueryParameters {
    /**
     * Include count of items
     */
    count?: boolean;
    /**
     * Expand related entities
     */
    expand?: string[];
    /**
     * Filter items by property values
     */
    filter?: string;
    /**
     * Order items by property values
     */
    orderby?: string[];
    /**
     * Select properties to be returned
     */
    select?: string[];
}
/**
 * Builds and executes requests for operations under /users/{user-id}/mailFolders/{mailFolder-id}/childFolders/{mailFolder-id1}/messages/{message-id}/attachments
 */
export class AttachmentsRequestBuilder extends BaseRequestBuilder {
    /**
     * The Count property
     */
    public get count(): CountRequestBuilder {
        return new CountRequestBuilder(this.pathParameters, this.requestAdapter);
    }
    /**
     * Gets an item from the graphtypescriptv4.utilities.users.item.mailFolders.item.childFolders.item.messages.item.attachments.item collection
     * @param attachmentId The unique identifier of attachment
     * @returns a AttachmentItemRequestBuilder
     */
    public byAttachmentId(attachmentId: string) : AttachmentItemRequestBuilder {
        if(!attachmentId) throw new Error("attachmentId cannot be undefined");
        const urlTplParams = getPathParameters(this.pathParameters);
        urlTplParams["attachment%2Did"] = attachmentId
        return new AttachmentItemRequestBuilder(urlTplParams, this.requestAdapter);
    };
    /**
     * Instantiates a new AttachmentsRequestBuilder and sets the default values.
     * @param pathParameters The raw url or the Url template parameters for the request.
     * @param requestAdapter The request adapter to use to execute the requests.
     */
    public constructor(pathParameters: Record<string, unknown> | string | undefined, requestAdapter: RequestAdapter) {
        super(pathParameters, requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/childFolders/{mailFolder%2Did1}/messages/{message%2Did}/attachments{?%24filter,%24count,%24orderby,%24select,%24expand}");
    };
    /**
     * Retrieve a list of attachment objects attached to a message. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of AttachmentCollectionResponse
     * @see {@link https://learn.microsoft.com/graph/api/message-list-attachments?view=graph-rest-1.0|Find more info here}
     */
    public get(requestConfiguration?: RequestConfiguration<AttachmentsRequestBuilderGetQueryParameters> | undefined) : Promise<AttachmentCollectionResponse | undefined> {
        const requestInfo = this.toGetRequestInformation(
            requestConfiguration
        );
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendAsync<AttachmentCollectionResponse>(requestInfo, createAttachmentCollectionResponseFromDiscriminatorValue, errorMapping);
    };
    /**
     * Use this API to add an attachment to a message. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. You can add an attachment to an existing message by posting to its attachments collection, or you canadd an attachment to a message that is being created and sent on the fly. This operation limits the size of the attachment you can add to under 3 MB. This API is available in the following national cloud deployments.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a Promise of Attachment
     * @see {@link https://learn.microsoft.com/graph/api/message-post-attachments?view=graph-rest-1.0|Find more info here}
     */
    public post(body: Attachment, requestConfiguration?: RequestConfiguration<object> | undefined) : Promise<Attachment | undefined> {
        const requestInfo = this.toPostRequestInformation(
            body, requestConfiguration
        );
        const errorMapping = {
            "4XX": createODataErrorFromDiscriminatorValue,
            "5XX": createODataErrorFromDiscriminatorValue,
        } as Record<string, ParsableFactory<Parsable>>;
        return this.requestAdapter.sendAsync<Attachment>(requestInfo, createAttachmentFromDiscriminatorValue, errorMapping);
    };
    /**
     * Retrieve a list of attachment objects attached to a message. This API is available in the following national cloud deployments.
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toGetRequestInformation(requestConfiguration?: RequestConfiguration<AttachmentsRequestBuilderGetQueryParameters> | undefined) : RequestInformation {
        const requestInfo = new RequestInformation(HttpMethod.GET, this.urlTemplate, this.pathParameters);
        requestInfo.configure(requestConfiguration, attachmentsRequestBuilderGetQueryParametersMapper);
        requestInfo.tryAddRequestHeaders("Accept", "application/json");
        return requestInfo;
    };
    /**
     * Use this API to add an attachment to a message. An attachment can be one of the following types: All these types of attachment resources are derived from the attachmentresource. You can add an attachment to an existing message by posting to its attachments collection, or you canadd an attachment to a message that is being created and sent on the fly. This operation limits the size of the attachment you can add to under 3 MB. This API is available in the following national cloud deployments.
     * @param body The request body
     * @param requestConfiguration Configuration for the request such as headers, query parameters, and middleware options.
     * @returns a RequestInformation
     */
    public toPostRequestInformation(body: Attachment, requestConfiguration?: RequestConfiguration<object> | undefined) : RequestInformation {
        if(!body) throw new Error("body cannot be undefined");
        const requestInfo = new RequestInformation(HttpMethod.POST, this.urlTemplate, this.pathParameters);
        requestInfo.configure(requestConfiguration, undefined);
        requestInfo.tryAddRequestHeaders("Accept", "application/json");
        requestInfo.setContentFromParsable(this.requestAdapter, "application/json", body, serializeAttachment);
        return requestInfo;
    };
    /**
     * Returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
     * @param rawUrl The raw URL to use for the request builder.
     * @returns a attachmentsRequestBuilder
     */
    public withUrl(rawUrl: string) : AttachmentsRequestBuilder {
        if(!rawUrl) throw new Error("rawUrl cannot be undefined");
        return new AttachmentsRequestBuilder(rawUrl, this.requestAdapter);
    };
}
const attachmentsRequestBuilderGetQueryParametersMapper: Record<string, string> = {
    "count": "%24count",
    "expand": "%24expand",
    "filter": "%24filter",
    "orderby": "%24orderby",
    "select": "%24select",
};
// tslint:enable
// eslint-enable
