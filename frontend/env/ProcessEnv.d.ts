/// <reference types="node" />

declare namespace NodeJS {
    interface ProcessEnv {
        readonly API_HOST: string;
    }
}