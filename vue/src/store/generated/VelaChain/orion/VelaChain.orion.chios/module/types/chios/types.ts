/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.orion.chios";

export interface LiquidityProvider {
  /** address of provider */
  creator: string;
  /** list of pool shares owned by provider */
  liquid: PoolShares | undefined;
}

export interface LiquidityProviders {
  /** Name of associated pool */
  name: string;
  /** liquidity providers */
  LiquidProviders: LiquidityProvider[];
}

export interface PoolAsset {
  /** symbol contains denom */
  symbol: string;
  /** skd.Int amout of pool asset */
  amount: string;
}

export interface PoolAssets {
  pa: PoolAsset[];
}

export interface PoolShares {
  /** symbol contains share denom */
  symbol: string;
  /** amount of shares */
  amount: string;
}

export interface Pool {
  /**
   * name of pool to index by
   * structure as: tokenA-tokenB-...-tokenN
   * in a lexigraphical ordering
   */
  id: string;
  /** pool assets owned by pool */
  poolAssets: PoolAssets | undefined;
  /** pool shares live in chain (minted and not yet burned) */
  poolShares: PoolShares | undefined;
  /**
   * liquidity providers for the pool
   * pool shares should equal the sum of
   * liquidity providers' shares amounts
   */
  liquidityProviders: LiquidityProviders | undefined;
  /**
   * fee to swap using pool
   * TODO figure out how fee (charged in native tokens) should
   * be determined for a pool between two external tokens
   */
  swapFee: string;
  /**
   * fee to exit the pool
   * TODO determine how fee is paid - could burn % before
   * exchanging shares for assets
   */
  exitFee: string;
}

const baseLiquidityProvider: object = { creator: "" };

export const LiquidityProvider = {
  encode(message: LiquidityProvider, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.liquid !== undefined) {
      PoolShares.encode(message.liquid, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LiquidityProvider {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLiquidityProvider } as LiquidityProvider;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.liquid = PoolShares.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LiquidityProvider {
    const message = { ...baseLiquidityProvider } as LiquidityProvider;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.liquid !== undefined && object.liquid !== null) {
      message.liquid = PoolShares.fromJSON(object.liquid);
    } else {
      message.liquid = undefined;
    }
    return message;
  },

  toJSON(message: LiquidityProvider): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.liquid !== undefined &&
      (obj.liquid = message.liquid
        ? PoolShares.toJSON(message.liquid)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<LiquidityProvider>): LiquidityProvider {
    const message = { ...baseLiquidityProvider } as LiquidityProvider;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.liquid !== undefined && object.liquid !== null) {
      message.liquid = PoolShares.fromPartial(object.liquid);
    } else {
      message.liquid = undefined;
    }
    return message;
  },
};

const baseLiquidityProviders: object = { name: "" };

