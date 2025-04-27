package config

type HTTPServer struct
{
	Address string
}

type Config struct{
	Env string
	StoragePath string
	HTTPServer

}