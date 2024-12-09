// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.3.0
//   protoc               unknown
// source: model/v1/finding.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "model.v1";

export interface Finding {
  protocol: string;
  severity: Finding_Severity;
  metadata: { [key: string]: string };
  type: Finding_FindingType;
  alertId: string;
  name: string;
  description: string;
  timestamp: string;
  uniqueKey: string;
  routeInfo: RouteInfo | undefined;
}

export enum Finding_Severity {
  UNKNOWN = 0,
  INFO = 1,
  LOW = 2,
  MEDIUM = 3,
  HIGH = 4,
  CRITICAL = 5,
  UNRECOGNIZED = -1,
}

export function finding_SeverityFromJSON(object: any): Finding_Severity {
  switch (object) {
    case 0:
    case "UNKNOWN":
      return Finding_Severity.UNKNOWN;
    case 1:
    case "INFO":
      return Finding_Severity.INFO;
    case 2:
    case "LOW":
      return Finding_Severity.LOW;
    case 3:
    case "MEDIUM":
      return Finding_Severity.MEDIUM;
    case 4:
    case "HIGH":
      return Finding_Severity.HIGH;
    case 5:
    case "CRITICAL":
      return Finding_Severity.CRITICAL;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Finding_Severity.UNRECOGNIZED;
  }
}

