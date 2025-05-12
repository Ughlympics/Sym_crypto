from gcd import extended_gcd

def comparison(a, b):
    n = 961
    d, x, y = extended_gcd(a, n)

    if b % d != 0:
        return []

    if d == 1:
        _, inv_a1, _ = extended_gcd(a1, n1)
        x0 = (b1 * inv_a1) % n1
        return x0
    else:
        b1 = b // d
        a1 = a // d
        n1 = n // d
        _, inv_a1, _ = extended_gcd(a1, n1)
        x0 = (b1 * inv_a1) % n1
        solutions = [(x0 + i * n1) % n for i in range(d)]
        return solutions