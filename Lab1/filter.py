import re

def clean_text_file(input_file, output_file, output_file_no_spaces):
    with open(input_file, 'r', encoding='utf-8') as f:
        text = f.read()

    cleaned_text = re.sub(r'[^a-zA-Zа-яА-Я\s]', '', text).lower()
    cleaned_text = re.sub(r'[\n\t\r]', '', cleaned_text)
    cleaned_text = cleaned_text.replace('ъ', 'ь')

    cleaned_text_no_spaces = cleaned_text.replace(' ', '')

    with open(output_file, 'w', encoding='utf-8') as f:
        f.write(cleaned_text)
    
    with open(output_file_no_spaces, 'w', encoding='utf-8') as f:
        f.write(cleaned_text_no_spaces)
    
    print(f'Files created successfully:\n1. {output_file} (with spaces)\n2. {output_file_no_spaces} (no spaces)')

input_filename = 'Noviy_zavet_utf.txt'  
output_filename = 'Noviy_zavet_f.txt'  
output_filename_no_spaces = 'Noviy_zavet_f_no_spaces.txt'  

clean_text_file(input_filename, output_filename, output_filename_no_spaces)
