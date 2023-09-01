package config

type Config struct {
	PrivateKeyHex    string `mapstructure:"privateKey"`
	EthereumURL      string `mapstructure:"ethereumUrl"`
	EthAmount        string `mapstructure:"ethSendAmount"`
	TokenAmount      string `mapstructure:"tokenSendAmount"`
	TokenAddress     string `mapstructure:"tokenAddress"`
	GrpcServerUrl    string `mapstructure:"grpcServerUrl"`
	GatewayServerUrl string `mapstructure:"gatewayServerUrl"`
	MongoUri         string `mapstructure:"mongoUri"`
	Limit            uint16 `mapstructure:"limit"`
	LimitUnit        string `mapstructure:"limitUnit"`
}
