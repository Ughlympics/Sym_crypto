from scipy.stats import multinomial

x = [134, 98, 66]
n = sum(x)
p = [214/474, 150/474, 110/474]

prob = multinomial.pmf(x, n=n, p=p)
print(f"Probability of observing {x} in a multinomial distribution with n={n} and p={p}: {prob:.4f}")
# The output will be the probability of observing the given counts in a multinomial distribution with the specified parameters.