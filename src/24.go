def calculate_area(shape):
    if shape == "rectangle":
        length = float(input("Enter the length: "))
        width = float(input("Enter the width: "))
        area = length * width
        print(f"The area of the rectangle is: {area}")
    elif shape == "circle":
        radius = float(input("Enter the radius: "))
        area = 3.14 * (radius ** 2)
        print(f"The area of the circle is: {area:.2f}")
    else:
        print("Invalid shape.")
