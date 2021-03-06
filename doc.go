// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

/*
Package termui is a library designed for creating command line UI. For more info, goto http://github.com/xianwangs/termui

A simplest example:
    package main

    import ui "github.com/xianwangs/termui"

    func main() {
        if err:=ui.Init(); err != nil {
            panic(err)
        }
        defer ui.Close()

        g := ui.NewGauge()
        g.Percent = 50
        g.Width = 50
        g.BorderLabel = "Gauge"

        ui.Render(g)

        ui.Loop()
    }
*/
package termui
