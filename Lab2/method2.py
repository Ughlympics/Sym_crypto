def read_file(file_path):
    with open(file_path, 'r', encoding='utf-8') as f:
        return ''.join(filter(str.isalpha, f.read().upper()))

def coincidence_index(text, max_r=20):
    n = len(text)
    stats = {}

    for r in range(1, max_r + 1):
        D_r = 0
        for i in range(n - r):
            if text[i] == text[i + r]:
                D_r += 1
        stats[r] = D_r

    return stats


file_path = 'var_13.txt'  
text = read_file(file_path)
coincidence_stats = coincidence_index(text, max_r=30)

for r, D_r in coincidence_stats.items():
    print(f"r = {r}, D_r = {D_r}")

best_r = max(coincidence_stats, key=coincidence_stats.get)
print(f"\nMaximum value found D_r = {coincidence_stats[best_r]} at r = {best_r}")