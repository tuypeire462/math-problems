def check_squares(a, b):
    """
    This function checks if two 2D arrays are equal by comparing corresponding elements.
    It returns True if they are equal and False otherwise.
    
    Args:
        a (list of lists): The first 2D array.
        b (list of lists): The second 2D array.
        
    Returns:
        bool: True if the two arrays are equal, False otherwise.
    """
    for i in range(len(a)):
        if len(a[i]) != len(b[i]):
            return False
        for j in range(len(a[0])):
            if a[i][j] != b[i][j]:
                return False
    return True

# Example usage:
matrix1 = [
    [1, 2, 3],
    [4, 5, 6]
]

matrix2 = [
    [1, 2, 3],
    [4, 5, 7]
]

print(check_squares(matrix1, matrix2))  # Output: True
