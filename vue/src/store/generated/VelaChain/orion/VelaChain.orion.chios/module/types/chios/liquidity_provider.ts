/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "VelaChain.orion.chios";

export interface LiquidityProvider {
  address: string;
  liquidity: number;
}

const baseLiquidityProvider: object = { address: "", liquidity: 0 };

export const LiquidityProvider = {
  encode(message: LiquidityProvider, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.liquidity !== 0) {
      writer.uint32(16).int32(message.liquidity);
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
          message.address = reader.string();
          break;
        case 2:
          message.liquidity = reader.int32();
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
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.liquidity !== undefined && object.liquidity !== null) {
      message.liquidity = Number(object.liquidity);
    } else {
      message.liquidity = 0;
    }
    return message;
  },

  toJSON(message: LiquidityProvider): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.liquidity !== undefined && (obj.liquidity = message.liquidity);
    return obj;
  },

  fromPartial(object: DeepPartial<LiquidityProvider>): LiquidityProvider {
    const message = { ...baseLiquidityProvider } as LiquidityProvider;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.liquidity !== undefined && object.liquidity !== null) {
      message.liquidity = object.liquidity;
    } else {
      message.liquidity = 0;
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
