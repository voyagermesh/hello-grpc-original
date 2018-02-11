#!/usr/bin/env python


# http://stackoverflow.com/a/14050282
def check_antipackage():
    from sys import version_info
    sys_version = version_info[:2]
    found = True
    if sys_version < (3, 0):
        # 'python 2'
        from pkgutil import find_loader
        found = find_loader('antipackage') is not None
    elif sys_version <= (3, 3):
        # 'python <= 3.3'
        from importlib import find_loader
        found = find_loader('antipackage') is not None
    else:
        # 'python >= 3.4'
        from importlib import util
        found = util.find_spec('antipackage') is not None
    if not found:
        print('Install missing package "antipackage"')
        print('Example: pip install git+https://github.com/ellisonbg/antipackage.git#egg=antipackage')
        from sys import exit
        exit(1)
check_antipackage()

# ref: https://github.com/ellisonbg/antipackage
import antipackage
from github.appscode.libbuild import libbuild

import datetime
import io
import json
import os
import os.path
import shutil
import socket
import subprocess
import sys
import tempfile
from collections import OrderedDict
from os.path import expandvars


# Debian package
# https://gist.github.com/rcrowley/3728417
libbuild.REPO_ROOT = expandvars('$GOPATH') + '/src/github.com/appscode/hello-grpc'
BUILD_METADATA = libbuild.metadata(libbuild.REPO_ROOT)
libbuild.BIN_MATRIX = {
    'hello-grpc': {
        'type': 'go',
        'go_version': True,
        'distro': {
            'alpine': ['amd64']
        }
    },
}
libbuild.BUCKET_MATRIX = {
    'dev': 'gs://appscode-dev'
}


def call(cmd, stdin=None, cwd=libbuild.REPO_ROOT):
    print(cmd)
    return subprocess.call([expandvars(cmd)], shell=True, stdin=stdin, cwd=cwd)


def die(status):
    if status:
        sys.exit(status)


def version():
    # json.dump(BUILD_METADATA, sys.stdout, sort_keys=True, indent=2)
    for k in sorted(BUILD_METADATA):
        print(k + '=' + BUILD_METADATA[k])


def fmt():
    libbuild.ungroup_go_imports('pkg', '*.go')
    die(call('goimports -w pkg *.go'))
    call('gofmt -s -w pkg *.go')


def vet():
    call('go vet ./pkg/... *.go')


def gen_protos():
    # Generate protos
    die(call('./hack/make.sh', cwd=libbuild.REPO_ROOT + '/_proto'))
    #Move generated go files to api.
    call('rm -rf apis', cwd=libbuild.REPO_ROOT + '/pkg')
    shutil.copytree(libbuild.REPO_ROOT + '/_proto', libbuild.REPO_ROOT + '/pkg/apis', ignore=ignore_most)
    call('find . -type d -empty -delete', cwd=libbuild.REPO_ROOT + '/pkg/apis')
    call("find . -type f -name '*.go' -delete", cwd=libbuild.REPO_ROOT + '/_proto')


def gen_js_client():
    die(call('./hack/make.sh', cwd=libbuild.REPO_ROOT + '/client/js'))


def gen_extpoints():
    die(call('go generate main.go'))


def gen():
    gen_protos()
    gen_js_client()
    gen_extpoints()


def ignore_most(folder, files):
    ignore_list = []
    for file in files:
        full_path = os.path.join(folder, file)
        if not os.path.isdir(full_path):
            if not file.endswith(".go"):
                ignore_list.append(file)
    return ignore_list


def build_cmd(name):
    cfg = libbuild.BIN_MATRIX[name]
    if cfg['type'] == 'go':
        if 'distro' in cfg:
            for goos, archs in cfg['distro'].items():
                for goarch in archs:
                    libbuild.go_build(name, goos, goarch, main='*.go')
        else:
            libbuild.go_build(name, libbuild.GOHOSTOS, libbuild.GOHOSTARCH, main='*.go')


def build_cmds():
    gen()
    fmt()
    for name in libbuild.BIN_MATRIX.keys():
        build_cmd(name)


def build(name=None):
    if name:
        cfg = libbuild.BIN_MATRIX[name]
        if cfg['type'] == 'go':
            gen()
            fmt()
            build_cmd(name)
    else:
        build_cmds()


def push(name=None):
    if name:
        bindir = libbuild.REPO_ROOT + '/dist/' + name
        push_bin(bindir)
    else:
        dist = libbuild.REPO_ROOT + '/dist'
        for name in os.listdir(dist):
            d = dist + '/' + name
            if os.path.isdir(d):
                push_bin(d)


def push_bin(bindir):
    call('rm -f *.md5', cwd=bindir)
    call('rm -f *.sha1', cwd=bindir)
    for f in os.listdir(bindir):
        if os.path.isfile(bindir + '/' + f):
            libbuild.upload_to_cloud(bindir, f, BUILD_METADATA['version'])


def update_registry():
    libbuild.update_registry(BUILD_METADATA['version'])


def install():
    die(call('GO15VENDOREXPERIMENT=1 ' + libbuild.GOC + ' install .'))


def default():
    gen()
    fmt()
    die(call('GO15VENDOREXPERIMENT=1 ' + libbuild.GOC + ' install .'))


if __name__ == "__main__":
    if len(sys.argv) > 1:
        # http://stackoverflow.com/a/834451
        # http://stackoverflow.com/a/817296
        globals()[sys.argv[1]](*sys.argv[2:])
    else:
        default()
