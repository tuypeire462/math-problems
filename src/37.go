def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n ** 0.5) + 1):
        if n % i == 0:
            return False
    return True

def largest_prime_factor(n):
    largest = -1
    while n % 2 == 0:
        n //= 2
        largest = 2
    for i in range(3, int(n ** 0.5) + 1, 2):
        while n % i == 0:
            n //= i
            largest = max(largest, i)
    if n > 2:
        largest = max(largest, n)
    return largest

# Example usage
num = 100
print(f"The largest prime factor of {num} is: {largest_prime_factor(num)}")
