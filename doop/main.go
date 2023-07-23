package main

import (
	"os"
)

func main() {
	vf := true
	vfb := true
	arg := os.Args[1:]
	if len(arg) == 3 {
		if arg[0][0] == '-' {
			if !IsNumeric(arg[0][1:]) {
				vf = false
			}
		} else if !IsNumeric(arg[0]) {
			vf = false
		}
		if arg[2][0] == '-' {
			if !IsNumeric(arg[2][1:]) {
				vfb = false
			}
		} else if !IsNumeric(arg[2]) {
			vfb = false
		}
	}
	if vfb && vf && len(arg) == 3 {
		if arg[1] == "+" || arg[1] == "-" || arg[1] == "/" || arg[1] == "*" || arg[1] == "%" {
			Zzz(Atoi, arg)
		}
	}
}

func Zzz(f func(string) int, a []string) {
	var result int
	var tab [2]int
	vf := false
	var r string
	v := "No division by 0"
	z := "No modulo by 0"
	tab[0] = Atoi(a[0])
	tab[1] = Atoi(a[2])
	if tab[0] > 9223372036854775800 || tab[1] > 9223372036854775800 || tab[0] < -9223372036854775800 || tab[1] < -9223372036854775800 {
		vf = true
		return
	} else if a[1] == "-" {
		result = tab[0] - tab[1]
	} else if a[1] == "+" {
		result = tab[0] + tab[1]
	} else if a[1] == "*" {
		result = tab[0] * tab[1]
	} else if a[1] == "/" {
		if tab[1] == 0 {
			os.Stdout.WriteString(v)
			os.Stdout.WriteString("\n")
			vf = true
		} else {
			result = tab[0] / tab[1]
		}
	} else if a[1] == "%" {
		if tab[1] == 0 {
			os.Stdout.WriteString(z)
			os.Stdout.WriteString("\n")
			vf = true
		} else {
			result = tab[0] % tab[1]
		}
	}
	if !vf {
		r = Convertintstr(result)
		os.Stdout.WriteString(r)
		os.Stdout.WriteString("\n")
	}
}

func Convertintstr(n int) string {
	mod := 0
	var nb int = n
	var nombre []rune
	str := ""
	if n < 0 {
		for i := 0; i > n; i-- {
			if nb/10 == 0 {
				str += "-"
				nb *= -1
				nb += 48
				nombre = append(nombre, rune(nb))
				for j := len(nombre) - 1; j >= 0; j-- {
					str += string(nombre[j])
				}
				return str
			} else {
				mod = nb % 10
				nombre = append(nombre, rune(mod*-1+48))
				nb = (nb - mod) / 10
			}
		}
	} else if n == 0 {
		str += "0"
	} else {
		for i := 0; i < n; i++ {
			if nb/10 == 0 {
				nb += 48
				nombre = append(nombre, rune(nb))
				for j := len(nombre) - 1; j >= 0; j-- {
					str += string(nombre[j])
				}
				return str
			} else {
				mod = nb % 10
				nombre = append(nombre, rune(mod+48))
				nb = (nb - mod) / 10
			}
		}
	}
	return str
}

func Atoi(s string) int {
	o := 0
	c := 0
	a_s := []rune(s)
	if len(s) == 0 {
	} else {
		for _, word := range a_s {
			if word >= 48 && word <= 57 {
				for i := '0'; i < word; i++ {
					c++
				}
				o = o*10 + c
				c = 0
			} else {
				o = 0
				break
			}
		}
		if a_s[0] == 43 {
			for index, word := range a_s {
				if index == 0 {
				} else if word >= 48 && word <= 57 {
					for i := '0'; i < word; i++ {
						c++
					}
					o = o*10 + c
					c = 0
				} else {
					o = 0
					break
				}
			}
		}
		if a_s[0] == 45 {
			for index, word := range a_s {
				if index == 0 {
				} else if word >= 48 && word <= 57 {
					for i := '0'; i < word; i++ {
						c++
					}
					o = o*10 - c
					c = 0
				} else {
					o = 0
					break
				}
			}
		}
	}
	return o
}

func IsNumeric(s string) bool {
	z := []rune(s)
	for i := 0; i < len(s); i++ {
		if z[i] >= '0' && z[i] <= '9' {
		} else {
			return false
		}
	}
	return true
}
