import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Order extends jspb.Message {
  getId(): string;
  setId(value: string): Order;

  getType(): string;
  setType(value: string): Order;

  getHospital(): string;
  setHospital(value: string): Order;

  getDepartment(): string;
  setDepartment(value: string): Order;

  getTargetTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTargetTime(value?: google_protobuf_timestamp_pb.Timestamp): Order;
  hasTargetTime(): boolean;
  clearTargetTime(): Order;

  getPatient(): string;
  setPatient(value: string): Order;

  getCompanion(): string;
  setCompanion(value: string): Order;

  getFee(): number;
  setFee(value: number): Order;

  getNotes(): string;
  setNotes(value: string): Order;

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
    type: string,
    hospital: string,
    department: string,
    targetTime?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    patient: string,
    companion: string,
    fee: number,
    notes: string,
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

