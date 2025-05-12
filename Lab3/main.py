from gcd import extended_gcd
from collections import Counter
from comparison import comparison
import sys
sys.path.append('E:/Sym_cr/Lab3/')
import bigrams_search
import itertools
from bigrams_search import most_common_bigrams, find_keys

# Часто встречающиеся в русском буквы:
FREQ_RU = {'о', 'е', 'а', 'и', 'н', 'т'}

# Алфавит без ъ и ё
ALPHABET = 'абвгдежзийклмнопрстуфхцчшщьыэюя'

def modinv(a, m):
    g, x, _ = extended_gcd(a, m)
    return x % m if g == 1 else None

# def is_likely_russian(text, top_n=6, threshold=4):
#     freqs = Counter(text)
#     most_common = [char for char, _ in freqs.most_common(top_n)]
#     return len(set(most_common) & FREQ_RU) >= threshold

def is_likely_russian(text, threshold=0.6):
    text = text.lower()
    total_letters = sum(c.isalpha() for c in text)
    
    if total_letters == 0:
        return False

    # Частотные буквы русского языка
    frequent_letters = {'о', 'е', 'а', 'и', 'н', 'т', 'с', 'л', 'в', 'р'}
    letter_counts = Counter(filter(str.isalpha, text))

    # Доля частотных русских букв
    frequent_total = sum(letter_counts[l] for l in frequent_letters if l in letter_counts)
    freq_ratio = frequent_total / total_letters

    # Проверка на мусорные символы
    bad_chars = set("qwertyuiopasdfghjklzxcvbnm0123456789!@#$%^&*<>/\\")
    if any(c in bad_chars for c in text):
        return False

    # Простой порог: если доля частотных букв > threshold (по умолчанию 50%)
    return freq_ratio > threshold

def decrypt_bigram(text, key):
    a, b = key
    m = len(ALPHABET) ** 2
    a_inv = modinv(a, m)
    if a_inv is None:
        return None

    bigram_to_index = {
        a1 + a2: i * len(ALPHABET) + j
        for i, a1 in enumerate(ALPHABET)
        for j, a2 in enumerate(ALPHABET)
    }
    index_to_bigram = {v: k for k, v in bigram_to_index.items()}

    result = []
    for i in range(0, len(text) - 1, 2):
        bg = text[i:i+2]
        if bg in bigram_to_index:
            y = bigram_to_index[bg]
            x = (a_inv * (y - b)) % m
            result.append(index_to_bigram[x])
        else:
            result.append(bg)
    return ''.join(result)





def main():
    input_file = "16.txt"
    
    # Читаем текст
    with open(input_file, encoding="utf-8") as f:
        cipher_text = f.read().replace("\n", "").lower()

    
    cnt = most_common_bigrams("16.txt")
    print("Top 5 most common bigrams:")
    for bigram in cnt:
        print(f"{bigram}")

    k = find_keys(cnt)
    print("Possible keys:") 

    for key in k:
        decrypted = decrypt_bigram(cipher_text, key)
        if decrypted and is_likely_russian(decrypted):
            print(f"Ключ знайдено: {key}")
            with open("result.txt", "w", encoding="utf-8") as out:
                out.write(decrypted)
            break
    else:
        print("Не знайдено ключ.")

if __name__ == "__main__":
    main()
    