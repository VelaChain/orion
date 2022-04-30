/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { PoolShares, PoolAssets, PoolAsset } from "../chios/types";

export const protobufPackage = "VelaChain.orion.chios";

export interface MsgCreatePairPool {
  creator: string;
  denomA: string;
  amountA: string;
  denomB: string;
  amountB: string;
  sharesOut: string;
}

export interface MsgCreatePairPoolResponse {
  poolId: string;
  shares: PoolShares | undefined;
}

export interface MsgJoinPairPool {
  creator: string;
  denomA: string;
  amountA: string;
  denomB: string;
  amountB: string;
  sharesOut: string;
}

export interface MsgJoinPairPoolResponse {
  poolId: string;
  shares: PoolShares | undefined;
}

export interface MsgExitPairPool {
  creator: string;
  shareDenom: string;
}

export interface MsgExitPairPoolResponse {
  poolId: string;
  assets: PoolAssets | undefined;
}

export interface MsgSwapPair {
  creator: string;
  denomIn: string;
  amountIn: string;
  denomOut: string;
  minAmountOut: string;
}

export interface MsgSwapPairResponse {
  creator: string;
  assetOut: PoolAsset | undefined;
}

export interface MsgAddLiquidityPair {
  creator: string;
  denomA: string;
  amountA: string;
  denomB: string;
  amountB: string;
  sharesOut: string;
}

export interface MsgAddLiquidityPairResponse {
  poolId: string;
  shares: PoolShares | undefined;
}

export interface MsgRemoveLiquidityPair {
  creator: string;
  sharesDenom: string;
  sharesAmount: string;
}

export interface MsgRemoveLiquidityPairResponse {
  creator: string;
  assets: PoolAssets | undefined;
}

const baseMsgCreatePairPool: object = {
  creator: "",
  denomA: "",
  amountA: "",
  denomB: "",
  amountB: "",
  sharesOut: "",
};

export const MsgCreatePairPool = {
  encode(message: MsgCreatePairPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.amountA !== "") {
      writer.uint32(26).string(message.amountA);
    }
    if (message.denomB !== "") {
      writer.uint32(34).string(message.denomB);
    }
    if (message.amountB !== "") {
      writer.uint32(42).string(message.amountB);
    }
    if (message.sharesOut !== "") {
      writer.uint32(50).string(message.sharesOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePairPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePairPool } as MsgCreatePairPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.amountA = reader.string();
          break;
        case 4:
          message.denomB = reader.string();
          break;
        case 5:
          message.amountB = reader.string();
          break;
        case 6:
          message.sharesOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePairPool {
    const message = { ...baseMsgCreatePairPool } as MsgCreatePairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = String(object.sharesOut);
    } else {
      message.sharesOut = "";
    }
    return message;
  },

  toJSON(message: MsgCreatePairPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.sharesOut !== undefined && (obj.sharesOut = message.sharesOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePairPool>): MsgCreatePairPool {
    const message = { ...baseMsgCreatePairPool } as MsgCreatePairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = object.sharesOut;
    } else {
      message.sharesOut = "";
    }
    return message;
  },
};

const baseMsgCreatePairPoolResponse: object = { poolId: "" };

