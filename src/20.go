import sys

def main():
    while True:
        try:
            n = int(input("Enter a positive integer: "))
            break
        except ValueError:
            print("Invalid input. Please enter a positive integer.")

if __name__ == "__main__":
    main()
