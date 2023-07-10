up = 0
down = 1
left = 2
right = 3


def initialize_grid():
    matrix = []
    ant_positionx = 5
    ant_positiony = 5
    direction = up
    for row in range(10):
        matrix.append([])
        for col in range(10):
            if row == ant_positiony and col == ant_positionx:
                matrix[row].append("o")
            else:
                matrix[row].append(".")
    return matrix

print(initialize_grid())