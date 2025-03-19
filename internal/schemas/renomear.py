#!/usr/bin/env python3

import os
import re
import argparse

def camel_to_snake(name: str) -> str:
    """Converte um nome de arquivo de CamelCase para snake_case."""
    name = re.sub(r'(?<!^)(?=[A-Z])', '_', name).lower()
    return name

def rename_files_in_directory(directory: str, extension: str):
    """Renomeia arquivos em CamelCase para snake_case em um diretório específico."""
    for filename in os.listdir(directory):
        if filename.endswith(extension):
            name, ext = os.path.splitext(filename)
            new_name = camel_to_snake(name) + ext
            old_path = os.path.join(directory, filename)
            new_path = os.path.join(directory, new_name)
            os.rename(old_path, new_path)
            print(f'Renomeado: {filename} -> {new_name}')

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Renomeia arquivos de CamelCase para snake_case.")
    parser.add_argument("directory", type=str, help="Diretório onde os arquivos estão localizados.")
    parser.add_argument("extension", type=str, help="Extensão dos arquivos a serem renomeados, incluindo o ponto (ex: .go).")
    args = parser.parse_args()
    
    rename_files_in_directory(args.directory, args.extension)
