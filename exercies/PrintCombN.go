package piscine

import (
	"strconv"
)

//PrintCombN prints all possible combinations of n different digits in ascending order.
func PrintCombN(n int) string {
	reponce := ""

	if n == 0 {

		reponce = "\n"

	}

	if n == 1 {

		for a := 0; a <= 9; a++ {

			reponce = reponce + strconv.Itoa(a)

			if a < 9 {
				reponce = reponce + ", "

			}

		}

	}

	if n == 2 {

		for a := 0; a <= 8; a++ {
			for b := a + 1; b <= 9; b++ {

				if b > a {

					reponce = reponce + strconv.Itoa(a)
					reponce = reponce + strconv.Itoa(b)

				}

				if a == 8 && b == 9 {

					reponce = reponce + "\n"

				} else {

					reponce = reponce + ", "
				}
			}

		}

	}

	if n == 3 {

		for a := 0; a <= 7; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {

					if c > b && b > a {

						reponce = reponce + strconv.Itoa(a)
						reponce = reponce + strconv.Itoa(b)
						reponce = reponce + strconv.Itoa(c)

					}

					if a == 7 && b == 8 && c == 9 {

						reponce = reponce + "\n"

					} else {

						reponce = reponce + ", "
					}
				}

			}

		}

	}

	if n == 4 {

		for a := 0; a <= 6; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {

						if d > c && c > b && b > a {

							reponce = reponce + strconv.Itoa(a)
							reponce = reponce + strconv.Itoa(b)
							reponce = reponce + strconv.Itoa(c)
							reponce = reponce + strconv.Itoa(d)

						}

						if a == 6 && b == 7 && c == 8 && d == 9 {

							reponce = reponce + "\n"

						} else {

							reponce = reponce + ", "
						}
					}

				}

			}

		}
	}

	if n == 5 {

		for a := 0; a <= 5; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {

							if e > d && d > c && c > b && b > a {

								reponce = reponce + strconv.Itoa(a)
								reponce = reponce + strconv.Itoa(b)
								reponce = reponce + strconv.Itoa(c)
								reponce = reponce + strconv.Itoa(d)
								reponce = reponce + strconv.Itoa(e)

							}

							if a == 5 && b == 6 && c == 7 && d == 8 && e == 9 {

								reponce = reponce + "\n"

							} else {

								reponce = reponce + ", "
							}

						}

					}

				}

			}

		}
	}

	if n == 6 {

		for a := 0; a <= 4; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {

								if f > e && e > d && d > c && c > b && b > a {

									reponce = reponce + strconv.Itoa(a)
									reponce = reponce + strconv.Itoa(b)
									reponce = reponce + strconv.Itoa(c)
									reponce = reponce + strconv.Itoa(d)
									reponce = reponce + strconv.Itoa(e)
									reponce = reponce + strconv.Itoa(f)

								}

								if a == 4 && b == 5 && c == 6 && d == 7 && e == 8 && f == 9 {

									reponce = reponce + "\n"

								} else {

									reponce = reponce + ", "
								}

							}

						}

					}

				}

			}

		}

	}

	if n == 7 {

		for a := 0; a <= 3; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {

									if g > f && f > e && e > d && d > c && c > b && b > a {

										reponce = reponce + strconv.Itoa(a)
										reponce = reponce + strconv.Itoa(b)
										reponce = reponce + strconv.Itoa(c)
										reponce = reponce + strconv.Itoa(d)
										reponce = reponce + strconv.Itoa(e)
										reponce = reponce + strconv.Itoa(f)
										reponce = reponce + strconv.Itoa(g)

									}

									if a == 3 && b == 4 && c == 5 && d == 6 && e == 7 && f == 8 && g == 9 {

										reponce = reponce + "\n"

									} else {

										reponce = reponce + ", "
									}
								}

							}

						}

					}

				}

			}

		}
	}

	if n == 8 {

		for a := 0; a <= 2; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {
									for h := g + 1; h <= 9; h++ {

										if h > g && g > f && f > e && e > d && d > c && c > b && b > a {

											reponce = reponce + strconv.Itoa(a)
											reponce = reponce + strconv.Itoa(b)
											reponce = reponce + strconv.Itoa(c)
											reponce = reponce + strconv.Itoa(d)
											reponce = reponce + strconv.Itoa(e)
											reponce = reponce + strconv.Itoa(f)
											reponce = reponce + strconv.Itoa(g)
											reponce = reponce + strconv.Itoa(h)

										}

										if a == 2 && b == 3 && c == 4 && d == 5 && e == 6 && f == 7 && g == 8 && h == 9 {

											reponce = reponce + "\n"

										} else {

											reponce = reponce + ", "
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

	if n == 9 {

		for a := 0; a <= 2; a++ {
			for b := a + 1; b <= 9; b++ {
				for c := b + 1; c <= 9; c++ {
					for d := c + 1; d <= 9; d++ {
						for e := d + 1; e <= 9; e++ {
							for f := e + 1; f <= 9; f++ {
								for g := f + 1; g <= 9; g++ {
									for h := g + 1; h <= 9; h++ {
										for i := h + 1; i <= 9; i++ {

											if i > h && h > g && g > f && f > e && e > d && d > c && c > b && b > a {

												reponce = reponce + strconv.Itoa(a)
												reponce = reponce + strconv.Itoa(b)
												reponce = reponce + strconv.Itoa(c)
												reponce = reponce + strconv.Itoa(d)
												reponce = reponce + strconv.Itoa(e)
												reponce = reponce + strconv.Itoa(f)
												reponce = reponce + strconv.Itoa(g)
												reponce = reponce + strconv.Itoa(h)
												reponce = reponce + strconv.Itoa(i)
											}

											if a == 1 && b == 2 && c == 3 && d == 4 && e == 5 && f == 6 && g == 7 && h == 8 && i == 9 {

												reponce = reponce + "\n"

											} else {

												reponce = reponce + ", "
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
	return reponce
}
