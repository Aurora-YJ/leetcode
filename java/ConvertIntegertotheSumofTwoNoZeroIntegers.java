public class ConvertIntegertotheSumofTwoNoZeroIntegers {
    
    public  int[] getNoZeroIntegers(int n) {
        
        for (int i = 1; i <= n; i++) {
            String j = Integer.toString(i);
            String j1 = Integer.toString(n-i);
            if (!j.contains("0") && !j1.contains("0")) {
                return new int[]{i,n-i};
            }
        }
        return new int[]{};
    }
    public static void main(String[] args) {
        ConvertIntegertotheSumofTwoNoZeroIntegers h = new ConvertIntegertotheSumofTwoNoZeroIntegers();
        int[] t = h.getNoZeroIntegers(7);

        if (t.length == 0) {
            System.out.print("");
            return;
        }
        System.out.println(  "[" + t[0]  + "," + t[1] + "]");
    }
}


