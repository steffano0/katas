import copy
import random

def initialize_grid():
    matrix = []
    for row in range(10):
        matrix.append([])
        for col in range(10):
            _ = random.random()
            if _ < 0.5:
                matrix[row].append("x")
            else:
                matrix[row].append(".")

    return matrix       

print(initialize_grid())

