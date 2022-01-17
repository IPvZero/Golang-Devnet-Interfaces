package main

import (
        "encoding/json"
        "flag"
        "fmt"

        "github.com/scrapli/scrapligo/driver/base"
        "github.com/scrapli/scrapligo/driver/core"
)

func main() {
        arg := flag.String(
                "file",
                "intf.fsm",
                "_",
        )
        flag.Parse()

        d, err := core.NewCoreDriver(
                "ios-xe-mgmt.cisco.com",
                "cisco_iosxe",
                base.WithPort(22),
                base.WithAuthStrictKey(false),
                base.WithAuthUsername("developer"),
                base.WithAuthPassword("C1sco12345"),
                base.WithTransportType("standard"),
        )

        if err != nil {
                fmt.Printf("failed to create driver; error: %+v\n", err)
                return
        }

        err = d.Open()
        if err != nil {
                fmt.Printf("failed to open driver; error: %+v\n", err)
                return
        }
        defer d.Close()

        r, err := d.SendCommand("show interfaces")
        if err != nil {
                fmt.Printf("failed to send command; error: %+v\n", err)
                return
        }

        parsedOut, err := r.TextFsmParse(*arg)
        if err != nil {
                fmt.Printf("failed to parse command; error: %+v\n", err)
                return
        }
        //fmt.Println(parsedOut)
        s, _ := json.MarshalIndent(parsedOut, "", "\t")
        fmt.Print(string(s))

        fmt.Println("\n****\n")
        for i := 0; i < len(parsedOut); i++ {
                fmt.Println("INTERFACE:", parsedOut[i]["INTERFACE"], "IP:", parsedOut[i]["IP_ADDRESS"])
        }

}
