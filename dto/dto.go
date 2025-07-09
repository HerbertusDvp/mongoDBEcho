package dto

type CategoriaDto struct {
	Nombre string `json:"nombre"`
}

type GenericoDto struct {
	Estado  string `json:"estado"`
	Mensaje string `json:"mensaje"`
}

type ProductoDto struct {
	Nombre      string `json:"nombre"`
	Precio      int    `json:"precio"`
	Stock       int    `json:"stock"`
	Descripcion string `json:"descripcion"`
	CategoriaID string `json:"categoria_id"`
}

type UsuarioDto struct {
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Telefono string `json:"telefono"`
	Password string `json:"password"`
}
