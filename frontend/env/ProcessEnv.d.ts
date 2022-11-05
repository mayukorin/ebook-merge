/// <reference types="node" />

declare namespace NodeJS {
    interface ProcessEnv {
        readonly API_HOST: string;
        readonly FIREBASE_API_KEY: string;
        readonly FIREBASE_AUTH_DOMAIN: string;
    }
}