# \VoiceprintApi

All URIs are relative to *https://vpc.speakin.mobi*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](VoiceprintApi.md#Delete) | **Post** /cloudapi/v1beta/voiceprint/delete | 
[**Query**](VoiceprintApi.md#Query) | **Post** /cloudapi/v1beta/voiceprint/query | 
[**Register**](VoiceprintApi.md#Register) | **Post** /cloudapi/v1beta/voiceprint/register | 
[**Threshold**](VoiceprintApi.md#Threshold) | **Post** /cloudapi/v1beta/voiceprint/threshold | 
[**Verify**](VoiceprintApi.md#Verify) | **Post** /cloudapi/v1beta/voiceprint/verify | 
[**Verify1ton**](VoiceprintApi.md#Verify1ton) | **Post** /cloudapi/v1beta/voiceprint/verify1ton | 
[**Verifytopn**](VoiceprintApi.md#Verifytopn) | **Post** /cloudapi/v1beta/voiceprint/verifytopn | 


# **Delete**
> RespVoiceprintDeleteResponse Delete(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintDeleteRequest** | [**optional.Interface of VoiceprintDeleteRequest**](VoiceprintDeleteRequest.md)|  | 

### Return type

[**RespVoiceprintDeleteResponse**](RespVoiceprintDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Query**
> RespVoiceprintQueryResponse Query(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***QueryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QueryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintQueryRequest** | [**optional.Interface of VoiceprintQueryRequest**](VoiceprintQueryRequest.md)|  | 

### Return type

[**RespVoiceprintQueryResponse**](RespVoiceprintQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Register**
> RespVoiceprintRegisterResponse Register(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegisterOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegisterOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintRegisterRequest** | [**optional.Interface of VoiceprintRegisterRequest**](VoiceprintRegisterRequest.md)|  | 

### Return type

[**RespVoiceprintRegisterResponse**](RespVoiceprintRegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Threshold**
> RespVoiceprintThresholdResponse Threshold(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ThresholdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ThresholdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintThresholdRequest** | [**optional.Interface of VoiceprintThresholdRequest**](VoiceprintThresholdRequest.md)|  | 

### Return type

[**RespVoiceprintThresholdResponse**](RespVoiceprintThresholdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verify**
> RespVoiceprintVerifyResponse Verify(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprintVerifyRequest** | [**optional.Interface of VoiceprintVerifyRequest**](VoiceprintVerifyRequest.md)|  | 

### Return type

[**RespVoiceprintVerifyResponse**](RespVoiceprintVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verify1ton**
> RespVoiceprint1tonVerifyResponse Verify1ton(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***Verify1tonOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Verify1tonOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprint1tonVerifyRequest** | [**optional.Interface of Voiceprint1tonVerifyRequest**](Voiceprint1tonVerifyRequest.md)|  | 

### Return type

[**RespVoiceprint1tonVerifyResponse**](RespVoiceprint1tonVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verifytopn**
> RespVoiceprinttopnVerifyResponse Verifytopn(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VerifytopnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifytopnOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **voiceprinttopnVerifyRequest** | [**optional.Interface of VoiceprinttopnVerifyRequest**](VoiceprinttopnVerifyRequest.md)|  | 

### Return type

[**RespVoiceprinttopnVerifyResponse**](RespVoiceprinttopnVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
