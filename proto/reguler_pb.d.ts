import * as jspb from 'google-protobuf'



export class RegulerCreateRequest extends jspb.Message {
  getCustomerid(): string;
  setCustomerid(value: string): RegulerCreateRequest;

  getBulan(): string;
  setBulan(value: string): RegulerCreateRequest;

  getTahun(): string;
  setTahun(value: string): RegulerCreateRequest;

  getInfaq(): number;
  setInfaq(value: number): RegulerCreateRequest;

  getZakat(): number;
  setZakat(value: number): RegulerCreateRequest;

  getShadaqah(): number;
  setShadaqah(value: number): RegulerCreateRequest;

  getIkhsan(): number;
  setIkhsan(value: number): RegulerCreateRequest;

  getSabil(): number;
  setSabil(value: number): RegulerCreateRequest;

  getTabunganfitrah(): number;
  setTabunganfitrah(value: number): RegulerCreateRequest;

  getTabunganqurban(): number;
  setTabunganqurban(value: number): RegulerCreateRequest;

  getBina(): string;
  setBina(value: string): RegulerCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegulerCreateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegulerCreateRequest): RegulerCreateRequest.AsObject;
  static serializeBinaryToWriter(message: RegulerCreateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegulerCreateRequest;
  static deserializeBinaryFromReader(message: RegulerCreateRequest, reader: jspb.BinaryReader): RegulerCreateRequest;
}

export namespace RegulerCreateRequest {
  export type AsObject = {
    customerid: string,
    bulan: string,
    tahun: string,
    infaq: number,
    zakat: number,
    shadaqah: number,
    ikhsan: number,
    sabil: number,
    tabunganfitrah: number,
    tabunganqurban: number,
    bina: string,
  }
}

export class RegulerResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): RegulerResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegulerResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegulerResponse): RegulerResponse.AsObject;
  static serializeBinaryToWriter(message: RegulerResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegulerResponse;
  static deserializeBinaryFromReader(message: RegulerResponse, reader: jspb.BinaryReader): RegulerResponse;
}

export namespace RegulerResponse {
  export type AsObject = {
    message: string,
  }
}

export class RegulerGetResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): RegulerGetResponse;

  getTotalPage(): number;
  setTotalPage(value: number): RegulerGetResponse;

  getDataList(): Array<RegulerCreateRequest>;
  setDataList(value: Array<RegulerCreateRequest>): RegulerGetResponse;
  clearDataList(): RegulerGetResponse;
  addData(value?: RegulerCreateRequest, index?: number): RegulerCreateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegulerGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RegulerGetResponse): RegulerGetResponse.AsObject;
  static serializeBinaryToWriter(message: RegulerGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegulerGetResponse;
  static deserializeBinaryFromReader(message: RegulerGetResponse, reader: jspb.BinaryReader): RegulerGetResponse;
}

export namespace RegulerGetResponse {
  export type AsObject = {
    message: string,
    totalPage: number,
    dataList: Array<RegulerCreateRequest.AsObject>,
  }
}

export class RegulerGetRequest extends jspb.Message {
  getPage(): number;
  setPage(value: number): RegulerGetRequest;

  getPerPage(): number;
  setPerPage(value: number): RegulerGetRequest;

  getSearch(): string;
  setSearch(value: string): RegulerGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegulerGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegulerGetRequest): RegulerGetRequest.AsObject;
  static serializeBinaryToWriter(message: RegulerGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegulerGetRequest;
  static deserializeBinaryFromReader(message: RegulerGetRequest, reader: jspb.BinaryReader): RegulerGetRequest;
}

export namespace RegulerGetRequest {
  export type AsObject = {
    page: number,
    perPage: number,
    search: string,
  }
}

export class RegulerUpdateRequest extends jspb.Message {
  getCustomerid(): string;
  setCustomerid(value: string): RegulerUpdateRequest;

  getBulantahun(): string;
  setBulantahun(value: string): RegulerUpdateRequest;

  getInfaq(): string;
  setInfaq(value: string): RegulerUpdateRequest;

  getZakat(): string;
  setZakat(value: string): RegulerUpdateRequest;

  getShadaqah(): string;
  setShadaqah(value: string): RegulerUpdateRequest;

  getIkhsan(): string;
  setIkhsan(value: string): RegulerUpdateRequest;

  getSabil(): string;
  setSabil(value: string): RegulerUpdateRequest;

  getTabunganfitrah(): string;
  setTabunganfitrah(value: string): RegulerUpdateRequest;

  getTabunganqurban(): string;
  setTabunganqurban(value: string): RegulerUpdateRequest;

  getBina(): string;
  setBina(value: string): RegulerUpdateRequest;

  getId(): string;
  setId(value: string): RegulerUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RegulerUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RegulerUpdateRequest): RegulerUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: RegulerUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RegulerUpdateRequest;
  static deserializeBinaryFromReader(message: RegulerUpdateRequest, reader: jspb.BinaryReader): RegulerUpdateRequest;
}

export namespace RegulerUpdateRequest {
  export type AsObject = {
    customerid: string,
    bulantahun: string,
    infaq: string,
    zakat: string,
    shadaqah: string,
    ikhsan: string,
    sabil: string,
    tabunganfitrah: string,
    tabunganqurban: string,
    bina: string,
    id: string,
  }
}

