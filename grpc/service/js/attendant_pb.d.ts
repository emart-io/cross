import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Attendant extends jspb.Message {
  getId(): string;
  setId(value: string): Attendant;

  getName(): string;
  setName(value: string): Attendant;

  getPassword(): string;
  setPassword(value: string): Attendant;

  getTelephone(): string;
  setTelephone(value: string): Attendant;

  getIcon(): string;
  setIcon(value: string): Attendant;

  getSignature(): string;
  setSignature(value: string): Attendant;

  getCert(): Certification | undefined;
  setCert(value?: Certification): Attendant;
  hasCert(): boolean;
  clearCert(): Attendant;

  getShopsList(): Array<Shop>;
  setShopsList(value: Array<Shop>): Attendant;
  clearShopsList(): Attendant;
  addShops(value?: Shop, index?: number): Shop;

  getAnnotationsMap(): jspb.Map<string, string>;
  clearAnnotationsMap(): Attendant;

  getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreated(value?: google_protobuf_timestamp_pb.Timestamp): Attendant;
  hasCreated(): boolean;
  clearCreated(): Attendant;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Attendant.AsObject;
  static toObject(includeInstance: boolean, msg: Attendant): Attendant.AsObject;
  static serializeBinaryToWriter(message: Attendant, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Attendant;
  static deserializeBinaryFromReader(message: Attendant, reader: jspb.BinaryReader): Attendant;
}

export namespace Attendant {
  export type AsObject = {
    id: string,
    name: string,
    password: string,
    telephone: string,
    icon: string,
    signature: string,
    cert?: Certification.AsObject,
    shopsList: Array<Shop.AsObject>,
    annotationsMap: Array<[string, string]>,
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Certification extends jspb.Message {
  getFullname(): string;
  setFullname(value: string): Certification;

  getIdcardno(): string;
  setIdcardno(value: string): Certification;

  getImagesList(): Array<string>;
  setImagesList(value: Array<string>): Certification;
  clearImagesList(): Certification;
  addImages(value: string, index?: number): Certification;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Certification.AsObject;
  static toObject(includeInstance: boolean, msg: Certification): Certification.AsObject;
  static serializeBinaryToWriter(message: Certification, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Certification;
  static deserializeBinaryFromReader(message: Certification, reader: jspb.BinaryReader): Certification;
}

export namespace Certification {
  export type AsObject = {
    fullname: string,
    idcardno: string,
    imagesList: Array<string>,
  }
}

export class Address extends jspb.Message {
  getId(): string;
  setId(value: string): Address;

  getAttendantid(): string;
  setAttendantid(value: string): Address;

  getContact(): string;
  setContact(value: string): Address;

  getTelephone(): string;
  setTelephone(value: string): Address;

  getLocation(): string;
  setLocation(value: string): Address;

  getDefault(): boolean;
  setDefault(value: boolean): Address;

  getAnnotationsMap(): jspb.Map<string, string>;
  clearAnnotationsMap(): Address;

  getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreated(value?: google_protobuf_timestamp_pb.Timestamp): Address;
  hasCreated(): boolean;
  clearCreated(): Address;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Address.AsObject;
  static toObject(includeInstance: boolean, msg: Address): Address.AsObject;
  static serializeBinaryToWriter(message: Address, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Address;
  static deserializeBinaryFromReader(message: Address, reader: jspb.BinaryReader): Address;
}

export namespace Address {
  export type AsObject = {
    id: string,
    attendantid: string,
    contact: string,
    telephone: string,
    location: string,
    pb_default: boolean,
    annotationsMap: Array<[string, string]>,
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Shop extends jspb.Message {
  getId(): string;
  setId(value: string): Shop;

  getName(): string;
  setName(value: string): Shop;

  getDescription(): string;
  setDescription(value: string): Shop;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Shop.AsObject;
  static toObject(includeInstance: boolean, msg: Shop): Shop.AsObject;
  static serializeBinaryToWriter(message: Shop, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Shop;
  static deserializeBinaryFromReader(message: Shop, reader: jspb.BinaryReader): Shop;
}

export namespace Shop {
  export type AsObject = {
    id: string,
    name: string,
    description: string,
  }
}

export class Memo extends jspb.Message {
  getId(): string;
  setId(value: string): Memo;

  getAttendantid(): string;
  setAttendantid(value: string): Memo;

  getTitle(): string;
  setTitle(value: string): Memo;

  getContent(): string;
  setContent(value: string): Memo;

  getLocation(): string;
  setLocation(value: string): Memo;

  getAnnotationsMap(): jspb.Map<string, string>;
  clearAnnotationsMap(): Memo;

  getCreated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreated(value?: google_protobuf_timestamp_pb.Timestamp): Memo;
  hasCreated(): boolean;
  clearCreated(): Memo;

  getUpdated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setUpdated(value?: google_protobuf_timestamp_pb.Timestamp): Memo;
  hasUpdated(): boolean;
  clearUpdated(): Memo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Memo.AsObject;
  static toObject(includeInstance: boolean, msg: Memo): Memo.AsObject;
  static serializeBinaryToWriter(message: Memo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Memo;
  static deserializeBinaryFromReader(message: Memo, reader: jspb.BinaryReader): Memo;
}

export namespace Memo {
  export type AsObject = {
    id: string,
    attendantid: string,
    title: string,
    content: string,
    location: string,
    annotationsMap: Array<[string, string]>,
    created?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    updated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

