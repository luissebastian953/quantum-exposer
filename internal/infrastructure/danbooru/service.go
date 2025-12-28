package danbooru

import (
	"os"
	"quantum-exposer/internal/infrastructure/network"
	"quantum-exposer/internal/usecase"
	"time"
)

// Cloudflare's public DNS server
const CustomDNSServer = "1.1.1.1:53"

func InitializeDanbooruService() usecase.PostRepository {
	return NewDanbooruAPIRepository(
		network.NewConfiguratedHTTPClient(30*time.Second, CustomDNSServer),
		os.Getenv("EXPOSING_URL"),
		os.Getenv("EXPOSING_USERNAME"),
		os.Getenv("EXPOSING_API_KEY"),
	)
}
