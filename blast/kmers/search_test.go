package kmers

import (
	"reflect"
	"testing"
)

func TestSearchKmersInDb(t *testing.T) {
	kmers := []string{"ACT", "TGA"}
	db := "ACTTGAACT"
	expected := []KmerSearchResult{
		{Kmer: "ACT", DbPosition: 0},
		{Kmer: "TGA", DbPosition: 3},
		{Kmer: "ACT", DbPosition: 6},
	}

	got := SearchKmersInDb(kmers, db)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, got)
	}
}
