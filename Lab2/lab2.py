import os
from collections import Counter

def index_of_coincidence(text):
    alphabet = "абвгдежзийклмнопрстуфхцчшщьыъэюя"
    text = ''.join([char for char in text if char in alphabet])  
    
    if len(text) == 0:
        return 0.0
    
    freq = Counter(text)
    n = len(text)
    ic = sum(f * (f - 1) for f in freq.values()) / (n * (n - 1))
    return ic

def process_directory(directory):
    for filename in os.listdir(directory):
        if filename.endswith(".txt"):
            filepath = os.path.join(directory, filename)
            try:
                with open(filepath, 'r', encoding='utf-8') as file:
                    text = file.read()
                ic = index_of_coincidence(text)
                print(f"Index for '{filename}': {ic:.4f}")
            except FileNotFoundError:
                print(f"Error: file {filename} ot found.")

if __name__ == "__main__":
    directory = "E:\Sym_cr\Lab2"
    process_directory(directory)