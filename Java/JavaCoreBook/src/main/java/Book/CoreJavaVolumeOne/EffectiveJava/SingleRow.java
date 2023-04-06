package Book.CoreJavaVolumeOne.EffectiveJava;


import java.util.function.Supplier;

// 实现单列模式
public class SingleRow {
    // 实现常量实例，采用饿汉模式
    public static final SingleRow InitSingleRow = new SingleRow();
    // 使用transient实现不允许序列化
    private transient String Name = "StarYuhen";

    // 默认的显式构造方法是共有的
    private SingleRow() {
        this.Name = "Test";
    }

    //     实现静态工厂方法 不使用泛型静态了
    public static SingleRow CreateSingleRow() {
        return InitSingleRow;
    }


    public static void main(String[] args) {
        // 同时使用lambed
        Supplier<SingleRow> supplier = SingleRow::CreateSingleRow;
        System.out.println(supplier.toString());
    }


}
