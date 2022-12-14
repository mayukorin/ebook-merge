/* tslint:disable */
/* eslint-disable */
/**
 * e-book-merge
 * 電子書籍を一元管理
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    CreateGmailApiOAuth2TokenRequest,
    CreateGmailApiOAuth2TokenRequestFromJSON,
    CreateGmailApiOAuth2TokenRequestToJSON,
    CreateGmailApiOAuth2TokenResponse,
    CreateGmailApiOAuth2TokenResponseFromJSON,
    CreateGmailApiOAuth2TokenResponseToJSON,
    GmailApiConsentPageUrl,
    GmailApiConsentPageUrlFromJSON,
    GmailApiConsentPageUrlToJSON,
} from '../models';

export interface CreateGmailApiOauth2TokenRequest {
    body: CreateGmailApiOAuth2TokenRequest;
}

/**
 * 
 */
export class OAuth2TokenApi extends runtime.BaseAPI {

    /**
     * GmailAPIの同意画面のURLを生成
     * confirm-gmail-api
     */
    async confirmGmailApiRaw(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<GmailApiConsentPageUrl>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/generate-consent-page-url-of-gmail-api`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GmailApiConsentPageUrlFromJSON(jsonValue));
    }

    /**
     * GmailAPIの同意画面のURLを生成
     * confirm-gmail-api
     */
    async confirmGmailApi(initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<GmailApiConsentPageUrl> {
        const response = await this.confirmGmailApiRaw(initOverrides);
        return await response.value();
    }

    /**
     * GmailAPIのトークンを生成
     * create-gmail-api-oauth2-token
     */
    async createGmailApiOauth2TokenRaw(requestParameters: CreateGmailApiOauth2TokenRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<CreateGmailApiOAuth2TokenResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createGmailApiOauth2Token.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.apiKey) {
            headerParameters["Authorization"] = this.configuration.apiKey("Authorization"); // Authorization authentication
        }

        const response = await this.request({
            path: `/generate-oauth2-token-of-gmail-api`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateGmailApiOAuth2TokenRequestToJSON(requestParameters.body),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateGmailApiOAuth2TokenResponseFromJSON(jsonValue));
    }

    /**
     * GmailAPIのトークンを生成
     * create-gmail-api-oauth2-token
     */
    async createGmailApiOauth2Token(requestParameters: CreateGmailApiOauth2TokenRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<CreateGmailApiOAuth2TokenResponse> {
        const response = await this.createGmailApiOauth2TokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
