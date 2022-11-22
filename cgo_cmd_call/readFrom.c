#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

static void cmd_read(const char* c) {
    void (*fn)();
    void *h = dlopen("./cgo_cmd.so", RTLD_LAZY);
    if (!h) {
        fprintf(stderr, "Error: %s\n", dlerror());
        return;
    }

    *(void**)(&fn) = dlsym(h, "ReadFromCMD");
    if (!fn) {
        fprintf(stderr, "Error: %s\n", dlerror());
        dlclose(h);
        return;
    }
    fn(c);
    dlclose(h);
}
