// source: infrastructure/rpc/cqrs/link/v1/link_query.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = (function() { return this || window || global || self || Function('return this')(); }).call(null);

var domain_link_cqrs_v1_link_pb = require('../../../../../domain/link_cqrs/v1/link_pb.js');
goog.object.extend(proto, domain_link_cqrs_v1_link_pb);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.GetRequest', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.GetResponse', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.ListRequest', null, global);
goog.exportSymbol('proto.infrastructure.rpc.cqrs.link.v1.ListResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.GetRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.GetRequest.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.GetRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.GetResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.GetResponse.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.GetResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.ListRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.ListRequest.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.ListRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.infrastructure.rpc.cqrs.link.v1.ListResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.infrastructure.rpc.cqrs.link.v1.ListResponse.displayName = 'proto.infrastructure.rpc.cqrs.link.v1.ListResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.GetRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    hash: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.GetRequest;
  return proto.infrastructure.rpc.cqrs.link.v1.GetRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setHash(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.GetRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getHash();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string hash = 1;
 * @return {string}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.prototype.getHash = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetRequest} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.GetRequest.prototype.setHash = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.GetResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    link: (f = msg.getLink()) && domain_link_cqrs_v1_link_pb.LinkView.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.GetResponse;
  return proto.infrastructure.rpc.cqrs.link.v1.GetResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_cqrs_v1_link_pb.LinkView;
      reader.readMessage(value,domain_link_cqrs_v1_link_pb.LinkView.deserializeBinaryFromReader);
      msg.setLink(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.GetResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLink();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_cqrs_v1_link_pb.LinkView.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link_cqrs.v1.LinkView link = 1;
 * @return {?proto.domain.link_cqrs.v1.LinkView}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.getLink = function() {
  return /** @type{?proto.domain.link_cqrs.v1.LinkView} */ (
    jspb.Message.getWrapperField(this, domain_link_cqrs_v1_link_pb.LinkView, 1));
};


/**
 * @param {?proto.domain.link_cqrs.v1.LinkView|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.setLink = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.GetResponse} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.clearLink = function() {
  return this.setLink(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.GetResponse.prototype.hasLink = function() {
  return jspb.Message.getField(this, 1) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.ListRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    filter: jspb.Message.getFieldWithDefault(msg, 1, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.ListRequest;
  return proto.infrastructure.rpc.cqrs.link.v1.ListRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilter(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.ListRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilter();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string filter = 1;
 * @return {string}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.prototype.getFilter = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListRequest} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.ListRequest.prototype.setFilter = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.infrastructure.rpc.cqrs.link.v1.ListResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    links: (f = msg.getLinks()) && domain_link_cqrs_v1_link_pb.LinksView.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.infrastructure.rpc.cqrs.link.v1.ListResponse;
  return proto.infrastructure.rpc.cqrs.link.v1.ListResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new domain_link_cqrs_v1_link_pb.LinksView;
      reader.readMessage(value,domain_link_cqrs_v1_link_pb.LinksView.deserializeBinaryFromReader);
      msg.setLinks(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.infrastructure.rpc.cqrs.link.v1.ListResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getLinks();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      domain_link_cqrs_v1_link_pb.LinksView.serializeBinaryToWriter
    );
  }
};


/**
 * optional domain.link_cqrs.v1.LinksView links = 1;
 * @return {?proto.domain.link_cqrs.v1.LinksView}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.getLinks = function() {
  return /** @type{?proto.domain.link_cqrs.v1.LinksView} */ (
    jspb.Message.getWrapperField(this, domain_link_cqrs_v1_link_pb.LinksView, 1));
};


/**
 * @param {?proto.domain.link_cqrs.v1.LinksView|undefined} value
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse} returns this
*/
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.setLinks = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.infrastructure.rpc.cqrs.link.v1.ListResponse} returns this
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.clearLinks = function() {
  return this.setLinks(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.infrastructure.rpc.cqrs.link.v1.ListResponse.prototype.hasLinks = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.infrastructure.rpc.cqrs.link.v1);
