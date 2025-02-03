package main

import (
	"fmt"
)

func SARSAlignment() {
	fmt.Println("Reading in SARS and SARS-CoV-2 genomes.")
	sars1 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV_genome.fasta")
	sars2 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV-2_genome.fasta")

	match := 1.0
	mismatch := 1.0
	gap := 3.0

	sarsAlignment := GlobalAlignment(sars1, sars2, match, mismatch, gap)
	fmt.Println("Alignment run, now writing to file.")

	WriteAlignmentToFASTA(sarsAlignment, "Output/sarsAlignment.txt")
	fmt.Println("Alignment written to file. Exiting...")
}

func HemoglobinAlignment() {
	zebrafish := ReadFASTAFile("Data/Hemoglobin/Danio_rerio_hemoglobin.fasta")
	cow := ReadFASTAFile("Data/Hemoglobin/Bos_taurus_hemoglobin.fasta")
	gorilla := ReadFASTAFile("Data/Hemoglobin/Gorilla_gorilla_hemoglobin.fasta")
	human := ReadFASTAFile("Data/Hemoglobin/Homo_sapiens_hemoglobin.fasta")

	// if mismatch > gap*2, then we will never have any mismatches since you can just go left and up rather than going diagonal, which doesn't make sense for a mutation
	// inserting or deleting is more harmful to the organism than a mismatch which may have a very similar chemical signature, so it is less likely, therefore we penalize gaps more
	match := 1.0
	mismatch := 1.0
	gap := 3.0

	alignment_1 := GlobalAlignment(zebrafish, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment_1, "Output/human_zebrafish_hemoglobin.txt")

	alignment_2 := GlobalAlignment(cow, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment_2, "Output/human_cow_hemoglobin.txt")

	alignment_3 := GlobalAlignment(gorilla, human, match, mismatch, gap)
	WriteAlignmentToFASTA(alignment_3, "Output/gorilla_human_hemoglobin.txt")
}

func SARSSpikeLocalAlignment() {
	fmt.Println("Reading in SARS spike protein and SARS-CoV-2 genome.")
	sars1 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV_genome_spike_protein.fasta")
	sars2 := ReadFASTAFile("Data/Coronaviruses/SARS-CoV-2_genome.fasta")

	match := 1.0
	mismatch := 1.0
	gap := 1.0

	spikeAlignment, start1, end1, start2, end2 := LocalAlignment(sars1, sars2, match, mismatch, gap)
	fmt.Println("Alignment run, now writing to file.")

	WriteLocalAlignmentToFASTA(spikeAlignment, "Output/sarsSpikeAlignment.txt", start1, end1, start2, end2)
	fmt.Println("Alignment written to file. Exiting...")
}

func main() {
	// fmt.Println("Sequence alignment!")

	// HemoglobinAlignment()
	// SARSAlignment()
	// SARSSpikeLocalAlignment()

	// maxScore, align1, align2 := AffineAlignment(1, 3, 2, 1, "GA", "GTTA")
	// maxScore, align1, align2 := AffineAlignment(1, 5, 5, 1, "GAT", "AT")
	// maxScore, align1, align2 := AffineAlignment(1, 2, 3, 2, "CAGGT", "TAC")
	// maxScore, align1, align2 := AffineAlignment(2, 3, 3, 2, "GTTCCAGGTA", "CAGTAGTCGT")
	// fmt.Println(maxScore)
	// fmt.Println(align1)
	// fmt.Println(align2)

	n, m := 10, 8
	edges := []Edge{
		{Coordinate{1, 1}, Coordinate{1, 2}, 1},
		{Coordinate{1, 3}, Coordinate{1, 4}, 1},
		{Coordinate{2, 3}, Coordinate{2, 4}, 1},
		{Coordinate{2, 2}, Coordinate{3, 2}, 1},
		{Coordinate{3, 5}, Coordinate{3, 6}, 1},
		{Coordinate{4, 2}, Coordinate{5, 2}, 1},
		{Coordinate{4, 3}, Coordinate{4, 4}, 1},
		{Coordinate{4, 5}, Coordinate{4, 6}, 1},
		{Coordinate{5, 2}, Coordinate{5, 3}, 1},
		{Coordinate{6, 0}, Coordinate{6, 1}, 1},
		{Coordinate{8, 0}, Coordinate{8, 1}, 2},
		{Coordinate{9, 2}, Coordinate{9, 3}, 1},
		{Coordinate{9, 4}, Coordinate{9, 5}, 1},
		{Coordinate{9, 6}, Coordinate{9, 7}, 1},
	}
	path := ManhattanTourist(n, m, edges)

	DisplayPathAscii(n, m, path)
}
