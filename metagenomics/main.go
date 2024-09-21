package main

import "fmt"

func main() {
	fmt.Println("Metagenomics")

	// read in samples from file
	fmt.Println("Reading in samples from file...")
	allMaps := ReadSamplesFromDirectory("Data\\2019_Samples")
	fmt.Println("Samples read sucessfully!")

	// generate richness/simpson's maps, write to file
	fmt.Println("Generating richness and simpson's maps...")
	WriteRichnessMapToFile(RichnessMap(allMaps), "Matrices\\RichnessMap_2019.csv")
	WriteSimpsonsMapToFile(SimpsonsMap(allMaps), "Matrices\\SimpsonsMap_2019.csv")
	fmt.Println("Richness and Simpson's maps generated!")

	// generate beta diversity matrices, write to file
	fmt.Println("Generating beta diversity matrices...")
	names, jaccard := BetaDiversityMatrix(allMaps, "Jaccard")
	_, braycurtis := BetaDiversityMatrix(allMaps, "Bray-Curtis")
	WriteBetaDiversityMatrixToFile(jaccard, names, "Matrices\\BrayCurtisBetaDiversityMatrix_2019.csv")
	WriteBetaDiversityMatrixToFile(braycurtis, names, "Matrices\\JaccardBetaDiversityMatrix_2019.csv")
	fmt.Println("Finished generating diversity matrices!")

	// visualize data w/ R
}
