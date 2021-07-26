package models

import "aluraflix-api/db"

type Video struct {
	ID        int
	Titulo    string
	Descricao string
	URL       string
}

func ListarTodosOsVideos() []Video {
	db := db.ConectaComBancoDeDados()

	query, err := db.Query("select * from videos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	v := Video{}
	videos := []Video{}

	for query.Next() {
		var id int
		var titulo, descricao, url string

		err = query.Scan(&id, &titulo, &descricao, &url)
		if err != nil {
			panic(err.Error())
		}

		v.ID = id
		v.Titulo = titulo
		v.Descricao = descricao
		v.URL = url

		videos = append(videos, v)
	}

	defer db.Close()
	return videos
}
