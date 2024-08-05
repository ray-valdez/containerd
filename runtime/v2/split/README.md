# Split Shim

## Introduction
Split is a containerd shim for driving the split api from the owner side. 

## Building 

### Build Containerd
```bash
$ GODEBUG=1 make bin/containerd  BUILDTAGS="no_btrfs"
```

Install `bin/containerd` in your path

### Build the split-shim
To create the container-shim-split-v2 executable
```bash
$ make
```

To install in `/usr/local/bin`
```bash
$ sudo make install
```

Add split-shim to `containerd`'s configuration file
```
[plugins]
  [plugins."io.containerd.grpc.v1.cri"]
    [plugins."io.containerd.grpc.v1.cri".containerd]
      [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.split]
          runtime_type = "io.containerd.split.v2"
```