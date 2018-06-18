package url

import (
	"math/rand"
	"net/url"
	"time"
)

const (
	tamanho  = 5
	simbolos = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_-+"
)

type Url struct {
	Id      string
	Criacao time.Time
	Destino string
}

type Repositorio interface {
	IdExiste(id string) bool
	BuscarPorId(id string) *Url
	BuscarPorUrl(url string) *Url
	RegistartClick(id string)
	Salvar(url Url) error
}

var repo Repositorio

func BuscarOuCriarNovaUrl(destino string) (
	u *Url,
	nova bool,
	err error,
) {
	if u = repo.BuscarPorUrl(destino); u != nil {
		return u, false, nil
	}

	if _, err = url.ParseRequestURI(destino); err != nil {
		return nil, false, err
	}

	url := Url{gerarId(), time.Now(), destino}
	repo.Salvar(url)
	return &url, true, nil
}

func gerarId() string {
	novoId := func() string {
		id := make([]byte, tamanho, tamanho)
		for i := range id {
			id[i] = simbolos[rand.Intn(len(simbolos))]
		}
		return string(id)
	}

	for {
		if id := novoId(); !repo.IdExiste(id) {
			return id
		}
	}
}

func (r *repositorioMemoria) RegistrarClick(id string) {
	r.clicks[id] += 1
}

func Buscar(id string) *Url {
	return repo.BuscarPorId(id)
}

func ConfigurarRepositorio(r Repositorio) {
	repo = r
}

func NovoRepositorioMemoria() *repositorioMemoria {
	return &repositorioMemoria{
		make(map[string]*Url),
		make(map[string])int,
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
	repo = &repositorioMemoria{make(map[string]*Url)}
}
