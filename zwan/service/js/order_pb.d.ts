import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Order extends jspb.Message {
  getId(): string;
  setId(value: string): Order;

  getName(): string;
  setName(value: string): Order;

  getPriority(): number;
  setPriority(value: number): Order;

  getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreated(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasCreated(): boolean;
  clearCreated(): Order;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Order.AsObject;
  static toObject(includeInstance: boolean, msg: Order): Order.AsObject;
  static serializeBinaryToWriter(message: Order, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Order;
  static deserializeBinaryFromReader(message: Order, reader: jspb.BinaryReader): Order;
}

export namespace Order {
  export type AsObject = {
    id: string,
    name: string,
    priority: number,
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class OrderRequest extends jspb.Message {
  getId(): string;
  setId(value: string): OrderRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): OrderRequest.AsObject;
  static toObject(includeInstance: boolean, msg: OrderRequest): OrderRequest.AsObject;
  static serializeBinaryToWriter(message: OrderRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): OrderRequest;
  static deserializeBinaryFromReader(message: OrderRequest, reader: jspb.BinaryReader): OrderRequest;
}

export namespace OrderRequest {
  export type AsObject = {
    id: string,
  }
}

