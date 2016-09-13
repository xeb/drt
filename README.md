# drt

**Docker Run Tool** (or **drt** pronounced *dirt*)is a program designed to allow running CLI utilities which are distributed as containers.  

# Example
Let's assume that you want to run a small [TensorFlow](https://github.com/tensorflow/tensorflow) script -- like [tf_script.py](samples/tensorflow/tf_script.py), but you don't want to download all the TensorFlow dependencies.  Assuming you at least have [Docker](http://docker.com) running on your system, you can use **drt** to run a TensorFlow program.

```

$ cd samples/tensorflow

$ drt run drt.yaml tf_script.py
docker run --rm -v /Users/xeb/.marks/gocode/drt:/pwd tensorflow/tensorflow:latest /usr/bin/python /pwd/samples/tensorflow/tf_script.py
Unable to find image 'tensorflow/tensorflow:latest' locally
latest: Pulling from tensorflow/tensorflow
759d6771041e: Downloading [=======================================>           ] 52.43 MB/65.69 MB
8836b825667b: Download complete
c2f5e51744e6: Download complete
a3ed95caeb02: Download complete
b1a230d2f7d7: Downloading [============================>                      ] 65.96 MB/117.2 MB
f4bd848af23b: Download complete
9ff2d28f3bea: Downloading [===========================>                       ] 37.31 MB/66.78 MB
de3ac9a732b6: Waiting
ba1ff7fd8338: Waiting
8b638dd1001b: Waiting
4c3f9ea22d7e: Waiting
ba1ff7fd8338: Pull complete
8b638dd1001b: Pull complete
4c3f9ea22d7e: Pull complete
...
Digest: sha256:8ae5229583adf18c1e50ac7fbf4c25301aef186eb3cc7c23d85999328a72ac4b
Status: Downloaded newer image for tensorflow/tensorflow:latest
[[19 22]
 [43 50]]
$

```

At least that is the idea.  I'll probably start working on registering "run manifests" in Git (like Homebrew does) and building independent binaries so you can ship YOUR OWN CLIs with just a docker image and a small "drt-based" binary
