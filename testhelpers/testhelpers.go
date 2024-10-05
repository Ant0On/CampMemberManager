package testhelpers

type TestOptions struct {
	Hosts        []string `envconfig:"HOSTS" default:"127.0.0.1"`
	KeySpace     string   `envconfig:"KEY_SPACE" default:"test_key_space"`
	ProtoVersion int      `envconfig:"PROTO_VERSION" default:"4" required:"true"`
	Consistency  string   `envconfig:"CONSISTENCY" default:"QUORUM"`
}
