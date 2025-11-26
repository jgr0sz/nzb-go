package src

//Retrieves the message ID from a segment.
func MessageID(segment *Segment) string {
	return segment.ID
}

//Retrieves the number from a segment.
func Number(segment *Segment) int {
	return segment.Number
}

//Retrieves the size of the segment in bytes.
func SegmentSize(segment *Segment) int {
	return segment.Bytes
}