package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for i := 0; i <= 9; i++ {
			z01.PrintRune(rune(i) + '0')
			if i != 9 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	} else if n == 2 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				if i < j {
					z01.PrintRune(rune(i) + '0')
					z01.PrintRune(rune(j) + '0')
					if i == 8 && j == 9 {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	} else if n == 3 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					if i < j && j < k {
						z01.PrintRune(rune(i) + '0')
						z01.PrintRune(rune(j) + '0')
						z01.PrintRune(rune(k) + '0')
						if i == 7 && j == 8 && k == 9 {
							break
						}
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	} else if n == 4 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						if i < j && j < k && k < l {
							z01.PrintRune(rune(i) + '0')
							z01.PrintRune(rune(j) + '0')
							z01.PrintRune(rune(k) + '0')
							z01.PrintRune(rune(l) + '0')
							if i == 6 && j == 7 && k == 8 && l == 9 {
								break
							}
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	} else if n == 5 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for a := 0; a <= 9; a++ {
							if i < j && j < k && k < l && l < a {
								z01.PrintRune(rune(i) + '0')
								z01.PrintRune(rune(j) + '0')
								z01.PrintRune(rune(k) + '0')
								z01.PrintRune(rune(l) + '0')
								z01.PrintRune(rune(a) + '0')
								if i == 5 && j == 6 && k == 7 && l == 8 && a == 9 {
									break
								}
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	} else if n == 6 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for a := 0; a <= 9; a++ {
							for b := 0; b <= 9; b++ {
								if i < j && j < k && k < l && l < a && a < b {
									z01.PrintRune(rune(i) + '0')
									z01.PrintRune(rune(j) + '0')
									z01.PrintRune(rune(k) + '0')
									z01.PrintRune(rune(l) + '0')
									z01.PrintRune(rune(a) + '0')
									z01.PrintRune(rune(b) + '0')
									if i == 4 && j == 5 && k == 6 && l == 7 && a == 8 && b == 9 {
										break
									}
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	} else if n == 7 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for a := 0; a <= 9; a++ {
							for b := 0; b <= 9; b++ {
								for c := 0; c <= 9; c++ {
									if i < j && j < k && k < l && l < a && a < b && b < c {
										z01.PrintRune(rune(i) + '0')
										z01.PrintRune(rune(j) + '0')
										z01.PrintRune(rune(k) + '0')
										z01.PrintRune(rune(l) + '0')
										z01.PrintRune(rune(a) + '0')
										z01.PrintRune(rune(b) + '0')
										z01.PrintRune(rune(c) + '0')
										if i == 3 && j == 4 && k == 5 && l == 6 && a == 7 && b == 8 && c == 9 {
											break
										}
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 8 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for a := 0; a <= 9; a++ {
							for b := 0; b <= 9; b++ {
								for c := 0; c <= 9; c++ {
									for d := 0; d <= 9; d++ {
										if i < j && j < k && k < l && l < a && a < b && b < c && c < d {
											z01.PrintRune(rune(i) + '0')
											z01.PrintRune(rune(j) + '0')
											z01.PrintRune(rune(k) + '0')
											z01.PrintRune(rune(l) + '0')
											z01.PrintRune(rune(a) + '0')
											z01.PrintRune(rune(b) + '0')
											z01.PrintRune(rune(c) + '0')
											z01.PrintRune(rune(d) + '0')
											if i == 2 && j == 3 && k == 4 && l == 5 && a == 6 && b == 7 && c == 8 && d == 9 {
												break
											}
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 9 {
		for i := 0; i <= 9; i++ {
			for j := 0; j <= 9; j++ {
				for k := 0; k <= 9; k++ {
					for l := 0; l <= 9; l++ {
						for a := 0; a <= 9; a++ {
							for b := 0; b <= 9; b++ {
								for c := 0; c <= 9; c++ {
									for d := 0; d <= 9; d++ {
										for f := 0; f <= 9; f++ {
											if i < j && j < k && k < l && l < a && a < b && b < c && c < d && d < f {
												z01.PrintRune(rune(i) + '0')
												z01.PrintRune(rune(j) + '0')
												z01.PrintRune(rune(k) + '0')
												z01.PrintRune(rune(l) + '0')
												z01.PrintRune(rune(a) + '0')
												z01.PrintRune(rune(b) + '0')
												z01.PrintRune(rune(c) + '0')
												z01.PrintRune(rune(d) + '0')
												z01.PrintRune(rune(f) + '0')
												if i == 1 && j == 2 && k == 3 && l == 4 && a == 5 && b == 6 && c == 7 && d == 8 && f == 9 {
													break
												}
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
