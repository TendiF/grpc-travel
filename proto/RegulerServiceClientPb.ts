/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v0.0.0
// source: proto/reguler.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_reguler_pb from '../proto/reguler_pb';


export class RegulersServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorCreate = new grpcWeb.MethodDescriptor(
    '/proto.RegulersService/Create',
    grpcWeb.MethodType.UNARY,
    proto_reguler_pb.RegulerCreateRequest,
    proto_reguler_pb.RegulerResponse,
    (request: proto_reguler_pb.RegulerCreateRequest) => {
      return request.serializeBinary();
    },
    proto_reguler_pb.RegulerResponse.deserializeBinary
  );

  create(
    request: proto_reguler_pb.RegulerCreateRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_reguler_pb.RegulerResponse>;

  create(
    request: proto_reguler_pb.RegulerCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerResponse) => void): grpcWeb.ClientReadableStream<proto_reguler_pb.RegulerResponse>;

  create(
    request: proto_reguler_pb.RegulerCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.RegulersService/Create',
        request,
        metadata || {},
        this.methodDescriptorCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.RegulersService/Create',
    request,
    metadata || {},
    this.methodDescriptorCreate);
  }

  methodDescriptorGet = new grpcWeb.MethodDescriptor(
    '/proto.RegulersService/Get',
    grpcWeb.MethodType.UNARY,
    proto_reguler_pb.RegulerGetRequest,
    proto_reguler_pb.RegulerGetResponse,
    (request: proto_reguler_pb.RegulerGetRequest) => {
      return request.serializeBinary();
    },
    proto_reguler_pb.RegulerGetResponse.deserializeBinary
  );

  get(
    request: proto_reguler_pb.RegulerGetRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_reguler_pb.RegulerGetResponse>;

  get(
    request: proto_reguler_pb.RegulerGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerGetResponse) => void): grpcWeb.ClientReadableStream<proto_reguler_pb.RegulerGetResponse>;

  get(
    request: proto_reguler_pb.RegulerGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.RegulersService/Get',
        request,
        metadata || {},
        this.methodDescriptorGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.RegulersService/Get',
    request,
    metadata || {},
    this.methodDescriptorGet);
  }

  methodDescriptorUpdate = new grpcWeb.MethodDescriptor(
    '/proto.RegulersService/Update',
    grpcWeb.MethodType.UNARY,
    proto_reguler_pb.RegulerUpdateRequest,
    proto_reguler_pb.RegulerResponse,
    (request: proto_reguler_pb.RegulerUpdateRequest) => {
      return request.serializeBinary();
    },
    proto_reguler_pb.RegulerResponse.deserializeBinary
  );

  update(
    request: proto_reguler_pb.RegulerUpdateRequest,
    metadata: grpcWeb.Metadata | null): Promise<proto_reguler_pb.RegulerResponse>;

  update(
    request: proto_reguler_pb.RegulerUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerResponse) => void): grpcWeb.ClientReadableStream<proto_reguler_pb.RegulerResponse>;

  update(
    request: proto_reguler_pb.RegulerUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_reguler_pb.RegulerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.RegulersService/Update',
        request,
        metadata || {},
        this.methodDescriptorUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.RegulersService/Update',
    request,
    metadata || {},
    this.methodDescriptorUpdate);
  }

}

