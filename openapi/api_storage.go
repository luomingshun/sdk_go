/*
 * 声纹云api
 *
 * api document
 *
 * API version: /cloudapi/v1beta
 * Contact: xiachengjia@speakin.mobi
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"os"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type StorageApiService service

/*
StorageApiService
下载文件
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucket
 * @param key
@return *os.File
*/
func (a *StorageApiService) Download(ctx context.Context, bucket string, key string) (*os.File, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cloudapi/v1beta/storage/download"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("bucket", parameterToString(bucket, ""))
	localVarQueryParams.Add("key", parameterToString(key, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/octet-stream"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}


	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarHttpResponse.Body, localVarHttpResponse.Header.Get("Content-Type"));
		if err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v *os.File
			err = a.client.decode(&v, localVarHttpResponse.Body, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
StorageApiService
上传文件
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param bucket
 * @param type_
 * @param timestamp
 * @param optional nil or *UploadOpts - Optional Parameters:
 * @param "Body" (optional.Interface of *os.File) -  上传文件
 * @param "DurationMS" (optional.Int64) - 
@return CallUploadResponse
*/

type UploadOpts struct {
    Body optional.Interface
    DurationMS optional.Int64
}

func (a *StorageApiService) Upload(ctx context.Context, bucket string, type_ string, timestamp int64, localVarOptionals *UploadOpts) (CallUploadResponse, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue CallUploadResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/cloudapi/v1beta/storage/upload"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("bucket", parameterToString(bucket, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	if localVarOptionals != nil && localVarOptionals.DurationMS.IsSet() {
		localVarQueryParams.Add("durationMS", parameterToString(localVarOptionals.DurationMS.Value(), ""))
	}
	localVarQueryParams.Add("timestamp", parameterToString(timestamp, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream", "multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {
		localVarPostBody = localVarOptionals.Body.Value()
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}


	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarHttpResponse.Body, localVarHttpResponse.Header.Get("Content-Type"));
		if err != nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v CallUploadResponse
			err = a.client.decode(&v, localVarHttpResponse.Body, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
