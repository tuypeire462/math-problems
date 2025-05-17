import sys
def main():
    # Example: solution to a mathematical problem
    print("Solution to 1+2=3: 5")

if __name__ == "__main__":
    if (sys.version_info[0] < 3):
        from distutils.core import setup
        from glob import iglob
        import os
        print('This code requires Python version >= 3.6')
        sys.exit(1)

    main()
