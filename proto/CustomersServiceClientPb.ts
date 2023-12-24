/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.5.0
// 	protoc              v0.0.0
// source: proto/customers.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as proto_customers_pb from '../proto/customers_pb'; // proto import: "proto/customers.proto"


export class CustomersServiceClient {
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
    '/proto.CustomersService/Create',
    grpcWeb.MethodType.UNARY,
    proto_customers_pb.CustomerCreateRequest,
    proto_customers_pb.CustomerResponse,
    (request: proto_customers_pb.CustomerCreateRequest) => {
      return request.serializeBinary();
    },
    proto_customers_pb.CustomerResponse.deserializeBinary
  );

  create(
    request: proto_customers_pb.CustomerCreateRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_customers_pb.CustomerResponse>;

  create(
    request: proto_customers_pb.CustomerCreateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerResponse) => void): grpcWeb.ClientReadableStream<proto_customers_pb.CustomerResponse>;

  create(
    request: proto_customers_pb.CustomerCreateRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.CustomersService/Create',
        request,
        metadata || {},
        this.methodDescriptorCreate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.CustomersService/Create',
    request,
    metadata || {},
    this.methodDescriptorCreate);
  }

  methodDescriptorGet = new grpcWeb.MethodDescriptor(
    '/proto.CustomersService/Get',
    grpcWeb.MethodType.UNARY,
    proto_customers_pb.CustomerGetRequest,
    proto_customers_pb.CustomerGetResponse,
    (request: proto_customers_pb.CustomerGetRequest) => {
      return request.serializeBinary();
    },
    proto_customers_pb.CustomerGetResponse.deserializeBinary
  );

  get(
    request: proto_customers_pb.CustomerGetRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_customers_pb.CustomerGetResponse>;

  get(
    request: proto_customers_pb.CustomerGetRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerGetResponse) => void): grpcWeb.ClientReadableStream<proto_customers_pb.CustomerGetResponse>;

  get(
    request: proto_customers_pb.CustomerGetRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerGetResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.CustomersService/Get',
        request,
        metadata || {},
        this.methodDescriptorGet,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.CustomersService/Get',
    request,
    metadata || {},
    this.methodDescriptorGet);
  }

  methodDescriptorUpdate = new grpcWeb.MethodDescriptor(
    '/proto.CustomersService/Update',
    grpcWeb.MethodType.UNARY,
    proto_customers_pb.CustomerUpdateRequest,
    proto_customers_pb.CustomerResponse,
    (request: proto_customers_pb.CustomerUpdateRequest) => {
      return request.serializeBinary();
    },
    proto_customers_pb.CustomerResponse.deserializeBinary
  );

  update(
    request: proto_customers_pb.CustomerUpdateRequest,
    metadata?: grpcWeb.Metadata | null): Promise<proto_customers_pb.CustomerResponse>;

  update(
    request: proto_customers_pb.CustomerUpdateRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerResponse) => void): grpcWeb.ClientReadableStream<proto_customers_pb.CustomerResponse>;

  update(
    request: proto_customers_pb.CustomerUpdateRequest,
    metadata?: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: proto_customers_pb.CustomerResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/proto.CustomersService/Update',
        request,
        metadata || {},
        this.methodDescriptorUpdate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/proto.CustomersService/Update',
    request,
    metadata || {},
    this.methodDescriptorUpdate);
  }

}

