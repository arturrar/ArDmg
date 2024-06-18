package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    var name string
    fmt.Println("Введите имя файла содержащего 1 матрицу:")
    fmt.Scanf("%s\n", &name)

    fileName := name + ".txt"
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Ошибка при открытии файла:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    matrixA := [][]int{}

    for scanner.Scan() {
        line := scanner.Text()
        numbersStr := strings.Fields(line)
        row := []int{}

        for _, numStr := range numbersStr {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                fmt.Println("Ошибка при преобразовании строки в число:", err)
                return
            }
            row = append(row, num)
        }
        matrixA = append(matrixA, row)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка при чтении содержимого файла:", err)
        return
    }

    fmt.Println("Полученная матрица:", matrixA, "\n")

    fmt.Println("Введите имя файла содержащего 2 матрицу:")
    fmt.Scanf("%s\n", &name)

    fileName = name + ".txt"
    file, err = os.Open(fileName)
    if err != nil {
        fmt.Println("Ошибка при открытии файла:", err)
        return
    }
    defer file.Close()

    scanner = bufio.NewScanner(file)
    matrixB := [][]int{}

    for scanner.Scan() {
        line := scanner.Text()
        numbersStr := strings.Fields(line)
        row := []int{}

        for _, numStr := range numbersStr {
            num, err := strconv.Atoi(numStr)
            if err != nil {
                fmt.Println("Ошибка при преобразовании строки в число:", err)
                return
            }
            row = append(row, num)
        }
        matrixB = append(matrixB, row)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Ошибка при чтении содержимого файла:", err)
        return
    }

    fmt.Println("Полученная матрица:", matrixB, "\n")

    fmt.Println("Выберите операцию над матрицами:\n1)Сложение\n2)Вычитание\n3)Умножение")
    fmt.Scanf("%s\n", &name)

    switch name {
    case "1":
    sumMatrix := make([][]int, len(matrixA))
    for i := range sumMatrix {
        sumMatrix[i] = make([]int, len(matrixA[0]))
        for j := range sumMatrix[i] {
            sumMatrix[i][j] = matrixA[i][j] + matrixB[i][j]
        }
    }
    fmt.Println("\nСумма матриц:", sumMatrix, "\n")

    case "2":
    sumMatrix := make([][]int, len(matrixA))
    for i := range sumMatrix {
        sumMatrix[i] = make([]int, len(matrixA[0]))
        for j := range sumMatrix[i] {
            sumMatrix[i][j] = matrixA[i][j] - matrixB[i][j]
        }
    }
    fmt.Println("\nРазница матриц:", sumMatrix, "\n")

    case "3":
            resultMatrix := make([][]int, len(matrixA))
    for i := range resultMatrix {
        resultMatrix[i] = make([]int, len(matrixB[0]))
        for j := range resultMatrix[i] {
            for k := range matrixA[0] {
                resultMatrix[i][j] += matrixA[i][k] * matrixB[k][j]
            }
        }
    }
    fmt.Println("\nПроизведение матриц:", resultMatrix, "\n")

    default:
        fmt.Println("Ошибка")
    }
}
