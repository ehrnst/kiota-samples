using Kiota.Abstractions;
using Kiota.Abstractions.Serialization;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties.Item;
namespace Graphdotnetv4.Users.MailFolders.Messages.MultiValueExtendedProperties {
    public class MultiValueExtendedPropertiesRequestBuilder {
        public MultiValueLegacyExtendedPropertyRequestBuilder this[string position] { get {
            return new MultiValueLegacyExtendedPropertyRequestBuilder { HttpCore = HttpCore, SerializerFactory = SerializerFactory, CurrentPath = CurrentPath + PathSegment  + "/" + position};
        } }
        public async Task<MultiValueExtendedPropertiesResponse> GetAsync(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreateGetRequestInfo(
                q, h
            );
            return await HttpCore.SendAsync<MultiValueExtendedPropertiesResponse>(requestInfo, responseHandler);
        }
        public RequestInfo CreateGetRequestInfo(Action<GetQueryParameters> q = default, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.GET,
                URI = new Uri(CurrentPath + PathSegment),
            };
            if (q != null) {
                var qParams = new GetQueryParameters();
                q.Invoke(qParams);
                qParams.AddQueryParameters(requestInfo.QueryParameters);
            }
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        public async Task<MultiValueLegacyExtendedProperty> PostAsync(MultiValueLegacyExtendedProperty body, Action<IDictionary<string, string>> h = default, IResponseHandler responseHandler = default) {
            var requestInfo = CreatePostRequestInfo(
                body, h
            );
            return await HttpCore.SendAsync<MultiValueLegacyExtendedProperty>(requestInfo, responseHandler);
        }
        public RequestInfo CreatePostRequestInfo(MultiValueLegacyExtendedProperty body, Action<IDictionary<string, string>> h = default) {
            var requestInfo = new RequestInfo {
                HttpMethod = HttpMethod.POST,
                URI = new Uri(CurrentPath + PathSegment),
            };
            requestInfo.SetJsonContentFromParsable(body, SerializerFactory);
            h?.Invoke(requestInfo.Headers);
            return requestInfo;
        }
        private string PathSegment { get; } = "/multiValueExtendedProperties";
        public string CurrentPath { get; set; }
        public IHttpCore HttpCore { get; set; }
        public Func<string, ISerializationWriter> SerializerFactory { get; set; }
        public class GetQueryParameters : QueryParametersBase {
            public int? Top { get; set; }
            public int? Skip { get; set; }
            public string Search { get; set; }
            public string Filter { get; set; }
            public bool? Count { get; set; }
            public string[] Orderby { get; set; }
            public string[] Select { get; set; }
            public string[] Expand { get; set; }
        }
    }
}
