// source: order.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
goog.object.extend(proto, google_protobuf_empty_pb);
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
goog.exportSymbol('proto.zwan.Order', null, global);
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
proto.zwan.Order = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.zwan.Order, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.zwan.Order.displayName = 'proto.zwan.Order';
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
proto.zwan.Order.prototype.toObject = function(opt_includeInstance) {
  return proto.zwan.Order.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.zwan.Order} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zwan.Order.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: jspb.Message.getFieldWithDefault(msg, 1, ""),
    hospital: jspb.Message.getFieldWithDefault(msg, 2, ""),
    department: jspb.Message.getFieldWithDefault(msg, 3, ""),
    targetTime: (f = msg.getTargetTime()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    patient: jspb.Message.getFieldWithDefault(msg, 5, ""),
    companion: jspb.Message.getFieldWithDefault(msg, 6, ""),
    notes: jspb.Message.getFieldWithDefault(msg, 7, ""),
    created: (f = msg.getCreated()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
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
 * @return {!proto.zwan.Order}
 */
proto.zwan.Order.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.zwan.Order;
  return proto.zwan.Order.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.zwan.Order} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.zwan.Order}
 */
proto.zwan.Order.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setHospital(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setDepartment(value);
      break;
    case 4:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTargetTime(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setPatient(value);
      break;
    case 6:
      var value = /** @type {string} */ (reader.readString());
      msg.setCompanion(value);
      break;
    case 7:
      var value = /** @type {string} */ (reader.readString());
      msg.setNotes(value);
      break;
    case 8:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setCreated(value);
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
proto.zwan.Order.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.zwan.Order.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.zwan.Order} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.zwan.Order.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getHospital();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getDepartment();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getTargetTime();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getPatient();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getCompanion();
  if (f.length > 0) {
    writer.writeString(
      6,
      f
    );
  }
  f = message.getNotes();
  if (f.length > 0) {
    writer.writeString(
      7,
      f
    );
  }
  f = message.getCreated();
  if (f != null) {
    writer.writeMessage(
      8,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string id = 1;
 * @return {string}
 */
proto.zwan.Order.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string hospital = 2;
 * @return {string}
 */
proto.zwan.Order.prototype.getHospital = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setHospital = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string department = 3;
 * @return {string}
 */
proto.zwan.Order.prototype.getDepartment = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setDepartment = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional google.protobuf.Timestamp target_time = 4;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.zwan.Order.prototype.getTargetTime = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 4));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.zwan.Order} returns this
*/
proto.zwan.Order.prototype.setTargetTime = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.clearTargetTime = function() {
  return this.setTargetTime(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.zwan.Order.prototype.hasTargetTime = function() {
  return jspb.Message.getField(this, 4) != null;
};


/**
 * optional string patient = 5;
 * @return {string}
 */
proto.zwan.Order.prototype.getPatient = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setPatient = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional string companion = 6;
 * @return {string}
 */
proto.zwan.Order.prototype.getCompanion = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setCompanion = function(value) {
  return jspb.Message.setProto3StringField(this, 6, value);
};


/**
 * optional string notes = 7;
 * @return {string}
 */
proto.zwan.Order.prototype.getNotes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 7, ""));
};


/**
 * @param {string} value
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.setNotes = function(value) {
  return jspb.Message.setProto3StringField(this, 7, value);
};


/**
 * optional google.protobuf.Timestamp created = 8;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.zwan.Order.prototype.getCreated = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 8));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.zwan.Order} returns this
*/
proto.zwan.Order.prototype.setCreated = function(value) {
  return jspb.Message.setWrapperField(this, 8, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.zwan.Order} returns this
 */
proto.zwan.Order.prototype.clearCreated = function() {
  return this.setCreated(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.zwan.Order.prototype.hasCreated = function() {
  return jspb.Message.getField(this, 8) != null;
};


goog.object.extend(exports, proto.zwan);