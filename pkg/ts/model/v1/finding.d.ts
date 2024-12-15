import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "model.v1";
export interface Finding {
    protocol: string;
    severity: Finding_Severity;
    metadata: {
        [key: string]: string;
    };
    type: Finding_FindingType;
    alertId: string;
    name: string;
    description: string;
    timestamp: string;
    uniqueKey: string;
    routeInfo: RouteInfo | undefined;
}
export declare enum Finding_Severity {
    UNKNOWN = 0,
    INFO = 1,
    LOW = 2,
    MEDIUM = 3,
    HIGH = 4,
    CRITICAL = 5,
    UNRECOGNIZED = -1
}
export declare function finding_SeverityFromJSON(object: any): Finding_Severity;
export declare function finding_SeverityToJSON(object: Finding_Severity): string;
export declare enum Finding_FindingType {
    UNKNOWN_TYPE = 0,
    EXPLOIT = 1,
    SUSPICIOUS = 2,
    DEGRADED = 3,
    INFORMATION = 4,
    SCAM = 5,
    UNRECOGNIZED = -1
}
export declare function finding_FindingTypeFromJSON(object: any): Finding_FindingType;
export declare function finding_FindingTypeToJSON(object: Finding_FindingType): string;
export interface Finding_MetadataEntry {
    key: string;
    value: string;
}
export interface RouteInfo {
    userName: string;
    botName: string;
}
export declare const Finding: MessageFns<Finding>;
export declare const Finding_MetadataEntry: MessageFns<Finding_MetadataEntry>;
export declare const RouteInfo: MessageFns<RouteInfo>;
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P : P & {
    [K in keyof P]: Exact<P[K], I[K]>;
} & {
    [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
};
export interface MessageFns<T> {
    encode(message: T, writer?: BinaryWriter): BinaryWriter;
    decode(input: BinaryReader | Uint8Array, length?: number): T;
    fromJSON(object: any): T;
    toJSON(message: T): unknown;
    create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
    fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
export {};
