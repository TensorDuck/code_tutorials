public class fibSequence {
    private int a = 0, b = 0, total = 0;

    public fibSequence() {
    }

    public int getNext() {
        total = a + b;
        if (a == 0 & b == 0) {
            a = 1;
        } else {
            a = b;
            b = total;
        }
        return total;
    }
}