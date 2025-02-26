package license

type License struct {
	source        string
	LicenseType   string
	RawText       string
	FormattedText string
}

func NewLicense(source string, licenseType string, rawText string) *License {
	return &License{
		source:      source,
		LicenseType: licenseType,
		RawText:     rawText,
	}
}
