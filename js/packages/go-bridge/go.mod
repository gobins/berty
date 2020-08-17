go 1.14

module berty.tech/berty/v2/js/packages/go-bridge

replace berty.tech/berty/v2/go => ../../../go

require golang.org/x/mobile v0.0.0-20200629153529-33b80540585f

replace (
	bazil.org/fuse => bazil.org/fuse v0.0.0-20200117225306-7b5117fecadc // specific version for iOS building
	github.com/agl/ed25519 => github.com/agl/ed25519 v0.0.0-20170116200512-5312a6153412 // latest commit before the author shutdown the repo; see https://github.com/golang/go/issues/20504
	github.com/ipld/go-ipld-prime => github.com/ipld/go-ipld-prime v0.0.2-0.20191108012745-28a82f04c785 // specific version needed indirectly
	github.com/ipld/go-ipld-prime-proto => github.com/ipld/go-ipld-prime-proto v0.0.0-20191113031812-e32bd156a1e5 // specific version needed indirectly
	github.com/libp2p/go-libp2p-kbucket => github.com/libp2p/go-libp2p-kbucket v0.4.2 // specific version needed indirectly
)
