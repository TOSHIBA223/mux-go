# \ExportsApi

All URIs are relative to *https://api.mux.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListExports**](ExportsApi.md#ListExports) | **Get** /data/v1/exports | List property video view export links
[**ListExportsViews**](ExportsApi.md#ListExportsViews) | **Get** /data/v1/exports/views | List available property view exports


# **ListExports**
> ListExportsResponse ListExports(ctx, )
List property video view export links

Deprecated: The API has been replaced by the list-exports-views API call.  Lists the available video view exports along with URLs to retrieve them. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListExportsResponse**](ListExportsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListExportsViews**
> ListVideoViewExportsResponse ListExportsViews(ctx, )
List available property view exports

Lists the available video view exports along with URLs to retrieve them.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListVideoViewExportsResponse**](ListVideoViewExportsResponse.md)

### Authorization

[accessToken](../README.md#accessToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