export const MsgCreatePairPoolResponse = {
  encode(
    message: MsgCreatePairPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolId !== "") {
      writer.uint32(10).string(message.poolId);
    }
    if (message.shares !== undefined) {
      PoolShares.encode(message.shares, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreatePairPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreatePairPoolResponse,
    } as MsgCreatePairPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = reader.string();
          break;
        case 2:
          message.shares = PoolShares.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePairPoolResponse {
    const message = {
      ...baseMsgCreatePairPoolResponse,
    } as MsgCreatePairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreatePairPoolResponse): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.shares !== undefined &&
      (obj.shares = message.shares
        ? PoolShares.toJSON(message.shares)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreatePairPoolResponse>
  ): MsgCreatePairPoolResponse {
    const message = {
      ...baseMsgCreatePairPoolResponse,
    } as MsgCreatePairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },
};

const baseMsgJoinPairPool: object = {
  creator: "",
  denomA: "",
  amountA: "",
  denomB: "",
  amountB: "",
  sharesOut: "",
};

export const MsgJoinPairPool = {
  encode(message: MsgJoinPairPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.amountA !== "") {
      writer.uint32(26).string(message.amountA);
    }
    if (message.denomB !== "") {
      writer.uint32(34).string(message.denomB);
    }
    if (message.amountB !== "") {
      writer.uint32(42).string(message.amountB);
    }
    if (message.sharesOut !== "") {
      writer.uint32(50).string(message.sharesOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinPairPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgJoinPairPool } as MsgJoinPairPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.amountA = reader.string();
          break;
        case 4:
          message.denomB = reader.string();
          break;
        case 5:
          message.amountB = reader.string();
          break;
        case 6:
          message.sharesOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinPairPool {
    const message = { ...baseMsgJoinPairPool } as MsgJoinPairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = String(object.sharesOut);
    } else {
      message.sharesOut = "";
    }
    return message;
  },

  toJSON(message: MsgJoinPairPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.sharesOut !== undefined && (obj.sharesOut = message.sharesOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgJoinPairPool>): MsgJoinPairPool {
    const message = { ...baseMsgJoinPairPool } as MsgJoinPairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = object.sharesOut;
    } else {
      message.sharesOut = "";
    }
    return message;
  },
};

const baseMsgJoinPairPoolResponse: object = { poolId: "" };

export const MsgJoinPairPoolResponse = {
  encode(
    message: MsgJoinPairPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolId !== "") {
      writer.uint32(10).string(message.poolId);
    }
    if (message.shares !== undefined) {
      PoolShares.encode(message.shares, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgJoinPairPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgJoinPairPoolResponse,
    } as MsgJoinPairPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = reader.string();
          break;
        case 2:
          message.shares = PoolShares.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgJoinPairPoolResponse {
    const message = {
      ...baseMsgJoinPairPoolResponse,
    } as MsgJoinPairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: MsgJoinPairPoolResponse): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.shares !== undefined &&
      (obj.shares = message.shares
        ? PoolShares.toJSON(message.shares)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgJoinPairPoolResponse>
  ): MsgJoinPairPoolResponse {
    const message = {
      ...baseMsgJoinPairPoolResponse,
    } as MsgJoinPairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },
};

const baseMsgExitPairPool: object = { creator: "", shareDenom: "" };

export const MsgExitPairPool = {
  encode(message: MsgExitPairPool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.shareDenom !== "") {
      writer.uint32(18).string(message.shareDenom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgExitPairPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgExitPairPool } as MsgExitPairPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.shareDenom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExitPairPool {
    const message = { ...baseMsgExitPairPool } as MsgExitPairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.shareDenom !== undefined && object.shareDenom !== null) {
      message.shareDenom = String(object.shareDenom);
    } else {
      message.shareDenom = "";
    }
    return message;
  },

  toJSON(message: MsgExitPairPool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.shareDenom !== undefined && (obj.shareDenom = message.shareDenom);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgExitPairPool>): MsgExitPairPool {
    const message = { ...baseMsgExitPairPool } as MsgExitPairPool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.shareDenom !== undefined && object.shareDenom !== null) {
      message.shareDenom = object.shareDenom;
    } else {
      message.shareDenom = "";
    }
    return message;
  },
};

const baseMsgExitPairPoolResponse: object = { poolId: "" };

export const MsgExitPairPoolResponse = {
  encode(
    message: MsgExitPairPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolId !== "") {
      writer.uint32(10).string(message.poolId);
    }
    if (message.assets !== undefined) {
      PoolAssets.encode(message.assets, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgExitPairPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgExitPairPoolResponse,
    } as MsgExitPairPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = reader.string();
          break;
        case 2:
          message.assets = PoolAssets.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgExitPairPoolResponse {
    const message = {
      ...baseMsgExitPairPoolResponse,
    } as MsgExitPairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    if (object.assets !== undefined && object.assets !== null) {
      message.assets = PoolAssets.fromJSON(object.assets);
    } else {
      message.assets = undefined;
    }
    return message;
  },

  toJSON(message: MsgExitPairPoolResponse): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.assets !== undefined &&
      (obj.assets = message.assets
        ? PoolAssets.toJSON(message.assets)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgExitPairPoolResponse>
  ): MsgExitPairPoolResponse {
    const message = {
      ...baseMsgExitPairPoolResponse,
    } as MsgExitPairPoolResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    if (object.assets !== undefined && object.assets !== null) {
      message.assets = PoolAssets.fromPartial(object.assets);
    } else {
      message.assets = undefined;
    }
    return message;
  },
};

const baseMsgSwapPair: object = {
  creator: "",
  denomIn: "",
  amountIn: "",
  denomOut: "",
  minAmountOut: "",
};

export const MsgSwapPair = {
  encode(message: MsgSwapPair, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomIn !== "") {
      writer.uint32(18).string(message.denomIn);
    }
    if (message.amountIn !== "") {
      writer.uint32(26).string(message.amountIn);
    }
    if (message.denomOut !== "") {
      writer.uint32(34).string(message.denomOut);
    }
    if (message.minAmountOut !== "") {
      writer.uint32(42).string(message.minAmountOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapPair {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapPair } as MsgSwapPair;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomIn = reader.string();
          break;
        case 3:
          message.amountIn = reader.string();
          break;
        case 4:
          message.denomOut = reader.string();
          break;
        case 5:
          message.minAmountOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapPair {
    const message = { ...baseMsgSwapPair } as MsgSwapPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomIn !== undefined && object.denomIn !== null) {
      message.denomIn = String(object.denomIn);
    } else {
      message.denomIn = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = String(object.amountIn);
    } else {
      message.amountIn = "";
    }
    if (object.denomOut !== undefined && object.denomOut !== null) {
      message.denomOut = String(object.denomOut);
    } else {
      message.denomOut = "";
    }
    if (object.minAmountOut !== undefined && object.minAmountOut !== null) {
      message.minAmountOut = String(object.minAmountOut);
    } else {
      message.minAmountOut = "";
    }
    return message;
  },

  toJSON(message: MsgSwapPair): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomIn !== undefined && (obj.denomIn = message.denomIn);
    message.amountIn !== undefined && (obj.amountIn = message.amountIn);
    message.denomOut !== undefined && (obj.denomOut = message.denomOut);
    message.minAmountOut !== undefined &&
      (obj.minAmountOut = message.minAmountOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwapPair>): MsgSwapPair {
    const message = { ...baseMsgSwapPair } as MsgSwapPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomIn !== undefined && object.denomIn !== null) {
      message.denomIn = object.denomIn;
    } else {
      message.denomIn = "";
    }
    if (object.amountIn !== undefined && object.amountIn !== null) {
      message.amountIn = object.amountIn;
    } else {
      message.amountIn = "";
    }
    if (object.denomOut !== undefined && object.denomOut !== null) {
      message.denomOut = object.denomOut;
    } else {
      message.denomOut = "";
    }
    if (object.minAmountOut !== undefined && object.minAmountOut !== null) {
      message.minAmountOut = object.minAmountOut;
    } else {
      message.minAmountOut = "";
    }
    return message;
  },
};

const baseMsgSwapPairResponse: object = { creator: "" };

export const MsgSwapPairResponse = {
  encode(
    message: MsgSwapPairResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.assetOut !== undefined) {
      PoolAsset.encode(message.assetOut, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgSwapPairResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgSwapPairResponse } as MsgSwapPairResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.assetOut = PoolAsset.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSwapPairResponse {
    const message = { ...baseMsgSwapPairResponse } as MsgSwapPairResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.assetOut !== undefined && object.assetOut !== null) {
      message.assetOut = PoolAsset.fromJSON(object.assetOut);
    } else {
      message.assetOut = undefined;
    }
    return message;
  },

  toJSON(message: MsgSwapPairResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.assetOut !== undefined &&
      (obj.assetOut = message.assetOut
        ? PoolAsset.toJSON(message.assetOut)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgSwapPairResponse>): MsgSwapPairResponse {
    const message = { ...baseMsgSwapPairResponse } as MsgSwapPairResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.assetOut !== undefined && object.assetOut !== null) {
      message.assetOut = PoolAsset.fromPartial(object.assetOut);
    } else {
      message.assetOut = undefined;
    }
    return message;
  },
};

const baseMsgAddLiquidityPair: object = {
  creator: "",
  denomA: "",
  amountA: "",
  denomB: "",
  amountB: "",
  sharesOut: "",
};

export const MsgAddLiquidityPair = {
  encode(
    message: MsgAddLiquidityPair,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.amountA !== "") {
      writer.uint32(26).string(message.amountA);
    }
    if (message.denomB !== "") {
      writer.uint32(34).string(message.denomB);
    }
    if (message.amountB !== "") {
      writer.uint32(42).string(message.amountB);
    }
    if (message.sharesOut !== "") {
      writer.uint32(50).string(message.sharesOut);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddLiquidityPair {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddLiquidityPair } as MsgAddLiquidityPair;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.amountA = reader.string();
          break;
        case 4:
          message.denomB = reader.string();
          break;
        case 5:
          message.amountB = reader.string();
          break;
        case 6:
          message.sharesOut = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquidityPair {
    const message = { ...baseMsgAddLiquidityPair } as MsgAddLiquidityPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = String(object.amountA);
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = String(object.amountB);
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = String(object.sharesOut);
    } else {
      message.sharesOut = "";
    }
    return message;
  },

  toJSON(message: MsgAddLiquidityPair): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.amountA !== undefined && (obj.amountA = message.amountA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.amountB !== undefined && (obj.amountB = message.amountB);
    message.sharesOut !== undefined && (obj.sharesOut = message.sharesOut);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddLiquidityPair>): MsgAddLiquidityPair {
    const message = { ...baseMsgAddLiquidityPair } as MsgAddLiquidityPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.amountA !== undefined && object.amountA !== null) {
      message.amountA = object.amountA;
    } else {
      message.amountA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.amountB !== undefined && object.amountB !== null) {
      message.amountB = object.amountB;
    } else {
      message.amountB = "";
    }
    if (object.sharesOut !== undefined && object.sharesOut !== null) {
      message.sharesOut = object.sharesOut;
    } else {
      message.sharesOut = "";
    }
    return message;
  },
};

const baseMsgAddLiquidityPairResponse: object = { poolId: "" };

export const MsgAddLiquidityPairResponse = {
  encode(
    message: MsgAddLiquidityPairResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.poolId !== "") {
      writer.uint32(10).string(message.poolId);
    }
    if (message.shares !== undefined) {
      PoolShares.encode(message.shares, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddLiquidityPairResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddLiquidityPairResponse,
    } as MsgAddLiquidityPairResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.poolId = reader.string();
          break;
        case 2:
          message.shares = PoolShares.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddLiquidityPairResponse {
    const message = {
      ...baseMsgAddLiquidityPairResponse,
    } as MsgAddLiquidityPairResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = String(object.poolId);
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromJSON(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },

  toJSON(message: MsgAddLiquidityPairResponse): unknown {
    const obj: any = {};
    message.poolId !== undefined && (obj.poolId = message.poolId);
    message.shares !== undefined &&
      (obj.shares = message.shares
        ? PoolShares.toJSON(message.shares)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddLiquidityPairResponse>
  ): MsgAddLiquidityPairResponse {
    const message = {
      ...baseMsgAddLiquidityPairResponse,
    } as MsgAddLiquidityPairResponse;
    if (object.poolId !== undefined && object.poolId !== null) {
      message.poolId = object.poolId;
    } else {
      message.poolId = "";
    }
    if (object.shares !== undefined && object.shares !== null) {
      message.shares = PoolShares.fromPartial(object.shares);
    } else {
      message.shares = undefined;
    }
    return message;
  },
};

const baseMsgRemoveLiquidityPair: object = {
  creator: "",
  sharesDenom: "",
  sharesAmount: "",
};

export const MsgRemoveLiquidityPair = {
  encode(
    message: MsgRemoveLiquidityPair,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.sharesDenom !== "") {
      writer.uint32(18).string(message.sharesDenom);
    }
    if (message.sharesAmount !== "") {
      writer.uint32(26).string(message.sharesAmount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveLiquidityPair {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveLiquidityPair } as MsgRemoveLiquidityPair;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.sharesDenom = reader.string();
          break;
        case 3:
          message.sharesAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveLiquidityPair {
    const message = { ...baseMsgRemoveLiquidityPair } as MsgRemoveLiquidityPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.sharesDenom !== undefined && object.sharesDenom !== null) {
      message.sharesDenom = String(object.sharesDenom);
    } else {
      message.sharesDenom = "";
    }
    if (object.sharesAmount !== undefined && object.sharesAmount !== null) {
      message.sharesAmount = String(object.sharesAmount);
    } else {
      message.sharesAmount = "";
    }
    return message;
  },

  toJSON(message: MsgRemoveLiquidityPair): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.sharesDenom !== undefined &&
      (obj.sharesDenom = message.sharesDenom);
    message.sharesAmount !== undefined &&
      (obj.sharesAmount = message.sharesAmount);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRemoveLiquidityPair>
  ): MsgRemoveLiquidityPair {
    const message = { ...baseMsgRemoveLiquidityPair } as MsgRemoveLiquidityPair;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.sharesDenom !== undefined && object.sharesDenom !== null) {
      message.sharesDenom = object.sharesDenom;
    } else {
      message.sharesDenom = "";
    }
    if (object.sharesAmount !== undefined && object.sharesAmount !== null) {
      message.sharesAmount = object.sharesAmount;
    } else {
      message.sharesAmount = "";
    }
    return message;
  },
};

const baseMsgRemoveLiquidityPairResponse: object = { creator: "" };

export const MsgRemoveLiquidityPairResponse = {
  encode(
    message: MsgRemoveLiquidityPairResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.assets !== undefined) {
      PoolAssets.encode(message.assets, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveLiquidityPairResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveLiquidityPairResponse,
    } as MsgRemoveLiquidityPairResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.assets = PoolAssets.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveLiquidityPairResponse {
    const message = {
      ...baseMsgRemoveLiquidityPairResponse,
    } as MsgRemoveLiquidityPairResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.assets !== undefined && object.assets !== null) {
      message.assets = PoolAssets.fromJSON(object.assets);
    } else {
      message.assets = undefined;
    }
    return message;
  },

  toJSON(message: MsgRemoveLiquidityPairResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.assets !== undefined &&
      (obj.assets = message.assets
        ? PoolAssets.toJSON(message.assets)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRemoveLiquidityPairResponse>
  ): MsgRemoveLiquidityPairResponse {
    const message = {
      ...baseMsgRemoveLiquidityPairResponse,
    } as MsgRemoveLiquidityPairResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.assets !== undefined && object.assets !== null) {
      message.assets = PoolAssets.fromPartial(object.assets);
    } else {
      message.assets = undefined;
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** CreatePairPool defines a method for creating a pool for two assets */
  CreatePairPool(
    request: MsgCreatePairPool
  ): Promise<MsgCreatePairPoolResponse>;
  /** JoinPairPool defines a method for joining an existing pool */
  JoinPairPool(request: MsgJoinPairPool): Promise<MsgJoinPairPoolResponse>;
  /** ExitPairPool defines a method for entirely leaving a pool */
  ExitPairPool(request: MsgExitPairPool): Promise<MsgExitPairPoolResponse>;
  /** SwapPair defines a method for swapping two assets using a PairPool */
  SwapPair(request: MsgSwapPair): Promise<MsgSwapPairResponse>;
  /** AddLiquidityPair defines a method for adding liquidity to a PairPool for an existing provider */
  AddLiquidityPair(
    request: MsgAddLiquidityPair
  ): Promise<MsgAddLiquidityPairResponse>;
  /** RemoveLiquidityPair defines a method for removing some liquidity from a PairPool but not all */
  RemoveLiquidityPair(
    request: MsgRemoveLiquidityPair
  ): Promise<MsgRemoveLiquidityPairResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePairPool(
    request: MsgCreatePairPool
  ): Promise<MsgCreatePairPoolResponse> {
    const data = MsgCreatePairPool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "CreatePairPool",
      data
    );
    return promise.then((data) =>
      MsgCreatePairPoolResponse.decode(new Reader(data))
    );
  }

  JoinPairPool(request: MsgJoinPairPool): Promise<MsgJoinPairPoolResponse> {
    const data = MsgJoinPairPool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "JoinPairPool",
      data
    );
    return promise.then((data) =>
      MsgJoinPairPoolResponse.decode(new Reader(data))
    );
  }

  ExitPairPool(request: MsgExitPairPool): Promise<MsgExitPairPoolResponse> {
    const data = MsgExitPairPool.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "ExitPairPool",
      data
    );
    return promise.then((data) =>
      MsgExitPairPoolResponse.decode(new Reader(data))
    );
  }

  SwapPair(request: MsgSwapPair): Promise<MsgSwapPairResponse> {
    const data = MsgSwapPair.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "SwapPair",
      data
    );
    return promise.then((data) => MsgSwapPairResponse.decode(new Reader(data)));
  }

  AddLiquidityPair(
    request: MsgAddLiquidityPair
  ): Promise<MsgAddLiquidityPairResponse> {
    const data = MsgAddLiquidityPair.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "AddLiquidityPair",
      data
    );
    return promise.then((data) =>
      MsgAddLiquidityPairResponse.decode(new Reader(data))
    );
  }

  RemoveLiquidityPair(
    request: MsgRemoveLiquidityPair
  ): Promise<MsgRemoveLiquidityPairResponse> {
    const data = MsgRemoveLiquidityPair.encode(request).finish();
    const promise = this.rpc.request(
      "VelaChain.orion.chios.Msg",
      "RemoveLiquidityPair",
      data
    );
    return promise.then((data) =>
      MsgRemoveLiquidityPairResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
