// package: infrastructure.rpc.cqrs.link.v1
// file: infrastructure/rpc/cqrs/link/v1/link_query.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as domain_link_cqrs_v1_link_pb from "../../../../../domain/link_cqrs/v1/link_pb";

export class GetRequest extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: GetRequest): GetRequest.AsObject;

    static serializeBinaryToWriter(message: GetRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): GetRequest;

    static deserializeBinaryFromReader(message: GetRequest, reader: jspb.BinaryReader): GetRequest;

    getHash(): string;

    setHash(value: string): GetRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): GetRequest.AsObject;
}

export namespace GetRequest {
    export type AsObject = {
        hash: string,
    }
}

export class GetResponse extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: GetResponse): GetResponse.AsObject;

    static serializeBinaryToWriter(message: GetResponse, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): GetResponse;

    static deserializeBinaryFromReader(message: GetResponse, reader: jspb.BinaryReader): GetResponse;

    hasLink(): boolean;

    clearLink(): void;

    getLink(): domain_link_cqrs_v1_link_pb.LinkView | undefined;

    setLink(value?: domain_link_cqrs_v1_link_pb.LinkView): GetResponse;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): GetResponse.AsObject;
}

export namespace GetResponse {
    export type AsObject = {
        link?: domain_link_cqrs_v1_link_pb.LinkView.AsObject,
    }
}

export class ListRequest extends jspb.Message {
    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: ListRequest): ListRequest.AsObject;

    static serializeBinaryToWriter(message: ListRequest, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): ListRequest;

    static deserializeBinaryFromReader(message: ListRequest, reader: jspb.BinaryReader): ListRequest;

    getFilter(): string;

    setFilter(value: string): ListRequest;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): ListRequest.AsObject;
}

export namespace ListRequest {
    export type AsObject = {
        filter: string,
    }
}

export class ListResponse extends jspb.Message {

    static extensions: { [key: number]: jspb.ExtensionFieldInfo<jspb.Message> };
    static extensionsBinary: { [key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message> };

    static toObject(includeInstance: boolean, msg: ListResponse): ListResponse.AsObject;

    static serializeBinaryToWriter(message: ListResponse, writer: jspb.BinaryWriter): void;

    static deserializeBinary(bytes: Uint8Array): ListResponse;

    static deserializeBinaryFromReader(message: ListResponse, reader: jspb.BinaryReader): ListResponse;

    hasLinks(): boolean;

    clearLinks(): void;

    getLinks(): domain_link_cqrs_v1_link_pb.LinksView | undefined;

    setLinks(value?: domain_link_cqrs_v1_link_pb.LinksView): ListResponse;

    serializeBinary(): Uint8Array;

    toObject(includeInstance?: boolean): ListResponse.AsObject;
}

export namespace ListResponse {
    export type AsObject = {
        links?: domain_link_cqrs_v1_link_pb.LinksView.AsObject,
    }
}
