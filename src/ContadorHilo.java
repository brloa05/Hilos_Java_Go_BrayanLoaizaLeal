class ContadorHilo extends Thread {

    private final long inicio;
    private final long fin;

    public ContadorHilo(long inicio, long fin, String nombre) {
        super(nombre);
        this.inicio = inicio;
        this.fin = fin;
    }

    @Override
    public void run() {
        long tiempoInicio = System.nanoTime();

        for (long i = inicio; i <= fin; i++) {
            System.out.println(getName() + " -> " + i);
        }

        System.out.printf("%s terminó en %.3f seg%n",
                getName(),
                (System.nanoTime() - tiempoInicio) / 1_000_000_000.0);
    }
}
