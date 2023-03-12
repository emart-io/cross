import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as order_pb from './order_pb';


export class OrdersClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: order_pb.Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: order_pb.Order) => void
  ): grpcWeb.ClientReadableStream<order_pb.Order>;

  get(
    request: order_pb.OrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: order_pb.Order) => void
  ): grpcWeb.ClientReadableStream<order_pb.Order>;

  update(
    request: order_pb.Order,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: order_pb.Order) => void
  ): grpcWeb.ClientReadableStream<order_pb.Order>;

  list(
    request: order_pb.Order,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<order_pb.Order>;

  delete(
    request: order_pb.OrderRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: google_protobuf_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<google_protobuf_empty_pb.Empty>;

}

export class OrdersPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  add(
    request: order_pb.Order,
    metadata?: grpcWeb.Metadata
  ): Promise<order_pb.Order>;

  get(
    request: order_pb.OrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<order_pb.Order>;

  update(
    request: order_pb.Order,
    metadata?: grpcWeb.Metadata
  ): Promise<order_pb.Order>;

  list(
    request: order_pb.Order,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<order_pb.Order>;

  delete(
    request: order_pb.OrderRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<google_protobuf_empty_pb.Empty>;

}

