from collections import Counter
from itertools import permutations
from itertools import combinations
from math import gcd
from gcd import extended_gcd
from comparison import comparison

alphabet = 'абвгдежзийклмнопрстуфхцчшщьыэюя'
letter_to_index = {char: idx for idx, char in enumerate(alphabet)}

def bigram_to_number(bigram):
    return 31 * letter_to_index[bigram[0]] + letter_to_index[bigram[1]]

def most_common_bigrams(filepath):
    with open(filepath, 'r', encoding='utf-8') as file:
        text = file.read()

    bigrams = [text[i:i+2] for i in range(len(text)-1)]

    counter = Counter(bigrams)

    return [bigram for bigram, _ in counter.most_common(5)]


def find_keys (bigrams):
    keys = set()
    print("The most frequent Bigrams:", bigrams)

    common_bigrams = ['ст', 'но', 'то', 'на', 'ен'] 
    known_nums = [bigram_to_number(b) for b in bigrams]
    common_nums = [bigram_to_number(b) for b in common_bigrams]

    for (y1, y2) in combinations(known_nums, 2):
        dy = (y1 - y2) % 961

        for (x1, x2) in permutations(common_nums, 2):
            dx = (x1 - x2) % 961
            if gcd(dx, 961) != 1:
                continue

            a_list = comparison(dx, dy)
            if not a_list:
                continue
            if isinstance(a_list, int):
                a_list = [a_list]

            for a in a_list:
                b = (y1 - a * x1) % 961
                keys.add((a, b))



    print("Keys:", keys)
    keys = sorted(keys)
    # print("Possible keys:")
    # for key in keys:
    #     print(key)


    return list(keys)


