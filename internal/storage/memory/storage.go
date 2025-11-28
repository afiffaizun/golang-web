package memory

import (
	"github.com/afiffaizun/golang-web/internal/material"
	"github.com/afiffaizun/golang-web/internal/note"
)

var materials = []material.Material{
	{ID: 1, Title: "Belajar Go Dasar", Summary: "Tipe data, fungsi, struktur kontrol"},
}

func GetAllMaterials() []material.Material {
	return materials
}

func GetMaterialByID(id int) (material.Material, bool) {
	for _, m := range materials {
		if m.ID == id {
			return m, true
		}
	}
	return material.Material{}, false
}

func AddMaterial(newMat material.Material) material.Material {
	newMat.ID = len(materials) + 1
	materials = append(materials, newMat)
	return newMat
}

var notes = []note.Note{}

func GetNotesByMaterialID(materialID int) []note.Note {
	var result []note.Note
	for _, n := range notes {
		if n.MaterialID == materialID {
			result = append(result, n)
		}
	}
	return result
}

func AddNote(newNote note.Note) note.Note {
	newNote.ID = len(notes) + 1
	notes = append(notes, newNote)
	return newNote
}
