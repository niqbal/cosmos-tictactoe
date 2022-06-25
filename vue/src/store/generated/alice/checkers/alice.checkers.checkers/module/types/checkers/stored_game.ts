/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.checkers";

export interface StoredGame {
  index: string;
  creator: string;
  idValue: number;
  game: string;
  crossPlayer: string;
  circlePlayer: string;
}

const baseStoredGame: object = {
  index: "",
  creator: "",
  idValue: 0,
  game: "",
  crossPlayer: "",
  circlePlayer: "",
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creator !== "") {
      writer.uint32(18).string(message.creator);
    }
    if (message.idValue !== 0) {
      writer.uint32(24).uint64(message.idValue);
    }
    if (message.game !== "") {
      writer.uint32(34).string(message.game);
    }
    if (message.crossPlayer !== "") {
      writer.uint32(42).string(message.crossPlayer);
    }
    if (message.circlePlayer !== "") {
      writer.uint32(50).string(message.circlePlayer);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creator = reader.string();
          break;
        case 3:
          message.idValue = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.game = reader.string();
          break;
        case 5:
          message.crossPlayer = reader.string();
          break;
        case 6:
          message.circlePlayer = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = Number(object.idValue);
    } else {
      message.idValue = 0;
    }
    if (object.game !== undefined && object.game !== null) {
      message.game = String(object.game);
    } else {
      message.game = "";
    }
    if (object.crossPlayer !== undefined && object.crossPlayer !== null) {
      message.crossPlayer = String(object.crossPlayer);
    } else {
      message.crossPlayer = "";
    }
    if (object.circlePlayer !== undefined && object.circlePlayer !== null) {
      message.circlePlayer = String(object.circlePlayer);
    } else {
      message.circlePlayer = "";
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creator !== undefined && (obj.creator = message.creator);
    message.idValue !== undefined && (obj.idValue = message.idValue);
    message.game !== undefined && (obj.game = message.game);
    message.crossPlayer !== undefined &&
      (obj.crossPlayer = message.crossPlayer);
    message.circlePlayer !== undefined &&
      (obj.circlePlayer = message.circlePlayer);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.idValue !== undefined && object.idValue !== null) {
      message.idValue = object.idValue;
    } else {
      message.idValue = 0;
    }
    if (object.game !== undefined && object.game !== null) {
      message.game = object.game;
    } else {
      message.game = "";
    }
    if (object.crossPlayer !== undefined && object.crossPlayer !== null) {
      message.crossPlayer = object.crossPlayer;
    } else {
      message.crossPlayer = "";
    }
    if (object.circlePlayer !== undefined && object.circlePlayer !== null) {
      message.circlePlayer = object.circlePlayer;
    } else {
      message.circlePlayer = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
