public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World My Old Friend");
        fibSequence x = new fibSequence();
        for (int i = 0; i < 10; i++) {
            System.out.println(x.getNext());
        }
    }

}
