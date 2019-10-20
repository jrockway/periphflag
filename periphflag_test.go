package periphflag

import (
	"errors"
	"flag"
	"reflect"
	"testing"

	"periph.io/x/periph/conn"
	"periph.io/x/periph/conn/physic"
	"periph.io/x/periph/conn/spi"
	"periph.io/x/periph/conn/spi/spireg"
)

type fakePort struct{}
type fakeConn struct{}

func (p fakePort) String() string                      { return "fake spi" }
func (p fakePort) Close() error                        { return nil }
func (p fakePort) LimitSpeed(f physic.Frequency) error { return nil }
func (p fakePort) Connect(f physic.Frequency, mode spi.Mode, bits int) (spi.Conn, error) {
	return fakeConn{}, nil
}
func (c fakeConn) String() string                 { return "fake connection" }
func (c fakeConn) TxPackets(p []spi.Packet) error { return errors.New("not implemented") }
func (c fakeConn) Tx(w, r []byte) error           { return errors.New("not implemented") }
func (c fakeConn) Duplex() conn.Duplex            { return conn.DuplexUnknown }

func TestSPIDevVar(t *testing.T) {
	spireg.Register("foo", []string{"bar", "baz"}, 1, func() (spi.PortCloser, error) { return fakePort{}, nil })
	fs := new(flag.FlagSet)
	var spi string
	SPIDevVarOnFlagSet(fs, &spi, "spi", "bar", "spi bus to use")
	if err := fs.Parse(nil); err != nil {
		t.Fatalf("parse: %v", err)
	}

	var flags []flag.Flag
	fs.VisitAll(func(orig *flag.Flag) { f := *orig; f.Value = nil; flags = append(flags, f) })

	if got, want := flags, []flag.Flag{{Name: "spi", Usage: "spi bus to use; available devices: [foo bar baz]", Value: nil, DefValue: "foo"}}; !reflect.DeepEqual(got, want) {
		t.Errorf("flags:\n  got: %#v\n want: %#v", got, want)
	}
}
