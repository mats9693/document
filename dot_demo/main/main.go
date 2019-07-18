package main

import (
    "fmt"
    "github.com/scryinfo/dot/dot"
    "github.com/scryinfo/dot/dots/line"
    "github.com/scryinfo/dp_app_dot_demo/dots/summary"
    "github.com/scryinfo/dp_app_dot_demo/dots/summary/sw"
    "github.com/scryinfo/scryg/sutils/ssignal"
    "go.uber.org/zap"
    "os"
)

func main() {
    var logger = dot.Logger()
    l, err := line.BuildAndStart(func(l dot.Line) error {
        l.PreAdd(summary.SumTypeLive()...)
        return nil
    })

    if err != nil {
        dot.Logger().Debugln("Line init failed. ", zap.NamedError("", err))
        return
    }

    ssignal.WaitCtrlC(func(s os.Signal) bool {
        return false //quit
    })
}
