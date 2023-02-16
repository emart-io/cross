import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as attendant_pb from './attendant_pb';


export class AttendantsClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Attendant) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  get(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Attendant) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  update(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Attendant) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  delete(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  login(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Attendant) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  certificate(
    request: attendant_pb.Attendant,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Attendant) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

}

export class AddressesClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Address) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Address>;

  get(
    request: attendant_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Address) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Address>;

  update(
    request: attendant_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Address) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Address>;

  delete(
    request: attendant_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Address>;

}

export class MemosClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Memo>;

  get(
    request: attendant_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Memo>;

  update(
    request: attendant_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: attendant_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<attendant_pb.Memo>;

  delete(
    request: attendant_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Memo>;

}

export class AttendantsPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Attendant>;

  get(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Attendant>;

  update(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Attendant>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Attendant>;

  delete(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  login(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Attendant>;

  certificate(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Attendant>;

}

export class AddressesPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Address>;

  get(
    request: attendant_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Address>;

  update(
    request: attendant_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Address>;

  delete(
    request: attendant_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Address>;

}

export class MemosPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: attendant_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Memo>;

  get(
    request: attendant_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Memo>;

  update(
    request: attendant_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<attendant_pb.Memo>;

  delete(
    request: attendant_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  list(
    request: attendant_pb.Attendant,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<attendant_pb.Memo>;

}

