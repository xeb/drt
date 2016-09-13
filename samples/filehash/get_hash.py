#!/usr/bin/env python
import os,sys
import hashlib

def get_file_hash(file):
    BLOCKSIZE = 65536
    hasher = hashlib.md5()
    with open(file, 'rb') as afile:
        buf = afile.read(BLOCKSIZE)
        while len(buf) > 0:
            hasher.update(buf)
            buf = afile.read(BLOCKSIZE)
        return hasher.hexdigest()

if __name__ == "__main__":
    print(get_file_hash(sys.argv[1]))
