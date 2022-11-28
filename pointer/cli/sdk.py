import ctypes
so = ctypes.CDLL('/root/goProject/src/NTI-SDK/test/cmd/v3/cli/libsdk.so')

c = so.initSDK("/root/goProject/src/NTI-SDK/conf/config.ini")
queryRes = so.query(c, "123")
so.destroy(c)
