# periphflag

This package provides functions for naming hardware buses in flags that produce context-specific
usage messages. For example, instead of using a `flag.StringVar` and getting a usage message like:

    $ foo -help
    Usage of foo:
      -spi string
            spi bus to use

You will instead get a list of devices that are available:

    $ foo -help
    Usage of foo:
      -spi string
            spi bus to use; available devices: [FT232H] (default "FT232H")

This should help people get started with your program a bit more easily.

For information about hardware devices, see [periph.io](https://periph.io/).

For full documentation, read the [godoc](https://godoc.org/github.com/jrockway/periphflag).
