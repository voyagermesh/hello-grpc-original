#!/usr/bin/python

import os
import sys


def reindex(root):
    """root: path to the root of the git root."""
    apis = []
    for dir, _, files in os.walk(root + '/apis'):
        for file in files:
            rel = os.path.relpath(dir + '/' + file, root)
            # print rel
            apis.append(rel)

    content = """// Code generated by ./hack/browserify.py
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var _ = require('lodash');

"""
    # module exports
    content += 'var apis = _.merge({},\n'
    for api in sorted(apis):
        content += "    require('./{}'),\n".format(api)
    content = content[:-2]
    content += '\n);\n'
    content += 'module.exports = apis.appscode.hello;\n'
    with open(root + '/index.js', 'w') as f:
        return f.write(content)


if __name__ == "__main__":
    if len(sys.argv) > 1:
        reindex(os.path.abspath(sys.argv[1]))
    else:
        reindex(os.path.expandvars('$GOPATH') + '/src/voyagermesh.dev/hello-grpc/client/js')
