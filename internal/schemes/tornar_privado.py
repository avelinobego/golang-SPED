#!/usr/bin/env python3

import re
import sys

def tornar_privado(match):
    # Captura o nome da estrutura (primeiro grupo)
    nome_publico = match.group(1)
    # Converte a primeira letra para minúscula
    nome_privado = nome_publico[0].lower() + nome_publico[1:]
    # Retorna a linha com o tipo modificado
    return f"type {nome_privado} struct {{"

def processar_arquivo(caminho):
    # Expressão regular para encontrar a declaração de tipo de struct
    regex = re.compile(r'^type\s+([A-Z]\w+)\s+struct\s*\{', re.MULTILINE)

    # Lê o conteúdo do arquivo
    with open(caminho, 'r', encoding='utf-8') as f:
        conteudo = f.read()

    # Substitui as declarações públicas por privadas
    novo_conteudo = regex.sub(tornar_privado, conteudo)

    # Reescreve o mesmo arquivo com o novo conteúdo
    with open(caminho, 'w', encoding='utf-8') as f:
        f.write(novo_conteudo)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Uso: python script.py arquivo.go")
        sys.exit(1)

    arquivo = sys.argv[1]
    processar_arquivo(arquivo)
    print(f"Arquivo '{arquivo}' reescrito com sucesso.")

