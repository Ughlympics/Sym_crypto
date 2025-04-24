alphabet = ['а', 'б', 'в', 'г', 'д', 'е', 'ж', 'з', 'и', 'й', 'к', 'л', 'м', 'н', 'о', 'п', 'р',
            'с', 'т', 'у', 'ф', 'х', 'ц', 'ч', 'ш', 'щ', 'ъ', 'ы', 'ь', 'э', 'ю', 'я']
m = len(alphabet)

def open_txt(input_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        return ''.join([i.strip() for i in f.readlines()])

def split_text(ciphertext, r):
    blocks = ['' for _ in range(r)]
    for i, char in enumerate(ciphertext):
        blocks[i % r] += char
    return blocks

def fre_symbol(text):
    return [text.count(letter) for letter in alphabet]

def recover_key_by_freq(blocks, main_letter='о'):
    key = ''
    main_index = alphabet.index(main_letter)
    for block in blocks:
        freqs = fre_symbol(block)
        most_freq_index = freqs.index(max(freqs))
        shift = (most_freq_index - main_index) % m
        key += alphabet[shift]
    return key

def recover_key_Mg(blocks):
    p = {'а': 0.07625, 'б': 0.01722, 'в': 0.04346, 'г': 0.01702, 'д': 0.02981, 'е': 0.08367, 'ж': 0.01066, 'з': 0.01715, 
         'и': 0.07300, 'й': 0.01046, 'к': 0.03165, 'л': 0.04980, 'м': 0.03010, 'н': 0.07126, 'о': 0.11518, 'п': 0.02743, 
         'р': 0.04354, 'с': 0.05253, 'т': 0.06379, 'у': 0.02710, 'ф': 0.00155, 'х': 0.00873, 'ц': 0.00394, 'ч': 0.01495, 
         'ш': 0.00762, 'щ': 0.00396, 'ъ': 0.00023, 'ы': 0.02030, 'ь': 0.02043, 'э': 0.00297, 'ю': 0.00550, 'я': 0.01857}
    p_values = [p[char] for char in alphabet]
    
    key = ''
    for block in blocks:
        block_len = len(block)
        block_freq = [block.count(c) / block_len for c in alphabet]
        scores = []
        for g in range(m):
            score = sum(p_values[t] * block_freq[(t + g) % m] for t in range(m))
            scores.append(score)
        best_shift = scores.index(max(scores))
        key += alphabet[best_shift]
    return key

def find_index(text):
    return [alphabet.index(char) for char in text]

def decipher_text(ciphertext, key):
    text_ind = find_index(ciphertext)
    key_ind = find_index(key)
    full_key = (key_ind * ((len(text_ind) // len(key_ind)) + 1))[:len(text_ind)]
    return ''.join([alphabet[(ci - ki) % m] for ci, ki in zip(text_ind, full_key)])

def save_txt(text, filename):
    with open(filename, 'w', encoding='utf-8') as f:
        f.write(text)


ciphertext = open_txt("var_13.txt")
r = 17
blocks = split_text(ciphertext, r)

key_freq = recover_key_by_freq(blocks)
decrypted_by_freq = decipher_text(ciphertext, key_freq)
save_txt(decrypted_by_freq, "decrypted_by_freq.txt")

key_mg = recover_key_Mg(blocks)
decrypted_by_mg = decipher_text(ciphertext, key_mg)
save_txt(decrypted_by_mg, "decrypted_by_mg.txt")

print(f"Frequency method key: {key_freq}")
print(f"method key M(g): {key_mg}")