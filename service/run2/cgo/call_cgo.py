from ctypes import cdll
cur = cdll.LoadLibrary('/root/goProject/src/CGO/service/run2/cgo/libsdk.so')
cur.start()
cur.search()