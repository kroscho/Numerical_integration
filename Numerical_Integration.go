package main

import (
	"fmt"
	"math"
)

var variant int = 3
var eps float64 = math.Pow(10, -8)

//исходная функция
func f(x float64) float64 {
	if variant == 3 {
		return math.Pow(5, x) - 3
	} else {
		return math.Pow(3, x) - 2 * x + 5
	}
}

//производная исходной функции
func dF(x float64) float64 {
	if variant ==3 {
		return math.Log(5)*math.Pow(5,x)
	} else {
		return math.Log(3) * math.Pow(3,x)-2
	}
}

//Вычисление интеграла на определенном шаге
func partIntegral(N int, a, b, h float64, method, k int) float64 {
	var sum float64 = 0
	switch method {
	//формула трапеций
	case 1:
		{
			for i := 0; i < N; i++ {
				var _a float64 = a + float64(i)*h
				var _b float64 = a + float64(i+1)*h
				sum += (_b - _a) * (f(_a) + f(_b)) / 2.
			}
			k += N / 2
			break
		}
	//модифицированная формула трапеций
	case 2:
		{
			for i := 1; i < N; i++ {
				sum += f(a + float64(i)*h)
			}
			sum = h*((f(a)+f(b))/2.+sum) + (h*h)/12.*(dF(a)-dF(b))
			k += N / 2
			break
		}
	//формула Симпсона
	case 3:
		{
			for i:=0; i<N; i++ {
				var _a float64 = a+float64(i)*h
				var _b float64 = a+float64(i+1)*h
				sum += ((_b-_a)/6)*(f(_a)+4*f((_a+_b)/2)+f(_b))
			}
			k+=N
			break
		}
	//формула Гаусса
	case 4:
		{
			var k, a02, a1 float64
			k=0
			a02 = 5./9
			a1 = 8./9
			for i:=0; i<N; i++ {
				k = a + (2*float64(i) + 1)*h/2
				sum += h/2.*(a02*f(k-h*math.Sqrt(0.6)/2.)+a1*f(k)+a02*f(k+h*math.Sqrt(0.6)/2.))
				k+=3
			}
			break
		}
	}
	return sum
}

func print(N int, h, sum, err, k float64) {
	fmt.Println(N, "  |  ", h, "  |  ", sum, "   |   ", err, "   |   ", k )
}

//вычисление интеграла
func calculateIntegral(a, b float64, method int) {
	var N, count int
	var h, sum, term, err, tet, s0, s1, k float64
	N = 1
	count = 0
	h = 1
	sum = 0
	err = 1
	switch method {
	case 1:
		{
			fmt.Println("Trapezium method:")
			break
		}
	case 2:
		{
			fmt.Println("Modified trapezium method:")
			break
		}
	case 3:
		{
			fmt.Println("Simpson Method:")
			tet = 1./15
			break
		}
	case 4:
		{
			fmt.Println("Gauss method:")
			tet = 1./63
			break
		}
	}
	fmt.Println("N     |    h    |    Integral  | Error estimate |  k   ")

	s0 = partIntegral(N, a, b, h, method, count)
	s1 = partIntegral(N*2, a, b, h, method, count)
	k=0

	for math.Abs(err)>eps {
		term = sum
		h = (b - a)/float64(N)
		sum = partIntegral(N, a, b, h, method, count)
		err = (sum - term) * tet
		//эмпирическая оценка порядка аппроксимации
		k = math.Log((sum-s0)/(s1-s0)-1.)/math.Log(0.5)
		s0=s1
		s1=sum
		print(N, h, sum, err, k)
		N*=2
	}
	if method!=4 {
		count += 2
	}
	fmt.Println("Result: ", sum, "\nNumber of requests: ", count)
}

func main() {
	var a, b float64
	a=1
	b=2
	calculateIntegral(a, b, 1)
	calculateIntegral(a, b, 2)
	calculateIntegral(a, b, 3)
	calculateIntegral(a, b, 4)
}