/*
 * Copyright (c) 2019, 2024, ducesoft and/or its affiliates. All rights reserved.
 * DUCESOFT PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 *
 */

package webapp

import (
	"embed"
	"fmt"
	"github.com/opendatav/mesh/client/golang/cause"
	httpx "github.com/opendatav/mesh/client/golang/http"
	"github.com/opendatav/mesh/client/golang/macro"
	"net/http"
	"os"
)

func init() {
	macro.Provide((*http.FileSystem)(nil), &webapp{
		FS: &httpx.StaticFileSystem{
			FS:      home,
			Home:    "static",
			Name:    "webapp.home",
			Pattern: "/mesh",
		},
	})
	//macro.Provide((*http.FileSystem)(nil), &favicon{
	//	FS: &httpx.StaticFileSystem{
	//		FS:      home,
	//		Home:    "static",
	//		Name:    "webapp.favicon",
	//		Pattern: "/favicon.ico",
	//	},
	//})
	//macro.Provide((*http.FileSystem)(nil), &library{
	//	FS: &httpx.StaticFileSystem{
	//		FS:      home,
	//		Home:    "static",
	//		Name:    "webapp.library",
	//		Pattern: "/lib",
	//	},
	//})
}

//go:embed static
var home embed.FS

type webapp struct {
	FS httpx.VFS
}

func (that *webapp) Att() *macro.Att {
	return that.FS.Att()
}

func (that *webapp) Open(name string) (http.File, error) {
	if f, err := that.FS.Open(name); os.IsNotExist(cause.DeError(err)) {
		return that.FS.Open("/")
	} else {
		return f, err
	}
}

type favicon struct {
	FS httpx.VFS
}

func (that *favicon) Att() *macro.Att {
	return that.FS.Att()
}

func (that *favicon) Open(name string) (http.File, error) {
	return that.FS.Open(name)
}

type library struct {
	FS httpx.VFS
}

func (that *library) Att() *macro.Att {
	return that.FS.Att()
}

func (that *library) Open(name string) (http.File, error) {
	return that.FS.Open(fmt.Sprintf("/lib%s", name))
}
