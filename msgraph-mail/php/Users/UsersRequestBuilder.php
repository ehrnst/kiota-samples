<?php

namespace Microsoft\Graph\Users;

use Microsoft\Graph\Users\Item\UserRequestBuilder;
use Microsoft\Kiota\Abstractions\RequestAdapter;

class UsersRequestBuilder 
{
    /** @var array $pathParameters Path parameters for the request */
    private array $pathParameters;
    
    /** @var IRequestAdapter $requestAdapter The request adapter to use to execute the requests. */
    private RequestAdapter $requestAdapter;
    
    /** @var string $urlTemplate Url template to use to build the URL for the current request builder */
    private string $urlTemplate;
    
    /**
     * Instantiates a new UsersRequestBuilder and sets the default values.
     * @param array $pathParameters Path parameters for the request
     * @param RequestAdapter $requestAdapter The request adapter to use to execute the requests.
    */
    public function __construct(array $pathParameters, RequestAdapter $requestAdapter) {
        if (is_null($pathParameters)) {
            throw new \Exception('$pathParameters cannot be null');
        }
        if (is_null($requestAdapter)) {
            throw new \Exception('$requestAdapter cannot be null');
        }
        $this->urlTemplate = 'https://graph.microsoft.com/v1.0/users';
        $this->requestAdapter = $requestAdapter;
        $this->pathParameters = $pathParameters;
    }

}
