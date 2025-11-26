package src

import "time"

//Function to convert and retrieve the Unix-timestamped File field into UTC format.
func DatePosted(file File) time.Time {
	return time.Unix(file.Date, 0).UTC()
}

//Function to retrieve the poster of a file.
func Poster(file File) string {
	return file.Poster
}

//Function to retrieve the segments of a file.
func Segments(file File) []Segment {
	return file.Segments
}

//Adds up all of the segment byte sizes of a file, returning its overall size.
func FileSize(file File) int {
	var segmentByteSize int
	for _, s := range file.Segments {
		segmentByteSize += s.Bytes
	}
	return segmentByteSize
}
