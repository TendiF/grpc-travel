import * as jspb from 'google-protobuf'

import * as proto_reguler_pb from '../proto/reguler_pb'; // proto import: "proto/reguler.proto"


export class CustomerCreateRequest extends jspb.Message {
  getDataList(): Array<Customer>;
  setDataList(value: Array<Customer>): CustomerCreateRequest;
  clearDataList(): CustomerCreateRequest;
  addData(value?: Customer, index?: number): Customer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerCreateRequest): CustomerCreateRequest.AsObject;
  static serializeBinaryToWriter(message: CustomerCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerCreateRequest;
  static deserializeBinaryFromReader(message: CustomerCreateRequest, reader: jspb.BinaryReader): CustomerCreateRequest;
}

export namespace CustomerCreateRequest {
  export type AsObject = {
    dataList: Array<Customer.AsObject>,
  }
}

export class Customer extends jspb.Message {
  getNik(): string;
  setNik(value: string): Customer;

  getNama(): string;
  setNama(value: string): Customer;

  getJk(): string;
  setJk(value: string): Customer;

  getStatus(): string;
  setStatus(value: string): Customer;

  getStatuskk(): string;
  setStatuskk(value: string): Customer;

  getNokk(): string;
  setNokk(value: string): Customer;

  getTanggallahir(): string;
  setTanggallahir(value: string): Customer;

  getKotakab(): string;
  setKotakab(value: string): Customer;

  getKecamatan(): string;
  setKecamatan(value: string): Customer;

  getDesakelurahan(): string;
  setDesakelurahan(value: string): Customer;

  getKampung(): string;
  setKampung(value: string): Customer;

  getRt(): string;
  setRt(value: string): Customer;

  getRw(): string;
  setRw(value: string): Customer;

  getKol(): string;
  setKol(value: string): Customer;

  getSyahidan(): string;
  setSyahidan(value: string): Customer;

  getPj(): string;
  setPj(value: string): Customer;

  getNote(): string;
  setNote(value: string): Customer;

  getId(): string;
  setId(value: string): Customer;

  getReguler(): proto_reguler_pb.RegulerCreateRequest | undefined;
  setReguler(value?: proto_reguler_pb.RegulerCreateRequest): Customer;
  hasReguler(): boolean;
  clearReguler(): Customer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Customer.AsObject;
  static toObject(includeInstance: boolean, msg: Customer): Customer.AsObject;
  static serializeBinaryToWriter(message: Customer, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Customer;
  static deserializeBinaryFromReader(message: Customer, reader: jspb.BinaryReader): Customer;
}

export namespace Customer {
  export type AsObject = {
    nik: string,
    nama: string,
    jk: string,
    status: string,
    statuskk: string,
    nokk: string,
    tanggallahir: string,
    kotakab: string,
    kecamatan: string,
    desakelurahan: string,
    kampung: string,
    rt: string,
    rw: string,
    kol: string,
    syahidan: string,
    pj: string,
    note: string,
    id: string,
    reguler?: proto_reguler_pb.RegulerCreateRequest.AsObject,
  }
}

export class CustomerSortParam extends jspb.Message {
  getNokk(): number;
  setNokk(value: number): CustomerSortParam;

  getPj(): number;
  setPj(value: number): CustomerSortParam;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerSortParam.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerSortParam): CustomerSortParam.AsObject;
  static serializeBinaryToWriter(message: CustomerSortParam, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerSortParam;
  static deserializeBinaryFromReader(message: CustomerSortParam, reader: jspb.BinaryReader): CustomerSortParam;
}

export namespace CustomerSortParam {
  export type AsObject = {
    nokk: number,
    pj: number,
  }
}

export class CustomerResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): CustomerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerResponse): CustomerResponse.AsObject;
  static serializeBinaryToWriter(message: CustomerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerResponse;
  static deserializeBinaryFromReader(message: CustomerResponse, reader: jspb.BinaryReader): CustomerResponse;
}

export namespace CustomerResponse {
  export type AsObject = {
    message: string,
  }
}

