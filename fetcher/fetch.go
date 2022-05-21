package fetcher

const (
	urlPrefix  = ``
	waitSecond = 1
)

// Fetch scrape target
func Fetch() error {

	return runChromedp(false)
}
