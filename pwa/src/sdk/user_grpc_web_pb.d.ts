import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as user_pb from './user_pb';


export class UsersClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.User) => void
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  get(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.User) => void
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  update(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.User) => void
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  delete(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  login(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.User) => void
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  certificate(
    request: user_pb.User,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.User) => void
  ): grpcWeb.ClientReadableStream<user_pb.User>;

}

export class AddressesClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Address) => void
  ): grpcWeb.ClientReadableStream<user_pb.Address>;

  get(
    request: user_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Address) => void
  ): grpcWeb.ClientReadableStream<user_pb.Address>;

  update(
    request: user_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Address) => void
  ): grpcWeb.ClientReadableStream<user_pb.Address>;

  delete(
    request: user_pb.Address,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.Address>;

}

export class MemosClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<user_pb.Memo>;

  get(
    request: user_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<user_pb.Memo>;

  update(
    request: user_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: user_pb.Memo) => void
  ): grpcWeb.ClientReadableStream<user_pb.Memo>;

  delete(
    request: user_pb.Memo,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.Memo>;

}

export class UsersPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.User>;

  get(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.User>;

  update(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.User>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.User>;

  delete(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  login(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.User>;

  certificate(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.User>;

}

export class AddressesPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Address>;

  get(
    request: user_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Address>;

  update(
    request: user_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Address>;

  delete(
    request: user_pb.Address,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.Address>;

}

export class MemosPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: user_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Memo>;

  get(
    request: user_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Memo>;

  update(
    request: user_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<user_pb.Memo>;

  delete(
    request: user_pb.Memo,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

  list(
    request: user_pb.User,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<user_pb.Memo>;

}

