#!/usr/bin/env python
import tensorflow as tf

m1 = tf.constant([[1,2],[3,4]])
m2 = tf.constant([[5,6],[7,8]])
mout = tf.matmul(m1, m2)
with tf.Session() as sess:
    result = sess.run(mout)
    print(result)
    sess.close()
