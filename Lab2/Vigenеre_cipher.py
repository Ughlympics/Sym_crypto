import sys

def vigenere_encrypt(text, key):
    alphabet = "абвгдежзийклмнопрстуфхцчшщьыъэюя"
    text = text.lower().replace("ё", "е")
    key = key.lower().replace("ё", "е")
    encrypted_text = ""
    key_index = 0

    for char in text:
        if char in alphabet:
            shift = alphabet.index(key[key_index % len(key)])
            new_letter = alphabet[(alphabet.index(char) + shift) % len(alphabet)]
            encrypted_text += new_letter
            key_index += 1
        else:
            encrypted_text += char
    
    return encrypted_text

def process_file(input_filename, output_filename, key):
    try:
        with open(input_filename, 'r', encoding='utf-8') as infile:
            text = infile.read()
        encrypted_text = vigenere_encrypt(text, key)
        with open(output_filename, 'w', encoding='utf-8') as outfile:
            outfile.write(encrypted_text)
        print(f"The file has been successfully encrypted and saved in {output_filename}")
    except FileNotFoundError:
        print("Error: Input file not found.")

if __name__ == "__main__":
    input_file = "random_text.txt"
    output_file = "random_text_20.txt"
    key = "божехраникороляиегон"
    process_file(input_file, output_file, key)

#шифрувалося повідомленням: божехраникороляиегонарод