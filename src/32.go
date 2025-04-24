1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20

def find_largest_prime_factor(number):
    """
    Find the largest prime factor of a given number.

    :param number: An integer
    :return: The largest prime factor of the number
    """
    if number <= 1:
        return None
    i = 2
    while i * i <= number:
        if number % i:
            i += 1
        else:
            number //= i
    if number > 1:
        return number
    else:
        return None

# Example usage:
print(find_largest_prime_factor(60))  # Output: 5
