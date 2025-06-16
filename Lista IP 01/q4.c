#include <stdio.h>
#include <math.h>

int main(){

    float salario, consumo, custoKW, custoConsumo, Desconto;

    scanf("%f %f", &salario, &consumo);

    custoKW = salario * 7 / 1000;

    custoConsumo = custoKW * consumo;

    Desconto = custoConsumo * 0.9;

    printf("Custo por kW: R$: %.2f\nCusto consumo: R$: %.2f\nCusto com desconto R$: %.2f\n", custoKW, custoConsumo, Desconto);
}