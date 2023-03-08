package Book.CoreJavaVolumeOne;

import java.util.ArrayList;

// 默认超类 Object
public class SuperClass {
    // 使用保护标志属性,指名此方法只有自己和子类可以使用
    protected String name = "StarYuhen";
    // 测试equals
    private final static DeepCopy DeepCopyPrivate = new DeepCopy();

    // 测试equals超类
    public void Compare(DeepCopy deepCopy) {
        if (deepCopy.equals(DeepCopyPrivate)) System.out.println("不同对象的散列码不可能一样");
        // 同时输出一下散列码
        /*
         散列码是根据超类的方法计算的，同时使用内存地址

         而equals必须具有以下特性：一致性，传递性，自反性，对称性，非空性
            其中非空性指的是这个方法即使是null也会是安全的，也就是自行判断过

         比对散列码是不行的，因为在Java中有着哈希冲突的概念，即使在不同的对象，不同的类有可能是一样的散列码

         同时每一个对象都有一个散列码，每一个类都有一个虚拟表(VMT)用于重载解析
         */
        System.out.println(deepCopy.hashCode());
        System.out.println(DeepCopyPrivate.hashCode());

    }


    // 设置包装器，其中的自动拆箱和自动装箱
    /*
    // 注意空指针和性能问题
    同时使用包装类可以使用==比对了，并且可以优化部分代码
     */
    public void Wrapper() {
        // 使用包装类
        Integer num = 10;
        Boolean flash = true;
        // 自动拆箱
        int num1 = num;
        boolean flash1 = flash;
        // 自动装箱,将基本类型转化为基础类
        Integer num2 = num1 + 10;
        Boolean flash2 = flash1 && false;

    }

    // 使用泛型数组
    @SuppressWarnings("uncheckd") // 使用这个标记，可以告诉编译器忽略错误，可以用于强制类型转换
    public void Array() {
        // 初始化数组,可以省略右边的类型，同时简易使用var
        // 使用初始容器长度，或者不需要，但是不设定长度会消耗性能
        ArrayList<DeepCopy> deepCopies = new ArrayList<>();
        System.out.println(deepCopies.getClass().getName() + "@" + Integer.toHexString(hashCode()));
        System.out.printf("现在数组的长度:%s \n", deepCopies.size());

        DeepCopy deepCopy = new DeepCopy();
        for (int i = 0; i <= 95; i++) {
            deepCopies.add(deepCopy);
        }
        System.out.printf("增加91个数据后数组的长度:%s \n", deepCopies.size());
        // 使用函数在运行时增加限制,100个的长度
        deepCopies.ensureCapacity(100);

        // 同时使用函数回收数组长度,设置为现在占用的长度，
        deepCopies.trimToSize();

        deepCopies.remove(55);

        // 强制Gc，观看代码情况 trimToSize 此函数只会在gc后回收生效
        System.gc();
        System.out.printf("现在数组的长度:%s \n", deepCopies.size());

        /*
         结果：
            现在数组的长度:0
            增加91个数据后数组的长度:96
            现在数组的长度:95

            会发现回收了，就这样
         */

    }

    // 新增一个参数可变的方法
    public void TestCell(String string, String... args) {
        // 这样可以接收多个String类型的参数，String... args 可变Object... args
    }


    public static void main(String[] args) {
        DeepCopy deepCopy = new DeepCopy();
        SuperClass superClass = new SuperClass();
        superClass.Compare(deepCopy);
        superClass.Array();
    }
}
