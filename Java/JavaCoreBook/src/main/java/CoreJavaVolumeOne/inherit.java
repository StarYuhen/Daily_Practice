package CoreJavaVolumeOne;

/**
 * 用于测试Java的继承操作
 * 继承前面的深拷贝类，同时测试使用字段和方法
 *
 * @author StarYuhen
 * @version 1.0.1
 */
public class inherit extends DeepCopy {
    // 自定义获取父类字段方法，这是错误的,单纯获取不行
//    public int getAge() {
//        return 0;
//    }

    // 正确方法使用，super属性获取方法
    public int getAge() {
        System.out.println(super.getAge());
        return 0;
    }

    // 不可以直接使用父类的getAge，会循环崩溃,因为会一直调用自身
//    public int getAge() {
//        System.out.println(getAge());
//        return 0;
//    }
}
