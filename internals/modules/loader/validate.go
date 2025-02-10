package loader

// import (
// 	"fmt"
// 	"strings"
// )

/**
given a URL we should validate it and if it is a valid URL return true
else return false
we should also parse if the url or path is for JSR, NPM or the CDN.
we should also parse if the url or path is for a local file
So when the url or path is for a remote file we should return false

* We should also return the type of the file
- JSR
- NPM
- CDN
- Local
*/

func ValidateURL(url string) {
	// isNPM := strings.HasPrefix(url, "npm:")
	// isJSR := strings.HasPrefix(url, "jsr:")

	println("Validating URL:", url)
	// what file format is it?
}
