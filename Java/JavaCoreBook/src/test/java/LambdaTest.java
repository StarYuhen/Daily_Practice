import CoreJavaVolumeOne.Factory;
import org.junit.Test;

import java.util.function.BiPredicate;

interface TestLambda {
    void TestSystem(int a);
}

public class LambdaTest {
    // 一个普通的lambda表达式，采用函数方式。即使没有参数也需要()
    @Test
    public void ObjectTest() {
//        Runnable runnable = () -> {
//            for (int i = 0; i < 10; i++) {
//                System.out.println("Hello, world!");
//            }
//        };
//        Thread thread = new Thread(runnable);
//        thread.start();

        // 经过无参数可以判断有参数的情况,并声明函数接口
        TestLambda listener = e -> System.out.println("Test");
        listener.TestSystem(6);


    }

    @Test
    // 创建lambda
    public void InterfaceTest() {
        Factory factory = new Factory();
        // 比对两个factory对象是否相同，对象和实例方法的lambda表达式
        BiPredicate<Factory, Factory> equalsFunction = Factory::equals;
        System.out.println(equalsFunction.test(factory, factory));


    }
}
