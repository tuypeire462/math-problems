def find_sum(start, end):
    total = 0
    while start <= end:
        if start % 2 == 1 and (start + 1) % 3 != 0:
            total += start
        else:
            total += start + (end - start)
        start += 1
    return total

print(find_sum(1, 5))
