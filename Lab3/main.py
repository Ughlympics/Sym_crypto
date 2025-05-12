from gcd import extended_gcd
from collections import Counter
from comparison import comparison
import sys
sys.path.append('E:/Sym_cr/Lab3/')
import bigrams_search
import itertools
from bigrams_search import most_common_bigrams, find_keys





def main():
    cnt = most_common_bigrams("v8.txt")
    print("Top 5 most common bigrams:")
    for bigram in cnt:
        print(f"{bigram}")

    k = find_keys(cnt)
    print("Possible keys:") 

if __name__ == "__main__":
    main()
    