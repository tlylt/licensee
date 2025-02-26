package license

type License struct {
	LicenseType   string
	RawText       string
	FormattedText string
}

func NewLicense(licenseType string, rawText string) *License {
	return &License{}
}
