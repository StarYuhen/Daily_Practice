package CoreJavaVolumeOne;

// definition java interface product Println
interface Product {
    void show();
}

enum ShowType {
    Name, Age, Test
}


// definition Java ,is factory model product
class ProductA implements Product {
    // log productA
    @Override
    public void show() {
        System.out.println("ProductA");
    }
}

class ProductB implements Product {
    // log productB
    @Override
    public void show() {
        System.out.println("ProductB");
    }
}


// factory model
public class Factory {
    // init function
    public static Product CreateFactory(String string) {
        if (string.equals("A")) {
            return new ProductA();
        } else if (string.equals("B")) {
            return new ProductB();
        }
        return null;
    }

    public static void main(String[] args) {
        Product factoryA = Factory.CreateFactory("A");
        Product factoryB = Factory.CreateFactory("B");
        factoryA.show();
        factoryB.show();
    }

}



