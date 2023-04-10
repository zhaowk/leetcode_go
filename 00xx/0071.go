package _0xx

func simplifyPath(path string) string {

	lp := len(path)

	if lp < 2 {
		return path
	}

	paths := make([]string, 0)

	start, end := 1, 0

	for i, v := range path {
		if v == '/' {
			end = i
			if end < start {
				continue
			}
			if end == start+1 {
				if path[start:end] != "." {
					paths = append(paths, path[start:end])
				}
			} else if end == start+2 {
				if path[start:end] != ".." {
					paths = append(paths, path[start:end])
				} else {
					if len(paths) > 0 {
						paths = paths[:len(paths)-1]
					}
				}
			} else if end != start {
				paths = append(paths, path[start:end])
			}
			start = end + 1
		}
	}

	if end != len(path)-1 {
		if path[start:] == ".." {
			if len(paths) > 0 {
				paths = paths[0 : len(paths)-1]
			}
		} else if path[start:len(path)] == "." {
			//
		} else {
			paths = append(paths, path[start:])
		}
	}

	if len(paths) == 0 {
		return "/"
	}

	tmp := ""
	for _, p := range paths {
		tmp += "/" + p
	}
	return tmp
}