export const LiquidityProviders = {
  encode(
    message: LiquidityProviders,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.LiquidProviders) {
      LiquidityProvider.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LiquidityProviders {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLiquidityProviders } as LiquidityProviders;
    message.LiquidProviders = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.LiquidProviders.push(
            LiquidityProvider.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): LiquidityProviders {
    const message = { ...baseLiquidityProviders } as LiquidityProviders;
    message.LiquidProviders = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (
      object.LiquidProviders !== undefined &&
      object.LiquidProviders !== null
    ) {
      for (const e of object.LiquidProviders) {
        message.LiquidProviders.push(LiquidityProvider.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: LiquidityProviders): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.LiquidProviders) {
      obj.LiquidProviders = message.LiquidProviders.map((e) =>
        e ? LiquidityProvider.toJSON(e) : undefined
      );
    } else {
      obj.LiquidProviders = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<LiquidityProviders>): LiquidityProviders {
    const message = { ...baseLiquidityProviders } as LiquidityProviders;
    message.LiquidProviders = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (
      object.LiquidProviders !== undefined &&
      object.LiquidProviders !== null
    ) {
      for (const e of object.LiquidProviders) {
        message.LiquidProviders.push(LiquidityProvider.fromPartial(e));
      }
    }
    return message;
  },
};

const basePoolAsset: object = { symbol: "", amount: "" };

export const PoolAsset = {
  encode(message: PoolAsset, writer: Writer = Writer.create()): Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PoolAsset {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePoolAsset } as PoolAsset;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolAsset {
    const message = { ...basePoolAsset } as PoolAsset;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: PoolAsset): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<PoolAsset>): PoolAsset {
    const message = { ...basePoolAsset } as PoolAsset;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const basePoolAssets: object = {};

export const PoolAssets = {
  encode(message: PoolAssets, writer: Writer = Writer.create()): Writer {
    for (const v of message.pa) {
      PoolAsset.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PoolAssets {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePoolAssets } as PoolAssets;
    message.pa = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pa.push(PoolAsset.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolAssets {
    const message = { ...basePoolAssets } as PoolAssets;
    message.pa = [];
    if (object.pa !== undefined && object.pa !== null) {
      for (const e of object.pa) {
        message.pa.push(PoolAsset.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: PoolAssets): unknown {
    const obj: any = {};
    if (message.pa) {
      obj.pa = message.pa.map((e) => (e ? PoolAsset.toJSON(e) : undefined));
    } else {
      obj.pa = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<PoolAssets>): PoolAssets {
    const message = { ...basePoolAssets } as PoolAssets;
    message.pa = [];
    if (object.pa !== undefined && object.pa !== null) {
      for (const e of object.pa) {
        message.pa.push(PoolAsset.fromPartial(e));
      }
    }
    return message;
  },
};

const basePoolShares: object = { symbol: "", amount: "" };

export const PoolShares = {
  encode(message: PoolShares, writer: Writer = Writer.create()): Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PoolShares {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePoolShares } as PoolShares;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PoolShares {
    const message = { ...basePoolShares } as PoolShares;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = String(object.symbol);
    } else {
      message.symbol = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: PoolShares): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<PoolShares>): PoolShares {
    const message = { ...basePoolShares } as PoolShares;
    if (object.symbol !== undefined && object.symbol !== null) {
      message.symbol = object.symbol;
    } else {
      message.symbol = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const basePool: object = { id: "", swapFee: "", exitFee: "" };

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.poolAssets !== undefined) {
      PoolAssets.encode(message.poolAssets, writer.uint32(18).fork()).ldelim();
    }
    if (message.poolShares !== undefined) {
      PoolShares.encode(message.poolShares, writer.uint32(26).fork()).ldelim();
    }
    if (message.liquidityProviders !== undefined) {
      LiquidityProviders.encode(
        message.liquidityProviders,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.swapFee !== "") {
      writer.uint32(42).string(message.swapFee);
    }
    if (message.exitFee !== "") {
      writer.uint32(50).string(message.exitFee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.poolAssets = PoolAssets.decode(reader, reader.uint32());
          break;
        case 3:
          message.poolShares = PoolShares.decode(reader, reader.uint32());
          break;
        case 4:
          message.liquidityProviders = LiquidityProviders.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.swapFee = reader.string();
          break;
        case 6:
          message.exitFee = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id);
    } else {
      message.id = "";
    }
    if (object.poolAssets !== undefined && object.poolAssets !== null) {
      message.poolAssets = PoolAssets.fromJSON(object.poolAssets);
    } else {
      message.poolAssets = undefined;
    }
    if (object.poolShares !== undefined && object.poolShares !== null) {
      message.poolShares = PoolShares.fromJSON(object.poolShares);
    } else {
      message.poolShares = undefined;
    }
    if (
      object.liquidityProviders !== undefined &&
      object.liquidityProviders !== null
    ) {
      message.liquidityProviders = LiquidityProviders.fromJSON(
        object.liquidityProviders
      );
    } else {
      message.liquidityProviders = undefined;
    }
    if (object.swapFee !== undefined && object.swapFee !== null) {
      message.swapFee = String(object.swapFee);
    } else {
      message.swapFee = "";
    }
    if (object.exitFee !== undefined && object.exitFee !== null) {
      message.exitFee = String(object.exitFee);
    } else {
      message.exitFee = "";
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.poolAssets !== undefined &&
      (obj.poolAssets = message.poolAssets
        ? PoolAssets.toJSON(message.poolAssets)
        : undefined);
    message.poolShares !== undefined &&
      (obj.poolShares = message.poolShares
        ? PoolShares.toJSON(message.poolShares)
        : undefined);
    message.liquidityProviders !== undefined &&
      (obj.liquidityProviders = message.liquidityProviders
        ? LiquidityProviders.toJSON(message.liquidityProviders)
        : undefined);
    message.swapFee !== undefined && (obj.swapFee = message.swapFee);
    message.exitFee !== undefined && (obj.exitFee = message.exitFee);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = "";
    }
    if (object.poolAssets !== undefined && object.poolAssets !== null) {
      message.poolAssets = PoolAssets.fromPartial(object.poolAssets);
    } else {
      message.poolAssets = undefined;
    }
    if (object.poolShares !== undefined && object.poolShares !== null) {
      message.poolShares = PoolShares.fromPartial(object.poolShares);
    } else {
      message.poolShares = undefined;
    }
    if (
      object.liquidityProviders !== undefined &&
      object.liquidityProviders !== null
    ) {
      message.liquidityProviders = LiquidityProviders.fromPartial(
        object.liquidityProviders
      );
    } else {
      message.liquidityProviders = undefined;
    }
    if (object.swapFee !== undefined && object.swapFee !== null) {
      message.swapFee = object.swapFee;
    } else {
      message.swapFee = "";
    }
    if (object.exitFee !== undefined && object.exitFee !== null) {
      message.exitFee = object.exitFee;
    } else {
      message.exitFee = "";
    }
    return message;
  },
};

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
