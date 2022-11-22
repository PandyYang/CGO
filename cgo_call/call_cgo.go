package main

/*
#cgo LDFLAGS: -ldl
#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

static void callFromLib() {
    void (*fn)();
    void *h = dlopen("./cgo.so", RTLD_LAZY);
    if (!h) {
        fprintf(stderr, "Error: %s\n", dlerror());
        return;
    }

    *(void**)(&fn) = dlsym(h, "RunLib");
    if (!fn) {
        fprintf(stderr, "Error: %s\n", dlerror());
        dlclose(h);
        return;
    }

    fn();
    dlclose(h);
}

*/
import "C"

func main() {
	C.callFromLib()
}
