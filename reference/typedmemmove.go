// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reference

import "unsafe"

// typedmemmove copies a value of type t to dst from src.
//go:noescape
func typedmemmove(reflectRtype, dst, src unsafe.Pointer)
