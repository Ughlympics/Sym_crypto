import re
import math
import collections
import xlsxwriter

def compute_entropy(freq_dict, total_count):
    entropy = -sum((count / total_count) * math.log2(count / total_count) for count in freq_dict.values())
    return entropy

def process_text_file(input_file, output_file):
    with open(input_file, 'r', encoding='utf-8') as f:
        text = f.read().lower()
    
    letter_freq = collections.Counter(text)
    bigram_freq = collections.Counter(text[i:i+2] for i in range(len(text)-1))
    total_letters = sum(letter_freq.values())
    total_bigrams = sum(bigram_freq.values())
    
    H1 = compute_entropy(letter_freq, total_letters)
    H2 = compute_entropy(bigram_freq, total_bigrams) / 2
    
    workbook = xlsxwriter.Workbook(output_file)
    worksheet = workbook.add_worksheet()
    row = 0
    
    worksheet.write(row, 0, 'Буква')
    worksheet.write(row, 1, 'Частота')
    row += 1
    for letter, count in letter_freq.items():
        worksheet.write(row, 0, repr(letter)) 
        worksheet.write(row, 1, count / total_letters)
        row += 1
    
    row += 1
    worksheet.write(row, 0, 'Биграмма')
    worksheet.write(row, 1, 'Частота')
    row += 1
    for bigram, count in bigram_freq.items():
        worksheet.write(row, 0, repr(bigram)) 
        worksheet.write(row, 1, count / total_bigrams)
        row += 1
    
    row += 1
    worksheet.write(row, 0, 'H1')
    worksheet.write(row, 1, H1)
    row += 1
    worksheet.write(row, 0, 'H2')
    worksheet.write(row, 1, H2)
    
    workbook.close()
    print(f'File {output_file} created successfully!')

input_filename = 'Noviy_zavet_f.txt'
output_filename = 'frequency.xlsx'
process_text_file(input_filename, output_filename)