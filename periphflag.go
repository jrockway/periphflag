// Package periphflag integrates periph.io autodiscovery with flags, so that "yourcommand -help"
// makes more sense to the end user.
package periphflag

import (
	"flag"
	"fmt"

	"periph.io/x/periph/conn/spi/spireg"
)

// SPIDevVar defines a flag that acceps the name of a SPI bus, with some information that is
// displayed to the end-user replaced with real discovered values.  The arguments are the same as a
// standard flag.StringVar.  The default value is looked up against discovered buses and their
// aliases, and the real device name is substituted.  The usage message is changed to have a list of
// discovered bus names after the usage message.
//
// In order to populate devices, you must have already initialized periph.io with host.Init() or
// hostextra.Init().  Because you need to check the error from that before proceeding, you will
// typically not declare these variables at the package level, and instead declare them near the top
// of your main function:
//
//   package main
//
//   var (
//       someFlag = flag.String("send", "foo", "string to send to the device")
//       spi string
//   )
//
//   func main() {
//       if _, err := host.Init(); err != nil {
//           log.Fatal("periph init: %v", err)
//       }
//       periphflag.SPIDevVar(&spi, "spi", "", "spi bus to use")
//       flag.Parse()
//
//       port, err := spireg.Open(spi)
//       if err != nil {
//           flag.Usage()
//           log.Fatal("open spi port %q: %v", spi, err)
//       }
//       defer port.Close()
//       ...
//   }
func SPIDevVar(p *string, name, value, usage string) {
	SPIDevVarOnFlagSet(flag.CommandLine, p, name, value, usage)
}

// SPIDevVarOnFlagSet defines a flag in a manner identical to SPIDevVar on the provided flag.FlagSet.
func SPIDevVarOnFlagSet(f *flag.FlagSet, p *string, name, value, usage string) {
	devices := spireg.All()
	var def string
	var names []string
	for _, d := range devices {
		names = append(names, d.Name)
		names = append(names, d.Aliases...)
		if def == "" {
			if d.Name == value {
				def = d.Name
			} else {
				for _, a := range d.Aliases {
					if a == value {
						def = d.Name
						break
					}
				}
			}
		}
	}
	if def == "" && len(names) > 0 {
		def = names[0]
	}
	f.StringVar(p, name, def, fmt.Sprintf("%s; available devices: %v", usage, names))
}