export function finding_SeverityToJSON(object: Finding_Severity): string {
  switch (object) {
    case Finding_Severity.UNKNOWN:
      return "UNKNOWN";
    case Finding_Severity.INFO:
      return "INFO";
    case Finding_Severity.LOW:
      return "LOW";
    case Finding_Severity.MEDIUM:
      return "MEDIUM";
    case Finding_Severity.HIGH:
      return "HIGH";
    case Finding_Severity.CRITICAL:
      return "CRITICAL";
    case Finding_Severity.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum Finding_FindingType {
  UNKNOWN_TYPE = 0,
  EXPLOIT = 1,
  SUSPICIOUS = 2,
  DEGRADED = 3,
  INFORMATION = 4,
  SCAM = 5,
  UNRECOGNIZED = -1,
}

export function finding_FindingTypeFromJSON(object: any): Finding_FindingType {
  switch (object) {
    case 0:
    case "UNKNOWN_TYPE":
      return Finding_FindingType.UNKNOWN_TYPE;
    case 1:
    case "EXPLOIT":
      return Finding_FindingType.EXPLOIT;
    case 2:
    case "SUSPICIOUS":
      return Finding_FindingType.SUSPICIOUS;
    case 3:
    case "DEGRADED":
      return Finding_FindingType.DEGRADED;
    case 4:
    case "INFORMATION":
      return Finding_FindingType.INFORMATION;
    case 5:
    case "SCAM":
      return Finding_FindingType.SCAM;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Finding_FindingType.UNRECOGNIZED;
  }
}

export function finding_FindingTypeToJSON(object: Finding_FindingType): string {
  switch (object) {
    case Finding_FindingType.UNKNOWN_TYPE:
      return "UNKNOWN_TYPE";
    case Finding_FindingType.EXPLOIT:
      return "EXPLOIT";
    case Finding_FindingType.SUSPICIOUS:
      return "SUSPICIOUS";
    case Finding_FindingType.DEGRADED:
      return "DEGRADED";
    case Finding_FindingType.INFORMATION:
      return "INFORMATION";
    case Finding_FindingType.SCAM:
      return "SCAM";
    case Finding_FindingType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Finding_MetadataEntry {
  key: string;
  value: string;
}

export interface RouteInfo {
  userName: string;
  botName: string;
}

function createBaseFinding(): Finding {
  return {
    protocol: "",
    severity: 0,
    metadata: {},
    type: 0,
    alertId: "",
    name: "",
    description: "",
    timestamp: "",
    uniqueKey: "",
    routeInfo: undefined,
  };
}

export const Finding: MessageFns<Finding> = {
  encode(message: Finding, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.protocol !== "") {
      writer.uint32(10).string(message.protocol);
    }
    if (message.severity !== 0) {
      writer.uint32(16).int32(message.severity);
    }
    Object.entries(message.metadata).forEach(([key, value]) => {
      Finding_MetadataEntry.encode({ key: key as any, value }, writer.uint32(26).fork()).join();
    });
    if (message.type !== 0) {
      writer.uint32(32).int32(message.type);
    }
    if (message.alertId !== "") {
      writer.uint32(42).string(message.alertId);
    }
    if (message.name !== "") {
      writer.uint32(50).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(58).string(message.description);
    }
    if (message.timestamp !== "") {
      writer.uint32(66).string(message.timestamp);
    }
    if (message.uniqueKey !== "") {
      writer.uint32(74).string(message.uniqueKey);
    }
    if (message.routeInfo !== undefined) {
      RouteInfo.encode(message.routeInfo, writer.uint32(82).fork()).join();
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Finding {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinding();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.protocol = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 16) {
            break;
          }

          message.severity = reader.int32() as any;
          continue;
        }
        case 3: {
          if (tag !== 26) {
            break;
          }

          const entry3 = Finding_MetadataEntry.decode(reader, reader.uint32());
          if (entry3.value !== undefined) {
            message.metadata[entry3.key] = entry3.value;
          }
          continue;
        }
        case 4: {
          if (tag !== 32) {
            break;
          }

          message.type = reader.int32() as any;
          continue;
        }
        case 5: {
          if (tag !== 42) {
            break;
          }

          message.alertId = reader.string();
          continue;
        }
        case 6: {
          if (tag !== 50) {
            break;
          }

          message.name = reader.string();
          continue;
        }
        case 7: {
          if (tag !== 58) {
            break;
          }

          message.description = reader.string();
          continue;
        }
        case 8: {
          if (tag !== 66) {
            break;
          }

          message.timestamp = reader.string();
          continue;
        }
        case 9: {
          if (tag !== 74) {
            break;
          }

          message.uniqueKey = reader.string();
          continue;
        }
        case 10: {
          if (tag !== 82) {
            break;
          }

          message.routeInfo = RouteInfo.decode(reader, reader.uint32());
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Finding {
    return {
      protocol: isSet(object.protocol) ? gt.String(object.protocol) : "",
      severity: isSet(object.severity) ? finding_SeverityFromJSON(object.severity) : 0,
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
      type: isSet(object.type) ? finding_FindingTypeFromJSON(object.type) : 0,
      alertId: isSet(object.alertId) ? gt.String(object.alertId) : "",
      name: isSet(object.name) ? gt.String(object.name) : "",
      description: isSet(object.description) ? gt.String(object.description) : "",
      timestamp: isSet(object.timestamp) ? gt.String(object.timestamp) : "",
      uniqueKey: isSet(object.uniqueKey) ? gt.String(object.uniqueKey) : "",
      routeInfo: isSet(object.routeInfo) ? RouteInfo.fromJSON(object.routeInfo) : undefined,
    };
  },

  toJSON(message: Finding): unknown {
    const obj: any = {};
    if (message.protocol !== "") {
      obj.protocol = message.protocol;
    }
    if (message.severity !== 0) {
      obj.severity = finding_SeverityToJSON(message.severity);
    }
    if (message.metadata) {
      const entries = Object.entries(message.metadata);
      if (entries.length > 0) {
        obj.metadata = {};
        entries.forEach(([k, v]) => {
          obj.metadata[k] = v;
        });
      }
    }
    if (message.type !== 0) {
      obj.type = finding_FindingTypeToJSON(message.type);
    }
    if (message.alertId !== "") {
      obj.alertId = message.alertId;
    }
    if (message.name !== "") {
      obj.name = message.name;
    }
    if (message.description !== "") {
      obj.description = message.description;
    }
    if (message.timestamp !== "") {
      obj.timestamp = message.timestamp;
    }
    if (message.uniqueKey !== "") {
      obj.uniqueKey = message.uniqueKey;
    }
    if (message.routeInfo !== undefined) {
      obj.routeInfo = RouteInfo.toJSON(message.routeInfo);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Finding>, I>>(base?: I): Finding {
    return Finding.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Finding>, I>>(object: I): Finding {
    const message = createBaseFinding();
    message.protocol = object.protocol ?? "";
    message.severity = object.severity ?? 0;
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = gt.String(value);
      }
      return acc;
    }, {});
    message.type = object.type ?? 0;
    message.alertId = object.alertId ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.timestamp = object.timestamp ?? "";
    message.uniqueKey = object.uniqueKey ?? "";
    message.routeInfo = (object.routeInfo !== undefined && object.routeInfo !== null)
      ? RouteInfo.fromPartial(object.routeInfo)
      : undefined;
    return message;
  },
};

function createBaseFinding_MetadataEntry(): Finding_MetadataEntry {
  return { key: "", value: "" };
}

export const Finding_MetadataEntry: MessageFns<Finding_MetadataEntry> = {
  encode(message: Finding_MetadataEntry, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): Finding_MetadataEntry {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFinding_MetadataEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.key = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.value = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Finding_MetadataEntry {
    return {
      key: isSet(object.key) ? gt.String(object.key) : "",
      value: isSet(object.value) ? gt.String(object.value) : "",
    };
  },

  toJSON(message: Finding_MetadataEntry): unknown {
    const obj: any = {};
    if (message.key !== "") {
      obj.key = message.key;
    }
    if (message.value !== "") {
      obj.value = message.value;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Finding_MetadataEntry>, I>>(base?: I): Finding_MetadataEntry {
    return Finding_MetadataEntry.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Finding_MetadataEntry>, I>>(object: I): Finding_MetadataEntry {
    const message = createBaseFinding_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseRouteInfo(): RouteInfo {
  return { userName: "", botName: "" };
}

export const RouteInfo: MessageFns<RouteInfo> = {
  encode(message: RouteInfo, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.userName !== "") {
      writer.uint32(10).string(message.userName);
    }
    if (message.botName !== "") {
      writer.uint32(18).string(message.botName);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): RouteInfo {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRouteInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.userName = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.botName = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): RouteInfo {
    return {
      userName: isSet(object.userName) ? gt.String(object.userName) : "",
      botName: isSet(object.botName) ? gt.String(object.botName) : "",
    };
  },

  toJSON(message: RouteInfo): unknown {
    const obj: any = {};
    if (message.userName !== "") {
      obj.userName = message.userName;
    }
    if (message.botName !== "") {
      obj.botName = message.botName;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<RouteInfo>, I>>(base?: I): RouteInfo {
    return RouteInfo.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<RouteInfo>, I>>(object: I): RouteInfo {
    const message = createBaseRouteInfo();
    message.userName = object.userName ?? "";
    message.botName = object.botName ?? "";
    return message;
  },
};

declare const self: any | undefined;
declare const window: any | undefined;
declare const global: any | undefined;
const gt: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
