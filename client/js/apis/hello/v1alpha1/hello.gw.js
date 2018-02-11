// Code generated by protoc-gen-grpc-js-client
// source: hello.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function helloInto(p, conf) {
    path = '/apis/hello/v1alpha1/into/json'
    return xhr(path, 'GET', conf, p);
}

var services = {
    hello: {
        into: helloInto
    }
};

module.exports = {appscode: {hello: {v1alpha1: services}}};
