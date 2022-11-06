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

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface GmailApiConsentPageUrl
 */
export interface GmailApiConsentPageUrl {
    /**
     * 
     * @type {string}
     * @memberof GmailApiConsentPageUrl
     */
    googleConcentPageUrl: string;
}

export function GmailApiConsentPageUrlFromJSON(json: any): GmailApiConsentPageUrl {
    return GmailApiConsentPageUrlFromJSONTyped(json, false);
}

export function GmailApiConsentPageUrlFromJSONTyped(json: any, ignoreDiscriminator: boolean): GmailApiConsentPageUrl {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'googleConcentPageUrl': json['google_concent_page_url'],
    };
}

export function GmailApiConsentPageUrlToJSON(value?: GmailApiConsentPageUrl | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'google_concent_page_url': value.googleConcentPageUrl,
    };
}
