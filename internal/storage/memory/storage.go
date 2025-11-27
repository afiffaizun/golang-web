package memory

import "github.com/afiffaizun/golang-web/internal/material"

var materials = []material.Material{
	{ID: 1, Title: "Belajar Go Dasar", Summary: "Tipe data, fungsi, struktur kontrol"},
}

func GetAllMaterials() []material.Material {
	return materials
}

func GetMaterialByID(id int) (material.Material, bool) {
	for _, m := range materials{
		if m.ID == id {
			return m, true
		}
	}
	return material.Material{}, false
}
