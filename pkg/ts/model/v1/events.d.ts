import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";
export declare const protobufPackage = "model.v1";
export interface BlockEvent {
    blockHash: string;
    blockNumber: string;
    network: BlockEvent_Network | undefined;
    block: BlockEvent_Block | undefined;
}
export interface BlockEvent_Network {
    chainId: string;
}
export interface BlockEvent_Block {
    difficulty: string;
    extraData: string;
    gasLimit: string;
    gasUsed: string;
    hash: string;
    logsBloom: string;
    miner: string;
    mixHash: string;
    nonce: string;
    number: string;
    parentHash: string;
    receiptsRoot: string;
    sha3Uncles: string;
    size: string;
    stateRoot: string;
    timestamp: string;
    totalDifficulty: string;
    transactions: string[];
    transactionsRoot: string;
    uncles: string[];
    baseFeePerGas: string;
}
export interface TransactionEvent {
    transaction: TransactionEvent_Transaction | undefined;
    network: TransactionEvent_Network | undefined;
    addresses: {
        [key: string]: boolean;
    };
    block: TransactionEvent_Block | undefined;
    isContractDeployment: boolean;
    contractAddress: string;
    txAddresses: {
        [key: string]: boolean;
    };
    receipt: TransactionReceipt | undefined;
}
export interface TransactionEvent_Network {
    chainId: string;
}
export interface TransactionEvent_Block {
    blockHash: string;
    blockNumber: string;
    blockTimestamp: string;
    baseFeePerGas: string;
}
export interface TransactionEvent_Transaction {
    type: string;
    nonce: string;
    gasPrice: string;
    gas: string;
    value: string;
    input: string;
    v: string;
    r: string;
    s: string;
    to: string;
    hash: string;
    from: string;
    maxFeePerGas: string;
    maxPriorityFeePerGas: string;
}
export interface TransactionEvent_AddressesEntry {
    key: string;
    value: boolean;
}
export interface TransactionEvent_TxAddressesEntry {
    key: string;
    value: boolean;
}
export interface TransactionLog {
    address: string;
    topics: string[];
    data: string;
    blockNumber: string;
    transactionHash: string;
    transactionIndex: string;
    blockHash: string;
    logIndex: string;
    removed: boolean;
}
export interface TransactionReceipt {
    blobGasUsed?: number | undefined;
    blockHash: string;
    blockNumber: number;
    contractAddress?: string | undefined;
    cumulativeGasUsed: number;
    effectiveGasPrice: number;
    from: string;
    gasUsed: number;
    logs: TransactionLog[];
    logsBloom: string;
    root?: string | undefined;
    status: string;
    to: string;
    transactionHash: string;
    transactionIndex: number;
    type: string;
}
export declare const BlockEvent: MessageFns<BlockEvent>;
export declare const BlockEvent_Network: MessageFns<BlockEvent_Network>;
export declare const BlockEvent_Block: MessageFns<BlockEvent_Block>;
export declare const TransactionEvent: MessageFns<TransactionEvent>;
export declare const TransactionEvent_Network: MessageFns<TransactionEvent_Network>;
export declare const TransactionEvent_Block: MessageFns<TransactionEvent_Block>;
export declare const TransactionEvent_Transaction: MessageFns<TransactionEvent_Transaction>;
export declare const TransactionEvent_AddressesEntry: MessageFns<TransactionEvent_AddressesEntry>;
export declare const TransactionEvent_TxAddressesEntry: MessageFns<TransactionEvent_TxAddressesEntry>;
export declare const TransactionLog: MessageFns<TransactionLog>;
export declare const TransactionReceipt: MessageFns<TransactionReceipt>;
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
