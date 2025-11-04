package models

type Book struct {
	Id				int
    Title           string
    Author          string
    AuthorFirst     string
    AuthorLast      string
    Genre           string
    Series          string
    Description     string
    PublishDate     string
    Publisher       string
    EanIsbn         string
    UpcIsbn         string
    Pages           uint16 
    Ddc             string
    CoverStyle      string
    SprayedEdges    bool
    SpecialEd       bool
    FirstEd         bool
    Signed          bool
    Location        string
}
