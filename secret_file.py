
import os
import base64

xkey = bytearray(os.environ["XKEY"],'UTF-8')

with open("s1.txt") as f:
    content = "".join(f.readlines())
    content = content[:-1] + base64.encodebytes(xkey).decode("UTF-8")
    print(content)

    