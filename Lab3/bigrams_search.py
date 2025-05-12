from collections import Counter
from itertools import permutations
from math import gcd
from gcd import extended_gcd
from comparison import comparison

alphabet = 'абвгдеёжзийклмнопрстуфхцчшщыьэюя'
letter_to_index = {char: idx for idx, char in enumerate(alphabet)}

def bigram_to_number(bigram):
    return 31 * letter_to_index[bigram[0]] + letter_to_index[bigram[1]]

def most_common_bigrams(filepath):
    with open(filepath, 'r', encoding='utf-8') as file:
        text = file.read()

    bigrams = [text[i:i+2] for i in range(len(text)-1)]

    counter = Counter(bigrams)

    return [bigram for bigram, _ in counter.most_common(5)]

# top5 = most_common_bigrams("v8.txt")
# for bigram, count in top5:
#     print(f"{bigram}: {count}")

def find_keys (bigrams):
    keys = set()

    common_bigrams = ['ст', 'но', 'ен', 'то', 'на'] 
    known_nums = [bigram_to_number(b) for b in bigrams]
    common_nums = [bigram_to_number(b) for b in common_bigrams]
    for (x1, x2) in permutations(common_nums, 2):
        dx = (x1 - x2) % 961
        if gcd(dx, 961) != 1:
            continue 

        for (y1, y2) in permutations(known_nums, 2):
            dy = (y1 - y2) % 961
            a_list = comparison(dx, dy)
            if isinstance(a_list, int):
                a_list = [a_list]

            for a in a_list:
                b = (y1 - a * x1) % 961
                keys.add((a, b))

            #######################
            # dy = (y1 - y2) % 961

            # k, inv_dx, _ = extended_gcd(dx, 961)

            # if k != 1:
            #     continue

            # a = comparison(dx, dy)
            # b = (y1 - a * x1) % 961

            # keys.append((a, b))


    print("Possible keys:")
    for a, b in keys:
        print(f"a: {a}, b: {b}")        

    return keys


