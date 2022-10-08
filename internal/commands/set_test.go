package commands

import (
	"context"
	"github.com/shoenig/test/must"
	"os"
	"testing"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/shoenig/envy/internal/keyring"
	"github.com/shoenig/envy/internal/output"
	"github.com/shoenig/envy/internal/safe"
	"github.com/shoenig/envy/internal/setup"
	"github.com/shoenig/secrets"
)

func TestSetCmd_Ops(t *testing.T) {

	db := newDBFile(t)
	defer cleanupFile(t, db)

	w := output.New(os.Stdout, os.Stdout)
	cmd := NewSetCmd(setup.New(db, w))

	must.Eq(t, setCmdName, cmd.Name())
	must.Eq(t, setCmdSynopsis, cmd.Synopsis())
	must.Eq(t, setCmdUsage, cmd.Usage())
}

func TestSetCmd_Execute(t *testing.T) {

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	ring.EncryptMock.When(secrets.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(secrets.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	box.SetMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{8, 8, 8, 8, 8, 8},
			"bar": safe.Encrypted{9, 9, 9, 9},
		},
	}).Return(nil)

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"set", "myNS", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitSuccess, rc)
	must.Eq(t, "", b.String())
	must.Eq(t, "stored 2 items in \"myNS\"\n", a.String())
}

func TestSetCmd_Execute_ioError(t *testing.T) {
	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	box.SetMock.Expect(&safe.Namespace{
		Name: "myNS",
		Content: map[string]safe.Encrypted{
			"foo": safe.Encrypted{8, 8, 8, 8, 8, 8},
			"bar": safe.Encrypted{9, 9, 9, 9},
		},
	}).Return(errors.New("io error"))

	ring.EncryptMock.When(secrets.New("abc123")).Then(safe.Encrypted{8, 8, 8, 8, 8, 8})
	ring.EncryptMock.When(secrets.New("1234")).Then(safe.Encrypted{9, 9, 9, 9})

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	fs, args := setupFlagSet(t, []string{"set", "myNS", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitFailure, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to update namespace: io error\n", b.String())
}

func TestSetCmd_Execute_badNS(t *testing.T) {

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	// e.g. forgot to specify namespace
	fs, args := setupFlagSet(t, []string{"set", "foo=abc123", "bar=1234"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: unable to parse args: namespace uses non-word characters\n", b.String())
}

func TestSetCmd_Execute_noVars(t *testing.T) {

	box := safe.NewBoxMock(t)
	defer box.MinimockFinish()

	ring := keyring.NewRingMock(t)
	defer ring.MinimockFinish()

	a, b, w := newWriter()

	pc := &setCmd{
		writer: w,
		ex:     newExtractor(ring),
		box:    box,
	}

	// e.g. reminder to use purge to remove namespace
	fs, args := setupFlagSet(t, []string{"set", "ns1"})
	pc.SetFlags(fs)
	ctx := context.Background()
	rc := pc.Execute(ctx, fs, args)

	must.Eq(t, subcommands.ExitUsageError, rc)
	must.Eq(t, "", a.String())
	must.Eq(t, "envy: use 'purge' to remove namespace\n", b.String())
}
