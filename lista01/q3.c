#include <stdio.h>
#include <math.h>

int main(){
    int centenas, dezenas, unidades, numero;
    int quadrado;
    int i = 0;
    while( i == 0 ){

        int q = 0;

        scanf("%d", &centenas);
        scanf("%d", &dezenas);
        scanf("%d", &unidades);

        if( centenas >= 0 && centenas <= 9){
            q++;
        }
        if ( dezenas >= 0 && dezenas <= 9){
            q++;
        }
        if ( unidades >= 0 && unidades <= 9){
            q++;
        }
        if( q == 3){
            i++;
        } else {
            printf("DIGITO INVALIDO");
            return 0;
        }
    }
    
    numero = 100 * centenas + 10 * dezenas + unidades;

    quadrado = pow(numero, 2);
     
    printf("%d%d%d, %d\n", centenas, dezenas, unidades, quadrado );
}