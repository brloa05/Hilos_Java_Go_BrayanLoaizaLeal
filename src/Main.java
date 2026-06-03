import java.util.Scanner;

public class Main {

    public static void main(String[] args) throws InterruptedException {

        Scanner sc = new Scanner(System.in);

        System.out.print("Cantidad de hilos: ");
        int numHilos = sc.nextInt();

        long limite = 50_000_000L;
        long bloque = limite / numHilos;

        Thread[] hilos = new Thread[numHilos];
        long inicioPrograma = System.nanoTime();

        for (int i = 0; i < numHilos; i++) {
            long inicio = i * bloque + 1;
            long fin = (i == numHilos - 1) ? limite : (i + 1) * bloque;
            hilos[i] = new ContadorHilo(inicio, fin, "Hilo-" + (i + 1));
            hilos[i].start();
        }

        for (Thread hilo : hilos) {
            hilo.join();
        }

        long finPrograma = System.nanoTime();
        System.out.printf("%nTiempo total: %.3f seg%n",
                (finPrograma - inicioPrograma) / 1_000_000_000.0);
    }
}