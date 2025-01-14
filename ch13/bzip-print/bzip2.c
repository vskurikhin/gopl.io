// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 362.
// This is the version that appears in print,
// but it does not comply with the proposed
// rules for passing pointers between Go and C.
// (https://github.com/golang/proposal/blob/master/design/12416-cgo-pointers.md)
// See gopl.io/ch13/bzip for an updated version.
//
// patched by Victor N. Skurikhin.
// Reason:
// link: running cc failed: exit status 1
// ld: warning: -no_pie is deprecated when targeting new OS versions
// duplicate symbol '_bz2compress'
// ld: 1 duplicate symbol for architecture x86_64
// clang: error: linker command failed with exit code 1

//!+
/* This file is gopl.io/ch13/bzip/bzip2.c,         */
/* a simple wrapper for libbzip2 suitable for cgo. */
#include <bzlib.h>

int bz2compress_res_confl(bz_stream *s, int action,
                char *in, unsigned *inlen, char *out, unsigned *outlen) {
  s->next_in = in;
  s->avail_in = *inlen;
  s->next_out = out;
  s->avail_out = *outlen;
  int r = BZ2_bzCompress(s, action);
  *inlen -= s->avail_in;
  *outlen -= s->avail_out;
  return r;
}

//!-
