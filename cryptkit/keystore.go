package cryptkit

//go:generate minimock -i github.com/insolar/vanilla/cryptkit.PublicKeyStore -o . -s _mock.go -g

type PublicKeyStore interface {
	PublicKeyStore()
}

type SecretKeyStore interface {
	PrivateKeyStore()
	AsPublicKeyStore() PublicKeyStore
}

//go:generate minimock -i github.com/insolar/vanilla/cryptkit.KeyStoreFactory -o . -s _mock.go -g

type KeyStoreFactory interface {
	CreatePublicKeyStore(SigningKeyHolder) PublicKeyStore
}
