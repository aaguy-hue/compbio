package msps

type MaximalSegmentPair struct {
	DbStart    int     // Start of the MSP string in the database
	DbEnd      int     // End of the MSP string in the database
	QStart     int     // Start of the MSP string in the query sequence
	QEnd       int     // End of the MSP string in the query sequence
	Score      int     // Score of the segment pair
	DbSequence string  // The actual sequence from the database
	QSequence  string  // The actual sequence from the query
	PValue     float32 // P-value of the MSP
}
