def fibonacci(n):
    if n <= 0:
        return "Invalid input"
    elif n == 1:
        return [0]
    elif n == 2:
        return [0, 1]

    sequence = fibonacci(n - 1)
    last_two_elements = sequence[-1] + sequence[-2]
    sequence.append(last_two_elements)

    return sequence

# Example usage
print(fibonacci(10))
