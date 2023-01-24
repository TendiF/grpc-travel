import * as jspb from 'google-protobuf'



export class CustomerCreateRequest extends jspb.Message {
  getNo(): string;
  setNo(value: string): CustomerCreateRequest;

  getNik(): string;
  setNik(value: string): CustomerCreateRequest;

  getNama(): string;
  setNama(value: string): CustomerCreateRequest;

  getStatuskk(): string;
  setStatuskk(value: string): CustomerCreateRequest;

  getNokk(): string;
  setNokk(value: string): CustomerCreateRequest;

  getTanggallahir(): string;
  setTanggallahir(value: string): CustomerCreateRequest;

  getUsia(): string;
  setUsia(value: string): CustomerCreateRequest;

  getKotakab(): string;
  setKotakab(value: string): CustomerCreateRequest;

  getKecamatan(): string;
  setKecamatan(value: string): CustomerCreateRequest;

  getDesakelurahan(): string;
  setDesakelurahan(value: string): CustomerCreateRequest;

  getKampung(): string;
  setKampung(value: string): CustomerCreateRequest;

  getRt(): string;
  setRt(value: string): CustomerCreateRequest;

  getRw(): string;
  setRw(value: string): CustomerCreateRequest;

  getKol(): string;
  setKol(value: string): CustomerCreateRequest;

  getSyahidan(): string;
  setSyahidan(value: string): CustomerCreateRequest;

  getPj(): string;
  setPj(value: string): CustomerCreateRequest;

  getNote(): string;
  setNote(value: string): CustomerCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CustomerCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: CustomerCreateRequest): CustomerCreateRequest.AsObject;
  static serializeBinaryToWriter(message: CustomerCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CustomerCreateRequest;
  static deserializeBinaryFromReader(message: CustomerCreateRequest, reader: jspb.BinaryReader): CustomerCreateRequest;
}

export namespace CustomerCreateRequest {
  export type AsObject = {
    no: string,
    nik: string,
    nama: string,
    statuskk: string,
    nokk: string,
    tanggallahir: string,
    usia: string,
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

  getDataList(): Array<CustomerCreateRequest>;
  setDataList(value: Array<CustomerCreateRequest>): CustomerGetResponse;
  clearDataList(): CustomerGetResponse;
  addData(value?: CustomerCreateRequest, index?: number): CustomerCreateRequest;

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
    dataList: Array<CustomerCreateRequest.AsObject>,
  }
}

export class CustomerGetRequest extends jspb.Message {
  getPage(): number;
  setPage(value: number): CustomerGetRequest;

  getPerPage(): number;
  setPerPage(value: number): CustomerGetRequest;

  getSearch(): string;
  setSearch(value: string): CustomerGetRequest;

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
  }
}

export class CustomerUpdateRequest extends jspb.Message {
  getNo(): string;
  setNo(value: string): CustomerUpdateRequest;

  getNik(): string;
  setNik(value: string): CustomerUpdateRequest;

  getNama(): string;
  setNama(value: string): CustomerUpdateRequest;

  getStatuskk(): string;
  setStatuskk(value: string): CustomerUpdateRequest;

  getNokk(): string;
  setNokk(value: string): CustomerUpdateRequest;

  getTanggallahir(): string;
  setTanggallahir(value: string): CustomerUpdateRequest;

  getUsia(): string;
  setUsia(value: string): CustomerUpdateRequest;

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
    no: string,
    nik: string,
    nama: string,
    statuskk: string,
    nokk: string,
    tanggallahir: string,
    usia: string,
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

