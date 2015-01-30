package models


type Profile struct {
  Name    string
  Hobbies []string `xml:"Hobbies>Hobby"`
}