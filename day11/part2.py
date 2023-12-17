def find_empty_space(space):
    empty_rows = []
    empty_cols = []
    
    # Find the empty rows.
    for i in range(0, len(space)):
        if "#" not in space[i]:
            empty_rows.append(i)
    
    # Find the empty columns.
    for i in range(0, len(space[0])):
        if "#" not in [s[i] for s in space]:
            empty_cols.append(i)
    
    return set(empty_rows), set(empty_cols)


def sum_paths(space, expansion_factor):
    empty_rows, empty_cols = find_empty_space(space)
    
    # Replace the # markers with numbers to keep track of them and figure out
    # where they are.
    galaxies = {}
    galaxy_n = 1
    for row in range(0, len(space)):
        for col in range (0, len(space[0])):
            if space[row][col] == "#":
                space[row][col] = str(galaxy_n)
                galaxies[galaxy_n] =  (row, col)
                galaxy_n += 1

    # Now walk through the full set of galaxies and compute how far apart they are.
    path_lengths = []
    for galaxy, coord in galaxies.items():
        for next_galaxy, next_coord in galaxies.items():
            if next_galaxy <= galaxy:
                continue
            
            if coord[0] < next_coord[0]:
                traversed_rows = [i for i in range(coord[0], next_coord[0])]
            else:
                traversed_rows = [i for i in range(next_coord[0], coord[0])]
            
            if coord[1] < next_coord[1]:
                traversed_cols = [i for i in range(coord[1], next_coord[1])]
            else:
                traversed_cols = [i for i in range(next_coord[1], coord[1])]
            
            # Rather than actually expanding space like part1, use Euclidean distance
            # again but expand it by the number of empty lines we crossed.
            row_dist = abs(next_coord[0] - coord[0])
            crossed_empty_rows = len(empty_rows.intersection(traversed_rows))
            row_dist += (crossed_empty_rows * expansion_factor) - crossed_empty_rows
            
            col_dist = abs(next_coord[1] - coord[1])
            crossed_empty_cols = len(empty_cols.intersection(traversed_cols))
            col_dist += (crossed_empty_cols * expansion_factor) - crossed_empty_cols
    
            dist = row_dist + col_dist
            path_lengths.append(dist)
    
    return sum(path_lengths)


if __name__ == "__main__":
    with open("input.txt") as f:
        space = []
        for line in f:
            space.append([c for c in line.strip()])
    
    print(f"part2 sum of paths: {sum_paths(space, expansion_factor=1000000)}")
