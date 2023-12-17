def expand(space):
    expanded_space = []
    
    for i in range(0, len(space)):
        if "#" in space[i]:
            expanded_space.append(space[i])
        else:
            expanded_space.append(space[i])
            expanded_space.append(space[i])
    
    return expanded_space


def sum_paths(space):
    # Expand the rows.
    expanded_space = expand(space)

    # Expand the columns first by rotating...
    rotated_space = []
    for i in range(0, len(expanded_space[0])):
        rotated_space.append([s[i] for s in expanded_space])
    # ...then doing the expansion...
    rotated_space = expand(rotated_space)
    # ...and rotating them back.
    expanded_space = []
    for i in range(0, len(rotated_space[0])):
        expanded_space.append([s[i] for s in rotated_space])

    # Replace the # markers with numbers to keep track of them and figure out
    # where they are.
    galaxies = {}
    galaxy_n = 1
    for row in range(0, len(expanded_space)):
        for col in range (0, len(expanded_space[0])):
            if expanded_space[row][col] == "#":
                expanded_space[row][col] = str(galaxy_n)
                galaxies[galaxy_n] =  (row, col)
                galaxy_n += 1

    # Now walk through the full set of galaxies and compute how far apart they are.
    path_lengths = []
    for galaxy, coord in galaxies.items():
        for next_galaxy, next_coord in galaxies.items():
            if next_galaxy <= galaxy:
                continue
        
            # Just do Euclidean distance between the two points.
            dist = abs(next_coord[0] - coord[0]) + abs(next_coord[1] - coord[1])
            path_lengths.append(dist)
    
    return sum(path_lengths)


if __name__ == "__main__":
    with open("input.txt") as f:
        space = []
        for line in f:
            space.append([c for c in line.strip()])
    
    print(f"part1 sum of paths: {sum_paths(space)}")
