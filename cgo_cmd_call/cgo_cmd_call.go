package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

/*
#cgo LDFLAGS: -ldl
#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>
#include <string.h>

void cmd_read(const char *c) {
    void (*go)(char *);
	void *h;
	char *error;

	h = dlopen("./cgo_cmd.so", RTLD_LAZY);
    if (!h) {
        fprintf(stderr, "Error: %s\n", dlerror());
        return;
    }

    go = dlsym(h, "ReadFromCMD");
	if ((error = dlerror()) != NULL)  {
            fputs(error, stderr);
            exit(1);
    }

    go(c);
    dlclose(h);
}

*/
import "C"

func main() {
	fFile := pflag.StringP("file", "F", "", "file to read")
	pflag.Parse()

	if len(*fFile) > 0 {
		f := *fFile
		C.cmd_read(C.CString(f))
	} else {
		fmt.Println("no file to read")
	}
}
