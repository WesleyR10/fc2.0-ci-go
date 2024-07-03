package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := soma(5, 3)
    if resultado != 8 {
        t.Errorf("soma(5, 3) = %d; esperado 8", resultado)
    }
}

func TestSubtracao(t *testing.T) {
    resultado := subtracao(5, 3)
    if resultado != 2 {
        t.Errorf("subtracao(5, 3) = %d; esperado 2", resultado)
    }
}

func TestMultiplicacao(t *testing.T) {
    resultado := multiplicacao(5, 3)
    if resultado != 15 {
        t.Errorf("multiplicacao(5, 3) = %d; esperado 15", resultado)
    }
}

func TestDivisao(t *testing.T) {
    resultado := divisao(6, 3)
    if resultado != 2 {
        t.Errorf("divisao(6, 3) = %d; esperado 2", resultado)
    }

    // Teste de divisão por zero
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("divisao(5, 0) deveria causar panic")
        }
    }()
    divisao(5, 0)
}