export class CustomerGetResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): CustomerGetResponse;

  getTotalPage(): number;
  setTotalPage(value: number): CustomerGetResponse;

  getDataList(): Array<Customer>;
  setDataList(value: Array<Customer>): CustomerGetResponse;
  clearDataList(): CustomerGetResponse;
  addData(value?: Customer, index?: number): Customer;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerGetResponse): CustomerGetResponse.AsObject;
  static serializeBinaryToWriter(message: CustomerGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerGetResponse;
  static deserializeBinaryFromReader(message: CustomerGetResponse, reader: jspb.BinaryReader): CustomerGetResponse;
}

export namespace CustomerGetResponse {
  export type AsObject = {
    message: string,
    totalPage: number,
    dataList: Array<Customer.AsObject>,
  }
}

export class CustomerGetRequest extends jspb.Message {
  getPage(): number;
  setPage(value: number): CustomerGetRequest;

  getPerPage(): number;
  setPerPage(value: number): CustomerGetRequest;

  getSearch(): string;
  setSearch(value: string): CustomerGetRequest;

  getTahun(): string;
  setTahun(value: string): CustomerGetRequest;

  getBulan(): string;
  setBulan(value: string): CustomerGetRequest;

  getSort(): CustomerSortParam | undefined;
  setSort(value?: CustomerSortParam): CustomerGetRequest;
  hasSort(): boolean;
  clearSort(): CustomerGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerGetRequest): CustomerGetRequest.AsObject;
  static serializeBinaryToWriter(message: CustomerGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerGetRequest;
  static deserializeBinaryFromReader(message: CustomerGetRequest, reader: jspb.BinaryReader): CustomerGetRequest;
}

export namespace CustomerGetRequest {
  export type AsObject = {
    page: number,
    perPage: number,
    search: string,
    tahun: string,
    bulan: string,
    sort?: CustomerSortParam.AsObject,
  }
}

export class CustomerUpdateRequest extends jspb.Message {
  getNik(): string;
  setNik(value: string): CustomerUpdateRequest;

  getNama(): string;
  setNama(value: string): CustomerUpdateRequest;

  getStatus(): string;
  setStatus(value: string): CustomerUpdateRequest;

  getStatuskk(): string;
  setStatuskk(value: string): CustomerUpdateRequest;

  getNokk(): string;
  setNokk(value: string): CustomerUpdateRequest;

  getTanggallahir(): string;
  setTanggallahir(value: string): CustomerUpdateRequest;

  getKotakab(): string;
  setKotakab(value: string): CustomerUpdateRequest;

  getKecamatan(): string;
  setKecamatan(value: string): CustomerUpdateRequest;

  getDesakelurahan(): string;
  setDesakelurahan(value: string): CustomerUpdateRequest;

  getKampung(): string;
  setKampung(value: string): CustomerUpdateRequest;

  getRt(): string;
  setRt(value: string): CustomerUpdateRequest;

  getRw(): string;
  setRw(value: string): CustomerUpdateRequest;

  getKol(): string;
  setKol(value: string): CustomerUpdateRequest;

  getSyahidan(): string;
  setSyahidan(value: string): CustomerUpdateRequest;

  getPj(): string;
  setPj(value: string): CustomerUpdateRequest;

  getNote(): string;
  setNote(value: string): CustomerUpdateRequest;

  getId(): string;
  setId(value: string): CustomerUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerUpdateRequest): CustomerUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: CustomerUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerUpdateRequest;
  static deserializeBinaryFromReader(message: CustomerUpdateRequest, reader: jspb.BinaryReader): CustomerUpdateRequest;
}

export namespace CustomerUpdateRequest {
  export type AsObject = {
    nik: string,
    nama: string,
    status: string,
    statuskk: string,
    nokk: string,
    tanggallahir: string,
    kotakab: string,
    kecamatan: string,
    desakelurahan: string,
    kampung: string,
    rt: string,
    rw: string,
    kol: string,
    syahidan: string,
    pj: string,
    note: string,
    id: string,
  }
}

