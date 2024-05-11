package main

import (
	"testing"
)

func TestLerSitesDoArquivoComSites(t *testing.T) {
	sites := LerSitesDoArquivo("sites.txt")

	if len(sites) == 0 {
		t.Error("Falha no teste de leitura de sites do arquivo")
	}

	for _, site := range sites {
		if site == "" {
			t.Error("Falha no teste de validação dos sites lidos")
		}
	}
}

func TestLerSitesDoArquivoSemSites(t *testing.T) {
	arquivoVazio := "arquivoVazio.txt"
	sites := LerSitesDoArquivo(arquivoVazio)

	if len(sites) > 0 {
		t.Error("Falha no teste de leitura de sites do arquivo")
	}

	for _, site := range sites {
		if site != "" {
			t.Error("Falha no teste de validação dos sites lidos")
		}
	}
}

func TestTestaSite(t *testing.T) {
	siteOnline := "http://www.google.com"
	siteOffline := "http://www.siteinexistente.com"

	if !TestaSite(siteOnline) {
		t.Error("Falha no teste de monitorar site, o site deveria estar online")
	}

	if TestaSite(siteOffline) {
		t.Error("Falha no teste de monitorar site, o site deveria estar offline")
	}
}
