import * as jspb from 'google-protobuf'


export class UserRequest extends jspb.Message {
  getFirstName(): string;
  setFirstName(value: string): UserRequest;

  getLastName(): string;
  setLastName(value: string): UserRequest;

  getGender(): string;
  setGender(value: string): UserRequest;

  getEmail(): string;
  setEmail(value: string): UserRequest;

  getPassword(): string;
  setPassword(value: string): UserRequest;

  getAddress(): string;
  setAddress(value: string): UserRequest;

  getUsername(): string;
  setUsername(value: string): UserRequest;

  getRole(): Role;
  setRole(value: Role): UserRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserRequest): UserRequest.AsObject;
  static serializeBinaryToWriter(message: UserRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserRequest;
  static deserializeBinaryFromReader(message: UserRequest, reader: jspb.BinaryReader): UserRequest;
}

export namespace UserRequest {
  export type AsObject = {
    firstName: string,
    lastName: string,
    gender: string,
    email: string,
    password: string,
    address: string,
    username: string,
    role: Role,
  }
}

export class UserResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UserResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserResponse): UserResponse.AsObject;
  static serializeBinaryToWriter(message: UserResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserResponse;
  static deserializeBinaryFromReader(message: UserResponse, reader: jspb.BinaryReader): UserResponse;
}

export namespace UserResponse {
  export type AsObject = {
    message: string,
  }
}

export class UserUpdateRequest extends jspb.Message {
  getId(): string;
  setId(value: string): UserUpdateRequest;

  getFirstName(): string;
  setFirstName(value: string): UserUpdateRequest;

  getLastName(): string;
  setLastName(value: string): UserUpdateRequest;

  getGender(): string;
  setGender(value: string): UserUpdateRequest;

  getEmail(): string;
  setEmail(value: string): UserUpdateRequest;

  getPassword(): string;
  setPassword(value: string): UserUpdateRequest;

  getAddress(): string;
  setAddress(value: string): UserUpdateRequest;

  getUsername(): string;
  setUsername(value: string): UserUpdateRequest;

  getRole(): Role;
  setRole(value: Role): UserUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserUpdateRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserUpdateRequest): UserUpdateRequest.AsObject;
  static serializeBinaryToWriter(message: UserUpdateRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserUpdateRequest;
  static deserializeBinaryFromReader(message: UserUpdateRequest, reader: jspb.BinaryReader): UserUpdateRequest;
}

export namespace UserUpdateRequest {
  export type AsObject = {
    id: string,
    firstName: string,
    lastName: string,
    gender: string,
    email: string,
    password: string,
    address: string,
    username: string,
    role: Role,
  }
}

export class UserGetRequest extends jspb.Message {
  getPage(): number;
  setPage(value: number): UserGetRequest;

  getPerPage(): number;
  setPerPage(value: number): UserGetRequest;

  getSearch(): string;
  setSearch(value: string): UserGetRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetRequest): UserGetRequest.AsObject;
  static serializeBinaryToWriter(message: UserGetRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetRequest;
  static deserializeBinaryFromReader(message: UserGetRequest, reader: jspb.BinaryReader): UserGetRequest;
}

export namespace UserGetRequest {
  export type AsObject = {
    page: number,
    perPage: number,
    search: string,
  }
}

export class UserGetResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): UserGetResponse;

  getTotalPage(): number;
  setTotalPage(value: number): UserGetResponse;

  getDataList(): Array<UserUpdateRequest>;
  setDataList(value: Array<UserUpdateRequest>): UserGetResponse;
  clearDataList(): UserGetResponse;
  addData(value?: UserUpdateRequest, index?: number): UserUpdateRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserGetResponse.AsObject;
  static toObject(includeInstance: boolean, msg: UserGetResponse): UserGetResponse.AsObject;
  static serializeBinaryToWriter(message: UserGetResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserGetResponse;
  static deserializeBinaryFromReader(message: UserGetResponse, reader: jspb.BinaryReader): UserGetResponse;
}

export namespace UserGetResponse {
  export type AsObject = {
    message: string,
    totalPage: number,
    dataList: Array<UserUpdateRequest.AsObject>,
  }
}

export class userDeleteRequest extends jspb.Message {
  getId(): string;
  setId(value: string): userDeleteRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): userDeleteRequest.AsObject;
  static toObject(includeInstance: boolean, msg: userDeleteRequest): userDeleteRequest.AsObject;
  static serializeBinaryToWriter(message: userDeleteRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): userDeleteRequest;
  static deserializeBinaryFromReader(message: userDeleteRequest, reader: jspb.BinaryReader): userDeleteRequest;
}

export namespace userDeleteRequest {
  export type AsObject = {
    id: string,
  }
}

export class userLoginRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): userLoginRequest;

  getPassword(): string;
  setPassword(value: string): userLoginRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): userLoginRequest.AsObject;
  static toObject(includeInstance: boolean, msg: userLoginRequest): userLoginRequest.AsObject;
  static serializeBinaryToWriter(message: userLoginRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): userLoginRequest;
  static deserializeBinaryFromReader(message: userLoginRequest, reader: jspb.BinaryReader): userLoginRequest;
}

export namespace userLoginRequest {
  export type AsObject = {
    username: string,
    password: string,
  }
}

export class userLoginResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): userLoginResponse;

  getToken(): string;
  setToken(value: string): userLoginResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): userLoginResponse.AsObject;
  static toObject(includeInstance: boolean, msg: userLoginResponse): userLoginResponse.AsObject;
  static serializeBinaryToWriter(message: userLoginResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): userLoginResponse;
  static deserializeBinaryFromReader(message: userLoginResponse, reader: jspb.BinaryReader): userLoginResponse;
}

export namespace userLoginResponse {
  export type AsObject = {
    message: string,
    token: string,
  }
}

export class userProfileRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): userProfileRequest.AsObject;
  static toObject(includeInstance: boolean, msg: userProfileRequest): userProfileRequest.AsObject;
  static serializeBinaryToWriter(message: userProfileRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): userProfileRequest;
  static deserializeBinaryFromReader(message: userProfileRequest, reader: jspb.BinaryReader): userProfileRequest;
}

export namespace userProfileRequest {
  export type AsObject = {
  }
}

export class userProfileResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): userProfileResponse;

  getProfile(): UserUpdateRequest | undefined;
  setProfile(value?: UserUpdateRequest): userProfileResponse;
  hasProfile(): boolean;
  clearProfile(): userProfileResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): userProfileResponse.AsObject;
  static toObject(includeInstance: boolean, msg: userProfileResponse): userProfileResponse.AsObject;
  static serializeBinaryToWriter(message: userProfileResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): userProfileResponse;
  static deserializeBinaryFromReader(message: userProfileResponse, reader: jspb.BinaryReader): userProfileResponse;
}

export namespace userProfileResponse {
  export type AsObject = {
    message: string,
    profile?: UserUpdateRequest.AsObject,
  }
}

export enum Role { 
  ADMIN = 0,
  USER = 1,
}
