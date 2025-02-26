package analyzer

import "strings"

var licenseTypes = []string{
	"MIT",
	"Apache",
	"GPL",
	"BSD",
	"ISC",
	"Unlicense",
}

func LicenseType(rawText string) string {
	for _, licenseType := range licenseTypes {
		if strings.Contains(rawText, licenseType) {
			return licenseType
		}
	}
	return "Unknown"
}
