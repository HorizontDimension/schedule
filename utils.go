package schedule

import (
	"fmt"
	"strings"
	"time"
)

// Given a filename such as file.csv, file.md.tar, or file, add a timestamp
// file.csv -> file_2014-02-14.csv
// file.md.tar -> file_2014-02-14.md.tar
// file -> file_2014-02-14
// The timestamp will always be prepended before the first period
func TimestampFilename(t time.Time, filename string) string {
	return TimestampFilenameWith(t, filename, `2006-01-02`)
}

// Same as timestamp filename above but allow custom timestamp layout
func TimestampFilenameWith(t time.Time, filename, layout string) string {
	tokens := strings.SplitN(filename, ".", 2)
	// Careful, tokens[1] may not exist
	if len(tokens) == 2 {
		return fmt.Sprintf("%s_%s.%s", tokens[0], t.Format(layout), tokens[1])
	}
	return fmt.Sprintf("%s_%s", tokens[0], t.Format(layout))
}
