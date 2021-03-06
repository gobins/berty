package main

import (
	"context"
	"flag"
	"fmt"

	"berty.tech/berty/v2/go/cmd/berty/mini"
	"berty.tech/berty/v2/go/internal/config"
	"berty.tech/berty/v2/go/pkg/errcode"
	"berty.tech/go-orbit-db/cache/cacheleveldown"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/shibukawa/configdir"
	"go.uber.org/zap"
)

func miniCommand() *ffcli.Command {
	var miniFlags = flag.NewFlagSet("mini demo client", flag.ExitOnError)
	miniFlags.StringVar(&opts.miniGroup, "g", opts.miniGroup, "group to join, leave empty to create a new group")
	miniFlags.StringVar(&opts.datastorePath, "d", opts.datastorePath, "datastore base directory")
	miniFlags.UintVar(&opts.miniPort, "p", opts.miniPort, "default IPFS listen port")
	miniFlags.StringVar(&opts.remoteDaemonAddr, "r", opts.remoteDaemonAddr, "remote berty daemon")
	miniFlags.StringVar(&opts.rdvpMaddr, "rdvp", opts.rdvpMaddr, "rendezvous point maddr")
	miniFlags.BoolVar(&opts.miniInMemory, "inmem", opts.miniInMemory, "disable persistence")

	return &ffcli.Command{
		Name:      "mini",
		ShortHelp: "start a terminal-based mini berty client (not fully compatible with the app)",
		FlagSet:   miniFlags,
		Exec: func(ctx context.Context, args []string) error {
			cleanup := globalPreRun()
			defer cleanup()

			if !opts.miniInMemory && opts.datastorePath == cacheleveldown.InMemoryDirectory {
				storagePath := configdir.New("Berty", "Mini")
				storageDirs := storagePath.QueryFolders(configdir.Global)
				if len(storageDirs) == 0 {
					return fmt.Errorf("no storage path found")
				}

				if err := storageDirs[0].CreateParentDir(""); err != nil {
					return err
				}

				opts.datastorePath = storageDirs[0].Path
			}

			rootDS, dsLock, err := getRootDatastore(opts.datastorePath)
			if err != nil {
				return errcode.TODO.Wrap(err)
			}
			if dsLock != nil {
				defer func() { _ = dsLock.Unlock() }()
			}
			defer rootDS.Close()

			rdvpeer, err := parseRdvpMaddr(ctx, opts.rdvpMaddr, opts.logger)
			if err != nil {
				return errcode.TODO.Wrap(err)
			}

			l := zap.NewNop()
			if opts.logToFile != "" {
				l = opts.logger
			}

			err = mini.Main(ctx, &mini.Opts{
				RemoteAddr:      opts.remoteDaemonAddr,
				GroupInvitation: opts.miniGroup,
				Port:            opts.miniPort,
				RootDS:          rootDS,
				Logger:          l,
				POIDebug:        opts.poiDebug,
				Bootstrap:       config.BertyDev.Bootstrap,
				RendezVousPeer:  rdvpeer,
				DisplayName:     opts.displayName,
				LocalDiscovery:  opts.localDiscovery,
			})
			if err != nil {
				return errcode.TODO.Wrap(err)
			}
			return nil
		},
	}
}
