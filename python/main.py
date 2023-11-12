import sys
import os
sys.path.insert(0, os.getcwd() + '/ten')

import ten

def main():
    ten.diag_diff([[1, 2, 3], [4, 5, 6], [9, 8, 9]])

if __name__ == "__main__":
    main()