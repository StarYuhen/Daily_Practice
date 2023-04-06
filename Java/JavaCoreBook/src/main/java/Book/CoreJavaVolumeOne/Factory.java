package Book.CoreJavaVolumeOne;

import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

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
public class Factory implements Cloneable, Comparable {
    // init function
    public static @Nullable Product CreateFactory(String string) {
        if (string.equals("A")) {
            return new ProductA();
        } else if (string.equals("B")) {
            return new ProductB();
        }
        return null;
    }

    public static void main(String[] args) {
        // 同时这个变量也是拥有多态的概念
        // 多态:在程序运行时，根据对象的实际类型来决定调用哪个子类的方法，这就是多态的表现。
        Product factoryA = Factory.CreateFactory("A");
        Product factoryB = Factory.CreateFactory("B");
        factoryA.show();
        factoryB.show();
    }

    @Override
    public int compareTo(@NotNull Object o) {
        return 0;
    }
}



