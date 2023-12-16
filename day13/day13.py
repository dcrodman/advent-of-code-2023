def part1(game):
    # Check for horizontal mirrors.
    rows = find_mirror(game)
    
    # Now check for vertical mirrors by rotating the entire game
    # repeating the same algorithm.
    rotated_game = []
    for i in range(0, len(game[0])):
        rotated_game.append([row[i] for row in game])
    cols = find_mirror(rotated_game)
    
    return rows, cols


def find_mirror(game):
    for i in range(0, len(game)-1):
        # Did we find the start of the reflection?
        if game[i] == game[i+1]:
            # Walk to either end of the grid and as long as the reflections
            # match up until that point then we have the mirror.
            found = True
            for j in range(1, len(game)-1-i):
                if i-j < 0:
                    break
                if game[i-j] != game[i+1+j]:
                    found = False
                    break
            if found:
                return i+1
    return 0


def part2(game, mirror_row, mirror_col):
    # Check for horizontal mirrors.
    rows = find_mirror_smudges(game, mirror_row)
    
    # Now check for vertical mirrors by rotating the entire game
    # repeating the same algorithm.
    rotated_game = []
    for i in range(0, len(game[0])):
        rotated_game.append([row[i] for row in game])
    cols = find_mirror_smudges(rotated_game, mirror_col)
    
    return rows, cols


def find_mirror_smudges(game, known_mirror):
    for i in range(0, len(game)-1):
        # Skip the mirror we found in part1 to find the smudge.
        if i == known_mirror:
            continue
        
        d = diff(game[i], game[i+1])
        if d <= 1:
            found = True
            for j in range(1, len(game)-1-i):
                if i-j < 0:
                    break
                
                d += diff(game[i-j], game[i+1+j])
                # We only expect one smudge, so if we've exceeded that then bail.
                if d > 1:
                    found = False
                    break
            if found:
                return i+1
    return 0


def diff(row1, row2):
    d = 0
    for i in range(len(row1)):
        if row1[i] != row2[i]:
            d += 1
    return d


def parse(filename):
    with open(filename) as f:
        games = []
        game = []
        for line in f:    
            if len(line) == 1:
                games.append(game)
                game = []
                row = []
            else:
                row = []
                for c in line.strip():
                    row.append(c)
                game.append(row)
        games.append(game)
        
    return games


if __name__ == "__main__":
    games = parse("input.txt")
    
    p1_total_rows, p1_total_cols = 0, 0
    p2_total_rows, p2_total_cols = 0, 0
    for i in range(0, len(games)):
        game = games[i]

        row, col = part1(game)
        p1_total_rows += row
        p1_total_cols += col
        
        row, col = part2(game, row-1, col-1)
        p2_total_rows += row
        p2_total_cols += col
        
        
    print(f"part1 answer: {p1_total_cols + (p1_total_rows * 100)}")
    print(f"part2 answer: {p2_total_cols + (p2_total_rows * 100)}") 
