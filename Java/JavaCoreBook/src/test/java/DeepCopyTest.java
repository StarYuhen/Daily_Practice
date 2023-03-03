import CoreJavaVolumeOne.DeepCopy;
import org.junit.Test;

public class DeepCopyTest {
    // test deep copy memory address
    @Test
    public void TestDeepCopy() {
        DeepCopy deep1 = new DeepCopy();
        // Quote DeepCopy
        DeepCopy deep2 = deep1;
        deep2.setAge(90);
        System.out.println(deep1 == deep2);
        // 注意，这里获取的是deep1的值，而我设置的是deep2的值，这就是浅拷贝，改变属性会改变源头引用
        System.out.println(deep1.getAge());
    }

    @Test
    public void TestShallow() throws CloneNotSupportedException {
        DeepCopy deep1 = new DeepCopy();
        DeepCopy deep2 = deep1.clone();
        // DeepCopy deep3 = new DeepCopy();
        System.out.println(deep1 == deep2);
        // update value
        deep2.setAge(90);
        System.out.println(deep1.getAge());
    }
//    // test Parameter Quote
//    public void TestParameter(){
//        DeepCopy deepCopy=new DeepCopy();
//        deepCopy.
//    }
}
